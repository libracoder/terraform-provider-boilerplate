package person

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
