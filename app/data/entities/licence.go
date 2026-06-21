package entities

import (
	"time"

	"github.com/bitwormhole/markets/app/data/dxo"
)

type BaseLicence struct {

	// id
	// ID dxo.LicenceID

	Base

	// fields

	Code dxo.LicenceCode

	Type dxo.LicenceType

	// HolderID   dxo.CompanyID
	// HolderCode dxo.CompanyCode
	// HolderName dxo.CompanyName

	IssuerName    string
	IssuerAddress string

	SubjectAddress string
	SubjectName    dxo.CompanyName
	SubjectID      dxo.CompanyID
	SubjectCode    dxo.CompanyCode

	NotBefore time.Time
	NotAfter  time.Time

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	URI dxo.URI `gorm:"unique"`

	Remarks string
}

type Licence struct {

	// id
	ID dxo.LicenceID

	BaseLicence
}

// type UserLicence struct {
// 	// id
// 	ID dxo.UserLicenceID
// 	BaseLicence
// }
