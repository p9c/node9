package pkgs

import (
	"reflect"
	"strings"
)

func Name(dtype interface{}) string {
	name := reflect.TypeOf(dtype).PkgPath()
	name = strings.TrimPrefix(
		name,
		"github.com/p9c/node9/",
	)
	return name
}
