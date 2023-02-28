package types

// DefaultGenesisState returns the default genesis state for the x/oracle
// module.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}
