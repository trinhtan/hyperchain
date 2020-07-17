/** @format */

const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');
const cookieParser = require('cookie-parser');
const logger = require('morgan');
const mongoose = require('mongoose');
const helmet = require('helmet');
const checkJWT = require('./middlewares/check-jwt');
const fileUpload = require('express-fileupload');

var app = express();

require('dotenv').config();

const networkRoutes = require('./routes/network');
const chaincodeRoutes = require('./routes/chaincode');
const authRoutes = require('./routes/auth');
const adminRoutes = require('./routes/admin');

// Connect database
mongoose.connect(
  process.env.MONGODB_URI,
  { useUnifiedTopology: true, useNewUrlParser: true },
  (error) => {
    if (error) console.log(error);
  }
);
mongoose.set('useCreateIndex', true);

app.use(express.json({ limit: '5mb' }));

app.use(
  fileUpload({
    createParentPath: true,
  })
);

// security
app.use(helmet());

// show log
app.use(logger('dev'));

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cookieParser());

// set up cors to allow us to accept requests from our client
app.use(
  cors({
    origin: '*', // allow to server to accept request from different origin
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    credentials: true, // allow session cookie from browser to pass through
  })
);

// Set up routes
app.use('/auth', authRoutes);
app.use('/network', checkJWT, networkRoutes);
app.use('/chaincode', checkJWT, chaincodeRoutes);
app.use('/admin', checkJWT, adminRoutes);

// catch 404 and forward to error handler
app.use((req, res, next) => {
  const err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// error handler
app.use((err, req, res) => {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
