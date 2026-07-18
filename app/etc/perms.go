package etc

import (
	"net/http"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security/permissions"
)

type FastPermissions struct {

	//starter:component

	_as func(permissions.Registry) //starter:as(".")

}

// ListRegistrations implements [permissions.Registry].
func (inst *FastPermissions) ListRegistrations() []*permissions.Registration {

	b := new(innerPermListBuilder)

	b.Add(http.MethodGet, "/api/v1/my/products/:id", "admin,root,owner,any,anonym")
	b.Add(http.MethodPut, "/api/v1/my/products/:id", "admin,root,owner,any,anonym")
	b.Add(http.MethodDelete, "/api/v1/my/products/:id", "admin,root,owner,any,anonym")

	b.Add(http.MethodGet, "/api/v1/sessions/current", "admin,root,any,all,user,anonym")
	b.Add(http.MethodPost, "/api/v1/auth/login", "admin,root,any,all,anonym")

	return b.Build()
}

func (inst *FastPermissions) _impl() permissions.Registry {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerPermListBuilder struct {
	list []*permissions.Registration
}

func (inst *innerPermListBuilder) Add(method string, path string, roles rbac.RoleNameList) {

	r1 := &permissions.Registration{}
	r1.Method = method
	r1.Path = path
	r1.Enabled = true
	r1.Priority = 999
	r1.Roles = roles

	inst.list = append(inst.list, r1)
}

func (inst *innerPermListBuilder) Build() []*permissions.Registration {
	return inst.list
}

////////////////////////////////////////////////////////////////////////////////
// EOF
