把 $\textit{num}$ 转成字符串 $s$，从左到右找第一个不是 $9$ 的字符，把这个字符都替换成 $9$，得到最大数。

同理找第一个不是 $0$ 的字符，替换成 $0$ 得到最小数，由于 $s[0]$ 一定不是 $0$，所以替换它就行。

附：[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)

```py [sol1-Python3]
class Solution:
    def minMaxDifference(self, num: int) -> int:
        mx = num
        s = str(num)
        for c in s:
            if c != '9':
                mx = int(s.replace(c, '9'))
                break
        return mx - int(s.replace(s[0], '0'))
```

```go [sol1-Go]
func minMaxDifference(num int) int {
	mx := num
	s := strconv.Itoa(num)
	for _, c := range s {
		if c != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "9"))
			break
		}
	}
	mn, _ := strconv.Atoi(strings.ReplaceAll(s, s[:1], "0"))
	return mx - mn
}
```

### 复杂度分析

- 时间复杂度：$O(\log \textit{num})$。
- 空间复杂度：$O(\log \textit{num})$。
