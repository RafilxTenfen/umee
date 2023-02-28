package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgExample{}

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
func (m *MsgExample) ValidateBasic() error {
	return nil
}

// Signers returns the addrs of signers that must sign.
// CONTRACT: All signatures must be present to be valid.
// CONTRACT: Returns addrs in some deterministic order.
func (m *MsgExample) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
