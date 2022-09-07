package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// DisassociateInstance decouples an allocated Floating IP from an instance
func DisassociateInstance(client *golangsdk.ServiceClient, serverID string, opts DisassociateOptsBuilder) (r DisassociateResult) {
	b, err := opts.ToFloatingIPDisassociateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(disassociateURL(client, serverID), b, nil, nil)
	return
}
