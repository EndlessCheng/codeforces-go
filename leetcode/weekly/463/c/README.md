## 分析

考虑删除 $\textit{nums}$ 的某个子数组 $a$（$a$ 的元素和是 $k$ 的倍数）。$a$ 的更内部元素如何删除，是否会影响到 $a$ 整体的删除？

比如 $k=4$，考虑删除子数组 $a=[2,1,3,2]$。有两种删除方案：

- 先删除更内部的 $[1,3]$，得到 $[2,2]$，再删除 $[2,2]$。
- 直接删除 $[2,1,3,2]$。

这两种删除方案都可以删除 $[2,1,3,2]$。不会因为先删除了 $[2,2]$，导致子数组 $a=[2,1,3,2]$ 无法删除。

一般地，如果 $a$ 的元素和是 $k$ 的倍数，且删除的更内部的子数组元素和也是 $k$ 的倍数，那么删除更内部元素后，$a$ 的剩余元素之和是不变的，仍然为 $k$ 的倍数。

这个性质意味着，我们可以首先考虑删除 $\textit{nums}$ 的最右边的子数组，无需担心这个子数组内部元素对删除的影响。

## 寻找子问题

考虑 $\textit{nums}[n-1]$ 删或不删：

- 不删，问题变成前缀 $[0,n-2]$ 的最小和。
- 删，设子数组的左端点为 $j$，删除子数组 $[j,n-1]$ 后，问题变成前缀 $[0,j-1]$ 的最小和。

删或不删都会把原问题变成一个**和原问题相似的、规模更小的子问题**。

## 状态定义与状态转移方程

根据上面的讨论，定义 $f[i+1]$ 表示前缀 $[0,i]$ 的最小和。其中 $+1$ 是为了留出位置给 $f[0]$ 表示空前缀。

考虑 $\textit{nums}[i]$ 删或不删：

- 不删，问题变成前缀 $[0,i-1]$ 的最小和，即 $f[i]$，加上 $\textit{nums}[i]$，得到 $f[i+1] = f[i] + \textit{nums}[i]$。
- 删，可能有多个子数组的左端点，满足子数组的元素和是 $k$ 的倍数。设这些左端点为 $j_1,j_2,\dots,j_m$。删除子数组 $[j,i]$ 后，问题变成前缀 $[0,j-1]$ 的最小和，即 $f[j]$。取最小值，得 $f[i+1] = \min\limits_{j\in \{j_1,j_2,\dots,j_m\}} f[j]$。

现在的问题是，如何快速计算 $\min\limits_{j\in \{j_1,j_2,\dots,j_m\}} f[j]$？

首先要找到这些 $j_1,j_2,\dots,j_m$。计算 $\textit{nums}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组 $\textit{sum}$，每一项都模 $k$。比如当 $\textit{sum}[i+1] = 6$ 时，左边还有 $\textit{sum}[j_1] = \textit{sum}[j_2] = 6$，即子数组 $[j_1,i]$ 和子数组 $[j_2,i]$ 的元素和模 $k$ 都是 $0$。考虑维护这些 $\textit{sum}[j] = 6$ 对应的 $f[j]$ 的最小值，记作 $\textit{minF}$，比如 $\textit{minF}[6]$ 表示所有模 $k$ 为 $6$ 的前缀和对应的状态值的最小值。如此一来，转移方程就可以优化成 $f[i+1] = \textit{minF}[\textit{sum}[i+1]]$。

不删和删取最小值，得

$$
f[i+1] = \min(f[i] + \textit{nums}[i], \textit{minF}[\textit{sum}[i+1]])
$$

然后更新 $\textit{sum}[i+1]$ 对应的 $f[i+1]$ 的最小值，即用上面计算的 $f[i+1]$ 更新 $\textit{minF}[\textit{sum}[i+1]]$ 的最小值。

初始值：$f[0] = 0$，空前缀的元素和等于 $0$。

答案：$f[n]$。

代码实现时，$\textit{sum}$ 和 $f$ 都可以优化为一个变量。

为什么可以一遍累加一边取模，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 数组实现

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minArraySum(self, nums: list[int], k: int) -> int:
        min_f = [inf] * k
        min_f[0] = 0  # s[0] = 0，对应的 f[0] = 0
        f = s = 0
        for x in nums:
            s = (s + x) % k
            # 不删除 x，那么转移来源为 f + x
            # 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            # 其中剩余前缀的元素和模 k 等于 s，对应的 f 值的最小值记录在 min_f[s] 中
            f = min(f + x, min_f[s])
            # 维护前缀和 s 对应的最小和，由于上面计算了 min，这里无需再计算 min
            min_f[s] = f
        return f
```

```java [sol-Java]
class Solution {
    public long minArraySum(int[] nums, int k) {
        long[] minF = new long[k];
        Arrays.fill(minF, Long.MAX_VALUE);
        minF[0] = 0; // sum[0] = 0，对应的 f[0] = 0
        long f = 0;
        int sum = 0;
        for (int x : nums) {
            sum = (sum + x) % k;
            // 不删除 x，那么转移来源为 f + x
            // 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            // 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
            f = Math.min(f + x, minF[sum]);
            // 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
            minF[sum] = f;
        }
        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minArraySum(vector<int>& nums, int k) {
        vector<long long> min_f(k, LLONG_MAX);
        min_f[0] = 0; // sum[0] = 0，对应的 f[0] = 0
        long long f = 0;
        int sum = 0;
        for (int x : nums) {
            sum = (sum + x) % k;
            // 不删除 x，那么转移来源为 f + x
            // 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            // 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 min_f[sum] 中
            f = min(f + x, min_f[sum]);
            // 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
            min_f[sum] = f;
        }
        return f;
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k int) int64 {
	minF := make([]int, k)
	// sum[0] = 0，对应的 f[0] = 0
	for i := 1; i < k; i++ {
		minF[i] = math.MaxInt
	}
	f, sum := 0, 0
	for _, x := range nums {
		sum = (sum + x) % k
		// 不删除 x，那么转移来源为 f + x
		// 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
		// 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
		f = min(f+x, minF[sum])
		// 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
		minF[sum] = f
	}
	return int64(f)
}
```

## 哈希表实现

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minArraySum(self, nums: list[int], k: int) -> int:
        min_f = defaultdict(lambda: inf)
        min_f[0] = 0  # s[0] = 0，对应的 f[0] = 0
        f = s = 0
        for x in nums:
            s = (s + x) % k
            # 不删除 x，那么转移来源为 f + x
            # 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            # 其中剩余前缀的元素和模 k 等于 s，对应的 f 值的最小值记录在 min_f[s] 中
            f = min(f + x, min_f[s])
            # 维护前缀和 s 对应的最小和，由于上面计算了 min，这里无需再计算 min
            min_f[s] = f
        return f
```

```java [sol-Java]
class Solution {
    public long minArraySum(int[] nums, int k) {
        Map<Integer, Long> minF = new HashMap<>();
        minF.put(0, 0L); // sum[0] = 0，对应的 f[0] = 0
        long f = 0;
        int sum = 0;
        for (int x : nums) {
            sum = (sum + x) % k;
            // 不删除 x，那么转移来源为 f + x
            // 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            // 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
            f = Math.min(f + x, minF.getOrDefault(sum, Long.MAX_VALUE));
            // 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
            minF.put(sum, f);
        }
        return f;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minArraySum(vector<int>& nums, int k) {
        unordered_map<int, long long> min_f = {{0, 0}}; // sum[0] = 0，对应的 f[0] = 0
        long long f = 0;
        int sum = 0;
        for (int x : nums) {
            sum = (sum + x) % k;
            // 不删除 x
            f += x;
            // 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
            // 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 min_f[sum] 中
            if (min_f.contains(sum)) {
                f = min(f, min_f[sum]);
            }
            // 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
            min_f[sum] = f;
        }
        return f;
    }
};
```

```go [sol-Go]
func minArraySum(nums []int, k int) int64 {
	minF := map[int]int{0: 0} // sum[0] = 0，对应的 f[0] = 0
	f, sum := 0, 0
	for _, x := range nums {
		sum = (sum + x) % k
		// 不删除 x
		f += x
		// 删除以 x 结尾的子数组，问题变成剩余前缀的最小和
		// 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
		if mn, ok := minF[sum]; ok {
			f = min(f, mn)
		}
		// 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
		minF[sum] = f
	}
	return int64(f)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+k)$ 或 $\mathcal{O}(n)$。

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
