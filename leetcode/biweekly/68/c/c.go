package main

/* 全部变为左括号或右括号

首先 $s$ 长度不能为奇数，此时应直接返回 $\texttt{false}$。

然后就是用括号问题的经典技巧了：通过一个变量记录括号的平衡度。

以变左括号为例，将所有可以变化的括号变为左括号，如果这样中间还是会出现平衡度小于 $0$ 的情况，那么就返回 $\texttt{false}$。多余的平衡度可以通过将左括号变成右括号来实现

*/

// github.com/EndlessCheng/codeforces-go
func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}

	x := 0
	for i, ch := range s {
		if ch == '(' || locked[i] == '0' { // 能变左就变左
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}

	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' { // 能变右就变右
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	return true
}
