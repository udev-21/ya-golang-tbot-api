package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewDeleteMyCommands() *DeleteMyCommands {
	return &DeleteMyCommands{}
}

type DeleteMyCommands struct {
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string               `json:"language_code,omitempty"`
}

func (smc *DeleteMyCommands) Endpoint() string {
	return "deleteMyCommands"
}

func (smc *DeleteMyCommands) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(smc)
}

func (smc *DeleteMyCommands) WithScope(scope types.BotCommandScope) {
	smc.Scope = scope
}

func (smc *DeleteMyCommands) WithLanguageCode(langCode string) {
	smc.LanguageCode = &langCode
}
