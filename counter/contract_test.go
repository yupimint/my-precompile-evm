// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package counter

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	tests = map[string]testutils.PrecompileTest{
		"insufficient gas for getCounter should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				input, err := PackGetCounter()
				require.NoError(t, err)
				return input
			},
			SuppliedGas: GetCounterGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"readOnly incrementCounter should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				input, err := PackIncrementCounter()
				require.NoError(t, err)
				return input
			},
			SuppliedGas: IncrementCounterGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"insufficient gas for incrementCounter should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				input, err := PackIncrementCounter()
				require.NoError(t, err)
				return input
			},
			SuppliedGas: IncrementCounterGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"readOnly setCounter should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput = big.NewInt(1)
				input, err := PackSetCounter(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetCounterGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"insufficient gas for setCounter should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput = big.NewInt(1)
				input, err := PackSetCounter(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetCounterGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
	}
)

// TestCounterRun tests the Run function of the precompile contract.
func TestCounterRun(t *testing.T) {
	// Run tests.
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Run(t, Module, state.NewTestStateDB(t))
		})
	}
}

func BenchmarkCounter(b *testing.B) {
	// Benchmark tests.
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			test.Bench(b, Module, state.NewTestStateDB(b))
		})
	}
}
