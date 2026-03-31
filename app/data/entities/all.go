package entities

func ListAll(prefix string) []any {

	innerUpdateTableNamePrefix(prefix)

	list := make([]any, 0)

	list = append(list, new(Company))
	list = append(list, new(Licence))
	list = append(list, new(Manufacturer))
	list = append(list, new(Product))
	list = append(list, new(Shop))
	list = append(list, new(SKU))
	list = append(list, new(Standard))
	list = append(list, new(Trademark))

	return list
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string = ""

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
// EOF
