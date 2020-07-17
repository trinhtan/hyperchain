/** @format */

const mongoose = require('mongoose');
const { Schema } = mongoose;

const ProjectSchema = new Schema({
  owner: {
    type: String,
    trim: true,
    lowercase: true,
    match: [/^[a-zA-Z0-9]+$/, 'is invalid'],
  },

  domain: {
    type: String,
    trim: true,
    required: [true, "can't be blank"],
    match: [/^[a-zA-Z]+$/, 'is invalid'],
  },

  status: {
    type: Number,
  },

  chaincodeVersion: {
    type: Number,
    trim: true,
    required: [true, "can't be blank"],
  },

  newChaincode: {
    type: Boolean,
    required: [true, "can't be blank"],
  },

  chaincodeLanguage: {
    type: Number,
  },

  chaincodeByConfig: {
    type: Boolean,
  },

  createDate: {
    type: Date,
  },

  orgs: {
    type: [Object],
    required: true,
  },

  networkLog: {
    type: String,
  },

  chaincodeLog: {
    type: String,
  },
});

const Project = mongoose.model('Project', ProjectSchema);

module.exports = Project;
