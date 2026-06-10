## 引入

在上一题 [3956. 非重叠子数组最大和 I](https://leetcode.cn/problems/maximum-sum-of-m-non-overlapping-subarrays-i/) 中，我们实际上解决了如下问题：

- 从 $\textit{nums}$ 中选择**恰好** $x$ 个不重叠子数组，所选元素之和的最大值 $f(x)$。

$f(x)$ 的折线图是什么样的？

例如 $\textit{nums}=[3,1,-4,1,-5,9]$：

- 选 $0$ 个子数组，和为 $0$。
- 选 $1$ 个子数组，最优是选 $[9]$，和为 $9$。
- 选 $2$ 个子数组，最优是选 $[3,1]$ 和 $[9]$，和为 $13$。
- 选 $3$ 个子数组，最优是选 $[3,1]$、$[1]$ 和 $[9]$，和为 $14$。
- 选 $4$ 个子数组，最优是选 $[3]$、$[1]$、$[1]$ 和 $[9]$，和为 $14$。
- 选 $5$ 个子数组，最优是选 $[3]$、$[1]$、$[-4]$、$[1]$ 和 $[9]$，和为 $10$。
- 选 $6$ 个子数组，最优是选 $[3]$、$[1]$、$[-4]$、$[1]$、$[-5]$ 和 $[9]$，和为 $5$。

![lc3956.svg](https://pic.leetcode.cn/1780887600-kdBkoq-lc3956.svg){:width=500px}

这是个凸函数，在 $x=3$ 或者 $x=4$ 时取到最大值。

万丈高楼平地起。在上一题中，为了解决恰好选 $x$ 个子数组的问题，我们必须先解决恰好选 $1,2,\ldots,x-1$ 个子数组的问题。当 $x=m$ 时，时间复杂度为 $\mathcal{O}(nm)$。如果题目没有 $m$ 的约束，我们可以用单调队列优化至 $\mathcal{O}(n)$ 时间复杂度（见后文）。

这说明：

- 一方面，直接求解 $x=m$ 的问题是困难的。
- 另一方面，解决没有 $m$ 约束的问题是容易的，但算出的子数组个数可能超过 $m$。

能否在「没有 $m$ 约束」的基础上，引入一些条件，从而可以调控我们算出的子数组个数？

## 拉格朗日松弛 (Lagrangian Relaxation)

在凸优化中，一个常用的技巧是增加一个**惩罚项**。对于本题，每选一个子数组，就把元素和减少 $k$。

对于「没有 $m$ 约束」的问题，$k$ 越大，多选一个子数组的收益更可能是负的，所以选的子数组越少；反之，$k$ 越小，选的子数组越多。这样，我们就能通过**二分** $k$，从而让我们算出的子数组个数恰好等于 $m$。

> **注**：国内算法竞赛圈把这个技巧叫做「WQS 二分」。

本题是「至多」选 $m$ 个子数组，我们可以先算出没有约束时的子数组个数，如果 $\le m$，那么就直接找到了答案。否则必须「恰好」选 $m$ 个子数组。这样就把「至多」转化成了「恰好」。

> 如果题目还要求至少选 $\textit{least}$ 个子数组，怎么做？见文末的「变形题」。

## 没有 m 约束

对于「没有 $m$ 约束」的问题，我们除了要计算所选元素之和的最大值，还需要计算**选了多少个子数组**，这样才能与 $m$ 比大小，从而通过二分调控 $k$ 的值。

类似上一题，定义 $f[i]$ 表示从前 $i$ 个数（下标 $[0,i-1]$）中选一些不重叠的子数组，所选元素之和的最大值，但每选一个子数组，就要从元素之和中减去 $k$。

状态转移方程为

$$
f[i] = \max\left\{f[i-1], \max_{L=i-r}^{i-\ell} f[L] + s[i] - s[L] - k \right\}
$$

定义 $\textit{cnt}[i]$ 表示 $f[i]$ 对应着我们选了 $\textit{cnt}[i]$ 个子数组。

- 如果从 $f[i-1]$ 转移到 $f[i]$，那么没选子数组，所以 $\textit{cnt}[i] = \textit{cnt}[i-1]$。
- 如果从 $f[L]$ 转移到 $f[i]$，那么选了一个子数组 $[L,i-1]$，所以 $\textit{cnt}[i] = \textit{cnt}[L] + 1$。

如果 $f[i-1] = f[L]$，那么 $\textit{cnt}[i]$ 应该等于 $\textit{cnt}[i-1]$，还是等于 $\textit{cnt}[L] + 1$？选小的，还是选大的？

都可以。对于本题，由于我们要先判断 $k=0$ 时的子数组个数是否 $\le m$，可以优先选 $\textit{cnt}$ 值更小的，这样更可能不用跑后面的二分。

初始值 $f[0] = 0$。

最终选了 $\textit{cnt}[n]$ 个子数组，所选元素之和减去 $\textit{cnt}[n]$ 个 $k$ 后的最大值为 $f[n]$，即所选元素之和的最大值为 $f[n] + \textit{cnt}[n]\cdot k$。

## 细节

#### 1)

可能有多种选子数组的方式，都得到了相同的 $f[n]$。例如在 $k=10$ 的情况下，选 $3$ 个、$4$ 个还是 $5$ 个子数组，都得到了 $f[n]=100$。

在 DP 值相等时，我们的策略是优先选 $\textit{cnt}$ 值更小的。所以可能会出现 $m=4$，但最终二分出来的 $\textit{cnt}[n] = 3$ 的情况，怎么办？

在 $k=10$ 的情况下，既然选 $3$ 个子数组还是选 $4$ 个子数组都是一样的 $f[n]=100$，那么直接认为 $f[n]$ 对应着我们选了 $m=4$ 个子数组（这样答案也更大）。最终答案为 $f[n]$ 加上减去的 $m$ 个 $k$，即 $f[n] + mk$。

#### 2)

下面代码采用开区间二分。使用闭区间或者半闭半开区间也是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。由于 $k=0$ 的情况我们已经提前计算过了，所以二分中的 $k=0$ 对应的子数组个数一定大于 $m$，不满足要求。
- 开区间右端点初始值：$S+1$，其中 $S$ 是 $\textit{nums}$ 中的正数之和。当 $k=S$ 时，选一个子数组的收益 $\le 0$，所以算法只会选一个子数组，一定满足要求。由于我的实现是在二分中更新答案（这样可以避免在二分结束后再跑一次 DP），所以右端点初始值多加了 $1$，以保证二分覆盖到 $k=S$ 的情况。

```py [sol-Python3]
max = lambda a, b: b if b > a else a  # 手写 max 更快

class Solution:
    def maximumSum(self, nums: list[int], m: int, l: int, r: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))  # nums 的前缀和

        # 没有 m 约束，但每选一个子数组就要把元素和减少 k
        def dp_without_limit(k: int) -> tuple[int, int]:
            f = [(0, 0)] * (n + 1)  # (DP 值, 子数组个数)，其中子数组个数取反，方便比大小
            q = deque()
            res = (-inf, 0)

            for i in range(l, n + 1):
                # 1. 入
                j = i - l
                v = (f[j][0] - s[j], f[j][1])
                while q and (f[q[-1]][0] - s[q[-1]], f[q[-1]][1]) <= v:
                    q.pop()
                q.append(j)

                # 2. 更新答案
                j = q[0]
                choose = (f[j][0] - s[j] + s[i] - k, f[j][1] - 1)  # 注意子数组个数取反了，加一变成减一
                # choose 保证我们至少选了一个子数组
                res = max(res, choose)  # DP 值相等时，子数组个数小的更优

                # 更新 DP
                f[i] = max(f[i - 1], choose)

                # 3. 出，下一轮循环队首离开窗口
                if j <= i - r:
                    q.popleft()

            return res[0], -res[1]

        res, cnt = dp_without_limit(0)
        if cnt <= m:  # 直接满足题目要求
            return res

        # 现在专注于解决「选恰好 m 个子数组」的问题
        ans = 0
        pos_sum = sum(x for x in nums if x > 0)  # nums 中的正数之和
        left, right = 0, pos_sum + 1
        while left + 1 < right:
            k = (left + right) // 2
            res, cnt = dp_without_limit(k)
            if cnt <= m:
                ans = res + m * k  # 见题解【细节 1】
                right = k
            else:
                left = k
        return ans
```

```java [sol-Java]
class Solution {
    // DP 值, 子数组个数
    private record Pair(long f, int cnt) {
    }

    // 相等的时候，子数组个数更大的劣
    private boolean less(Pair a, Pair b) {
        return a.f < b.f || a.f == b.f && a.cnt > b.cnt;
    }

    public long maximumSum(int[] nums, int m, int l, int r) {
        int n = nums.length;
        long[] s = new long[n + 1]; // nums 的前缀和
        long posSum = 0; // nums 中的正数之和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
            if (nums[i] > 0) {
                posSum += nums[i];
            }
        }

        Pair res0 = dpWithoutLimit(0, n, l, r, s);
        if (res0.cnt <= m) { // 直接满足题目要求
            return res0.f;
        }

        // 现在专注于解决「选恰好 m 个子数组」的问题
        long ans = 0;
        long left = 0;
        long right = posSum + 1;
        while (left + 1 < right) {
            long k = left + (right - left) / 2;
            Pair res = dpWithoutLimit(k, n, l, r, s);
            if (res.cnt <= m) {
                ans = res.f + m * k; // 见题解【细节 1】
                right = k;
            } else {
                left = k;
            }
        }
        return ans;
    }

    // 没有 m 约束，但每选一个子数组就要把元素和减少 k
    private Pair dpWithoutLimit(long k, int n, int l, int r, long[] s) {
        Pair[] f = new Pair[n + 1];
        Arrays.fill(f, 0, l, new Pair(0, 0));
        Deque<Integer> q = new ArrayDeque<>();
        Pair res = new Pair(Long.MIN_VALUE, 0);

        for (int i = l; i <= n; i++) {
            // 1. 入
            int j = i - l;
            Pair v = new Pair(f[j].f - s[j], f[j].cnt);
            while (!q.isEmpty() && less(new Pair(f[q.peekLast()].f - s[q.peekLast()], f[q.peekLast()].cnt), v)) {
                q.pollLast();
            }
            q.addLast(j);

            // 2. 更新答案
            j = q.peekFirst();
            Pair choose = new Pair(f[j].f - s[j] + s[i] - k, f[j].cnt + 1);
            if (less(res, choose)) {
                // choose 保证我们至少选了一个子数组
                res = choose;
            }

            // 更新 DP
            f[i] = less(f[i - 1], choose) ? choose : f[i - 1];

            // 3. 出，下一轮循环队首离开窗口
            if (j <= i - r) {
                q.pollFirst();
            }
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSum(vector<int>& nums, int m, int l, int r) {
        int n = nums.size();
        vector<long long> s(n + 1); // nums 的前缀和
        long long pos_sum = 0; // nums 中的正数之和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
            if (nums[i] > 0) {
                pos_sum += nums[i];
            }
        }

        // 没有 m 约束，但每选一个子数组就要把元素和减少 k
        using Pair = pair<long long, int>; // (DP 值, 子数组个数)，其中子数组个数取反，方便比大小
        auto dp_without_limit = [&](long long k) -> Pair {
            vector<Pair> f(n + 1);
            deque<int> q;
            Pair res = {LLONG_MIN, 0};

            for (int i = l; i <= n; i++) {
                // 1. 入
                int j = i - l;
                Pair v = {f[j].first - s[j], f[j].second};
                while (!q.empty() && Pair{f[q.back()].first - s[q.back()], f[q.back()].second} <= v) {
                    q.pop_back();
                }
                q.push_back(j);

                // 2. 更新答案
                j = q.front();
                Pair choose = {f[j].first - s[j] + s[i] - k, f[j].second - 1}; // 注意子数组个数取反了，加一变成减一
                // choose 保证我们至少选了一个子数组
                res = max(res, choose); // DP 值相等时，子数组个数小的更优

                // 更新 DP
                f[i] = max(f[i - 1], choose);

                // 3. 出，下一轮循环队首离开窗口
                if (j <= i - r) {
                    q.pop_front();
                }
            }

            res.second *= -1; // 恢复成正数
            return res;
        };

        auto [res, cnt] = dp_without_limit(0);
        if (cnt <= m) { // 直接满足题目要求
            return res;
        }

        // 现在专注于解决「选恰好 m 个子数组」的问题
        long long ans = 0;
        long long left = 0, right = pos_sum + 1;
        while (left + 1 < right) {
            long long k = left + (right - left) / 2;
            auto [res, cnt] = dp_without_limit(k);
            if (cnt <= m) {
                ans = res + m * k; // 见题解【细节 1】
                right = k;
            } else {
                left = k;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type pair struct{ f, cnt int } // DP 值, 子数组个数

// 相等的时候，子数组个数更大的劣
func less(a, b pair) bool {
	return a.f < b.f || a.f == b.f && a.cnt > b.cnt
}

func maximumSum(nums []int, m, l, r int) int64 {
	n := len(nums)
	s := make([]int, n+1) // nums 的前缀和
	posSum := 0 // nums 中的正数之和
	for i, x := range nums {
		s[i+1] = s[i] + x
		if x > 0 {
			posSum += x
		}
	}

	// 没有 m 约束，但每选一个子数组就要把元素和减少 k
	dpWithoutLimit := func(k int) pair {
		f := make([]pair, n+1)
		q := []int{}
		res := pair{math.MinInt, 0}

		for i := l; i <= n; i++ {
			// 1. 入
			j := i - l
			v := pair{f[j].f - s[j], f[j].cnt}
			for len(q) > 0 && less(pair{f[q[len(q)-1]].f - s[q[len(q)-1]], f[q[len(q)-1]].cnt}, v) {
				q = q[:len(q)-1]
			}
			q = append(q, j)

			// 2. 更新答案
			choose := pair{f[q[0]].f - s[q[0]] + s[i] - k, f[q[0]].cnt + 1}
			if less(res, choose) {
				// choose 保证我们至少选了一个子数组
				res = choose
			}

			// 更新 DP
			if less(f[i-1], choose) {
				f[i] = choose
			} else { // 不选
				f[i] = f[i-1]
			}

			// 3. 出，下一轮循环队首离开窗口
			if q[0] <= i-r {
				q = q[1:]
			}
		}

		return res
	}

	res0 := dpWithoutLimit(0)
	if res0.cnt <= m { // 直接满足题目要求
		return int64(res0.f)
	}

	// 现在专注于解决「选恰好 m 个子数组」的问题
	ans := 0
	sort.Search(posSum, func(k int) bool {
		k++
		res := dpWithoutLimit(k)
		if res.cnt <= m {
			ans = res.f + m*k // 见题解【细节 1】
			return true
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log S)$，其中 $n$ 是 $\textit{nums}$ 的长度，$S$ 是 $\textit{nums}$ 中的正数之和。二分 $\mathcal{O}(\log S)$ 次，每次跑一个 $\mathcal{O}(n)$ 的 DP。
- 空间复杂度：$\mathcal{O}(n)$。

## 变形题

本题至少要选一个子数组，如果改成至少要选 $\textit{least}$ 个子数组（$\textit{least}$ 是额外传入的一个非负整数），怎么做？

**答**：首先，跑一遍 $k=0$ 的 DP 算法（允许选零个子数组），设我们选了 $c$ 个子数组。

分类讨论：

- 如果 $\textit{least}\le c \le m$，那么直接算出了答案。
- 如果 $c > m$，那么和本题一样，转化成恰好选 $m$ 个子数组的问题。
- 如果 $c < \textit{least}$，则转化成恰好选 $\textit{least}$ 个子数组的问题。此时惩罚值 $k$ 必须是负数，从而让 DP 算法选出更多的子数组。闭区间二分下界是 $\textit{nums}$ 中的所有负数之和。

## 专题训练

见下面动态规划题单的「**§11.8 WQS 二分优化 DP**」。

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
