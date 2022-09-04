下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

由于数据范围很小，我们可以枚举所有大小为 $\textit{cols}$ 的列编号的集合，对于每个集合，遍历 $\textit{mat}$，统计所有 $1$ 被覆盖的行的个数，个数的最大值即为答案。

代码实现时，我们可以用二进制表示集合，二进制的第 $i$ 位为 $1$ 表示 $i$ 在集合中，为 $0$ 表示 $i$ 不在集合中。这样可以用二进制枚举集合，同时把 $\textit{mat}$ 的每一行也用二进制表示，从而做到 $O(1)$ 判断行中的所有 $1$ 是否被覆盖。

```py [sol1-Python3]
class Solution:
    def maximumRows(self, mat: List[List[int]], cols: int) -> int:
        ans = 0
        mask = [sum(v << j for j, v in enumerate(row)) for i, row in enumerate(mat)]
        for set in range(1 << len(mat[0])):
            if set.bit_count() == cols:  # 集合的大小等于 cols，符合题目要求
                ans = max(ans, sum(row & set == row for row in mask))  # row & set = row 表示 row 是 set 的子集，所有 1 都被覆盖
        return ans
```

```go [sol1-Go]
func maximumRows(mat [][]int, cols int) (ans int) {
	m, n := len(mat), len(mat[0])
	mask := make([]int, m)
	for i, row := range mat {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 0; set < 1<<n; set++ {
		if bits.OnesCount(uint(set)) != cols { // 跳过大小不等于 cols 的集合
			continue
		}
		cnt := 0
		for _, row := range mask {
			if row&set == row { // row 是 set 的子集，所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

上面的代码有很多无效枚举，即大小不等于 $\textit{cols}$ 的集合，如何优化呢？

通过使用 Gosper's Hack，我们可以在 $O(1)$ 的时间内找到下一个大小为 $\textit{cols}$ 的集合。

我会在下午的直播中介绍这个算法。

#### 复杂度分析

- 时间复杂度：$O(m\cdot C_{n}^{\textit{cols}})$。其中 $m$ 为 $\textit{mat}$ 的行数，$n$ 为 $\textit{mat}$ 的列数。
- 空间复杂度：$O(m)$。

```py [sol2-Python3]
class Solution:
    def maximumRows(self, mat: List[List[int]], cols: int) -> int:
        ans = 0
        mask = [sum(v << j for j, v in enumerate(row)) for i, row in enumerate(mat)]
        set = (1 << cols) - 1
        while set < 1 << len(mat[0]):
            ans = max(ans, sum(row & set == row for row in mask))  # row & set = row 表示 row 是 set 的子集，所有 1 都被覆盖
            lb = set & -set
            x = set + lb
            set = (set ^ x) // lb >> 2 | x
        return ans
```

```go [sol2-Go]
func maximumRows(mat [][]int, cols int) (ans int) {
	m, n := len(mat), len(mat[0])
	mask := make([]int, m)
	for i, row := range mat {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 1<<cols - 1; set < 1<<n; {
		cnt := 0
		for _, row := range mask {
			if row&set == row { // row 是 set 的子集，所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
		lb := set & -set
		x := set + lb
		// 下式等价于 set = (set^x)/lb>>2 | x
		set = (set^x)>>bits.TrailingZeros(uint(lb))>>2 | x
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
