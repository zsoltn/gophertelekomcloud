package defsecrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List will return a collection of default rules.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("os-security-group-default-rules"), func(r pagination.PageResult) pagination.Page {
		return DefaultRulePage{pagination.SinglePageBase(r)}
	})
}
