package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type MediaLink struct {

	// id
	ID dxo.MediaLinkID `json:"id"`

	Base

	Sum  lang.Hex `json:"sum"`
	Size int64    `json:"size"`
}

type MediaObject struct {

	// id
	ID dxo.MediaObjectID `json:"id"`

	Base

	Sum  lang.Hex `json:"sum"`
	Size int64    `json:"size"`
}
