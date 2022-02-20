package main

// github.com/EndlessCheng/codeforces-go
func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	ans := []byte{}
	for i := 25; i >= 0; i-- { // 找一个最大的字母
		if cnt[i] == 0 {
			continue
		}
	next:
		for j := 0; j < repeatLimit && cnt[i] > 0; j++ { // 填充 min(repeatLimit, cnt[i]) 个字母 i
			cnt[i]--
			ans = append(ans, 'a'+byte(i))
		}
		if cnt[i] == 0 { // i 用完了，找下一个更小的字母
			continue
		}
		for j := i - 1; j >= 0; j-- { // 插入一个字母 j，这样可以继续填 i
			if cnt[j] > 0 {
				cnt[j]--
				ans = append(ans, 'a'+byte(j))
				goto next
			}
		}
	}
	return string(ans)
}
