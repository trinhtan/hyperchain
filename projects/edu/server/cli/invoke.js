/** @format */

"use strict";

const argv = require("yargs").argv;
const conn = require("../fabric/network");
const uuidv4 = require("uuid/v4");

/**
 * Invoke function of chaincode
 * @param  {String} orgMSP  Org Name (default: student)
 * @param  {String} func  Function Name (required)
 * @param  {String} username User Name (required)
 */

async function main() {
  try {
    if (!argv.func || !argv.username || !argv.org) {
      console.log(`Parameter func or username or org cannot undefined`);
      return;
    }

    let func = argv.func.toString();
    let username = argv.username.toString();
    let org = argv.org.toString();

    const networkObj = await conn.connectToNetwork(org, username, true);
    if (func === "CreateCourse") {

      let CourseID = uuidv4();
      let Name = argv.Name.toString();

      let courseCreated = {
        CourseID : CourseID ,
        Name: Name,
      };

      await conn.CreateCourse(networkObj, courseCreated);

    } else if (func === "CreateSubject") {

      let SubjectID = uuidv4();
      let Name = argv.Name.toString();

      let subjectCreated = {
        SubjectID : SubjectID ,
        Name: Name,
      };

      await conn.CreateSubject(networkObj, subjectCreated);

    } else if (func === "CreateClass") {

      let ClassID = uuidv4();
      let Name = argv.Name.toString();

      let classCreated = {
        ClassID : ClassID ,
        Name: Name,
      };

      await conn.CreateClass(networkObj, classCreated);

    } else if (func === "UpdateCourse") {

      let CourseID = argv.CourseID.toString();;
      let Name = argv.Name.toString();

      let courseUpdated = {
        CourseID : CourseID ,
        Name: Name,
      };

      await conn.UpdateCourse(networkObj, courseUpdated);
    } else if (func === "UpdateSubject") {

      let SubjectID = argv.SubjectID.toString();;
      let Name = argv.Name.toString();

      let subjectUpdated = {
        SubjectID : SubjectID ,
        Name: Name,
      };

      await conn.UpdateSubject(networkObj, subjectUpdated);
    } else if (func === "UpdateClass") {

      let ClassID = argv.ClassID.toString();;
      let Name = argv.Name.toString();

      let classUpdated = {
        ClassID : ClassID ,
        Name: Name,
      };

      await conn.UpdateClass(networkObj, classUpdated);
    } else {
      console.log("Failed!");
      process.exit(0);
    }
    console.log('Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to submit transaction: ${error}`);
    process.exit(1);
  }
}

main();