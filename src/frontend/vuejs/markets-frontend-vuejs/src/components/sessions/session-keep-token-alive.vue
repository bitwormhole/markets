<script lang="js">
// session-keep-token-alive.vue

import LibGetterJS from '@/js/getter'
import { useSessionStore } from '@/stores/sessions'

const theSessionStore = useSessionStore()
const theTimerInterval = 1000 * 10

export default {
  name: 'session-keep-token-alive',

  components: {},

  computed: {
    theCurrentToken() {
      return theSessionStore.current_token
    },

    theTokenNotAfter() {
      let gett = LibGetterJS.NewGetter()
      gett.From(this.theCurrentToken).AsNumber().Get('not_after')
      return gett.Result(0)
    },

    theTokenNotBefore() {
      let gett = LibGetterJS.NewGetter()
      gett.From(this.theCurrentToken).AsNumber().Get('not_before')
      return gett.Result(0)
    },
  },

  data() {
    const timerStartedAt = 0
    const timerStoppedAt = 0
    const timerRunningID = 0
    return {
      timerStartedAt,
      timerStoppedAt,
      timerRunningID,
    }
  },

  methods: {
    now() {
      let d = new Date()
      return d.getTime()
    },

    innerTryGetStatus() {
      if (this.isNeedGetStatus()) {
        this.innerDoGetStatus()
      }
    },

    innerTryPostKeepAlive() {
      if (this.isNeedKeepAlive()) {
        this.innerDoPostKeepAlive()
      }
    },

    innerDoGetStatus() {
      let p = theSessionStore.fetch()
      p.then(() => {
        this.innerTryPostKeepAlive()
      })
      return p
    },

    innerDoPostKeepAlive() {
      let p = theSessionStore.keepAlive()
      return p
    },

    computeTimeStampAt(t0, t100, percent) {
      if (0 < t0 && t0 < t100) {
        let diff = t100 - t0
        return diff * percent + t0
      }
      return 0
    },

    isNeedKeepAlive() {
      let tn = this.now()
      let t0 = this.theTokenNotBefore
      let t100 = this.theTokenNotAfter
      let t80 = this.computeTimeStampAt(t0, t100, 0.8)
      let t200 = this.computeTimeStampAt(t0, t100, 2.0)
      return t80 < tn && tn < t200
    },

    isNeedGetStatus() {
      let tn = this.now()
      let t0 = this.theTokenNotBefore
      let t100 = this.theTokenNotAfter
      let t80 = this.computeTimeStampAt(t0, t100, 0.8)
      let t200 = this.computeTimeStampAt(t0, t100, 2.0)
      return t80 < tn && tn < t200
    },

    onTimer() {
      // let d = new Date()
      // console.log('keep-token-alive: on-timer', d)
      this.innerTryGetStatus()
    },

    handleTimerCallback(timer_id) {
      if (timer_id != this.timerRunningID) {
        return
      }
      const interval = theTimerInterval
      setTimeout(() => {
        this.onTimer()
        this.handleTimerCallback(timer_id)
      }, interval)
    },

    startTimer() {
      let t = this.now()

      console.log('keep-token-alive: start timer @id=', t)

      this.timerStartedAt = t
      this.timerRunningID = t
      this.timerStoppedAt = 0

      this.handleTimerCallback(t)
    },

    stopTimer() {
      console.log('keep-token-alive: stop timer')
      let t = this.now()
      // this.timerStartedAt = t
      this.timerRunningID = 0
      this.timerStoppedAt = t
    },
  },

  mounted() {
    this.startTimer()
  },

  unmounted() {
    this.stopTimer()
  },

  props: {},
}
</script>

<style scoped></style>

<template>
  <span></span>
</template>
