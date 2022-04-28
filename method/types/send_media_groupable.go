package types

type SendMediaGroupable interface {
	IsSendMediaGroupable()
}

type BaseSendMediaGroupable struct{}

func (b *BaseSendMediaGroupable) IsSendMediaGroupable() {}
