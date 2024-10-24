package reserved

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	FuncUrn       string         `json:"-"`
	Count         *int           `json:"count" required:"true"`
	IdleMode      *bool          `json:"idle_mode"`
	TacticsConfig *TacticsConfig `json:"tactics_config"`
}

type TacticsConfig struct {
	CronConfigs []CronConfig `json:"cron_configs"`
}
type CronConfig struct {
	Name        string `json:"name,omitempty"`
	Cron        string `json:"cron,omitempty"`
	Count       int    `json:"count,omitempty"`
	StartTime   int    `json:"start_time,omitempty"`
	ExpiredTime int    `json:"expired_time,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*FuncReservedRespUpdate, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "reservedinstances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncReservedRespUpdate
	return &res, extract.Into(raw.Body, &res)
}

type FuncReservedRespUpdate struct {
	TacticsConfig *TacticsConfig `json:"tactics_config"`
	IdleMode      bool           `json:"idle_mode"`
	Count         int            `json:"count"`
}
