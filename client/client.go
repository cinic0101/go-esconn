package client

import (
	"github.com/cinic0101/go-esconn/api"
	"github.com/cinic0101/go-esconn/config"
)

type Client struct {
	*api.API
}

func New(config *config.Config) (*Client, error) {
	return &Client{
		api.New(config),
	}, nil
}