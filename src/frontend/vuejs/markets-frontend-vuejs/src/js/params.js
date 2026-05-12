
// params.js
// 这个模块是用于快捷处理 "参数集合" 的工具类型


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


function Params() {
  this.table = {}
}

Params.prototype = {

  Export(dst) {
    if (dst == null) {
      dst = {}
    }
    let src = this.table;
    copyKeyValueTable(src, dst)
    return dst;
  },

  Import(src) {
    if (src == null) {
      return;
    }
    let dst = this.table;
    copyKeyValueTable(src, dst)
  },

}

////////////////////////////////////////////////////////////////////////////////

export default {

  name: "lib-params-js",

  NewParams() {
    return new Params();
  }
}

////////////////////////////////////////////////////////////////////////////////
// EOF 
