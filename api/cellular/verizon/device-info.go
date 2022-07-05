package verizon

import (
	"fmt"

	"github.com/sberryman/playground-golang/utils"
	"github.com/sberryman/playground-golang/utils/s2s"
)

type Devices struct {
	HasMoreData bool     `json:"hasMoreData"`
	Devices     []Device `json:"devices"`
}
type CarrierInformations struct {
	CarrierName string `json:"carrierName"`
	ServicePlan string `json:"servicePlan"`
	State       string `json:"state"`
}
type DeviceIds struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}
type ExtendedAttributes struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
type Device struct {
	AccountName         string                `json:"accountName"`
	BillingCycleEndDate string                `json:"billingCycleEndDate"`
	CarrierInformations []CarrierInformations `json:"carrierInformations"`
	Connected           bool                  `json:"connected"`
	CreatedAt           string                `json:"createdAt"`
	DeviceIds           []DeviceIds           `json:"deviceIds"`
	ExtendedAttributes  []ExtendedAttributes  `json:"extendedAttributes"`
	GroupNames          []string              `json:"groupNames"`
	IPAddress           string                `json:"ipAddress"`
	LastActivationBy    string                `json:"lastActivationBy"`
	LastActivationDate  string                `json:"lastActivationDate"`
	LastConnectionDate  string                `json:"lastConnectionDate"`
}

func RetrieveDeviceInformation(iccid string) (Device, error) {
	resp := Devices{}

	err := s2s.NewOp(
		&s2s.Op{
			Hostname: utils.Config.Cellular.ClusterHostname,
			Route:    "/verizon/api/m2m/v1/devices/actions/list",
			ReqBody: map[string]interface{}{
				"deviceId": map[string]string{
					"id":   iccid,
					"kind": "iccid",
				},
			},
			Resp: &resp,
		},
	).Post()

	if len(resp.Devices) != 1 {
		return Device{}, fmt.Errorf("no device found for iccid %s", iccid)
	}

	return resp.Devices[0], err
}
