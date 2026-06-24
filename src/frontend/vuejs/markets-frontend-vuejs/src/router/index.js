import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },

    { path: '/login', component: () => import('@/views/user/login.vue') },


    ////////////////////////////////////////////////////////////////////////////
    // /admin/*

    { path: '/admin', component: () => import('@/views/admin/index.vue') },


    // /admin/companies/*

    { path: '/admin/companies', component: () => import('@/views/admin/companies/index.vue') },
    { path: '/admin/companies/add', component: () => import('@/views/admin/companies/company-add-view.vue') },
    { path: '/admin/companies/query', component: () => import('@/views/admin/companies/company-query-view.vue') },
    { path: '/admin/companies/:id/edit', component: () => import('@/views/admin/companies/company-edit-view.vue') },
    { path: '/admin/companies/:id/detail', component: () => import('@/views/admin/companies/company-detail-view.vue') },
    { path: '/admin/companies/:id', component: () => import('@/views/admin/companies/company-detail-view.vue') },

    // /admin/licences/*

    { path: '/admin/licences', component: () => import('@/views/admin/licences/index.vue') },
    { path: '/admin/licences/add', component: () => import('@/views/admin/licences/licence-add-view.vue') },


    // /admin/manufacturers/*

    { path: '/admin/manufacturers/', component: () => import('@/views/admin/manufacturers/index.vue') },
    { path: '/admin/manufacturers/add', component: () => import('@/views/admin/manufacturers/manufacturer-add-view.vue') },

    // /admin/products/*

    { path: '/admin/products/', component: () => import('@/views/admin/products/index.vue') },
    { path: '/admin/products/add', component: () => import('@/views/admin/products/product-add-view.vue') },
    { path: '/admin/products/edit', component: () => import('@/views/admin/products/product-edit-view.vue') },
    { path: '/admin/products/detail', component: () => import('@/views/admin/products/product-detail-view.vue') },
    { path: '/admin/products/query', component: () => import('@/views/admin/products/product-query-view.vue') },


    // /admin/shops/*

    { path: '/admin/shops/', component: () => import('@/views/admin/shops/index.vue') },
    { path: '/admin/shops/add', component: () => import('@/views/admin/shops/shop-add-view.vue') },
    { path: '/admin/shops/edit', component: () => import('@/views/admin/shops/shop-edit-view.vue') },
    { path: '/admin/shops/detail', component: () => import('@/views/admin/shops/shop-detail-view.vue') },
    { path: '/admin/shops/query', component: () => import('@/views/admin/shops/shop-query-view.vue') },

    // /admin/skus/*

    { path: '/admin/skus', component: () => import('@/views/admin/skus/index.vue') },
    { path: '/admin/skus/add', component: () => import('@/views/admin/skus/sku-add-view.vue') },


    // /admin/standards/*

    { path: '/admin/standards', component: () => import('@/views/admin/standards/index.vue') },
    { path: '/admin/standards/add', component: () => import('@/views/admin/standards/standard-add-view.vue') },

    // /admin/trademarks/*

    { path: '/admin/trademarks', component: () => import('@/views/admin/trademarks/index.vue') },
    { path: '/admin/trademarks/add', component: () => import('@/views/admin/trademarks/trademark-add-view.vue') },

    // /admin/users/*

    { path: '/admin/users', component: () => import('@/views/admin/users/index.vue') },
    { path: '/admin/users/add', component: () => import('@/views/admin/users/user-add-view.vue') },
    { path: '/admin/users/query', component: () => import('@/views/admin/users/user-query-view.vue') },
    { path: '/admin/users/:id/edit', component: () => import('@/views/admin/users/user-edit-view.vue') },
    { path: '/admin/users/:id/detail', component: () => import('@/views/admin/users/user-detail-view.vue') },


    ////////////////////////////////////////////////////////////////////////////
    // /dev/*


    { path: '/dev', component: () => import('@/views/developer/index.vue') },
    { path: '/dev/ui', component: () => import('@/views/developer/full-ui-list-view.vue') },


    // /dev/examples/*

    { path: '/dev/examples/', component: () => import('@/views/developer/examples/index.vue') },
    { path: '/dev/examples/add', component: () => import('@/views/developer/examples/example-add-view.vue') },
    { path: '/dev/examples/detail', component: () => import('@/views/developer/examples/example-detail-view.vue') },
    { path: '/dev/examples/edit', component: () => import('@/views/developer/examples/example-edit-view.vue') },
    { path: '/dev/examples/query', component: () => import('@/views/developer/examples/example-query-view.vue') },



    ////////////////////////////////////////////////////////////////////////////
    // /(user)/*

    { path: '/home', component: () => import('@/views/user/index.vue') },

    // /(user)/examples/* 

    { path: '/examples/', component: () => import('@/views/user/examples/index.vue') },
    { path: '/examples/add', component: () => import('@/views/user/examples/example-add-view.vue') },
    { path: '/examples/detail', component: () => import('@/views/user/examples/example-detail-view.vue') },
    { path: '/examples/edit', component: () => import('@/views/user/examples/example-edit-view.vue') },
    { path: '/examples/query', component: () => import('@/views/user/examples/example-query-view.vue') },



    ////////////////////////////////////////////////////////////////////////////
    // /my/*


    { path: '/my/companies', component: () => import('@/views/my/companies.vue') },
    { path: '/my/licences', component: () => import('@/views/my/licences.vue') },
    { path: '/my/manufacturers', component: () => import('@/views/my/manufacturers.vue') },
    { path: '/my/products', component: () => import('@/views/my/products.vue') },
    { path: '/my/shops', component: () => import('@/views/my/shops.vue') },
    { path: '/my/skus', component: () => import('@/views/my/skus.vue') },
    { path: '/my/standards', component: () => import('@/views/my/standards.vue') },
    { path: '/my/trademarks', component: () => import('@/views/my/trademarks.vue') },

    { path: '/my/favorites', component: () => import('@/views/my/favorites.vue') },


    ////////////////////////////////////////////////////////////////////////////
    // /(share)/* 

    // table

    { path: '/companies', component: () => import('@/views/share/companies/index.vue') },
    { path: '/licences', component: () => import('@/views/share/licences/index.vue') },
    { path: '/manufacturers', component: () => import('@/views/share/manufacturers/index.vue') },
    { path: '/products', component: () => import('@/views/share/products/index.vue') },
    { path: '/shops', component: () => import('@/views/share/shops/index.vue') },
    { path: '/skus', component: () => import('@/views/share/skus/index.vue') },
    { path: '/standards', component: () => import('@/views/share/standards/index.vue') },
    { path: '/trademarks', component: () => import('@/views/share/trademarks/index.vue') },


    // add

    { path: '/companies/add', component: () => import('@/views/share/companies/company-add-view.vue') },
    { path: '/licences/add', component: () => import('@/views/share/licences/licence-add-view.vue') },
    { path: '/manufacturers/add', component: () => import('@/views/share/manufacturers/manufacturer-add-view.vue') },
    { path: '/products/add', component: () => import('@/views/share/products/product-add-view.vue') },
    { path: '/shops/add', component: () => import('@/views/share/shops/shop-add-view.vue') },
    { path: '/skus/add', component: () => import('@/views/share/skus/sku-add-view.vue') },
    { path: '/standards/add', component: () => import('@/views/share/standards/standard-add-view.vue') },
    { path: '/trademarks/add', component: () => import('@/views/share/trademarks/trademark-add-view.vue') },

    // detail

    { path: '/companies/:id/detail', component: () => import('@/views/share/companies/company-detail-view.vue') },
    { path: '/licences/:id/detail', component: () => import('@/views/share/licences/licence-detail-view.vue') },
    { path: '/manufacturers/:id/detail', component: () => import('@/views/share/manufacturers/manufacturer-detail-view.vue') },
    { path: '/products/:id/detail', component: () => import('@/views/share/products/product-detail-view.vue') },
    { path: '/shops/:id/detail', component: () => import('@/views/share/shops/shop-detail-view.vue') },
    { path: '/skus/:id/detail', component: () => import('@/views/share/skus/sku-detail-view.vue') },
    { path: '/standards/:id/detail', component: () => import('@/views/share/standards/standard-detail-view.vue') },
    { path: '/trademarks/:id/detail', component: () => import('@/views/share/trademarks/trademark-detail-view.vue') },


    ////////////////////////////////////////////////////////////////////////////
    // /other/*


    { path: '/other/', component: () => import('@/views/AboutView.vue') }


  ],
})

export default router
