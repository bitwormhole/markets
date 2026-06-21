package companies

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

func ComputeUri(o *DTO) dxo.URI {

	uid := o.Owner
	code := o.Code.String()
	com := utils.NewEntityUriComputer()

	com.SetUser(uid).SetType("companies").SetCode(code)

	return com.URI()
}

func UpdateUri(d *DTO) {
	uri := ComputeUri(d)
	d.URI = uri
}

func ComputeUriWithEntity(o *Entity) dxo.URI {

	uid := o.Owner
	code := o.Code.String()
	com := utils.NewEntityUriComputer()

	com.SetUser(uid).SetType("companies").SetCode(code)

	return com.URI()
}

func UpdateUriWithEntity(e *Entity) {
	uri := ComputeUriWithEntity(e)
	e.URI = uri
}

////////////////////////////////////////////////////////////////////////////////
// EOF
