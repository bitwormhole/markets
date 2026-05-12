

import { useAxiosLib } from './axios.js'
import { defineStore } from 'pinia'

import lib_pinia_js from '@/js/pinia.js'
import lib_getter_js from '@/js/getter.js'


////////////////////////////////////////////////////////////////////////////////

const thePiniaInst = lib_pinia_js.GetPinia()
const theAxiosLib = useAxiosLib(thePiniaInst)

////////////////////////////////////////////////////////////////////////////////


const state = function () {
  return {
    inner_items: [],
    inner_pagination: {
      total: 0,
      offset: 0,
      limit: 5,
    },
  }
}

const getters = {
  items(state) {
    return state.inner_items;
  },

  pagination(state) {
    return state.inner_pagination;
  },
}

const actions = {

  fetch(config) {

    let method = 'GET'
    let url = '/api/v1/standards'
    let getter = lib_getter_js.NewGetter();

    let params = getter.From(config).Get('params').Result({})

    let ret = theAxiosLib.execute({ method, url, params })

    ret.then((res) => {

      let vo = res.data;

      this.inner_items = getter.From(vo).Get('standards').Result([]);
      this.inner_pagination = getter.From(vo).Get('pagination').Result({});

    });  // .catch((res) => { })

    return ret;
  },
}

////////////////////////////////////////////////////////////////////////////////


export const useStandardStore = defineStore('standards', {
  actions, getters, state,
})

////////////////////////////////////////////////////////////////////////////////
// EOF 
