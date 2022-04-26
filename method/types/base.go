package types

type ReplyMarkup interface {
	IsReplyMarkup()
}

type ReplyMarkuper struct {
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *ReplyMarkuper) WithReplyMarkup(replyMarkup ReplyMarkup) {
	m.ReplyMarkup = replyMarkup
}

type ParseModer struct {
	ParseMode *string `json:"parse_mode,omitempty"`
}

const (
	ParseModeMarkdownV2 string = "MarkdownV2"
	ParseModeHTML       string = "HTML"
	ParseModeMarkdown   string = "Markdown"
)

func (m *ParseModer) WithParseModeHTML() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeHTML
}

func (m *ParseModer) WithParseModeParseModeMarkdownV2() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdownV2
}

func (m *ParseModer) WithParseModeParseModeMarkdown() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdown
}

type Durationer struct {
	Duration *int64 `json:"duration,omitempty"`
}

func (d *Durationer) WithDuration(duration int64) {
	d.Duration = &duration
}
