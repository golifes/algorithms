package main

import (
	"fmt"
	"sort"
)

func main() {
	//s := "12334"
	//for k, v := range s {
	//	fmt.Println(k, v, string(v))
	//}
	area := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(area)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 2; i < n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3

}

func threeSum(nums []int) [][]int {

	length := len(nums)
	if length <= 2 {
		return nil
	}
	sort.Ints(nums)

	var res [][]int
	for i := 0; i <= length-2; i++ {
		left, right := i+1, length-1
		start := -nums[i]
		if -start > 0 {
			break
		}

		// 去重：如果此数已经选取过，
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}
		for left < right {
			_right := nums[right]
			_left := nums[left]

			if _right+_left == start {
				res = append(res, []int{nums[i], _right, _left})
				left += 1
				right -= 1
				// 去重：第二个数和第三个数也不重复选取
				for (left < right) && (nums[left] == nums[left-1]) {
					left += 1
				}
				// 去重：第二个数和第三个数也不重复选取
				for (left < right) && (nums[right] == nums[right+1]) {
					right -= 1
				}

			} else if (_right + _left) < start {
				left += 1
			} else {
				right -= 1
			}

		}
	}
	return res
}
