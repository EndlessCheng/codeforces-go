下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，不仅讲做法，还会教你如何一步步思考，记得关注哦~

---

由于题目保证 $\textit{words}$ 中的字符串互不相同，所以可以遍历 $\textit{words}$，对于 $w=\textit{words}[i]$：

- 如果前面遇到了 $w$ 反转的字符串，那么就找到了一个匹配。
- 如果前面没有遇到，那么把 $w$ 加入一个哈希表（或者数组），方便后面快速判断是否有对应的字符串。

### 思考题

1. 如果 $\textit{words}$ 中有相同字符串要怎么做？
2. 如果 $\textit{words}[i]$ 的长度大于 $2$ 要怎么做？

下午的直播中会讲讲这两个思考题要怎么做。

```py [sol-Python3]
class Solution:
    def maximumNumberOfStringPairs(self, words: List[str]) -> int:
        ans, vis = 0, set()
        for s in words:
            if s[::-1] in vis:
                ans += 1
            else:
                vis.add(s)
        return ans
```

```go [sol-Go]
func maximumNumberOfStringPairs(words []string) (ans int) {
	vis := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if vis[y][x] {
			ans++
		} else {
			vis[x][y] = true
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{words}$ 的长度。字符串长度视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。
