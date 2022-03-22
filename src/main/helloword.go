package main

import (
	"fmt"
	"helloGo/src/practice"
)

func main() {
	fmt.Println("hello!")
	practice.TestPrint()
	practice.ModifyA()
	practice.PrintA()

	i := 10
	practice.ModifyVal(i)
	fmt.Println(i)

	practice.ModifyValWithPointer(&i)
	fmt.Println(i)

	student := practice.Student{"小红", 10, practice.School{"长征小学"}}
	practice.ModifyStudent(&student)
	fmt.Println(student.Name)
	fmt.Println(student.School.Name)

	student.ModifySelf("KEKE", 25, "西南大学")
	fmt.Println(student.Name)
	fmt.Println(student.School.Name)
}
