package main

import (
	"encoding/json"
	"fmt"
)

type   Student struct {

	Name string `json:"name"`
	Age int `json:"age"`


}
type  Class struct {

	Banji string `json:"banji"`
	Students []Student

}




func main() {

	str:= `{"banji":"一班","Students":[{"name":"mack1","age":1},{"name":"mack2","age":2},{"name":"mack3","age":3},{"name":"mack4","age":4},{"name":"mack5","age":5},{"name":"mack6","age":6},{"name":"mack7","age":7},
{"name":"mack8","age":8},{"name":"mack9","age":9},{"name":"mack10","age":10}]}`
	c := Class{}
	error := json.Unmarshal([]byte(str),&c)
	if error!=nil {

		fmt.Print(error)
	}else{

		fmt.Println(c)
	}
}