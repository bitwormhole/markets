package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/mimetypes"
)

type MediaObject struct {

	// id
	ID dxo.MediaObjectID

	Base

	// fields

	Name          string
	ContentType   mimetypes.Type
	ContentLength int64

	ContentSum lang.Hex `gorm:"unique"` // a sha-256-sum

}
