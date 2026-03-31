package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/products"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type ProductController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service products.Service //starter:inject("#")

}

func (inst *ProductController) _impl() libgin.Controller {
	return inst
}

func (inst *ProductController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ProductController) route(rp libgin.RouterProxy) error {

	rp = rp.For("products")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.POST("", inst.handlePostOne)
	rp.PUT(":id", inst.handlePutOne)
	rp.DELETE(":id", inst.handleDeleteOne)

	return nil
}

func (inst *ProductController) handleGetOne(gc *gin.Context) {

	req := new(myProductRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *ProductController) handleGetList(gc *gin.Context) {

	req := new(myProductRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *ProductController) handlePostOne(gc *gin.Context) {

	req := new(myProductRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *ProductController) handlePutOne(gc *gin.Context) {

	req := new(myProductRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doUpdateItem)
}

func (inst *ProductController) handleDeleteOne(gc *gin.Context) {

	req := new(myProductRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doRemoveItem)
}

////////////////////////////////////////////////////////////////////////////////

type myProductRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *ProductController

	q     products.Query
	id    dxo.ProductID
	body1 vo.Products
	body2 vo.Products
}

func (inst *myProductRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.ProductID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(products.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myProductRequest) innerGetQuery(c *gin.Context, q *products.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myProductRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myProductRequest) doGetList() error {

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

func (inst *myProductRequest) doGetOne() error {

	it := &dto.Product{}

	inst.body2.Items = []*dto.Product{it}
	return nil
}

func (inst *myProductRequest) doInsertItem() error {

	it1 := inst.body1.Items[0]
	ser := inst.controller.Service
	ctx := inst.context

	it2, err := ser.Insert(ctx, it1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Product{it2}
	return nil
}

func (inst *myProductRequest) doUpdateItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Product{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Product{it1, it2}
	return nil
}

func (inst *myProductRequest) doRemoveItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Product{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Product{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
