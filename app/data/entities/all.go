package entities

func ListAll(prefix string) []any {

	innerUpdateTableNamePrefix(prefix)

	list := make([]any, 0)

	list = append(list, new(Company))
	list = append(list, new(Domain))
	list = append(list, new(Licence))
	list = append(list, new(Manufacturer))
	list = append(list, new(Product))
	list = append(list, new(Shop))
	list = append(list, new(SKU))
	list = append(list, new(Standard))
	list = append(list, new(Trademark))

	// list = append(list, new(UserCompany))
	// list = append(list, new(UserLicence))
	// list = append(list, new(UserManufacturer))
	// list = append(list, new(UserProduct))
	// list = append(list, new(UserShop))
	// list = append(list, new(UserSKU))
	// list = append(list, new(UserStandard))
	// list = append(list, new(UserTrademark))

	return list
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string = ""

var theUserTableNameTag string = "user_"

func innerUpdateTableNamePrefix(prefix string) {

	if prefix == "" {
		return
	}

	if theTableNamePrefix == "" {
		theTableNamePrefix = prefix
	}

}

////////////////////////////////////////////////////////////////////////////////

func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}

func (Company) TableName() string {
	return theTableNamePrefix + "companies"
}

func (Domain) TableName() string {
	return theTableNamePrefix + "domains"
}

func (Licence) TableName() string {
	return theTableNamePrefix + "licences"
}

func (Manufacturer) TableName() string {
	return theTableNamePrefix + "manufacturers"
}

func (Product) TableName() string {
	return theTableNamePrefix + "products"
}

func (Shop) TableName() string {
	return theTableNamePrefix + "shops"
}

func (SKU) TableName() string {
	return theTableNamePrefix + "skus"
}

func (Standard) TableName() string {
	return theTableNamePrefix + "standards"
}

func (Trademark) TableName() string {
	return theTableNamePrefix + "trademarks"
}

////////////////////////////////////////////////////////////////////////////////

// func (UserCompany) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "companies"
// }

// func (UserLicence) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "licences"
// }

// func (UserManufacturer) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "manufacturers"
// }

// func (UserProduct) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "products"
// }

// func (UserShop) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "shops"
// }

// func (UserSKU) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "skus"
// }

// func (UserStandard) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "standards"
// }

// func (UserTrademark) TableName() string {
// 	return theTableNamePrefix + theUserTableNameTag + "trademarks"
// }

////////////////////////////////////////////////////////////////////////////////
// EOF
