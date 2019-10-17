package conte

import (
	"sync/atomic"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/urfave/cli"

	"github.com/p9c/pod/app/appdata"
	"github.com/p9c/pod/cmd/node/rpc"
	"github.com/p9c/pod/cmd/node/state"
	`github.com/p9c/pod/pkg/chain/config/netparams`
	"github.com/p9c/pod/pkg/pod"
	"github.com/p9c/pod/pkg/wallet"
)

type _dtype int

var _d _dtype

// Xt as in conte.Xt stores all the common state data used in pod
type Xt struct {
	// App is the heart of the application system,
	// this creates and initialises it.
	App *cli.App
	// Config is the pod all-in-one server config
	Config *pod.Config
	// StateCfg is a reference to the main node state configuration struct
	StateCfg *state.Config
	// ActiveNet is the active net parameters
	ActiveNet *netparams.Params
	// DataDir is the default data dir
	DataDir string
	// Node is the run state of the node
	Node *atomic.Value
	// NodeKill is the killswitch for the Node
	NodeKill chan struct{}
	// TestNode is the run state of the TestNode
	TestNode *atomic.Value
	// TestNodeKill is the killswitch for the TestNode
	TestNodeKill chan struct{}
	// Wallet is the run state of the wallet
	Wallet *atomic.Value
	// WalletKill is the killswitch for the Wallet
	WalletKill chan struct{}
	// // Window is the fyne window when running GUI
	// Window *fyne.Window
	// RPCServer is needed to directly query data
	RPCServer *rpc.Server
	// WalletServer is needed to query the wallet
	WalletServer *wallet.Wallet
	// Scribble DB
	DB *scribble.Driver
	// RealNode is the main node
	RealNode *rpc.Node
}

// GetNewContext returns a fresh new context
func GetNewContext(appName string, subtext string) *Xt {
	return &Xt{
		App:      cli.NewApp(),
		Config:   pod.EmptyConfig(),
		StateCfg: new(state.Config),
		DataDir:  appdata.Dir(appName, false),
	}
}
