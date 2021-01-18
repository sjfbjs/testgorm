package main

import "fmt"

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}
func (t *T) M2() {
	t.Name = "name2"
}

type Member struct {
	id     int
	name   string
	email  string
	gender int
	age    int
}

func Change(m1 Member, m2 *Member) {
	m1.name = "小明"
	m2.name = "小红"
}

func test(test map[int32]string) {
	// 添加
	test[5] = "e"
	// 修改
	test[1] = "aaa"
	// 删除
	delete(test, 2)
}

func main() {
	fmt.Println("修改t1===")
	t1 := T{"t1"}
	fmt.Printf("M1调用前地址：%p", &t1)
	//
	fmt.Println("M1调用前：", t1.Name)
	fmt.Printf("M1调用前地址：%p", &t1)

	t1.M1()
	fmt.Println("M1调用后：", t1.Name)
	fmt.Printf("M1调用后地址：%p", &t1)

	//
	fmt.Println("M2调用前：", t1.Name)
	fmt.Printf("M2调用前地址：%p", &t1)
	t1.M2()
	fmt.Println("M2调用后：", t1.Name)
	fmt.Printf("M2调用后地址：%p", &t1)
	fmt.Println("修改t2===")
	t2 := &T{"t2"}
	fmt.Println("M1调用前：", t2.Name)
	t2.M1()
	fmt.Println("M1调用后：", t2.Name)
	fmt.Println("M2调用前：", t2.Name)
	t2.M2()
	fmt.Println("M2调用后：", t2.Name)

	//
	fmt.Println("----------------------")
	//空值类型
	m1 := Member{}
	//指针类型
	m2 := new(Member)
	//函数改变值--------------   func(struct)
	/*
		默认的函数传参传结构体进去，结构体本身不会进行任何改变
	    因为函数传的参数会拷贝一份
	 */
	Change(m1, m2)
	fmt.Println(m1)
	fmt.Println(m2)

	//a :=
	//fmt.Println(a)
	fmt.Println("----------------------")
	testMap := map[int32] string{}
	testMap[1] = "a"
	testMap[2] = "b"
	testMap[3] = "c"
	fmt.Println("原始值:", testMap)
	test(testMap)
	fmt.Println("处理后的值", testMap)
	type Intf interface {
		M1()
		M2()
	}
	//var t3 T = T{"t1"}
	var t3  = T{"t3"}
	// t3的指针实现了M2方法
	fmt.Println("M1调用前：", t3.Name)
	t3.M1()
	fmt.Println("M1调用后：", t3.Name)

	fmt.Println("M2调用前：", t3.Name)
	t3.M2()
	fmt.Println("M2调用后：", t3.Name)

	// var  t4 Intf = t3  错误写法
	/*

	**/
	var it Intf = &t3
	it.M1()
	it.M2()
	//fmt.Println("M2调用后：", it.Name)

}
