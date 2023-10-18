package main

import "fmt"

func main() {
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}

	for _, num := range nums {
		if num%2 == 0 && num != 0 {
			fmt.Println(num, "Even")
			continue
		}
		fmt.Println(num, "Odd")
	}

}
