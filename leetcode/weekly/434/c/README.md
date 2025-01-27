## 方法一：枚举 + 状态机 DP

考虑把子数组中等于 $\textit{target}$ 的元素都变成 $k$。

由于 $\textit{nums}$ 至多有 $50$ 种不同元素，可以枚举 $\textit{target}=1,2,3,\ldots,50$（或者 $\textit{nums}$ 中的不同元素）。

$\textit{nums}$ 可以分为三段：

1. 左：被修改的子数组的左边。
2. 中：被修改的子数组。
3. 右：被修改的子数组的右边。

用状态机 DP 计算 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 最多有多少个数可以等于 $k$：

- $f[i+1][0]$ 表示左，或者说 $\textit{nums}[i]$ 在被修改的子数组的左侧，此时只能统计等于 $k$ 的元素个数。
- $f[i+1][1]$ 表示左+中，或者说 $\textit{nums}[i]$ 在被修改的子数组中，此时只能统计等于 $\textit{target}$ 的元素个数，这些数被修改成 $k$。
- $f[i+1][2]$ 表示左+中+右，或者说 $\textit{nums}[i]$ 在被修改的子数组的右侧，此时只能统计等于 $k$ 的元素个数。

从左到右遍历 $\textit{nums}$，设 $x=\textit{nums}[i]$，考虑转移来源：

- 「左」只能从「左」转移过来。如果 $x=k$，那么不选白不选，问题变成 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 最多有多少个数可以等于 $k$，即状态 $f[i][0]$，所以有转移方程 $f[i+1][0] = f[i][0] + 1$。如果 $x\ne k$，那么 $f[i+1][0] = f[i][0]$。
- 「左+中」可以从「左+中」或者「左」转移过来。同上，问题变成 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 最多有多少个数可以等于 $k$。如果 $x=\textit{target}$，那么 $f[i+1][1] = \max(f[i][1], f[i][0]) + 1$，否则 $f[i+1][1] = \max(f[i][1], f[i][0])$。这里从 $f[i][1]$ 转移过来，表示 $\textit{nums}[i-1]$ 也在被修改的子数组中；从 $f[i][0]$ 转移过来，表示 $\textit{nums}[i]$ 是被修改的子数组的第一个数。
- 「左+中+右」可以从「左+中+右」或者「左+中」转移过来。同上，问题变成 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 最多有多少个数可以等于 $k$。如果 $x=k$，那么 $f[i+1][2] = \max(f[i][2], f[i][1]) + 1$，否则 $f[i+1][2] = \max(f[i][2], f[i][1])$。这里从 $f[i][2]$ 转移过来，表示 $\textit{nums}[i-1]$ 也在被修改的子数组的右边；从 $f[i][1]$ 转移过来，表示 $\textit{nums}[i-1]$ 是被修改的子数组的最后一个数。

初始值 $f[0][0] = 0, f[0][1] = f[0][2] = -\infty$。想象 $\textit{nums}$ 第一个数左边还有一个位置，这个位置只能在「左」。

答案为 $\max(f[n][1], f[n][2])$。最后一个数可以在「中」也可以在「右」。

代码实现时，第一个维度可以优化掉，三个状态按照 $f_2\to f_1\to f_0$ 的顺序倒着更新，理由同 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/)。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15sFNewEia/?t=13m50s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        ans = 0
        for target in set(nums):
            f0, f1, f2 = 0, -inf, -inf
            for x in nums:
                f2 = max(f2, f1) + (x == k)
                f1 = max(f1, f0) + (x == target)
                f0 += (x == k)
            ans = max(ans, f1, f2)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequency(int[] nums, int k) {
        Set<Integer> set = new HashSet<>();
        for (int x : nums) {
            set.add(x);
        }

        int ans = 0;
        for (int target : set) {
            int f0 = 0;
            int f1 = Integer.MIN_VALUE;
            int f2 = Integer.MIN_VALUE;
            for (int x : nums) {
                f2 = Math.max(f2, f1) + (x == k ? 1 : 0);
                f1 = Math.max(f1, f0) + (x == target ? 1 : 0);
                f0 += (x == k ? 1 : 0);
            }
            ans = Math.max(ans, Math.max(f1, f2));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequency(vector<int>& nums, int k) {
        unordered_set<int> st(nums.begin(), nums.end());
        int ans = 0;
        for (int target : st) {
            int f0 = 0, f1 = INT_MIN, f2 = INT_MIN;
            for (int x : nums) {
                f2 = max(f2, f1) + (x == k);
                f1 = max(f1, f0) + (x == target);
                f0 += (x == k);
            }
            ans = max({ans, f1, f2});
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFrequency(nums []int, k int) (ans int) {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}

	for target := range set {
		f0, f1, f2 := 0, math.MinInt, math.MinInt
		for _, x := range nums {
			f2 = max(f2, f1) + b2i(x == k)
			f1 = max(f1, f0) + b2i(x == target)
			f0 += b2i(x == k)
		}
		ans = max(ans, f1, f2)
	}
	return
}

func b2i(b bool) int { if b { return 1 }; return 0 }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 中的不同元素个数。
- 空间复杂度：$\mathcal{O}(U)$。

## 方法二：状态机 DP 优化（一次遍历）

$f_0$ 和 $f_2$ 的定义不变。

设 $x=\textit{nums}[i]$。用 $f_1[x]$ 存储方法一中的 $\textit{target}=x$ 时的 $f_1$ 状态，转移方程为

$$
f_1[x] = \max(f_1[x], f_0) + 1
$$

其余 $f_1[y]\ (y\ne x)$ 不变（懒更新）。 

额外定义 $\textit{maxF}_1$ 表示 $f_1[x]$ 的最大值。

那么 $f_2$ 的转移方程为

$$
f_2 = \max(f_2, \textit{maxF}_1) + [x=k]
$$

这里为了方便把所有 $f_1[x]$ 初始化成 $0$。其实无论是方法一还是方法二，$f_1$ 初始化成 $0$ 或者任意负数都可以，毕竟要和 $f_0$ 这个非负数取最大值。

```py [sol-Python3]
class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        f0 = 0
        f1 = defaultdict(int)
        max_f1 = f2 = -inf
        for x in nums:
            f2 = max(f2, max_f1) + (x == k)
            f1[x] = max(f1[x], f0) + 1
            f0 += (x == k)
            max_f1 = max(max_f1, f1[x])
        return max(max_f1, f2)
```

```java [sol-Java]
class Solution {
    public int maxFrequency(int[] nums, int k) {
        int f0 = 0;
        int[] f1 = new int[51];
        int f2 = Integer.MIN_VALUE;
        int maxF1 = Integer.MIN_VALUE;
        for (int x : nums) {
            f2 = Math.max(f2, maxF1) + (x == k ? 1 : 0);
            f1[x] = Math.max(f1[x], f0) + 1;
            f0 += (x == k ? 1 : 0);
            maxF1 = Math.max(maxF1, f1[x]);
        }
        return Math.max(maxF1, f2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequency(vector<int>& nums, int k) {
        int f0 = 0, f1[51]{}, f2 = INT_MIN;
        int max_f1 = INT_MIN;
        for (int x : nums) {
            f2 = max(f2, max_f1) + (x == k);
            f1[x] = max(f1[x], f0) + 1;
            f0 += (x == k);
            max_f1 = max(max_f1, f1[x]);
        }
        return max(max_f1, f2);
    }
};
```

```go [sol-Go]
func maxFrequency(nums []int, k int) int {
	f0 := 0
	f1 := [51]int{}
	f2 := math.MinInt
	maxF1 := math.MinInt
	for _, x := range nums {
		f2 = max(f2, maxF1)
		f1[x] = max(f1[x], f0) + 1
		if x == k {
			f2++
			f0++
		}
		maxF1 = max(maxF1, f1[x])
	}
	return max(maxF1, f2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 中的不同元素个数。创建数组需要 $\mathcal{O}(U)$ 的时间。
- 空间复杂度：$\mathcal{O}(U)$。

## 变形题

改成可以修改**两个**不相交的子数组呢？

在 [视频讲解](https://www.bilibili.com/video/BV15sFNewEia/?t=13m50s) 第三题的最后讲了怎么做。

更多相似题目，见下面动态规划题单中的「**五、状态机 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
