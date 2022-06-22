package person

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"age": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"city": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
