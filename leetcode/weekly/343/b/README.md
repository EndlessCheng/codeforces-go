下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

首先遍历 $\textit{mat}$，用一个 $\textit{pos}$ 数组记录 $\textit{mat}[i][j]$ 在 $\textit{mat}$ 中的位置。

然后遍历 $\textit{arr}[i]$，同时用两个数组 $\textit{rowCnt}$ 和 $\textit{colCnt}$ 记录每行每列的涂色个数。

如果出现某一行或某一列上都被涂色的情况，就返回 $i$。

```py [sol1-Python3]
class Solution:
    def firstCompleteIndex(self, arr: List[int], mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        pos = [0] * (m * n + 1)
        for i, row in enumerate(mat):
            for j, x in enumerate(row):
                pos[x] = (i, j)
        row_cnt = [0] * m
        col_cnt = [0] * n
        for i, x in enumerate(arr):
            r, c = pos[x]
            row_cnt[r] += 1
            col_cnt[c] += 1
            if row_cnt[r] == n or col_cnt[c] == m:
                return i
```

```go [sol1-Go]
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	type pair struct{ r, c int }
	pos := make([]pair, m*n+1)
	for i, row := range mat {
		for j, x := range row {
			pos[x] = pair{i, j}
		}
	}

	rowCnt := make([]int, m)
	colCnt := make([]int, n)
	for i, x := range arr {
		p := pos[x]
		rowCnt[p.r]++
		colCnt[p.c]++
		if rowCnt[p.r] == n || colCnt[p.c] == m {
			return i
		}
	}
	return -1
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m+n)$。
