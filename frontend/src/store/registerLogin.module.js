import { authService } from '../services/registerLogin.service';
import router from '../router';

const user = JSON.parse(localStorage.getItem('user'));
const state = user
  ? {
      status: { loggedIn: true },
      user,
    }
  : {
      status: {},
      user: null,
    };

const actions = {
  async login({ dispatch, commit }, { username, password }) {
    try {
      let user = await authService.login(username, password);
      if (user) {
        commit('loginSuccess', user);
        router.push('/network');
      }
    } catch (error) {
      commit('loginFailure', error);
      dispatch('alert/alertAuthor', error, { root: true });
    }
  },

  async register({ dispatch, commit }, { username, password, email }) {
    try {
      let response = await authService.register(username, password, email);
      if (response.success) {
        commit('registerSuccess');
        setTimeout(() => {
          dispatch(
            'alert/success',
            { alert: false, message: 'Registration successful' },
            { root: true }
          );
        });
        router.push('/login');
      }
    } catch (error) {
      commit('loginFailure', error);
      commit('registerFailure');
      //   dispatch('alert/alertAuthor', error, { root: true });
    }
  },

  logout() {
    authService.logout();
    router.push('/login');
  },
};

const mutations = {
  loginSuccess(state, user) {
    state.status = { loggedIn: true };
    state.user = user;
  },
  loginFailure(state) {
    state.status = {};
  },
  registerSuccess(state) {
    state.status = {};
  },
  registerFailure(state) {
    state.status = {};
  },
  logout(state) {
    state.status = {};
    state.user = null;
  },
};

export const registerLogin = {
  namespaced: true,
  state,
  actions,
  mutations,
};
