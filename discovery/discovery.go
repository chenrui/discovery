package discovery

import (
	"context"

	"github.com/Bilibili/discovery/conf"
	"github.com/Bilibili/discovery/lib/http"
	"github.com/Bilibili/discovery/registry"
)

// Discovery discovery.
type Discovery struct {
	c        *conf.Config
	client   *http.Client
	registry *registry.Registry
	nodes    *registry.Nodes
}

// New get a discovery.
func New(c *conf.Config) (d *Discovery, cancel context.CancelFunc) {
	d = &Discovery{
		c:        c,
		client:   http.NewClient(c.HTTPClient),
		registry: registry.NewRegistry(),
		nodes:    registry.NewNodes(c),
	}
	d.syncUp()
	cancel = d.regSelf()
	go d.nodesproc()
	return
}
