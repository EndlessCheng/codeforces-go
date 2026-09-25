package main

// github.com/EndlessCheng/codeforces-go
func evaluate(s string, knowledge [][]string) string {
	mp := make(map[string]string, len(knowledge)) // 预分配空间
	for _, k := range knowledge {
		mp[k[0]] = k[1]
	}

	ans := []byte{}
	left := -1
	for i, ch := range s {
		if ch == '(' {
			left = i // 记录未配对左括号的位置
		} else if ch == ')' {
			// 替换左右括号之间的子串
			t, ok := mp[s[left+1:i]]
			if !ok {
				t = "?"
			}
			ans = append(ans, t...)
			left = -1
		} else if left < 0 { // ch 不在括号中
			ans = append(ans, byte(ch))
		}
	}
	return string(ans)
}
