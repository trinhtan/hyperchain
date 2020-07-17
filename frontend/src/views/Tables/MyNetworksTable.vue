<template>
  <div class="card shadow" :class="type === 'dark' ? 'bg-default' : ''">
    <div class="card-header border-0" :class="type === 'dark' ? 'bg-transparent' : ''">
      <div class="row align-items-center">
        <div class="col">
          <h3 class="mb-0" :class="type === 'dark' ? 'text-white' : ''">{{ title }}</h3>
        </div>
        <div class="col text-right">
          <router-link to="/newNetwork">
            <base-button type="primary" size="sm">New Network</base-button>
          </router-link>
        </div>
      </div>
    </div>

    <div class="table-responsive">
      <base-table
        class="table align-items-center table-flush"
        :class="type === 'dark' ? 'table-dark' : ''"
        :thead-classes="type === 'dark' ? 'thead-dark' : 'thead-light'"
        tbody-classes="list"
        :data="tableData"
      >
        <template slot="columns">
          <th>Domain</th>
          <th>Create Date</th>
          <th>Org Number</th>
          <th>Peer Number</th>
          <th>Chaincode version</th>
          <th>Status</th>
          <th>Completion</th>
          <th></th>
        </template>

        <template slot-scope="{ row }">
          <th scope="row">
            <div class="media align-items-center">
              <div class="media-body">
                <span class="name mb-0 text-sm">{{ row.domain }}</span>
              </div>
            </div>
          </th>
          <td class="budget">{{ row.createDate }}</td>
          <td class="budget">{{ row.orgs.length }}</td>
          <td class="budget">{{ row.peerNumber }}</td>
          <td class="budget">{{ row.chaincodeVersion }}</td>
          <td>
            <badge class="badge-dot mr-4" :type="row.statusType">
              <i :class="`bg-${row.statusType}`"></i>
              <span class="status">{{ row.status }}</span>
            </badge>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{ row.completion }}%</span>
              <div>
                <base-progress
                  :type="row.statusType"
                  :show-percentage="false"
                  class="pt-0"
                  :value="row.completion"
                />
              </div>
            </div>
          </td>

          <td class="text-right">
            <template>
              <router-link :to="'/network/' + row.domain">
                <base-button outline type="primary" size="sm" title="More Information"
                  >...</base-button
                >
              </router-link>
            </template>
          </td>
        </template>
      </base-table>
    </div>

    <div
      class="card-footer d-flex justify-content-end"
      :class="type === 'dark' ? 'bg-transparent' : ''"
    >
      <base-pagination :total="5"></base-pagination>
    </div>
  </div>
</template>
<script>
import { mapActions, mapState } from 'vuex';
export default {
  name: 'MyNetworksTable',
  props: {
    type: {
      type: String,
    },
    title: String,
  },
  data() {
    return {
      tableData: [],
    };
  },
  computed: {
    ...mapState('user', ['myNetworks']),
  },
  watch: {
    myNetworks: function() {
      this.tableData = this.myNetworks;
    },
  },
  methods: {
    ...mapActions('user', ['getMyNetworks']),
  },
};
</script>
<style></style>
