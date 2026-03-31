package test4markets
import (
    pb1fbf5dd6 "github.com/bitwormhole/markets/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pb1fbf5dd6.ExampleCase in package:github.com/bitwormhole/markets/src/test/golang/unittestcases
//
// id:com-b1fbf5dd6711c61d-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pb1fbf5dd67_unittestcases_ExampleCase struct {
}

func (inst* pb1fbf5dd67_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b1fbf5dd6711c61d-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb1fbf5dd67_unittestcases_ExampleCase) new() any {
    return &pb1fbf5dd6.ExampleCase{}
}

func (inst* pb1fbf5dd67_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb1fbf5dd6.ExampleCase)
	nop(ie, com)

	


    return nil
}


