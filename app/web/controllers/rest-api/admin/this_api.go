package admin

import (
	"github.com/bitwormhole/markets/app/web/controllers"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
)

func router4api(rp libgin.RouterProxy, path string) libgin.RouterProxy {

	return rp.For("/api/v1/admin/" + path)
}

func TryGetPagination(c *gin.Context) (*rbac.Pagination, error) {

	return controllers.TryGetPagination(c)
}
