下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

暴力枚举即可。

```py [sol1-Python3]
class Solution:
    def twoEditWords(self, queries: List[str], dictionary: List[str]) -> List[str]:
        ans = []
        for q in queries:
            for s in dictionary:
                if sum(x != y for x, y in zip(q, s)) <= 2:
                    ans.append(q)
                    break
        return ans
```

```go [sol1-Go]
func twoEditWords(queries, dictionary []string) (ans []string) {
	for _, q := range queries {
		for _, s := range dictionary {
			c := 0
			for i := range s {
				if q[i] != s[i] {
					c++
				}
			}
			if c <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(mkn)$，其中 $m$ 为 $\textit{queries}$ 的长度，$k$ 为 $\textit{dictionary}$ 的长度，$n$ 为 $\textit{queries}[i]$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
