import Vue from 'vue'
import Router from 'vue-router'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _bb2594ac = () => interopDefault(import('..\\pages\\comment.vue' /* webpackChunkName: "pages_comment" */))
const _e71ddb18 = () => interopDefault(import('..\\pages\\login.vue' /* webpackChunkName: "pages_login" */))
const _18a381e8 = () => interopDefault(import('..\\pages\\register.vue' /* webpackChunkName: "pages_register" */))

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: decodeURI('/'),
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/comment",
    component: _bb2594ac,
    name: "comment"
  }, {
    path: "/login",
    component: _e71ddb18,
    name: "login"
  }, {
    path: "/register",
    component: _18a381e8,
    name: "register"
  }],

  fallback: false
}

export function createRouter () {
  return new Router(routerOptions)
}
