package campaignmonitor

// Attachment contains the base64 encoded content of the attachment as well as the name and content type (e.g. application/pdf)
type Attachment struct {
	Content string
	Name    string
	Type    string
}

// SmartEmailResponse contains the status and id of the smart email for a recipient
type SmartEmailResponse struct {
	Status    string
	MessageID string
	Recipient string
}
type SmartEmailRequest struct {
	To                  []string
	CC                  []string
	BCC                 []string
	Attachments         []Attachment
	Data                map[string]interface{}
	AddRecipientsToList bool
	ConsentToTrack      bool
}
