package licences

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/web/dto"
)

type ID = dxo.LicenceCode

type Entity = entities.Licence

type DTO = dto.Licence

type Type = dxo.LicenceType

////////////////////////////////////////////////////////////////////////////////

const (
	TypeYYZZ Type = dxo.LicenceYYZZ // 营业执照
	TypeSC   Type = dxo.LicenceSC   // 食品生产许可证
	TypeXK   Type = dxo.LicenceXK   // 生产许可证
	TypeWS   Type = dxo.LicenceWS   // 卫生许可证
	TypeTM   Type = dxo.LicenceTM   // 商标注册证书
)

////////////////////////////////////////////////////////////////////////////////
