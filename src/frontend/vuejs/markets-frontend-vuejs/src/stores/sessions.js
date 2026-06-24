// sessions.js

import { useAxiosLib } from './axios.js'
import { defineStore } from 'pinia'

import lib_pinia_js from '@/js/pinia.js'
import lib_getter_js from '@/js/getter.js'
import lib_base64_js from '@/js/base64.js'


////////////////////////////////////////////////////////////////////////////////

const thePiniaInst = lib_pinia_js.GetPinia()
const theAxiosLib = useAxiosLib(thePiniaInst)

////////////////////////////////////////////////////////////////////////////////

function makeHttpBasicAuthHeader(user, pass) {

  let plain = user + ':' + pass;
  let b64 = lib_base64_js.Encode(plain)
  return "Basic " + b64
}

////////////////////////////////////////////////////////////////////////////////




const state = function () {
  return {
    current_session_info: {
      a: 1,
      b: 2,
    },
  }
}

const getters = {
  current(state) {
    return state.current_session_info;
  },
}

const actions = {

  fetch() {

    let method = 'GET'
    let url = '/api/v1/sessions'

    return theAxiosLib.execute({ method, url }).then((res) => {
      let vo = res.data
      let items = vo.sessions
      this.current_session_info = items[0]
    })
  },

  login_basic(config) {

    let username = config.username;
    let password = config.password;
    let http_auth_hdr = makeHttpBasicAuthHeader(username, password)

    let method = 'POST'
    let url = '/api/v1/auth/login'
    let headers = {
      Authorization: http_auth_hdr
    }
    let data = {}
    return theAxiosLib.execute({ method, url, headers, data }).then((res) => {
      let vo = res.sessions
      this.current_session_info.xxxx_vo = vo
    })
  },

  logout() {
    return this.exit()
  },

  exit() {
    let method = 'POST'
    let url = '/api/v1/sessions/exit'
    let headers = {}
    let data = {}
    return theAxiosLib.execute({ method, url, headers, data })
  },

  keepAlive() {
    return Promise.reject('no impl')
  },

}


export const useSessionStore = defineStore('sessions', {
  actions, getters, state,
})
