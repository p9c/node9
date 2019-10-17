// Package app is a multi-function universal binary that does all the things.
//
// Parallelcoin Pod
//
// This is the heart of configuration and coordination of
// the parts that compose the parallelcoin Pod - Ctl, Node and Wallet, and
// the extended, combined Shell and the webview GUI.
package app

import (
	"fmt"
	"os"

	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/log"
)

const (
	appName           = "pod"
	confExt           = ".json"
	podConfigFilename = appName + confExt
	// ctlAppName           = "ctl"
	// ctlConfigFilename    = ctlAppName + confExt
	// nodeAppName          = "node"
	// nodeConfigFilename   = nodeAppName + confExt
	// walletAppName        = "wallet"
	// walletConfigFilename = walletAppName + confExt
)

// Main is the entrypoint for the pod AiO suite
func Main() int {
	cx := conte.GetNewContext(appName, "main")
	cx.App = getApp(cx)
	log.DEBUG("running App")
	e := cx.App.Run(os.Args)
	if e != nil {
		fmt.Println("Pod ERROR:", e)
		return 1
	}
	return 0
}
