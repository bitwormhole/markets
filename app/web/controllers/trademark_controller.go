package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/trademarks"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type TrademarkController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder   //starter:inject("#")
	Service trademarks.Service //starter:inject("#")

}

func (inst *TrademarkController) _impl() libgin.Controller {
	return inst
}

func (inst *TrademarkController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *TrademarkController) route(rp libgin.RouterProxy) error {

	rp = rp.For("trademarks")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *TrademarkController) handleGetOne(gc *gin.Context) {

	req := new(myTrademarkRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *TrademarkController) handleGetList(gc *gin.Context) {

	req := new(myTrademarkRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *TrademarkController) handlePutItem(gc *gin.Context) {

	req := new(myTrademarkRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *TrademarkController) handlePostItem(gc *gin.Context) {

	req := new(myTrademarkRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *TrademarkController) handleDemo(gc *gin.Context) {

	req := new(myTrademarkRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myTrademarkRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *TrademarkController

	q     trademarks.Query
	id    dxo.TrademarkID
	body1 vo.Trademarks
	body2 vo.Trademarks
}

func (inst *myTrademarkRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.TrademarkID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(trademarks.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myTrademarkRequest) innerGetQuery(c *gin.Context, q *trademarks.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myTrademarkRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myTrademarkRequest) doGetList() error {

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

func (inst *myTrademarkRequest) doGetOne() error {

	it := &dto.Trademark{}

	inst.body2.Items = []*dto.Trademark{it}
	return nil
}

func (inst *myTrademarkRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Trademark{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Trademark{it1, it2}
	return nil
}

func (inst *myTrademarkRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Trademark{o2}
	return nil
}

func (inst *myTrademarkRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
