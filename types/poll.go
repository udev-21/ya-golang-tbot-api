package types

type Poll struct {
	ID                    int64           `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       *int64          `json:"correct_option_id,omitempty"`
	Explanation           *string         `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int64          `json:"open_period,omitempty"`
	CloseDate             *int64          `json:"close_date,omitempty"`
}
