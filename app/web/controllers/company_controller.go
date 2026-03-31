package controllers

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/companies"
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type CompanyController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder  //starter:inject("#")
	Service companies.Service //starter:inject("#")

}

func (inst *CompanyController) _impl() libgin.Controller {
	return inst
}

func (inst *CompanyController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *CompanyController) route(rp libgin.RouterProxy) error {

	rp = rp.For("companies")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *CompanyController) handleGetOne(gc *gin.Context) {

	req := new(myCompanyRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *CompanyController) handleGetList(gc *gin.Context) {

	req := new(myCompanyRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false
	req.wantRequestQuery = true

	req.execute(req.doGetList)
}

func (inst *CompanyController) handlePutItem(gc *gin.Context) {

	req := new(myCompanyRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *CompanyController) handlePostItem(gc *gin.Context) {

	req := new(myCompanyRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *CompanyController) handleDemo(gc *gin.Context) {

	req := new(myCompanyRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myCompanyRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *CompanyController

	q     companies.Query
	id    dxo.CompanyID
	body1 vo.Companies
	body2 vo.Companies
}

func (inst *myCompanyRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.CompanyID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(companies.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myCompanyRequest) innerGetQuery(c *gin.Context, q *companies.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myCompanyRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myCompanyRequest) doGetList() error {

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

func (inst *myCompanyRequest) doGetOne() error {

	it := &dto.Company{}

	inst.body2.Items = []*dto.Company{it}
	return nil
}

func (inst *myCompanyRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Company{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Company{it1, it2}
	return nil
}

func (inst *myCompanyRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Company{o2}
	return nil

}

func (inst *myCompanyRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
