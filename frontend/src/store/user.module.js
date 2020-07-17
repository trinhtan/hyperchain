import { userService } from '../services/user.service';
import router from '../router';

const state = {
  myNetworks: [],
  networkDetail: null,
};

const actions = {
  async getMyNetworks({ dispatch, commit }) {
    try {
      let myNetworks = await userService.getMyNetworks();
      commit('getMyNetworks', myNetworks);
      return myNetworks;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async getNetworkDetail({ dispatch, commit }, { domain }) {
    try {
      let data = await userService.getNetworkDetail(domain);
      return data;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async upNetwork({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.upNetwork(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async downNetwork({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.downNetwork(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async removeNetwork({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.removeNetwork(domain);
      router.push('/network');
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async addNetwork({ dispatch, commit }, { form }) {
    try {
      let response = await userService.addNetwork(form);
      router.push('/network');
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async generateChaincode({ dispatch, commit }, { form }) {
    try {
      let response = await userService.generateChaincode(form);
      router.push('/network/' + form.domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async installChaincodeByConfig({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.installChaincodeByConfig(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async downloadServer({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.downloadServer(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async downloadChaincode({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.downloadChaincode(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async uploadChaincodeFile({ dispatch, commit }, { domain, file }) {
    try {
      let response = await userService.uploadChaincodeFile(domain, file);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async upgradeChaincodeByConfig({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.upgradeChaincodeByConfig(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async installChaincode({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.installChaincode(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },

  async upgradeChaincode({ dispatch, commit }, { domain }) {
    try {
      let response = await userService.upgradeChaincode(domain);
      return response;
    } catch (error) {
      if (error.response.status === 403) {
        router.push('/login');
      }
    }
  },
};

const mutations = {
  getMyNetworks(state, myNetworks) {
    for (let i = 0; i < myNetworks.length; i++) {
      if (myNetworks[i].status === 1) {
        myNetworks[i].status = 'Config format';
        myNetworks[i].statusType = 'warning';
        myNetworks[i].completion = 33;
      }
      if (myNetworks[i].status === 2) {
        myNetworks[i].status = 'Implemented';
        myNetworks[i].statusType = 'info';
        myNetworks[i].completion = 66;
      }
      if (myNetworks[i].status === 3) {
        myNetworks[i].status = 'Installed chaincode';
        myNetworks[i].statusType = 'success';
        myNetworks[i].completion = 100;
      }

      myNetworks[i].peerNumber = 0;
      for (let j = 0; j < myNetworks[i].orgs.length; j++) {
        myNetworks[i].peerNumber += myNetworks[i].orgs[j].peerNumber;
      }
    }
    state.myNetworks = myNetworks;
  },

  getNetworkDetail(state, network) {
    state.networkDetail = network;
  },
};

export const user = {
  namespaced: true,
  state,
  actions,
  mutations,
};

// function capitalizeFirstLetter(string) {
//   return string.charAt(0).toUpperCase() + string.slice(1);
// }
