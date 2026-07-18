package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/mimetypes"
)

type MediaLink struct {

	// id
	ID dxo.MediaLinkID

	Base

	// fields

	Name string

	ContentType   mimetypes.Type
	ContentLength int64
	ContentSum    lang.Hex // a sha-256-sum

	TargetID   dxo.MediaObjectID
	TargetUUID lang.UUID
}
