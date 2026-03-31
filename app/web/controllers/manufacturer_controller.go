package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/manufacturers"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type ManufacturerController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder      //starter:inject("#")
	Service manufacturers.Service //starter:inject("#")

}

func (inst *ManufacturerController) _impl() libgin.Controller {
	return inst
}

func (inst *ManufacturerController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ManufacturerController) route(rp libgin.RouterProxy) error {

	rp = rp.For("manufacturers")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *ManufacturerController) handleGetOne(gc *gin.Context) {

	req := new(myManufacturerRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *ManufacturerController) handleGetList(gc *gin.Context) {

	req := new(myManufacturerRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *ManufacturerController) handlePutItem(gc *gin.Context) {

	req := new(myManufacturerRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *ManufacturerController) handlePostItem(gc *gin.Context) {

	req := new(myManufacturerRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *ManufacturerController) handleDemo(gc *gin.Context) {

	req := new(myManufacturerRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myManufacturerRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *ManufacturerController

	q     manufacturers.Query
	id    dxo.ManufacturerID
	body1 vo.Manufacturers
	body2 vo.Manufacturers
}

func (inst *myManufacturerRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.ManufacturerID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(manufacturers.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myManufacturerRequest) innerGetQuery(c *gin.Context, q *manufacturers.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myManufacturerRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myManufacturerRequest) doGetList() error {

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

func (inst *myManufacturerRequest) doGetOne() error {

	it := &dto.Manufacturer{}

	inst.body2.Items = []*dto.Manufacturer{it}
	return nil
}

func (inst *myManufacturerRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Manufacturer{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Manufacturer{it1, it2}
	return nil
}

func (inst *myManufacturerRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Manufacturer{o2}
	return nil

}

func (inst *myManufacturerRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
