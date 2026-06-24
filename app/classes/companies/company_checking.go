package companies

import "github.com/starter-go/v0/subjects"

func PreCheckDTO(o *DTO) *subjects.Checking {

	ch := new(subjects.Checking)

	ch.Type = "todo..."
	ch.Target = o
	ch.Owner = o.Owner

	return ch
}

func PreCheckEntity(o *Entity) *subjects.Checking {

	ch := new(subjects.Checking)

	ch.Type = "todo..."
	ch.Target = o
	ch.Owner = o.Owner

	return ch
}
