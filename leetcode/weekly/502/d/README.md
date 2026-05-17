计算 $\textit{nums}$ 的**后缀数组** $\textit{sa}$ 和**高度数组** $\textit{height}$。

其中 $\textit{height}[i]$ 定义为后缀 $\textit{sa}[i]$ 和后缀 $\textit{sa}[i-1]$ 的 LCP。

对于后缀 $\textit{sa}[i]$，它与其他后缀的最长公共前缀，来自后缀 $\textit{sa}[i]$ 和 $\textit{sa}[i-1]$ 的 LCP，以及后缀 $\textit{sa}[i]$ 和 $\textit{sa}[i+1]$ 的 LCP，这二者的最大值，即

$$
\max(\textit{height}[i],\textit{height}[i+1])
$$

这个值再加一，即为后缀 $\textit{sa}[i]$ 中的最短**唯一**前缀（可以用反证法证明，如果不唯一，那么 $\textit{height}[i]$ 或 $\textit{height}[i+1]$ 还能更大，矛盾）。

注意 $\max(\textit{height}[i],\textit{height}[i+1])+1$ 不能超过后缀 $\textit{sa}[i]$ 的长度，即

$$
\max(\textit{height}[i],\textit{height}[i+1])+1 \le n - \textit{sa}[i]
$$

如果上式成立，用 $\max(\textit{height}[i],\textit{height}[i+1])+1$ 更新答案的最小值。

> 注：本题也可以用二分答案 + 字符串哈希做，时间复杂度 $\mathcal{O}(n\log n)$。

[本题视频讲解](https://www.bilibili.com/video/BV18gLE6VETZ/?t=43m37s)，欢迎点赞关注~

```go
func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 把 1 个 int 拆成 3 个 byte（题目保证 nums[i] <= 1e5），从而可以调用库函数计算后缀数组
	tmp := make([]byte, 0, n*3)
	for _, x := range nums {
		tmp = append(tmp, byte(x>>16), byte(x>>8), byte(x))
	}

	type _tp struct {
		_  []byte
		sa []int32
	}
	_sa := (*_tp)(unsafe.Pointer(suffixarray.New(tmp))).sa

	sa := make([]int32, 0, n)
	for _, p := range _sa {
		if p%3 == 0 { // 是 3 的倍数的 _sa[i] 就对应着 nums 的 sa[i]
			sa = append(sa, p/3)
		}
	}

	// 计算后缀名次数组
	// 后缀 nums[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 nums 在后缀字典序中的排名，rank[n-1] 即 nums[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, n)
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(nums[sa[i]:], nums[sa[i-1]:])  (i > 0)
	// 获取 nums[i] 所在位置的高度：height[rank[i]]
	height := make([]int, n)
	h := 0
	// 计算 nums 与 nums[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 nums[1:] 与 nums[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 nums[2:] 与 nums[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 nums[n-1:] 与 nums[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 nums[0] 和 nums[sa[rank[0]-1]] 这两个数
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && nums[i+h] == nums[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	ans := n
	for i, h := range height {
		// 对于后缀 nums[sa[i]:]，其长为 uniqueLength 的前缀是唯一的
		uniqueLength := h + 1
		if i < n-1 {
			uniqueLength = max(h, height[i+1]) + 1
		}
		// 注意 uniqueLength 不能超过后缀 nums[sa[i]:] 的长度
		if uniqueLength <= n-int(sa[i]) {
			ans = min(ans, uniqueLength)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。注：库函数用的是 SA-IS 算法。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面字符串题单的「**四、字符串哈希**」和「**八、后缀数组/后缀自动机**」。

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
