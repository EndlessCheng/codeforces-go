思路是用最大堆来模拟，每次将堆顶减半，累加每次减半的值，直到不低于总和的一半。

虽然浮点数可以通过本题，但是本着能不用浮点就不用浮点的想法（毕竟浮点数计算会产生误差），将每个数都乘上一个 $2$ 的幂次（比如 $2^{20}$），因为可以 [证明](https://leetcode.cn/problems/minimum-operations-to-halve-array-sum/solution/onsuan-fa-by-hqztrue-jalf/) 每个数除 $2$ 的次数不会超过 $20$。

这样就可以愉快地用整数 + 堆来模拟了。

```py [sol-Python3]
class Solution:
    def halveArray(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            nums[i] <<= 20
        heapify_max(nums)

        ans = 0
        half = sum(nums) // 2
        while half > 0:
            half -= nums[0] // 2
            heapreplace_max(nums, nums[0] // 2)
            ans += 1
        return ans
```

```go [sol-Go]
func halveArray(nums []int) (ans int) {
	half := 0
	for i := range nums {
		nums[i] <<= 20
		half += nums[i]
	}

	h := hp{nums}
	heap.Init(&h)
	for half /= 2; half > 0; ans++ {
		half -= h.IntSlice[0] / 2
		h.IntSlice[0] /= 2
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice } // 继承 sort.IntSlice 的方法
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。原地修改，只用到常数额外空间。

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
