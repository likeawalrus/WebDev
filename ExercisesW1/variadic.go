package main

import "fmt"

func main() {
	fmt.Println( findGreatest(1,2,3,4,5) )

}

func findGreatest(nums ...int) (int) {
		
	biggest := 0

	for _, num := range nums {
		if num > biggest {
			biggest = num
		}
	}
	return biggest
}
