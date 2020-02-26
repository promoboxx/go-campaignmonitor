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
	CreateCampaignFromTemplate(ctx context.Context, req CreateCampaignFromTemplateRequest) (string, glitch.DataError)
	SendCampaign(ctx context.Context, campaignID string, req SendCampaignRequest) glitch.DataError
	CreateSegment(ctx context.Context, listID string, req CreateSegmentRequest) (string, glitch.DataError)
	ImportSubscribers(ctx context.Context, listID string, req ImportSubscribersRequest) (*ImportSubscriberResponse, glitch.DataError)
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

func readResponse(status int, response []byte, ret interface{}) glitch.DataError {
	if status >= 200 && status <= 299 {
		err := json.Unmarshal(response, ret)
		if err != nil {
			return glitch.NewDataError(err, ErrorJSON, fmt.Sprintf("could not unmarshal response: %s", response))
		}
		return nil
	}
	cmErr := Error{}
	err := json.Unmarshal(response, &cmErr)
	if err != nil {
		return glitch.NewDataError(err, ErrorJSON, fmt.Sprintf("could not unmarshal error response: %s", response))
	}
	return glitch.NewDataError(&cmErr, fmt.Sprintf("%d", cmErr.Code), cmErr.Message)
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

	gErr = readResponse(status, resBy, &res)
	if gErr != nil {
		return nil, gErr
	}
	return res, nil
}

// CreateCampaignFromTemplate will create a campaign from the provided data and return the template ID
// https://www.campaignmonitor.com/api/campaigns/#creating-draft-campaign
func (sc *serviceClient) CreateCampaignFromTemplate(ctx context.Context, req CreateCampaignFromTemplateRequest) (string, glitch.DataError) {
	slug := fmt.Sprintf("campaigns/%s/fromtemplate.json", sc.clientID)

	reader, gErr := client.ObjectToJSONReader(&req)
	if gErr != nil {
		return "", gErr
	}

	status, resBy, gErr := sc.c.MakeRequest(ctx, http.MethodPost, slug, nil, nil, reader)
	if gErr != nil {
		return "", gErr
	}

	var res string
	gErr = readResponse(status, resBy, &res)
	if gErr != nil {
		return "", gErr
	}
	return res, nil
}

// SendCampaign will trigger the send of a specific campaign at a specific time.  If SendDate is nil it will send immediately
// https://www.campaignmonitor.com/api/campaigns/#sending-draft-campaign
func (sc *serviceClient) SendCampaign(ctx context.Context, campaignID string, req SendCampaignRequest) glitch.DataError {
	slug := fmt.Sprintf("campaigns/%s/send.json", campaignID)

	p := req.process()
	reader, gErr := client.ObjectToJSONReader(&p)
	if gErr != nil {
		return gErr
	}

	status, resBy, gErr := sc.c.MakeRequest(ctx, http.MethodPost, slug, nil, nil, reader)
	if gErr != nil {
		return gErr
	}

	if status < 200 || status > 299 {
		cmErr := Error{}
		err := json.Unmarshal(resBy, &cmErr)
		if err != nil {
			return glitch.NewDataError(err, ErrorJSON, fmt.Sprintf("could not unmarshal error response: %s", resBy))
		}
		return glitch.NewDataError(&cmErr, fmt.Sprintf("%d", cmErr.Code), cmErr.Message)
	}
	return nil
}

// CreateSegment creates a segment on a list with the provided rules
// Rules within a group are OR'd together
// RuleGroups are AND'd together
// https://www.campaignmonitor.com/api/segments/#creating-a-segment
func (sc *serviceClient) CreateSegment(ctx context.Context, listID string, req CreateSegmentRequest) (string, glitch.DataError) {
	slug := fmt.Sprintf("segments/%s.json", listID)

	reader, gErr := client.ObjectToJSONReader(&req)
	if gErr != nil {
		return "", gErr
	}

	status, resBy, gErr := sc.c.MakeRequest(ctx, http.MethodPost, slug, nil, nil, reader)
	if gErr != nil {
		return "", gErr
	}

	var res string
	gErr = readResponse(status, resBy, &res)
	if gErr != nil {
		return "", gErr
	}
	return res, nil
}

// ImportSubscribers bulk upserts subscribers into the list provided.
// https://www.campaignmonitor.com/api/subscribers/#importing-many-subscribers
func (sc *serviceClient) ImportSubscribers(ctx context.Context, listID string, req ImportSubscribersRequest) (*ImportSubscriberResponse, glitch.DataError) {
	slug := fmt.Sprintf("subscribers/%s/import.json", listID)

	reader, gErr := client.ObjectToJSONReader(&req)
	if gErr != nil {
		return nil, gErr
	}

	status, resBy, gErr := sc.c.MakeRequest(ctx, http.MethodPost, slug, nil, nil, reader)
	if gErr != nil {
		return nil, gErr
	}

	var res ImportSubscriberResponse
	gErr = readResponse(status, resBy, &res)
	if gErr != nil {
		return nil, gErr
	}
	return &res, nil
}
