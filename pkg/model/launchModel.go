package model

import "github.com/eFlink/launch-dashboard-server/pkg"

type Launch struct {
	pkg.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}
