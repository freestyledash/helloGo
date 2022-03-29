package practice

import "fmt"

func TestPanic() {
	defer func() {
		fmt.Println("enter defer")
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("recover success")
		}
	}()

	panic("there is a panic")
}
