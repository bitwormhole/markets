package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/examples"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type ExampleController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service examples.Service //starter:inject("#")

}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ExampleController) route(rp libgin.RouterProxy) error {

	rp = rp.For("examples")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *ExampleController) handleGetOne(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *ExampleController) handleGetList(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false
	req.wantRequestQuery = true

	req.execute(req.doGetList)
}

func (inst *ExampleController) handlePutItem(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *ExampleController) handlePostItem(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *ExampleController) handleDemo(gc *gin.Context) {

	req := new(myExampleRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *ExampleController

	q     examples.Query
	id    dxo.ExampleID
	body1 vo.Examples
	body2 vo.Examples
}

func (inst *myExampleRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(examples.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myExampleRequest) innerGetQuery(c *gin.Context, q *examples.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myExampleRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myExampleRequest) doGetList() error {

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

func (inst *myExampleRequest) doGetOne() error {

	it := &dto.Example{}

	inst.body2.Items = []*dto.Example{it}
	return nil
}

func (inst *myExampleRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Example{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Example{it1, it2}
	return nil
}

func (inst *myExampleRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Example{o2}
	return nil

}

func (inst *myExampleRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
