const mongoose = require('mongoose');
const { Schema } = mongoose;

const PortSchema = new Schema({
  portNumber: {
    type: String,
    trim: true,
    required: [true, "can't be blank"],
  },
  domain: {
    type: String,
  },
  org: {
    type: String,
  },
  service: {
    type: String,
  },
  portType: {
    type: Number,
  },
});

const Port = mongoose.model('Port', PortSchema);

module.exports = Port;
