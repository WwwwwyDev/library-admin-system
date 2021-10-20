import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/home',
    children: [{
      path: 'home',
      name: '首页',
      component: () => import('@/views/home/index'),
      meta: {
        title: '首页',
        icon: 'home'
      }
    }]
  },
  {
    path: '/owncenter',
    component: Layout,
    redirect: '/owncenter/index',
    hidden: true,
    children: [{
      path: 'index',
      name: '个人中心',
      component: () => import('@/views/owncenter/index'),
      meta: {
        title: '个人中心'
      }
    }]
  },
]

export const permissionRoutes = [{
    path: '/system',
    component: Layout,
    redirect: '/system/user',
    name: '系统管理',
    meta: {
      title: '系统管理',
      icon: 'system',
      roles: ['superadmin']
    },
    children: [{
      path: 'user',
      name: '用户管理',
      component: () => import('@/views/system/user/index'),
      meta: {
        title: '用户管理',
        icon: 'user'
      }
    }, {
      path: 'loginlog',
      name: '登录日志',
      component: () => import('@/views/system/loginlog/index'),
      meta: {
        title: '登录日志',
        icon: 'log'
      }
    }, ]
  },
  {
    path: '/book',
    component: Layout,
    redirect: '/book',
    meta: {
      title: '图书馆',
      icon: 'book',
      roles: ['superadmin', 'admin', 'bookadmin']
    },
    children: [{
      path: '',
      name: '藏书管理',
      component: () => import('@/views/book/index'),
      meta: {
        title: '藏书管理',
        icon: 'book'
      }
    },{
      path: 'type',
      name: '类目管理',
      component: () => import('@/views/book/type/index'),
      meta: {
        title: '类目管理',
        icon: 'type'
      }
    }]
  },
  {
    path: '/lend',
    component: Layout,
    redirect: '/lend/lend',
    children: [{
      path: 'lend',
      name: '借阅管理',
      component: () => import('@/views/lend/index'),
      meta: {
        title: '借阅管理',
        icon: 'lend',
        roles: ['superadmin', 'admin', 'lendadmin']
      }
    }]
  },
  {
    path: '/vip',
    component: Layout,
    redirect: '/vip',
    children: [{
      path: 'vip',
      name: '会员管理',
      component: () => import('@/views/vip/index'),
      meta: {
        title: '会员管理',
        icon: 'vip',
        roles: ['superadmin', 'admin', 'vipadmin']
      }
    }]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]
const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
