import Vue from 'vue';
import Router from 'vue-router';
import DashboardLayout from '@/layout/DashboardLayout';
import AuthLayout from '@/layout/AuthLayout';
Vue.use(Router);

export default new Router({
  linkExactActiveClass: 'active',
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: 'dashboard',
      component: DashboardLayout,
      children: [
        {
          path: '/network/:domain',
          name: 'dashboard',
          component: () => import('./views/Dashboard.vue'),
        },
        {
          path: '/icons',
          name: 'icons',
          component: () => import('./views/Icons.vue'),
        },
        {
          path: '/profile',
          name: 'profile',
          component: () => import('./views/UserProfile.vue'),
        },
        {
          path: '/maps',
          name: 'maps',
          component: () => import('./views/Maps.vue'),
        },
        {
          path: '/network',
          name: 'myNetworks',
          component: () => import('./views/MyNetworks.vue'),
        },
        {
          path: '/newNetwork',
          name: 'newNetwork',
          component: () => import('./views/NewNetwork.vue'),
        },
        {
          path: '/newChaincode',
          name: 'newChaincode',
          component: () => import('./views/NewChaincode.vue'),
        },
      ],
    },
    {
      path: '/',
      redirect: 'login',
      component: AuthLayout,
      children: [
        {
          path: '/login',
          name: 'login',
          component: () => import('./views/Login.vue'),
        },
        {
          path: '/register',
          name: 'register',
          component: () => import('./views/Register.vue'),
        },
      ],
    },
  ],
});
