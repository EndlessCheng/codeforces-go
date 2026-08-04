package main

import "strings"

func minRemoveToMakeValid(s string) string {
	t := []byte(s)
	st := []int{}

	for i, ch := range t {
		if ch == '(' {
			st = append(st, i) // 记录左括号的下标
		} else if ch == ')' {
			if len(st) > 0 {
				st = st[:len(st)-1] // 左右括号匹配
			} else {
				t[i] = '-' // 右括号没有与之匹配的左括号，标记为移除
			}
		}
	}
	for _, i := range st {
		t[i] = '-' // 栈中剩下的左括号没有与之匹配的右括号，标记为移除
	}

	return strings.Replace(string(t), "-", "", -1)
}
