[视频讲解](https://www.bilibili.com/video/BV1QT41127kJ/) 包含**二维差分模板**的讲解，欢迎收看~

---

读者需要先了解二维前缀和的思想，可以看看 [304. 二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/)。

二维差分可以结合二维前缀和与一维差分的思想推导出来。当我们对一个左上角在 $(x1,y1)$，右下角在 $(x2,y2)$ 矩形区域全部增加 $x$ 时，相当于在二维差分矩阵上对 $(x1,y1)$ 增加 $x$，对 $(x1,y2+1)$ 和 $(x2+1,y1)$ 减少 $x$，由于这样两个地方都减少了 $x$，我们还需要在 $(x2+1,y2+1)$ 处增加 $x$，读者可以用二维前缀和对比体会这一做法。

更新结束后，我们需要从二维差分矩阵还原出二维计数矩阵，这可以通过对二维差分矩阵求二维前缀和求出。

```py [sol1-Python3]
class Solution:
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        m = n

        # 二维差分模板
        diff = [[0] * (n + 1) for _ in range(m + 1)]
        for r1, c1, r2, c2 in queries:
            r2 += 1
            c2 += 1
            diff[r1][c1] += 1
            diff[r1][c2] -= 1
            diff[r2][c1] -= 1
            diff[r2][c2] += 1

        # 用二维前缀和复原
        ans = [[0] * (n + 1) for _ in range(m + 1)]
        for i, row in enumerate(diff[:n]):
            for j, x in enumerate(row[:m]):
                ans[i + 1][j + 1] = ans[i + 1][j] + ans[i][j + 1] - ans[i][j] + x
        del ans[0]
        for row in ans:
            del row[0]
        return ans
```

```go [sol1-Go]
func rangeAddQueries(n int, queries [][]int) [][]int {
	m := n

	// 二维差分模板
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, m+1)
	}
	update := func(r1, c1, r2, c2, x int) {
		r2++
		c2++
		diff[r1][c1] += x
		diff[r1][c2] -= x
		diff[r2][c1] -= x
		diff[r2][c2] += x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原
	ans := make([][]int, n+1)
	ans[0] = make([]int, m+1)
	for i, row := range diff[:n] {
		ans[i+1] = make([]int, m+1)
		for j, x := range row[:m] {
			ans[i+1][j+1] = ans[i+1][j] + ans[i][j+1] - ans[i][j] + x
		}
	}
	ans = ans[1:]
	for i, row := range ans {
		ans[i] = row[1:]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n^2+q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n^2)$。

#### 相似题目

- [2132. 用邮票贴满网格图](https://leetcode.cn/problems/stamping-the-grid/)
