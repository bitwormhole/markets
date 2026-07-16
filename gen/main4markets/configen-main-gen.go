package main4markets

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p03b29322a1_ilicence_LicenceDaoImpl{})
    inst.register(&p03b29322a1_ilicence_LicenceServiceImpl{})
    inst.register(&p0e0d72f717_imedia_MediaIOFilter{})
    inst.register(&p0e0d72f717_imedia_MediaLinkDao{})
    inst.register(&p0e0d72f717_imedia_MediaLinkService{})
    inst.register(&p0e0d72f717_imedia_MediaObjectDao{})
    inst.register(&p0e0d72f717_imedia_MediaObjectService{})
    inst.register(&p19c8c70a5b_imanufacturer_ManufacturerDaoImpl{})
    inst.register(&p19c8c70a5b_imanufacturer_ManufacturerServiceImpl{})
    inst.register(&p4928c59769_iexamples_ExampleDaoImpl{})
    inst.register(&p4928c59769_iexamples_ExampleServiceImpl{})
    inst.register(&p4adf7a7b64_marketdb_TheGroup{})
    inst.register(&p4da71880f9_idb_MarketDBAgent{})
    inst.register(&p8b634539d1_istandard_StandardDaoImpl{})
    inst.register(&p8b634539d1_istandard_StandardServiceImpl{})
    inst.register(&p97fb530732_admin_CompanyController{})
    inst.register(&p97fb530732_admin_ExampleController{})
    inst.register(&p97fb530732_admin_LicenceController{})
    inst.register(&p97fb530732_admin_ManufacturerController{})
    inst.register(&p97fb530732_admin_ProductController{})
    inst.register(&p97fb530732_admin_ShopController{})
    inst.register(&p97fb530732_admin_SkuController{})
    inst.register(&p97fb530732_admin_StandardController{})
    inst.register(&p97fb530732_admin_TrademarkController{})
    inst.register(&p97fb530732_admin_UserController{})
    inst.register(&pd0d159fec5_ishop_ShopDaoImpl{})
    inst.register(&pd0d159fec5_ishop_ShopServiceImpl{})
    inst.register(&pdab695db7d_iproduct_ProductDaoImpl{})
    inst.register(&pdab695db7d_iproduct_ProductServiceImpl{})
    inst.register(&pdb5a6eac05_icompany_CompanyDaoImpl{})
    inst.register(&pdb5a6eac05_icompany_CompanyServiceImpl{})
    inst.register(&pdfe7f411a5_isku_SkuDaoImpl{})
    inst.register(&pdfe7f411a5_isku_SkuServiceImpl{})
    inst.register(&pe12b812cec_user_CompanyController{})
    inst.register(&pe12b812cec_user_ExampleController{})
    inst.register(&pe12b812cec_user_LicenceController{})
    inst.register(&pe12b812cec_user_ManufacturerController{})
    inst.register(&pe12b812cec_user_ProductController{})
    inst.register(&pe12b812cec_user_ShopController{})
    inst.register(&pe12b812cec_user_SkuController{})
    inst.register(&pe12b812cec_user_StandardController{})
    inst.register(&pe12b812cec_user_TrademarkController{})
    inst.register(&pe12b812cec_user_UserController{})
    inst.register(&pecc835d0da_itrademark_TrademarkDaoImpl{})
    inst.register(&pecc835d0da_itrademark_TrademarkServiceImpl{})


    return nil
}
