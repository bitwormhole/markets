package utils

import (
	"strings"

	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/rbac"
)

// func ComputeEntityUri(enType string, code string, uid rbac.UserID) dxo.URI {

// 	b := new(strings.Builder)

// 	b.WriteString("uri://")
// 	b.WriteString(uid.String())
// 	b.WriteString("@")
// 	b.WriteString(enType)
// 	b.WriteString("/")
// 	b.WriteString(code)

// 	str := b.String()
// 	return dxo.URI(str)
// }

func NewEntityUriComputer() *EntityUriComputer {
	com := new(EntityUriComputer)
	com.domain = "markets"
	return com
}

////////////////////////////////////////////////////////////////////////////////

type EntityUriComputer struct {
	code   string
	eType  string
	domain dxo.DomainName
	uid    rbac.UserID
}

func (inst *EntityUriComputer) SetType(t string) *EntityUriComputer {
	inst.eType = t
	return inst
}

func (inst *EntityUriComputer) SetDomain(dn dxo.DomainName) *EntityUriComputer {
	inst.domain = dn
	return inst
}

func (inst *EntityUriComputer) SetCode(code string) *EntityUriComputer {
	inst.code = code
	return inst
}

func (inst *EntityUriComputer) SetUser(uid rbac.UserID) *EntityUriComputer {
	inst.uid = uid
	return inst
}

func (inst *EntityUriComputer) URI() dxo.URI {

	b := new(strings.Builder)
	dn := inst.domain
	code := inst.code
	ty := inst.eType
	uid := inst.uid.String()

	if dn == "" {
		dn = "markets"
	}
	if ty == "" {
		ty = "untyped"
	}
	if code == "" {
		code = "na"
	}

	b.WriteString("uri://")
	b.WriteString(uid)
	b.WriteString("@")
	b.WriteString(dn.String())
	b.WriteString("/")
	b.WriteString(ty)
	b.WriteString("/")
	b.WriteString(code)

	str := strings.ToLower(b.String())
	return dxo.URI(str)
}
