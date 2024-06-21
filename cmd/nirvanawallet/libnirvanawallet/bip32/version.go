package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// NirvanaMainnetPrivate is the version that is used for
// nirvana mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var NirvanaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// NirvanaMainnetPublic is the version that is used for
// nirvana mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var NirvanaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// NirvanaTestnetPrivate is the version that is used for
// nirvana testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var NirvanaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// NirvanaTestnetPublic is the version that is used for
// nirvana testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var NirvanaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// NirvanaDevnetPrivate is the version that is used for
// nirvana devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var NirvanaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// NirvanaDevnetPublic is the version that is used for
// nirvana devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var NirvanaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// NirvanaSimnetPrivate is the version that is used for
// nirvana simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var NirvanaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// NirvanaSimnetPublic is the version that is used for
// nirvana simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var NirvanaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case NirvanaMainnetPrivate:
		return NirvanaMainnetPublic, nil
	case NirvanaTestnetPrivate:
		return NirvanaTestnetPublic, nil
	case NirvanaDevnetPrivate:
		return NirvanaDevnetPublic, nil
	case NirvanaSimnetPrivate:
		return NirvanaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case NirvanaMainnetPrivate:
		return true
	case NirvanaTestnetPrivate:
		return true
	case NirvanaDevnetPrivate:
		return true
	case NirvanaSimnetPrivate:
		return true
	}

	return false
}
