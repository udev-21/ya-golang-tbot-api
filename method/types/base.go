package types

type RawJsonPayloader interface {
	RawJsonPayload() (map[string]interface{}, error)
}
type ReplyMarkuper struct {
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

type ReplyMarkup interface {
	IsReplyMarkup()
}

func (m *ReplyMarkuper) WithKeyboard(replyMarkup ReplyMarkup) {
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
