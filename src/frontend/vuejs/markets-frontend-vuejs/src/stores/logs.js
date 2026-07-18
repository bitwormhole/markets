// import { ref, computed } from 'vue'
import { defineStore } from 'pinia'


const a_item_demo = {
  level: "info", // [trace|debug|info|warn|error|fatal]
  message: "",
  error: "",
  status: 200,
  timestamp: 123456789,
}


function isErrorItem(item) {
  if (item == null) {
    return false
  }
  let level = item.level;
  return ((level == 'error') || (level == 'fatal'))
}



////////////////////////////////////////////////////////////////////////////////


const state = function () {

  const m_list = [];
  const m_revision = 0;
  const m_latest = {}
  const m_limit = 99; // list 中最多可以包含的条目数量

  return { m_list, m_revision, m_limit, m_latest }
}

const getters = {
  all(state) {
    return state.m_list;
  },

  latest(state) {
    return state.m_latest;
  },

  revision(state) {
    return state.m_revision;
  },

}

const actions = {
  push(item) {
    let st = this.$state;
    if (item == null) {
      return Promise.reject("log is nil")
    }
    if (isErrorItem(item)) {
      console.error(item)
    } else {
      console.log(item)
    }
    st.m_list.push(item)
    st.m_latest = item
    st.m_revision++;
    return Promise.resolve(item)
  },
}


export const useLogStore = defineStore('logs', {
  actions, getters, state,
})
