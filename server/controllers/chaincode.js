/** @format */

'use strict';
// const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const path = require('path');
const mongoose = require('mongoose');
const Port = require('../models/Port');
require('dotenv').config();

// Connect database
mongoose.connect(
  process.env.MONGODB_URI,
  { useUnifiedTopology: true, useNewUrlParser: true },
  (error) => {
    if (error) console.log(error);
  }
);

mongoose.set('useCreateIndex', true);

exports.checkNameAndProperties = async function (config) {
  try {
    for (let i = 0; i < config.objects.length; i++) {
      if (!config.objects[i].name) {
        throw new Error(`Org[${i}]'s name must be not null!`);
      }

      config.objects[i].name = config.objects[i].name.toLowerCase();

      if (!config.objects[i].name_many) {
        throw new Error(`Org[${i}]'s name_many must be not null!`);
      }

      config.objects[i].name_many = config.objects[i].name_many.toLowerCase();

      if (!Array.isArray(config.objects[i].properties)) {
        throw new Error(`${config.objects[i].name}'s properties must be array!`);
      }

      for (let j = 0; j < config.objects[i].properties.length; j++) {
        let nameCap = capitalizeFirstLetter(config.objects[i].properties[j].name);

        if (nameCap == capitalizeFirstLetter(config.objects[i].name) + 'Id') {
          throw new Error(`${config.objects[i].name}'s properties[${j}] must be not ${nameCap}!`);
        }

        config.objects[i].properties[j].name = nameCap;
        config.objects[i].properties[j].type = 'string';
      }
    }

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error.message,
    };
    return response;
  }
};

exports.checkCreate = async function (config, orgs) {
  try {
    let objects = config.objects;

    for (let i = 0; i < objects.length; i++) {
      if (!objects[i].create.function_name) {
        throw new Error(`${objects[i].name}'s create method function_name must be not null!`);
      }

      let msp = await this.checkMSP(objects[i].create, orgs);
      // objects[i].create.msp = msp;
      objects[i].create.msp = orgs;

      // config.objects[i] = objects[i];
    }
    config.objects = objects;
    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };
    return response;
  }
};

exports.checkUpdate = async function (config, orgs) {
  try {
    let objects = config.objects;

    for (let i = 0; i < objects.length; i++) {
      if (typeof objects[i].update === 'undefined') {
        throw new Error(`${objects[i].name}'s update method must be defined!`);
      }

      if (objects[i].update) {
        if (!objects[i].update.function_name) {
          throw new Error(`${objects[i].name}'s update method function_name must be not null!`);
        }

        let msp = await this.checkMSP(objects[i].update, orgs);
        // objects[i].update.msp = msp;
        objects[i].update.msp = orgs;
        config.objects[i] = objects[i];
      }
    }

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };
    return response;
  }
};

exports.checkInclusion = async function (config, orgs) {
  try {
    let objects = config.objects;
    for (let i = 0; i < objects.length; i++) {
      if (typeof objects[i].inclusion === 'undefined') {
        throw new Error(`${objects[i].name}'s inclusion must be defined!`);
      }
      if (objects[i].inclusion) {
        for (let j = 0; j < objects[i].inclusion.length; j++) {
          if (!objects[i].inclusion[j].child_obj_name) {
            throw new Error(
              `${objects[i].name}'s inclusion[${j}] - child_obj_name must be not null!`
            );
          }

          if (
            !objects[i].inclusion[j].add.function_name ||
            !objects[i].inclusion[j].remove.function_name
          ) {
            throw new Error(
              `Add and Remove method of ${objects[i].name}'s inclusion[${j}] must be difined`
            );
          }

          // objects[i].inclusion[j].add.msp = await this.checkMSP(objects[i].inclusion[j].add, orgs);

          // objects[i].inclusion[j].remove.msp = await this.checkMSP(
          //   objects[i].inclusion[j].remove,
          //   orgs
          // );

          objects[i].inclusion[j].add.msp = orgs;
          objects[i].inclusion[j].remove.msp = orgs;

          let childObjIndex = await this.getObjIndex(
            config,
            objects[i].inclusion[j].child_obj_name
          );

          let isInclusedIndex = await this.getIsInclusedIndex(
            objects[childObjIndex],
            objects[i].name
          );

          objects[i].inclusion[j].child_obj_name_many = objects[childObjIndex].name_many;
          objects[i].inclusion[j].field = capitalizeFirstLetter(objects[childObjIndex].name_many);

          objects[childObjIndex].is_inclused[isInclusedIndex].parent_field = capitalizeFirstLetter(
            objects[childObjIndex].name_many
          );

          objects[childObjIndex].is_inclused[isInclusedIndex].parent_obj_name_many =
            objects[i].name_many;
        }
      }
    }
    config.objects = objects;

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };

    return response;
  }
};

exports.checkIsDependenced = async function (config, orgs) {
  try {
    let objects = config.objects;
    for (let i = 0; i < objects.length; i++) {
      if (typeof objects[i].is_dependenced === 'undefined') {
        throw new Error(`${objects[i].name}'s is_dependenced must be defined!`);
      }
      if (objects[i].is_dependenced) {
        for (let j = 0; j < objects[i].is_dependenced.length; j++) {
          if (!objects[i].is_dependenced[j].parent_obj_name) {
            throw new Error(
              `${objects[i].name}'s is_dependenced[${j}]-parent_obj_name must be not null!`
            );
          }

          let parentObjIndex = await this.getObjIndex(
            config,
            objects[i].is_dependenced[j].parent_obj_name
          );

          let dependenceIndex = await this.getDependenceIndex(
            objects[parentObjIndex],
            objects[i].name
          );

          objects[i].is_dependenced[j].parent_obj_name_many = objects[parentObjIndex].name_many;

          objects[i].is_dependenced[j].parent_field = capitalizeFirstLetter(objects[i].name_many);

          objects[i].is_dependenced[j].field =
            capitalizeFirstLetter(objects[parentObjIndex].name) + 'Id';

          objects[parentObjIndex].dependence[dependenceIndex].field = capitalizeFirstLetter(
            objects[i].name_many
          );

          objects[parentObjIndex].dependence[dependenceIndex].child_obj_name_many =
            objects[i].name_many;
        }
      }
    }
    config.objects = objects;

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };

    return response;
  }
};

exports.checkMatch = async function (config, orgs) {
  try {
    let objects = config.objects;
    for (let i = 0; i < objects.length; i++) {
      if (typeof objects[i].match === 'undefined') {
        throw new Error(`${objects[i].name}'s match must be defined!`);
      }
      if (objects[i].match) {
        for (let j = 0; j < objects[i].match.length; j++) {
          if (!objects[i].match[j].des_obj_name) {
            throw new Error(`${objects[i].name}'s match[${j}]-des_obj_name must be not null!`);
          }

          if (!objects[i].match[j].add.function_name || !objects[i].match[j].remove.function_name) {
            throw new Error(
              `Add and Remove method of ${objects[i].name}'s match[${j}] must be difined`
            );
          }

          // objects[i].match[j].add.msp = await this.checkMSP(objects[i].match[j].add, orgs);

          // objects[i].match[j].remove.msp = await this.checkMSP(objects[i].match[j].remove, orgs);

          objects[i].match[j].add.msp = orgs;
          objects[i].match[j].remove.msp = orgs;

          let desObjIndex = await this.getObjIndex(config, objects[i].match[j].des_obj_name);

          let isMatchedIndex = await this.getIsMatchedIndex(objects[desObjIndex], objects[i].name);

          objects[i].match[j].des_obj_name_many = objects[desObjIndex].name_many;

          objects[i].match[j].field = capitalizeFirstLetter(objects[desObjIndex].name_many);

          objects[i].match[j].des_field = capitalizeFirstLetter(objects[i].name_many);

          objects[desObjIndex].is_matched[isMatchedIndex].sou_obj_name_many = objects[i].name_many;

          objects[desObjIndex].is_matched[isMatchedIndex].field = capitalizeFirstLetter(
            objects[i].name_many
          );

          objects[desObjIndex].is_matched[isMatchedIndex].sou_field = capitalizeFirstLetter(
            objects[desObjIndex].name_many
          );
        }
      }
    }
    config.objects = objects;

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };

    return response;
  }
};

exports.checkOwner = async function (config, orgs) {
  try {
    let objects = config.objects;
    for (let i = 0; i < objects.length; i++) {
      if (objects[i].owner) {
        for (let j = 0; j < objects[i].owner.length; j++) {
          if (!objects[i].owner[j].child_obj_name) {
            throw new Error(`${objects[i].name}'s owner[${j}]-des_obj_name must be not null!`);
          }

          if (
            !objects[i].owner[j].add.function_name ||
            !objects[i].owner[j].remove.function_name ||
            !objects[i].owner[j].change.function_name
          ) {
            throw new Error(
              `Add and Remove and Change method of ${objects[i].name}'s owner[${j}] must be difined`
            );
          }

          // objects[i].owner[j].add.msp = await this.checkMSP(objects[i].owner[j].add, orgs);

          // objects[i].owner[j].remove.msp = await this.checkMSP(objects[i].owner[j].remove, orgs);

          // objects[i].owner[j].change.msp = await this.checkMSP(objects[i].owner[j].change, orgs);

          objects[i].owner[j].add.msp = orgs;
          objects[i].owner[j].remove.msp = orgs;
          objects[i].owner[j].change.msp = orgs;

          let childObjIndex = await this.getObjIndex(config, objects[i].owner[j].child_obj_name);

          let isOwnedIndex = await this.getIsOwnedIndex(objects[childObjIndex], objects[i].name);

          objects[i].owner[j].child_obj_name_many = objects[childObjIndex].name_many;

          objects[i].owner[j].field = capitalizeFirstLetter(objects[childObjIndex].name_many);

          objects[i].owner[j].child_obj_field = capitalizeFirstLetter(objects[i].name) + 'Id';

          objects[childObjIndex].is_owned[isOwnedIndex].parent_field = capitalizeFirstLetter(
            objects[childObjIndex].name_many
          );

          objects[childObjIndex].is_owned[isOwnedIndex].field =
            capitalizeFirstLetter(objects[i].name) + 'Id';

          objects[childObjIndex].is_owned[isOwnedIndex].parent_obj_name_many = objects[i].name_many;
        }
      }
    }
    config.objects = objects;

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };

    return response;
  }
};

exports.checkArray = async function (config) {
  try {
    for (let i = 0; i < config.objects.length; i++) {
      if (!Array.isArray(config.objects[i].inclusion)) {
        throw new Error(`inclusion for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].is_inclused)) {
        throw new Error(`is_inclused for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].dependence)) {
        throw new Error(`dependence for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].is_dependenced)) {
        throw new Error(`is_dependenced for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].match)) {
        throw new Error(`match for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].is_matched)) {
        throw new Error(`is_matched for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].owner)) {
        throw new Error(`owner for ${config.objects[i].name} must be array!`);
      }

      if (!Array.isArray(config.objects[i].is_owned)) {
        throw new Error(`is_owned for ${config.objects[i].name} must be array!`);
      }
    }

    let response = {
      success: true,
      config: config,
      msg: 'Valid format',
    };
    return response;
  } catch (error) {
    let response = {
      success: false,
      config: config,
      msg: error.message,
    };

    return response;
  }
};

exports.checkMSP = async function (func, orgs) {
  if (func.msp == []) {
    return orgs;
  }
  if (!Array.isArray(func.msp)) {
    throw new Error(`MSP for ${func.function_name} must be array!`);
  }
  for (let i = 0; i < func.msp.length; i++) {
    if (!orgs.includes(func.msp[i])) {
      throw new Error(`${func.msp[i]} MSP for ${func.function_name} does not exist in orgs array!`);
    }
  }
  return func.msp;
};

exports.getIsInclusedIndex = async function (object, parentObjName) {
  for (let i = 0; i < object.is_inclused.length; i++) {
    if (object.is_inclused[i].parent_obj_name === parentObjName) {
      return i;
    }
  }
  throw new Error(`No ${object.name}-is_inclused match ${parentObjName}-inclusion`);
};

exports.getDependenceIndex = async function (object, childObjName) {
  for (let i = 0; i < object.dependence.length; i++) {
    if (object.dependence[i].child_obj_name === childObjName) {
      return i;
    }
  }
  throw new Error(`No ${object.name}-dependence match ${childObjName}-is_dependenced`);
};

exports.getIsOwnedIndex = async function (object, parentObjName) {
  for (let i = 0; i < object.is_owned.length; i++) {
    if (object.is_owned[i].parent_obj_name === parentObjName) {
      return i;
    }
  }
  throw new Error(`No ${object.name}-is_owned match ${parentObjName}-owner`);
};

exports.getIsMatchedIndex = async function (object, souObjName) {
  for (let i = 0; i < object.is_matched.length; i++) {
    if (object.is_matched[i].sou_obj_name === souObjName) {
      return i;
    }
  }
  throw new Error(`No ${object.name}-is_matched match ${souObjName}-match`);
};

exports.getObjIndex = async function (config, objName) {
  let objects = config.objects;
  for (let i = 0; i < objects.length; i++) {
    if (objects[i].name === objName) {
      return i;
    }
  }
  throw new Error(`Object ${objName} does not exist!`);
};

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}
