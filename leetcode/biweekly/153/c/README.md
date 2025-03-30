题目给出的式子有子数组和，我们先用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 将其简化。

定义 $\textit{sumNum}[i+1]$ 表示 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素和。

定义 $s[i+1]$ 表示 $\textit{cost}[0]$ 到 $\textit{cost}[i]$ 的元素和。

题目给出的式子转换成

$$
\begin{aligned}
    & (\textit{sumNum}[r+1] + k\cdot i) \cdot (s[r+1] - s[l])      \\
={} & \textit{sumNum}[r+1] \cdot (s[r+1] - s[l]) + k\cdot i \cdot (s[r+1] - s[l])       \\
\end{aligned}
$$

如果能把式子中的 $i$ 去掉，我们就可以写一个 $\mathcal{O}(n^2)$ 的划分型 DP。

横看成岭侧成峰，换一个角度看待 $i \cdot (s[r+1] - s[l])$。

假设划分成了三段，$\textit{cost}$ 的子数组和分别为 $A,B,C$。

这三段的 $i \cdot (s[r+1] - s[l])$ 分别为 $A,2B,3C$，累加得

$$
\begin{aligned}
    & A+2B+3C      \\
={} & (A+B+C) + (B+C) + C        \\
\end{aligned}
$$

如此变形后，我们可以把 $A+B+C$ 当作第一段的 $i \cdot (s[r+1] - s[l])$，把 $B+C$ 当作第二段的 $i \cdot (s[r+1] - s[l])$，把 $C$ 当作第三段的 $i \cdot (s[r+1] - s[l])$。

换句话说，我们可以跨越时空，把未来要计算的内容，放到现在计算！

式子中的 $i \cdot (s[r+1] - s[l])$ 可以替换成 $s[n] - s[l]$，因为 $A+B+C,B+C,C$ 都是 $\textit{cost}$ 的后缀和。

题目给出的式子替换成

$$
\textit{sumNum}[r+1] \cdot (s[r+1] - s[l]) + k\cdot (s[n] - s[l])
$$

> 注意上式和原式并不一定相等，但计算所有子数组的上式之和后，是相等的。

根据 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/)「§5.2 最优划分」，定义 $f[i+1]$ 表示下标 $[0,i]$ 分割后的最小总代价。

枚举最后一个子数组的左端点 $j$，问题变成下标 $[0,j-1]$ 分割后的最小总代价，即 $f[j]$。其中 $j$ 最小是 $0$，最大是 $i$。

取最小值，有

$$
f[i+1] = \min_{j=0}^{i} f[j] + \textit{sumNum}[i+1] \cdot (s[i+1] - s[j]) + k\cdot (s[n] - s[j])
$$

初始值 $f[0]=0$。

答案为 $f[n]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int], cost: List[int], k: int) -> int:
        n = len(nums)
        s = list(accumulate(cost, initial=0))  # cost 的前缀和
        f = [0] * (n + 1)
        for i, sum_num in enumerate(accumulate(nums), 1):  # 这里把 i 加一了，下面不用加一
            f[i] = min(f[j] + sum_num * (s[i] - s[j]) + k * (s[n] - s[j])
                       for j in range(i))
        return f[n]
```

```java [sol-Java]
class Solution {
    public long minimumCost(int[] nums, int[] cost, int k) {
        int n = nums.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + cost[i]; // cost 的前缀和
        }

        long[] f = new long[n + 1];
        int sumNum = 0;
        for (int i = 1; i <= n; i++) { // 注意这里 i 从 1 开始，下面不用把 i 加一
            sumNum += nums[i - 1];
            f[i] = Long.MAX_VALUE;
            for (int j = 0; j < i; j++) {
                f[i] = Math.min(f[i], f[j] + (long) sumNum * (s[i] - s[j]) + k * (s[n] - s[j]));
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(vector<int>& nums, vector<int>& cost, int k) {
        int n = nums.size();
        vector<int> s(n + 1);
        partial_sum(cost.begin(), cost.end(), s.begin() + 1); // cost 的前缀和

        vector<long long> f(n + 1, LLONG_MAX);
        f[0] = 0;
        int sum_num = 0;
        for (int i = 1; i <= n; i++) { // 注意这里 i 从 1 开始，下面不用把 i 加一
            sum_num += nums[i - 1];
            for (int j = 0; j < i; j++) {
                f[i] = min(f[i], f[j] + 1LL * sum_num * (s[i] - s[j]) + k * (s[n] - s[j]));
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minimumCost(nums, cost []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, c := range cost {
		s[i+1] = s[i] + c // cost 的前缀和
	}

	f := make([]int, n+1)
	sumNum := 0
	for i, x := range nums {
		i++ // 这里把 i 加一了，下面不用加一
		sumNum += x
		res := math.MaxInt
		for j := range i {
			res = min(res, f[j]+sumNum*(s[i]-s[j])+k*(s[n]-s[j]))
		}
		f[i] = res
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
