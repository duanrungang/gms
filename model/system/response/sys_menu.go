package response

import (
	system2 "gms/model/system"
)

type SysMenusResponse struct {
	Menus []system2.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system2.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system2.SysBaseMenu `json:"menu"`
}
