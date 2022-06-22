package person

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}
