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


function innerGetFirstItem(list, value_default) {
  if (list != null) {
    for (let idx in list) {
      let it = list[idx]
      if (it == null) {
        continue
      }
      return it
    }
  }
  return value_default;
}


////////////////////////////////////////////////////////////////////////////////


const state = function () {

  const m_session = {}
  const m_token = {}
  const current_session_info = {
    a: 1,
    b: 2,
  }

  return { m_session, m_token, current_session_info }
}

const getters = {

  current(state) {
    return state.current_session_info;
  },

  current_session(state) {
    return state.m_session;
  },

  current_token(state) {
    return state.m_token;
  },
}

const actions = {

  fetch() {
    let method = 'GET'
    let url = '/api/v1/sessions/current'

    return theAxiosLib.execute({ method, url }).then((res) => {
      let vo = res.data
      let tt = innerGetFirstItem(vo.tokens, {})
      let ss = innerGetFirstItem(vo.sessions, {})

      this.current_session_info = ss
      this.m_session = ss
      this.m_token = tt
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
    let method = 'POST'
    let url = '/api/v1/sessions/keep-alive'
    let headers = {}
    let data = {}
    return theAxiosLib.execute({ method, url, headers, data })
  },

}


export const useSessionStore = defineStore('sessions', {
  actions, getters, state,
})
