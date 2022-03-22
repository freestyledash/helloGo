package practice

import "fmt"

var a = 10
var b = 10

func TestPrint() {
	fmt.Println("123123")
}

func ModifyA() {
	a = 20
}

func PrintA() {
	fmt.Println(a)
}

func ModifyVal(i int) {
	i++
}

func ModifyValWithPointer(i *int) {
	*i = *i + 1
}

type Student struct {
	Name   string
	Age    int
	School School
}

type School struct {
	Name string
}

func ModifyStudent(student *Student) {
	student.Name = "张言琦"
	student.Age = 10
	student.School.Name = "西南大学"
}
