package uuid_count

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
