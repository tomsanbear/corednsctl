package control

import (
	"context"
	"net"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

// CoreControl is our main struct to control the world of remote control
type CoreControl struct {
	Next     plugin.Handler
	listener net.Listener
}

// New provides a new CoreControl object, and starts the server
func New() (CoreControl, error) {
	var cc CoreControl
	return cc, nil
}

// Close the server down, should be handled when caddy turns off
func (cc *CoreControl) Close() {
	err := cc.listener.Close()
	if err != nil {
		log.Warningf("failed to close https server: %v", err)
	}
}

// Name is the name of the plugin
func (cc *CoreControl) Name() string { return "control" }

// ServeDNS implements the Handler interface.
func (cc *CoreControl) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// Need to bring up the server

	// Record response to get status code and size of the reply.
	status, err := plugin.NextOrFailure(cc.Name(), cc.Next, ctx, w, r)

	return status, err
}

func (cc *CoreControl) RegisterCommand(cmd ControlCommand) {
	// Register the command in the registry, make available to be used
}
