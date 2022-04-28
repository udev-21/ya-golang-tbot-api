package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewGetMyDefaultAdministratorRights() *GetMyDefaultAdministratorRights {
	return &GetMyDefaultAdministratorRights{}
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

func (smdar *GetMyDefaultAdministratorRights) Endpoint() string {
	return "getMyDefaultAdministratorRights"
}

func (smdar *GetMyDefaultAdministratorRights) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(smdar)
}

func (smdar *GetMyDefaultAdministratorRights) WithForChannels() {
	smdar.ForChannels = true
}
