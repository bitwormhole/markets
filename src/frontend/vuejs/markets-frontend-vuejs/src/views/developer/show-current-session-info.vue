<script lang="js">
import { ElButton, ElInput } from 'element-plus'

import { useAxiosLib } from '@/stores/axios'

import MyDebugBox from '@/components/widgets/debug-box/index.vue'

const theAxiosLib = useAxiosLib()

export default {
  name: 'views-dev-show-current-session-info',

  components: { ElButton, MyDebugBox },

  computed: {
    theTokenStr() {
      let o = this.m_token
      return JSON.stringify(o, null, '\t')
    },
    theSessionStr() {
      let o = this.m_session
      return JSON.stringify(o, null, '\t')
    },
  },

  data() {
    const m_token = {}
    const m_session = {}
    const m_tmp_text = ''
    return { m_token, m_session, m_tmp_text }
  },

  methods: {
    fetch() {
      let method = 'GET'
      let url = '/api/v1/sessions/current'
      let p = theAxiosLib.execute({ method, url })
      p.then((res) => {
        this.handleResponseVO(res.data)
      })
      return p
    },

    handleResponseVO(vo) {
      let session = this.getFirstItem(vo.sessions)
      let token = this.getFirstItem(vo.tokens)
      this.rewriteTermInfoToDTO(session, session.started_at, session.expired_at)
      this.rewriteTermInfoToDTO(token, token.not_before, token.not_after)
      this.m_session = session
      this.m_token = token
    },

    getFirstItem(list) {
      if (list == null) {
        return null
      }
      for (let idx in list) {
        let it = list[idx]
        if (it != null) {
          return it
        }
      }
      return null
    },

    formatTimestamp(ts) {
      let d = new Date()
      d.setTime(ts)
      return d.toString()
    },

    rewriteTermInfoToDTO(dto, t1, t2) {
      let max_age = t2 - t1
      dto['str_term_t1'] = this.formatTimestamp(t1)
      dto['str_term_t2'] = this.formatTimestamp(t2)
      dto['str_term_max_age'] = max_age
    },
  },

  mounted() {
    this.fetch()
  },

  props: {},
}
</script>

<style scoped></style>

<template>
  <frame-for-developer>
    <h1>Show Current Session Info</h1>

    <ElButton @click="fetch"> Refresh </ElButton>

    <MyDebugBox title="Token">
      <pre> {{ theTokenStr }} </pre>
    </MyDebugBox>

    <MyDebugBox title="Session">
      <pre> {{ theSessionStr }} </pre>
    </MyDebugBox>

    <MyDebugBox title="临时笔记">
      <ElInput v-model="m_tmp_text" type="textarea"></ElInput>
    </MyDebugBox>
  </frame-for-developer>
</template>
