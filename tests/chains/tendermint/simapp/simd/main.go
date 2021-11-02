package main

import (
	"os"

	"github.com/Electronic-Signatures-Industries/yui-relayer/tests/tendermint/simapp"
	"github.com/Electronic-Signatures-Industries/yui-relayer/tests/tendermint/simapp/simd/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
