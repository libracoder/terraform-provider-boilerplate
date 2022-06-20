package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	serverOpts := &plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return interna()
		},
	}
	plugin.Serve(serverOpts)
}
