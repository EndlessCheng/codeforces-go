由于 $n\le 1500$，我们可以枚举子数组的左右端点。

外层循环枚举左端点，内层循环枚举右端点，不断向右扩大子数组长度。

贪心地，**把子数组内较小的数与子数组外较大的数交换**。所以我们要用两个数据结构，分别能快速计算：

- 子数组内的前 $k$ **小**元素和。
- 子数组外的前 $k$ **大**元素和。

这可以用**值域树状数组**维护。原理见 [OI-wiki](https://oi-wiki.org/ds/fenwick/#%E5%8D%95%E7%82%B9%E4%BF%AE%E6%94%B9%E6%9F%A5%E8%AF%A2%E5%85%A8%E5%B1%80%E7%AC%AC-k-%E5%B0%8F)。

但如果子数组内的第 $k$ 小比子数组外的第 $k$ 大还要大，那么不能交换。这意味着，实际交换次数可能小于 $k$。

难道要二分交换次数吗？这样总体时间复杂度是 $\mathcal{O}(n^2\log ^2 n)$，太慢了。

注意到，当子数组长度增加一，只有一个元素从外面进入子数组，交换次数至多增加或减少 $1$。

设元素进入子数组前，需要交换 $\textit{needSwap}$ 次。

- 如果子数组内的第 $\textit{needSwap}+1$ 个数 $<$ 子数组外的第 $\textit{needSwap}+1$ 个数，那么 $\textit{needSwap}$ 加一。
- 如果子数组内的第 $\textit{needSwap}$ 个数 $\ge$ 子数组外的第 $\textit{needSwap}$ 个数，那么 $\textit{needSwap}$ 减一。

[本题视频讲解](https://www.bilibili.com/video/BV1ptJw6hENZ/?t=20m58s)，欢迎点赞关注~

其他语言稍后补充。

```go
type pair struct{ cnt, sum int }
type fenwick []pair

func (t fenwick) update(i, num, val int) {
	for ; i < len(t); i += i & -i {
		t[i].cnt += num
		t[i].sum += val
	}
}

// 返回第 k 小的数（k 从 1 开始）
func (t fenwick) kth(k int, sorted []int) int {
	i := 0
	for b := 1 << 10; b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			i = nxt
		}
	}
	return sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (t fenwick) preSum(k int, sorted []int) (res int) {
	i := 0
	for b := 1 << 10; b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			res += t[nxt].sum
			i = nxt
		}
	}
	// 加上剩下的
	res += sorted[i] * k
	return
}

func maxSum(nums []int, k int) int64 {
	// 离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	n := len(nums)

	rank := make([]int, n)
	outTreeAll := make(fenwick, m+1)
	totalSum := 0
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		outTreeAll.update(rank[i], 1, x)
		totalSum += x
	}

	ans := math.MinInt

	// 枚举子数组左右端点
	for left := range nums {
		inTree := make(fenwick, m+1)
		outTree := slices.Clone(outTreeAll)
		needSwap := 0
		subSum := 0

		for right := left; right < n; right++ {
			// 更新子数组内外数据
			x := nums[right]
			rk := rank[right]
			subSum += x
			inTree.update(rk, 1, x)
			outTree.update(rk, -1, -x)

			ok := false
			sz := right - left + 1
			if needSwap < k && needSwap < sz && needSwap < n-sz {
				// 能否再交换一次
				if inTree.kth(needSwap+1, sorted) < outTree.kth(n-sz-needSwap, sorted) {
					ok = true
					needSwap++
				}
			}

			if !ok && needSwap > 0 {
				// 是否要减少交换次数
				if inTree.kth(needSwap, sorted) >= outTree.kth(n-sz-needSwap+1, sorted) {
					needSwap--
				}
			}

			// 计算通过交换导致的元素和的变化量
			delta := 0
			if needSwap > 0 {
				inSum := inTree.preSum(needSwap, sorted)
				outSum := totalSum - subSum - outTree.preSum(n-sz-needSwap, sorted)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}
```

## 优化

设 $\textit{nums}$ 中的正数个数为 $p$。跑一个定长滑窗，如果存在一个长为 $p$ 的窗口，其中 $正数个数+k\ge p$，那么可以把所有正数都聚在一起，此时直接返回所有正数之和（这是答案的上界）。 

```py [sol-Python3]

```

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§8.1 树状数组**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
