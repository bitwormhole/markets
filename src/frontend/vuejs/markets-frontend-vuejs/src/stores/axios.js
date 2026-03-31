// import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'

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
    return axios(config)
  },
}


export const useAxiosLib = defineStore('axios', {
  actions, getters, state,
})
