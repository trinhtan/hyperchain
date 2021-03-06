/** @format */

"use strict";

const FabricCAServices = require("fabric-ca-client");
const { FileSystemWallet, X509WalletMixin } = require("fabric-network");
const fs = require("fs");
const path = require("path");
const argv = require("yargs").argv;

/**
 * Create admin for Org
 * @param  {String} org - Org Name
 */

async function main() {
  try {
    let org;
    let username;

    if (!argv.org) {
      throw new Error(`org cannot undefined`);
    } else {
      org = argv.org.toString();
    }

    username = "admin" + org;

    let msp = await changeCaseFirstLetter(org);

    const ccpPath = path.resolve(
      __dirname,
      "..",
      "connection-files",
      `connection-${org}.json`
    );

    const ccpJSON = fs.readFileSync(ccpPath, "utf8");
    const ccp = JSON.parse(ccpJSON);

    // Create a new CA client for interacting with the CA.
    const caInfo = ccp.certificateAuthorities[`certificate-${org}-ca`];
    const caTLSCACerts = caInfo.tlsCACerts.pem;
    const ca = new FabricCAServices(
      caInfo.url,
      { trustedRoots: caTLSCACerts, verify: false },
      caInfo.caName
    );

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), `/wallet/wallet-${org}`);
    const wallet = new FileSystemWallet(walletPath);

    // Check to see if we've already enrolled the admin user.
    const adminExists = await wallet.exists(username);
    if (adminExists) {
      console.log(
        `An identity for the admin user ${username} already exists in the wallet`
      );
      return;
    }

    const enrollment = await ca.enroll({
      enrollmentID: "admin",
      enrollmentSecret: "adminpw",
    });
    const identity = await X509WalletMixin.createIdentity(
      `${msp}MSP`,
      enrollment.certificate,
      enrollment.key.toBytes()
    );
    await wallet.import(username, identity);
    console.log(
      `Successfully enrolled admin user ${username} and imported it into the wallet-${org}`
    );
    process.exit(0);
  } catch (error) {
    console.error(`Failed to enroll admin: ${error}`);
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