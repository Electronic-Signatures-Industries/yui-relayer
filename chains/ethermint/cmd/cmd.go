package cmd

import (
	"github.com/Electronic-Signatures-Industries/yui-relayer/config"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func TendermintCmd(m codec.Codec, ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tendermint",
		Short: "manage tendermint configurations",
	}

	cmd.AddCommand(
		configCmd(m),
		keysCmd(ctx),
		lightCmd(ctx),
	)

	return cmd
}
