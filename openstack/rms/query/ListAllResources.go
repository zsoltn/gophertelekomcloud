package query

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllOpts struct {
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
	// Specifies the region ID.
	RegionId string `q:"region_id"`
	// Specifies the resource type
	Type string `q:"type"`
	// Specifies the resource ID.
	Id string `q:"id"`
	// Specifies the resource name.
	Name string `q:"name"`
	// Specifies tags. The format is key or key=value.
	Tags []string `q:"tags"`
}

func ListAllResources(client *golangsdk.ServiceClient, domainId string, opts ListAllOpts) ([]Resource, error) {
	// GET /v1/resource-manager/domains/{domain_id}/all-resources
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "all-resources").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractResources(pages)
}
