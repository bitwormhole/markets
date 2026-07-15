<script>
import { ElPagination } from 'element-plus'
import LibLocatorJS from '@/js/locator.js'
import LibGetterJS from '@/js/getter.js'

export default {
  name: 'widgets-common-pager',

  computed: {
    thePagination() {
      const def = {
        page: 1,
        size: 5,
        total: 0,
      }
      let gett = LibGetterJS.NewGetter()
      return gett.From(this.pagination).Get('').Result(def)
    },

    limit() {
      const def = 10
      let gett = LibGetterJS.NewGetter()
      let p = this.thePagination
      let str = gett.From(p).Get('size').Result(def)
      return Number(str)
    },

    offset() {
      let gett = LibGetterJS.NewGetter()
      let p = this.thePagination
      let size = gett.From(p).Get('size').Result(10)
      let page = gett.From(p).Get('page').Result(1)
      page = Number(page)
      size = Number(size)
      return size * (page - 1)
    },

    total() {
      let gett = LibGetterJS.NewGetter()
      let p = this.thePagination
      let total = gett.From(p).Get('total').Result(0)
      total = Number(total)
      return total
    },

    currentPage() {
      let gett = LibGetterJS.NewGetter()
      let p = this.thePagination
      let page = gett.From(p).Get('page').Result(1)
      page = Number(page)
      return page
    },
  },

  data() {
    let locator = LibLocatorJS.NewLocator(this)
    return { locator }
  },

  methods: {
    init() {
      this.locator = LibLocatorJS.NewLocator(this)
    },

    now() {
      let d = new Date()
      return d.getTime()
    },

    on_pager_limit(size) {
      let limit = size
      let offset = this.offset
      this.updatePageAt(limit, offset)
    },

    on_pager_offset(current) {
      let limit = this.limit
      let offset = (current - 1) * limit
      this.updatePageAt(limit, offset)
    },

    updatePageAt(limit, offset) {
      let locator = this.locator
      // let now = this.now()
      if (limit == null) {
        limit = this.limit
      }
      if (offset == null) {
        offset = this.offset
      }
      locator.Set('limit', limit)
      locator.Set('offset', offset)
      // locator.Set('timestamp', now)
      locator.Apply()

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
    pagination: Object, // a pagination object { page , size, total }
  },
}
</script>

<style></style>

<template>
  <div>
    <ElPagination
      :page-size="limit"
      :page-sizes="[3, 10, 20, 50, 100, 200]"
      :total="total"
      :current-page="currentPage"
      background
      layout="total , sizes  ,prev, pager, next"
      @current-change="on_pager_offset"
      @size-change="on_pager_limit"
    ></ElPagination>
  </div>
</template>
