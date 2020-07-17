package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestInstancesCreation(test *testing.T) {
	stub := InitChaincode(test)

	Invoke(test, stub, "CreateSubject", "Subject01", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "GetSubject", "Subject01")
	Invoke(test, stub, "UpdateSubject", "Subject01", "Trinh Van Tan", "Hoang Ngoc Phuc")
	Invoke(test, stub, "GetSubject", "Subject01")
	Invoke(test, stub, "CreateCourse", "Course01", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "GetCourse", "Course01")
	Invoke(test, stub, "UpdateCourse", "Course01", "Trinh Van Tan", "Hoang Ngoc Phuc")
	Invoke(test, stub, "GetCourse", "Course01")
	Invoke(test, stub, "AddSubjectToCourse", "Course01", "Subject01")
	Invoke(test, stub, "GetCourse", "Course01")
	Invoke(test, stub, "GetSubjectsOfCourse", "Course01")
	Invoke(test, stub, "CreateSubject", "Subject02", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "AddSubjectToCourse", "Course01", "Subject02")
	Invoke(test, stub, "RemoveSubjectFromCourse", "Course01", "Subject02")
	Invoke(test, stub, "GetSubjectsOfCourse", "Course01")
	Invoke(test, stub, "CreateClass", "Subject01", "Class01", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "GetAllClasses")
	Invoke(test, stub, "CreateTeacher", "Teacher01", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "CreateTeacher", "Teacher02", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "GetTeacher", "Teacher01")
	Invoke(test, stub, "AssignClassToTeacher", "Teacher01", "Class01")
	Invoke(test, stub, "GetClassesOfTeacher", "Teacher01")
	Invoke(test, stub, "GetAllClasses")
	// Invoke(test, stub, "UnassignClassFromTeacher", "Teacher01", "Class01")
	Invoke(test, stub, "GetClassesOfTeacher", "Teacher01")
	Invoke(test, stub, "GetAllClasses")
	Invoke(test, stub, "CreateClass", "Subject01", "Class02", "Hoang Ngoc Phuc", "Trinh Van Tan")
	Invoke(test, stub, "GetTeacherOfClass", "Class01")
	Invoke(test, stub, "GetSubjectOfClass", "Class02")
	Invoke(test, stub, "GetClassesOfSubject", "Subject01")
	Invoke(test, stub, "ChangeTeacher", "Class01", "Teacher02")
	Invoke(test, stub, "GetTeacherOfClass", "Class01")
	Invoke(test, stub, "GetClassesOfTeacher", "Teacher02")
}

func InitChaincode(test *testing.T) *shim.MockStub {
	stub := shim.NewMockStub("testingStub", new(SmartContract))
	result := stub.MockInit("000", nil)

	if result.Status != shim.OK {
		test.FailNow()
	}
	return stub
}

func Invoke(test *testing.T, stub *shim.MockStub, function string, args ...string) {
	cc_args := make([][]byte, 1+len(args))
	cc_args[0] = []byte(function)

	for i, arg := range args {
		cc_args[i+1] = []byte(arg)
	}
	result := stub.MockInvoke("000", cc_args)
	fmt.Println("Call:	", function, "(", strings.Join(args, ", "), ")")
	fmt.Println("RetCode:	", result.Status)
	fmt.Println("RetMsg:	", result.Message)
	fmt.Println("Payload:	", string(result.Payload))

	if result.Status != shim.OK {
		test.FailNow()
	}
}
