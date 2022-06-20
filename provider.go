package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/libracoder/terraform-provider-biolerplate/resources/uuid_count"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"example_server": uuid_count.ResourceServer(),
		},
	}
}
