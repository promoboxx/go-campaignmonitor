package campaignmonitor

import "net/http"

type authRoundTripper struct {
	baseTripper http.RoundTripper
	apiKey      string
}

func (a *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(a.apiKey, "x")
	return a.baseTripper.RoundTrip(req)
}
