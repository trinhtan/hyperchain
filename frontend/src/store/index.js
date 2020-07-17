import Vue from 'vue';
import Vuex from 'vuex';
import { registerLogin } from './registerLogin.module';
import { user } from './user.module';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: { registerLogin, user },
});
