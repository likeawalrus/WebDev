package main

import "fmt"

func main() {


	test := func(numb int) (int, bool) {
		var isE bool
		var numbee int
		numbee = numb/2
		if numb%2 == 0 {
			isE = true
		} else {
			isE = false
		}
		return numbee, isE
	}



	var iinput int
	fmt.Scan(&iinput)
	var numbe, pants= test(iinput)
	fmt.Println("half(", iinput, ") returns (", numbe, ",", pants,")")
}

func isEven(numb int) (int, bool) {
	var isE bool
	var numbee int
	numbee = numb/2
	if numb%2 == 0 {
		isE = true
	} else {
		isE = false
	}
	return numbee, isE
}
