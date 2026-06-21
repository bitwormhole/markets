package standards

import (
	"testing"
)

func TestChecking(t *testing.T) {

	obj := new(DTO)
	ch := NewChecking(obj)

	t.Logf("type = %s", ch.Type)

}
