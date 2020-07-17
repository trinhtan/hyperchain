/** @format */

"use strict";

const argv = require("yargs").argv;
const {
  FileSystemWallet,
  Gateway,
  X509WalletMixin,
} = require("fabric-network");
const path = require("path");

/**
 * Register user for org
 * @param  {String} org  Org Name (default: student)
 * @param  {String} username User Name (required)
 */

async function main() {
  try {
    let username;
    let org;

    if (!argv.org) {
      throw new Error(`org cannot undefined`);
    } else {
      org = argv.org.toString();
    }

    if (!argv.username) {
      throw new Error(`username cannot undefined`);
    } else {
      username = argv.username.toString();
    }

    let msp = await changeCaseFirstLetter(org);
    let orgAdmin = "admin" + org;

    const ccpPath = path.resolve(
      __dirname,
      "..",
      "connection-files",
      `connection-${org}.json`
    );

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), `/wallet/wallet-${org}`);
    const wallet = new FileSystemWallet(walletPath);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.exists(username);
    if (userExists) {
      console.log(
        `An identity for the user ${username} already exists in the wallet-${org}`
      );
      return;
    }

    // Check to see if we've already enrolled the admin user.
    const adminExists = await wallet.exists(orgAdmin);
    if (!adminExists) {
      console.log(`Admin user ${orgAdmin} does not exist in the wallet`);
      console.log("Run the enrollAdmin.js application before retrying");
      return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet,
      identity: orgAdmin,
      discovery: { enabled: true, asLocalhost: true },
    });

    // Get the CA client object from the gateway for interacting with the CA.
    const ca = gateway.getClient().getCertificateAuthority();
    const adminIdentity = gateway.getCurrentIdentity();

    //Register the user, enroll the user, and import the new identity into the wallet.
    const secret = await ca.register(
      {
        affiliation: "",
        enrollmentID: username,
        role: "client",
        attrs: [{ name: "username", value: username, ecert: true }],
      },
      adminIdentity
    );

    const enrollment = await ca.enroll({
      enrollmentID: username,
      enrollmentSecret: secret,
    });

    const userIdentity = X509WalletMixin.createIdentity(
      `${msp}MSP`,
      enrollment.certificate,
      enrollment.key.toBytes()
    );

    await wallet.import(username, userIdentity);

    console.log(
      `Successfully registered and enrolled user ${username} and imported it into the wallet`
    );

    await gateway.disconnect();
    process.exit(0);
  } catch (error) {
    console.error(`Failed to register user: ${error}`);
    process.exit(1);
  }
}

function changeCaseFirstLetter(params) {
  if (typeof params === "string") {
    return params.charAt(0).toUpperCase() + params.slice(1);
  }
  return null;
}

main();