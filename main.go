package main

import (
	"encoding/json"
	"fmt"
)

type TA struct {
	A int64
	B int
}
func main() {
	var tmp = TA{
		A:111111111111111111,
		B:1,
	}
	tmpBytes, _:= json.Marshal(tmp)
	var tmp1 = make(map[string]interface{})
	json.Unmarshal(tmpBytes, &tmp1)

	tmp1Bytes,_:=json.Marshal(tmp1)
	fmt.Println(string(tmp1Bytes))

}
