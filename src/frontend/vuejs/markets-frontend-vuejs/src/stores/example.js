// import { ref, computed } from 'vue'
import { defineStore } from 'pinia'


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
  action1() {
    return Promise.resolve("ok")
  },
}


export const useCounterStore = defineStore('example', {
  actions, getters, state,
})
