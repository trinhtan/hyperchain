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
	Subjects	[]string
}

type Subject struct {
	SubjectID string
	Name	string
	Des	string
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

	if function == "AddSubjectToCourse" {
		return AddSubjectToCourse(stub, args)
	}

	if function == "RemoveSubjectToCourse" {
		return RemoveSubjectToCourse(stub, args)
	}

	if function == "GetSubjectsOfCourse" {
		return GetSubjectsOfCourse(stub, args)
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

	return shim.Error("Invalid Smart Contract function name!")
}

func CreateCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
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

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
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


func AddSubjectToCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseID := args[0]
	SubjectID := args[1]

	keyCourse := "Course-" + CourseID
	course, errObjExist := getCourse(stub, keyCourse)

	if errObjExist != nil {
		return shim.Error("This course does not exist - " + CourseID + "!")
	}

	keySubject := "Subject-" + SubjectID
	_, errChildObjExist := getSubject(stub, keySubject)

	if errChildObjExist != nil {
		return shim.Error("This subject does not exist - " + SubjectID + "!")
	}

	course.Subjects = append(course.Subjects, SubjectID)

	courseAsBytes, errMarshal := json.Marshal(course)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func RemoveSubjectToCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	MSPID, errMSP := cid.GetMSPID(stub)

	if errMSP != nil {
		return shim.Error("Can not get MSPID!")
	}

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseID := args[0]
	SubjectID := args[1]

	keyCourse := "Course-" + CourseID
	course, errObjExist := getCourse(stub, keyCourse)

	if errObjExist != nil {
		return shim.Error("This course does not exist - " + CourseID + "!")
	}

	keySubject := "Subject-" + SubjectID
	_, errChildObjExist := getSubject(stub, keySubject)

	if errChildObjExist !=  nil {
		return shim.Error("This subject does not exist - " + SubjectID + "!")
	}

	var i int
	var lenSubjects = len(course.Subjects)
	for i = 0; i < lenSubjects; i++ {
		if course.Subjects[i] == SubjectID {
			break
		}
	}

	copy(course.Subjects[i:], course.Subjects[i+1:])
	course.Subjects[lenSubjects-1] = ""
	course.Subjects = course.Subjects[:lenSubjects-1]

	courseAsBytes, errMarshal := json.Marshal(course)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func GetSubjectsOfCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseID := args[0]

	keyCourse := "Course-" + CourseID
	course, errObjExist := getCourse(stub, keyCourse)

	if errObjExist != nil {
		return shim.Error("Can not get course!")
	}

	var result []Subject
	var i int

	for i = 0; i < len(course.Subjects); i++ {

		subject, errChildObjExist := getSubject(stub, "Subject-"+course.Subjects[i])
		if errChildObjExist != nil {
			return shim.Error("Subject does not exist - " + course.Subjects[i] + " !")
		}
		result = append(result, subject)
	}

	jsonRow, errMarshal := json.Marshal(result)
	if errMarshal != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
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

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	SubjectID := args[0]
	Name := args[1]
	Des := args[2]

	keySubject := "Subject-" + SubjectID
	_, errObjExist := getSubject(stub, keySubject)

	if errObjExist == nil {
		return shim.Error("This subjectID already exists - " + SubjectID + "!")
	}


	var subject = Subject{SubjectID: SubjectID, Name: Name, Des: Des, }

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

	if  MSPID != "AcademyMSP" &&  MSPID != "StudentMSP" {
		return shim.Error("Permission Denied!")
	}

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	SubjectID := args[0]
	Name := args[1]
	Des := args[2]

	keySubject := "Subject-" + SubjectID
	subject, errObjExist := getSubject(stub, keySubject)

	if errObjExist != nil {
		return shim.Error("This subject does not exist - " + SubjectID + "!")
	}

	subject.Name = Name
	subject.Des = Des

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

func main() {

	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error createing new Smart Contract: %s", err)
	}
}