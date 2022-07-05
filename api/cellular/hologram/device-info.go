package hologram

import (
	"fmt"

	"github.com/sberryman/playground-golang/utils"
	"github.com/sberryman/playground-golang/utils/s2s"
)

type DeviceList struct {
	Success   bool     `json:"success"`
	Limit     int      `json:"limit"`
	Size      int      `json:"size"`
	Continues bool     `json:"continues"`
	Lastid    int      `json:"lastid"`
	Data      []Device `json:"data"`
}
type Plan struct {
	ID          int    `json:"id"`
	Zone        string `json:"zone"`
	Name        string `json:"name"`
	Data        int    `json:"data"`
	Cost        string `json:"cost"`
	Sms         string `json:"sms"`
	Overage     string `json:"overage"`
	AccountTier string `json:"account_tier"`
}
type Cellular struct {
	ID                  int    `json:"id"`
	Sim                 string `json:"sim"`
	Imsi                int64  `json:"imsi"`
	Msisdn              string `json:"msisdn"`
	Carrierid           int    `json:"carrierid"`
	State               string `json:"state"`
	Whenclaimed         string `json:"whenclaimed"`
	Whenexpires         string `json:"whenexpires"`
	Overagelimit        int    `json:"overagelimit"`
	DataThreshold       int    `json:"data_threshold"`
	Smslimit            int    `json:"smslimit"`
	Pin                 string `json:"pin"`
	Puk                 string `json:"puk"`
	Apn                 string `json:"apn"`
	Plan                Plan   `json:"plan"`
	LastNetworkUsed     string `json:"last_network_used"`
	LastConnectTime     string `json:"last_connect_time"`
	CurBillingDataUsed  int    `json:"cur_billing_data_used"`
	LastBillingDataUsed int    `json:"last_billing_data_used"`
}
type Links struct {
	Cellular []Cellular `json:"cellular"`
}
type Lastsession struct {
	Linkid                interface{} `json:"linkid"` // string or a number?
	Bytes                 int         `json:"bytes"`
	SessionBegin          string      `json:"session_begin"`
	SessionEnd            string      `json:"session_end"`
	Imei                  string      `json:"imei"`
	Cellid                interface{} `json:"cellid"` // string or a number?
	Tadig                 string      `json:"tadig"`
	Lac                   interface{} `json:"lac"` // string or a number?
	NetworkName           string      `json:"network_name"`
	RadioAccessTechnology string      `json:"radio_access_technology"`
	Active                bool        `json:"active"`
}
type Device struct {
	ID              int           `json:"id"`
	Orgid           int           `json:"orgid"`
	Name            string        `json:"name"`
	Type            string        `json:"type"`
	Whencreated     string        `json:"whencreated"`
	Phonenumber     string        `json:"phonenumber"`
	PhonenumberCost string        `json:"phonenumber_cost"`
	Tunnelable      int           `json:"tunnelable"`
	Imei            string        `json:"imei"`
	ImeiSv          string        `json:"imei_sv"`
	Hidden          int           `json:"hidden"`
	Tags            []interface{} `json:"tags"`
	Links           Links         `json:"links"`
	Lastsession     Lastsession   `json:"lastsession"`
	Model           string        `json:"model"`
	Manufacturer    string        `json:"manufacturer"`
}

func ListDeviceByIccid(iccid string) (Device, error) {
	resp := DeviceList{}

	err := s2s.NewOp(
		&s2s.Op{
			Hostname: utils.Config.Cellular.ClusterHostname,
			Route:    "/hologram/api/1/devices",
			ReqQuery: map[string]string{
				"orgid": "21766",
				"sim":   iccid,
			},
			Resp: &resp,
		},
	).Get()

	if len(resp.Data) != 1 {
		return Device{}, fmt.Errorf("no device found for iccid %s", iccid)
	}

	return resp.Data[0], err
}
