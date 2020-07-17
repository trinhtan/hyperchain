/** @format */

const router = require('express').Router();
const jwt = require('jsonwebtoken');
const User = require('../models/User');
const bcrypt = require('bcryptjs');
const { body, validationResult } = require('express-validator');
let secretJWT = require('../configs/secret').secret;
const USER_ROLES = require('../configs/constant').USER_ROLES;

// Register
router.post(
  '/register',
  [
    body('email').isEmail(),
    body('username').not().isEmpty().trim().escape(),
    body('password').not().isEmpty().trim().isLength({ min: 6 }),
  ],
  async (req, res, next) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let user = await User.findOne({ username: req.body.username });

      if (user) {
        return res.status(409).json({
          success: false,
          msg: 'Username already exist!',
        });
      }

      user = await User.findOne({ email: req.body.email });

      if (user) {
        return res.status(409).json({
          success: false,
          msg: 'Email already exist!',
        });
      }

      user = new User({
        username: req.body.username,
        password: req.body.password,
        email: req.body.email,
        role: USER_ROLES.USER,
      });

      await user.save();

      return res.status(200).json({
        success: true,
        msg: 'Register Successfully!',
      });
    } catch (error) {
      return res.status(500).json({
        success: false,
        msg: error.message,
      });
    }
  }
);

// Login
router.post(
  '/login',
  [
    body('username').not().isEmpty().trim().escape(),
    body('password').not().isEmpty().trim().isLength({ min: 6 }),
  ],
  async (req, res, next) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(422).json({ errors: errors.array() });
    }

    try {
      let user = await User.findOne({ username: req.body.username });

      if (!user) {
        return res.status(403).json({
          success: false,
          msg: 'Username Or Password Is Incorrect!',
        });
      }

      let validPassword = await bcrypt.compare(req.body.password, user.password);

      if (!validPassword) {
        return res.status(403).json({
          success: false,
          msg: 'Username Or Password Is Incorrect!',
        });
      }

      let token = jwt.sign(
        {
          user: user,
        },
        secretJWT
      );

      return res.status(200).json({
        success: true,
        username: user.username,
        msg: 'Login Successfully!',
        token: token,
        role: user.role,
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
