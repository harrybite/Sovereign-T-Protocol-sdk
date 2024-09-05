package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				IdList: []types.Id{
					{
						Did: "0",
					},
					{
						Did: "1",
					},
				},
				UniquekeyList: []types.Uniquekey{
					{
						Key: "0",
					},
					{
						Key: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated id",
			genState: &types.GenesisState{
				IdList: []types.Id{
					{
						Did: "0",
					},
					{
						Did: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated uniquekey",
			genState: &types.GenesisState{
				UniquekeyList: []types.Uniquekey{
					{
						Key: "0",
					},
					{
						Key: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
