package campaignmonitor

import (
	"time"

	"github.com/promoboxx/go-client/client"
)

const (
	baseURL     = "https://api.createsend.com/api/v3.2/"
	serviceName = "campaignmonitor"
)

// Client is a client that can interact with campaign monitor
type Client interface {
}

type serviceClient struct {
	c      client.BaseClient
	apiKey string
}

// NewClient will create a new Client for communicating with campaign monitor
// it requires an apikey for authentication and a a timeout to control how long it waits for requests to finish
func NewClient(apiKey string, timeout time.Duration) Client {
	finder := func(serviceName string) (string, error) {
		return baseURL, nil
	}
	return &serviceClient{c: client.NewBaseClient(finder, serviceName, true, timeout, nil), apiKey: apiKey}
}
