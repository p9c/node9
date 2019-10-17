package save

import (
	"encoding/json"
	"io/ioutil"

	"github.com/p9c/node9/app/apputil"
	"github.com/p9c/node9/pkg/log"
	"github.com/p9c/node9/pkg/pod"
)

func Pod(c *pod.Config) (success bool) {
	log.INFO("saving configuration to", *c.ConfigFile)
	if yp, e := json.MarshalIndent(c,"", "  "); e == nil {
		apputil.EnsureDir(*c.ConfigFile)
		if e := ioutil.WriteFile(*c.ConfigFile, yp, 0600); e != nil {
			log.ERROR(e)
			return
		}
		return true
	}
	return
}
