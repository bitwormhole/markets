package marketdb

import (
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/starter-go/libgorm"
)

type TheGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Enabled bool //starter:inject("${datagroup.markets.enabled}")

	Alias  string //starter:inject("${datagroup.markets.alias}")
	Prefix string //starter:inject("${datagroup.markets.table-name-prefix}")
	Source string //starter:inject("${datagroup.markets.datasource}")
	URI    string //starter:inject("${datagroup.markets.uri}")

}

// Prototypes implements libgorm.Group.
func (inst *TheGroup) Prototypes() []any {
	prefix := inst.Prefix
	return entities.ListAll(prefix)
}

// Groups implements libgorm.GroupRegistry.
func (inst *TheGroup) Groups() []*libgorm.GroupRegistration {

	r1 := &libgorm.GroupRegistration{

		Enabled: inst.Enabled,

		Alias:  inst.Alias,
		Prefix: inst.Prefix,
		Source: inst.Source,
		URI:    inst.URI,

		Group: inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

func (inst *TheGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}
