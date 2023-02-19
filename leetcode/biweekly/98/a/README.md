下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

把 $\textit{num}$ 转成字符串，枚举每个字符，替换成 $9$ 得到最大数，替换成 $0$ 得到最小数。

```py [sol1-Python3]
class Solution:
    def minMaxDifference(self, num: int) -> int:
        mx, mn = 0, num
        s = str(num)
        for c in s:
            mx = max(mx, int(s.replace(c, '9')))
            mn = min(mn, int(s.replace(c, '0')))
        return mx - mn
```

```go [sol1-Go]
func minMaxDifference(num int) int {
	mx, mn := 0, num
	s := strconv.Itoa(num)
	for _, c := range s {
		t := strings.ReplaceAll(s, string(c), "9")
		x, _ := strconv.Atoi(t)
		mx = max(mx, x)
		t = strings.ReplaceAll(s, string(c), "0")
		x, _ = strconv.Atoi(t)
		mn = min(mn, x)
	}
	return mx - mn
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(\log^2 \textit{num})$。
- 空间复杂度：$O(\log \textit{num})$。
