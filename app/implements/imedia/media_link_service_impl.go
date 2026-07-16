package imedia

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/media/mlinks"
)

type MediaLinkService struct {

	//starter:component

	_as func(mlinks.Service) //starter:as("#")

}

// Find implements [mlinks.Service].
func (inst *MediaLinkService) Find(cc context.Context, id mlinks.ID) (*mlinks.DTO, error) {
	panic("unimplemented")
}

// Insert implements [mlinks.Service].
func (inst *MediaLinkService) Insert(cc context.Context, item *mlinks.DTO) (*mlinks.DTO, error) {
	panic("unimplemented")
}

func (inst *MediaLinkService) _impl() mlinks.Service {
	return inst
}
