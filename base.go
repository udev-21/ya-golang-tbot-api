package golangtbotapi

type HandlerFunc func(Context) error

type Middleware func(HandlerFunc) HandlerFunc
