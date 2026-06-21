package utils

import (
	"context"
	"errors"

	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

// func GetSubjectGetter(ctx context.Context) (subjects.Getter, error) {
// 	sub, err := subjects.GetCurrent(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sub.DoGet()
// }

func NewSubjectHolder(ctx context.Context) *SubjectHolder {
	sh := new(SubjectHolder)
	return sh.init(ctx)
}

////////////////////////////////////////////////////////////////////////////////

type innerSubjectCache struct {
	getter  subjects.Getter
	setter  subjects.Setter
	checker subjects.Checker
	sub     subjects.Subject
	uid     rbac.UserID
}

////////////////////////////////////////////////////////////////////////////////

type SubjectHolder struct {
	context context.Context

	useSetter  bool
	useGetter  bool
	useChecker bool

	errlist []error

	cache *innerSubjectCache
}

func (inst *SubjectHolder) UseGetter() *SubjectHolder {
	inst.useGetter = true
	return inst
}

func (inst *SubjectHolder) UseSetter() *SubjectHolder {
	inst.useSetter = true
	return inst
}

func (inst *SubjectHolder) UseChecker() *SubjectHolder {
	inst.useChecker = true
	return inst
}

func (inst *SubjectHolder) Subject() subjects.Subject {
	c := inst.getCacheReq()
	ret := c.sub
	if ret == nil {
		panic("SubjectHolder: subject is nil")
	}
	return ret
}

func (inst *SubjectHolder) Checker() subjects.Checker {
	c := inst.getCacheReq()
	ret := c.checker
	if ret == nil {
		panic("SubjectHolder: checker is nil")
	}
	return ret
}

func (inst *SubjectHolder) Getter() subjects.Getter {
	c := inst.getCacheReq()
	ret := c.getter
	if ret == nil {
		panic("SubjectHolder: getter is nil")
	}
	return ret
}

func (inst *SubjectHolder) Setter() subjects.Setter {
	c := inst.getCacheReq()
	ret := c.setter
	if ret == nil {
		panic("SubjectHolder: setter is nil")
	}
	return ret
}

func (inst *SubjectHolder) UID() rbac.UserID {
	c := inst.getCacheReq()
	return c.uid
}

func (inst *SubjectHolder) load() (*innerSubjectCache, error) {

	ctx := inst.context
	c := new(innerSubjectCache)

	sub, err := subjects.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}

	if inst.useChecker {
		checker, err := sub.DoCheck()
		if err != nil {
			return nil, err
		}
		c.checker = checker
	}

	if inst.useGetter {
		getter, err := sub.DoGet()
		if err != nil {
			return nil, err
		}
		uid := getter.GetUserID()
		c.getter = getter
		c.uid = uid
	}

	if inst.useSetter {
		setter, err := sub.DoSet()
		if err != nil {
			return nil, err
		}
		c.setter = setter
	}

	c.sub = sub

	return c, nil
}

func (inst *SubjectHolder) getCache() (*innerSubjectCache, error) {
	c := inst.cache
	if c == nil {
		c2, err := inst.load()
		if err != nil {
			return nil, err
		}
		c = c2
		inst.cache = c2
	}
	return c, nil
}

func (inst *SubjectHolder) getCacheReq() *innerSubjectCache {
	c, err := inst.getCache()
	if err != nil {
		panic(err)
	}
	return c
}

func (inst *SubjectHolder) Error() error {
	all := inst.errlist
	count := len(all)
	if count < 1 {
		return nil
	} else if count == 1 {
		return all[0]
	}
	return errors.Join(all...)
}

func (inst *SubjectHolder) init(ctx context.Context) *SubjectHolder {
	inst.context = ctx
	return inst
}

func (inst *SubjectHolder) addErr(err error) {
	if err == nil {
		return
	}
	inst.errlist = append(inst.errlist, err)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
