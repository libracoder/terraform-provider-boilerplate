package uuid_count

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"net/http"
)

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	uuidCount := d.Get("uuid_count").(string)

	d.SetId(uuidCount)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + uuidCount)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resourceServerRead(d, m)
}
