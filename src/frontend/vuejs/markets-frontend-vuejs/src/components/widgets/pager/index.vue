<script>
import { ElPagination } from 'element-plus'
import LibLocatorJS from '@/js/locator.js'
import LibGetterJS from '@/js/getter.js'

const theGetter = LibGetterJS.NewGetter()

export default {
  name: 'widgets-common-pager',

  components: { ElPagination },

  computed: {
    offset() {
      let size = this.limit
      let current = this.currentPage
      return size * (current - 1)
    },

    limit() {
      const def = 10
      let g = theGetter.Clone().From(this.modelValue)
      g = g.Get('pagination/size').AsNumber()
      return g.Result(def)
    },

    currentPage() {
      const def = 1
      let g = theGetter.Clone().From(this.modelValue)
      g = g.Get('pagination/page').AsNumber()
      return g.Result(def)
    },

    total() {
      const def = 0
      let g = theGetter.Clone().From(this.modelValue)
      g = g.Get('pagination/total').AsNumber()
      return g.Result(def)
    },
  },

  data() {
    const locator = LibLocatorJS.NewLocator(this)
    return { locator }
  },

  methods: {
    init() {
      let l = this.locator
      l.Init(this, this.modelValue)
    },

    now() {
      let d = new Date()
      return d.getTime()
    },

    on_pager_size_change(size) {
      let limit = size
      let offset = this.offset
      this.updatePageAt(offset, limit)
    },

    on_pager_current_change(current) {
      let limit = this.limit
      let offset = (current - 1) * limit
      this.updatePageAt(offset, limit)
    },

    updatePageAt(offset, limit) {
      let locator = this.locator
      let bank = locator.GetBank('want', true)
      bank.SetValue('limit', limit)
      bank.SetValue('offset', offset)

      this.fireReload()
    },

    fireReload() {
      this.$emit('reload')
    },
  },

  mounted() {
    this.init()
  },

  props: {
    modelValue: Object, // a 'VMO' (View-Model-Object)
  },
}
</script>

<style></style>

<template>
  <div>
    <ElPagination
      :page-size="limit"
      :page-sizes="[5, 10, 20, 50, 100, 200, 500]"
      :total="total"
      :current-page="currentPage"
      background
      layout="total , sizes  ,prev, pager, next"
      @current-change="on_pager_current_change"
      @size-change="on_pager_size_change"
    ></ElPagination>
  </div>
</template>
