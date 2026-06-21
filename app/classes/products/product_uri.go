package products

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/data/dxo"
)

func ComputeUri(o *DTO) dxo.URI {

	uid := o.Owner
	code := o.Code.String()
	com := utils.NewEntityUriComputer()

	com.SetUser(uid).SetType("products").SetCode(code)

	return com.URI()
}
