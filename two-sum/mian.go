package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2,7,11,15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	//获取控制台输入
	nums, target := input()

	result := niuBiSolution(nums, target)

	fmt.Println("result:", result)
}

//我的解决方案
func mySolution(nums []int, target int) (result []int) {
	numLen := len(nums)
OK:
	for i := 0; i < numLen; i++ {
		t := target - nums[i]
		for j := i + 1; j < numLen; j++ {
			if t == nums[j] {
				result[0] = i
				result[1] = j
				break OK
			}
			continue
		}
	}
	return
}

//牛逼的解决办法
//牛逼之处：简化复杂度
func niuBiSolution(nums []int, target int) (result []int) {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}

		// 把b和i的值，存入map
		index[b] = i
	}

	return nil
}

func input() ([]int, int) {
	var (
		numsStr string
		target  int
	)

	fmt.Printf("请输入数组内容(如：1,2,3,4): ")
	fmt.Scanln(&numsStr)
	fmt.Printf("请输入target: ")
	fmt.Scanln(&target)
	return transform(numsStr), target
}

func transform(str string) (nums []int) {
	numsStr := strings.Split(str, ",")
	for _, v := range numsStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return
}
