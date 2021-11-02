package module

import (
	"github.com/Electronic-Signatures-Industries/yui-relayer/chains/ethermint"
	"github.com/Electronic-Signatures-Industries/yui-relayer/chains/ethermint/cmd"
	"github.com/Electronic-Signatures-Industries/yui-relayer/config"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/spf13/cobra"
)

type Module struct{}

var _ config.ModuleI = (*Module)(nil)

// Name returns the name of the module
func (Module) Name() string {
	return "ethermint"
}

// RegisterInterfaces register the module interfaces to protobuf Any.
func (Module) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	ethermint.RegisterInterfaces(registry)
}

// GetCmd returns the command
func (Module) GetCmd(ctx *config.Context) *cobra.Command {
	return cmd.EthermintCmd(ctx.Codec, ctx)
}