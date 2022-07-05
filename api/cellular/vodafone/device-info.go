package vodafone

import (
	"fmt"

	"github.com/sberryman/playground-golang/utils"
	"github.com/sberryman/playground-golang/utils/s2s"
)

type DeviceDetails struct {
	GetDeviceDetailsv2Response GetDeviceDetailsv2Response `json:"getDeviceDetailsv2Response"`
}
type ReturnCode struct {
	MajorReturnCode string `json:"majorReturnCode"`
	MinorReturnCode string `json:"minorReturnCode"`
}
type DeviceInformationItem struct {
	ItemName  string `json:"itemName"`
	ItemType  string `json:"itemType"`
	ItemValue string `json:"itemValue,omitempty"`
}
type DeviceInformationList struct {
	DeviceInformationItem []DeviceInformationItem `json:"deviceInformationItem"`
}
type ApnListItem struct {
	ApnName         string `json:"apnName"`
	StaticIPAddress string `json:"staticIpAddress"`
}
type ApnList struct {
	ApnListItem ApnListItem `json:"apnListItem"`
}
type DeviceReturn struct {
	ReturnCode             ReturnCode            `json:"returnCode"`
	DeviceID               string                `json:"deviceId"`
	CustomerServiceProfile string                `json:"customerServiceProfile"`
	State                  string                `json:"state"`
	BaseCountry            string                `json:"baseCountry"`
	DeviceInformationList  DeviceInformationList `json:"deviceInformationList"`
	ApnList                ApnList               `json:"apnList"`
}
type GetDeviceDetailsv2Response struct {
	Return DeviceReturn `json:"return"`
}

func ListDeviceByIccid(iccid string) (DeviceReturn, error) {
	resp := DeviceDetails{}

	err := s2s.NewOp(
		&s2s.Op{
			Hostname: utils.Config.Cellular.ClusterHostname,
			Route:    "/vodafone/m2m/v1/devices/" + iccid,
			Resp:     &resp,
		},
	).Get()

	if resp.GetDeviceDetailsv2Response.Return.ReturnCode.MajorReturnCode != "000" {
		return DeviceReturn{}, fmt.Errorf("no device found for iccid %s", iccid)
	}

	return resp.GetDeviceDetailsv2Response.Return, err
}
