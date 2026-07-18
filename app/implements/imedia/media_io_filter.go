package imedia

import (
	"github.com/bitwormhole/markets/app/classes/media/mlinks"
	"github.com/bitwormhole/markets/app/classes/media/mobjects"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/vlog"
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

	err := next.Put(o)
	if err != nil {
		return err
	}

	err = inst.onPutOK(o)
	if err != nil {
		return err
	}

	return nil
}

func (inst *MediaIOFilter) onPutOK(o *objects.IOContext) error {

	ctx := o.CC
	mo := new(mobjects.DTO)
	ml := new(mlinks.DTO)
	src := o.GetWant(true)
	oService := inst.ObjectService

	meta := &src.Meta
	sum1 := meta.Sum.Hex()
	sum2 := sum1.Bytes()

	mo.Name = meta.Name
	mo.ContentSum = sum1
	mo.ContentLength = meta.Length
	mo.ContentType = mimetypes.Type(meta.Type)

	ml.Name = meta.Name
	ml.ContentSum = sum1
	ml.ContentLength = meta.Length
	ml.ContentType = mimetypes.Type(meta.Type)

	// find | create

	hasOlder, err := oService.ContainsSum(ctx, sum2)
	if err != nil {
		return err
	}
	if hasOlder {
		older, err := oService.FindBySum(ctx, sum2)
		if err != nil {
			return err
		}
		mo = older
	} else {
		mo1, err := oService.Insert(ctx, mo)
		if err != nil {
			return err
		}
		mo = mo1
	}

	ml.TargetID = mo.ID
	ml.TargetUUID = mo.UUID

	ml, err = inst.LinkService.Insert(ctx, ml)
	if err != nil {
		return err
	}

	if vlog.IsDebugEnabled() {
		id := ml.ID
		uuid := ml.UUID
		csize := ml.ContentLength
		ctype := ml.ContentType
		csum := ml.ContentSum
		vlog.Debug("[media id:%d uuid:'%s' sum:'%s' type:'%s' size:%d ]", id, uuid, csum, ctype, csize)
	}

	return nil
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
