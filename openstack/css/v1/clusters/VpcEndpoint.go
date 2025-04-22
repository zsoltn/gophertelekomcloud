package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnablePrivateDnsOpts struct {
	// Indicates whether to enable the internal DNS name .
	EndpointWithDnsName *bool `json:"endpointWithDnsName"`
}

// EnableVpcEndpoint function is used to Enable VCP endpoint of the cluster.
func EnableVpcEndpoint(client *golangsdk.ServiceClient, clusterID string, opts []EnablePrivateDnsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("clusters", clusterID, "vpcepservice", "open")
	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}

// EnableVpcEndpoint function is used to Disable VCP endpoint of the cluster.
func DisableVpcEndpoint(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "vpcepservice", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
