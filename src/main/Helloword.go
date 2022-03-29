package main

import "helloGo/src/algorithm"

func main() {
	//fmt.Println("hello!")
	//practice.TestPrint()
	//practice.ModifyA()
	//practice.PrintA()
	//
	//i := 10
	//practice.ModifyVal(i)
	//fmt.Println(i)
	//
	//practice.ModifyValWithPointer(&i)
	//fmt.Println(i)
	//
	//student := practice.Student{"小红", 10, practice.School{"长征小学"}}
	//practice.ModifyStudent(&student)
	//fmt.Println(student.Name)
	//fmt.Println(student.School.Name)
	//
	//student.ModifySelf("KEKE", 25, "西南大学")
	//fmt.Println(student.Name)
	//fmt.Println(student.School.Name)
	//
	//var a int64 = 10
	//var trigger bool = true
	//stringChan := make(chan string)
	//practice.HaveATest(&a, &trigger, stringChan)
	//o := 0
	//for {
	//	// 从通道中取出数据, 此处会阻塞直到信道中返回数据
	//	message := <-stringChan
	//	// 打印数据
	//	fmt.Println(message)
	//	o++
	//	if o >= 3 {
	//		break
	//	}
	//}
	//fmt.Println(a)
	//
	//practice.TestPanic()

	//web.StartServer()
	//fmt.Println("server Start")

	algorithm.HaveAMapTest()
}
