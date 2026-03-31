package webunit

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

type WebMessage struct {
	Headers map[string]string

	ContentType string

	Body []byte
}

////////////////////////////////////////////////////////////////////////////////

type WebRequest struct {
	WebMessage

	Method string

	Location string
}

////////////////////////////////////////////////////////////////////////////////

type WebResponse struct {
	WebMessage

	StatusCode int

	StatusText string
}

////////////////////////////////////////////////////////////////////////////////

type WebUnitRunner struct {
	method string
}

func (inst *WebUnitRunner) Run() error {

	return fmt.Errorf("no")
}

func (inst *WebUnitRunner) RunT(t *testing.T) {

	err := inst.Run()
	if err != nil {
		t.Error(err)
	}

}
