package volumeattach

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns a Pager that allows you to iterate over a collection of
// VolumeAttachments.
func List(client *golangsdk.ServiceClient, serverID string) pagination.Pager {
	return pagination.NewPager(client, listURL(client, serverID), func(r pagination.PageResult) pagination.Page {
		return VolumeAttachmentPage{pagination.SinglePageBase(r)}
	})
}
