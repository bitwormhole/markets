<script lang="js">
// import lib_getter_js from '@/js/getter'
// import lib_params_js from '@/js/params'

import { useAxiosLib } from '@/stores/axios'
import LibLocatorJS from '@/js/locator.js'
import LibGetterJS from '@/js/getter.js'

const theAxiosLib = useAxiosLib()
const theGetter = LibGetterJS.NewGetter()

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

    theVMO() {
      return this.modelValue
    },

    theRevision() {
      let g = theGetter.Clone().From(this.modelValue)
      g = g.Get('revision').AsNumber()
      return g.Result(0)
    },

    theItemListFieldName() {
      let name = this.items
      if (name == null || name == '') {
        name = 'items'
      }
      return name
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
      let vmo = this.theVMO
      locator.Init(this, vmo)

      locator.GetBank('params', true)
      locator.GetBank('pagination', true)
      locator.GetBank('want', true)

      let d = locator.GetBank('default', true)
      let q = locator.GetBank('query', true)

      d.SetValue('limit', 10)
      d.SetValue('offset', 0)

      q.Import(this.$route.query)
    },

    fetch() {
      let method = 'GET'
      let url = this.path
      let params = this.prepareParams()

      let p = theAxiosLib.execute({ method, url, params })

      p.then((res) => {
        let vo = res.data
        this.handleResponseVO(vo)
      })
      this.rewriteParams(params)
      return p
    },

    filterParams(src, dst) {
      for (let key in src) {
        let value = src[key]
        if (value == null) {
          continue
        }
        if (value == '') {
          continue
        }
        dst[key] = value
      }
    },

    prepareParams() {
      let l = this.locator
      let bank_name_list = ['default', 'query', 'want']
      let dst = {}

      for (let idx in bank_name_list) {
        let name = bank_name_list[idx]
        let bank = l.GetBank(name, true)
        let tmp = {}
        tmp = bank.Export(tmp)
        this.filterParams(tmp, dst)
      }

      return dst
    },

    rewriteParams(params) {
      this.rewriteParamsToVMO(params)
      this.rewriteParamsToLocation(params)
    },

    rewriteParamsToLocation(params) {
      let query = params
      this.$router.push({ query })
    },

    rewriteParamsToVMO(params) {
      let l = this.locator
      let bank = l.GetBank('params', true)
      bank.Import(params)
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

    handleResponsePagination(pagination) {
      let vmo = this.modelValue
      vmo['pagination'] = pagination
      this.fireOnPage(pagination)
    },

    handleResponseItems(items) {
      let vmo = this.modelValue
      vmo['items'] = items
      this.fireOnItems(items)
    },

    handleResponseVO(vo) {
      let gett = theGetter.Clone().Reset()
      let items_field_name = this.theItemListFieldName
      let items = gett.From(vo).Get(items_field_name).AsAny().Result([])
      let page = gett.From(vo).Get('pagination').AsAny().Result({})

      this.handleResponsePagination(page)
      this.handleResponseItems(items)

      this.fireOnData(vo)
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
      this.fetch()
    }
  },

  beforeUpdate() {
    // this.checkSumUpdate()
  },

  watch: {
    theSum() {
      // this.checkSumUpdate()
    },

    theRevision() {
      // this.checkSumUpdate()
      this.fetch()
    },
  },

  props: {
    modelValue: Object, // a 'VMO' (view-model-object)
    path: String, // the URL.path of REST-api
    auto: Boolean, // call fetch on-started , if true
    items: String, // the vo-field-name of items
  },
}
</script>

<style></style>

<template>
  <span class="main-table-loader"> </span>
</template>
