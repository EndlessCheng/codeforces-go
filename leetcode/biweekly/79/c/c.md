设点 $i$ 的度数（与点 $i$ 相邻的城市数）为 $\textit{deg}[i]$，点 $i$ 被安排的整数值为 $p[i]$，问题即最大化

$$
\sum_{i=0}^{n-1} \textit{deg}[i]\cdot p[i]
$$

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728) 可知，$\textit{deg}$ 最小的安排 $1$，次小的安排 $2$，依此类推。因此排序后累加即得到答案。

```Python [sol1-Python3]
class Solution:
    def maximumImportance(self, n: int, roads: List[List[int]]) -> int:
        deg = [0] * n
        for x, y in roads:
            deg[x] += 1
            deg[y] += 1
        deg.sort()
        return sum(i * d for i, d in enumerate(deg, 1))
```

```go [sol1-Go]
func maximumImportance(n int, roads [][]int) (ans int64) {
	deg := make([]int, n)
	for _, r := range roads {
		deg[r[0]]++
		deg[r[1]]++
	}
	sort.Ints(deg)
	for i, d := range deg {
		ans += int64(i+1) * int64(d)
	}
	return
}
```
