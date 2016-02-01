package main

import (
	"fmt"
)

func main(){
	var ssmall int
	var llarge int
	fmt.Println("Enter a large number")
	fmt.Scan(&llarge)
	fmt.Println("Enter a small number")
	fmt.Scan(&ssmall)
	fmt.Println("remainder")
	fmt.Println(llarge%ssmall)
}