package dxo

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/starter-go/rbac"
)

////////////////////////////////////////////////////////////////////////////////

type URI string

func (u URI) String() string {
	return string(u)
}

func (u URI) Resolve(dst *UniformResourceID) error {

	if dst == nil {
		return fmt.Errorf("URI.Resolve(dst): param 'dst' is nil")
	}

	str := u.String()
	u2, err := url.Parse(str)
	if err != nil {
		return nil
	}

	strType := u2.Hostname()
	strUser := u2.User.Username()
	strCode := u2.Path

	if strUser != "" {
		numUser, err := strconv.ParseInt(strUser, 10, 64)
		if err != nil {
			return nil
		}
		dst.User = rbac.UserID(numUser)
	}

	dst.Type = strType
	dst.Code = strCode

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type UniformResourceID struct {
	User rbac.UserID
	Type string
	Code string
}

func (inst *UniformResourceID) URI() URI {

	// like 'uri://user@type/id'

	t := inst.Type
	u := inst.User
	c := inst.Code
	b := new(strings.Builder)

	b.WriteString("uri://")
	if u > 0 {
		user := strconv.FormatInt(int64(u), 10)
		b.WriteString(user)
		b.WriteRune('@')
	}
	b.WriteString(t)
	b.WriteRune('/')
	b.WriteString(c)

	str := b.String()
	return URI(str)
}

func (inst *UniformResourceID) String() string {

	tmp := inst.URI()
	return tmp.String()
}

////////////////////////////////////////////////////////////////////////////////
