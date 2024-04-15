package main

import (
	"fmt"
	"reflect"
)

type CustomInt int

var myCustomInt CustomInt
var i int

func main() {

	i = 5

	typeI := reflect.TypeOf(i)
	valueI := reflect.ValueOf(i)

	fmt.Printf("type: %v, value: %v\n", typeI, valueI)

	myCustomInt = 10

	typeOfMyCustomInt := reflect.TypeOf(myCustomInt)
	valueOfMyCustomInt := reflect.ValueOf(myCustomInt)

	originalTypeOfMyCustomInt := valueOfMyCustomInt.Kind() // может различать только basic types
	fmt.Printf("type: %v, value: %v, original type: %v\n", typeOfMyCustomInt, valueOfMyCustomInt, originalTypeOfMyCustomInt)

	val := valueOfMyCustomInt.Interface() // можно получить значение
	fmt.Printf("CustomInt value is: %v\n", val)
}
