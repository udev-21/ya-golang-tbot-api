package types

type RawJsonPayloader interface {
	RawJsonPayload() (map[string]interface{}, error)
}
type ViaReplyMarkup struct {
	ReplyMarkup ReplyMarkup `json:"-"`
}

type ReplyMarkup interface {
	RawJsonPayloader
}

// replyMarkup can be either InlineKeyboardMarkup
func (m *ViaReplyMarkup) WithKeyboard(replyMarkup ReplyMarkup) {
	m.ReplyMarkup = replyMarkup
}

type ViaParseMode struct {
	ParseMode *string `json:"parse_mode,omitempty"`
}

const (
	ParseModeMarkdownV2 string = "MarkdownV2"
	ParseModeHTML       string = "HTML"
	ParseModeMarkdown   string = "Markdown"
)

func (m *ViaParseMode) WithParseModeHTML() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeHTML
}

func (m *ViaParseMode) WithParseModeParseModeMarkdownV2() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdownV2

}

func (m *ViaParseMode) WithParseModeParseModeMarkdown() {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdown
}
