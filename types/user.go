package types

// User - This object represents a Telegram user or bot.
// https://core.telegram.org/bots/api#user
type User struct {
	ID                      int64   `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
}
