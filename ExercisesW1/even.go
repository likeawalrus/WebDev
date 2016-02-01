package main

import (
	"fmt"
)

func main(){
	var rem int
	for i:= 0
		i<=100
		i++ {
		rem = i%2
		if rem == 0 {
			fmt.Println(i)
		}
	}
}