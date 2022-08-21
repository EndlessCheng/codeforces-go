下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

用差分数组 $\textit{diff}$ 表示一段区间上的更新，即在 $\textit{start}_i$ 变化量增加了 $x$，在 $\textit{end}_i+1$ 变化量减少了 $x$（类比导数的概念）。

这里 $x=2\cdot\textit{direction}_i-1$，把 $0$ 和 $1$ 变成 $-1$ 和 $1$。

然后从小到大遍历 $\textit{diff}$，累加变化量为 $\textit{shift}$（类比积分的概念），这样对于第 $i$ 个字符，其移位值就是 $\textit{shift}$。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
c2i = {c: i for i, c in enumerate(ascii_lowercase)}

class Solution:
    def shiftingLetters(self, s: str, shifts: List[List[int]]) -> str:
        diff = [0] * (len(s) + 1)
        for st, end, dir in shifts:
            diff[st] += dir * 2 - 1
            diff[end + 1] -= dir * 2 - 1
        return ''.join(ascii_lowercase[(c2i[c] + shift) % 26] for c, shift in zip(s, accumulate(diff)))
```

```go [sol1-Go]
func shiftingLetters(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)
	for _, p := range shifts {
		x := p[2]*2 - 1 // 0 和 1 变成 -1 和 1
		diff[p[0]] += x
		diff[p[1]+1] -= x
	}
	t, shift := []byte(s), 0
	for i, b := range t {
		shift = (shift+diff[i])%26 + 26 // 防一手负数
		t[i] = (b-'a'+byte(shift))%26 + 'a'
	}
	return string(t)
}
```
