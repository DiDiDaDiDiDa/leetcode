package main

import (
	"fmt"
	"strings"
)

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := input()
	result := mySolution(s)
	fmt.Println("result:", result)
}

//我的解决方案
func mySolution(s string) bool {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		if s == "" {
			return true
		}
	}
	return false
}

//牛逼的解决办法
//牛逼之处：不用像上面一定要遍历这么多次
func niuBiSolution(s string) bool {
	for {
		old := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
	}
	return false
}

func input() string {

	var s string
	fmt.Printf("请输入字符串: ")
	fmt.Scanln(&s)
	return s
}
