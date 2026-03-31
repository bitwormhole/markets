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
    inst.register(&p19c8c70a5b_imanufacturer_ManufacturerDaoImpl{})
    inst.register(&p19c8c70a5b_imanufacturer_ManufacturerServiceImpl{})
    inst.register(&p48d083251a_controllers_CompanyController{})
    inst.register(&p48d083251a_controllers_ExampleController{})
    inst.register(&p48d083251a_controllers_LicenceController{})
    inst.register(&p48d083251a_controllers_ManufacturerController{})
    inst.register(&p48d083251a_controllers_ProductController{})
    inst.register(&p48d083251a_controllers_ShopController{})
    inst.register(&p48d083251a_controllers_SkuController{})
    inst.register(&p48d083251a_controllers_StandardController{})
    inst.register(&p48d083251a_controllers_TrademarkController{})
    inst.register(&p4928c59769_iexamples_ExampleDaoImpl{})
    inst.register(&p4928c59769_iexamples_ExampleServiceImpl{})
    inst.register(&p4adf7a7b64_marketdb_TheGroup{})
    inst.register(&p4da71880f9_idb_MarketDBAgent{})
    inst.register(&p8b634539d1_istandard_StandardDaoImpl{})
    inst.register(&p8b634539d1_istandard_StandardServiceImpl{})
    inst.register(&pd0d159fec5_ishop_ShopDaoImpl{})
    inst.register(&pd0d159fec5_ishop_ShopServiceImpl{})
    inst.register(&pdab695db7d_iproduct_ProductDaoImpl{})
    inst.register(&pdab695db7d_iproduct_ProductServiceImpl{})
    inst.register(&pdb5a6eac05_icompany_CompanyDaoImpl{})
    inst.register(&pdb5a6eac05_icompany_CompanyServiceImpl{})
    inst.register(&pdfe7f411a5_isku_SkuDaoImpl{})
    inst.register(&pdfe7f411a5_isku_SkuServiceImpl{})
    inst.register(&pecc835d0da_itrademark_TrademarkDaoImpl{})
    inst.register(&pecc835d0da_itrademark_TrademarkServiceImpl{})


    return nil
}
