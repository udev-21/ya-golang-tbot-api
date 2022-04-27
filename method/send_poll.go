package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendPoll(question string, options []string) *Poll {
	return &Poll{
		Question: question,
		Options:  options,
	}
}

// https://core.telegram.org/bots/api#sendpoll
type Poll struct {
	Question              string                 `json:"question"`
	Options               []string               `json:"options"`
	IsAnonymous           *bool                  `json:"is_anonymous,omitempty"`
	Type                  *string                `json:"type,omitempty"`
	AllowsMultipleAnswers bool                   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID       *int64                 `json:"correct_option_id,omitempty"`
	Explanation           *string                `json:"explanation,omitempty"`
	ExplanationParseMode  *string                `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities   []*types.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int64                 `json:"open_period,omitempty"`
	CloseDate             *int64                 `json:"close_date,omitempty"`
	IsClosed              bool                   `json:"is_closed,omitempty"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (p *Poll) Endpoint() string {
	return "sendPoll"
}

func (p *Poll) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(p)
}

func (p *Poll) WithIsAnonymous(isAnonymous bool) {
	p.IsAnonymous = &isAnonymous
}

func (p *Poll) WithTypeQuiz() {
	if p.Type == nil {
		p.Type = new(string)
	}
	*p.Type = "quiz"
}

func (p *Poll) WithAllowsMultipleAnswers() {
	p.AllowsMultipleAnswers = true
}

func (p *Poll) WithCorrectOptionID(correctOptionID int64) {
	p.CorrectOptionID = &correctOptionID
}

func (p *Poll) WithExplanation(explanation string) {
	p.Explanation = &explanation
}

func (p *Poll) WithExplanationParseModeHTML() {
	if p.ExplanationParseMode == nil {
		p.ExplanationParseMode = new(string)
	}
	*p.ExplanationParseMode = myTypes.ParseModeHTML
}

func (p *Poll) WithExplanationParseModeMarkdown() {
	if p.ExplanationParseMode == nil {
		p.ExplanationParseMode = new(string)
	}
	*p.ExplanationParseMode = myTypes.ParseModeMarkdown
}

func (p *Poll) WithExplanationParseModeMarkdownV2() {
	if p.ExplanationParseMode == nil {
		p.ExplanationParseMode = new(string)
	}
	*p.ExplanationParseMode = myTypes.ParseModeMarkdownV2
}

func (p *Poll) WithExplanationEntities(entities []*types.MessageEntity) {
	p.ExplanationEntities = entities
}

func (p *Poll) WithOpenPeriod(openPeriod int64) {
	p.OpenPeriod = &openPeriod
}

func (p *Poll) WithCloseDate(closeDate int64) {
	p.CloseDate = &closeDate
}

func (p *Poll) WithIsClosed() {
	p.IsClosed = true
}
