统计 $s$ 中每个字母的个数 $\textit{cnt}$。要使字典序最大，应优先选择最大的字母 $i$，选择 $\min(\textit{cnt}[i], \textit{repeatLimit})$ 个。

如果字母 $i$ 已被选完，那么继续选择下一个较小的字母。

如果字母 $i$ 未被选完，为了不让字母 $i$ 连续出现次数超过 $\textit{repeatLimit}$，我们可以选择一个较小的字母放在后面，然后再继续选择字母 $i$。

```go
func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	ans := []byte{}
next:
	for {
		for i := 25; i >= 0; i-- { // 找一个最大的字母
			if cnt[i] == 0 {
				continue
			}
			for j := 0; j < repeatLimit && cnt[i] > 0; j++ { // 填充 min(repeatLimit, cnt[i]) 个字母 i
				cnt[i]--
				ans = append(ans, 'a'+byte(i))
			}
			if cnt[i] == 0 { // i 用完了，找下一个更小的字母
				continue next
			}
			for j := i - 1; j >= 0; j-- { // 插入一个字母 j，这样可以继续填 i
				if cnt[j] > 0 {
					cnt[j]--
					ans = append(ans, 'a'+byte(j))
					continue next
				}
			}
			break // 后面没有字母了
		}
		return string(ans)
	}
}
```
