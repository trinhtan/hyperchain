/** @format */

const router = require("express").Router();
const { check, body, validationResult } = require("express-validator");
const Project = require("../models/Project");
const Port = require("../models/Port");

const STATUS = require("../configs/constant").PROJECT_STATUS;
const fs = require("fs");
const mongoose = require("mongoose");
const exec = require("child_process").exec;
const networkController = require("../controllers/network");
const USER_ROLES = require("../configs/constant").USER_ROLES;

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

      if (user.role !== USER_ROLES.ADMIN) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project does not exist!",
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

      exec(script, async (err) => {
        if (err) {
          return res.status(500).json({
            success: false,
            msg: data,
          });
        }

        network.status = STATUS.CONFIG_FORMAT;
        network.chaincodeVersion = 0;
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

      if (user.role !== USER_ROLES.ADMIN) {
        return res.status(403).json({
          success: false,
          msg: "Permission Denied!",
        });
      }

      let network = await Project.findOne({ domain: domain });

      if (!network) {
        return res.status(409).json({
          success: false,
          msg: "Project Does Not Exist!",
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

module.exports = router;
