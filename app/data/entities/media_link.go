package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type MediaLink struct {

	// id
	ID dxo.MediaLinkID

	Base

	// fields

	Size int64

	Sum lang.Hex // a sha-256-sum

	TargetID dxo.MediaObjectID

	TargetUUID lang.UUID
}
