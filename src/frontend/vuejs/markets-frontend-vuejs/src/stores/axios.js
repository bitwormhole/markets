// import { ref, computed } from 'vue'

import lib_pinia_js from '@/js/pinia.js'
import axios from 'axios'
import { useLogStore } from './logs'
import { defineStore } from 'pinia'


import LibGetterJS from '@/js/getter'



////////////////////////////////////////////////////////////////////////////////

const pinia = lib_pinia_js.GetPinia()

const theLogStore = useLogStore(pinia)

const theGetterJS = LibGetterJS.NewGetter();


////////////////////////////////////////////////////////////////////////////////


function now() {
  let d = new Date();
  return d.getTime()
}


function loadLogItemWithVO(vo, log_item) {

  let g1 = theGetterJS.Clone().Reset().From(vo)
  let t0 = now();

  let time = g1.Get('time').Result(t0.toString());
  let timestamp = g1.Get('timestamp').Result(t0);
  let message = g1.Get('message').Result('');
  let error = g1.Get('error').Result('');

  let tmp = { time, timestamp, message, error }

  for (let key in tmp) {
    log_item[key] = tmp[key]
  }
}



function makeLogItemForResult(res) {

  // if (ok)

  let g1 = theGetterJS.Clone().Reset().From(res)
  let vo = g1.Get('data').Result({})
  let method = g1.Get('config/method').Result('http.GET')
  let level = 'info'
  let status = res.status;
  let status_text = res.statusText;

  method = method.toUpperCase();

  let dst = { status, status_text, level, method }
  loadLogItemWithVO(vo, dst)

  return dst
}


function makeLogItemForError(res) {

  // if (error)

  let gett = theGetterJS.Clone().Reset().From(res)
  let vo = gett.Get('response/data').Result({})
  let method = gett.Get('config/method').Result('http.GET')
  let level = 'error'
  let status = res.status;

  method = method.toUpperCase();

  let dst = { status, level, method }
  loadLogItemWithVO(vo, dst)

  if (dst.message == '' && dst.error == '') {
    dst.message = res.message;
    dst.error = res.code;
  }

  return dst
}


////////////////////////////////////////////////////////////////////////////////


const state = function () {
  return {
    "state1": 1,
  }
}

const getters = {
  foo(state) {
    return state.state1;
  },
}

const actions = {
  execute(config) {
    let p = axios(config)
    p.then((res) => {
      let item = makeLogItemForResult(res)
      theLogStore.push(item)
    }).catch((res) => {
      let item = makeLogItemForError(res)
      theLogStore.push(item)
    })
    return p
  },
}

////////////////////////////////////////////////////////////////////////////////


export const useAxiosLib = defineStore('axios', {
  actions, getters, state,
})


////////////////////////////////////////////////////////////////////////////////
// EOF 
