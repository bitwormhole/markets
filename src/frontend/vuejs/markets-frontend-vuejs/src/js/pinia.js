// pinia.js
//
// 这个模块用于管理全局唯一的 pinia-store 实例

import { createPinia } from 'pinia'

////////////////////////////////////////////////////////////////////////////////


function PiniaStoreHolder() {

  let p = createPinia();
  this.pinia = p;
}

PiniaStoreHolder.theInst = null;

PiniaStoreHolder.getInstance = function () {
  let inst = PiniaStoreHolder.theInst;
  if (inst == null) {
    inst = new PiniaStoreHolder();
    PiniaStoreHolder.theInst = inst;
  }
  return inst;
}

PiniaStoreHolder.prototype = {

  getPiniaInstance: function () {
    return this.pinia;
  },

}

////////////////////////////////////////////////////////////////////////////////

export default {

  name: "lib-pinia-js",

  GetPinia() {
    let holder = PiniaStoreHolder.getInstance();
    return holder.getPiniaInstance();
  },

}
