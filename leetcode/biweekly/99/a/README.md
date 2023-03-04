下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

两个要点：

1. 数字要均匀分；
2. 把小的数字放高位，大的数字放低位。

```py [sol1-Python3]
class Solution:
    def splitNum(self, num: int) -> int:
        a = [[], []]
        for i, c in enumerate(sorted(list(str(num)))):
            a[i % 2].append(c)
        return int(''.join(a[0])) + int(''.join(a[1]))
```

```go [sol1-Go]
func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	a := [2][]byte{}
	for i, c := range s {
		a[i&1] = append(a[i&1], c)
	}
	x, _ := strconv.Atoi(string(a[0]))
	y, _ := strconv.Atoi(string(a[1]))
	return x + y
}
```

### 复杂度分析

- 时间复杂度：$O(m\log m)$，其中 $m$ 为 $\textit{num}$ 转成字符串后的长度。
- 空间复杂度：$O(m)$。
