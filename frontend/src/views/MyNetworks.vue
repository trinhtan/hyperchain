<template>
  <div>
    <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
      <!-- Card stats -->
      <div class="row">
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Network Number"
            type="gradient-red"
            :sub-title="networkNumber"
            icon="ni ni-atom"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2">
                <!-- <i class="fa fa-arrow-up"></i>  -->
                Latest:
              </span>
              <span class="text-nowrap">{{ latest }}</span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Org Number"
            type="gradient-info"
            :sub-title="orgNumber"
            icon="ni ni-chart-bar-32"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2">
                <!-- <i class="fa fa-arrow-up"></i> 54.8% -->
              </span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Peer Number"
            type="gradient-orange"
            :sub-title="peerNumber"
            icon="ni ni-bulb-61"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-success mr-2">
                <!-- <i class="fa fa-arrow-up"></i> 12.18% -->
              </span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
        <div class="col-xl-3 col-lg-6">
          <stats-card
            title="Total storage"
            type="gradient-green"
            :sub-title="totalStorage + ' Gi'"
            icon="ni ni-money-coins"
            class="mb-4 mb-xl-0"
          >
            <template slot="footer">
              <span class="text-danger mr-2">
                <!-- <i class="fa fa-arrow-down"></i> 5.72% -->
              </span>
              <span class="text-nowrap"></span>
            </template>
          </stats-card>
        </div>
      </div>
    </base-header>

    <div class="container-fluid mt--7">
      <!-- <div class="row">
        <div class="col">
          <networks-table title="Light Table"></networks-table>
        </div>
      </div>-->
      <div class="row mt-5">
        <div class="col">
          <!-- <networks-table type="dark" title="Dark Table"></networks-table> -->
          <MyNetworksTable type="dark" title="My Networks" />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import MyNetworksTable from "./Tables/MyNetworksTable";
import { mapActions, mapState } from "vuex";

export default {
  name: "MyNetworks",
  components: {
    MyNetworksTable
  },
  data() {
    return {
      networkNumber: "",
      orgNumber: "",
      peerNumber: "",
      totalStorage: "",
      latest: ""
    };
  },
  computed: {
    ...mapState("user", ["myNetworks"])
  },
  methods: {
    ...mapActions("user", ["getMyNetworks"])
  },
  async created() {
    await this.getMyNetworks();
    let latestDate = "";
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

        countStorage +=
          this.myNetworks[i].orgs[j].storage *
          this.myNetworks[i].orgs[j].peerNumber;
      }
    }
    this.orgNumber = countOrg.toString();
    this.totalStorage = countStorage.toString();
    this.peerNumber = countNumber.toString();
  }
};
</script>
<style></style>
