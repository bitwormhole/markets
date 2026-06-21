package skus

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/data/dxo"
)

func ComputeUri(o *DTO) dxo.URI {

	code := o.Code.String()
	com := utils.NewEntityUriComputer()
	dn := o.Domain
	uid := o.Owner

	com.SetUser(uid).SetType("skus").SetCode(code).SetDomain(dn)

	return com.URI()
}
