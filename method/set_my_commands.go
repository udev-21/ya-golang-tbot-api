package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetMyCommands(commands []types.BotCommand) *SetMyCommands {
	return &SetMyCommands{
		Commands: commands,
	}
}

// https://core.telegram.org/bots/api#setmycommands
type SetMyCommands struct {
	Commands     []types.BotCommand    `json:"commands"`
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string               `json:"language_code,omitempty"`
}

func (smc *SetMyCommands) Endpoint() string {
	return "setMyCommands"
}

func (smc *SetMyCommands) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(smc)
}

func (smc *SetMyCommands) WithScope(scope types.BotCommandScope) {
	smc.Scope = scope
}

func (smc *SetMyCommands) WithLanguageCode(langCode string) {
	smc.LanguageCode = &langCode
}
