package control

import (
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/mholt/caddy"
)

var (
	log = clog.NewWithPlugin("control")
)

func init() {
	caddy.RegisterPlugin("control", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	// Setup up the server and other tings here
	return nil
}

func parse(c *caddy.Controller) (cc CoreControl, err error) {
	return cc, nil
}
