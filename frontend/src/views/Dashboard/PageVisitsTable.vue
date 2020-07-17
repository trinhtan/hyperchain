<template>
  <div class="card">
    <div class="card-header border-0">
      <div class="row align-items-center">
        <div class="col">
          <h3 class="mb-0">Information</h3>
        </div>
        <div class="col text-right">
          <a href="#!" class="btn btn-sm btn-primary">Download Server</a>
        </div>
      </div>
    </div>

    <div class="table-responsive">
      <base-table thead-classes="thead-light" :data="detail.orgs">
        <template slot="columns">
          <th>Org name</th>
          <th>Peer Number</th>
          <th>Total Storage</th>
          <th>Certificate Authority</th>
          <th></th>
        </template>

        <template slot-scope="{row}">
          <th scope="row">{{row.name}}</th>
          <td>{{row.peerNumber}}</td>
          <td>{{row.peerNumber * row.storage}}</td>
          <td>{{row.ca}}</td>
          <td>
            <base-button outline type="primary" size="sm">...</base-button>
          </td>
        </template>
      </base-table>
    </div>
  </div>
</template>
<script>
import { mapActions, mapState } from "vuex";

export default {
  name: "page-visits-table",
  data() {
    return {
      tableData: [
        {
          page: "/argon/",
          visitors: "4,569",
          unique: "340",
          bounceRate: "46,53%",
          bounceRateDirection: "up"
        },
        {
          page: "/argon/index.html",
          visitors: "3,985",
          unique: "319",
          bounceRate: "46,53%",
          bounceRateDirection: "down"
        },
        {
          page: "/argon/charts.html",
          visitors: "3,513",
          unique: "294",
          bounceRate: "36,49%",
          bounceRateDirection: "down"
        },
        {
          page: "/argon/tables.html",
          visitors: "2,050",
          unique: "147",
          bounceRate: "50,87%",
          bounceRateDirection: "up"
        },
        {
          page: "/argon/profile.html",
          visitors: "1,795",
          unique: "190",
          bounceRate: "46,53%",
          bounceRateDirection: "down"
        }
      ],
      detail: null
    };
  },
  computed: {
    ...mapState("user", ["networkDetail"])
  },
  methods: {
    ...mapActions("user", ["getNetworkDetail"])
  },
  async created() {
    this.detail = await this.getNetworkDetail({
      domain: this.$route.params.domain
    });
  }
};
</script>
<style>
</style>
