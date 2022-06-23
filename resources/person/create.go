package person

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	person := d.Get("name").(string)

	d.SetId(person)

	//// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	//resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + uuidCount)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()

	return resourceServerRead(d, m)
}
