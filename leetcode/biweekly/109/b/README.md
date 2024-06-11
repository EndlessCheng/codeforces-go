请看 [视频讲解](https://www.bilibili.com/video/BV1AM4y1x7r4/) 第二题。

普通写法：

```py
class Solution:
    def sortVowels(self, s: str) -> str:
        a = sorted(c for c in s if c in "aeiouAEIOU")
        t = list(s)
        j = 0
        for i, c in enumerate(t):
            if c in "aeiouAEIOU":
                t[i] = a[j]
                j += 1
        return ''.join(t)
```

位运算写法：

```go
func sortVowels(s string) string {
	a := []byte{}
	for _, c := range s {
		if 2130466>>(c&31)&1 > 0 {
			a = append(a, byte(c))
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	t, j := []byte(s), 0
	for i, c := range t {
		if 2130466>>(c&31)&1 > 0 {
			t[i] = a[j]
			j++
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
