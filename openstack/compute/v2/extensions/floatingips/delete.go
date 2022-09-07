package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previous allocated Floating IP.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("os-floating-ips", id), nil)
	return
}
