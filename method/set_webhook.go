package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetWebhook(url string) *SetWebhook {
	return &SetWebhook{
		Url: url,
	}
}

// https://core.telegram.org/bots/api#setwebhook
type SetWebhook struct {
	Url            string             `json:"url"`
	Certificate    myTypes.Uploadable `json:"certificate,omitempty"`
	IpAddress      string             `json:"ip_address,omitempty"`
	MaxConnections int64              `json:"max_connections,omitempty"`
	AllowedUpdates []string           `json:"allowed_updates,omitempty"`
}

func (sw *SetWebhook) Endpoint() string {
	return "setWebhook"
}

func (sw *SetWebhook) Params() (myTypes.Params, error) {
	old, err := utils.ConvertToMapStringInterface(sw)
	if err != nil {
		return nil, err
	}
	if sw.Certificate != nil {
		old["certificate"] = sw.Certificate
	}
	return old, nil
}

func (sw *SetWebhook) Files() []myTypes.Uploadable {
	if sw.Certificate != nil {
		return []myTypes.Uploadable{sw.Certificate}
	}
	return nil
}

func (sw *SetWebhook) WithCertificate(file myTypes.Uploadable) {
	sw.Certificate = file
}

func (sw *SetWebhook) WithIpAddress(ipAddress string) {
	sw.IpAddress = ipAddress
}

func (sw *SetWebhook) WithMaxConnections(maxConnections int64) {
	sw.MaxConnections = maxConnections
}

func (sw *SetWebhook) WithAllowedUpdates(allowedUpdates []string) {
	copy(sw.AllowedUpdates, allowedUpdates)
}
