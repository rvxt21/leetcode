package solutions

type Leetcode struct {
}

func New() Leetcode {
	return Leetcode{}
}

func (l Leetcode) TwoSum(nums []int, target int) []int {
	//Difficulty: easy, Title: 1. Two Sum
	indices := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := indices[complement]; ok {
			return []int{j, i}
		}
		indices[num] = i
	}
	return []int{}
}
