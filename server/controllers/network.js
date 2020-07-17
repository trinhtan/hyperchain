/** @format */

'use strict';
const mongoose = require('mongoose');
const Port = require('../models/Port');
const PORT_TYPE = require('../configs/constant').PORT_TYPE;
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

exports.checkNetwork = async function (orgs) {
  try {
    let names = [];
    for (var i = 0; i < orgs.length; i++) {
      orgs[i].name = orgs[i].name.toLowerCase();
      if (!orgs[i].name) {
        throw new Error(`Organization name cannot null!`);
      }
      if (names.includes(orgs[i].name)) {
        throw new Error(`Organizations cannot have the same name`);
      }
      names.push(orgs[i].name);
      if (!orgs[i].peer_number || orgs[i].peer_number <= 0) {
        throw new Error(`Peer number of ${orgs[i].name} muse more than 0!`);
      }
      if (!orgs[i].storage) {
        throw new Error(`Storage of ${orgs[i].name} muse not be null!`);
      }
    }
    let response = {
      success: true,
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

exports.checkPeerPort = async function (config) {
  let orgs = config.orgs;
  for (var i = 0; i < orgs.length; i++) {
    orgs[i].peers = [];
    for (var j = 0; j < orgs[i].peer_number; j++) {
      for (var k = 30000; k <= 32000; k++) {
        let port = await Port.findOne({ portNumber: k.toString() });
        if (port) {
          continue;
        }

        orgs[i].peers.push({
          peer: 'peer' + j.toString(),
          port: k,
        });

        port = new Port({
          portNumber: k.toString(),
          domain: config.domain,
          org: orgs[i].name,
          portType: PORT_TYPE.PEER_EXTERNAL,
          service: 'peer' + j.toString(),
        });

        await port.save();

        break;
      }
    }

    config.orgs[i] = orgs[i];
  }
  orgs = config.orgs;

  for (var i = 0; i < orgs.length; i++) {
    for (var j = 0; j < orgs[i].peer_number; j++) {
      for (var k = 30000; k <= 32000; k++) {
        let eventPort = await Port.findOne({ portNumber: k.toString() });
        if (eventPort) {
          continue;
        }

        let peer = {
          peer: orgs[i].peers[j].peer,
          port: orgs[i].peers[j].port,
          event_port: k,
        };

        orgs[i].peers[j] = peer;

        eventPort = new Port({
          portNumber: k.toString(),
          domain: config.domain,
          org: orgs[i].name,
          service: 'peer' + j.toString(),
          portType: PORT_TYPE.PEER_EVENT,
        });

        await eventPort.save();

        break;
      }
    }

    config.orgs[i] = orgs[i];
  }

  return config;
};

exports.checkCAPort = async function (config) {
  let orgs = config.orgs;
  for (var i = 0; i < orgs.length; i++) {
    for (var k = 30000; k <= 32000; k++) {
      let port = await Port.findOne({ portNumber: k.toString() });
      if (port) {
        continue;
      }

      orgs[i].ca = {
        port: k,
        target_port: 7054,
      };

      port = new Port({
        portNumber: k.toString(),
        domain: config.domain,
        org: orgs[i].name,
        service: 'ca',
        portType: PORT_TYPE.CA_EXTERNAL,
      });

      await port.save();

      break;
    }

    config.orgs[i] = orgs[i];
  }

  return config;
};

exports.checkOrdererPort = async function (config) {
  for (var k = 30000; k <= 32000; k++) {
    let port = await Port.findOne({ portNumber: k.toString() });
    if (port) {
      continue;
    }

    config.orderer = {
      port: k,
    };

    port = new Port({
      portNumber: k.toString(),
      domain: config.domain,
      portType: PORT_TYPE.ORDERER_PORT,
      service: 'orderer',
    });
    await port.save();
    break;
  }

  return config;
};
