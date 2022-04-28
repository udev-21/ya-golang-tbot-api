package types

type DisableNotificationer struct {
	DisableNotification bool `json:"disable_notification,omitempty"`
}

func (dn *DisableNotificationer) WithDisableNotification() {
	dn.DisableNotification = true
}
