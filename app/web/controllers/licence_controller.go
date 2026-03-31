package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/licences"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type LicenceController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service licences.Service //starter:inject("#")

}

func (inst *LicenceController) _impl() libgin.Controller {
	return inst
}

func (inst *LicenceController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *LicenceController) route(rp libgin.RouterProxy) error {

	rp = rp.For("licences")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *LicenceController) handleGetOne(gc *gin.Context) {

	req := new(myLicenceRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *LicenceController) handleGetList(gc *gin.Context) {

	req := new(myLicenceRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *LicenceController) handlePutItem(gc *gin.Context) {

	req := new(myLicenceRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *LicenceController) handlePostItem(gc *gin.Context) {

	req := new(myLicenceRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *LicenceController) handleDemo(gc *gin.Context) {

	req := new(myLicenceRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myLicenceRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *LicenceController

	q     licences.Query
	id    dxo.LicenceID
	body1 vo.Licences
	body2 vo.Licences
}

func (inst *myLicenceRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.LicenceID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(licences.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myLicenceRequest) innerGetQuery(c *gin.Context, q *licences.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myLicenceRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myLicenceRequest) doGetList() error {

	q := &inst.q
	ctx := inst.context
	ser := inst.controller.Service

	list, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	inst.body2.Items = list
	inst.body2.Pagination = &q.Pagination
	return nil
}

func (inst *myLicenceRequest) doGetOne() error {

	it := &dto.Licence{}

	inst.body2.Items = []*dto.Licence{it}
	return nil
}

func (inst *myLicenceRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Licence{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Licence{it1, it2}
	return nil
}

func (inst *myLicenceRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Licence{o2}
	return nil

}

func (inst *myLicenceRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
