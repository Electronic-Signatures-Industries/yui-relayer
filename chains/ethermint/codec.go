package ethermint

import (
	"github.com/Electronic-Signatures-Industries/yui-relayer/core"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	evmtypes "github.com/tharsis/ethermint/types"
)

type ExtensionOptionsWeb3TxI interface{}

// RegisterInterfaces register the module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
	// RegisterInterfaces register the Ethermint key concrete types.
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})

	// RegisterInterfaces registers the tendermint concrete client-related
	// implementations and interfaces.
	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&evmtypes.EthAccount{},
	)
	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&evmtypes.EthAccount{},
	)
	registry.RegisterInterface(
		"ethermint.v1.ExtensionOptionsWeb3Tx",
		(*ExtensionOptionsWeb3TxI)(nil),
		&evmtypes.ExtensionOptionsWeb3Tx{},
	)
}
