module github.com/Electronic-Signatures-Industries/yui-relayer

go 1.15

require (
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ibc-go v1.2.2
	github.com/datachainlab/ibc-mock-client v0.0.0-20210801010718-05f8b1087574
	github.com/ethereum/go-ethereum v1.10.11
	github.com/gogo/protobuf v1.3.3
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20210801023756-05047b73f120
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	github.com/tharsis/ethermint v0.7.2
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil => github.com/btcsuite/btcutil v1.0.2
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
