package golangtbotapi

import "fmt"

func newError(err string) error {
	return fmt.Errorf("golangtbotapi: %s", err)
}
