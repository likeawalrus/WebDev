package main

import (
	"fmt"
)

func main(){
	var rem5 int
	var rem3 int
	for i:= 1
		i<=100
		i++ {
		rem5 = i%5
		rem3 = i%3
		if rem5 == 0 && rem3 == 0 {
			fmt.Println("fizzbuzz")
		} else if rem5 == 0 {
			fmt.Println("Buzz")
		} else if rem3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}