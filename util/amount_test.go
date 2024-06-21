// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util_test

import (
	"math"
	"testing"

	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/constants"
)

func TestAmountCreation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		valid    bool
		expected Amount
	}{
		// Positive tests.
		{
			name:     "zero",
			amount:   0,
			valid:    true,
			expected: 0,
		},
		{
			name:     "max producible",
			amount:   4961e6,
			valid:    true,
			expected: Amount(constants.MaxSompi),
		},
		{
			name:     "one hundred",
			amount:   100,
			valid:    true,
			expected: 100 * constants.SompiPerNirvana,
		},
		{
			name:     "fraction",
			amount:   0.01234567,
			valid:    true,
			expected: 1234567,
		},
		{
			name:     "rounding up",
			amount:   54.999999999999943157,
			valid:    true,
			expected: 55 * constants.SompiPerNirvana,
		},
		{
			name:     "rounding down",
			amount:   55.000000000000056843,
			valid:    true,
			expected: 55 * constants.SompiPerNirvana,
		},

		// Negative tests.
		{
			name:   "not-a-number",
			amount: math.NaN(),
			valid:  false,
		},
		{
			name:   "-infinity",
			amount: math.Inf(-1),
			valid:  false,
		},
		{
			name:   "+infinity",
			amount: math.Inf(1),
			valid:  false,
		},
	}

	for _, test := range tests {
		a, err := NewAmount(test.amount)
		switch {
		case test.valid && err != nil:
			t.Errorf("%v: Positive test Amount creation failed with: %v", test.name, err)
			continue
		case !test.valid && err == nil:
			t.Errorf("%v: Negative test Amount creation succeeded (value %v) when should fail", test.name, a)
			continue
		}

		if a != test.expected {
			t.Errorf("%v: Created amount %v does not match expected %v", test.name, a, test.expected)
			continue
		}
	}
}

func TestAmountUnitConversions(t *testing.T) {
	tests := []struct {
		name      string
		amount    Amount
		unit      AmountUnit
		converted float64
		s         string
	}{
		{
			name:      "MNVA",
			amount:    Amount(constants.MaxSompi),
			unit:      AmountMegaNVA,
			converted: 4961,
			s:         "4961 MNVA",
		},
		{
			name:      "kNVA",
			amount:    44433322211100,
			unit:      AmountKiloNVA,
			converted: 444.33322211100,
			s:         "444.333222111 kNVA",
		},
		{
			name:      "NVA",
			amount:    44433322211100,
			unit:      AmountNVA,
			converted: 444333.22211100,
			s:         "444333.222111 NVA",
		},
		{
			name:      "mNVA",
			amount:    44433322211100,
			unit:      AmountMilliNVA,
			converted: 444333222.11100,
			s:         "444333222.111 mNVA",
		},
		{

			name:      "μNVA",
			amount:    44433322211100,
			unit:      AmountMicroNVA,
			converted: 444333222111.00,
			s:         "444333222111 μNVA",
		},
		{

			name:      "sompi",
			amount:    44433322211100,
			unit:      AmountSompi,
			converted: 44433322211100,
			s:         "44433322211100 Sompi",
		},
		{

			name:      "non-standard unit",
			amount:    44433322211100,
			unit:      AmountUnit(-1),
			converted: 4443332.2211100,
			s:         "4443332.22111 1e-1 NVA",
		},
	}

	for _, test := range tests {
		f := test.amount.ToUnit(test.unit)
		if f != test.converted {
			t.Errorf("%v: converted value %v does not match expected %v", test.name, f, test.converted)
			continue
		}

		s := test.amount.Format(test.unit)
		if s != test.s {
			t.Errorf("%v: format '%v' does not match expected '%v'", test.name, s, test.s)
			continue
		}

		// Verify that Amount.ToNVA works as advertised.
		f1 := test.amount.ToUnit(AmountNVA)
		f2 := test.amount.ToNVA()
		if f1 != f2 {
			t.Errorf("%v: ToNVA does not match ToUnit(AmountNVA): %v != %v", test.name, f1, f2)
		}

		// Verify that Amount.String works as advertised.
		s1 := test.amount.Format(AmountNVA)
		s2 := test.amount.String()
		if s1 != s2 {
			t.Errorf("%v: String does not match Format(AmountNVA): %v != %v", test.name, s1, s2)
		}
	}
}

func TestAmountMulF64(t *testing.T) {
	tests := []struct {
		name string
		amt  Amount
		mul  float64
		res  Amount
	}{
		{
			name: "Multiply 0.1 NVA by 2",
			amt:  100e5, // 0.1 NVA
			mul:  2,
			res:  200e5, // 0.2 NVA
		},
		{
			name: "Multiply 0.2 NVA by 0.02",
			amt:  200e5, // 0.2 NVA
			mul:  1.02,
			res:  204e5, // 0.204 NVA
		},
		{
			name: "Round down",
			amt:  49, // 49 Sompi
			mul:  0.01,
			res:  0,
		},
		{
			name: "Round up",
			amt:  50, // 50 Sompi
			mul:  0.01,
			res:  1, // 1 Sompi
		},
		{
			name: "Multiply by 0.",
			amt:  1e8, // 1 NVA
			mul:  0,
			res:  0, // 0 NVA
		},
		{
			name: "Multiply 1 by 0.5.",
			amt:  1, // 1 Sompi
			mul:  0.5,
			res:  1, // 1 Sompi
		},
		{
			name: "Multiply 100 by 66%.",
			amt:  100, // 100 Sompi
			mul:  0.66,
			res:  66, // 66 Sompi
		},
		{
			name: "Multiply 100 by 66.6%.",
			amt:  100, // 100 Sompi
			mul:  0.666,
			res:  67, // 67 Sompi
		},
		{
			name: "Multiply 100 by 2/3.",
			amt:  100, // 100 Sompi
			mul:  2.0 / 3,
			res:  67, // 67 Sompi
		},
	}

	for _, test := range tests {
		a := test.amt.MulF64(test.mul)
		if a != test.res {
			t.Errorf("%v: expected %v got %v", test.name, test.res, a)
		}
	}
}
