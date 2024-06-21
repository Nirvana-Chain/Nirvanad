package testutils

import (
	"fmt"
	"testing"

	"github.com/Nirvana-Chain/nirvanad/cmd/nirvanawallet/keys"
)

// ForAllPaths runs the passed testFunc with all available derivation paths
func ForAllPaths(t *testing.T, testFunc func(*testing.T, uint32)) {
	for i := uint32(1); i <= keys.LastVersion; i++ {
		version := fmt.Sprintf("v%d", i)
		t.Run(version, func(t *testing.T) {
			t.Logf("Running test for wallet version %d", i)
			testFunc(t, i)
		})
	}
}
