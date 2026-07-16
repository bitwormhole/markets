package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type MediaObject struct {

	// id
	ID dxo.MediaObjectID

	Base

	// fields

	Size int64
	Sum  lang.Hex // a sha-256-sum
}
