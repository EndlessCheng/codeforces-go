提供一个简单的思路。

题目要我们把数组切成若干段，每一段的元素和都 $\le b$。如果首尾两段的元素和 $\le b$，那么首尾两段可以合并成一段。

**枚举**最后一刀切在 $i$，我们需要知道两个信息：

1. 一共切了多少刀，或者说前 $i$ 个数分成了多少段。
2. 第一刀切在哪。如果首尾两段的元素和 $\le b$，那么分的段数可以减少一。

由于数组元素非负，段的右端点越大，左端点也越大，所以固定右端点时，左端点的最小值可以用滑动窗口计算。

一共切了多少刀，可以用 DP 算：

- 定义 $f_i$ 表示前 $i$ 个数至少要分成多少段。
- 设前 $i$ 个数的最后一段的左端点为 $\ell$，那么问题变成前 $\ell$ 个数至少要分成多少段，即 $f_i = f_{\ell} + 1$。初始值 $f_0=0$。

第一刀切在哪，也可以用 DP 算：

- 定义 $\textit{begin}_i$ 表示前 $i$ 个数的第二段的左端点。
- 如果 $f_i=2$，那么 $\textit{begin}_i=\ell$；否则，$\textit{begin}_i=\textit{begin}_\ell$。

设前缀和数组为 $s$。由于 $i$ 越大 $f_i$ 也越大，所以当我们发现 $s_n - (s_i - s_{\textit{begin}_{\ell}})\le b$ 时，说明首尾两段可以拼成一段，直接输出 $\max(f_i,2)$ 即可。这里和 $2$ 取最大值是因为存在找到第一段后剩余元素就 $\le b$ 的情况，此时 $f_i=1$ 但答案为 $2$。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, q, b int
	Fscan(in, &n, &q)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &sum[i])
		sum[i] += sum[i-1] // 前缀和
	}
	total := sum[n]

	f := make([]int, n+1)
	begin := make([]int, n+1)
	for range q {
		Fscan(in, &b)
		if total <= b { // 只有一段
			Println(1)
			continue
		}
		l := 0
		for i := 1; ; i++ {
			for sum[i]-sum[l] > b { // 滑动窗口
				l++
			}
			f[i] = f[l] + 1 // 前 i 个数至少要分 f[i] 段
			if f[i] == 2 {
				begin[i] = l // 记录第二段的开始位置
			} else {
				begin[i] = begin[l]
			}
			// 剩余部分 <= b，说明首尾两段可以拼成一段，或者相当于把剩余的后缀拼到了第一段，所以答案就是 f[i]
			if total-(sum[i]-sum[begin[i]]) <= b {
				Println(max(f[i], 2)) // 可能找到第一段就触发 if 了，此时答案是 2
				break
			}
		}
	}
}
```

时间复杂度：$\mathcal{O}(nq)$。

相似题目：[3464. 正方形上的点之间的最大距离](https://leetcode.cn/problems/maximize-the-distance-between-points-on-a-square/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
