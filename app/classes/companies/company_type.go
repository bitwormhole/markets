package companies

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/web/dto"
)

type ID = dxo.CompanyID

type Code = dxo.CompanyCode

type Name = dxo.CompanyName

type Entity = entities.Company

type DTO = dto.Company
