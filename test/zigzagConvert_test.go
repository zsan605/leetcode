package test

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	var s = "PAYPALISHIRING"
	var numRows = 4
	//
	//var s = "ab"
	//var numRows = 1
	fmt.Println(convert(s, numRows))
}

func TestCreate(t *testing.T) {
	data :=  map[string]string{
		"eeadfas":"3333",
		"aaaa":"111",
		"bbb":"222",
	}
	Create(data)
}