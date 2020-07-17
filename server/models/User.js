/** @format */

const mongoose = require("mongoose");
const { Schema } = mongoose;
const bcrypt = require("bcryptjs");

const UserSchema = new Schema({
  email: {
    type: String,
    trim: true,
    unique: true,
    lowercase: true,
    match: [/^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/, "format is invalid!"],
  },
  username: {
    type: String,
    trim: true,
    unique: true,
    lowercase: true,
    match: [/^[a-zA-Z0-9]+$/, "format is invalid!"],
  },
  password: {
    type: String,
  },
  role: {
    type: Number,
  },
});

UserSchema.pre("save", function (next) {
  const SALTROUNDS = 10; // or another integer in that ballpark
  var user = this;

  if (!user.isModified("password")) return next();

  bcrypt.genSalt(SALTROUNDS, (err, salt) => {
    if (err) {
      return next(err);
    }

    bcrypt.hash(user.password, SALTROUNDS, (error, hash) => {
      if (error) {
        return next(error);
      }

      user.password = hash;
      next();
    });
  });
});

const User = mongoose.model("User", UserSchema);

module.exports = User;
