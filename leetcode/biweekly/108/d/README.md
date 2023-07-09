下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

在 $m$ 和 $n$ 都很大的情况下，网格中有大量的 $2\times 2$ 的子矩阵是没有黑色格子的。只需要考虑有黑色格子的子矩阵。

如果 $(x,y)$ 处有黑色格子，那么子矩阵左上角在 $(x-1,y-1),(x-1,y),(x,y-1),(x,y)$ 都是包含这个黑色格子的，统计这些子矩阵中有多少黑色格子，加到答案中。

代码实现时，注意不要重复统计，可以用哈希表 $\textit{vis}$ 来记录统计过的子矩阵左上角。

最后不含黑色格子的子矩阵个数就是

$$
(m-1)\cdot (n-1) - \text{len}(\textit{vis})
$$

```py [sol-Python3]
class Solution:
    def countBlackBlocks(self, m: int, n: int, coordinates: List[List[int]]) -> List[int]:
        s = set(map(tuple, coordinates))
        arr = [0] * 5
        vis = set()
        for x, y in coordinates:
            for i in range(max(x - 1, 0), min(x + 1, m - 1)):
                for j in range(max(y - 1, 0), min(y + 1, n - 1)):
                    if (i, j) not in vis:
                        vis.add((i, j))
                        cnt = ((i, j) in s) + ((i, j + 1) in s) + \
                              ((i + 1, j) in s) + ((i + 1, j + 1) in s)
                        arr[cnt] += 1
        arr[0] = (m - 1) * (n - 1) - len(vis)
        return arr
```

```go [sol-Go]
func countBlackBlocks(m, n int, coordinates [][]int) []int64 {
	type pair struct{ x, y int }
	set := make(map[pair]int, len(coordinates))
	for _, p := range coordinates {
		set[pair{p[0], p[1]}] = 1
	}

	arr := make([]int64, 5)
	vis := make(map[pair]bool, len(set)*4)
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := max(x-1, 0); i <= x && i < m-1; i++ {
			for j := max(y-1, 0); j <= y && j < n-1; j++ {
				if !vis[pair{i, j}] {
					vis[pair{i, j}] = true
					cnt := set[pair{i, j}] + set[pair{i, j + 1}] +
						   set[pair{i + 1, j}] + set[pair{i + 1, j + 1}]
					arr[cnt]++
				}
			}
		}
	}
	arr[0] = int64(m-1)*int64(n-1) - int64(len(vis))
	return arr
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k)$，其中 $k$ 为 $\textit{coordinates}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。
