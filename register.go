package control

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/mholt/caddy"
)

// RegisterCliOperation is called by plugins to register new commands to be provided to the cli/api
func RegisterCliOperation(c *caddy.Controller, cmds ...ControlCommand) {
	cc := dnsserver.GetConfig(c).Handler("control")
	if cc == nil {
		return
	}
	if x, ok := cc.(*CoreControl); !ok {
		for _, cmd := range cmds {
			x.RegisterCommand(cmd)
		}
	}
}
