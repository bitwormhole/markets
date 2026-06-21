package standards

import (
	"reflect"

	"github.com/starter-go/v0/subjects"
)

func NewChecking(o *DTO) *subjects.Checking {

	ch := new(subjects.Checking)
	ty := reflect.TypeOf(o)

	ch.Owner = o.Owner
	ch.Target = o
	ch.Type = ty.String()

	return ch
}
