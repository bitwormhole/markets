package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/mimetypes"
)

type MediaLink struct {

	// id
	ID dxo.MediaLinkID `json:"id"`

	Base

	Name          string         `json:"name"`
	ContentType   mimetypes.Type `json:"content_type"`
	ContentLength int64          `json:"content_length"`
	ContentSum    lang.Hex       `json:"content_sum"`

	TargetID   dxo.MediaObjectID `json:"target_id"`
	TargetUUID lang.UUID         `json:"target_uuid"`
}

type MediaObject struct {

	// id
	ID dxo.MediaObjectID `json:"id"`

	Base

	Name          string         `json:"name"`
	ContentType   mimetypes.Type `json:"content_type"`
	ContentLength int64          `json:"content_length"`
	ContentSum    lang.Hex       `json:"content_sum"`
}
