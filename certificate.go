package main

import (
	"encoding/json"
	"fmt"
    "errors"

	// "github.com/hyperledger/fabric/core/chaincode/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type SmartContract struct {
}

type Course struct {
	CourseId	string
	Description	string
	CourseName	string
	Subjects	[]string
	Certificates	[]string
	Students	[]string
}

type Subject struct {
	SubjectId	string
	Description	string
	SubjectName	string
	Classes	[]string
}

type Class struct {
	ClassId	string
	Description	string
	ClassName	string
	Scores	[]string
	SubjectId	string
	TeacherId	string
}

type Teacher struct {
	TeacherId	string
	Description	string
	TeacherName	string
	Classes	[]string
}

type Student struct {
	StudentId	string
	Description	string
	StudentName	string
	Certificates	[]string
	Courses	[]string
}

type Score struct {
	ScoreId	string
	Value	string
	ClassId	string
}

type Certificate struct {
	CertificateId	string
	Type	string
	CourseId	string
	StudentId	string
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

	if function == "RemoveSubjectFromCourse" {
		return RemoveSubjectFromCourse(stub, args)
	}

	if function == "GetSubjectsOfCourse" {
		return GetSubjectsOfCourse(stub, args)
	}

	if function == "GetCertificatesOfCourse" {
		return GetCertificatesOfCourse(stub, args)
	}

	if function == "GetStudentsOfCourse" {
		return GetStudentsOfCourse(stub, args)
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

	if function == "GetClassesOfSubject" {
		return GetClassesOfSubject(stub, args)
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

	if function == "GetScoresOfClass" {
		return GetScoresOfClass(stub, args)
	}

	if function == "GetSubjectOfClass" {
		return GetSubjectOfClass(stub, args)
	}
	if function == "GetTeacherOfClass" {
		return GetTeacherOfClass(stub, args)
	}

	if function == "GetClass"{
		return GetClass(stub, args)
	}

	if function == "GetAllClasses" {
		return GetAllClasses(stub)
	}

	if function == "CreateTeacher" {
		return CreateTeacher(stub, args)
	}

	if function == "UpdateTeacher" {
		return UpdateTeacher(stub, args)
	}
	if function == "GetClassesOfTeacher" {
		return GetClassesOfTeacher(stub, args)
	}

	if function == "AssignClassToTeacher" {
		return AssignClassToTeacher(stub, args)
	}

	if function == "UnassignClassFromTeacher" {
		return UnassignClassFromTeacher(stub, args)
	}

	if function == "ChangeTeacher" {
		return ChangeTeacher(stub, args)
	}

	if function == "GetTeacher"{
		return GetTeacher(stub, args)
	}

	if function == "GetAllTeachers" {
		return GetAllTeachers(stub)
	}

	if function == "CreateStudent" {
		return CreateStudent(stub, args)
	}

	if function == "UpdateStudent" {
		return UpdateStudent(stub, args)
	}

	if function == "GetCertificatesOfStudent" {
		return GetCertificatesOfStudent(stub, args)
	}

	if function == "StudentResgisterCourse" {
		return StudentResgisterCourse(stub, args)
	}

	if function == "StudentCancelCourse" {
		return StudentCancelCourse(stub, args)
	}

	if function == "GetCoursesOfStudent" {
		return GetCoursesOfStudent(stub, args)
	}

	if function == "GetStudent"{
		return GetStudent(stub, args)
	}

	if function == "GetAllStudents" {
		return GetAllStudents(stub)
	}

	if function == "CreateScore" {
		return CreateScore(stub, args)
	}

	if function == "UpdateScore" {
		return UpdateScore(stub, args)
	}

	if function == "GetClassOfScore" {
		return GetClassOfScore(stub, args)
	}

	if function == "GetScore"{
		return GetScore(stub, args)
	}

	if function == "GetAllScores" {
		return GetAllScores(stub)
	}

	if function == "CreateCertificate" {
		return CreateCertificate(stub, args)
	}

	if function == "GetCourseOfCertificate" {
		return GetCourseOfCertificate(stub, args)
	}

	if function == "GetStudentOfCertificate" {
		return GetStudentOfCertificate(stub, args)
	}

	if function == "GetCertificate"{
		return GetCertificate(stub, args)
	}

	if function == "GetAllCertificates" {
		return GetAllCertificates(stub)
	}

	return shim.Error("Invalid Smart Contract function name!")
}

func CreateCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	CourseId := args[0]
	Description := args[1]
	CourseName := args[2]
	keyCourse := "Course-" + CourseId
	_, err := getCourse(stub, keyCourse)

	if err == nil {
		return shim.Error("This courseId already exists - " + CourseId + "!")
	}

	var course = Course{CourseId: CourseId, Description: Description, CourseName: CourseName}

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convert data of this course to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func UpdateCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	CourseId := args[0]
	Description := args[1]
	CourseName := args[2]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("This course does not exist - " + CourseId + "!")
	}

	course.Description = Description
	course.CourseName = CourseName

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)
	return shim.Success(nil)
}


func AddSubjectToCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseId := args[0]
	SubjectId := args[1]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("This course does not exist - " + CourseId + "!")
	}

	keySubject := "Subject-" + SubjectId
	_, err = getSubject(stub, keySubject)

	if err != nil {
		return shim.Error("This subject does not exist - " + SubjectId + "!")
	}

	course.Subjects = append(course.Subjects, SubjectId)

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func RemoveSubjectFromCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	CourseId := args[0]
	SubjectId := args[1]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("This course does not exist - " + CourseId + "!")
	}

	keySubject := "Subject-" + SubjectId
	_, err = getSubject(stub, keySubject)

	if err !=  nil {
		return shim.Error("This subject does not exist - " + SubjectId + "!")
	}

	var i int
	var lenSubjects = len(course.Subjects)
	for i = 0; i < lenSubjects; i++ {
		if course.Subjects[i] == SubjectId {
			break
		}
	}

	copy(course.Subjects[i:], course.Subjects[i+1:])
	course.Subjects[lenSubjects-1] = ""
	course.Subjects = course.Subjects[:lenSubjects-1]

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func GetSubjectsOfCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseId := args[0]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course!")
	}

	var result []Subject
	var i int

	for i = 0; i < len(course.Subjects); i++ {

		subject, err := getSubject(stub, "Subject-"+course.Subjects[i])
		if err != nil {
			return shim.Error("Subject does not exist - " + course.Subjects[i] + " !")
		}
		result = append(result, subject)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func GetCertificatesOfCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseId := args[0]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course - " + CourseId + "!")
	}

	var result []Certificate
	var i int

	for i = 0; i < len(course.Certificates); i++ {

		certificate, err := getCertificate(stub, "Certificate-"+course.Certificates[i])
		if err != nil {
			return shim.Error("Certificate does not exist - " + course.Certificates[i] + " !")
		}
		result = append(result, certificate)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func GetStudentsOfCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseId := args[0]

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course - " + CourseId + " !")
	}

	var i int
	var result []Student

	for i = 0; i < len(course.Students); i++ {
		keyStudent := "Student-" + course.Students[i]
		student, err := getStudent(stub, keyStudent)	

		if err != nil {
			return shim.Error("Can not get student - " + course.Students[i] + " !")
		}

		result = append(result, student)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes !")
	}

	return shim.Success(jsonRow)
}

func GetCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CourseId := args[0]

	keyCourse := "Course-" + CourseId
	courseAsBytes, err := stub.GetState(keyCourse)

	if err != nil {
		return shim.Error("Can not get course!")
	}

	if courseAsBytes == nil {
		return shim.Error("This course does not exist - " + CourseId + "!")
	}

	return shim.Success(courseAsBytes)
}

func GetAllCourses(stub shim.ChaincodeStubInterface) sc.Response {

	allCourses, err := getListCourses(stub)

	if err != nil {
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

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListCourses(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Course-"
	endKey := "Course-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getCourse(stub shim.ChaincodeStubInterface, compoundKey string) (Course, error) {

	var course Course

	courseAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  course, errors.New("Can not get course!")
	}

	if courseAsBytes == nil {
		return  course, errors.New("This course does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(courseAsBytes, &course)

	return course, nil
}

func CreateSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	SubjectId := args[0]
	Description := args[1]
	SubjectName := args[2]
	keySubject := "Subject-" + SubjectId
	_, err := getSubject(stub, keySubject)

	if err == nil {
		return shim.Error("This subjectId already exists - " + SubjectId + "!")
	}

	var subject = Subject{SubjectId: SubjectId, Description: Description, SubjectName: SubjectName}

	subjectAsBytes, err := json.Marshal(subject)
	if err != nil {
		return shim.Error("Can not convert data of this subject to bytes!")
	}

	stub.PutState(keySubject, subjectAsBytes)

	return shim.Success(nil)
}

func UpdateSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	SubjectId := args[0]
	Description := args[1]
	SubjectName := args[2]

	keySubject := "Subject-" + SubjectId
	subject, err := getSubject(stub, keySubject)

	if err != nil {
		return shim.Error("This subject does not exist - " + SubjectId + "!")
	}

	subject.Description = Description
	subject.SubjectName = SubjectName

	subjectAsBytes, err := json.Marshal(subject)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keySubject, subjectAsBytes)
	return shim.Success(nil)
}


func GetClassesOfSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	SubjectId := args[0]

	keySubject := "Subject-" + SubjectId
	subject, err := getSubject(stub, keySubject)

	if err != nil {
		return shim.Error("Can not get subject - " + SubjectId + "!")
	}

	var result []Class
	var i int

	for i = 0; i < len(subject.Classes); i++ {

		class, err := getClass(stub, "Class-"+subject.Classes[i])
		if err != nil {
			return shim.Error("Class does not exist - " + subject.Classes[i] + " !")
		}
		result = append(result, class)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func GetSubject(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	SubjectId := args[0]

	keySubject := "Subject-" + SubjectId
	subjectAsBytes, err := stub.GetState(keySubject)

	if err != nil {
		return shim.Error("Can not get subject!")
	}

	if subjectAsBytes == nil {
		return shim.Error("This subject does not exist - " + SubjectId + "!")
	}

	return shim.Success(subjectAsBytes)
}

func GetAllSubjects(stub shim.ChaincodeStubInterface) sc.Response {

	allSubjects, err := getListSubjects(stub)

	if err != nil {
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

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListSubjects(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Subject-"
	endKey := "Subject-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getSubject(stub shim.ChaincodeStubInterface, compoundKey string) (Subject, error) {

	var subject Subject

	subjectAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  subject, errors.New("Can not get subject!")
	}

	if subjectAsBytes == nil {
		return  subject, errors.New("This subject does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(subjectAsBytes, &subject)

	return subject, nil
}

func CreateClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4!")
	}

	SubjectId := args[0]
	keySubject := "Subject-" + SubjectId
	subject, err := getSubject(stub, keySubject)
	if err != nil {
		return shim.Error("This subject does not exist - " + SubjectId + "!")
	} 
	
	ClassId := args[1]
	Description := args[2]
	ClassName := args[3]
	keyClass := "Class-" + ClassId
	_, err = getClass(stub, keyClass)

	if err == nil {
		return shim.Error("This classId already exists - " + ClassId + "!")
	}

	subject.Classes = append(subject.Classes, ClassId)
	subjectAsBytes, err := json.Marshal(subject)
	if err != nil {
		return shim.Error("Can not convert data of subject - " + SubjectId + " to bytes!")
	}

	var class = Class{ClassId: ClassId, Description: Description, ClassName: ClassName, SubjectId: SubjectId}

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data of this class to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)
	stub.PutState(keySubject, subjectAsBytes)

	return shim.Success(nil)
}

func UpdateClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	ClassId := args[0]
	Description := args[1]
	ClassName := args[2]

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)

	if err != nil {
		return shim.Error("This class does not exist - " + ClassId + "!")
	}

	class.Description = Description
	class.ClassName = ClassName

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)
	return shim.Success(nil)
}


func GetScoresOfClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClassId := args[0]

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)

	if err != nil {
		return shim.Error("Can not get class - " + ClassId + "!")
	}

	var result []Score
	var i int

	for i = 0; i < len(class.Scores); i++ {

		score, err := getScore(stub, "Score-"+class.Scores[i])
		if err != nil {
			return shim.Error("Score does not exist - " + class.Scores[i] + " !")
		}
		result = append(result, score)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func GetTeacherOfClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClassId := args[0]

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)

	if err != nil {
		return shim.Error("Can not get class - " + ClassId + "!")
	}

	if class.TeacherId == "" {
		return shim.Error("This class - " + ClassId + " hasn't got teacher yet!")
	}
 
	keyTeacher := "Teacher-" + class.TeacherId
	teacher, err := getTeacher(stub, keyTeacher)

	if err != nil {
		return shim.Error("Can not get teacher - " + class.TeacherId + "!")
	}

	teacherAsBytes, err := json.Marshal(teacher)
	if err != nil {
		return shim.Error("Can not convet data of teacher - " + class.TeacherId + " to bytes!")
	}

	return shim.Success(teacherAsBytes)
}

func GetSubjectOfClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClassId := args[0]

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)

	if err != nil {
		return shim.Error("Can not get class - " + ClassId + "!")
	}

	keySubject := "Subject-" + class.SubjectId
	subject, err := getSubject(stub, keySubject)

	if err != nil {
		return shim.Error("Can not get subject - " + class.SubjectId + "!")
	}

	subjectAsBytes, err := json.Marshal(subject)
	if err != nil {
		return shim.Error("Can not convet data of subject - " + class.SubjectId + " to bytes!")
	}

	return shim.Success(subjectAsBytes)
}

func GetClass(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClassId := args[0]

	keyClass := "Class-" + ClassId
	classAsBytes, err := stub.GetState(keyClass)

	if err != nil {
		return shim.Error("Can not get class!")
	}

	if classAsBytes == nil {
		return shim.Error("This class does not exist - " + ClassId + "!")
	}

	return shim.Success(classAsBytes)
}

func GetAllClasses(stub shim.ChaincodeStubInterface) sc.Response {

	allClasses, err := getListClasses(stub)

	if err != nil {
		return shim.Error("Can not get all classes!")
	}

	defer allClasses.Close()

	var tlist []Class
	var i int

	for i = 0; allClasses.HasNext(); i++ {

		record, err := allClasses.Next()

		if err != nil {
			return shim.Success(nil)
		}

		class := Class{}
		json.Unmarshal(record.Value, &class)
		tlist = append(tlist, class)
	}

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListClasses(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Class-"
	endKey := "Class-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getClass(stub shim.ChaincodeStubInterface, compoundKey string) (Class, error) {

	var class Class

	classAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  class, errors.New("Can not get class!")
	}

	if classAsBytes == nil {
		return  class, errors.New("This class does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(classAsBytes, &class)

	return class, nil
}

func CreateTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	TeacherId := args[0]
	Description := args[1]
	TeacherName := args[2]
	keyTeacher := "Teacher-" + TeacherId
	_, err := getTeacher(stub, keyTeacher)

	if err == nil {
		return shim.Error("This teacherId already exists - " + TeacherId + "!")
	}

	var teacher = Teacher{TeacherId: TeacherId, Description: Description, TeacherName: TeacherName}

	teacherAsBytes, err := json.Marshal(teacher)
	if err != nil {
		return shim.Error("Can not convert data of this teacher to bytes!")
	}

	stub.PutState(keyTeacher, teacherAsBytes)

	return shim.Success(nil)
}

func UpdateTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	TeacherId := args[0]
	Description := args[1]
	TeacherName := args[2]

	keyTeacher := "Teacher-" + TeacherId
	teacher, err := getTeacher(stub, keyTeacher)

	if err != nil {
		return shim.Error("This teacher does not exist - " + TeacherId + "!")
	}

	teacher.Description = Description
	teacher.TeacherName = TeacherName

	teacherAsBytes, err := json.Marshal(teacher)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyTeacher, teacherAsBytes)
	return shim.Success(nil)
}


func GetClassesOfTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	TeacherId := args[0]

	keyTeacher := "Teacher-" + TeacherId
	teacher, err := getTeacher(stub, keyTeacher)

	if err != nil {
		return shim.Error("Can not get teacher - " + TeacherId + "!")
	}

	var result []Class
	var i int

	for i = 0; i < len(teacher.Classes); i++ {

		class, err := getClass(stub, "Class-"+teacher.Classes[i])
		if err != nil {
			return shim.Error("Class does not exist - " + teacher.Classes[i] + " !")
		}
		result = append(result, class)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func AssignClassToTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	TeacherId := args[0]
	ClassId := args[1]

	keyTeacher := "Teacher-" + TeacherId
	teacher, err := getTeacher(stub, keyTeacher)
	if err != nil {
		return shim.Error("Can not get teacher - " + TeacherId + "!")
	}

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)
	if err != nil {
		return shim.Error("Can not get class - " + ClassId + "!")
	}

	if class.TeacherId != "" {
		return shim.Error("This class was owned by teacher - " + class.TeacherId + " !")
	}

	var i int 
	var checkExist = false
	for i = 0; i < len(teacher.Classes); i++ {
		if teacher.Classes[i] == ClassId {
			checkExist = true
			break
		}
	}

	if checkExist == true {
		return shim.Error("This class already belong to teacher!")
	}

	class.TeacherId = TeacherId
	teacher.Classes = append(teacher.Classes, ClassId)

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data of class - " + ClassId + "to bytes!")
	}

	teacherAsBytes, err := json.Marshal(teacher)
	if err != nil {
		return shim.Error("Can not convert data of teacher - " + TeacherId + " to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)
	stub.PutState(keyTeacher, teacherAsBytes)

	return shim.Success(nil)
}

func UnassignClassFromTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	TeacherId := args[0]
	ClassId := args[1]

	keyTeacher := "Teacher-" + TeacherId
	teacher, err := getTeacher(stub, keyTeacher)
	if err != nil {
		return shim.Error("Can not get teacher - " + TeacherId + "!")
	}

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)
	if err != nil {
		return shim.Error("Can not get class - " + ClassId + "!")
	}

	if class.TeacherId != TeacherId {
		return shim.Error("This class - " + ClassId + " is not belong to teacher - " + TeacherId + " !")
	}


	var i int 
	var leng = len(teacher.Classes)
	for i = 0; i < leng; i++ {
		if teacher.Classes[i] == ClassId {
			break
		}
	}

	copy(teacher.Classes[i:], teacher.Classes[i+1:])
	teacher.Classes[leng-1] = ""
	teacher.Classes = teacher.Classes[:leng-1]

	class.TeacherId = ""

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data of class - " + ClassId + "to bytes!")
	}

	teacherAsBytes, err := json.Marshal(teacher)
	if err != nil {
		return shim.Error("Can not convert data of teacher - " + TeacherId + " to bytes!")
	}

	stub.PutState(keyClass, classAsBytes)
	stub.PutState(keyTeacher, teacherAsBytes)

	return shim.Success(nil)
}

func ChangeTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ClassId := args[0]
	newTeacherId := args[1]

	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)
	if err != nil {
		return shim.Error("Can not get class - " + ClassId  + "!")
	}

	oldTeacherId := class.TeacherId
	keyOldTeacher := "Teacher-" + oldTeacherId
	oldTeacher, err := getTeacher(stub, keyOldTeacher)
	if err != nil {
		return shim.Error("Can not get old teacher - " + oldTeacherId  + "!")
	}

	var i int
	var leng = len(oldTeacher.Classes)
	for i = 0; i < leng; i++ {
		if oldTeacher.Classes[i] == ClassId {
			break
		}
	}

	copy(oldTeacher.Classes[i:], oldTeacher.Classes[i+1:])
	oldTeacher.Classes[leng-1] = ""
	oldTeacher.Classes = oldTeacher.Classes[:leng-1]

	keyNewTeacher := "Teacher-" + newTeacherId
	newTeacher, err := getTeacher(stub, keyNewTeacher)
	if err != nil {
		return shim.Error("Can not get new teacher - " + newTeacherId  + "!")
	}

	class.TeacherId = newTeacherId
	newTeacher.Classes = append(newTeacher.Classes, ClassId)

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data of class - " + ClassId + "to bytes!")
	}

	oldTeacherAsBytes, err := json.Marshal(oldTeacher)
	if err != nil {
		return shim.Error("Can not convert data of teacher - " + oldTeacherId + " to bytes!")
	}


	newTeacherAsBytes, err := json.Marshal(newTeacher)
	if err != nil {
		return shim.Error("Can not convert data of teacher - " + newTeacherId + " to bytes!")
	}


	stub.PutState(keyClass, classAsBytes)
	stub.PutState(keyOldTeacher, oldTeacherAsBytes)
	stub.PutState(keyNewTeacher, newTeacherAsBytes)

	return shim.Success(nil)
}

func GetTeacher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	TeacherId := args[0]

	keyTeacher := "Teacher-" + TeacherId
	teacherAsBytes, err := stub.GetState(keyTeacher)

	if err != nil {
		return shim.Error("Can not get teacher!")
	}

	if teacherAsBytes == nil {
		return shim.Error("This teacher does not exist - " + TeacherId + "!")
	}

	return shim.Success(teacherAsBytes)
}

func GetAllTeachers(stub shim.ChaincodeStubInterface) sc.Response {

	allTeachers, err := getListTeachers(stub)

	if err != nil {
		return shim.Error("Can not get all teachers!")
	}

	defer allTeachers.Close()

	var tlist []Teacher
	var i int

	for i = 0; allTeachers.HasNext(); i++ {

		record, err := allTeachers.Next()

		if err != nil {
			return shim.Success(nil)
		}

		teacher := Teacher{}
		json.Unmarshal(record.Value, &teacher)
		tlist = append(tlist, teacher)
	}

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListTeachers(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Teacher-"
	endKey := "Teacher-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getTeacher(stub shim.ChaincodeStubInterface, compoundKey string) (Teacher, error) {

	var teacher Teacher

	teacherAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  teacher, errors.New("Can not get teacher!")
	}

	if teacherAsBytes == nil {
		return  teacher, errors.New("This teacher does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(teacherAsBytes, &teacher)

	return teacher, nil
}

func CreateStudent(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	StudentId := args[0]
	Description := args[1]
	StudentName := args[2]
	keyStudent := "Student-" + StudentId
	_, err := getStudent(stub, keyStudent)

	if err == nil {
		return shim.Error("This studentId already exists - " + StudentId + "!")
	}

	var student = Student{StudentId: StudentId, Description: Description, StudentName: StudentName}

	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convert data of this student to bytes!")
	}

	stub.PutState(keyStudent, studentAsBytes)

	return shim.Success(nil)
}

func UpdateStudent(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	StudentId := args[0]
	Description := args[1]
	StudentName := args[2]

	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("This student does not exist - " + StudentId + "!")
	}

	student.Description = Description
	student.StudentName = StudentName

	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyStudent, studentAsBytes)
	return shim.Success(nil)
}


func GetCertificatesOfStudent(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	StudentId := args[0]

	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("Can not get student - " + StudentId + "!")
	}

	var result []Certificate
	var i int

	for i = 0; i < len(student.Certificates); i++ {

		certificate, err := getCertificate(stub, "Certificate-"+student.Certificates[i])
		if err != nil {
			return shim.Error("Certificate does not exist - " + student.Certificates[i] + " !")
		}
		result = append(result, certificate)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func StudentResgisterCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	StudentId := args[0]
	CourseId := args[1]

	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("Can not get student - " + StudentId + " !")
	}

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course - " + CourseId + " !")
	}

	var i int 
	for i = 0; i < len(student.Courses); i++ {
		if student.Courses[i] == CourseId {
			return shim.Error("This student - " +  StudentId + " already StudentResgisterCourse the course - " + CourseId + " !")
		}
	}

	student.Courses = append(student.Courses, CourseId)
	course.Students = append(course.Students, StudentId)

	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convet data of student - " + StudentId + " to bytes!")
	}

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convet data of course - " + CourseId + " to bytes!")
	}

	stub.PutState(keyStudent, studentAsBytes)
	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func StudentCancelCourse(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	StudentId := args[0]
	CourseId := args[1]

	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("Can not get student - " + StudentId + " !")
	}

	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course - " + CourseId + " !")
	}

	var i int
	var checkExist = false
	var leng = len(student.Courses) 
	for i = 0; i < leng; i++ {
		if student.Courses[i] == CourseId {
			checkExist = true
			break
		}
	}

	if checkExist == false {
		return shim.Error("This student - " +  StudentId + " does not exist in the course - " + CourseId + " !")
	}

	copy(student.Courses[i:], student.Courses[i+1:])
	student.Courses[leng-1] = ""
	student.Courses = student.Courses[:leng-1]

	leng = len(course.Students)
	for i = 0; i < leng; i++ {
		if course.Students[i] == StudentId {
			break
		}
	}

	copy(course.Students[i:], course.Students[i+1:])
	course.Students[leng-1] = ""
	course.Students = course.Students[:leng-1]
	
	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convet data of student - " + StudentId + " to bytes!")
	}

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convet data of course - " + CourseId + " to bytes!")
	}

	stub.PutState(keyStudent, studentAsBytes)
	stub.PutState(keyCourse, courseAsBytes)

	return shim.Success(nil)
}

func GetCoursesOfStudent(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	StudentId := args[0]

	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("Can not get student - " + StudentId + " !")
	}

	var i int
	var result []Course

	for i = 0; i < len(student.Courses); i++ {
		keyCourse := "Course-" + student.Courses[i]
		course, err := getCourse(stub, keyCourse)	

		if err != nil {
			return shim.Error("Can not get course - " + student.Courses[i] + " !")
		}

		result = append(result, course)
	}

	jsonRow, err := json.Marshal(result)
	if err != nil {
		return shim.Error("Can not convert data to bytes !")
	}

	return shim.Success(jsonRow)
}


func GetStudent(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	StudentId := args[0]

	keyStudent := "Student-" + StudentId
	studentAsBytes, err := stub.GetState(keyStudent)

	if err != nil {
		return shim.Error("Can not get student!")
	}

	if studentAsBytes == nil {
		return shim.Error("This student does not exist - " + StudentId + "!")
	}

	return shim.Success(studentAsBytes)
}

func GetAllStudents(stub shim.ChaincodeStubInterface) sc.Response {

	allStudents, err := getListStudents(stub)

	if err != nil {
		return shim.Error("Can not get all students!")
	}

	defer allStudents.Close()

	var tlist []Student
	var i int

	for i = 0; allStudents.HasNext(); i++ {

		record, err := allStudents.Next()

		if err != nil {
			return shim.Success(nil)
		}

		student := Student{}
		json.Unmarshal(record.Value, &student)
		tlist = append(tlist, student)
	}

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListStudents(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Student-"
	endKey := "Student-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getStudent(stub shim.ChaincodeStubInterface, compoundKey string) (Student, error) {

	var student Student

	studentAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  student, errors.New("Can not get student!")
	}

	if studentAsBytes == nil {
		return  student, errors.New("This student does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(studentAsBytes, &student)

	return student, nil
}

func CreateScore(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	ClassId := args[0]
	keyClass := "Class-" + ClassId
	class, err := getClass(stub, keyClass)
	if err != nil {
		return shim.Error("This class does not exist - " + ClassId + "!")
	} 
	
	ScoreId := args[1]
	Value := args[2]
	keyScore := "Score-" + ScoreId
	_, err = getScore(stub, keyScore)

	if err == nil {
		return shim.Error("This scoreId already exists - " + ScoreId + "!")
	}

	class.Scores = append(class.Scores, ScoreId)
	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convert data of class - " + ClassId + " to bytes!")
	}

	var score = Score{ScoreId: ScoreId, Value: Value, ClassId: ClassId}

	scoreAsBytes, err := json.Marshal(score)
	if err != nil {
		return shim.Error("Can not convert data of this score to bytes!")
	}

	stub.PutState(keyScore, scoreAsBytes)
	stub.PutState(keyClass, classAsBytes)

	return shim.Success(nil)
}

func UpdateScore(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	ScoreId := args[0]
	Value := args[1]

	keyScore := "Score-" + ScoreId
	score, err := getScore(stub, keyScore)

	if err != nil {
		return shim.Error("This score does not exist - " + ScoreId + "!")
	}

	score.Value = Value

	scoreAsBytes, err := json.Marshal(score)
	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	stub.PutState(keyScore, scoreAsBytes)
	return shim.Success(nil)
}


func GetClassOfScore(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ScoreId := args[0]

	keyScore := "Score-" + ScoreId
	score, err := getScore(stub, keyScore)

	if err != nil {
		return shim.Error("Can not get score - " + ScoreId + "!")
	}

	keyClass := "Class-" + score.ClassId
	class, err := getClass(stub, keyClass)

	if err != nil {
		return shim.Error("Can not get class - " + score.ClassId + "!")
	}

	classAsBytes, err := json.Marshal(class)
	if err != nil {
		return shim.Error("Can not convet data of class - " + score.ClassId + " to bytes!")
	}

	return shim.Success(classAsBytes)
}

func GetScore(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ScoreId := args[0]

	keyScore := "Score-" + ScoreId
	scoreAsBytes, err := stub.GetState(keyScore)

	if err != nil {
		return shim.Error("Can not get score!")
	}

	if scoreAsBytes == nil {
		return shim.Error("This score does not exist - " + ScoreId + "!")
	}

	return shim.Success(scoreAsBytes)
}

func GetAllScores(stub shim.ChaincodeStubInterface) sc.Response {

	allScores, err := getListScores(stub)

	if err != nil {
		return shim.Error("Can not get all scores!")
	}

	defer allScores.Close()

	var tlist []Score
	var i int

	for i = 0; allScores.HasNext(); i++ {

		record, err := allScores.Next()

		if err != nil {
			return shim.Success(nil)
		}

		score := Score{}
		json.Unmarshal(record.Value, &score)
		tlist = append(tlist, score)
	}

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListScores(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Score-"
	endKey := "Score-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getScore(stub shim.ChaincodeStubInterface, compoundKey string) (Score, error) {

	var score Score

	scoreAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  score, errors.New("Can not get score!")
	}

	if scoreAsBytes == nil {
		return  score, errors.New("This score does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(scoreAsBytes, &score)

	return score, nil
}

func CreateCertificate(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4!")
	}

	CourseId := args[0]
	keyCourse := "Course-" + CourseId
	course, err := getCourse(stub, keyCourse)
	if err != nil {
		return shim.Error("This course does not exist - " + CourseId + "!")
	} 
	StudentId := args[1]
	keyStudent := "Student-" + StudentId
	student, err := getStudent(stub, keyStudent)
	if err != nil {
		return shim.Error("This student does not exist - " + StudentId + "!")
	} 
	
	CertificateId := args[2]
	Type := args[3]
	keyCertificate := "Certificate-" + CertificateId
	_, err = getCertificate(stub, keyCertificate)

	if err == nil {
		return shim.Error("This certificateId already exists - " + CertificateId + "!")
	}

	course.Certificates = append(course.Certificates, CertificateId)
	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convert data of course - " + CourseId + " to bytes!")
	}
	student.Certificates = append(student.Certificates, CertificateId)
	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convert data of student - " + StudentId + " to bytes!")
	}

	var certificate = Certificate{CertificateId: CertificateId, Type: Type, CourseId: CourseId, StudentId: StudentId}

	certificateAsBytes, err := json.Marshal(certificate)
	if err != nil {
		return shim.Error("Can not convert data of this certificate to bytes!")
	}

	stub.PutState(keyCertificate, certificateAsBytes)
	stub.PutState(keyCourse, courseAsBytes)
	stub.PutState(keyStudent, studentAsBytes)

	return shim.Success(nil)
}



func GetCourseOfCertificate(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CertificateId := args[0]

	keyCertificate := "Certificate-" + CertificateId
	certificate, err := getCertificate(stub, keyCertificate)

	if err != nil {
		return shim.Error("Can not get certificate - " + CertificateId + "!")
	}

	keyCourse := "Course-" + certificate.CourseId
	course, err := getCourse(stub, keyCourse)

	if err != nil {
		return shim.Error("Can not get course - " + certificate.CourseId + "!")
	}

	courseAsBytes, err := json.Marshal(course)
	if err != nil {
		return shim.Error("Can not convet data of course - " + certificate.CourseId + " to bytes!")
	}

	return shim.Success(courseAsBytes)
}

func GetStudentOfCertificate(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CertificateId := args[0]

	keyCertificate := "Certificate-" + CertificateId
	certificate, err := getCertificate(stub, keyCertificate)

	if err != nil {
		return shim.Error("Can not get certificate - " + CertificateId + "!")
	}

	keyStudent := "Student-" + certificate.StudentId
	student, err := getStudent(stub, keyStudent)

	if err != nil {
		return shim.Error("Can not get student - " + certificate.StudentId + "!")
	}

	studentAsBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Can not convet data of student - " + certificate.StudentId + " to bytes!")
	}

	return shim.Success(studentAsBytes)
}

func GetCertificate(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	CertificateId := args[0]

	keyCertificate := "Certificate-" + CertificateId
	certificateAsBytes, err := stub.GetState(keyCertificate)

	if err != nil {
		return shim.Error("Can not get certificate!")
	}

	if certificateAsBytes == nil {
		return shim.Error("This certificate does not exist - " + CertificateId + "!")
	}

	return shim.Success(certificateAsBytes)
}

func GetAllCertificates(stub shim.ChaincodeStubInterface) sc.Response {

	allCertificates, err := getListCertificates(stub)

	if err != nil {
		return shim.Error("Can not get all certificates!")
	}

	defer allCertificates.Close()

	var tlist []Certificate
	var i int

	for i = 0; allCertificates.HasNext(); i++ {

		record, err := allCertificates.Next()

		if err != nil {
			return shim.Success(nil)
		}

		certificate := Certificate{}
		json.Unmarshal(record.Value, &certificate)
		tlist = append(tlist, certificate)
	}

	jsonRow, err := json.Marshal(tlist)

	if err != nil {
		return shim.Error("Can not convert data to bytes!")
	}

	return shim.Success(jsonRow)
}

func getListCertificates(stub shim.ChaincodeStubInterface) (shim.StateQueryIteratorInterface, error) {

	startKey := "Certificate-"
	endKey := "Certificate-zzzzzzzz"

	resultIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}

	return resultIter, nil
}

func getCertificate(stub shim.ChaincodeStubInterface, compoundKey string) (Certificate, error) {

	var certificate Certificate

	certificateAsBytes, err := stub.GetState(compoundKey)

	if err != nil {
		return  certificate, errors.New("Can not get certificate!")
	}

	if certificateAsBytes == nil {
		return  certificate, errors.New("This certificate does not exist - " + compoundKey + "!")
	}

	json.Unmarshal(certificateAsBytes, &certificate)

	return certificate, nil
}

func main() {

	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error createing new Smart Contract: %s", err)
	}
}