package pnl

import (
	"github.com/p9c/pod/cmd/gui/db"
	"github.com/p9c/pod/cmd/gui/mod"
)

func Panels(d db.DuOSdb) (c []mod.DuOScomp) {
	//c = append(c, sys.Boot())
	//c = append(c, sys.Dev())
	//c = append(c, sys.Display())

	//c = append(c, serv.Blocks())
	//c = append(c, panel.Send())
	//c = append(c, panel.AddressBook())
	//c = append(c, panel.Blocks())
	//c = append(c, panel.Peers())
	//c = append(c, panel.Settings())
	//c = append(c, panel.Transactions())
	//c = append(c, panel.TransactionsExcerpts())
	//c = append(c, panel.TimeBalance())
	//c = append(c, sys.Nav())
	//c = append(c, panel.Manager())
	//c = append(c, panel.LayoutConfig())
	//c = append(c, panel.Test())
	//c = append(c, panel.ChartA())
	//c = append(c, panel.ChartB())
	c = append(c, WalletStatus())
	c = append(c, Status())
	c = append(c, LocalHashRate())
	//c = append(c, NetworkHashRate())

	return c
}

func PanelsJs(cs mod.DuOScomps) (s string) {

	for _, c := range cs {

		cc := `var` + c.ID + ` = new Vue({
	el: '#` + c.ID + `',
	name: '` + c.ID + `',
	template: '` + c.Template + `',
	` + c.Js + `});`
		s = s + cc
	}
	return
}
