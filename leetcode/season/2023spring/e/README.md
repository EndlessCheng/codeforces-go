### 本题视频讲解

见[【力扣杯2023春·个人赛】](https://www.bilibili.com/video/BV1dg4y1j78A/) 第五题。

### 思路

> 注：题意中的「两颗异色棋子之间恰好只有一颗棋子」允许中间有空格子。

分类讨论，对于一行或者一列的 RB 序列，有以下 $7$ 种情况：

- 空。
- 只有一个 B。
- 只有一个 R。
- 连续多个 B。
- 连续多个 R。
- BR 交替，且以 B 结尾。
- BR 交替，且以 R 结尾。

> 为什么单独一个 B/R 也算一个状态？因为既可以变成连续多个 B/R，又可以变成 BR 交替。而「连续多个 B/R」和 「BR 交替」是无法互相转换的。

遍历棋盘的过程就是在不断生成序列的过程，那么向序列末尾添加 B 或者 R，这些状态就会互相转换，形成一个 $7\cdot 2$ 的**转换关系表**，用数组 $\textit{trans}$ 记录，其中 $-1$ 表示非法转换。（见代码）

由于有 $7$ 种情况，可以用 $3$ 个比特存储，所有列合并在一起就需要 $3m$ 个比特，这可以用一个二进制数 $\textit{mask}$ 来存储。

写一个记忆化搜索，$\textit{DFS}(r,\textit{mask})$ 表示枚举到第 $r$ 行，每列的 RB 序列的状态组合起来是 $\textit{mask}$ 时，继续向后枚举可以得到多少合法的方案。

考虑从下一行的 $\textit{DFS}(r+1,\textit{mask}')$ 转移过来，如何枚举 $\textit{mask}'$ 呢？

写一个**暴力搜索**来枚举，枚举第 $r$ 行怎么放棋子是合法的。定义 $\textit{dfs}(c,\textit{rowMask},\textit{colMask})$ 表示搜索到第 $c$ 列，当前行的 RB 序列状态是 $\textit{rowMask}$，每列的 RB 序列的状态组合起来是 $\textit{colMask}$：

- 如果当前符号是 B，则往当前行和当前列的末尾添加 B，根据 $\textit{trans}$ 来看是否合法，如果合法就继续搜索。
- 如果当前符号是 R，则往当前行和当前列的末尾添加 R，根据 $\textit{trans}$ 来看是否合法，如果合法就继续搜索。
- 如果当前符号是 ?，则往当前行和当前列的末尾添加 B、添加 R、或者什么也不加，根据 $\textit{trans}$ 来看是否合法，如果合法就继续搜索。
- 如果当前符号是 .，什么也不加，继续搜索。
- $c=m$ 时，表示对于当前行，搜索到一个合法的方案，也就是 $\textit{mask}'=\textit{colMask}$，那么可以继续递归 $\textit{DFS}(r+1,\textit{mask}')$。

当 $r=n$ 时，表示找到了一个合法的方案，返回 $1$。

递归入口：$\textit{DFS}(0,0)$。

> 注：代码实现时，如果 $n<m$ 就旋转棋盘，这样可以保证 $m\le 5$。

> 其它语言稍后补充。

```Python [sol1-Python3]
TRANS = (
    # (当前序列末尾添加 B 之后的状态，当前序列末尾添加 R 之后的状态)
    (1, 2),   # 0: 空
    (3, 6),   # 1: 一个 B
    (5, 4),   # 2: 一个 R
    (3, -1),  # 3: 连续多个 B
    (-1, 4),  # 4: 连续多个 R
    (-1, 6),  # 5: BR 交替，且以 B 结尾
    (5, -1),  # 6: BR 交替，且以 R 结尾
)

class Solution:
    def getSchemeCount(self, n: int, m: int, a: List[str]) -> int:
        if n < m:  # 保证 n >= m
            a = [list(col) for col in zip(*a)]
            n, m = m, n
        @cache
        def DFS(r: int, mask: int) -> int:
            if r == n:  # 找到 1 个合法方案
                return 1
            # 写一个爆搜，生成出所有的合法状态
            def dfs(c: int, row_mask: int, col_mask: int) -> int:
                if c == m:  # 方案合法
                    return DFS(r + 1, col_mask)  # 枚举下一行
                def nxt(color: int) -> int:
                    rm = TRANS[row_mask][color]  # 新的 rowMask
                    if rm < 0: return 0  # 非法
                    c3 = c * 3
                    cm = TRANS[(col_mask >> c3) & 7][color]  # 新的 colMask 的第 c 列
                    if cm < 0: return 0 # 非法
                    cm = col_mask & ~(7 << c3) | (cm << c3)  # 修改 colMask 的第 c 列
                    return dfs(c + 1, rm, cm)
                b = a[r][c]
                if b == 'B': return nxt(0)
                if b == 'R': return nxt(1)
                if b == '.': return dfs(c + 1, row_mask, col_mask)
                return dfs(c + 1, row_mask, col_mask) + nxt(0) + nxt(1)
            return dfs(0, 0, mask)
        return DFS(0, 0)
```

```go [sol1-Go]
// 每一行的含义：{当前序列末尾添加 B 后的状态，当前序列末尾添加 R 后的状态}
var trans = [7][2]int{
	{1, 2},  // 空
	{3, 6},  // 只有一个 B
	{5, 4},  // 只有一个 R
	{3, -1}, // 连续多个 B
	{-1, 4}, // 连续多个 R
	{-1, 6}, // BR 交替，且以 B 结尾
	{5, -1}, // BR 交替，且以 R 结尾
}

func getSchemeCount(n, m int, g []string) int64 {
	a := make([][]byte, n)
	for i, row := range g {
		a[i] = []byte(row)
	}
	if n < m {
		a = rotate(a) // 保证 n >= m
		n, m = m, n
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<(m*3))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var DFS func(int, int) int
	DFS = func(r, mask int) int {
		if r == n { // 找到 1 个合法方案
			return 1
		}
		ptr := &memo[r][mask]
		if *ptr != -1 {
			return *ptr
		}
		// 写一个爆搜，生成所有的合法状态
		var dfs func(int, int, int) int
		dfs = func(c, rowMask, colMask int) (res int) {
			if c == m { // 方案合法
				return DFS(r+1, colMask) // 枚举下一行
			}
			next := func(color int) int {
				rm := trans[rowMask][color] // 新的 rowMask
				if rm < 0 { // 非法
					return 0
				}
				c3 := c * 3
				cm := trans[colMask>>c3&7][color] // 新的 colMask 的第 c 列
				if cm < 0 { // 非法
					return 0
				}
				return dfs(c+1, rm, colMask&^(7<<c3)|cm<<c3) // 修改 colMask 的第 c 列
			}
			switch a[r][c] {
			case 'B': // 填 B
				return next(0)
			case 'R': // 填 R
				return next(1)
			case '?': // 留空 / 填 B / 填 R
				return dfs(c+1, rowMask, colMask) + next(0) + next(1)
			default: // 留空
				return dfs(c+1, rowMask, colMask)
			}
		}
		*ptr = dfs(0, 0, mask)
		return *ptr
	}
	return int64(DFS(0, 0))
}

func rotate(a [][]byte) [][]byte {
	n, m := len(a), len(a[0])
	b := make([][]byte, m)
	for i := range b {
		b[i] = make([]byte, n)
	}
	for i, r := range a {
		for j, v := range r {
			b[j][n-1-i] = v
		}
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n21^{m})$。这里保证 $m\le 5\le \sqrt{nm}$（如果 $n<m$ 则旋转棋盘）。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。这里状态个数为 $\mathcal{O}(n7^{m})$，单个状态的计算时间为 $\mathcal{O}(3^{m})$，因此时间复杂度为 $\mathcal{O}(n7^{m}3^{m})=\mathcal{O}(n21^{m})$。考虑到很多状态是非法的，在全为 ? 的情况下，$\textit{dfs}$ 最多调用 $5601518$ 次。
- 空间复杂度：$\mathcal{O}(n8^{m})$。为了方便用位运算，实际用了 $\mathcal{O}(n8^{m})$ 的空间来存储状态。
