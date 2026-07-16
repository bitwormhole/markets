package imedia

import (
	"github.com/bitwormhole/markets/app/classes/media/mlinks"
	"github.com/bitwormhole/markets/app/classes/media/mobjects"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type MediaIOFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	LinkService   mlinks.Service   //starter:inject("#")
	ObjectService mobjects.Service //starter:inject("#")
}

// Fetch implements [objects.DownloadFilter].
func (inst *MediaIOFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {
	// panic("unimplemented")

	return next.Fetch(o)
}

// Put implements [objects.UploadFilter].
func (inst *MediaIOFilter) Put(o *objects.IOContext, next objects.UploadFilterChain) error {
	// panic("unimplemented")

	return next.Put(o)
}

// ListFilters implements [objects.FilterRegistry].
func (inst *MediaIOFilter) ListFilters() []*objects.FilterRegistration {

	fr := &objects.FilterRegistration{

		Label:    "MediaIOFilter",
		Priority: layers.PriorityDB,
		Class:    "",
		Enabled:  true,

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{fr}
}

func (inst *MediaIOFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
