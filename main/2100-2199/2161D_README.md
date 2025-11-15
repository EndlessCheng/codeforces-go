$\mathcal{O}(n)$ 做法。

### 从特殊到一般

考虑下面两个弱化版的问题。

**问题一**：去掉 $i<j$ 的限制，怎么做？

这题是 [LC740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/)，[我的题解](https://leetcode.cn/problems/delete-and-earn/solutions/3061028/zhi-yu-da-jia-jie-she-pythonjavaccgojsru-e5gg/)。

**问题二**：有 $i<j$ 的限制，但 $a$ 中只有 $1$ 和 $2$，怎么做？

这种情况下，保留的数（子序列）中的 $2$ 必须在 $1$ 的左边。

倒着遍历 $a$，考虑以 $a_i$ 开头的子序列，定义 $f_i$ 表示以 $a_i$ 开头的子序列的最长长度。

- 如果 $a_i = 1$，那么子序列只能全为 $1$，$f_i = f_j + 1$，其中 $j$ 是上一个 $1$ 的下标。
- 如果 $a_i = 2$，那么把 $2$ 拼在以 $a_{i+1}$ 开头的子序列的前面，所以 $f_i = f_{i+1}+1$。

### 回到原题

总体上是个值域打家劫舍问题，按照值域从小到大计算。

定义 $f_i$ 表示由 $\le a_i$ 的元素组成的子序列的最长长度，其中 $a_i$ 必选，且等于 $a_i$ 的其余元素都在 $a_i$ 的右边。

如果 $a$ 中有 $1,2,3$，那么当我们计算 $a_i=3$ 的 $f_i$ 时，用「选或不选」思考，有两种决策：

- 子序列包含 $2$，那么需要知道满足 $j > i$ 且 $2\le a_j\le 3$ 的最大 $f_j$，把 $3$ 添加在 $f_j$ 对应子序列的前面，得到 $f_i = f_j + 1$。
- 子序列不包含 $2$，那么需要知道满足 $a_j < 2$ 的最大 $f_j$，且下标 $\ge i$ 的 $3$ 都可以选，得到 $f_i = f_j + c$，其中 $c$ 是下标 $\ge i$ 的 $3$ 的个数。
- 两种情况取最大值。

一般地，枚举 $x=1,2,\ldots,n$，设当前在计算 $a_i = x$ 的 $f_i$，我们需要维护：

- 满足 $j > i$ 且 $x-1\le a_j\le x$ 的最大 $f_j$，记作 $\textit{sufMax}$。可以用双指针倒着遍历等于 $x$ 的下标列表，以及等于 $x-1$ 的下标列表。
- 满足 $a_j \le x-2$ 的最大 $f_j$，记作 $\textit{mx}$。

转移方程为

$$
f_i = \max(\textit{sufMax}+1, \textit{mx} + c)
$$

其中 $c$ 是下标 $\ge i$ 的 $x$ 的个数。

答案为 $n-\max(f)$。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt";"os";"slices")

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([][]int, n+1)
		for i := range n {
			Fscan(in, &v)
			pos[v] = append(pos[v], i)
		}

		f := make([]int, n+1)
		mx := 0 // a[i] <= x-2 的最大 f[i]
		for x := 1; x <= n; x++ {
			p := pos[x-1] // 下标列表
			q := pos[x]
			sufMax := 0
			j := len(p) - 1 // 双指针倒着遍历下标列表 p 和 q
			for i := len(q) - 1; i >= 0; i-- {
				for j >= 0 && p[j] > q[i] {
					sufMax = max(sufMax, f[p[j]])
					j--
				}
				sufMax++
				// max(选 x-1, 不选 x-1)
				f[q[i]] = max(sufMax, mx+len(q)-i)
			}
			for _, i := range p {
				mx = max(mx, f[i])
			}
		}
		Fprintln(out, n-slices.Max(f))
	}
}
```

**时间复杂度**：$\mathcal{O}(n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
