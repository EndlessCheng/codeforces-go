三变量问题，一般**枚举中间的变量**最简单。为什么？对比一下：

- 枚举 $i$，后续计算中还需要保证 $j$ 和 $k$ 的位置关系。
- 枚举 $j$，那么 $i$ 和 $k$ 自动被 $j$ 隔开，互相独立，后续计算中无需关心 $i$ 和 $k$ 的位置关系。

枚举中间的 $j$，问题变成：

- 在 $[0,j-1]$ 中，$\textit{nums}[j]\cdot 2$ 的出现次数。
- 在 $[j+1,n-1]$ 中，$\textit{nums}[j]\cdot 2$ 的出现次数。
- 在这些出现次数中，左右两边各选一个。根据乘法原理，把这两个出现次数相乘，加到答案中。

用哈希表（或者数组）统计 $j$ 左右每个数的出现次数。

- 右边的元素出现次数，可以先统计整个数组，然后再次遍历数组，**撤销** $[0,j]$ 中统计的元素出现次数，即为 $[j+1,n-1]$ 中的元素出现次数。
- 左边的元素出现次数，可以一边遍历 $\textit{nums}$，一边统计。

由于答案不超过 $n\cdot 10^5\cdot 10^5 = 10^{15}$，可以只在返回时取模。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1qsMxz6EEd/?t=9m46s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def specialTriplets(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        suf = Counter(nums)

        ans = 0
        pre = defaultdict(int)  # 比 Counter 快
        for x in nums:  # x = nums[j]
            suf[x] -= 1  # 撤销
            # 现在 pre 中的是 [0,j-1]，suf 中的是 [j+1,n-1]
            ans += pre[x * 2] * suf[x * 2]
            pre[x] += 1
        return ans % MOD
```

```java [sol-Java]
class Solution {
    public int specialTriplets(int[] nums) {
        final int MOD = 1_000_000_007;
        Map<Integer, Integer> suf = new HashMap<>();
        for (int x : nums) {
            suf.merge(x, 1, Integer::sum); // suf[x]++
        }

        long ans = 0;
        Map<Integer, Integer> pre = new HashMap<>();
        for (int x : nums) { // x = nums[j]
            suf.merge(x, -1, Integer::sum); // suf[x]-- // 撤销
            // 现在 pre 中的是 [0,j-1]，suf 中的是 [j+1,n-1]
            ans += (long) pre.getOrDefault(x * 2, 0) * suf.getOrDefault(x * 2, 0);
            pre.merge(x, 1, Integer::sum); // pre[x]++
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int specialTriplets(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        unordered_map<int, int> suf;
        for (int x : nums) {
            suf[x]++;
        }

        long long ans = 0;
        unordered_map<int, int> pre;
        for (int x : nums) { // x = nums[j]
            suf[x]--; // 撤销
            // 现在 pre 中的是 [0,j-1]，suf 中的是 [j+1,n-1]
            ans += 1LL * pre[x * 2] * suf[x * 2];
            pre[x]++;
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
func specialTriplets(nums []int) (ans int) {
	const mod = 1_000_000_007
	suf := map[int]int{}
	for _, x := range nums {
		suf[x]++
	}

	pre := map[int]int{}
	for _, x := range nums { // x = nums[j]
		suf[x]-- // 撤销
		// 现在 pre 中的是 [0,j-1]，suf 中的是 [j+1,n-1]
		ans += pre[x*2] * suf[x*2]
		pre[x]++
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

见下面数据结构题单的「**§0.2 枚举中间**」和动态规划题单的「**专题：前后缀分解**」。

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
