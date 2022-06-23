package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/libracoder/terraform-provider-boilerplate/resources/person"
	"github.com/libracoder/terraform-provider-boilerplate/resources/uuid_count"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"boilerplate_uuid":   uuid_count.ResourceServer(),
			"boilerplate_person": person.ResourceServer(),
		},
	}
}
