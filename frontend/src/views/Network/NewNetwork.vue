<template>
  <div>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-2">
        <label>Domain</label>
      </div>
    </div>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-2">
        <base-input
          v-model="form.domain"
          id="domain"
          type="text"
          alternative
          placeholder="domain"
        ></base-input>
      </div>
      <!-- <div class="col-md-6">
        <base-input alternative placeholder="Disabled" disabled></base-input>
      </div>-->
    </div>
    <div class="row" v-if="form.orgs.length > 0">
      <div class="col-md-3"></div>
      <div class="col-md-2">
        <label>Org Name</label>
      </div>
      <div class="col-md-2">
        <label>Peer Number</label>
      </div>
      <div class="col-md-2">
        <label>Storage / peer</label>
      </div>
    </div>
    <template v-for="(org, index) in form.orgs">
      <div class="row">
        <div class="col-md-3"></div>
        <div class="col-md-2">
          <base-input
            type="text"
            alternative
            placeholder="org name"
            v-model="form.orgs[index].name"
          ></base-input>
        </div>
        <div class="col-md-2">
          <base-input
            type="number"
            max="3"
            min="1"
            alternative
            placeholder="peer number"
            v-model="form.orgs[index].peer_number"
          ></base-input>
        </div>
        <div class="col-md-2">
          <base-input
            type="number"
            max="3"
            min="1"
            alternative
            placeholder="storage / peer"
            v-model="form.orgs[index].storage"
          ></base-input>
        </div>
        <div class="col-md-1">
          <label></label>
          <base-button type="danger" @click="removeOrg(index)">X</base-button>
        </div>
        <div class="col-md-3"></div>
      </div>
    </template>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-2"></div>
      <div class="col-md-2"></div>
      <div class="col-md-2">
        <base-button outline type="info" @click="addOrg()">Add Org</base-button>
        <base-button outline type="success" @click="onSubmit()">Submit</base-button>
      </div>
      <div class="col-md-3"></div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'new-network-form',
  props: {
    type: {
      type: String,
    },
    title: String,
  },
  data() {
    return {
      form: {
        domain: '',
        orgs: [{ name: '', storage: 1, peer_number: 1 }],
      },
    };
  },
  methods: {
    ...mapActions('user', ['addNetwork']),
    onSubmit() {
      this.addNetwork({ form: this.form });
    },
    addOrg() {
      this.form.orgs.push({
        name: '',
        peer_number: 1,
        storage: 1,
      });
    },
    removeOrg(index) {
      this.form.orgs.splice(index, 1);
    },
  },
};
</script>

<style></style>
