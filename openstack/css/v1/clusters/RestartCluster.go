package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// PolicyCreateOpts contains options for creating a snapshot policy.
// This object is passed to the snapshots.PolicyCreate function.

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
// https://docs.otc.t-systems.com/cloud-search-service/api-ref/cluster_management_apis/restarting_a_cluster.html#css-03-0021
func RestartCluster(client *golangsdk.ServiceClient, clusterId string) (err error) {
	_, err = client.Post(client.ServiceURL("clusters", clusterId, "restart"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
