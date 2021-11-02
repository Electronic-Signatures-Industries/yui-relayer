package mock

import (
	"github.com/Electronic-Signatures-Industries/yui-relayer/core"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	mocktypes "github.com/datachainlab/ibc-mock-client/modules/light-clients/xx-mock/types"
)

// RegisterInterfaces register the module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	mocktypes.RegisterInterfaces(registry)

	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
}
