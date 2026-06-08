对于 DP 问题，思考「**最后一步发生了什么**」，可以帮助我们找到状态定义和状态转移方程。

- 如果没有选 $\textit{nums}[n-1]$，问题变成在前 $n-1$ 个数中选子数组。
- 如果选了 $\textit{nums}[n-1]$，还需要枚举子数组的左端点 $L$，问题变成在前 $L$ 个数（下标 $[0,L-1]$）中选子数组。

和划分型 DP 一样，定义 $f[i][j]$ 表示在 $\textit{nums}$ 的前 $j$ 个数（下标 $[0,j-1]$）中选出**恰好** $i$ 个连续子数组，所选元素之和的最大值。枚举恰好选 $i=1,2,\ldots,m$ 个子数组，答案就是 $f[i][n]$ 的最大值。

> 为什么把状态定义成「恰好」而不是「至多」？这样方便处理「至少一个」的要求。想一想，如果改成「至少两个」「至少三个」，哪种定义更好？

分类讨论：

- 如果没有选 $\textit{nums}[j-1]$，问题变成在 $\textit{nums}$ 的前 $j-1$ 个数中选出恰好 $i$ 个连续子数组，所选元素之和的最大值，即 $f[i][j-1]$。
- 如果选了 $\textit{nums}[j-1]$，还需要枚举子数组的左端点 $L = j-r,j-r+1,\ldots,j-\ell$，问题变成在 $\textit{nums}$ 的前 $L$ 个数中选出恰好 $i-1$ 个连续子数组，所选元素之和的最大值，即 $f[i-1][L]$，加上子数组 $[L,j-1]$ 的元素和。

设 $\textit{nums}$ 的**前缀和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。子数组 $[L,j-1]$ 的元素和可以表示为 $s[j] - s[L]$。

所有情况取最大值，得到状态转移方程

$$
f[i][j] = \max\left\{f[i][j-1], \max_{L=j-r}^{j-\ell} f[i-1][L] + s[j] - s[L] \right\}
$$

初始值：

- $f[0][j] = 0$。选 $0$ 个子数组，元素和为 $0$。
- $f[i][0] = -\infty\ (i \ge 1)$。无法从 $0$ 个数中选出 $i$ 个子数组，不合法。设置为 $-\infty$，这样计算 $\max$ 不会从不合法的状态转移过来。

枚举恰好选 $i=1,2,\ldots,m$ 个子数组，答案为 $\max\limits_{i=1}^{m} f[i][n]$。

这样做的时间复杂度为 $\mathcal{O}(nm(r-\ell))$，太慢了。

## 单调队列优化

式子变形，把和 $L$ 无关的 $s[j]$ 提到外面：

$$
\max_{L=j-r}^{j-\ell} f[i-1][L] + s[j] - s[L] = s[j] + \max_{L=j-r}^{j-\ell} f[i-1][L] - s[L]
$$

定义 $d_{i-1}[L] = f[i-1][L] - s[L]$，那么

$$
\max_{L=j-r}^{j-\ell} f[i-1][L] - s[L] = \max_{L=j-r}^{j-\ell} d_{i-1}[L]
$$

由于 $j$ 每增加 $1$，定长窗口 $[j-r,j-\ell]$ 就向右滑动 $1$，所以这是一个标准的 [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) 问题，请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

用单调队列优化后，就可以 $\mathcal{O}(1)$ 计算转移了。

[本题视频讲解](https://www.bilibili.com/video/BV1yfEx6YEBx/?t=9m10s) 详细介绍了如何思考这类问题，欢迎点赞关注~

## 写法一：二维 DP 数组

```py [sol-Python3]
class Solution:
    def maximumSum(self, nums: List[int], m: int, left: int, right: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))  # nums 的前缀和

        # f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
        f = [[-inf] * (n + 1) for _ in range(m + 1)]
        f[0] = [0] * (n + 1)

        for i in range(1, m + 1):
            q = deque()

            # 前 i 个子数组至少占用了 i * left 个位置
            for j in range(i * left, n + 1):
                # 1. 入
                k = j - left
                v = f[i - 1][k] - s[k]
                while q and f[i - 1][q[-1]] - s[q[-1]] <= v:
                    q.pop()
                q.append(k)

                # 2. 更新
                # 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                f[i][j] = max(f[i][j - 1], f[i - 1][q[0]] - s[q[0]] + s[j])

                # 3. 出，下一轮循环队首离开窗口
                if q[0] <= j - right:
                    q.popleft()

        # 枚举恰好选 i 个子数组
        return max(f[i][n] for i in range(1, m + 1))
```

```java [sol-Java]
class Solution {
    public long maximumSum(int[] nums, int m, int left, int right) {
        int n = nums.length;
        long[] s = new long[n + 1]; // nums 的前缀和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        // f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
        long[][] f = new long[m + 1][n + 1];
        for (int i = 1; i <= m; i++) {
            Arrays.fill(f[i], Long.MIN_VALUE / 2); // 防止溢出
        }
        long ans = Long.MIN_VALUE;

        for (int i = 1; i <= m; i++) {
            Deque<Integer> q = new ArrayDeque<>();

            // 前 i 个子数组至少占用了 i * left 个位置
            for (int j = i * left; j <= n; j++) {
                // 1. 入
                int k = j - left;
                long v = f[i - 1][k] - s[k];
                while (!q.isEmpty() && f[i - 1][q.peekLast()] - s[q.peekLast()] <= v) {
                    q.pollLast();
                }
                q.offerLast(k);

                // 2. 更新
                // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                f[i][j] = Math.max(f[i][j - 1], f[i - 1][q.peekFirst()] - s[q.peekFirst()] + s[j]);

                // 3. 出，下一轮循环队首离开窗口
                if (q.peekFirst() <= j - right) {
                    q.pollFirst();
                }
            }

            // 枚举恰好选 i 个子数组
            ans = Math.max(ans, f[i][n]);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSum(vector<int>& nums, int m, int left, int right) {
        int n = nums.size();
        vector<long long> s(n + 1); // nums 的前缀和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        // f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
        vector f(m + 1, vector<long long>(n + 1, LLONG_MIN / 2)); // 防止溢出
        ranges::fill(f[0], 0);
        long long ans = LLONG_MIN;

        for (int i = 1; i <= m; i++) {
            deque<int> q;

            // 前 i 个子数组至少占用了 i * left 个位置
            for (int j = i * left; j <= n; j++) {
                // 1. 入
                int k = j - left;
                long long v = f[i - 1][k] - s[k];
                while (!q.empty() && f[i - 1][q.back()] - s[q.back()] <= v) {
                    q.pop_back();
                }
                q.push_back(k);

                // 2. 更新
                // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                f[i][j] = max(f[i][j - 1], f[i - 1][q.front()] - s[q.front()] + s[j]);

                // 3. 出，下一轮循环队首离开窗口
                if (q.front() <= j - right) {
                    q.pop_front();
                }
            }

            // 枚举恰好选 i 个子数组
            ans = max(ans, f[i][n]);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maximumSum(nums []int, m, left, right int) int64 {
	n := len(nums)
	s := make([]int, n+1) // nums 的前缀和
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	// f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	ans := math.MinInt

	for i := 1; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MinInt / 2
		}
		q := []int{}

		// 前 i 个子数组至少占用了 i * left 个位置
		for j := i * left; j <= n; j++ {
			// 1. 入
			k := j - left
			v := f[i-1][k] - s[k]
			for len(q) > 0 && f[i-1][q[len(q)-1]]-s[q[len(q)-1]] <= v {
				q = q[:len(q)-1]
			}
			q = append(q, k)

			// 2. 更新
			// 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
			f[i][j] = max(f[i][j-1], f[i-1][q[0]]-s[q[0]]+s[j])

			// 3. 出，下一轮循环队首离开窗口
			if q[0] <= j-right {
				q = q[1:]
			}
		}

		// 枚举恰好选 i 个子数组
		ans = max(ans, f[i][n])
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个三重循环，但对于内层的两重循环，每个下标入队出队各至多一次，所以内层的两重循环是 $\mathcal{O}(n)$ 时间的。
- 空间复杂度：$\mathcal{O}(nm)$。

## 写法二：一维 DP 数组

```py [sol-Python3]
class Solution:
    def maximumSum(self, nums: List[int], m: int, left: int, right: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))  # nums 的前缀和
        f = [0] * (n + 1)
        ans = -inf

        for i in range(1, m + 1):
            nf = [-inf] * (n + 1)
            q = deque()

            # 前 i 个子数组至少占用了 i * left 个位置
            for j in range(i * left, n + 1):
                # 1. 入
                k = j - left
                v = f[k] - s[k]
                while q and f[q[-1]] - s[q[-1]] <= v:
                    q.pop()
                q.append(k)

                # 2. 更新
                # 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                nf[j] = max(nf[j - 1], f[q[0]] - s[q[0]] + s[j])

                # 3. 出，下一轮循环队首离开窗口
                if q[0] <= j - right:
                    q.popleft()

            f = nf
            ans = max(ans, f[n])

        return ans
```

```java [sol-Java]
class Solution {
    public long maximumSum(int[] nums, int m, int left, int right) {
        int n = nums.length;
        long[] s = new long[n + 1]; // nums 的前缀和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        long[] f = new long[n + 1];
        long ans = Long.MIN_VALUE;

        for (int i = 1; i <= m; i++) {
            long[] nf = new long[n + 1];
            Arrays.fill(nf, Long.MIN_VALUE / 2);
            Deque<Integer> q = new ArrayDeque<>();

            // 前 i 个子数组至少占用了 i * left 个位置
            for (int j = i * left; j <= n; j++) {
                // 1. 入
                int k = j - left;
                long v = f[k] - s[k];
                while (!q.isEmpty() && f[q.peekLast()] - s[q.peekLast()] <= v) {
                    q.pollLast();
                }
                q.offerLast(k);

                // 2. 更新
                // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                nf[j] = Math.max(nf[j - 1], f[q.peekFirst()] - s[q.peekFirst()] + s[j]);

                // 3. 出，下一轮循环队首离开窗口
                if (q.peekFirst() <= j - right) {
                    q.pollFirst();
                }
            }

            // 枚举恰好选 i 个子数组
            f = nf;
            ans = Math.max(ans, f[n]);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSum(vector<int>& nums, int m, int left, int right) {
        int n = nums.size();
        vector<long long> s(n + 1); // nums 的前缀和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        vector<long long> f(n + 1);
        long long ans = LLONG_MIN;

        for (int i = 1; i <= m; i++) {
            vector<long long> nf(n + 1, LLONG_MIN / 2);
            deque<int> q;

            // 前 i 个子数组至少占用了 i * left 个位置
            for (int j = i * left; j <= n; j++) {
                // 1. 入
                int k = j - left;
                long long v = f[k] - s[k];
                while (!q.empty() && f[q.back()] - s[q.back()] <= v) {
                    q.pop_back();
                }
                q.push_back(k);

                // 2. 更新
                // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
                nf[j] = max(nf[j - 1], f[q.front()] - s[q.front()] + s[j]);

                // 3. 出，下一轮循环队首离开窗口
                if (q.front() <= j - right) {
                    q.pop_front();
                }
            }

            // 枚举恰好选 i 个子数组
            f = move(nf);
            ans = max(ans, f[n]);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maximumSum(nums []int, m, left, right int) int64 {
	n := len(nums)
	s := make([]int, n+1) // nums 的前缀和
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	// f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
	f := make([]int, n+1)
	ans := math.MinInt

	for i := 1; i <= m; i++ {
		nf := make([]int, n+1)
		for j := range nf {
			nf[j] = math.MinInt / 2
		}
		q := []int{}

		// 前 i 个子数组至少占用了 i * left 个位置
		for j := i * left; j <= n; j++ {
			// 1. 入
			k := j - left
			v := f[k] - s[k]
			for len(q) > 0 && f[q[len(q)-1]]-s[q[len(q)-1]] <= v {
				q = q[:len(q)-1]
			}
			q = append(q, k)

			// 2. 更新
			// 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
			nf[j] = max(nf[j-1], f[q[0]]-s[q[0]]+s[j])

			// 3. 出，下一轮循环队首离开窗口
			if q[0] <= j-right {
				q = q[1:]
			}
		}

		// 枚举恰好选 i 个子数组
		f = nf
		ans = max(ans, f[n])
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个三重循环，但对于内层的两重循环，每个下标入队出队各至多一次，所以内层的两重循环是 $\mathcal{O}(n)$ 时间的。
- 空间复杂度：$\mathcal{O}(n)$。

## 更快的做法：WQS 二分

见下一题 [3957. 非重叠子数组最大和 II](https://leetcode.cn/problems/maximum-sum-of-m-non-overlapping-subarrays-ii/)，[我的题解](https://leetcode.cn/problems/maximum-sum-of-m-non-overlapping-subarrays-ii/solutions/3980778/la-ge-lang-ri-song-chi-wqs-er-fen-python-m2iw/)。

**注**：利用下一题的结论，对于本题的做法，可以在 $f[i][n]\le \textit{ans}$ 时，提前返回答案。

## 专题训练

见下面动态规划题单的「**§5.3 约束划分个数**」和「**§11.3 单调队列优化 DP**」。

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
