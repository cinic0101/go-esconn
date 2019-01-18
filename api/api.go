package api

import (
	"github.com/cinic0101/go-esconn/config"
)

type API struct {
	Config *config.Config
}

func New(config *config.Config) *API {
	return &API{
		Config: config,
	}
}