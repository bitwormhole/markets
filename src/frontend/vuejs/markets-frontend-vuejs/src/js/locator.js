// locator.js
// 这个模块是用于 计算当前页面的 path&query '/path?id=n&query=x' 

import { tr } from "element-plus/es/locales.mjs";
import pagination from "./pagination";


////////////////////////////////////////////////////////////////////////////////


function copyKeyValueTable(src, dst) {

  if (src == null || dst == null) {
    return;
  }

  for (let name in src) {
    let value = src[name];
    dst[name] = value;
  }

  return dst;
}


////////////////////////////////////////////////////////////////////////////////


const a_vmo_demo = {

  query: {
    k1: 'v1',
    k2: 'v2',
    k3: 'v3',
  },

  pagination: {
    page: 0,
    size: 0,
    total: 0,
  },

  items: [],

  revision: 0,
}

////////////////////////////////////////////////////////////////////////////////

function innerLocatorCore(ctx) {
  this.context = ctx;
  this.path = ''
  this.query = {}
  this.vmo = {}
}

innerLocatorCore.prototype = {

  init(ctx, vmo) {
    this.context = ctx;
    this.vmo = vmo;
  },

  load() {

    let ctx = this.context;
    let path = ctx.$route.path;
    let src = ctx.$route.query;
    let dst = this.query;

    copyKeyValueTable(src, dst)

    this.path = path
  },

  apply() {

    // let url = this.computeURL() 

    let ctx = this.context;
    let query = this.query;
    ctx.$router.push({ query })
  },


  computeURL() {

    let str = '';
    let path = this.path;
    let q = this.query;

    if (path != null) {
      str = str + path
    }

    if (q != null) {
      let keys = []
      for (let k in q) {
        keys.push(k + '')
      }
      keys.sort()
      let sep = '?'
      for (let i in keys) {
        let name = keys[i]
        let value = q[name]
        str = str + sep + name + '=' + value
        sep = '&'
      }
    }

    return str
  },

  setParam(name, value) {
    if (name == null) {
      return;
    }
    this.query[name] = value;
  },

  getParam(name, value_default) {
    let value = this.query[name];
    if (value == null) {
      value = value_default;
    }
    return value;
  },

  getBank(name, auto_make) {
    let vmo = this.vmo;
    let bank = new innerLocatorBankCore(name, vmo);
    if (auto_make) {
      bank.init()
    }
    return bank
  },

}


////////////////////////////////////////////////////////////////////////////////


function innerLocatorBankCore(name, vmo) {

  if (vmo == null) {
    vmo = {}
  }

  if (name == null) {
    name = ''
  }

  this.vmo = vmo
  this.bank_name = name
}

innerLocatorBankCore.prototype = {

  init() {
    this.getBankTable(true)
  },

  doImportValues(src) {
    let dst = this.getBankTable(true);
    copyKeyValueTable(src, dst)
    return null;
  },

  doExportValues(dst) {
    let src = this.getBankTable(false);
    if (dst == null) {
      dst = {}
    }
    copyKeyValueTable(src, dst)
    return dst
  },

  getBankTable(auto_make) {
    let name = this.bank_name + '';
    let vmo = this.getVmo()
    let t = vmo[name];
    if ((t == null) && auto_make) {
      t = {};
      vmo[name] = t;
    }
    return t
  },

  getVmo() {
    let o = this.vmo;
    if (o == null) {
      o = {}
      this.vmo = o
    }
    return o
  },


  getParam(name, value_default) {
    let table = this.getBankTable(false)
    let value = null;
    if (table != null) {
      value = table[name]
    }
    if (value == null) {
      value = value_default;
    }
    return value
  },

  setParam(name, value) {
    if (name == null || value == null) {
      return
    }
    let table = this.getBankTable(true)
    table[name] = value
  },

  names() {
    let dst = []
    let src = this.getBankTable(false);
    if (src == null) {
      return dst
    }
    for (let name in src) {
      dst.push(name)
    }
    return dst
  },

}


////////////////////////////////////////////////////////////////////////////////

function locatorBankFacade(bank_core) {
  this.core = bank_core
}

locatorBankFacade.prototype = {

  Import(src) {
    return this.core.doImportValues(src)
  },

  Export(dst) {
    return this.core.doExportValues(dst)
  },

  ListNames() {
    return this.core.names()
  },

  GetValue(name, value_default) {
    return this.core.getParam(name, value_default)
  },

  SetValue(name, value) {
    return this.core.setParam(name, value)
  },
}


////////////////////////////////////////////////////////////////////////////////


function Locator(context) {
  let c = new innerLocatorCore(context);
  c.load();
  this.core = c;
}

Locator.prototype = {

  Init(ctx, vmo) {
    this.core.init(ctx, vmo)
  },

  InitBank(name) {
    this.core.getBank(name, true)
  },

  Load() {
    this.core.load();
  },

  Apply() {
    this.core.apply();
  },

  GetParams() {
    return this.core.query;
  },

  Get(name, value_default) {
    return this.core.getParam(name, value_default)
  },

  Set(name, value) {
    return this.core.setParam(name, value)
  },

  GetBank(name, auto_make) {
    let core = this.core.getBank(name, auto_make);
    return new locatorBankFacade(core);
  },
}

////////////////////////////////////////////////////////////////////////////////

export default {

  name: "lib-locator-js",

  NewLocator(context) {
    return new Locator(context);
  }
}

////////////////////////////////////////////////////////////////////////////////
// EOF 
