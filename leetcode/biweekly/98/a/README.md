把 $\textit{num}$ 转成字符串，从左到右找第一个不是 $9$ 的字符，替换成 $9$ 得到最大数；同理找第一个不是 $0$ 的字符，替换成 $0$ 得到最小数。

附：[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)

```py [sol1-Python3]
class Solution:
    def minMaxDifference(self, num: int) -> int:
        mx = mn = num
        s = str(num)
        for c in s:
            if c != '9':
                mx = int(s.replace(c, '9'))
                break
        for c in s:
            if c != '0':
                mn = int(s.replace(c, '0'))
                break
        return mx - mn
```

```go [sol1-Go]
func minMaxDifference(num int) int {
	mx, mn := num, num
	s := strconv.Itoa(num)
	for _, c := range s {
		if c != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "9"))
			break
		}
	}
	for _, c := range s {
		if c != '0' {
			mn, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "0"))
			break
		}
	}
	return mx - mn
}
```

### 复杂度分析

- 时间复杂度：$O(\log \textit{num})$。
- 空间复杂度：$O(\log \textit{num})$。
