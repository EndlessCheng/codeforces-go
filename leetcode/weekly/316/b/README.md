## 方法一：暴力枚举

暴力枚举即可，可以在不包含因子 $k$ 时提前退出循环。

```py [sol-Python3]
class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(len(nums)):
            g = 0
            for j in range(i, len(nums)):
                g = gcd(g, nums[j])
                if g % k:
                    break
                if g == k:
                    ans += 1
        return ans
```

```go [sol-Go]
func subarrayGCD(nums []int, k int) (ans int) {
	for i := range nums {
		g := 0
		for _, x := range nums[i:] {
			g = gcd(g, x)
			if g%k > 0 {
				break
			}
			if g == k {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+\log U))$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。外层循环时，单看 $g=\textit{nums}[i]$，它因为求 GCD 减半的次数是 $\mathcal{O}(\log U)$ 次，因此内层循环的时间复杂度为 $\mathcal{O}(n+\log U)$，所以总的时间复杂度为 $\mathcal{O}(n(n+\log U))$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：利用 GCD 的性质

**前置知识**：[LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。

```py [sol-Python3]
class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        ans = 0
        a = []  # [GCD，相同 GCD 区间的右端点]
        i0 = -1
        for i, x in enumerate(nums):
            if x % k:  # 保证后续求的 GCD 都是 k 的倍数
                a = []
                i0 = i
                continue
            a.append([x, i])
            # 原地去重，因为相同的 GCD 都相邻在一起
            j = 0
            for p in a:
                p[0] = gcd(p[0], x)
                if a[j][0] != p[0]:
                    j += 1
                    a[j] = p
                else:
                    a[j][1] = p[1]
            del a[j + 1:]
            if a[0][0] == k:  # a[0][0] >= k
                ans += a[0][1] - i0
        return ans
```

```go [sol-Go]
func subarrayGCD(nums []int, k int) (ans int) {
	type result struct{ v, i int }
	var a []result
	i0 := -1
	for i, v := range nums {
		if v%k > 0 {
			a = nil
			i0 = i
			continue
		}
		for j, p := range a {
			a[j].v = gcd(p.v, v)
		}
		a = append(a, result{v, i})
		j := 0
		for _, q := range a[1:] {
			if a[j].v != q.v {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]
		if a[0].v == k {
			ans += a[0].i - i0
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。单看每个元素，它因为求 GCD 减半的次数是 $O(\log U)$ 次，并且每次去重的时间复杂度也为 $\mathcal{O}(\log U)$，因此时间复杂度为 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(\log U)$。

## 相似题目

- [Codeforces 475D. CGCDSSQ](https://codeforces.com/problemset/problem/475/D)

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
