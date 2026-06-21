package standards

import (
	"github.com/bitwormhole/markets/app/classes/utils"
)

func ComputeUri(o *DTO) URI {

	uid := o.Owner
	code := o.Code.String()
	com := utils.NewEntityUriComputer()

	com.SetUser(uid).SetType("standards").SetCode(code)

	return com.URI()
}
