package ethereum

import (
	"github.com/Electronic-Signatures-Industries/yui-relayer/core"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces register the module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
}
