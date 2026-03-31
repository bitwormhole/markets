package main4markets
import (
    p006e7bf70 "github.com/bitwormhole/markets/app/classes/companies"
    p18a88e167 "github.com/bitwormhole/markets/app/classes/examples"
    pe29ea031d "github.com/bitwormhole/markets/app/classes/licences"
    pfa6c9190d "github.com/bitwormhole/markets/app/classes/manufacturers"
    p4cebd7daf "github.com/bitwormhole/markets/app/classes/products"
    pb1845fb61 "github.com/bitwormhole/markets/app/classes/shops"
    p1c4227d85 "github.com/bitwormhole/markets/app/classes/skus"
    pb395bddcd "github.com/bitwormhole/markets/app/classes/standards"
    p1142a8d2a "github.com/bitwormhole/markets/app/classes/trademarks"
    p4adf7a7b6 "github.com/bitwormhole/markets/app/data/marketdb"
    pdb5a6eac0 "github.com/bitwormhole/markets/app/implements/icompany"
    p4da71880f "github.com/bitwormhole/markets/app/implements/idb"
    p4928c5976 "github.com/bitwormhole/markets/app/implements/iexamples"
    p03b29322a "github.com/bitwormhole/markets/app/implements/ilicence"
    p19c8c70a5 "github.com/bitwormhole/markets/app/implements/imanufacturer"
    pdab695db7 "github.com/bitwormhole/markets/app/implements/iproduct"
    pd0d159fec "github.com/bitwormhole/markets/app/implements/ishop"
    pdfe7f411a "github.com/bitwormhole/markets/app/implements/isku"
    p8b634539d "github.com/bitwormhole/markets/app/implements/istandard"
    pecc835d0d "github.com/bitwormhole/markets/app/implements/itrademark"
    p48d083251 "github.com/bitwormhole/markets/app/web/controllers"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type p4adf7a7b6.TheGroup in package:github.com/bitwormhole/markets/app/data/marketdb
//
// id:com-4adf7a7b6427a7b2-marketdb-TheGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p4adf7a7b64_marketdb_TheGroup struct {
}

func (inst* p4adf7a7b64_marketdb_TheGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4adf7a7b6427a7b2-marketdb-TheGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4adf7a7b64_marketdb_TheGroup) new() any {
    return &p4adf7a7b6.TheGroup{}
}

func (inst* p4adf7a7b64_marketdb_TheGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4adf7a7b6.TheGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.URI = inst.getURI(ie)


    return nil
}


func (inst*p4adf7a7b64_marketdb_TheGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.markets.enabled}")
}


func (inst*p4adf7a7b64_marketdb_TheGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.markets.alias}")
}


func (inst*p4adf7a7b64_marketdb_TheGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.markets.table-name-prefix}")
}


func (inst*p4adf7a7b64_marketdb_TheGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.markets.datasource}")
}


func (inst*p4adf7a7b64_marketdb_TheGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.markets.uri}")
}



// type pdb5a6eac0.CompanyDaoImpl in package:github.com/bitwormhole/markets/app/implements/icompany
//
// id:com-db5a6eac059e2d06-icompany-CompanyDaoImpl
// class:
// alias:alias-006e7bf7089b23af45a893677ef0db44-DAO
// scope:singleton
//
type pdb5a6eac05_icompany_CompanyDaoImpl struct {
}

func (inst* pdb5a6eac05_icompany_CompanyDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db5a6eac059e2d06-icompany-CompanyDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-006e7bf7089b23af45a893677ef0db44-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb5a6eac05_icompany_CompanyDaoImpl) new() any {
    return &pdb5a6eac0.CompanyDaoImpl{}
}

func (inst* pdb5a6eac05_icompany_CompanyDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb5a6eac0.CompanyDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pdb5a6eac05_icompany_CompanyDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*pdb5a6eac05_icompany_CompanyDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pdb5a6eac0.CompanyServiceImpl in package:github.com/bitwormhole/markets/app/implements/icompany
//
// id:com-db5a6eac059e2d06-icompany-CompanyServiceImpl
// class:
// alias:alias-006e7bf7089b23af45a893677ef0db44-Service
// scope:singleton
//
type pdb5a6eac05_icompany_CompanyServiceImpl struct {
}

func (inst* pdb5a6eac05_icompany_CompanyServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db5a6eac059e2d06-icompany-CompanyServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-006e7bf7089b23af45a893677ef0db44-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb5a6eac05_icompany_CompanyServiceImpl) new() any {
    return &pdb5a6eac0.CompanyServiceImpl{}
}

func (inst* pdb5a6eac05_icompany_CompanyServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb5a6eac0.CompanyServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pdb5a6eac05_icompany_CompanyServiceImpl) getDao(ie application.InjectionExt)p006e7bf70.DAO{
    return ie.GetComponent("#alias-006e7bf7089b23af45a893677ef0db44-DAO").(p006e7bf70.DAO)
}



// type p4da71880f.MarketDBAgent in package:github.com/bitwormhole/markets/app/implements/idb
//
// id:com-4da71880f9b3f74d-idb-MarketDBAgent
// class:
// alias:alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent
// scope:singleton
//
type p4da71880f9_idb_MarketDBAgent struct {
}

func (inst* p4da71880f9_idb_MarketDBAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4da71880f9b3f74d-idb-MarketDBAgent"
	r.Classes = ""
	r.Aliases = "alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4da71880f9_idb_MarketDBAgent) new() any {
    return &p4da71880f.MarketDBAgent{}
}

func (inst* p4da71880f9_idb_MarketDBAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4da71880f.MarketDBAgent)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)
    com.Alias = inst.getAlias(ie)


    return nil
}


func (inst*p4da71880f9_idb_MarketDBAgent) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p4da71880f9_idb_MarketDBAgent) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.markets.datasource}")
}



// type p4928c5976.ExampleDaoImpl in package:github.com/bitwormhole/markets/app/implements/iexamples
//
// id:com-4928c59769c8b81e-iexamples-ExampleDaoImpl
// class:
// alias:alias-18a88e167ebc1a98a958760b99ada5e3-DAO
// scope:singleton
//
type p4928c59769_iexamples_ExampleDaoImpl struct {
}

func (inst* p4928c59769_iexamples_ExampleDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4928c59769c8b81e-iexamples-ExampleDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-18a88e167ebc1a98a958760b99ada5e3-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4928c59769_iexamples_ExampleDaoImpl) new() any {
    return &p4928c5976.ExampleDaoImpl{}
}

func (inst* p4928c59769_iexamples_ExampleDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4928c5976.ExampleDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p4928c59769_iexamples_ExampleDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*p4928c59769_iexamples_ExampleDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p4928c5976.ExampleServiceImpl in package:github.com/bitwormhole/markets/app/implements/iexamples
//
// id:com-4928c59769c8b81e-iexamples-ExampleServiceImpl
// class:
// alias:alias-18a88e167ebc1a98a958760b99ada5e3-Service
// scope:singleton
//
type p4928c59769_iexamples_ExampleServiceImpl struct {
}

func (inst* p4928c59769_iexamples_ExampleServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4928c59769c8b81e-iexamples-ExampleServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-18a88e167ebc1a98a958760b99ada5e3-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4928c59769_iexamples_ExampleServiceImpl) new() any {
    return &p4928c5976.ExampleServiceImpl{}
}

func (inst* p4928c59769_iexamples_ExampleServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4928c5976.ExampleServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p4928c59769_iexamples_ExampleServiceImpl) getDao(ie application.InjectionExt)p18a88e167.DAO{
    return ie.GetComponent("#alias-18a88e167ebc1a98a958760b99ada5e3-DAO").(p18a88e167.DAO)
}



// type p03b29322a.LicenceDaoImpl in package:github.com/bitwormhole/markets/app/implements/ilicence
//
// id:com-03b29322a13acd82-ilicence-LicenceDaoImpl
// class:
// alias:alias-e29ea031d0d62a43df849346ae629515-DAO
// scope:singleton
//
type p03b29322a1_ilicence_LicenceDaoImpl struct {
}

func (inst* p03b29322a1_ilicence_LicenceDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-03b29322a13acd82-ilicence-LicenceDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-e29ea031d0d62a43df849346ae629515-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p03b29322a1_ilicence_LicenceDaoImpl) new() any {
    return &p03b29322a.LicenceDaoImpl{}
}

func (inst* p03b29322a1_ilicence_LicenceDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p03b29322a.LicenceDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p03b29322a1_ilicence_LicenceDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*p03b29322a1_ilicence_LicenceDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p03b29322a.LicenceServiceImpl in package:github.com/bitwormhole/markets/app/implements/ilicence
//
// id:com-03b29322a13acd82-ilicence-LicenceServiceImpl
// class:
// alias:alias-e29ea031d0d62a43df849346ae629515-Service
// scope:singleton
//
type p03b29322a1_ilicence_LicenceServiceImpl struct {
}

func (inst* p03b29322a1_ilicence_LicenceServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-03b29322a13acd82-ilicence-LicenceServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-e29ea031d0d62a43df849346ae629515-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p03b29322a1_ilicence_LicenceServiceImpl) new() any {
    return &p03b29322a.LicenceServiceImpl{}
}

func (inst* p03b29322a1_ilicence_LicenceServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p03b29322a.LicenceServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p03b29322a1_ilicence_LicenceServiceImpl) getDao(ie application.InjectionExt)pe29ea031d.DAO{
    return ie.GetComponent("#alias-e29ea031d0d62a43df849346ae629515-DAO").(pe29ea031d.DAO)
}



// type p19c8c70a5.ManufacturerDaoImpl in package:github.com/bitwormhole/markets/app/implements/imanufacturer
//
// id:com-19c8c70a5b89fc95-imanufacturer-ManufacturerDaoImpl
// class:
// alias:alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-DAO
// scope:singleton
//
type p19c8c70a5b_imanufacturer_ManufacturerDaoImpl struct {
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19c8c70a5b89fc95-imanufacturer-ManufacturerDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerDaoImpl) new() any {
    return &p19c8c70a5.ManufacturerDaoImpl{}
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19c8c70a5.ManufacturerDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p19c8c70a5b_imanufacturer_ManufacturerDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*p19c8c70a5b_imanufacturer_ManufacturerDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p19c8c70a5.ManufacturerServiceImpl in package:github.com/bitwormhole/markets/app/implements/imanufacturer
//
// id:com-19c8c70a5b89fc95-imanufacturer-ManufacturerServiceImpl
// class:
// alias:alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-Service
// scope:singleton
//
type p19c8c70a5b_imanufacturer_ManufacturerServiceImpl struct {
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19c8c70a5b89fc95-imanufacturer-ManufacturerServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerServiceImpl) new() any {
    return &p19c8c70a5.ManufacturerServiceImpl{}
}

func (inst* p19c8c70a5b_imanufacturer_ManufacturerServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19c8c70a5.ManufacturerServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p19c8c70a5b_imanufacturer_ManufacturerServiceImpl) getDao(ie application.InjectionExt)pfa6c9190d.DAO{
    return ie.GetComponent("#alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-DAO").(pfa6c9190d.DAO)
}



// type pdab695db7.ProductDaoImpl in package:github.com/bitwormhole/markets/app/implements/iproduct
//
// id:com-dab695db7dd81dd8-iproduct-ProductDaoImpl
// class:
// alias:alias-4cebd7daf2e45520cae84295f03244fb-DAO
// scope:singleton
//
type pdab695db7d_iproduct_ProductDaoImpl struct {
}

func (inst* pdab695db7d_iproduct_ProductDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dab695db7dd81dd8-iproduct-ProductDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-4cebd7daf2e45520cae84295f03244fb-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdab695db7d_iproduct_ProductDaoImpl) new() any {
    return &pdab695db7.ProductDaoImpl{}
}

func (inst* pdab695db7d_iproduct_ProductDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdab695db7.ProductDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pdab695db7d_iproduct_ProductDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*pdab695db7d_iproduct_ProductDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pdab695db7.ProductServiceImpl in package:github.com/bitwormhole/markets/app/implements/iproduct
//
// id:com-dab695db7dd81dd8-iproduct-ProductServiceImpl
// class:
// alias:alias-4cebd7daf2e45520cae84295f03244fb-Service
// scope:singleton
//
type pdab695db7d_iproduct_ProductServiceImpl struct {
}

func (inst* pdab695db7d_iproduct_ProductServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dab695db7dd81dd8-iproduct-ProductServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-4cebd7daf2e45520cae84295f03244fb-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdab695db7d_iproduct_ProductServiceImpl) new() any {
    return &pdab695db7.ProductServiceImpl{}
}

func (inst* pdab695db7d_iproduct_ProductServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdab695db7.ProductServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pdab695db7d_iproduct_ProductServiceImpl) getDao(ie application.InjectionExt)p4cebd7daf.DAO{
    return ie.GetComponent("#alias-4cebd7daf2e45520cae84295f03244fb-DAO").(p4cebd7daf.DAO)
}



// type pd0d159fec.ShopDaoImpl in package:github.com/bitwormhole/markets/app/implements/ishop
//
// id:com-d0d159fec579493e-ishop-ShopDaoImpl
// class:
// alias:alias-b1845fb6124970f80f51eed232048f58-DAO
// scope:singleton
//
type pd0d159fec5_ishop_ShopDaoImpl struct {
}

func (inst* pd0d159fec5_ishop_ShopDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d0d159fec579493e-ishop-ShopDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-b1845fb6124970f80f51eed232048f58-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd0d159fec5_ishop_ShopDaoImpl) new() any {
    return &pd0d159fec.ShopDaoImpl{}
}

func (inst* pd0d159fec5_ishop_ShopDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd0d159fec.ShopDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pd0d159fec5_ishop_ShopDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*pd0d159fec5_ishop_ShopDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pd0d159fec.ShopServiceImpl in package:github.com/bitwormhole/markets/app/implements/ishop
//
// id:com-d0d159fec579493e-ishop-ShopServiceImpl
// class:
// alias:alias-b1845fb6124970f80f51eed232048f58-Service
// scope:singleton
//
type pd0d159fec5_ishop_ShopServiceImpl struct {
}

func (inst* pd0d159fec5_ishop_ShopServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d0d159fec579493e-ishop-ShopServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-b1845fb6124970f80f51eed232048f58-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd0d159fec5_ishop_ShopServiceImpl) new() any {
    return &pd0d159fec.ShopServiceImpl{}
}

func (inst* pd0d159fec5_ishop_ShopServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd0d159fec.ShopServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pd0d159fec5_ishop_ShopServiceImpl) getDao(ie application.InjectionExt)pb1845fb61.DAO{
    return ie.GetComponent("#alias-b1845fb6124970f80f51eed232048f58-DAO").(pb1845fb61.DAO)
}



// type pdfe7f411a.SkuDaoImpl in package:github.com/bitwormhole/markets/app/implements/isku
//
// id:com-dfe7f411a5e2d3a8-isku-SkuDaoImpl
// class:
// alias:alias-1c4227d85bd69a939ed704015c6b7410-DAO
// scope:singleton
//
type pdfe7f411a5_isku_SkuDaoImpl struct {
}

func (inst* pdfe7f411a5_isku_SkuDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dfe7f411a5e2d3a8-isku-SkuDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-1c4227d85bd69a939ed704015c6b7410-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdfe7f411a5_isku_SkuDaoImpl) new() any {
    return &pdfe7f411a.SkuDaoImpl{}
}

func (inst* pdfe7f411a5_isku_SkuDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdfe7f411a.SkuDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pdfe7f411a5_isku_SkuDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*pdfe7f411a5_isku_SkuDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pdfe7f411a.SkuServiceImpl in package:github.com/bitwormhole/markets/app/implements/isku
//
// id:com-dfe7f411a5e2d3a8-isku-SkuServiceImpl
// class:
// alias:alias-1c4227d85bd69a939ed704015c6b7410-Service
// scope:singleton
//
type pdfe7f411a5_isku_SkuServiceImpl struct {
}

func (inst* pdfe7f411a5_isku_SkuServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dfe7f411a5e2d3a8-isku-SkuServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-1c4227d85bd69a939ed704015c6b7410-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdfe7f411a5_isku_SkuServiceImpl) new() any {
    return &pdfe7f411a.SkuServiceImpl{}
}

func (inst* pdfe7f411a5_isku_SkuServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdfe7f411a.SkuServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pdfe7f411a5_isku_SkuServiceImpl) getDao(ie application.InjectionExt)p1c4227d85.DAO{
    return ie.GetComponent("#alias-1c4227d85bd69a939ed704015c6b7410-DAO").(p1c4227d85.DAO)
}



// type p8b634539d.StandardDaoImpl in package:github.com/bitwormhole/markets/app/implements/istandard
//
// id:com-8b634539d1f18a12-istandard-StandardDaoImpl
// class:
// alias:alias-b395bddcd56e3a1ae142524b343de1b5-DAO
// scope:singleton
//
type p8b634539d1_istandard_StandardDaoImpl struct {
}

func (inst* p8b634539d1_istandard_StandardDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8b634539d1f18a12-istandard-StandardDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-b395bddcd56e3a1ae142524b343de1b5-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8b634539d1_istandard_StandardDaoImpl) new() any {
    return &p8b634539d.StandardDaoImpl{}
}

func (inst* p8b634539d1_istandard_StandardDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8b634539d.StandardDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*p8b634539d1_istandard_StandardDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*p8b634539d1_istandard_StandardDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8b634539d.StandardServiceImpl in package:github.com/bitwormhole/markets/app/implements/istandard
//
// id:com-8b634539d1f18a12-istandard-StandardServiceImpl
// class:
// alias:alias-b395bddcd56e3a1ae142524b343de1b5-Service
// scope:singleton
//
type p8b634539d1_istandard_StandardServiceImpl struct {
}

func (inst* p8b634539d1_istandard_StandardServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8b634539d1f18a12-istandard-StandardServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-b395bddcd56e3a1ae142524b343de1b5-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8b634539d1_istandard_StandardServiceImpl) new() any {
    return &p8b634539d.StandardServiceImpl{}
}

func (inst* p8b634539d1_istandard_StandardServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8b634539d.StandardServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p8b634539d1_istandard_StandardServiceImpl) getDao(ie application.InjectionExt)pb395bddcd.DAO{
    return ie.GetComponent("#alias-b395bddcd56e3a1ae142524b343de1b5-DAO").(pb395bddcd.DAO)
}



// type pecc835d0d.TrademarkDaoImpl in package:github.com/bitwormhole/markets/app/implements/itrademark
//
// id:com-ecc835d0da89413e-itrademark-TrademarkDaoImpl
// class:
// alias:alias-1142a8d2a87e67b3d0431083000e797d-DAO
// scope:singleton
//
type pecc835d0da_itrademark_TrademarkDaoImpl struct {
}

func (inst* pecc835d0da_itrademark_TrademarkDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecc835d0da89413e-itrademark-TrademarkDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-1142a8d2a87e67b3d0431083000e797d-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecc835d0da_itrademark_TrademarkDaoImpl) new() any {
    return &pecc835d0d.TrademarkDaoImpl{}
}

func (inst* pecc835d0da_itrademark_TrademarkDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecc835d0d.TrademarkDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenSer = inst.getUUIDGenSer(ie)


    return nil
}


func (inst*pecc835d0da_itrademark_TrademarkDaoImpl) getAgent(ie application.InjectionExt)p4adf7a7b6.Agent{
    return ie.GetComponent("#alias-4adf7a7b6427a7b2ee56dd9cad2335eb-Agent").(p4adf7a7b6.Agent)
}


func (inst*pecc835d0da_itrademark_TrademarkDaoImpl) getUUIDGenSer(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pecc835d0d.TrademarkServiceImpl in package:github.com/bitwormhole/markets/app/implements/itrademark
//
// id:com-ecc835d0da89413e-itrademark-TrademarkServiceImpl
// class:
// alias:alias-1142a8d2a87e67b3d0431083000e797d-Service
// scope:singleton
//
type pecc835d0da_itrademark_TrademarkServiceImpl struct {
}

func (inst* pecc835d0da_itrademark_TrademarkServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecc835d0da89413e-itrademark-TrademarkServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-1142a8d2a87e67b3d0431083000e797d-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecc835d0da_itrademark_TrademarkServiceImpl) new() any {
    return &pecc835d0d.TrademarkServiceImpl{}
}

func (inst* pecc835d0da_itrademark_TrademarkServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecc835d0d.TrademarkServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pecc835d0da_itrademark_TrademarkServiceImpl) getDao(ie application.InjectionExt)p1142a8d2a.DAO{
    return ie.GetComponent("#alias-1142a8d2a87e67b3d0431083000e797d-DAO").(p1142a8d2a.DAO)
}



// type p48d083251.CompanyController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-CompanyController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_CompanyController struct {
}

func (inst* p48d083251a_controllers_CompanyController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-CompanyController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_CompanyController) new() any {
    return &p48d083251.CompanyController{}
}

func (inst* p48d083251a_controllers_CompanyController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.CompanyController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_CompanyController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_CompanyController) getService(ie application.InjectionExt)p006e7bf70.Service{
    return ie.GetComponent("#alias-006e7bf7089b23af45a893677ef0db44-Service").(p006e7bf70.Service)
}



// type p48d083251.ExampleController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_ExampleController struct {
}

func (inst* p48d083251a_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_ExampleController) new() any {
    return &p48d083251.ExampleController{}
}

func (inst* p48d083251a_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_ExampleController) getService(ie application.InjectionExt)p18a88e167.Service{
    return ie.GetComponent("#alias-18a88e167ebc1a98a958760b99ada5e3-Service").(p18a88e167.Service)
}



// type p48d083251.LicenceController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-LicenceController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_LicenceController struct {
}

func (inst* p48d083251a_controllers_LicenceController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-LicenceController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_LicenceController) new() any {
    return &p48d083251.LicenceController{}
}

func (inst* p48d083251a_controllers_LicenceController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.LicenceController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_LicenceController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_LicenceController) getService(ie application.InjectionExt)pe29ea031d.Service{
    return ie.GetComponent("#alias-e29ea031d0d62a43df849346ae629515-Service").(pe29ea031d.Service)
}



// type p48d083251.ManufacturerController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-ManufacturerController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_ManufacturerController struct {
}

func (inst* p48d083251a_controllers_ManufacturerController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-ManufacturerController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_ManufacturerController) new() any {
    return &p48d083251.ManufacturerController{}
}

func (inst* p48d083251a_controllers_ManufacturerController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.ManufacturerController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_ManufacturerController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_ManufacturerController) getService(ie application.InjectionExt)pfa6c9190d.Service{
    return ie.GetComponent("#alias-fa6c9190d2cff4a5ebddc5bb2d680cf3-Service").(pfa6c9190d.Service)
}



// type p48d083251.ProductController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-ProductController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_ProductController struct {
}

func (inst* p48d083251a_controllers_ProductController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-ProductController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_ProductController) new() any {
    return &p48d083251.ProductController{}
}

func (inst* p48d083251a_controllers_ProductController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.ProductController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_ProductController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_ProductController) getService(ie application.InjectionExt)p4cebd7daf.Service{
    return ie.GetComponent("#alias-4cebd7daf2e45520cae84295f03244fb-Service").(p4cebd7daf.Service)
}



// type p48d083251.ShopController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-ShopController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_ShopController struct {
}

func (inst* p48d083251a_controllers_ShopController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-ShopController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_ShopController) new() any {
    return &p48d083251.ShopController{}
}

func (inst* p48d083251a_controllers_ShopController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.ShopController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_ShopController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_ShopController) getService(ie application.InjectionExt)pb1845fb61.Service{
    return ie.GetComponent("#alias-b1845fb6124970f80f51eed232048f58-Service").(pb1845fb61.Service)
}



// type p48d083251.SkuController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-SkuController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_SkuController struct {
}

func (inst* p48d083251a_controllers_SkuController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-SkuController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_SkuController) new() any {
    return &p48d083251.SkuController{}
}

func (inst* p48d083251a_controllers_SkuController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.SkuController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_SkuController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_SkuController) getService(ie application.InjectionExt)p1c4227d85.Service{
    return ie.GetComponent("#alias-1c4227d85bd69a939ed704015c6b7410-Service").(p1c4227d85.Service)
}



// type p48d083251.StandardController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-StandardController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_StandardController struct {
}

func (inst* p48d083251a_controllers_StandardController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-StandardController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_StandardController) new() any {
    return &p48d083251.StandardController{}
}

func (inst* p48d083251a_controllers_StandardController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.StandardController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_StandardController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_StandardController) getService(ie application.InjectionExt)pb395bddcd.Service{
    return ie.GetComponent("#alias-b395bddcd56e3a1ae142524b343de1b5-Service").(pb395bddcd.Service)
}



// type p48d083251.TrademarkController in package:github.com/bitwormhole/markets/app/web/controllers
//
// id:com-48d083251a7b2708-controllers-TrademarkController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p48d083251a_controllers_TrademarkController struct {
}

func (inst* p48d083251a_controllers_TrademarkController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-48d083251a7b2708-controllers-TrademarkController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p48d083251a_controllers_TrademarkController) new() any {
    return &p48d083251.TrademarkController{}
}

func (inst* p48d083251a_controllers_TrademarkController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p48d083251.TrademarkController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p48d083251a_controllers_TrademarkController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p48d083251a_controllers_TrademarkController) getService(ie application.InjectionExt)p1142a8d2a.Service{
    return ie.GetComponent("#alias-1142a8d2a87e67b3d0431083000e797d-Service").(p1142a8d2a.Service)
}


