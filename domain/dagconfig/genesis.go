// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/Nirvana-Chain/nirvanad/domain/consensus/model/externalapi"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/blockheader"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/subnetworks"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/transactionhelper"
	"github.com/kaspanet/go-muhash"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
	0x4e, 0x69, 0x72, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x77, 0x61,
	0x73, 0x20, 0x62, 0x6f, 0x72, 0x6e, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x73, 0x20, 0x63, 0x61, 0x72, 0x64,
	0x20, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x6e, 0x79, 0xa2, 0x5a, 0xb3, 0x38, 0x85, 0x25,
	0xc1, 0x25, 0x55, 0xd4, 0x7c, 0xa5, 0x9a, 0xc1,
	0xd7, 0x80, 0x45, 0xee, 0xde, 0x16, 0x27, 0x56,
	0xd3, 0xa5, 0x53, 0xcf, 0xb3, 0x98, 0x94, 0x3d,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x6b, 0x97, 0xb4, 0xa1, 0x68, 0xbe, 0x44, 0xed,
	0xc6, 0x37, 0xbe, 0xcc, 0x94, 0x00, 0xca, 0xfc,
	0x89, 0x49, 0x07, 0xd2, 0x26, 0xaf, 0xa9, 0xeb,
	0x8f, 0x90, 0x38, 0x5c, 0x4f, 0x2f, 0x58, 0x0b,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x18FEBB9AFB0,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
	0x6E, 0x69, 0x72, 0x76, 0x61, 0x6E, 0x61, 0x2D,
	0x64, 0x65, 0x76, 0x6E, 0x65, 0x74,
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x3a, 0x22, 0x94, 0x88, 0x40, 0xb1, 0x29, 0x99,
	0x50, 0xae, 0xa4, 0x87, 0x44, 0x43, 0x02, 0x76,
	0x13, 0x34, 0xc7, 0x8f, 0xd1, 0x15, 0x2e, 0x51,
	0x6f, 0x48, 0xff, 0xee, 0x3b, 0x5d, 0x12, 0x2b,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x27, 0x64, 0x68, 0x72, 0x49, 0xcf, 0xde, 0xb7,
	0xf5, 0xe2, 0x71, 0xfb, 0xc1, 0x5d, 0xd0, 0xd8,
	0x10, 0x45, 0xe1, 0x28, 0xcd, 0xbe, 0xd3, 0xbd,
	0xe1, 0xb2, 0xfa, 0x7a, 0x65, 0xab, 0xa7, 0x52,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		1,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x18F57AFDBF0,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
	0x6E, 0x69, 0x72, 0x76, 0x61, 0x6E, 0x61, 0x2D,
	0x73, 0x69, 0x6D, 0x6E, 0x65, 0x74,
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x85, 0x02, 0x8f, 0x0b, 0x0d, 0xd4, 0x46, 0xd6,
	0xb0, 0x9b, 0x42, 0x49, 0xff, 0x1e, 0xd9, 0x69,
	0x28, 0x9d, 0xdb, 0x15, 0x94, 0xaa, 0x73, 0x35,
	0x58, 0x30, 0xdd, 0x10, 0x1e, 0x56, 0x05, 0x12,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the development network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x9a, 0xa6, 0xcc, 0x13, 0x70, 0xf9, 0x8c, 0x13,
	0x67, 0x2b, 0xba, 0x08, 0x34, 0xc6, 0x54, 0x47,
	0x38, 0x83, 0x4a, 0xe0, 0xbd, 0x64, 0x84, 0x58,
	0xb7, 0x6a, 0x95, 0x7b, 0xb4, 0x6f, 0x24, 0xed,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		1,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x18F57AFDBF0,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
	0x6E, 0x69, 0x72, 0x76, 0x61, 0x6E, 0x61, 0x2D,
	0x74, 0x65, 0x73, 0x74, 0x6E, 0x65, 0x74,
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x25, 0xd5, 0x82, 0x20, 0xd2, 0xd2, 0x83, 0x4d,
	0xa7, 0x17, 0xbf, 0x37, 0xce, 0x4a, 0x93, 0x91,
	0xfe, 0xd9, 0x16, 0x68, 0x3b, 0x83, 0x2c, 0x9d,
	0x19, 0x68, 0xab, 0xe0, 0x05, 0x7d, 0x7b, 0x04,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xd8, 0x62, 0x85, 0x11, 0x61, 0xe9, 0xbd, 0xbb,
	0xe5, 0x02, 0xc4, 0x9e, 0x86, 0xe8, 0x75, 0xac,
	0xa9, 0xea, 0x6e, 0x20, 0x08, 0x36, 0x85, 0xee,
	0xff, 0xca, 0x2c, 0xb6, 0xbc, 0x69, 0xc3, 0x67,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		1,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x18F57AFDBF0,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
