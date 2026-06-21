package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type Licence struct {

	// id
	ID dxo.LicenceID `json:"id"`

	Base

	Code      dxo.LicenceCode `json:"code"`
	Type      dxo.LicenceType `json:"type"`
	Reference dxo.URL         `json:"reference"`
	URI       dxo.URI         `json:"uri"`

	NotBefore lang.Time `json:"not_before"`
	NotAfter  lang.Time `json:"not_after"`

	Remarks string `json:"remarks"`

	// issuer

	IssuerAddress string `json:"issuer_address"`
	IssuerName    string `json:"issuer_name"`

	// subject

	SubjectAddress string          `json:"subject_address"`
	SubjectCode    dxo.CompanyCode `json:"subject_code"`
	SubjectName    dxo.CompanyName `json:"subject_name"`
}

// func (inst *Licence) Complete() *Licence {
// 	inst.URI =
// 	return inst
// }
