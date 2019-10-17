package panel

import "github.com/p9c/pod/cmd/gui/vue/mod"

func Blocks() mod.DuoVUEcomp {
	return mod.DuoVUEcomp{
		IsApp:    true,
		Name:     "Blocks",
		ID:       "panelblocks",
		Version:  "0.0.1",
		CompType: "core",
		SubType:  "blocks",
		Js: `
	data () { return { 
	duoSystem,
       pageSettings: { pageSize: 10, pageSizes: [10,20,50,100], pageCount: 5 }

	}},
		`,
		Template: `<div class="rwrap">
        <ejs-grid :dataSource="this.duoSystem.blocks" height="100%" :allowPaging="true" :pageSettings='pageSettings'>
			<e-columns>
				<e-column field='height' headerText='Height' textAlign='Left' width=60></e-column>
				<e-column field='time' headerText='Time' textAlign='Center' width=90></e-column>
				<e-column field='hash' headerText='Hash' textAlign='Center' width=240></e-column>
				<e-column field='txnum' headerText='Transactions' textAlign='Right' width=90></e-column>
			</e-columns>
        </ejs-grid>
</div>`,
		Css: `
		
		`,
	}
}
