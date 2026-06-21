package user

import (
	"strconv"

	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/bitwormhole/markets/app/web/dto"
	"github.com/bitwormhole/markets/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/rbac-web-app/app/classes/users"
)

////////////////////////////////////////////////////////////////////////////////

type UserController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service users.Service    //starter:inject("#")

}

func (inst *UserController) _impl() libgin.Controller {
	return inst
}

func (inst *UserController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *UserController) route(rp libgin.RouterProxy) error {

	rp = router4api(rp, "users")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)
	rp.POST("", inst.handlePostItem)

	return nil
}

func (inst *UserController) handleGetOne(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *UserController) handleGetList(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *UserController) handlePutItem(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

func (inst *UserController) handlePostItem(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *UserController) handleDemo(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doDemo)
}

////////////////////////////////////////////////////////////////////////////////

type myUserRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	wantRequestPage  bool
	wantRequestQuery bool

	context    *gin.Context
	controller *UserController

	q     users.Query
	id    rbac.UserID
	body1 vo.Users
	body2 vo.Users
}

func (inst *myUserRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = rbac.UserID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		q := new(users.Query)
		err := inst.innerGetQuery(ctx, q)
		if err != nil {
			return err
		}
		inst.q = *q
	}

	return nil
}

func (inst *myUserRequest) innerGetQuery(c *gin.Context, q *users.Query) error {

	pg, err := TryGetPagination(c)
	if err != nil {
		return err
	}
	q.Pagination = *pg

	// todo ... more fields

	return nil
}

func (inst *myUserRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myUserRequest) doGetList() error {

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

func (inst *myUserRequest) doGetOne() error {

	it := &dto.User{}

	inst.body2.Items = []*dto.User{it}
	return nil
}

func (inst *myUserRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.User{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.User{it1, it2}
	return nil
}

func (inst *myUserRequest) doInsertItem() error {

	ser := inst.controller.Service
	ctx := inst.context
	o1 := inst.body1.Items[0]

	subHolder := utils.NewSubjectHolder(ctx).UseChecker()
	subHolder.Checker().AcceptAdmin()

	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.User{o2}
	return nil
}

func (inst *myUserRequest) doDemo() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
