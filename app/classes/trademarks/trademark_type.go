package trademarks

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/web/dto"
)

type ID = dxo.TrademarkID

type Entity = entities.Trademark

// type UE = entities.UserTrademark // the user_trademark_entity

type DTO = dto.Trademark
