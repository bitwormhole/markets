package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/rbac"
)

func TryGetPagination(c *gin.Context) (*rbac.Pagination, error) {

	const (
		base    = 10
		bitSize = 60
	)
	limit := int64(0)
	offset := int64(0)
	p := new(rbac.Pagination)
	r := new(numberRange)

	strOffset, hasOffset := c.GetQuery("offset")
	strLimit, hasLimit := c.GetQuery("limit")

	if hasLimit {
		num, err := strconv.ParseInt(strLimit, base, bitSize)
		if err != nil {
			return nil, err
		}
		limit = num
	}
	limit = r.normalizeMM(limit, 1, 1024*2)

	if hasOffset {
		num, err := strconv.ParseInt(strOffset, base, bitSize)
		if err != nil {
			return nil, err
		}
		offset = num
	}
	offset = r.normalizeMM(offset, 0, 0xffffffffffffff)

	p.Size = int(limit)
	p.Page = (offset / limit) + 1

	return p, nil
}

////////////////////////////////////////////////////////////////////////////////

type numberRange struct {
	min int64
	max int64
}

func (inst *numberRange) normalize(n int64) int64 {

	a := inst.min
	b := inst.max

	if n < a {
		n = a
	}

	if n > b {
		n = b
	}

	return n
}

func (inst *numberRange) normalizeMM(n, min, max int64) int64 {

	inst.min = min
	inst.max = max
	return inst.normalize(n)
}
