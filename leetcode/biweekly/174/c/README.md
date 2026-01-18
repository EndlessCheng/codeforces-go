**前置题目**：[560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)，[我的题解](https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/)。

想一想，最后一步发生了什么？

设整个数组的异或和为 $s$。

- 如果最后一段的异或和是 $\textit{target}_1$，分割出最后一段后，剩余元素（$\textit{nums}$ 的前缀）的异或和为 $s\oplus \textit{target}_1$，问题变成：分割一个异或和为 $s\oplus \textit{target}_1$ 的数组，满足题目异或和交替的要求，且最后一段的异或和是 $\textit{target}_2$ 的方案数。
- 如果最后一段的异或和是 $\textit{target}_2$，分割出最后一段后，剩余元素（$\textit{nums}$ 的前缀）的异或和为 $s\oplus \textit{target}_2$，问题变成：分割一个异或和为 $s\oplus \textit{target}_2$ 的数组，满足题目异或和交替的要求，且最后一段的异或和是 $\textit{target}_1$ 的方案数。

回顾 560 题的技巧，我们可以用哈希表记录每个前缀的信息：

- 定义 $f_1[s]$ 表示分割一个异或和为 $s$ 的前缀，满足题目异或和交替的要求，且最后一段的异或和是 $\textit{target}_1$ 的方案数。
- 定义 $f_2[s]$ 表示分割一个异或和为 $s$ 的前缀，满足题目异或和交替的要求，且最后一段的异或和是 $\textit{target}_2$ 的方案数。

从左到右遍历 $\textit{nums}$，同时计算前缀异或和 $\textit{preSum}$。根据上面的讨论，有如下转移方程

$$
\begin{aligned}
f_1[\textit{preSum}] &\ += f_2[\textit{preSum} \oplus \textit{target}_1]      \\
f_2[\textit{preSum}] &\ += f_1[\textit{preSum} \oplus \textit{target}_2]      \\
\end{aligned}
$$

注意这两个式子要同时计算。

初始值：$f_2[0] = 1$。相当于在第一段前面有一个空前缀，异或和为 $0$。这样我们计算第一段时，就可以让第一段的方案数是 $1$。

答案为最后一轮循环（更新 $f_1$ 和 $f_2$ 之前）的 $f_1[\textit{preSum} \oplus \textit{target}_2] + f_2[\textit{preSum} \oplus \textit{target}_1]$。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def alternatingXOR(self, nums: List[int], target1: int, target2: int) -> int:
        MOD = 1_000_000_007
        f1 = defaultdict(int)
        f2 = defaultdict(int)
        f2[0] = 1
        pre_sum = 0
        for i, x in enumerate(nums):
            pre_sum ^= x
            last1 = f2[pre_sum ^ target1]  # [0,i] 的最后一段的异或和是 target1 的方案数
            last2 = f1[pre_sum ^ target2]  # [0,i] 的最后一段的异或和是 target2 的方案数
            if i == len(nums) - 1:
                return (last1 + last2) % MOD
            f1[pre_sum] = (f1[pre_sum] + last1) % MOD
            f2[pre_sum] = (f2[pre_sum] + last2) % MOD
```

```java [sol-Java]
class Solution {
    public int alternatingXOR(int[] nums, int target1, int target2) {
        final int MOD = 1_000_000_007;
        Map<Integer, Integer> f1 = new HashMap<>();
        Map<Integer, Integer> f2 = new HashMap<>();
        f2.put(0, 1);
        int preSum = 0;
        for (int i = 0; ; i++) {
            preSum ^= nums[i];
            int last1 = f2.getOrDefault(preSum ^ target1, 0); // [0,i] 的最后一段的异或和是 target1 的方案数
            int last2 = f1.getOrDefault(preSum ^ target2, 0); // [0,i] 的最后一段的异或和是 target2 的方案数
            if (i == nums.length - 1) {
                return (last1 + last2) % MOD;
            }
            f1.put(preSum, (f1.getOrDefault(preSum, 0) + last1) % MOD);
            f2.put(preSum, (f2.getOrDefault(preSum, 0) + last2) % MOD);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int alternatingXOR(vector<int>& nums, int target1, int target2) {
        constexpr int MOD = 1'000'000'007;
        unordered_map<int, int> f1;
        unordered_map<int, int> f2 = {{0, 1}};
        int pre_sum = 0;
        for (int i = 0; ; i++) {
            pre_sum ^= nums[i];
            int last1 = f2[pre_sum ^ target1]; // [0,i] 的最后一段的异或和是 target1 的方案数
            int last2 = f1[pre_sum ^ target2]; // [0,i] 的最后一段的异或和是 target2 的方案数
            if (i == nums.size() - 1) {
                return (last1 + last2) % MOD;
            }
            f1[pre_sum] = (f1[pre_sum] + last1) % MOD;
            f2[pre_sum] = (f2[pre_sum] + last2) % MOD;
        }
    }
};
```

```go [sol-Go]
func alternatingXOR(nums []int, target1, target2 int) int {
	const mod = 1_000_000_007
	f1 := map[int]int{}
	f2 := map[int]int{0: 1}
	preSum := 0
	for i, x := range nums {
		preSum ^= x
		last1 := f2[preSum^target1] // [0,i] 的最后一段的异或和是 target1 的方案数
		last2 := f1[preSum^target2] // [0,i] 的最后一段的异或和是 target2 的方案数
		if i == len(nums)-1 {
			return (last1 + last2) % mod
		}
		f1[preSum] = (f1[preSum] + last1) % mod
		f2[preSum] = (f2[preSum] + last2) % mod
	}
	panic("unreachable")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§7.4 合法子序列 DP**」。

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
