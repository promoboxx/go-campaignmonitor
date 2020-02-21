package campaignmonitor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/promoboxx/go-glitch/glitch"

	"github.com/promoboxx/go-client/client"
)

const (
	baseURL     = "https://api.createsend.com/api/v3.2/"
	serviceName = "campaignmonitor"

	//error codes
	ErrorJSON = "ERROR_JSON"
)

// Client is a client that can interact with campaign monitor
type Client interface {
	SendSmartEmail(ctx context.Context, smartEmailID string, req SmartEmailRequest) ([]SmartEmailResponse, glitch.DataError)
}

type serviceClient struct {
	c        client.BaseClient
	clientID string
}

// NewClient will create a new Client for communicating with campaign monitor
// it requires an apikey for authentication and a a timeout to control how long it waits for requests to finish
func NewClient(apiKey string, clientID string, timeout time.Duration) Client {
	finder := func(serviceName string) (string, error) {
		return baseURL, nil
	}
	// the authRoundTripper will inject the api key into the request as basic auth
	rt := authRoundTripper{apiKey: apiKey, baseTripper: http.DefaultTransport}
	return &serviceClient{c: client.NewBaseClient(finder, serviceName, true, timeout, &rt), clientID: clientID}
}

// SendSmartEmail will send the specified smart email id with attachments and variable replacements specified in req to the people specified in req.
// Note that CM only allows a maximum of 25 recipients per request.
func (sc *serviceClient) SendSmartEmail(ctx context.Context, smartEmailID string, req SmartEmailRequest) ([]SmartEmailResponse, glitch.DataError) {
	slug := fmt.Sprintf("transactional/smartEmail/%s/send", smartEmailID)

	reader, gErr := client.ObjectToJSONReader(&req)
	if gErr != nil {
		return nil, gErr
	}
	var res []SmartEmailResponse
	status, resBy, gErr := sc.c.MakeRequest(ctx, http.MethodPost, slug, nil, nil, reader)
	if gErr != nil {
		return nil, gErr
	}

	if status >= 200 && status <= 299 {
		err := json.Unmarshal(resBy, &res)
		if err != nil {
			return nil, glitch.NewDataError(err, ErrorJSON, fmt.Sprintf("could not unmarshal response: %s", resBy))
		}
		return res, nil
	}
	cmErr := Error{}
	err := json.Unmarshal(resBy, &cmErr)
	if err != nil {
		return nil, glitch.NewDataError(err, ErrorJSON, fmt.Sprintf("could not unmarshal error response: %s", resBy))
	}
	return nil, glitch.NewDataError(&cmErr, fmt.Sprintf("%d", cmErr.Code), cmErr.Message)
}
