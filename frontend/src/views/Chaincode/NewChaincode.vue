<template>
  <div>
    <div v-if="this.progress === 1">
      <div class="row">
        <div class="col-md-3"></div>
        <div class="col-md-2"></div>
        <div class="col-md-2">
          <label>Domain</label>
        </div>
      </div>
      <div class="row">
        <div class="col-md-3"></div>
        <div class="col-md-2"></div>
        <div class="col-md-2">
          <select class="form-control" v-model="form.domain">
            <option v-for="network in this.networks">{{ network.domain }}</option>
          </select>
        </div>
        <div class="col-md-2">
          <base-button type="success" @click="handleDomain()">
            <i class="ni ni-bold-right"></i>
          </base-button>
        </div>
      </div>
    </div>
    <div class="row" v-if="this.progress === 2">
      <div class="col-md-5"></div>
      <div class="col-md-2">
        <base-input
          id="objName"
          placeholder="object name"
          @keyup.enter="handleAddObj()"
          v-model="newObj"
        ></base-input>
      </div>
      <div class="col-md-2">
        <base-button outline type="info" @click="handleAddObj()">
          <i class="ni ni-fat-add"></i>
        </base-button>
      </div>
    </div>
    <div class="row" v-if="this.progress === 2">
      <div v-for="(obj, index) in this.form.objects" class="col-xl-3 mt-5">
        <div class="card">
          <div class="card-header border-0">
            <div class="row align-items-center">
              <div class="col-xl-2">
                <base-button outline size="sm" type="danger" @click="handleRemoveObj(index)">
                  <i class="ni ni-fat-remove"></i>
                </base-button>
              </div>
              <div class="col-xl-5">
                <h3 class="mb-0">{{obj.name}}</h3>
              </div>
              <div v-if="form.objects.length > 1" class="col-xl-5">
                <base-dropdown>
                  <base-button
                    size="sm"
                    slot="title"
                    type="primary"
                    class="dropdown-toggle"
                  >+ Relationship</base-button>
                  <div
                    @click="handleAddRelationship(index, 'inclusion')"
                    class="dropdown-item"
                  >inclusion</div>
                  <div
                    @click="handleAddRelationship(index, 'dependence')"
                    class="dropdown-item"
                  >dependence</div>
                  <div @click="handleAddRelationship(index, 'owner')" class="dropdown-item">owner</div>
                  <div @click="handleAddRelationship(index, 'match')" class="dropdown-item">match</div>
                </base-dropdown>
              </div>
            </div>
          </div>

          <div class="table-responsive">
            <base-table thead-classes="thead-light" :data="obj.fields">
              <template slot-scope="{row}">
                <th scope="row" class="col-xl-5">{{row.name}}</th>
                <td class="col-xl-5">{{row.type}}</td>
                <td v-if="row.name !== capitalizeFirstLetter(obj.name) + 'ID'">
                  <base-button
                    outline
                    type="danger"
                    size="sm"
                    @click="handleRemoveField(index, row.name)"
                  >
                    <i class="ni ni-fat-remove"></i>
                  </base-button>
                </td>
              </template>
            </base-table>
          </div>
          <div v-if="obj.relationship !== ''" class="card-header border-0">
            <div class="row">
              <th class="col-xl-5">{{obj.relationship}}</th>
              <td class="col-xl-5">
                <base-dropdown>
                  <base-button
                    outline
                    size="sm"
                    slot="title"
                    type="secondary"
                    class="dropdown-toggle"
                  >Object</base-button>
                  <div
                    v-for="object in form.objects"
                    v-if="object.name !==  obj.name && !obj.relatedObj.includes(object.name)"
                    class="dropdown-item"
                    @click="handleRelationship(index, object.name)"
                  >{{object.name}}</div>
                </base-dropdown>
              </td>
              <td class="col-xl-2">
                <base-button
                  outline
                  type="danger"
                  size="sm"
                  @click="handleRemoveRelationship(index)"
                >
                  <i class="ni ni-fat-remove"></i>
                </base-button>
              </td>
            </div>
          </div>
          <div class="card-header border-0">
            <div class="row">
              <div class="col-xl-6">
                <base-input
                  id="fieldName"
                  placeholder="new field"
                  v-model="obj.newField"
                  @keyup.enter="handleAddNewField(index)"
                ></base-input>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row" v-if="this.progress === 2 && this.form.objects.length > 0">
      <div class="col-md-10"></div>
      <div class="col-md-2">
        <base-button type="success" @click="handleGenerate()">Generate Chaincode</base-button>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from "vuex";

export default {
  name: "new-chaincode-form",
  props: {
    type: {
      type: String
    },
    title: String
  },
  data() {
    return {
      form: {
        domain: "",
        objects: []
      },
      progress: 1,
      networks: [],
      newObj: "",
      newField: ""
    };
  },
  methods: {
    ...mapActions("user", [
      "submitChaincode",
      "getMyNetworks",
      "generateChaincode"
    ]),

    onSubmit() {
      this.addNetwork({ form: this.form });
    },

    handleDomain() {
      this.progress = 2;
    },

    handleAddObj() {
      if (this.newObj === "") {
        return;
      }

      this.newObj = this.newObj.toLowerCase();

      for (let i = 0; i < this.form.objects.length; i++) {
        if (this.form.objects[i].name === this.newObj) {
          return;
        }
      }
      this.form.objects.push({
        name: this.newObj,
        name_many: this.newObj + "s",
        create: {
          msp: [],
          function_name: "Create" + this.capitalizeFirstLetter(this.newObj)
        },
        update: {
          msp: [],
          function_name: "Update" + this.capitalizeFirstLetter(this.newObj)
        },
        delete: null,
        properties: [],
        inclusion: [],
        is_inclused: [],
        dependence: [],
        is_dependenced: [],
        owner: [],
        is_owned: [],
        match: [],
        is_matched: [],
        fields: [
          {
            name: this.capitalizeFirstLetter(this.newObj) + "ID",
            type: "String",
            typeRelationship: ""
          }
        ],
        relationship: "",
        relatedObj: []
      });
      this.newObj = "";
    },

    handleRemoveObj(index) {
      this.form.objects.splice(index, 1);
    },

    handleAddNewField(index) {
      if (this.form.objects[index].newField === "") {
        return;
      }
      this.form.objects[index].newField = this.form.objects[
        index
      ].newField.toLowerCase();
      this.form.objects[index].newField = this.capitalizeFirstLetter(
        this.form.objects[index].newField
      );

      for (let i = 0; i < this.form.objects[index].fields.length; i++) {
        if (
          this.form.objects[index].fields[i].name ===
          this.form.objects[index].newField
        ) {
          return;
        }
      }

      this.form.objects[index].fields.push({
        name: this.form.objects[index].newField,
        type: "String",
        typeRelationship: "none"
      });

      this.form.objects[index].properties.push({
        name: this.form.objects[index].newField,
        type: "string"
      });

      this.form.objects[index].newField = "";
    },

    handleRemoveField(index, fieldName) {
      let i = 0;
      for (i = 0; i < this.form.objects[index].fields.length; i++) {
        if (this.form.objects[index].fields[i].name === fieldName) {
          break;
        }
      }
      if (i == 0) {
        return;
      }

      if (this.form.objects[index].fields[i].typeRelationship == "none") {
        let j = 0;
        for (j = 0; j < this.form.objects[index].properties.length; j++) {
          if (this.form.objects[index].properties[j].name === fieldName) {
            break;
          }
        }

        this.form.objects[index].fields.splice(i, 1);
        this.form.objects[index].properties.splice(j, 1);
      } else if (
        this.form.objects[index].fields[i].typeRelationship == "inclusion"
      ) {
        let j = 0;
        for (j = 0; j < this.form.objects[index].inclusion.length; j++) {
          if (
            this.form.objects[index].inclusion[j].child_obj_name_many ===
            fieldName.toLowerCase()
          ) {
            break;
          }
        }

        this.form.objects[index].inclusion.splice(j, 1);

        let l = 0;

        for (l = 0; l < this.form.objects.length; l++) {
          if (this.form.objects[l].name_many === fieldName.toLowerCase()) {
            let k = 0;
            for (k = 0; k < this.form.objects[l].is_inclused.length; k++) {
              if (
                this.form.objects[l].is_inclused[k].parent_obj_name ===
                this.form.objects[index].name
              ) {
                break;
              }
            }

            this.form.objects[l].is_inclused.splice(k, 1);
            break;
          }
        }

        this.form.objects[index].fields.splice(i, 1);

        this.form.objects[index].relatedObj.splice(
          this.form.objects[index].relatedObj.indexOf(
            this.form.objects[l].name
          ),
          1
        );
        this.form.objects[l].relatedObj.splice(
          this.form.objects[l].relatedObj.indexOf(
            this.form.objects[index].name
          ),
          1
        );
      } else if (
        this.form.objects[index].fields[i].typeRelationship == "dependence"
      ) {
        let j = 0;
        for (j = 0; j < this.form.objects[index].dependence.length; j++) {
          if (
            this.form.objects[index].dependence[j].child_obj_name_many ===
            fieldName.toLowerCase()
          ) {
            break;
          }
        }

        this.form.objects[index].dependence.splice(j, 1);

        let l = 0;

        for (l = 0; l < this.form.objects.length; l++) {
          if (this.form.objects[l].name_many === fieldName.toLowerCase()) {
            let k = 0;
            for (k = 0; k < this.form.objects[l].is_dependenced.length; k++) {
              if (
                this.form.objects[l].is_dependenced[k].parent_obj_name ===
                this.form.objects[index].name
              ) {
                break;
              }
            }

            this.form.objects[l].is_dependenced.splice(k, 1);
            break;
          }
        }

        let m = 0;
        for (m = 0; m < this.form.objects[l].fields.length; m++) {
          if (
            this.form.objects[l].fields[m].name ===
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID"
          ) {
            break;
          }
        }

        this.form.objects[l].fields.splice(m, 1);
        this.form.objects[index].fields.splice(i, 1);

        this.form.objects[index].relatedObj.splice(
          this.form.objects[index].relatedObj.indexOf(
            this.form.objects[l].name
          ),
          1
        );
        this.form.objects[l].relatedObj.splice(
          this.form.objects[l].relatedObj.indexOf(
            this.form.objects[index].name
          ),
          1
        );
      } else if (
        this.form.objects[index].fields[i].typeRelationship == "owner"
      ) {
        let j = 0;
        for (j = 0; j < this.form.objects[index].owner.length; j++) {
          if (
            this.form.objects[index].owner[j].child_obj_name_many ===
            fieldName.toLowerCase()
          ) {
            break;
          }
        }

        this.form.objects[index].owner.splice(j, 1);

        let l = 0;

        for (l = 0; l < this.form.objects.length; l++) {
          if (this.form.objects[l].name_many === fieldName.toLowerCase()) {
            let k = 0;
            for (k = 0; k < this.form.objects[l].is_owned.length; k++) {
              if (
                this.form.objects[l].is_owned[k].parent_obj_name ===
                this.form.objects[index].name
              ) {
                break;
              }
            }

            this.form.objects[l].is_owned.splice(k, 1);
            break;
          }
        }

        let m = 0;
        for (m = 0; m < this.form.objects[l].fields.length; m++) {
          if (
            this.form.objects[l].fields[m].name ===
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID"
          ) {
            break;
          }
        }

        this.form.objects[l].fields.splice(m, 1);
        this.form.objects[index].fields.splice(i, 1);

        this.form.objects[index].relatedObj.splice(
          this.form.objects[index].relatedObj.indexOf(
            this.form.objects[l].name
          ),
          1
        );
        this.form.objects[l].relatedObj.splice(
          this.form.objects[l].relatedObj.indexOf(
            this.form.objects[index].name
          ),
          1
        );
      } else if (
        this.form.objects[index].fields[i].typeRelationship == "match"
      ) {
        let j = 0;
        for (j = 0; j < this.form.objects[index].owner.length; j++) {
          if (
            this.form.objects[index].match[j].des_obj_name_many ===
            fieldName.toLowerCase()
          ) {
            break;
          }
        }

        this.form.objects[index].match.splice(j, 1);

        let l = 0;

        for (l = 0; l < this.form.objects.length; l++) {
          if (this.form.objects[l].name_many === fieldName.toLowerCase()) {
            let k = 0;
            for (k = 0; k < this.form.objects[l].is_owned.length; k++) {
              if (
                this.form.objects[l].is_matched[k].sou_obj_name ===
                this.form.objects[index].name
              ) {
                break;
              }
            }

            this.form.objects[l].is_matched.splice(k, 1);
            break;
          }
        }

        let m = 0;
        for (m = 0; m < this.form.objects[l].fields.length; m++) {
          if (
            this.form.objects[l].fields[m].name ===
            this.capitalizeFirstLetter(this.form.objects[index].name_many)
          ) {
            break;
          }
        }

        this.form.objects[l].fields.splice(m, 1);
        this.form.objects[index].fields.splice(i, 1);

        this.form.objects[index].relatedObj.splice(
          this.form.objects[index].relatedObj.indexOf(
            this.form.objects[l].name
          ),
          1
        );
        this.form.objects[l].relatedObj.splice(
          this.form.objects[l].relatedObj.indexOf(
            this.form.objects[index].name
          ),
          1
        );
      }
    },

    handleAddRelationship(index, typeRelationship) {
      this.form.objects[index].relationship = typeRelationship;
    },

    handleRemoveRelationship(index) {
      this.form.objects[index].relationship = "";
    },

    handleRelationship(index, desObjName) {
      let i;
      for (i = 0; i < this.form.objects.length; i++) {
        if (this.form.objects[i].name === desObjName) {
          break;
        }
      }

      if (i == this.form.objects.length || i == index) {
        return;
      }

      if (this.form.objects[index].relationship === "inclusion") {
        this.form.objects[index].inclusion.push({
          child_obj_name: desObjName,
          child_obj_name_many: this.form.objects[i].name_many,
          add: {
            msp: [],
            function_name:
              "Add" +
              this.capitalizeFirstLetter(desObjName) +
              "To" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          },
          remove: {
            msp: [],
            function_name:
              "Remove" +
              this.capitalizeFirstLetter(desObjName) +
              "To" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          }
        });

        this.form.objects[i].is_inclused.push({
          parent_obj_name: this.form.objects[index].name,
          parent_obj_name_many: this.form.objects[index].name_many,
          parent_field: this.capitalizeFirstLetter(
            this.form.objects[i].name_many
          )
        });

        this.form.objects[index].fields.push({
          name: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          type: "[]String",
          typeRelationship: "inclusion"
        });
      } else if (this.form.objects[index].relationship === "dependence") {
        this.form.objects[index].dependence.push({
          child_obj_name: desObjName,
          child_obj_name_many: this.form.objects[i].name_many,
          field: this.capitalizeFirstLetter(this.form.objects[i].name_many)
        });

        this.form.objects[i].is_dependenced.push({
          parent_obj_name: this.form.objects[index].name,
          parent_obj_name_many: this.form.objects[index].name_many,
          parent_field: this.capitalizeFirstLetter(
            this.form.objects[i].name_many
          ),
          field:
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID"
        });

        this.form.objects[index].fields.push({
          name: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          type: "[]String",
          typeRelationship: "dependence"
        });

        this.form.objects[i].fields.push({
          name:
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID",
          type: "String",
          typeRelationship: "is_dependenced"
        });
      } else if (this.form.objects[index].relationship === "owner") {
        this.form.objects[index].owner.push({
          child_obj_name: desObjName,
          child_obj_name_many: this.form.objects[i].name_many,
          child_obj_field: this.capitalizeFirstLetter(
            this.form.objects[index].name_many
          ),
          add: {
            msp: [],
            function_name:
              "Add" +
              this.capitalizeFirstLetter(desObjName) +
              "To" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          },
          remove: {
            msp: [],
            function_name:
              "Remove" +
              this.capitalizeFirstLetter(desObjName) +
              "From" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          },
          change: {
            msp: [],
            function_name:
              "Change" +
              this.capitalizeFirstLetter(desObjName) +
              "Of" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          }
        });

        this.form.objects[i].is_owned.push({
          parent_obj_name: this.form.objects[index].name,
          parent_obj_name_many: this.form.objects[index].name_many,
          parent_field: this.capitalizeFirstLetter(
            this.form.objects[i].name_many
          ),
          field:
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID"
        });

        this.form.objects[index].fields.push({
          name: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          type: "[]String",
          typeRelationship: "owner"
        });

        this.form.objects[i].fields.push({
          name:
            this.capitalizeFirstLetter(this.form.objects[index].name) + "ID",
          type: "String",
          typeRelationship: "is_owned"
        });
      } else if (this.form.objects[index].relationship === "match") {
        this.form.objects[index].match.push({
          des_obj_name: desObjName,
          des_obj_name_many: this.form.objects[i].name_many,
          des_field: this.capitalizeFirstLetter(
            this.form.objects[index].name_many
          ),
          field: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          add: {
            msp: [],
            function_name:
              "Add" +
              this.capitalizeFirstLetter(desObjName) +
              "To" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          },
          remove: {
            msp: [],
            function_name:
              "Remove" +
              this.capitalizeFirstLetter(desObjName) +
              "From" +
              this.capitalizeFirstLetter(this.form.objects[index].name)
          }
        });

        this.form.objects[i].is_matched.push({
          sou_obj_name: this.form.objects[index].name,
          sou_obj_name_many: this.form.objects[index].name_many,
          sou_field: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          field: this.capitalizeFirstLetter(this.form.objects[index].name_many)
        });

        this.form.objects[index].fields.push({
          name: this.capitalizeFirstLetter(this.form.objects[i].name_many),
          type: "[]String",
          typeRelationship: "match"
        });

        this.form.objects[i].fields.push({
          name: this.capitalizeFirstLetter(this.form.objects[index].name_many),
          type: "[]String",
          typeRelationship: "is_matched"
        });
      }

      this.form.objects[index].relatedObj.push(desObjName);
      this.form.objects[i].relatedObj.push(this.form.objects[index].name);

      this.form.objects[index].relationship = "";
    },

    handleGenerate() {
      this.generateChaincode({ form: this.form });
    },

    capitalizeFirstLetter(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    }
  },

  async created() {
    let data = await this.getMyNetworks();
    this.progress = 1;

    for (let i = 0; i < data.length; i++) {
      if (
        data[i].status === "Implemented" ||
        data[i].status === "Installed chaincode"
      ) {
        this.networks.push(data[i]);
      }
    }
  }
};
</script>

<style>
#objName {
  text-transform: lowercase;
}

#fieldName {
  text-transform: lowercase;
}
</style>
