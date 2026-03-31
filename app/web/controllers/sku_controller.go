package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/skus"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type SkuController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service skus.Service     //starter:inject("#")

}

func (inst *SkuController) _impl() libgin.Controller {
	return inst
}

func (inst *SkuController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *SkuController) route(rp libgin.RouterProxy) error {

	rp = rp.For("skus")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *SkuController) handleGetOne(gc *gin.Context) {

	req := new(mySkuRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *SkuController) handleGetList(gc *gin.Context) {

	req := new(mySkuRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *SkuController) handlePutItem(gc *gin.Context) {

	req := new(mySkuRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *SkuController) handleDemo(gc *gin.Context) {

	req := new(mySkuRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

func (inst *SkuController) handlePostItem(gc *gin.Context) {

	req := new(mySkuRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

////////////////////////////////////////////////////////////////////////////////

type mySkuRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *SkuController

	q     skus.Query
	id    dxo.SkuID
	body1 vo.SKUs
	body2 vo.SKUs
}

func (inst *mySkuRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.SkuID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(skus.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *mySkuRequest) innerGetQuery(c *gin.Context, q *skus.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *mySkuRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *mySkuRequest) doGetList() error {

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

func (inst *mySkuRequest) doGetOne() error {

	it := &dto.SKU{}

	inst.body2.Items = []*dto.SKU{it}
	return nil
}

func (inst *mySkuRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.SKU{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.SKU{it1, it2}
	return nil
}

func (inst *mySkuRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.SKU{o2}
	return nil

}

func (inst *mySkuRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
