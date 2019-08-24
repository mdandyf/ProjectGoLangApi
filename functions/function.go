package functions

import (
	"fmt"
	"strconv"
)

var myInt8 uint8
var myBigInt uint64
var myUint8String string
var myUint64String string
var days []string
var myPerson Person

type Person struct {
	name    string
	age     int
	address string
}

func GetPrint(input string) {
	fmt.Println(input)
}

func GetHello() string {
	return "Hello World"
}

func GetUint8() string {
	myInt8 := 123
	myUint8String := "Hello Uint8 :"
	return (myUint8String + strconv.Itoa(myInt8))
}

func GetUint64() string {
	myBigInt := 12300
	myUint64String := "Hello Uint64 :"
	return (myUint64String + strconv.Itoa(myBigInt))
}

func GetArrays() [7]string {
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	return days
}

func GetMap() map[int]string {
	testDay := map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday",
	}
	return testDay
}

func GetStruct() Person {
	gina := Person{name: "Gina Dean", age: 25, address: "Jakarta - Indonesia"}
	return gina
}
