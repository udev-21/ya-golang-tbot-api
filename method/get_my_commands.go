package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewGetMyCommands() *GetMyCommands {
	return &GetMyCommands{}
}

type GetMyCommands struct {
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string               `json:"language_code,omitempty"`
}

func (gmc *GetMyCommands) Endpoint() string {
	return "getMyCommands"
}

func (gmc *GetMyCommands) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(gmc)
}

func (gmc *GetMyCommands) WithScope(scope types.BotCommandScope) {
	gmc.Scope = scope
}

func (gmc *GetMyCommands) WithLanguageCode(langCode string) {
	gmc.LanguageCode = &langCode
}
