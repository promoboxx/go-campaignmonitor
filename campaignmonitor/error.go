package campaignmonitor

import "fmt"

// Error holds CM error responses
type Error struct {
	Code       int32
	Message    string
	ResultData interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("CampaignMonitor error. Code: %d, message: %s", e.Code, e.Message)
}
