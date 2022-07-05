package cellular

import (
	"fmt"

	"github.com/sberryman/playground-golang/utils"
	"github.com/sberryman/playground-golang/utils/s2s"
)

type SimState string

const (
	Preflight   SimState = "preflight"
	Active      SimState = "active"
	Suspended   SimState = "suspended"
	Deactivated SimState = "deactivated"
	Unknown     SimState = "unknown"
)

type SimDataLimit struct {
	Limit     int64 `json:"limit,omitempty"`
	Used      int64 `json:"used,omitempty"`
	ResetAt   int64 `json:"reset_at,omitempty"`
	OverLimit bool  `json:"over_limit,omitempty"`
}

type SimDevice struct {
	Provider        string        `json:"provider"`
	Iccid           string        `json:"iccid"`
	Imei            string        `json:"imei,omitempty"`
	State           SimState      `json:"state"`
	CreatedAt       int64         `json:"created_at"`
	LastConnectedAt int64         `json:"last_connected_at"`
	DataLimit       *SimDataLimit `json:"data_limit,omitempty"`
}

type SimChangeStateResponse struct {
	Ok bool `json:"ok"`
}

func SimByIccid(iccid string) (SimDevice, error) {
	resp := SimDevice{}

	err := s2s.NewOp(
		&s2s.Op{
			Hostname: utils.Config.Cellular.ClusterHostname,
			Route:    "/sim/getByIccid",
			ReqBody: map[string]string{
				"iccid": iccid,
			},
			Resp: &resp,
		},
	).Post()

	return resp, err
}

func SimChangeStateByIccid(iccid string, action string) error {
	resp := SimChangeStateResponse{}

	err := s2s.NewOp(
		&s2s.Op{
			Hostname: utils.Config.Cellular.ClusterHostname,
			Route:    "/sim/changeStateByIccid",
			ReqBody: map[string]string{
				"iccid":  iccid,
				"action": action,
			},
			Resp: &resp,
		},
	).Post()

	if err != nil {
		return err
	}

	if !resp.Ok {
		return fmt.Errorf("failed to change state of sim %s", iccid)
	}

	return nil
}
