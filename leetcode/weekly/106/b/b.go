package main

// https://space.bilibili.com/206214
func minAddToMakeValid(s string) (ans int) {
	left := 0 // 未配对的左括号的个数
	for _, ch := range s {
		if ch == '(' {
			left++
		} else if left > 0 {
			left-- // 左右括号配对
		} else { // 右括号太多了
			ans++ // 在 ch 左边任意位置插入一个左括号，与 ch 配对
		}
	}

	// 循环结束后，如果 left > 0，说明有 left 个多余的左括号
	// 在 s 末尾插入 left 个右括号
	return ans + left
}
