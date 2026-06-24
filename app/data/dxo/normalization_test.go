package dxo

import (
	"testing"
)

func TestNormalizeCompanyCode(t *testing.T) {

	list := []CompanyCode{
		"",
		"0A1B2C3",
		"0a1b2c3d4",
		"0a1 b2c/3d4",
		"//123456++",
		"/xyz",
		"/XYZ",
	}

	for index, c1 := range list {
		c2 := c1.Normalize()
		t.Logf(" normalize: CompanyCode [%d]: %s   ==>  %s  \n", index, c1, c2)
	}

}

func TestNormalizeLicenceCode(t *testing.T) {

	list := []LicenceCode{
		"",
		"0A1B2C3",
		"0a1b2c3d4",
		"0a1 b2c/3d4",
		"//123456++",
		"/xyz",
		"/XYZ",
	}

	for index, c1 := range list {
		c2 := c1.Normalize()
		t.Logf(" normalize: CompanyCode [%d]: %s   ==>  %s  \n", index, c1, c2)
	}

}

func TestNormalizeStandardCode(t *testing.T) {

	list := []StandardCode{
		"",
		"GB 1234",
		"GB/T 6789",
		"GB/T 6789-2333",
		"0a1 b2c/3d4",
		"//123456++",
		"GB / T 6789 -  2333",
		"/xyz",
		"/XYZ",
	}

	for index, c1 := range list {
		c2 := c1.Normalize()
		t.Logf(" normalize: CompanyCode [%d]: %s   ==>  %s  \n", index, c1, c2)
	}

}

func TestNormalizeTrademarkCode(t *testing.T) {

	list := []TrademarkCode{
		"",
		"GB 1234",
		"GB/T 6789",
		"GB/T 6789-2333",
		"0a1 b2c/3d4",
		"//123456++",
		"GB / T 6789 -  2333",
		"/xyz",
		"/XYZ",
	}

	for index, c1 := range list {
		c2 := c1.Normalize()
		t.Logf(" normalize: CompanyCode [%d]: %s   ==>  %s  \n", index, c1, c2)
	}

}

func TestNormalizeProductCode(t *testing.T) {

	list := []ProductCode{
		"",
		"GB 1234",
		"GB/T 6789",
		"GB/T 6789-2333",
		"0a1 b2c/3d4",
		"//123456++",
		"GB / T 6789 -  2333",
		"/xyz",
		"/XYZ",
	}

	for index, c1 := range list {
		c2 := c1.Normalize()
		t.Logf(" normalize: CompanyCode [%d]: %s   ==>  %s  \n", index, c1, c2)
	}

}
