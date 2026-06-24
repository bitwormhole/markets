// import { ref, computed } from 'vue'

import lib_pinia_js from '@/js/pinia.js'
import axios from 'axios'
import { useLogStore } from './logs'
import { defineStore } from 'pinia'


////////////////////////////////////////////////////////////////////////////////

const pinia = lib_pinia_js.GetPinia()

const theLogStore = useLogStore(pinia)

////////////////////////////////////////////////////////////////////////////////


function now() {
  let d = new Date();
  return d.getTime()
}


function makeLogItemForResult(res) {

  let timestamp = now();
  let level = 'info'
  let status = 200;

  return { status, level, timestamp }
}


function makeLogItemForError(res) {


  let timestamp = now();
  let level = 'error'
  let status = res.status;

  return { status, level, timestamp }
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
