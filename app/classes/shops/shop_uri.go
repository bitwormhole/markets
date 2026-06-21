package shops

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/data/dxo"
)

func ComputeUri(o *DTO) dxo.URI {

	uid := o.Owner
	code := o.Name
	com := utils.NewEntityUriComputer()
	dn := o.Domain

	com.SetUser(uid).SetType("shops").SetCode(code).SetDomain(dn)

	return com.URI()
}
