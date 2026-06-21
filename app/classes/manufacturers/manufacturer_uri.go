package manufacturers

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/data/dxo"
)

func ComputeUri(o *DTO) dxo.URI {

	uid := o.Owner
	code := o.CompanyCode.String()
	com := utils.NewEntityUriComputer()

	com.SetUser(uid).SetType("manufacturers").SetCode(code)

	return com.URI()

}
