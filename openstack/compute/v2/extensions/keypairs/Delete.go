package keypairs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previous stored KeyPair from the server.
func Delete(client *golangsdk.ServiceClient, name string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("os-keypairs", name), nil)
	return
}
