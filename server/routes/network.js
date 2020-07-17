/** @format */

const router = require("express").Router();
const { body, validationResult } = require("express-validator");
const Project = require("../models/Project");
const Port = require("../models/Port");

const STATUS = require("../configs/constant").PROJECT_STATUS;
const fs = require("fs");
const mongoose = require("mongoose");
const exec = require("child_process").exec;
const networkController = require("../controllers/network");
const USER_ROLES = require("../configs/constant").USER_ROLES;
const CHAINCODE_LANGUAGES = require("../configs/constant").CHAINCODE_LANGUAGES;
// const spawn = require('child_process').spawn;

require("dotenv").config();

// Connect database
mongoose.connect(
  process.env.MONGODB_URI,
  { useUnifiedTopology: true, useNewUrlParser: true },
  (error) => {
    if (error) console.log(error);
  }
);

mongoose.set("useCreateIndex", true);

require("dotenv").config();

router.post(
  "/",
  [body("domain").not().isEmpty().trim().escape()],
  async (req, res, next) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let user = req.decoded.user;

      if (user.role !== USER_ROLES.USER) {
        return res.status(403).json({
          msg: "Permission Denied!",
        });
      }

      if (!req.body.domain) {
        return res.status(400).json({
          success: false,
          msg: "Domain must be not null!",
        });
      }

      let domain = req.body.domain;

      let network = await Project.findOne({ domain: domain });

      if (network) {
        return res.status(409).json({
          success: false,
          msg: "Domain already exists!",
        });
      }

      let checkNetwork = await networkController.checkNetwork(req.body.orgs);

      if (!checkNetwork.success) {
        return res.status(400).json({
          success: false,
          msg: checkNetwork.msg,
        });
      }

      let config = req.body;

      config = await networkController.checkPeerPort(config);

      config = await networkController.checkCAPort(config);

      config = await networkController.checkOrdererPort(config);

      if (!fs.existsSync(process.env.PROJECT_DIR + "/config/" + domain + "/")) {
        fs.mkdirSync(process.env.PROJECT_DIR + "/config/" + domain + "/");
      }

      let jsonFormat = JSON.stringify(config, null, 2);

      fs.writeFile(
        process.env.PROJECT_DIR +
          "/config/" +
          domain +
          "/network_json_format.json",
        jsonFormat,
        "utf8",
        function (err) {
          if (err) {
            return res.status(500).json({
              success: false,
              msg: `Can not write!`,
            });
          }
        }
      );

      let orgs = [];

      for (var i = 0; i < config.orgs.length; i++) {
        orgs.push({
          name: config.orgs[i].name,
          peerNumber: Number(config.orgs[i].peer_number),
          storage: Number(config.orgs[i].storage),
        });
      }

      network = new Project({
        owner: user.username,
        domain: domain,
        status: STATUS.CONFIG_FORMAT,
        chaincodeVersion: 0,
        orgs: orgs,
        newChaincode: false,
        createDate: Date.now(),
      });

      await network.save();

      return res.status(200).json({
        success: true,
        msg: `Upload Network Configuration Successfully!`,
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

router.post(
  "/upNetwork",
  [body("domain").not().isEmpty().trim().escape()],

  async (req, res) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }
    try {
      let user = req.decoded.user;
      let domain = req.body.domain;
      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project does not exist!",
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      if (network.status !== STATUS.CONFIG_FORMAT) {
        return res.status(400).json({
          success: false,
          msg: "Project was implemented!",
        });
      }

      let publiIp = process.env.PUBLIC_IP;
      let script = "cd .. && ./up_network.sh " + domain + " " + publiIp;

      exec(script, async (err, stdout, stderr) => {
        if (err) {
          console.log(err);
          network.networkLog = stdout;
          await network.save();
          let data = fs.readFileSync(
            "../projects/" + domain + "/log.txt",
            "utf8"
          );
          return res.status(500).json({
            success: false,
            msg: data,
          });
        }

        network.status = STATUS.IMPLEMENTED;
        network.networkLog = stdout;
        await network.save();

        return res.status(200).json({
          success: true,
          msg: "Up Network Successfully!",
          log: stdout,
        });
      });
    } catch (error) {
      console.log(error);
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

router.post(
  "/downNetwork",
  [body("domain").not().isEmpty().trim().escape()],

  async (req, res) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let domain = req.body.domain;
      let user = req.decoded.user;
      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project does not exist!",
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      if (
        network.status !== STATUS.IMPLEMENTED &&
        network.status !== STATUS.INSTALLED_CHAINCODE
      ) {
        return res.status(400).json({
          success: false,
          msg: "Project Was Down!",
        });
      }

      let script = "cd .. && ./down_network.sh " + domain;

      exec(script, async (err, stdout) => {
        if (err) {
          network.networkLog = stdout;
          await network.save();
          return res.status(500).json({
            success: false,
            msg: data,
          });
        }

        network.networkLog = stdout;
        network.status = STATUS.CONFIG_FORMAT;
        network.newChaincode = false;
        network.chaincodeVersion = 0;
        network.chaincodeLog = "";
        await network.save();

        return res.status(200).json({
          success: true,
          msg: "Down Network Successfully!",
        });
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

router.post(
  "/removeNetwork",
  [body("domain").not().isEmpty().trim().escape()],

  async (req, res) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let domain = req.body.domain;
      let user = req.decoded.user;
      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project Does Not Exist!",
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      if (network.status !== STATUS.CONFIG_FORMAT) {
        return res.status(400).json({
          success: false,
          msg: "Project Was Implemented!",
        });
      }

      await Project.deleteOne({ domain: domain });
      await Port.deleteMany({ domain: domain });

      return res.status(200).json({
        success: true,
        msg: "Remove Network Successfully!",
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

router.get("/", async (req, res) => {
  try {
    let user = req.decoded.user;

    let networks = await Project.find({ owner: user.username });

    return res.status(200).json({
      success: true,
      msg: `Successfully!`,
      networks: networks,
    });
  } catch (error) {
    return res.status(500).json({
      success: false,
      msg: "Internal Server Error!",
    });
  }
});

router.get(
  "/:domain",

  async (req, res) => {
    try {
      let user = req.decoded.user;
      let domain = req.params.domain;

      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project does not exist!",
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      let services = await Port.find({ domain: domain });

      let chaincodeFile;
      if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.GO) {
        chaincodeFile = domain + ".go";
      } else if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.NODE) {
        chaincodeFile = domain + ".js";
      } else if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.JAVA) {
        chaincodeFile = domain + ".java";
      }

      let chaincode;

      if (
        network.newChaincode ||
        network.status == STATUS.INSTALLED_CHAINCODE
      ) {
        chaincode = fs.readFileSync(
          "../projects/" +
            domain +
            "/network/deploy-kubernetes/config/chaincode/" +
            domain +
            "/" +
            chaincodeFile,
          "utf8"
        );
      }

      return res.status(200).json({
        success: true,
        msg: `Successfully!`,
        network: network,
        chaincode: chaincode ? chaincode : "",
        services: services,
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

router.post(
  "/genServerWithoutChaincode",
  [body("domain").not().isEmpty().trim().escape()],

  async (req, res) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }
    try {
      let domain = req.body.domain;
      let user = req.decoded.user;
      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Domain Does Not Exist!",
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      let script = "cd .. && ./gen_server_without_chaincode.sh " + domain;

      exec(script, async (err) => {
        if (err) {
          return res.status(500).json({
            success: false,
            msg: "Cannot generate server!",
          });
        }
        return res.status(200).json({
          success: true,
          msg: "Up network and generate server successfully!",
        });
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: "Internal Server Error!",
      });
    }
  }
);

module.exports = router;
