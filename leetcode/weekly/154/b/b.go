package main

import "slices"

// https://space.bilibili.com/206214
func reverseParentheses1(s string) string {
	i := 0

	var f func() []byte
	f = func() (res []byte) {
		for i < len(s) {
			ch := s[i]
			i++
			if ch == ')' { // 归
				slices.Reverse(res)
				return
			}
			if ch == '(' { // 递
				res = append(res, f()...)
			} else { // 字母
				res = append(res, ch)
			}
		}
		return
	}

	return string(f())
}

func reverseParentheses(s string) string {
	n := len(s)
	links := make([]int, n)
	st := []int{}
	for i, ch := range s {
		if ch == '(' {
			st = append(st, i)
		} else if ch == ')' {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			// 右括号与栈顶左括号配对，下标建立双向链接
			links[i] = j
			links[j] = i
		}
	}

	ans := []byte{}
	step := 1
	for i := 0; i < n; i += step {
		ch := s[i]
		if ch == '(' || ch == ')' {
			i = links[i] // 跳到对应的括号位置
			step = -step // 反向移动
		} else {
			ans = append(ans, ch)
		}
	}
	return string(ans)
}
