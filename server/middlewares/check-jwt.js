/** @format */

const jwt = require('jsonwebtoken');
let secretJWT = require('../configs/secret').secret;
const User = require('../models/User');

module.exports = (req, res, next) => {
  let token = req.headers['x-access-token'] || req.headers['authorization'];

  if (!token) {
    return res.status(403).json({
      success: false,
      message: 'No token provided',
    });
  }

  if (token.startsWith('Bearer ')) {
    token = token.slice(7, token.length);
  }

  jwt.verify(token, secretJWT, (err, decoded) => {
    if (err) {
      return res.status(401).json({
        success: false,
        message: 'Failed to authentication token',
      });
    }

    let user = User.findOne({
      username: req.body.username,
      password: decoded.user.password,
    });

    if (!user) {
      return res.status(409).json({
        success: false,
        msg: 'User does not exist!',
      });
    }

    req.decoded = decoded;
    next();
  });
};
