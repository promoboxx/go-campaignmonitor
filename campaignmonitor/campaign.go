package campaignmonitor

import "time"

// Singleline represents content for floating singleline tags.
type Singleline struct {
	Content string
	Href    *string
}

// Multiline represents content for floating multiline tags
type Multiline struct {
	Content string
}

// Image represents content for floating editable image tags
type Image struct {
	Content string
	Alt     *string
	Href    *string
}

// Item represents content for single/multi/images within a layout
type Item struct {
	Layout      *string
	Singlelines []Singleline
	Multilines  []Multiline
	Images      []Image
}

// Repeater represents content for repeaters. Each repeater should contain an Items collection
type Repeater struct {
	Items []Item
}

// TemplateContent holds the content to be replaced in the template
type TemplateContent struct {
	Singlelines []Singleline
	Multilines  []Multiline
	Images      []Image
	Repeaters   []Repeater
}

// CreateCampaignFromTemplateRequest contains all the data necessary for creating a campaign from a template
type CreateCampaignFromTemplateRequest struct {
	Name            string
	Subject         string
	FromName        string
	FromEmail       string
	ReplyTo         string
	ListIDs         []string
	SegmentIDs      []string
	TemplateID      string
	TemplateContent TemplateContent
}

// SendCampaignRequest contains the data to send a campaign.
// SendDate should be null to send immediately - it will be sent as the string "immediately"
type SendCampaignRequest struct {
	ConfirmationEmail string
	SendDate          *time.Time
}
