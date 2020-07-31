package configs

import (
	"fmt"

	"github.com/kumarabd/gokit/utils"
)

type LocalApp struct{}

func NewLocal() (AppHandler, error) {
	return &LocalApp{}, nil
}

// -------------------------------------------Application config methods----------------------------------------------------------------

func (l LocalApp) Get(result interface{}) error {
	v := fmt.Sprintf(`{
		"name":"%s",
		"config":"%s",
		"port":"%s",
		"proxy":"%s",
		"debug":"%s",
		"version":"%s"
	}`, appname, local, port, proxy, debug, version)
	return utils.Unmarshal(v, result)
}
