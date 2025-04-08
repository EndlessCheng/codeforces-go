### 公式推导

首先计算不操作（或者原地不动）时的答案，然后计算操作对答案的**最大增量**。

有向左移动和向右移动两种情况，但其实两种情况可以用同一个式子表示，推导过程如下：

设 $a$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组为 $s$。

- 向左移动：假设 $a_i$ 移动到 $a_j$，那么下标在 $[j,i-1]$ 中的数的下标都加一，同时 $a_i\cdot i$ 减少为 $a_i\cdot j$，所以答案增加 $s_i - s_j - a_i\cdot(i-j)$。
- 向右移动：假设 $a_i$ 移动到 $a_j$，那么下标在 $[i+1,j]$ 中的数的下标都减一，同时 $a_i\cdot i$ 增加为 $a_i\cdot j$，所以答案增加 
  
    $$
    \begin{aligned}
        & -(s_{j+1} - s_{i+1}) + a_i\cdot(j-i)      \\
    ={} & -(s_{j+1} - s_{i+1} + a_i) + a_i + a_i\cdot(j-i)        \\
    ={} & -(s_{j+1} - s_i) + a_i\cdot(j+1-i)        \\
    ={} & -(s_{j'} - s_i) + a_i\cdot(j'-i)        \\
    ={} & s_i-s_{j'} - a_i\cdot(i-j')        \\
    \end{aligned}
    $$
    
    其中 $j'=j+1$（变量替换）。

可以发现两个式子是一样的，所以问题变成最大化

$$
s_i - s_j - a_i\cdot(i-j)
$$

其中 $0\le i\le n-1$，$0\le j\le n$。（本文讨论的下标从 $0$ 开始）

### 几何意义

把上式改成点积的形式，这样我们能得到来自几何意义上的观察。

设向量 $\mathbf{p} = (a_i, 1)$，向量 $\mathbf{v}_j = (j,-s_j)$。

那么

$$
\begin{aligned}
    & s_i - s_j - a_i\cdot(i-j)      \\
={} & a_i\cdot j - s_j - a_i\cdot i + s_i         \\
={} & \mathbf{p}\cdot \mathbf{v}_j - a_i\cdot i + s_i       \\
\end{aligned}
$$

根据点积的几何意义，我们求的是 $\mathbf{v}_j$ 在 $\mathbf{p}$ 方向上的投影长度，再乘以 $\mathbf{p}$ 的模长 $||\mathbf{p}||$。由于 $||\mathbf{p}||$ 是个定值，所以要最大化投影长度。

考虑 $\mathbf{v}_j$ 的**上凸包**（用 Andrew 算法计算），在凸包内的点，就像是山坳，比凸包顶点的投影长度短。所以只需考虑凸包顶点。

这样有一个很好的性质：顺时针遍历凸包顶点，$\mathbf{p}\cdot \mathbf{v}_j$ 会先变大再变小（单峰函数）。那么要计算最大值，就类似 [852. 山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/)，**二分**首个「下坡」的位置，具体见 [我的题解](https://leetcode.cn/problems/peak-index-in-a-mountain-array/solutions/2984800/er-fen-gen-ju-shang-po-huan-shi-xia-po-p-uoev/)。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt";"os";"sort")

type vec struct{ x, y int }
func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, tot, mx, s int
	Fscan(in, &n)
	a := make([]int, n)
	q := []vec{{-1, 0}} // s[-1] = 0
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i] * (i + 1)
		s += a[i]
		v := vec{i, -s} // 由于 i 是单调递增的，求上凸包无需排序
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, v)
	}

	s = 0
	for i, x := range a {
		p := vec{x, 1}
		j := sort.Search(len(q)-1, func(j int) bool { return p.dot(q[j]) > p.dot(q[j+1]) }) // 二分找峰顶
		s += x
		mx = max(mx, p.dot(q[j])-x*i+s)
	}
	Print(tot + mx)
}
```

**时间复杂度**：$\mathcal{O}(n\log n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
