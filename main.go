package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	fmt.Println("Solutions for leetcode problems")
	leetcodeSol := solutions.New()
	res := leetcodeSol.TwoSum([]int{2, 7, 11, 12}, 9)
	fmt.Print(res)
}
