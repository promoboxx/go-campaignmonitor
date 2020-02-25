package campaignmonitor

// Rule describes a segmentation rule
type Rule struct {
	RuleType string
	Clause   string
}

// RuleGroup describes a set of rules to be OR'd together.  Each group is then AND'd together
type RuleGroup struct {
	Rules []Rule
}

// CreateSegmentRequest contains the necessary data to create a segment
type CreateSegmentRequest struct {
	Title      string
	RuleGroups []RuleGroup
}
