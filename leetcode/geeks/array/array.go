package main

import (
	"fmt"
	"math"
)

func moveZeroesV1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	l := len(nums)
	tmp := make([]int, l, l)
	idx := 0
	for _, v := range nums {
		if v != 0 {
			tmp[idx] = v
			idx += 1
		}
	}
	return tmp
}

func moveZeroesV2(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	nonZeroCount := 0
	for _, v := range nums {
		if v != 0 {
			nums[nonZeroCount] = v
			nonZeroCount += 1
		}
	}
	for i := nonZeroCount; i < length; i++ {
		nums[i] = 0
	}
}
func maxArea() int {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	length := len(height)
	var max float64
	i, j := 0, length-1
	for i < j {
		area := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
		max = math.Max(max, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return int(max)
}

func plusOne(digits []int) []int {

	length := len(digits)
	if length <= 0 {
		return nil
	}

	tmp := 0

	for i := length - 1; i >= 0; i-- {
		//fmt.Println(digits[i],math.Pow10(length-i-1),i,length-i)
		tmp += digits[i] * (10 ^ int(math.Pow10(length-i-1)))

	}
	tmp += 1
	//tmpArray := make([]int,1)
	for tmp > 0 {
		tmp = tmp / 10
		fmt.Println(tmp)
	}
	return nil
}

func main() {
	area := plusOne([]int{9, 9})
	println(area)
}
