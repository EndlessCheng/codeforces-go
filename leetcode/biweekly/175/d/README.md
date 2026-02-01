由于子数组的值是个分数，不好处理，我们可以先计算分子之和（或者说，把子数组的值乘以 $2$），最后返回时再除以 $2$。

## 划分型 DP

本题是标准的划分型 DP。根据 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/)「§5.3 约束划分个数」，定义 $f_{K,i}$ 表示把长为 $i$ 的前缀 $[0,i-1]$ 划分成恰好 $K$ 个子数组的分数（乘以 $2$）。

枚举最后一段子数组的左端点 $j$，问题变成把前缀 $[0,j-1]$ 划分成恰好 $K-1$ 个子数组的分数（乘以 $2$），即 $f_{K-1,j}$。

设 $\textit{nums}$ 的**前缀和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

把子数组 $[j,i-1]$ 的元素和用前缀和表示，即 $s_i - s_j$。

子数组的值（乘以 $2$）为

$$
(s_i - s_j)\cdot (s_i - s_j+1)
$$

取最小值，有

$$
f_{K,i} = \min_{j=0}^{i-1} f_{K-1,j} + (s_i - s_j)\cdot (s_i - s_j+1)
$$

初始值 $f_{0,0} = 0$，$f_{i,i-1}=\infty$。注意不需要初始化 $K<i-1$ 的状态，因为我们不会访问这些状态。

答案为 $f_{k,n}$。

这样做的时间复杂度是 $\mathcal{O}(n^2k)$，太慢了，如何优化？

## 斜率优化（凸包优化）

**注**：以下内容基于点积，和斜率优化本质是一样的。我推荐从点积的角度理解，相比斜率的角度，不需要做麻烦的式子变形，更方便。

**前置知识**：二维计算几何，凸包，Andrew 算法。

转移方程可以变形为

$$
f_{K,i} = s_i^2+s_i + \min_{j=0}^{i-1} f_{K-1,j} -2s_is_j + s_j^2-s_j
$$

把其中的

$$
f_{K-1,j} -2s_is_j + s_j^2-s_j
$$

改成点积的形式，这样我们能得到来自几何意义上的观察。

设向量 $\mathbf{v}_j = (s_j, f_{K-1,j} + s_j^2-s_j)$。

设向量 $\mathbf{p} = (-2s_i, 1)$。

那么我们求的是

$$
\min_{j=0}^{i-1} \mathbf{p}\cdot \mathbf{v}_j
$$

根据点积的几何意义，我们求的是 $\mathbf{v}_j$ 在 $\mathbf{p}$ 方向上的投影长度，再乘以 $\mathbf{p}$ 的模长 $||\mathbf{p}||$。由于 $||\mathbf{p}||$ 是个定值，所以要最小化投影长度。

考虑 $\mathbf{v}_j$ 的**下凸包**（用 Andrew 算法计算），在凸包内的点，比凸包顶点的投影长度长（注意 $\mathbf{p}$ 在凸包外面）。所以只需考虑凸包顶点。

> 由于 $s_j$ 是单调递增的，求下凸包无需排序。

这样有一个很好的性质：从左到右遍历凸包顶点，$\mathbf{p}\cdot \mathbf{v}_j$ 会先变小再变大（单峰函数）。那么要计算最小值，就类似 [852. 山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/)，**二分**首个「上坡」的位置，具体见 [我的题解](https://leetcode.cn/problems/peak-index-in-a-mountain-array/solutions/2984800/er-fen-gen-ju-shang-po-huan-shi-xia-po-p-uoev/)。

实际上不需要二分。由于 $-2s_i$ 是单调递减的，可以用单调队列维护凸包。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

其他语言稍后补充。

```py [sol-Python3]
# 注：由于代码使用了 Vec 类，跑得比较慢，直接在循环中做 dot det 更快
class Vec:
    __slots__ = 'x', 'y'

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __sub__(self, b: "Vec") -> "Vec":
        return Vec(self.x - b.x, self.y - b.y)

    def det(self, b: "Vec") -> int:
        return self.x * b.y - self.y * b.x

    def dot(self, b: "Vec") -> int:
        return self.x * b.x + self.y * b.y


class Solution:
    def minPartitionScore(self, nums: List[int], k: int) -> int:
        n = len(nums)
        pre = list(accumulate(nums, initial=0))
        f = [0] + [inf] * n

        for K in range(1, k + 1):
            s = pre[K - 1]
            q = [Vec(s, f[K - 1] + s * s - s)]
            for i in range(K, n - (k - K) + 1):
                s = pre[i]
                p = Vec(-2 * s, 1)
                while len(q) > 1 and p.dot(q[0]) >= p.dot(q[1]):
                    q = q[1:]

                v = Vec(s, f[i] + s * s - s)
                f[i] = p.dot(q[0]) + s * s + s

                while len(q) > 1 and (q[-1] - q[-2]).det(v - q[-1]) <= 0:
                    q.pop()
                q.append(v)

        return f[n] // 2
```

```go [sol-Go]
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x } // 如果乘法会溢出，用 detCmp
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func minPartitionScore(nums []int, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
	}

	for K := 1; K <= k; K++ {
		s := sum[K-1]
		q := []vec{{s, f[K-1] + s*s - s}}
		for i := K; i <= n-(k-K); i++ {
			s = sum[i]
			p := vec{-2 * s, 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}

			v := vec{s, f[i] + s*s - s}
			f[i] = p.dot(q[0]) + s*s + s

			// 读者可以把 detCmp 改成 det 感受下这个算法的效率
			// 目前 det 也能过，可以试试 hack 一下
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
	}

	return int64(f[n] / 2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-k)k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§5.3 约束划分个数**」和「**§11.7 斜率优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
