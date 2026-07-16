package imedia

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/media/mobjects"
)

type MediaObjectService struct {

	//starter:component

	_as func(mobjects.Service) //starter:as("#")

}

// Find implements [mobjects.Service].
func (inst *MediaObjectService) Find(cc context.Context, id mobjects.ID) (*mobjects.DTO, error) {
	panic("unimplemented")
}

// Insert implements [mobjects.Service].
func (inst *MediaObjectService) Insert(cc context.Context, item *mobjects.DTO) (*mobjects.DTO, error) {
	panic("unimplemented")
}

func (inst *MediaObjectService) _impl() mobjects.Service {
	return inst
}
