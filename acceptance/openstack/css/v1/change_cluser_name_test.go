package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// const timeout = 1200

func TestChangeClusterNameWorkflow(t *testing.T) {
	// POST /v1.0/{project_id}/clusters/{cluster_id}/changename
	// {
	// "displayName" : "ES-Test-new"
	// }
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	cssID := os.Getenv("CSS_ID")
	// cssID := clients.EnvOS.GetEnv("CSS_ID")

	if cssID == "" {
		t.Skip("`CSS_ID` need to be defined")
	}

	opts := clusters.ChangeClusterNameOpts{
		DisplayName: tools.RandomString("changed-css-", 4),
	}
	err = clusters.ChangeClusterName(client, opts, cssID)
	th.AssertNoErr(t, err)
}
