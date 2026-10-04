package main

// github.com/EndlessCheng/codeforces-go
func minInsertions1(s string) (ans int) {
	c, n := 0, len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			c++
		} else {
			if i+1 == n || s[i+1] != ')' {
				ans++
			} else {
				i++
			}
			if c > 0 {
				c--
			} else {
				ans++
			}
		}
	}
	ans += 2 * c
	return
}

func minInsertions(s string) (ans int) {
	n := len(s)
	left := 0 // 未配对的左括号个数

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++ // 未配对的左括号
			continue
		}

		if left > 0 {
			left-- // 左右括号配对
		} else {
			ans++ // 右括号太多了，补一个左括号
		}

		// 必须有两个连续的右括号，也就是 s[i+1] 必须也是右括号
		if i < n-1 && s[i+1] == ')' {
			i++
		} else {
			ans++ // 补一个右括号
		}
	}

	// 左括号太多了，补上缺失的右括号
	return ans + left*2
}
