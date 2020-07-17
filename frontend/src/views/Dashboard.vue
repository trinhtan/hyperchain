<template>
  <div
    v-loading.fullscreen.lock="loading"
    element-loading-text="Loading..."
    element-loading-background="rgba(255, 255, 255, 0.7)"
  >
    <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
      <div class="row">
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Network Number"
            type="gradient-red"
            v-bind:sub-title="networkNumber"
            icon="ni ni-atom"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2">Latest:</span>
              <span class="text-nowrap">{{ latest }}</span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Org Number"
            type="gradient-info"
            v-bind:sub-title="orgNumber"
            icon="ni ni-chart-bar-32"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2"></span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Peer Number"
            type="gradient-orange"
            v-bind:sub-title="peerNumber"
            icon="ni ni-bulb-61"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2"></span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Total storage"
            type="gradient-green"
            v-bind:sub-title="totalStorage + ' Gi'"
            icon="ni ni-money-coins"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-danger mr-2"></span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
      </div>
    </base-header>

    <div class="container-fluid mt--7">
      <div class="row">
        <div class="col-xl-7">
          <card id="log" type="default" header-classes="bg-transparent">
            <div slot="header" class="row align-items-center">
              <div class="col">
                <h5 class="h3 text-white mb-0">Chaincode preview</h5>
              </div>
              <div
                v-if="this.detail.network.status == 2 || this.detail.network.status == 3"
                class="col"
              >
                <ul class="nav nav-pills justify-content-end">
                  <label size="sm" for="chaincodeFile" class="btn btn-primary"
                    >Upload chaincode file</label
                  >
                  <input
                    id="chaincodeFile"
                    type="file"
                    style="visibility:hidden;"
                    ref="file"
                    @change="chaincodeFileChange"
                  />
                </ul>
              </div>
            </div>

            <span v-html="this.detail.chaincode" style="white-space: pre" id="log">{{
              this.detail.chaincode
            }}</span>
          </card>
        </div>
        <div class="col-xl-5">
          <card id="log" type="default" header-classes="bg-transparent">
            <div slot="header" class="row align-items-center">
              <div class="col">
                <h5 class="h3 text-white mb-0">Logs</h5>
              </div>
              <div class="col">
                <ul class="nav nav-pills justify-content-end">
                  <li class="nav-item mr-2 mr-md-0">
                    <a
                      class="nav-link py-2 px-3"
                      href="#"
                      @click.prevent="handleChangeLog('network')"
                    >
                      <span class="d-none d-md-block">Network's Log</span>
                    </a>
                  </li>
                  <li class="nav-item">
                    <a
                      class="nav-link py-2 px-3"
                      href="#"
                      @click.prevent="handleChangeLog('chaincode')"
                    >
                      <span class="d-none d-md-block">Chaincode's Log</span>
                    </a>
                  </li>
                </ul>
              </div>
            </div>

            <span v-html="this.log" style="white-space: pre-line" id="log"></span>
          </card>
        </div>
      </div>

      <div class="row mt-3">
        <div class="col-xl-7 ml-3">
          <div class="row">
            <base-button
              type="primary"
              v-if="
                (this.detail.network.status === 2 || this.detail.network.status === 3) &&
                  this.detail.network.newChaincode
              "
              @click="handleApplyChaincode()"
              >Install Chaincode</base-button
            >
            <base-button
              type="info"
              v-if="this.detail.network.status === 3 || this.detail.network.newChaincode"
              @click="handleDownloadChaincode()"
              >Download Chaincode</base-button
            >
          </div>
        </div>
        <div>
          <base-button type="info" v-if="this.detail.network.status === 1" @click="handleUp()"
            >Up</base-button
          >
          <base-button
            type="danger"
            v-if="this.detail.network.status === 2 || this.detail.network.status === 3"
            @click="handleDown()"
            >Down</base-button
          >
          <base-button
            type="default"
            v-if="this.detail.network.status === 1"
            @click="handleRemove()"
            >Remove</base-button
          >
        </div>
      </div>

      <div class="row mt-5">
        <div class="col-xl-7 mb-5 mb-xl-0">
          <!-- <page-visits-table></page-visits-table> -->
          <template>
            <div class="card">
              <div class="card-header border-0">
                <div class="row align-items-center">
                  <div class="col">
                    <h3 class="mb-0">Information</h3>
                  </div>
                  <div
                    class="col text-right"
                    v-if="this.detail.network.status === 2 || this.detail.network.status === 3"
                  >
                    <base-button @click="handleDownloadServer()" size="sm" type="primary"
                      >Download Server</base-button
                    >
                  </div>
                </div>
              </div>
              <div class="table-responsive">
                <base-table thead-classes="thead-light" :data="this.detail.network.orgs">
                  <template slot="columns">
                    <th>Org name</th>
                    <th>Peer Number</th>
                    <th>Total Storage</th>
                    <th>Services</th>
                    <th></th>
                  </template>

                  <template slot-scope="{ row }">
                    <th scope="row">{{ row.name }}</th>
                    <td>{{ row.peerNumber }}</td>
                    <td>{{ row.peerNumber * row.storage }}</td>
                    <td>
                      <div v-for="service in detail.services">
                        <badge v-if="service.org == row.name" type="default">{{
                          service.service
                        }}</badge>
                      </div>
                    </td>
                    <td>
                      <base-button outline type="primary" size="sm" @click="chooseOrg(row.name)"
                        >...</base-button
                      >
                    </td>
                  </template>
                </base-table>
              </div>
            </div>
          </template>
        </div>

        <div class="col-xl-5">
          <!-- <social-traffic-table></social-traffic-table> -->
          <template>
            <div class="card">
              <div class="card-header border-0">
                <div class="row align-items-center">
                  <div class="col">
                    <h3 class="mb-0">Ports</h3>
                  </div>
                </div>
              </div>
              <div class="table-responsive">
                <base-table thead-classes="thead-light" :data="services">
                  <template slot="columns">
                    <th>Service</th>
                    <th>Port</th>
                    <th>Port Type</th>
                  </template>

                  <template slot-scope="{ row }">
                    <th scope="row">{{ row.service }}</th>
                    <td>{{ row.portNumber }}</td>
                    <td>
                      <div class="d-flex align-items-center">
                        <badge type="primary">{{ row.portType }}</badge>
                      </div>
                    </td>
                  </template>
                </base-table>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import * as chartConfigs from '@/components/Charts/config';
import BarChart from '@/components/Charts/BarChart';

import { mapActions } from 'vuex';

let Convert = require('ansi-to-html');
let convert = new Convert();

export default {
  components: {
    BarChart,
  },
  data() {
    return {
      networkNumber: '',
      orgNumber: '',
      peerNumber: '',
      totalStorage: '',
      latest: '',
      detail: null,
      log: '',
      myNetworks: [],
      org: '',
      services: [],
      loading: false,
    };
  },
  methods: {
    ...mapActions('user', [
      'getNetworkDetail',
      'getMyNetworks',
      'upNetwork',
      'downNetwork',
      'removeNetwork',
      'installChaincodeByConfig',
      'upgradeChaincodeByConfig',
      'downloadServer',
      'downloadChaincode',
      'uploadChaincodeFile',
      'installChaincode',
      'upgradeChaincode',
    ]),

    chooseOrg(name) {
      this.services = [];
      this.org = name;
      for (let i = 0; i < this.detail.services.length; i++) {
        if (this.detail.services[i].org === name) {
          this.services.push(this.detail.services[i]);
        }
      }
    },

    handleChangeLog(choice) {
      if (choice === 'network') {
        this.log = this.detail.network.networkLog;
      } else if (choice === 'chaincode') {
        this.log = this.detail.network.chaincodeLog;
      } else return;
    },

    handleUp() {
      this.loading = true;
      this.upNetwork({ domain: this.detail.network.domain });
    },

    handleDown() {
      this.loading = true;
      this.downNetwork({ domain: this.detail.network.domain });
    },

    handleRemove() {
      this.removeNetwork({ domain: this.detail.network.domain });
    },

    handleApplyChaincode() {
      if (this.detail.network.status == 2) {
        this.loading = true;
        if (this.detail.network.chaincodeByConfig) {
          this.installChaincodeByConfig({ domain: this.detail.network.domain });
        } else {
          this.installChaincode({ domain: this.detail.network.domain });
        }
      } else if (this.detail.network.status == 3) {
        this.loading = true;

        if (this.detail.network.chaincodeByConfig) {
          this.upgradeChaincodeByConfig({ domain: this.detail.network.domain });
        } else {
          this.upgradeChaincode({ domain: this.detail.network.domain });
        }
      } else return;
    },

    handleDownloadServer() {
      this.downloadServer({ domain: this.detail.network.domain });
    },

    handleDownloadChaincode() {
      this.downloadChaincode({ domain: this.detail.network.domain });
    },

    chaincodeFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;

      if (!files.length) return;

      var chaincodeFile = files[0];

      if (
        !chaincodeFile ||
        (chaincodeFile.type !== 'text/x-go' &&
          chaincodeFile.type !== 'text/javascript' &&
          chaincodeFile.type !== 'text/java')
      ) {
        return;
      }

      this.uploadChaincodeFile({
        domain: this.detail.network.domain,
        file: chaincodeFile,
      });
    },
  },
  mounted() {
    this.initBigChart(0);
  },
  async created() {
    this.detail = await this.getNetworkDetail({
      domain: this.$route.params.domain,
    });

    this.org = this.detail.network.orgs[0].name;
    for (let i = 0; i < this.detail.services.length; i++) {
      if (this.detail.services[i].portType === 1) {
        this.detail.services[i].portType = 'Peer External';
      }
      if (this.detail.services[i].portType === 2) {
        this.detail.services[i].portType = 'Peer Event';
      }
      if (this.detail.services[i].portType === 3) {
        this.detail.services[i].portType = 'CA External';
      }

      if (this.detail.services[i].portType === 4) {
        this.detail.services[i].portType = 'Orderer External';
      }

      if (this.detail.services[i].org == this.org) {
        this.services.push(this.detail.services[i]);
      }
    }

    this.myNetworks = await this.getMyNetworks();
    let latestDate = '';
    this.networkNumber = this.myNetworks.length.toString();
    let countNumber = 0;
    let countStorage = 0;
    let countOrg = 0;
    for (let i = 0; i < this.myNetworks.length; i++) {
      if (this.myNetworks[i].createDate > latestDate) {
        this.latest = this.myNetworks[i].domain;
        latestDate = this.myNetworks[i].createDate;
      }
      countOrg += this.myNetworks[i].orgs.length;
      for (let j = 0; j < this.myNetworks[i].orgs.length; j++) {
        countNumber += this.myNetworks[i].orgs[j].peerNumber;

        countStorage += this.myNetworks[i].orgs[j].storage * this.myNetworks[i].orgs[j].peerNumber;
      }
    }
    this.orgNumber = countOrg.toString();
    this.totalStorage = countStorage.toString();
    this.peerNumber = countNumber.toString();

    this.detail.network.networkLog = convert.toHtml(this.detail.network.networkLog);
    this.detail.network.networkLog = this.detail.network.networkLog.split('#00A').join('#3ac252');

    this.detail.network.chaincodeLog = convert.toHtml(this.detail.network.chaincodeLog);
    this.detail.network.chaincodeLog = this.detail.network.chaincodeLog
      .split('#00A')
      .join('#3ac252');

    this.log = this.detail.network.networkLog;

    this.detail.chaincode = this.detail.chaincode
      .split('func ')
      .join("<span style='color:#3ac252'>func </span>");
  },
};
</script>
<style>
.button-wrapper {
  position: relative;
  width: 150px;
  text-align: center;
}

.button-wrapper span.label {
  position: relative;
  z-index: 0;
  display: inline-block;
  width: 100%;
  background: #00bfff;
  cursor: pointer;
  color: #fff;
  padding: 10px 0;
  text-transform: uppercase;
  font-size: 12px;
}

#upload {
  display: inline-block;
  position: absolute;
  z-index: 1;
  width: 100%;
  height: 50px;
  top: 0;
  left: 0;
  opacity: 0;
  cursor: pointer;
}
#log {
  min-height: 480px;
  max-height: 480px;
  overflow-y: scroll;
  color: #8fcfef;
}
</style>
