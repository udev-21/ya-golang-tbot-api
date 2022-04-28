package types

import "reflect"

type Params map[string]interface{}

func (p Params) AddNonNil(key string, value interface{}) {
	if value == nil ||
		(reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return
	}
	p[key] = value
}

func (p Params) AddNonNilNonZero(key string, value interface{}) {
	v := reflect.ValueOf(value)
	if value == nil || (v.Kind() == reflect.Ptr && v.IsNil()) || v.IsZero() {
		return
	}
	p[key] = value
}

type Sendable interface {
	Params() (Params, error)
	Endpoint() string
}

type UploadWithFiles interface {
	Sendable
	Filer
}

type Filer interface {
	Files() []Uploadable
}
