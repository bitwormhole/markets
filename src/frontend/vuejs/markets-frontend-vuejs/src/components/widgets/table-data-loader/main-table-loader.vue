<script lang="js">
// import lib_getter_js from '@/js/getter'
// import lib_params_js from '@/js/params'

import { useAxiosLib } from '@/stores/axios'
import LibLocatorJS from '@/js/locator.js'

const theAxiosLib = useAxiosLib()

export default {
  name: 'widgets-main-table-loader',

  components: {},

  computed: {
    theSum() {
      let locator = this.locator
      locator.Load()
      let limit = locator.Get('limit', -1)
      let offset = locator.Get('offset', -1)
      let rev = this.revision
      return limit + ':' + offset + ':' + rev
    },
  },

  data() {
    let locator = LibLocatorJS.NewLocator(this)
    const sum = 'init'
    return { locator, sum }
  },

  methods: {
    init() {
      let locator = this.locator
      let limit = locator.Get('limit', 10)
      let offset = locator.Get('offset', 0)

      limit = Number(limit)
      offset = Number(offset)

      locator.Set('limit', limit)
      locator.Set('offset', offset)

      if (limit < 1) {
        limit = 1
      }

      let current = Math.floor(offset / limit) + 1

      let p = {
        page: current,
        size: limit,
        total: 0,
      }
      this.fireOnPage(p)
    },

    fetch() {
      let locator = this.locator
      locator.Load()

      let method = 'GET'
      let url = this.path
      let params = locator.GetParams()

      let p = theAxiosLib.execute({ method, url, params })
      p.then((res) => {
        let vo = res.data
        this.fireOnData(vo)
        this.fireOnItems(vo.roles)
        this.fireOnPage(vo.pagination)
      })
      return p
    },

    now() {
      let d = new Date()
      return d.getTime()
    },

    tryUseDefaultValue(value, value_default) {
      if (value == null) {
        value = value_default
      }
      return value
    },

    fireOnData(vo) {
      vo = this.tryUseDefaultValue(vo, {})
      this.$emit('on-data', vo)
    },
    fireOnItems(list) {
      list = this.tryUseDefaultValue(list, [])
      this.$emit('on-items', list)
    },
    fireOnPage(pagination) {
      pagination = this.tryUseDefaultValue(pagination, {})
      this.$emit('on-page', pagination)
    },

    checkSumUpdate() {
      let sum1 = this.sum
      let sum2 = this.theSum
      if (sum1 == sum2) {
        return
      }
      this.sum = sum2
      this.reload()
    },

    reload() {
      console.log(' main-table-loader: reload() ')
      this.fetch()
    },
  },

  mounted() {
    this.init()
    if (this.auto) {
      this.checkSumUpdate()
    }
  },

  beforeUpdate() {
    // this.checkSumUpdate()
  },

  watch: {
    theSum() {
      // this.checkSumUpdate()
    },

    revision() {
      this.checkSumUpdate()
    },
  },

  props: {
    revision: Number,
    path: String,
    auto: Boolean,
  },
}
</script>

<style></style>

<template>
  <span class="main-table-loader"> </span>
</template>
