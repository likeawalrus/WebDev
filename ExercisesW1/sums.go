package main
//Problem 2 on projecteuler
//Add the total of all even fibonacci numbers below 4000000
import (
	"fmt"
)

func main(){
	first := 0
	second := 1
	third := 0
	sumeven := 0
	for third<4000000 {
		third = first+second
		if third%2 == 0 {
			sumeven	= sumeven + third
		}
		first = second
		second = third
	}
	fmt.Println(sumeven)
}