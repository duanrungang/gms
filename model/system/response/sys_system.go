package response

import "gms/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
