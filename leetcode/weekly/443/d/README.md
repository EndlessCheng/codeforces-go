滑动窗口中位数 + 划分型 DP（Python/Java/C++/Go）

---

## 把数组中的数都变成哪个数最优？

给定数组，每次操作可以把其中一个数加一/减一，把所有数都变成一样的，最少要操作多少次？

把所有数都变成数组的**中位数**是最优的。

[两种证明方法](https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/)。

## 滑动窗口中位数

见 [480. 滑动窗口中位数](https://leetcode.cn/problems/sliding-window-median/)，这可以用对顶堆做，具体见 [我的题解](https://leetcode.cn/problems/sliding-window-median/solutions/3628827/295-ti-lan-shan-chu-dui-pythonjavacgojsr-66ch/)。

## 如何计算操作次数？

需要额外维护**堆中元素之和**。

见[【一图秒懂】距离和](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。

本题 $j$ 为 $\textit{left}$（见 480 题解中的定义）的大小。

都变成 $\textit{left}$ 的堆顶 $v$，那么：

- 蓝色面积为 $v\cdot j - S_L$，其中 $S_L$ 为 $\textit{left}$ 的元素和。
- 绿色面积为 $S_R - v\cdot (x-j)$，其中 $S_R$ 为 $\textit{right}$ 的元素和。

操作次数即为面积之和

$$
v\cdot j - S_L + S_R - v\cdot (x-j)
$$

## 划分型 DP

根据 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/)「§5.3 约束划分个数」，定义 $f[i][j]$ 表示下标 $[0,j-1]$ 中选 $i$ 个长为 $x$ 的子数组的最小操作次数。

分类讨论：

- 子数组不含 $\textit{nums}[j-1]$，问题变成下标 $[0,j-2]$ 中选 $i$ 个长为 $x$ 的子数组的最小操作次数，即 $f[i][j-1]$。
- 子数组含 $\textit{nums}[j-1]$，子数组的左端点为 $j-x$，问题变成下标 $[0,j-x-1]$ 中选 $i-1$ 个长为 $x$ 的子数组的最小操作次数，即 $f[i-1][j-x]$。

取最小值，有

$$
f[i][j] =\min(f[i][j-1], f[i-1][j-x] + \textit{dis}[j-x]) 
$$

其中 $dis[i]$ 是滑动窗口算出的距离和（结果保存在子数组左端点）。

初始值 $f[0][0]=0,f[i][i\cdot x-1]=\infty$。

答案为 $f[k][n]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

其他语言稍后补充。

```go [sol-Go]
func minOperations(nums []int, x, k int) int64 {
	n := len(nums)
	dis := medianSlidingWindow(nums, x)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= k; i++ {
		f[i][i*x-1] = math.MaxInt
		for j := i * x; j <= n-(k-i)*x; j++ { // 左右留出足够空间给其他子数组
			f[i][j] = min(f[i][j-1], f[i-1][j-x]+dis[j-x]) // j-x 为子数组左端点
		}
	}
	return int64(f[k][n])
}

// 返回 nums 的所有长为 windowSize 的子数组的（到子数组中位数的）距离和
func medianSlidingWindow(nums []int, windowSize int) []int {
	ans := make([]int, len(nums)-windowSize+1)
	left := newLazyHeap()  // 最大堆（元素取反）
	right := newLazyHeap() // 最小堆

	for i, in := range nums {
		// 1. 进入窗口
		if left.size == right.size {
			left.push(-right.pushPop(in))
		} else {
			right.push(-left.pushPop(-in))
		}

		l := i + 1 - windowSize
		if l < 0 { // 窗口大小不足 k
			continue
		}

		// 2. 计算答案
		v := -left.top()
		s1 := v*left.size + left.sum // sum 取反
		s2 := right.sum - v*right.size
		ans[l] = s1 + s2

		// 3. 离开窗口
		out := nums[l]
		if out <= -left.top() {
			left.remove(-out)
			if left.size < right.size {
				left.push(-right.pop()) // 平衡两个堆的大小
			}
		} else {
			right.remove(out)
			if left.size > right.size+1 {
				right.push(-left.pop()) // 平衡两个堆的大小
			}
		}
	}

	return ans
}

func newLazyHeap() *lazyHeap {
	return &lazyHeap{removeCnt: map[int]int{}}
}

// 懒删除堆
type lazyHeap struct {
	sort.IntSlice
	removeCnt map[int]int // 每个元素剩余需要删除的次数
	size      int         // 实际大小
	sum       int         // 堆中元素和
}

// 必须实现的两个接口
func (h *lazyHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 删除
func (h *lazyHeap) remove(v int) {
	h.removeCnt[v]++ // 懒删除
	h.size--
	h.sum -= v
}

// 正式执行删除操作
func (h *lazyHeap) applyRemove() {
	for h.removeCnt[h.IntSlice[0]] > 0 {
		h.removeCnt[h.IntSlice[0]]--
		heap.Pop(h)
	}
}

// 查看堆顶
func (h *lazyHeap) top() int {
	h.applyRemove()
	return h.IntSlice[0]
}

// 出堆
func (h *lazyHeap) pop() int {
	h.applyRemove()
	h.size--
	h.sum -= h.IntSlice[0]
	return heap.Pop(h).(int)
}

// 入堆
func (h *lazyHeap) push(v int) {
	if h.removeCnt[v] > 0 {
		h.removeCnt[v]-- // 抵消之前的删除
	} else {
		heap.Push(h, v)
	}
	h.size++
	h.sum += v
}

// push(v) 然后 pop()
func (h *lazyHeap) pushPop(v int) int {
	if h.size > 0 && v > h.top() { // 最小堆，v 比堆顶大就替换堆顶
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n(n-kx))$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。可以滚动数组优化到 $\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
