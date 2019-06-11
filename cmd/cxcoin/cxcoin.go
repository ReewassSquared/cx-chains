/*
skycoin daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
	_ "net/http/pprof"
	"os"

	"github.com/amherag/skycoin/src/readable"
	"github.com/amherag/skycoin/src/skycoin"
	"github.com/amherag/skycoin/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.25.0-rc1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// CoinName name of coin
	CoinName = "cxcoin"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "af0b3942acdfcd890b2cedc5a4c2e473008f0f4f3422bdad914e6cb6bfdf14a9567f20676d4c7f0d0c4af10716d7f31b0794a288e8636976f5f23cb94134e77b01"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "23v7mT1uLpViNKZHh9aww4VChxizqKsNq4E"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02583e5ebbf85522474e0f17e681e62ca37910db6b8792763af4e97663c31a7984"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1426562704
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 100000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	}

	nodeConfig = skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		CoinName:                       CoinName,
		GenesisSignatureStr:            GenesisSignatureStr,
		GenesisAddressStr:              GenesisAddressStr,
		GenesisCoinVolume:              GenesisCoinVolume,
		GenesisTimestamp:               GenesisTimestamp,
		BlockchainPubkeyStr:            BlockchainPubkeyStr,
		BlockchainSeckeyStr:            BlockchainSeckeyStr,
		DefaultConnections:             DefaultConnections,
		PeerListURL:                    "https://127.0.0.1/peers.txt",
		Port:                           6000,
		WebInterfacePort:               6420,
		DataDirectory:                  "$HOME/.cxcoin",
		UnconfirmedBurnFactor:          2,
		UnconfirmedMaxTransactionSize:  65535,
		UnconfirmedMaxDropletPrecision: 3,
		CreateBlockBurnFactor:          2,
		CreateBlockMaxTransactionSize:  65535,
		CreateBlockMaxDropletPrecision: 3,
		MaxBlockSize:                   65535,
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := skycoin.NewCoin(skycoin.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}