package main

import (
	"github.com/turbot/steampipe-plugin-net/net"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: net.Plugin})
}
