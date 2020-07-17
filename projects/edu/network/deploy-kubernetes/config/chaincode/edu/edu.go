package main

import (
	"encoding/json"
	"fmt"
    "errors"

	"github.com/hyperledger/fabric/core/chaincode/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type SmartContract struct {
}

type Course struct {
	CourseID string
	Name	string
}

type Subject struct {
	SubjectID string
	Name	string
}

type Class struct {
	ClassID string
	Name	string
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "CreateCourse" {
		return CreateCourse(stub, args)
	}

	if function == "UpdateCourse" {
		return UpdateCourse(stub, args)
	}

	if function == "GetCourse"{
		return GetCourse(stub, args)
	}

	if function == "GetAllCourses" {
		return GetAllCourses(stub)
	}

	if function == "CreateSubject" {
		return CreateSubject(stub, args)
	}

	if function == "UpdateSubject" {
		return UpdateSubject(stub, args)
	}

	if function == "GetSubject"{
		return GetSubject(stub, args)
	}

	if function == "GetAllSubjects" {
		return GetAllSubjects(stub)
	}

	if function == "CreateClass" {
		return CreateClass(stub, args)
	}

	if function == "UpdateClass" {
		return UpdateClass(stub, args)
	}

	if function == "GetClass"{
		return GetClass(stub, args)
	}

	if function == "GetAllClasss" {
		return GetAllClasss(stub)
	}

	return shim.Error("Invalid Smart Contract function name!")
}

func CreateCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseID := args[0]
	Name := args[1]

	keyCourse := "Course-" + CourseID
	_, errObjExist := getCourse(stub, keyCourse)

	if errObjExist == nil {
		return shim.Error("This courseID already exists - " + CourseID + "!")
	}


	var course = Course{CourseID: CourseID, Name: Name, }

	courseAsBytes, errMarshal := json.Marshal(course)
	if errMarshal != nil {
		return shim.Error("Can not convert data of this course to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func UpdateCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseID := args[0]
	Name := args[1]

	keyCourse := "Course-" + CourseID
	course, errObjExist := getCourse(stub, keyCourse)

	if errObjExist != nil {
		return shim.Error("This course does not exist - " + CourseID + "!")
	}

	course.Name = Name

	courseAsBytes, errMarshal := json.Marshal(course)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)
	return shim.Success(nil)
}


func GetCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseID := args[0]

	keyCourse := "Course-" + CourseID
	courseAsBytes, errObjExist := stub.GetState(keyCourse)

	if errObjExist != nil {
		return shim.Error("Can not get course!")
	}

	if courseAsBytes == nil {
		return shim.Error("This course does not exist - " + CourseID + "!")
	}

	return shim.Success(courseAsBytes)
}

func GetAllCourses(stub shim.ChaincodeStubInterface) sc.Response {

	allCourses, errObjExist := getListCourses(stub)

	if errObjExist != nil {
		return shim.Error("Can not get all courses!")
	}

	defer allCourses.Close()

	var tlist []Course
	var i int

	for i = 0; allCourses.HasNext(); i++ {

		record, err := allCourses.Next()

		if err != nil {
			return shim.Success(nil)
		}

		course := Course{}
		json.Unmarshal(record.Value, &course)
		tlist = append(tlist, course)
	}

	jsonRow, errMarshal := json.Marshal(tlist)

	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListCourses(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Course-"
	endKey := "Course-zzzzzzzz"

	resultIter, errGetRange := stub.GetStateByRange(startKey, endKey)
	if errGetRange != nil {
		return nil, errGetRange
	}

	return resultIter, nil
}

func getCourse(stub shim.ChaincodeStubInterface, compoundKey string) (Course, error) {

	var course Course

	courseAsBytes, errState := stub.GetState(compoundKey)

	if errState != nil {
		return  course, errors.New("Can not get course!")
	}

	if courseAsBytes == nil {
		return  course, errors.New("This course does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(courseAsBytes, &course)

	return course, nil
}

func CreateSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	SubjectID := args[0]
	Name := args[1]

	keySubject := "Subject-" + SubjectID
	_, errObjExist := getSubject(stub, keySubject)

	if errObjExist == nil {
		return shim.Error("This subjectID already exists - " + SubjectID + "!")
	}


	var subject = Subject{SubjectID: SubjectID, Name: Name, }

	subjectAsBytes, errMarshal := json.Marshal(subject)
	if errMarshal != nil {
		return shim.Error("Can not convert data of this subject to bytes!")
	}

	stub.PutState(keySubject, subjectAsBytes)

	return shim.Success(nil)
}

func UpdateSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	SubjectID := args[0]
	Name := args[1]

	keySubject := "Subject-" + SubjectID
	subject, errObjExist := getSubject(stub, keySubject)

	if errObjExist != nil {
		return shim.Error("This subject does not exist - " + SubjectID + "!")
	}

	subject.Name = Name

	subjectAsBytes, errMarshal := json.Marshal(subject)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keySubject, subjectAsBytes)
	return shim.Success(nil)
}


func GetSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	SubjectID := args[0]

	keySubject := "Subject-" + SubjectID
	subjectAsBytes, errObjExist := stub.GetState(keySubject)

	if errObjExist != nil {
		return shim.Error("Can not get subject!")
	}

	if subjectAsBytes == nil {
		return shim.Error("This subject does not exist - " + SubjectID + "!")
	}

	return shim.Success(subjectAsBytes)
}

func GetAllSubjects(stub shim.ChaincodeStubInterface) sc.Response {

	allSubjects, errObjExist := getListSubjects(stub)

	if errObjExist != nil {
		return shim.Error("Can not get all subjects!")
	}

	defer allSubjects.Close()

	var tlist []Subject
	var i int

	for i = 0; allSubjects.HasNext(); i++ {

		record, err := allSubjects.Next()

		if err != nil {
			return shim.Success(nil)
		}

		subject := Subject{}
		json.Unmarshal(record.Value, &subject)
		tlist = append(tlist, subject)
	}

	jsonRow, errMarshal := json.Marshal(tlist)

	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListSubjects(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Subject-"
	endKey := "Subject-zzzzzzzz"

	resultIter, errGetRange := stub.GetStateByRange(startKey, endKey)
	if errGetRange != nil {
		return nil, errGetRange
	}

	return resultIter, nil
}

func getSubject(stub shim.ChaincodeStubInterface, compoundKey string) (Subject, error) {

	var subject Subject

	subjectAsBytes, errState := stub.GetState(compoundKey)

	if errState != nil {
		return  subject, errors.New("Can not get subject!")
	}

	if subjectAsBytes == nil {
		return  subject, errors.New("This subject does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(subjectAsBytes, &subject)

	return subject, nil
}

func CreateClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	ClassID := args[0]
	Name := args[1]

	keyClass := "Class-" + ClassID
	_, errObjExist := getClass(stub, keyClass)

	if errObjExist == nil {
		return shim.Error("This classID already exists - " + ClassID + "!")
	}


	var class = Class{ClassID: ClassID, Name: Name, }

	classAsBytes, errMarshal := json.Marshal(class)
	if errMarshal != nil {
		return shim.Error("Can not convert data of this class to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)

	return shim.Success(nil)
}

func UpdateClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "StudentMSP" &&  MSPID != "AcademyMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	ClassID := args[0]
	Name := args[1]

	keyClass := "Class-" + ClassID
	class, errObjExist := getClass(stub, keyClass)

	if errObjExist != nil {
		return shim.Error("This class does not exist - " + ClassID + "!")
	}

	class.Name = Name

	classAsBytes, errMarshal := json.Marshal(class)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)
	return shim.Success(nil)
}


func GetClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClassID := args[0]

	keyClass := "Class-" + ClassID
	classAsBytes, errObjExist := stub.GetState(keyClass)

	if errObjExist != nil {
		return shim.Error("Can not get class!")
	}

	if classAsBytes == nil {
		return shim.Error("This class does not exist - " + ClassID + "!")
	}

	return shim.Success(classAsBytes)
}

func GetAllClasss(stub shim.ChaincodeStubInterface) sc.Response {

	allClasss, errObjExist := getListClasss(stub)

	if errObjExist != nil {
		return shim.Error("Can not get all classs!")
	}

	defer allClasss.Close()

	var tlist []Class
	var i int

	for i = 0; allClasss.HasNext(); i++ {

		record, err := allClasss.Next()

		if err != nil {
			return shim.Success(nil)
		}

		class := Class{}
		json.Unmarshal(record.Value, &class)
		tlist = append(tlist, class)
	}

	jsonRow, errMarshal := json.Marshal(tlist)

	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListClasss(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Class-"
	endKey := "Class-zzzzzzzz"

	resultIter, errGetRange := stub.GetStateByRange(startKey, endKey)
	if errGetRange != nil {
		return nil, errGetRange
	}

	return resultIter, nil
}

func getClass(stub shim.ChaincodeStubInterface, compoundKey string) (Class, error) {

	var class Class

	classAsBytes, errState := stub.GetState(compoundKey)

	if errState != nil {
		return  class, errors.New("Can not get class!")
	}

	if classAsBytes == nil {
		return  class, errors.New("This class does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(classAsBytes, &class)

	return class, nil
}

func main() {

	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error createing new Smart Contract: %s", err)
	}
}