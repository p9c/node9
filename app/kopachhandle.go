package app

import (
	"sync"

	"github.com/urfave/cli"

	"github.com/p9c/node9/cmd/kopach"
	"github.com/p9c/node9/pkg/conte"
	"github.com/p9c/node9/pkg/util/interrupt"
)

func kopachHandle(cx *conte.Xt) func(c *cli.Context) (err error) {
	return func(c *cli.Context) (err error) {
		var wg sync.WaitGroup
		Configure(cx)
		quit := make(chan struct{})
		interrupt.AddHandler(func(){
			close(quit)
		})
		kopach.Main(cx, quit, &wg)
		wg.Wait()
		return
	}
}
