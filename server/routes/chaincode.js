/** @format */

const router = require('express').Router();
const { check, body, validationResult } = require('express-validator');
const Project = require('../models/Project');
const STATUS = require('../configs/constant').PROJECT_STATUS;
const fs = require('fs');
const mongoose = require('mongoose');
const exec = require('child_process').exec;
const chaincodeController = require('../controllers/chaincode');
const path = require('path');
const CHAINCODE_LANGUAGES = require('../configs/constant').CHAINCODE_LANGUAGES;

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

require('dotenv').config();

router.post('/upConfig', [body('domain').not().isEmpty().trim().escape()], async (req, res) => {
  let errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(422).json({ errors: errors.array() });
  }
  try {
    let domain = req.body.domain;
    let user = req.decoded.user;

    let network = await Project.findOne({ domain: domain });

    if (!network) {
      return res.status(400).json({
        success: false,
        msg: 'Domain does not exist!',
      });
    }

    if (network.owner !== user.username) {
      return res.status(403).json({
        success: false,
        msg: 'Permission Denied!',
      });
    }

    let orgs = [];
    for (let i = 0; i < network.orgs.length; i++) {
      orgs.push(network.orgs[i].name);
    }

    let config = req.body;

    let checkName = await chaincodeController.checkNameAndProperties(config);
    if (!checkName.success) {
      return res.status(400).json({
        success: false,
        msg: checkName.msg,
      });
    }
    config = checkName.config;

    let checkArray = await chaincodeController.checkArray(config);
    if (!checkArray.success) {
      return res.status(400).json({
        success: false,
        msg: checkArray.msg,
      });
    }

    let checkCreate = await chaincodeController.checkCreate(config, orgs);
    if (!checkCreate.success) {
      return res.status(400).json({
        success: false,
        msg: checkCreate.msg,
      });
    }

    config = checkCreate.config;

    let checkUpdate = await chaincodeController.checkUpdate(config, orgs);
    if (!checkUpdate.success) {
      return res.status(400).json({
        success: false,
        msg: checkUpdate.msg,
      });
    }

    config = checkUpdate.config;

    let checkInclusion = await chaincodeController.checkInclusion(config, orgs);
    if (!checkInclusion.success) {
      return res.status(400).json({
        success: false,
        msg: checkInclusion.msg,
      });
    }

    config = checkInclusion.config;

    let checkIsdependenced = await chaincodeController.checkIsDependenced(config, orgs);
    if (!checkIsdependenced.success) {
      return res.status(400).json({
        success: false,
        msg: checkIsdependenced.msg,
      });
    }
    config = checkIsdependenced.config;

    let checkMatch = await chaincodeController.checkMatch(config, orgs);
    if (!checkMatch.success) {
      return res.status(400).json({
        success: false,
        msg: checkMatch.msg,
      });
    }
    config = checkMatch.config;

    let checkOwner = await chaincodeController.checkOwner(config, orgs);
    if (!checkOwner.success) {
      return res.status(400).json({
        success: false,
        msg: checkOwner.msg,
      });
    }
    config = checkOwner.config;

    if (!fs.existsSync(process.env.PROJECT_DIR + '/config/' + domain + '/')) {
      fs.mkdirSync(process.env.PROJECT_DIR + '/config/' + domain + '/');
    }

    let jsonFormat = JSON.stringify(req.body, null, 2);

    fs.writeFile(
      process.env.PROJECT_DIR + '/config/' + domain + '/chaincode_json_format.json',
      jsonFormat,
      'utf8',
      function (err) {
        if (err) {
          return res.status(500).json({
            success: false,
            msg: `Can not write!`,
          });
        }
      }
    );

    let script = 'cd .. && ./gen_chaincode.sh ' + domain;

    exec(script, async (err) => {
      if (err) {
        return res.status(500).json({
          success: false,
          msg: 'Cannot generate chaincode now!',
        });
      }

      network.newChaincode = true;
      network.chaincodeLanguage = 1;
      network.chaincodeByConfig = true;
      await network.save();

      return res.status(200).json({
        success: true,
        msg: 'Generate chaincode successfully!',
      });
    });
  } catch (error) {
    return res.status(500).json({
      success: false,
      msg: 'Internal Server Error!',
    });
  }
});

router.post(
  '/installChaincode',
  [body('domain').not().isEmpty().trim().escape()],

  async (req, res) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let domain = req.body.domain;
      let user = req.decoded.user;
      let network = await Project.findOne({ domain: req.body.domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: 'Project does not exist!',
        });
      }

      if (user.username !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      if (network.status !== STATUS.IMPLEMENTED) {
        return res.status(409).json({
          success: false,
          msg: 'Project have not implemented yet!',
        });
      }

      let chaincodeVersion = network.chaincodeVersion + 1;
      chaincodeVersion = chaincodeVersion.toString() + '.0';

      let script = 'cd .. && ./install_chaincode.sh ' + domain + ' ' + chaincodeVersion;

      exec(script, async (err, stdout) => {
        network.chaincodeVersion = chaincodeVersion;
        network.chaincodeLog = stdout;

        if (err) {
          await network.save();
          let data = fs.readFileSync('../projects/' + domain + '/log.txt', 'utf8');
          data = data.toString();
          return res.status(500).json({
            success: false,
            msg: data,
          });
        }

        network.status = STATUS.INSTALLED_CHAINCODE;
        network.newChaincode = false;
        await network.save();

        return res.status(200).json({
          success: true,
          msg: 'Install chaincode successfully',
        });
      });
    } catch (error) {
      console.log(error);
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

router.post(
  '/genServerWithChaincodeByConfig',
  [body('domain').not().isEmpty().trim().escape()],

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
          msg: 'Domain Does Not Exist!',
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      let script = 'cd .. && ./gen_server_with_chaincode_by_config.sh ' + domain;

      exec(script, async (err) => {
        if (err) {
          return res.status(500).json({
            success: false,
            msg: 'Cannot generate server!',
          });
        }
        return res.status(200).json({
          success: true,
          msg: 'Install chaincode and generate server successfully!',
        });
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

router.post(
  '/upgradeChaincode',
  [body('domain').not().isEmpty().trim().escape()],

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
          msg: 'Domain does not exist!',
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      if (network.status !== STATUS.INSTALLED_CHAINCODE) {
        return res.status(400).json({
          success: false,
          msg: 'Must install chaincode before!',
        });
      }
      let chaincodeVersion = network.chaincodeVersion + 1;
      chaincodeVersion = chaincodeVersion.toString() + '.0';

      let script = 'cd .. && ./upgrade_chaincode.sh ' + domain + ' ' + chaincodeVersion;

      exec(script, async (err, stdout) => {
        network.chaincodeVersion = chaincodeVersion;
        network.chaincodeLog = stdout;

        if (err) {
          await network.save();
          let data = fs.readFileSync('../projects/' + domain + '/log.txt', 'utf8');
          data = data.toString();
          return res.status(500).json({
            success: false,
            msg: data,
          });
        }

        network.newChaincode = false;
        await network.save();
        return res.status(200).json({
          success: true,
          msg: 'Upgrade chaincode successfully',
        });
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

router.post(
  '/downloadServer',
  [body('domain').not().isEmpty().trim().escape()],

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
          msg: 'Domain does not exist!',
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      if (network.status !== STATUS.INSTALLED_CHAINCODE && network.status !== STATUS.IMPLEMENTED) {
        return res.status(400).json({
          success: false,
          msg: 'Must install chaincode before!',
        });
      }

      let file = process.env.PROJECT_DIR + '/projects/' + domain + '/server.zip';
      let fileName = path.basename(file);

      res.setHeader('Content-disposition', 'attachment; filename=' + fileName);

      res.setHeader('Content-type', 'application/zip');

      let filestream = fs.createReadStream(file);
      filestream.pipe(res);
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

router.post(
  '/downloadChaincode',
  [body('domain').not().isEmpty().trim().escape()],

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
          msg: 'Domain does not exist!',
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      if (network.status !== STATUS.INSTALLED_CHAINCODE && !network.newChaincode) {
        return res.status(400).json({
          success: false,
          msg: 'Must install chaincode before!',
        });
      }

      let fileName;
      if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.GO) {
        fileName = domain + '.go';
      } else if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.NODE) {
        fileName = domain + '.js';
      } else if (network.chaincodeLanguage == CHAINCODE_LANGUAGES.JAVA) {
        fileName = domain + '.java';
      }

      let file =
        process.env.PROJECT_DIR +
        '/projects/' +
        domain +
        '/network/deploy-kubernetes/config/chaincode/' +
        domain +
        '/' +
        fileName;

      let chaincodeFile = path.basename(file);

      res.setHeader('Content-disposition', 'attachment; filename=' + chaincodeFile);

      res.setHeader('Content-type', 'application/go');

      let filestream = fs.createReadStream(file);
      filestream.pipe(res);
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

router.post(
  '/uploadChaincodeFile',
  [body('domain').not().isEmpty().trim().escape()],
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
          msg: 'Domain does not exist!',
        });
      }

      if (network.owner !== user.username) {
        return res.status(403).json({
          success: false,
          msg: 'Permission Denied!',
        });
      }

      if (network.status !== STATUS.INSTALLED_CHAINCODE && network.status !== STATUS.IMPLEMENTED) {
        return res.status(400).json({
          success: false,
          msg: 'Project have not implemented yet!',
        });
      }

      if (!req.files) {
        return res.status(400).json({
          success: false,
          message: 'No file uploaded!',
        });
      }

      let chaincodeFile = req.files.chaincodeFile;

      if (chaincodeFile.mimetype !== 'text/x-go') {
        return res.status(400).json({
          success: false,
          message: 'Invalid file type!',
        });
      }

      chaincodeFile.mv(
        process.env.PROJECT_DIR +
          '/projects/' +
          domain +
          '/network/deploy-kubernetes/config/chaincode/' +
          domain +
          '/' +
          domain +
          '.go'
      );

      network.newChaincode = true;
      network.chaincodeLanguage = 1;
      network.chaincodeByConfig = false;

      await network.save();

      return res.status(200).json({
        success: true,
        msg: 'Upload chaincode successfully',
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: 'Internal Server Error!',
      });
    }
  }
);

module.exports = router;
