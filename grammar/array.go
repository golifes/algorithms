package main

import "fmt"

func main() {
	//arr1 := [3]int{1, 2, 3}
	arr2 := []int{1, 2, 3, 1, 4, 0, 1}
	sum := twoSum(arr2, 4)
	fmt.Println(sum)
	//arr3 := arr2[1:]
	//
	//arr4 := arr1[1:]
	//d := arr1[:]
	//
	//fmt.Println(arr2 == arr1)
	//
	//fmt.Println(d, arr3[:])
	//
	//fmt.Println(arr4)
	//arr5 := [][]int{{1, 2}, {3, 4}}
	//fmt.Println(arr5[len(arr5)-1][1])

}

func twoSum(arr []int, sum int) (ret [][]int) {
	arrMap := map[int]int{}
	for k, v := range arr {
		tmp := sum - v
		if _k, ok := arrMap[tmp]; ok {
			ret = append(ret, []int{k, _k})
			//if _k != k{
			//	ret = append(ret, []int{k, _k})
			//}else {
			//	ret = append(ret, []int{k, _k})
			//}
		} else {
			arrMap[v] = k
		}
	}

	return
}

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []byte

	for i := 0; i < length; i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if stack[0] != pairs[s[i]] {
				return false
			}
			stack = stack[1:]
		}
		//if pairs[s[i]] > 0 {
		//	if len(stack) == 0 || stack[0] != pairs[s[i]] {
		//		return false
		//	}
		//	//删除栈顶第一个女元素
		//	stack = stack[1:]
		//} else {
		//	stack = append(stack, s[i])
		//}
	}

	return len(stack) == 0

}
