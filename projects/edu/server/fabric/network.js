/** @format */

"use strict";
const {
  FileSystemWallet,
  Gateway,
  X509WalletMixin,
} = require("fabric-network");
const path = require("path");
require("dotenv").config();

exports.connectToNetwork = async function (org, username, cli = false) {
  try {
    let identity = username;

    const ccpPath = path.resolve(
      __dirname,
      "..",
      "connection-files",
      `connection-${org}.json`
    );
    let walletPath = path.join(process.cwd(), `/cli/wallet/wallet-${org}`);

    if (cli) {
      walletPath = path.join(process.cwd(), `/wallet/wallet-${org}`);
    }

    const wallet = new FileSystemWallet(walletPath);
    const userExists = await wallet.exists(identity);

    let networkObj;

    if (!userExists) {
      let response = {};
      response.error =
        "An identity for the user " +
        identity +
        " does not exist in the wallet. Register " +
        identity +
        " first";
      return response;
    } else {
      const gateway = new Gateway();

      await gateway.connect(ccpPath, {
        wallet: wallet,
        identity: identity,
        discovery: { enabled: false, asLocalhost: false },
      });

      const network = await gateway.getNetwork("educhannel");
      const contract = await network.getContract("edu");

      networkObj = {
        contract: contract,
        network: network,
        gateway: gateway,
        username: username,
      };
    }

    return networkObj;
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    process.exit(1);
  }
};

exports.query = async function (networkObj, func, args) {
  let response = {
    success: false,
    msg: "",
  };
  try {
    if (Array.isArray(args)) {
      response.msg = await networkObj.contract.evaluateTransaction(
        func,
        ...args
      );

      await networkObj.gateway.disconnect();
      response.success = true;
      return response;
    } else if (args) {
      response.msg = await networkObj.contract.evaluateTransaction(func, args);

      await networkObj.gateway.disconnect();
      response.success = true;
      return response;
    } else {
      response.msg = await networkObj.contract.evaluateTransaction(func);

      await networkObj.gateway.disconnect();
      response.success = true;
      return response;
    }
  } catch (error) {
    response.success = false;
    response.msg = error;
    return response;
  }
};


exports.CreateCourse = async function (networkObj, courseCreated) {
  if (
    !courseCreated.CourseID ||
    !courseCreated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "CreateCourse",
      courseCreated.CourseID,
      courseCreated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Create Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};

exports.CreateSubject = async function (networkObj, subjectCreated) {
  if (
    !subjectCreated.SubjectID ||
    !subjectCreated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "CreateSubject",
      subjectCreated.SubjectID,
      subjectCreated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Create Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};

exports.CreateClass = async function (networkObj, classCreated) {
  if (
    !classCreated.ClassID ||
    !classCreated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "CreateClass",
      classCreated.ClassID,
      classCreated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Create Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};

exports.UpdateCourse = async function (networkObj, courseUpdated) {
  if (
    !courseUpdated.CourseID ||
    !courseUpdated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "UpdateCourse",
      courseUpdated.CourseID,
      courseUpdated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Update Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};

exports.UpdateSubject = async function (networkObj, subjectUpdated) {
  if (
    !subjectUpdated.SubjectID ||
    !subjectUpdated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "UpdateSubject",
      subjectUpdated.SubjectID,
      subjectUpdated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Update Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};

exports.UpdateClass = async function (networkObj, classUpdated) {
  if (
    !classUpdated.ClassID ||
    !classUpdated.Name
  ) {
    let response = {};
    response.error =
      "Error! You need to fill all fields before you can register!";
    return response;
  }

  try {
    await networkObj.contract.submitTransaction(
      "UpdateClass",
      classUpdated.ClassID,
      classUpdated.Name
    );

    await networkObj.gateway.disconnect();
    let response = {
      success: true,
      msg: "Update Successfully!",
    };

    return response;
  } catch (error) {
    let response = {
      success: false,
      msg: error,
    };
    return response;
  }
};
