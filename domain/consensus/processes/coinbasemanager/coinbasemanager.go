package coinbasemanager

import (
	"math"

	"github.com/Nirvana-Chain/nirvanad/domain/consensus/model"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/model/externalapi"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/constants"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/hashset"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/subnetworks"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/transactionhelper"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/db/database"
	"github.com/pkg/errors"
)

type coinbaseManager struct {
	subsidyGenesisReward                    uint64
	preDeflationaryPhaseBaseSubsidy         uint64
	coinbasePayloadScriptPublicKeyMaxLength uint8
	genesisHash                             *externalapi.DomainHash
	deflationaryPhaseDaaScore               uint64
	deflationaryPhaseBaseSubsidy            uint64

	databaseContext     model.DBReader
	dagTraversalManager model.DAGTraversalManager
	ghostdagDataStore   model.GHOSTDAGDataStore
	acceptanceDataStore model.AcceptanceDataStore
	daaBlocksStore      model.DAABlocksStore
	blockStore          model.BlockStore
	pruningStore        model.PruningStore
	blockHeaderStore    model.BlockHeaderStore
}

func (c *coinbaseManager) ExpectedCoinbaseTransaction(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	coinbaseData *externalapi.DomainCoinbaseData) (expectedTransaction *externalapi.DomainTransaction, hasRedReward bool, err error) {

	ghostdagData, err := c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, true)
	if !database.IsNotFoundError(err) && err != nil {
		return nil, false, err
	}

	// If there's ghostdag data with trusted data we prefer it because we need the original merge set non-pruned merge set.
	if database.IsNotFoundError(err) {
		ghostdagData, err = c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, false)
		if err != nil {
			return nil, false, err
		}
	}

	acceptanceData, err := c.acceptanceDataStore.Get(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	daaAddedBlocksSet, err := c.daaAddedBlocksSet(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	txOuts := make([]*externalapi.DomainTransactionOutput, 0, len(ghostdagData.MergeSetBlues()))
	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	for _, blue := range ghostdagData.MergeSetBlues() {
		txOut, hasReward, err := c.coinbaseOutputForBlueBlock(stagingArea, blue, acceptanceDataMap[*blue], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		if hasReward {
			txOuts = append(txOuts, txOut)
		}
	}

	txOut, hasRedReward, err := c.coinbaseOutputForRewardFromRedBlocks(
		stagingArea, ghostdagData, acceptanceData, daaAddedBlocksSet, coinbaseData)
	if err != nil {
		return nil, false, err
	}

	if hasRedReward {
		txOuts = append(txOuts, txOut)
	}

	subsidy, err := c.CalcBlockSubsidy(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	payload, err := c.serializeCoinbasePayload(ghostdagData.BlueScore(), coinbaseData, subsidy)
	if err != nil {
		return nil, false, err
	}

	return &externalapi.DomainTransaction{
		Version:      constants.MaxTransactionVersion,
		Inputs:       []*externalapi.DomainTransactionInput{},
		Outputs:      txOuts,
		LockTime:     0,
		SubnetworkID: subnetworks.SubnetworkIDCoinbase,
		Gas:          0,
		Payload:      payload,
	}, hasRedReward, nil
}

func (c *coinbaseManager) daaAddedBlocksSet(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (
	hashset.HashSet, error) {

	daaAddedBlocks, err := c.daaBlocksStore.DAAAddedBlocks(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, err
	}

	return hashset.NewFromSlice(daaAddedBlocks...), nil
}

// coinbaseOutputForBlueBlock calculates the output that should go into the coinbase transaction of blueBlock
// If blueBlock gets no fee - returns nil for txOut
func (c *coinbaseManager) coinbaseOutputForBlueBlock(stagingArea *model.StagingArea,
	blueBlock *externalapi.DomainHash, blockAcceptanceData *externalapi.BlockAcceptanceData,
	mergingBlockDAAAddedBlocksSet hashset.HashSet) (*externalapi.DomainTransactionOutput, bool, error) {

	blockReward, err := c.calcMergedBlockReward(stagingArea, blueBlock, blockAcceptanceData, mergingBlockDAAAddedBlocksSet)
	if err != nil {
		return nil, false, err
	}

	if blockReward == 0 {
		return nil, false, nil
	}

	// the ScriptPublicKey for the coinbase is parsed from the coinbase payload
	_, coinbaseData, _, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(blockAcceptanceData.TransactionAcceptanceData[0].Transaction)
	if err != nil {
		return nil, false, err
	}

	txOut := &externalapi.DomainTransactionOutput{
		Value:           blockReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}

	return txOut, true, nil
}

func (c *coinbaseManager) coinbaseOutputForRewardFromRedBlocks(stagingArea *model.StagingArea,
	ghostdagData *externalapi.BlockGHOSTDAGData, acceptanceData externalapi.AcceptanceData, daaAddedBlocksSet hashset.HashSet,
	coinbaseData *externalapi.DomainCoinbaseData) (*externalapi.DomainTransactionOutput, bool, error) {

	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	totalReward := uint64(0)
	for _, red := range ghostdagData.MergeSetReds() {
		reward, err := c.calcMergedBlockReward(stagingArea, red, acceptanceDataMap[*red], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		totalReward += reward
	}

	if totalReward == 0 {
		return nil, false, nil
	}

	return &externalapi.DomainTransactionOutput{
		Value:           totalReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}, true, nil
}

func acceptanceDataFromArrayToMap(acceptanceData externalapi.AcceptanceData) map[externalapi.DomainHash]*externalapi.BlockAcceptanceData {
	acceptanceDataMap := make(map[externalapi.DomainHash]*externalapi.BlockAcceptanceData, len(acceptanceData))
	for _, blockAcceptanceData := range acceptanceData {
		acceptanceDataMap[*blockAcceptanceData.BlockHash] = blockAcceptanceData
	}
	return acceptanceDataMap
}

// CalcBlockSubsidy returns the subsidy amount a block at the provided blue score
// should have. This is mainly used for determining how much the coinbase for
// newly generated blocks awards as well as validating the coinbase for blocks
// has the expected value.
func (c *coinbaseManager) CalcBlockSubsidy(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (uint64, error) {
	if blockHash.Equal(c.genesisHash) {
		return c.subsidyGenesisReward, nil
	}
	blockDaaScore, err := c.daaBlocksStore.DAAScore(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}
	if blockDaaScore < c.deflationaryPhaseDaaScore {
		return c.preDeflationaryPhaseBaseSubsidy, nil
	}

	blockSubsidy := c.calcDeflationaryPeriodBlockSubsidy(blockDaaScore)
	return blockSubsidy, nil
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidy(blockDaaScore uint64) uint64 {
	// We define a year as 365.25 days and a month as 365.25 / 12 = 30.4375
	// secondsPerMonth = 30.4375 * 24 * 60 * 60
	const secondsPerMonth = 2629800
	// Note that this calculation implicitly assumes that block per second = 1 (by assuming daa score diff is in second units).
	monthsSinceDeflationaryPhaseStarted := (blockDaaScore - c.deflationaryPhaseDaaScore) / secondsPerMonth
	// Return the pre-calculated value from subsidy-per-month table
	return c.getDeflationaryPeriodBlockSubsidyFromTable(monthsSinceDeflationaryPhaseStarted)
}

/*
This table was pre-calculated by calling `calcDeflationaryPeriodBlockSubsidyFloatCalc` for all months until reaching 0 subsidy.
To regenerate this table, run `TestBuildSubsidyTable` in coinbasemanager_test.go (note the `deflationaryPhaseBaseSubsidy` therein)
*/
var subsidyByDeflationaryMonthTable = []uint64{
	4400000000, 4278340444, 4160044764, 4045019946, 3933175554, 3824423648, 3718678720, 3615857630, 3515879532, 3418665818, 3324140054, 3232227917, 3142857142, 3055957460, 2971460545, 2889299962, 2809411110, 2731731177, 2656199086, 2582755450, 2511342523, 2441904156, 2374385753, 2308734227, 2244897959,
	2182826757, 2122471818, 2063785687, 2006722221, 1951236555, 1897285061, 1844825321, 1793816087, 1744217254, 1695989823, 1649095876, 1603498542, 1559161969, 1516051298, 1474132633, 1433373015, 1393740396, 1355203615, 1317732372, 1281297205, 1245869467, 1211421302, 1177925626, 1145356101, 1113687121,
	1082893784, 1052951881, 1023837868, 995528854, 968002582, 941237408, 915212289, 889906762, 865300930, 841375447, 818111501, 795490800, 773495560, 752108486, 731312762, 711092039, 691430416, 672312434, 653723064, 635647687, 618072093, 600982462, 584365357, 568207714, 552496829,
	537220347, 522366259, 507922885, 493878868, 480223167, 466945045, 454034062, 441480066, 429273187, 417403827, 405862653, 394640592, 383728819, 373118756, 362802060, 352770620, 343016548, 333532175, 324310044, 315342904, 306623705, 298145590, 289901895, 281886137, 274092014,
	266513397, 259144329, 251979014, 245011820, 238237268, 231650031, 225244931, 219016932, 212961136, 207072782, 201347240, 195780010, 190366712, 185103092, 179985010, 175008443, 170169477, 165464308, 160889237, 156440665, 152115097, 147909130, 143819457, 139842864, 135976223,
	132216494, 128560721, 125006030, 121549626, 118188791, 114920883, 111743332, 108653640, 105649378, 102728184, 99887760, 97125873, 94440353, 91829086, 89290021, 86821161, 84420565, 82086345, 79816666, 77609743, 75463841, 73377274, 71348400, 69375624, 67457395,
	65592204, 63778587, 62015115, 60300403, 58633103, 57011904, 55435531, 53902744, 52412338, 50963142, 49554017, 48183853, 46851574, 45556133, 44296511, 43071717, 41880788, 40722788, 39596807, 38501960, 37437384, 36402244, 35395726, 34417038, 33465410,
	32540095, 31640365, 30765512, 29914848, 29087706, 28283434, 27501400, 26740989, 26001603, 25282661, 24583598, 23903864, 23242925, 22600260, 21975365, 21367749, 20776933, 20202453, 19643857, 19100706, 18572573, 18059044, 17559713, 17074189, 16602089,
	16143043, 15696689, 15262678, 14840666, 14430323, 14031326, 13643361, 13266124, 12899317, 12542652, 12195849, 11858635, 11530745, 11211921, 10901912, 10600476, 10307373, 10022376, 9745258, 9475803, 9213798, 8959037, 8711320, 8470453, 8236246,
	8008515, 7787080, 7571768, 7362409, 7158840, 6960898, 6768430, 6581284, 6399312, 6222372, 6050324, 5883033, 5720368, 5562200, 5408406, 5258864, 5113457, 4972070, 4834593, 4700917, 4570937, 4444551, 4321660, 4202166, 4085977,
	3973000, 3863147, 3756331, 3652469, 3551479, 3453280, 3357798, 3264955, 3174679, 3086900, 3001547, 2918555, 2837857, 2759390, 2683094, 2608906, 2536770, 2466629, 2398427, 2332110, 2267628, 2204928, 2143962, 2084682, 2027040,
	1970993, 1916495, 1863504, 1811979, 1761878, 1713162, 1665793, 1619734, 1574949, 1531401, 1489058, 1447886, 1407852, 1368925, 1331074, 1294270, 1258484, 1223687, 1189852, 1156953, 1124963, 1093858, 1063613, 1034204, 1005608,
	977803, 950767, 924479, 898917, 874062, 849894, 826395, 803545, 781327, 759723, 738717, 718292, 698431, 679119, 660342, 642083, 624330, 607067, 590282, 573961, 558091, 542659, 527655, 513065, 498879,
	485085, 471673, 458631, 445950, 433619, 421630, 409972, 398636, 387614, 376896, 366475, 356342, 346489, 336909, 327593, 318535, 309728, 301164, 292837, 284740, 276867, 269211, 261768, 254530, 247492,
	240649, 233995, 227525, 221234, 215117, 209169, 203385, 197762, 192294, 186977, 181807, 176780, 171892, 167139, 162518, 158024, 153655, 149406, 145275, 141258, 137353, 133555, 129862, 126271, 122780,
	119385, 116084, 112874, 109753, 106719, 103768, 100899, 98109, 95396, 92758, 90194, 87700, 85275, 82917, 80624, 78395, 76227, 74120, 72070, 70078, 68140, 66256, 64424, 62643, 60910,
	59226, 57589, 55996, 54448, 52943, 51479, 50055, 48671, 47325, 46017, 44745, 43507, 42304, 41135, 39997, 38891, 37816, 36770, 35754, 34765, 33804, 32869, 31960, 31077, 30217,
	29382, 28569, 27779, 27011, 26264, 25538, 24832, 24145, 23478, 22829, 22197, 21584, 20987, 20407, 19842, 19294, 18760, 18241, 17737, 17247, 16770, 16306, 15855, 15417, 14990,
	14576, 14173, 13781, 13400, 13029, 12669, 12319, 11978, 11647, 11325, 11012, 10707, 10411, 10123, 9843, 9571, 9307, 9049, 8799, 8556, 8319, 8089, 7865, 7648, 7436,
	7231, 7031, 6836, 6647, 6464, 6285, 6111, 5942, 5778, 5618, 5463, 5312, 5165, 5022, 4883, 4748, 4617, 4489, 4365, 4244, 4127, 4013, 3902, 3794, 3689,
	3587, 3488, 3391, 3298, 3206, 3118, 3031, 2948, 2866, 2787, 2710, 2635, 2562, 2491, 2422, 2355, 2290, 2227, 2165, 2105, 2047, 1990, 1935, 1882, 1830,
	1779, 1730, 1682, 1636, 1590, 1546, 1504, 1462, 1422, 1382, 1344, 1307, 1271, 1236, 1201, 1168, 1136, 1104, 1074, 1044, 1015, 987, 960, 933, 908,
	882, 858, 834, 811, 789, 767, 746, 725, 705, 685, 667, 648, 630, 613, 596, 579, 563, 548, 532, 518, 503, 489, 476, 463, 450,
	438, 425, 414, 402, 391, 380, 370, 359, 349, 340, 330, 321, 312, 304, 295, 287, 279, 271, 264, 257, 249, 243, 236, 229, 223,
	217, 211, 205, 199, 194, 188, 183, 178, 173, 168, 164, 159, 155, 150, 146, 142, 138, 134, 131, 127, 124, 120, 117, 114, 110,
	107, 104, 101, 99, 96, 93, 91, 88, 86, 83, 81, 79, 76, 74, 72, 70, 68, 66, 65, 63, 61, 59, 58, 56, 54,
	53, 52, 50, 49, 47, 46, 45, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 28, 27,
	26, 25, 25, 24, 23, 23, 22, 21, 21, 20, 20, 19, 18, 18, 17, 17, 16, 16, 16, 15, 15, 14, 14, 13, 13,
	13, 12, 12, 12, 11, 11, 11, 10, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 7, 7, 7, 6, 6,
	6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
}

func (c *coinbaseManager) getDeflationaryPeriodBlockSubsidyFromTable(month uint64) uint64 {
	if month >= uint64(len(subsidyByDeflationaryMonthTable)) {
		month = uint64(len(subsidyByDeflationaryMonthTable) - 1)
	}
	return subsidyByDeflationaryMonthTable[month]
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidyFloatCalc(month uint64) uint64 {
	baseSubsidy := c.deflationaryPhaseBaseSubsidy
	subsidy := float64(baseSubsidy) / math.Pow(1.4, float64(month)/12)
	return uint64(subsidy)
}

func (c *coinbaseManager) calcMergedBlockReward(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	blockAcceptanceData *externalapi.BlockAcceptanceData, mergingBlockDAAAddedBlocksSet hashset.HashSet) (uint64, error) {

	if !blockHash.Equal(blockAcceptanceData.BlockHash) {
		return 0, errors.Errorf("blockAcceptanceData.BlockHash is expected to be %s but got %s",
			blockHash, blockAcceptanceData.BlockHash)
	}

	if !mergingBlockDAAAddedBlocksSet.Contains(blockHash) {
		return 0, nil
	}

	totalFees := uint64(0)
	for _, txAcceptanceData := range blockAcceptanceData.TransactionAcceptanceData {
		if txAcceptanceData.IsAccepted {
			totalFees += txAcceptanceData.Fee
		}
	}

	block, err := c.blockStore.Block(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}

	_, _, subsidy, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(block.Transactions[transactionhelper.CoinbaseTransactionIndex])
	if err != nil {
		return 0, err
	}

	return subsidy + totalFees, nil
}

// New instantiates a new CoinbaseManager
func New(
	databaseContext model.DBReader,

	subsidyGenesisReward uint64,
	preDeflationaryPhaseBaseSubsidy uint64,
	coinbasePayloadScriptPublicKeyMaxLength uint8,
	genesisHash *externalapi.DomainHash,
	deflationaryPhaseDaaScore uint64,
	deflationaryPhaseBaseSubsidy uint64,

	dagTraversalManager model.DAGTraversalManager,
	ghostdagDataStore model.GHOSTDAGDataStore,
	acceptanceDataStore model.AcceptanceDataStore,
	daaBlocksStore model.DAABlocksStore,
	blockStore model.BlockStore,
	pruningStore model.PruningStore,
	blockHeaderStore model.BlockHeaderStore) model.CoinbaseManager {

	return &coinbaseManager{
		databaseContext: databaseContext,

		subsidyGenesisReward:                    subsidyGenesisReward,
		preDeflationaryPhaseBaseSubsidy:         preDeflationaryPhaseBaseSubsidy,
		coinbasePayloadScriptPublicKeyMaxLength: coinbasePayloadScriptPublicKeyMaxLength,
		genesisHash:                             genesisHash,
		deflationaryPhaseDaaScore:               deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy:            deflationaryPhaseBaseSubsidy,

		dagTraversalManager: dagTraversalManager,
		ghostdagDataStore:   ghostdagDataStore,
		acceptanceDataStore: acceptanceDataStore,
		daaBlocksStore:      daaBlocksStore,
		blockStore:          blockStore,
		pruningStore:        pruningStore,
		blockHeaderStore:    blockHeaderStore,
	}
}
