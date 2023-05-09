晚上 8:30[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

先简单记录一下思路，直播结束后继续更新题解和其它语言。

1. [二分](https://www.bilibili.com/video/BV1AP41137w7/)答案。
2. 贪心，每段城墙先向左膨胀，再向右膨胀，如果超过右侧的城墙则说明答案过大。如果都可以膨胀，则继续二分更大的答案。
3. 二分上界为中间剩余空间的平均值。

```go [sol1-Go]
func rampartDefensiveLine(rampart [][]int) (ans int) {
	n := len(rampart)
	leftSpace := rampart[n-1][0] - rampart[0][1]
	for _, p := range rampart[1 : n-1] {
		leftSpace -= p[1] - p[0]
	}
	return sort.Search(leftSpace/(n-2), func(mx int) bool {
		mx++
		preR := rampart[0][1]
		for i := 1; i < n-1; i++ {
			r := rampart[i][1]
			space := mx - (rampart[i][0] - preR)
			if space > 0 {
				r += space // 向右膨胀
				if r > rampart[i+1][0] { // 无法膨胀
					return true
				}
			}
			preR = r
		}
		return false
	})
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log (U/n))$，其中 $n$ 为 $\textit{rampart}$ 的长度，$U$ 为城墙范围。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
