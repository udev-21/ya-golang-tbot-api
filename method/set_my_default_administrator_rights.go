package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetMyDefaultAdministratorRights() *SetMyDefaultAdministratorRights {
	return &SetMyDefaultAdministratorRights{}
}

type SetMyDefaultAdministratorRights struct {
	Rights      *types.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                           `json:"for_channels,omitempty"`
}

func (smdar *SetMyDefaultAdministratorRights) Endpoint() string {
	return "setMyDefaultAdministratorRights"
}

func (smdar *SetMyDefaultAdministratorRights) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(smdar)
}

func (smdar *SetMyDefaultAdministratorRights) WithRights(rights types.ChatAdministratorRights) {
	smdar.Rights = &rights
}

func (smdar *SetMyDefaultAdministratorRights) WithForChannels() {
	smdar.ForChannels = true
}
