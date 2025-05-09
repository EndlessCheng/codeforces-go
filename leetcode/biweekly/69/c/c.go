package main

// github.com/EndlessCheng/codeforces-go
func longestPalindrome(words []string) (ans int) {
	cnt := [26][26]int{}
	for _, w := range words {
		cnt[w[0]-'a'][w[1]-'a']++
	}

	odd := 0 // 是否存在出现奇数次的 cnt[i][i]
	for i := range cnt {
		c := cnt[i][i]
		ans += c - c%2 // 保证结果是偶数，也可以写成 c &^ 1
		odd |= c % 2   // 存在出现奇数次的 cnt[i][i]
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2
		}
	}
	return (ans + odd) * 2 // 上面统计的是字符串个数，乘以 2 就是长度 
}
