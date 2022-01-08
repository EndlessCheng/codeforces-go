package main

// github.com/EndlessCheng/codeforces-go
func longestPalindrome(words []string) (ans int) {
	cnt := [26][26]int{}
	for _, s := range words {
		cnt[s[0]-'a'][s[1]-'a']++
	}
	odd := 0 // 是否有一个出现奇数次的 AA
	for i := 0; i < 26; i++ {
		c := cnt[i][i] // 相同字符
		ans += c &^ 1  // c &^ 1 等价于 c - c % 2
		odd |= c & 1
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2 // AB BA 取出现次数最小值
		}
	}
	return (ans + odd) * 2 // 上面的循环统计的是字符串个数，最后再乘 2
}

func min(a, b int) int { if a > b { return b }; return a }
