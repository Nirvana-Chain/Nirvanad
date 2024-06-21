package libnirvanawallet

import (
	"fmt"

	"github.com/Nirvana-Chain/nirvanad/cmd/nirvanawallet/libnirvanawallet/bip32"
	"github.com/Nirvana-Chain/nirvanad/domain/dagconfig"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
)

// CreateMnemonic creates a new bip-39 compatible mnemonic
func CreateMnemonic() (string, error) {
	const bip39BitSize = 256
	entropy, _ := bip39.NewEntropy(bip39BitSize)
	return bip39.NewMnemonic(entropy)
}

// Purpose and CoinType constants
const (
	SingleSignerPurpose = 44
	// Note: this is not entirely compatible to BIP 45 since
	// BIP 45 doesn't have a coin type in its derivation path.
	MultiSigPurpose = 45
	// Registered in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	CoinType = 121337
	// Wallet version 1 coin type
	CoinTypeV1 = 111111
)

func defaultPath(isMultisig bool, version uint32) string {
	purpose := SingleSignerPurpose
	if isMultisig {
		purpose = MultiSigPurpose
	}

	// Note: this is needed because initial fork was created
	// without changing the coin type in derivation path.
	if version == 1 {
		return fmt.Sprintf("m/%d'/%d'/0'", purpose, CoinTypeV1)
	}

	return fmt.Sprintf("m/%d'/%d'/0'", purpose, CoinType)
}

// MasterPublicKeyFromMnemonic returns the master public key with the correct derivation for the given mnemonic.
func MasterPublicKeyFromMnemonic(params *dagconfig.Params, mnemonic string, isMultisig bool, version uint32) (string, error) {
	path := defaultPath(isMultisig, version)
	extendedKey, err := extendedKeyFromMnemonicAndPath(mnemonic, path, params)
	if err != nil {
		return "", err
	}

	extendedPublicKey, err := extendedKey.Public()
	if err != nil {
		return "", err
	}

	return extendedPublicKey.String(), nil
}

func extendedKeyFromMnemonicAndPath(mnemonic string, path string, params *dagconfig.Params) (*bip32.ExtendedKey, error) {
	seed := bip39.NewSeed(mnemonic, "")
	version, err := versionFromParams(params)
	if err != nil {
		return nil, err
	}

	master, err := bip32.NewMasterWithPath(seed, version, path)
	if err != nil {
		return nil, err
	}

	return master, nil
}

func versionFromParams(params *dagconfig.Params) ([4]byte, error) {
	switch params.Name {
	case dagconfig.MainnetParams.Name:
		return bip32.NirvanaMainnetPrivate, nil
	case dagconfig.TestnetParams.Name:
		return bip32.NirvanaTestnetPrivate, nil
	case dagconfig.DevnetParams.Name:
		return bip32.NirvanaDevnetPrivate, nil
	case dagconfig.SimnetParams.Name:
		return bip32.NirvanaSimnetPrivate, nil
	}

	return [4]byte{}, errors.Errorf("unknown network %s", params.Name)
}
