package main

import (
	"log"

	ethereum "github.com/Electronic-Signatures-Industries/yui-relayer/chains/ethereum/module"
	ethermint "github.com/Electronic-Signatures-Industries/yui-relayer/chains/ethermint/module"
	tendermint "github.com/Electronic-Signatures-Industries/yui-relayer/chains/tendermint/module"
	"github.com/Electronic-Signatures-Industries/yui-relayer/cmd"
	mock "github.com/Electronic-Signatures-Industries/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		ethermint.Module{},
		ethereum.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
