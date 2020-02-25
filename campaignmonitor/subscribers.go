package campaignmonitor

// ImportSubscribersRequest holds all the data we need to send an import subscriber request.
type ImportSubscribersRequest struct {
	Subscribers                            []Subscriber
	Resubscribe                            bool
	RestartSubscriptionBasedAutoresponders bool
}

// ImportSubscriberResponse holds the response from CM on an import subscriber call
type ImportSubscriberResponse struct {
	FailureDetails              []interface{}
	TotalUniqueEmailsSubmitted  int32
	TotalExistingSubscribers    int32
	TotalNewSubscribers         int32
	DuplicateEmailsInSubmission []interface{}
}

// Subscribe contains all the data we need to send to add/upsert a subscriber in CM
type Subscriber struct {
	EmailAddress   string
	Name           string
	CustomFields   []KeyValue
	ConsentToTrack bool
}

// KeyValue holds a key value pair
type KeyValue struct {
	Key   string
	Value string
}
