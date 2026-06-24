package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

////////////////////////////////////////////////////////////////////////////////
// a set of all holders

var theCodeNormalizerHolders innerCodeNormalizerHolders

////////////////////////////////////////////////////////////////////////////////

type innerCodeNormalizerHolders struct {
	company  normalizers.CodeNormalizerHolder
	standard normalizers.CodeNormalizerHolder
	licence  normalizers.CodeNormalizerHolder
	tm       normalizers.CodeNormalizerHolder
	product  normalizers.CodeNormalizerHolder
}

func (inst *innerCodeNormalizerHolders) normalizerForCompany() normalizers.CodeNormalizer {

	// yyzz-code : '12A34B567CDE89Y'

	h := &inst.company

	return h.GetNormalizer(func(b *normalizers.CodeNormalizerBuilder) {
		b.UseDigit().UseUpperCase()
	})
}

func (inst *innerCodeNormalizerHolders) normalizerForStandard() normalizers.CodeNormalizer {

	// todo ...
	// 'GB12345'
	// 'GB/T12345'
	// 'GB/T12345-2333'

	h := &inst.standard

	return h.GetNormalizer(func(b *normalizers.CodeNormalizerBuilder) {
		b.UseDigit().UseUpperCase().UseMarks('/', '-')
	})
}

func (inst *innerCodeNormalizerHolders) normalizerForLicence() normalizers.CodeNormalizer {

	//like examples:
	// sc: 'SC123456'
	// xk: 'XK123456'
	// yyzz: '12A34B567CDE89Y'

	h := &inst.licence

	return h.GetNormalizer(func(b *normalizers.CodeNormalizerBuilder) {
		b.UseDigit().UseUpperCase()
	})
}

func (inst *innerCodeNormalizerHolders) normalizerForTM() normalizers.CodeNormalizer {

	//  注册号 : '45619538'

	h := &inst.tm

	return h.GetNormalizer(func(b *normalizers.CodeNormalizerBuilder) {
		b.UseDigit()
	})
}

func (inst *innerCodeNormalizerHolders) normalizerForProduct() normalizers.CodeNormalizer {

	// pure digit

	h := &inst.product

	return h.GetNormalizer(func(b *normalizers.CodeNormalizerBuilder) {
		b.UseDigit()
	})
}

////////////////////////////////////////////////////////////////////////////////

func normalizeCompanyCode(code CompanyCode) CompanyCode {

	holders := &theCodeNormalizerHolders
	normalizer := holders.normalizerForCompany()

	str1 := code.String()
	str2 := normalizer.Normalize(str1)
	return CompanyCode(str2)

}

func normalizeStandardCode(code StandardCode) StandardCode {

	holders := &theCodeNormalizerHolders
	normalizer := holders.normalizerForStandard()

	str1 := code.String()
	str2 := normalizer.Normalize(str1)
	return StandardCode(str2)
}

func normalizeLicenceCode(code LicenceCode) LicenceCode {

	holders := &theCodeNormalizerHolders
	normalizer := holders.normalizerForLicence()

	str1 := code.String()
	str2 := normalizer.Normalize(str1)
	return LicenceCode(str2)

}

func normalizeTrademarkCode(code TrademarkCode) TrademarkCode {

	holders := &theCodeNormalizerHolders
	normalizer := holders.normalizerForTM()

	str1 := code.String()
	str2 := normalizer.Normalize(str1)
	return TrademarkCode(str2)
}

func normalizeProductCode(code ProductCode) ProductCode {

	holders := &theCodeNormalizerHolders
	normalizer := holders.normalizerForProduct()

	str1 := code.String()
	str2 := normalizer.Normalize(str1)
	return ProductCode(str2)

}

////////////////////////////////////////////////////////////////////////////////
// EOF
