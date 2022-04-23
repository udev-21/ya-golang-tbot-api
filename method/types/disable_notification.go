package types

type DisableNotificationer struct {
	DisableNotification_ bool `json:"disable_notification,omitempty"`
}

func (dn *DisableNotificationer) DisableNotification() {
	dn.DisableNotification_ = true
}
