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

func main(){
	c := Class{
		Banji:"一班",
		Students: make([]Student,0),
	}
	for i :=1;i<=10 ;i++ {
		//指针写法
		//s :=&Student{}
		//s.Name = fmt.Sprintf("mack%v",i)
		//s.Age = i
		//c.Students = append(c.Students,*s)
		//实例化写法
		s := Student{
			Name: fmt.Sprintf("mack%v",i),
			Age:i,
		}
		c.Students = append(c.Students,s)
	}
	//fmt.Errorf(c)
	strByte, error := json.Marshal(c)
	if error != nil {
		fmt.Print(error)
	} else  {
		strJson := string(strByte)
		fmt.Println(strJson)
	}


}