package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// PolicyCreateOpts contains options for creating a snapshot policy.
// This object is passed to the snapshots.PolicyCreate function.
type ChangeClusterNameOpts struct {
	DisplayName string `json:"displayName" required:"true"`
}

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
func ChangeClusterName(client *golangsdk.ServiceClient, opts ChangeClusterNameOpts, clusterId string) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "changename"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
