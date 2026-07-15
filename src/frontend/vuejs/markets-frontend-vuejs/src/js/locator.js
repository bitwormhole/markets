// locator.js
// 这个模块是用于 计算当前页面的 path&query '/path?id=n&query=x' 


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

function innerLocatorCore(ctx) {
  this.context = ctx;
  this.path = ''
  this.query = {}
}

innerLocatorCore.prototype = {

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

}


////////////////////////////////////////////////////////////////////////////////


function Locator(context) {
  let c = new innerLocatorCore(context);
  c.load();
  this.core = c;
}

Locator.prototype = {

  // Export(dst) {
  //   if (dst == null) {
  //     dst = {}
  //   }
  //   let src = this.table;
  //   copyKeyValueTable(src, dst)
  //   return dst;
  // },

  // Import(src) {
  //   if (src == null) {
  //     return;
  //   }
  //   let dst = this.table;
  //   copyKeyValueTable(src, dst)
  // },


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
