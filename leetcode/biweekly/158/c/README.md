## 方法一：暴力枚举

枚举子数组的右端点 $i$，然后枚举子数组的左端点 $j=i,i-1,i-2,\ldots,0$。

在枚举的过程中，维护子数组 $[j,i]$ 的 GCD，记作 $g$。

由于一个数至多可以乘以 $2$ 一次，所以至多可以把 $g$ 变成 $2g$（比如子数组中的每个数都乘以 $2$）。

但没必要把每个数都乘以 $2$。比如子数组为 $[8,4,4,16]$，$g=4$，只需把中间两个 $4$ 乘以 $2$，就可以让子数组 GCD 从 $g=4$ 变成 $2g=8$。

所以关键在于，统计子数组中每个数的因子 $2$ 的个数的最小值，以及这个最小值的出现次数。如果最小值的出现次数 $\le k$，那么可以把 $g$ 变成 $2g$。

代码实现时，不需要写循环统计因子 $2$ 的个数，也不需要调用库函数统计尾零的个数，直接用 $\textit{nums}[j]$ 的 $\text{lowbit}$ 代替，因为 $\text{lowbit}$ 相当于每个数因子 $2$ 的乘积，可以改成求 $\text{lowbit}$ 的小值及其出现次数。

什么是 $\text{lowbit}$？请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

[本题视频讲解](https://www.bilibili.com/video/BV1rET9zsEsB/?t=11m24s)，欢迎点赞关注~

```py [sol-Python3]
max = lambda a, b: b if b > a else a  # 手写 max 更快

class Solution:
    def maxGCDScore(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(len(nums)):
            lowbit_min = inf
            lowbit_cnt = g = 0
            for j in range(i, -1, -1):
                x = nums[j]
                lb = x & -x
                if lb < lowbit_min:
                    lowbit_min, lowbit_cnt = lb, 1
                elif lb == lowbit_min:
                    lowbit_cnt += 1

                g = gcd(g, x)
                new_g = g * 2 if lowbit_cnt <= k else g
                ans = max(ans, new_g * (i - j + 1))
        return ans
```

```java [sol-Java]
class Solution {
    public long maxGCDScore(int[] nums, int k) {
        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int lowbitMin = Integer.MAX_VALUE;
            int lowbitCnt = 0;
            int g = 0;
            for (int j = i; j >= 0; j--) {
                int x = nums[j];
                int lb = x & -x;
                if (lb < lowbitMin) {
                    lowbitMin = lb;
                    lowbitCnt = 1;
                } else if (lb == lowbitMin) {
                    lowbitCnt++;
                }

                g = gcd(g, x);
                int newG = lowbitCnt <= k ? g * 2 : g;
                ans = Math.max(ans, (long) newG * (i - j + 1));
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxGCDScore(vector<int>& nums, int k) {
        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int lowbit_min = INT_MAX;
            int lowbit_cnt = 0;
            int g = 0;
            for (int j = i; j >= 0; j--) {
                int x = nums[j];
                int lb = x & -x;
                if (lb < lowbit_min) {
                    lowbit_min = lb;
                    lowbit_cnt = 1;
                } else if (lb == lowbit_min) {
                    lowbit_cnt++;
                }

                g = gcd(g, x);
                int new_g = lowbit_cnt <= k ? g * 2 : g;
                ans = max(ans, 1LL * new_g * (i - j + 1));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxGCDScore(nums []int, k int) int64 {
	ans := 0
	for i := range nums {
		lowbitMin, lowbitCnt := math.MaxInt, 0
		g := 0
		for j := i; j >= 0; j-- {
			x := nums[j]
			lb := x & -x
			if lb < lowbitMin {
				lowbitMin, lowbitCnt = lb, 1
			} else if lb == lowbitMin {
				lowbitCnt++
			}

			g = gcd(g, x)
			newG := g
			if lowbitCnt <= k {
				newG *= 2
			}
			ans = max(ans, newG*(i-j+1))
		}
	}
	return int64(ans)
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+\log U))$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。注意这里是加不是乘，因为 GCD 在循环中不会增加只会减少，每次要么不变，要么至少减半，减半次数就是 GCD 的循环次数之和，即 $\mathcal{O}(\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：logTrick

**前置知识**：[讲解（方法二）](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)。

根据 logTrick，固定右端点 $i$ 时，只有 $\mathcal{O}(\log U)$ 种不同的 GCD，其中 $U=\max(\textit{nums})$。并且，相同的 GCD 对应的左端点是连续的，也就是说左端点组成了一个区间。

设左端点 $j$ 在左开右闭区间 $(l,r]$ 中时，GCD 等于 $g$。

首先不操作，用 $g(i-l)$ 更新答案的最大值。

然后考虑能否通过操作让 $g$ 变成 $2g$。

区间中的最小的 $\text{lowbit}$，其实就是 $g$ 的 $\text{lowbit}$。

考虑 $[0,i]$ 中的 $\textit{nums}[i]$，设 $\textit{pos}$ 为 $\text{lowbit}$ 等于 $g$ 的 $\text{lowbit}$ 的 $\textit{nums}[i]$ 的下标列表。

考虑列表中的倒数第 $k+1$ 个下标 $i'$（如果不存在记作 $-1$），左端点最小值必须大于 $i'$，否则就没有足够的操作次数了。

因此，在操作 $k$ 次的情况下，左端点的最小值为

$$
\textit{minL} = \max(l+1, i'+1)
$$

> 注意 $(l,r]$ 是左开右闭区间。

如果 $\textit{minL}<r$，那么可以操作让 $g$ 变成 $2g$，用 $2g(i-\textit{minL})$ 更新答案的最大值。

```py [sol-Python3]
class Solution:
    def maxGCDScore(self, nums: List[int], k: int) -> int:
        lowbit_pos = defaultdict(list)

        ans = 0
        intervals = []  # 每个元素是一个三元组 (g, l, r)，表示区间 (l, r] 的 GCD 为 g
        for i, x in enumerate(nums):
            lowbit_pos[x & -x].append(i)

            # 更新已有区间的 GCD
            for p in intervals:
                p[0] = gcd(p[0], x)
            # 添加新元素作为新区间
            intervals.append([x, i - 1, i])

            # 去重（合并 g 相同的区间）
            idx = 1
            for j in range(1, len(intervals)):
                if intervals[j][0] != intervals[j - 1][0]:
                    intervals[idx] = intervals[j]
                    idx += 1
                else:
                    intervals[idx - 1][2] = intervals[j][2]
            del intervals[idx:]

            # 此时我们将区间 [0,i] 划分成了 len(intervals) 个左开右闭区间
            # 对于 intervals 中的 (l,r]，对于任意 j∈(l,r]，gcd(区间[j,i]) 的计算结果均为 g
            for g, l, r in intervals:
                # 不做任何操作
                ans = max(ans, g * (i - l))
                # 看看能否乘 2
                pos = lowbit_pos[g & -g]
                min_l = max(l, pos[-k - 1]) if len(pos) > k else l
                if min_l < r:  # 可以乘 2
                    ans = max(ans, g * 2 * (i - min_l))

        return ans
```

```java [sol-Java]
class Solution {
    public long maxGCDScore(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        int width = 32 - Integer.numberOfLeadingZeros(mx);

        List<Integer>[] lowbitPos = new List[width];
        Arrays.setAll(lowbitPos, i -> new ArrayList<>());

        int[][] intervals = new int[width + 1][3];
        int size = 0;

        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            int tz = Integer.numberOfTrailingZeros(x);
            lowbitPos[tz].add(i); // 用 tz 代替 x 的 lowbit

            // 更新每个区间的 gcd
            for (int j = 0; j < size; j++) {
                intervals[j][0] = gcd(intervals[j][0], x);
            }
            intervals[size][0] = x;
            intervals[size][1] = i - 1;
            intervals[size][2] = i;
            size++;

            // 合并 g 相同的区间（去重）
            int idx = 1;
            for (int j = 1; j < size; j++) {
                if (intervals[j][0] != intervals[j - 1][0]) {
                    intervals[idx][0] = intervals[j][0];
                    intervals[idx][1] = intervals[j][1];
                    intervals[idx][2] = intervals[j][2];
                    idx++;
                } else {
                    intervals[idx - 1][2] = intervals[j][2];
                }
            }
            size = idx;

            // 此时我们将区间 [0,i] 划分成了 size 个左开右闭区间
            // 对于 intervals 中的 (l,r]，对于任意 j∈(l,r]，gcd(区间[j,i]) 的计算结果均为 g
            for (int j = 0; j < size; j++) {
                int g = intervals[j][0];
                int l = intervals[j][1];
                int r = intervals[j][2];

                // 不做任何操作
                ans = Math.max(ans, (long) g * (i - l));

                // 看看能否乘 2
                List<Integer> pos = lowbitPos[Integer.numberOfTrailingZeros(g)];
                int minL = pos.size() > k ? Math.max(l, pos.get(pos.size() - k - 1)) : l;
                if (minL < r) { // 可以乘 2
                    ans = Math.max(ans, (long) g * 2 * (i - minL));
                }
            }
        }

        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxGCDScore(vector<int>& nums, int k) {
        int n = nums.size();
        int mx = bit_width((uint32_t) ranges::max(nums));
        vector<vector<int>> lowbit_pos(mx);

        struct Interval { int g, l, r; };  // 左开右闭 (l, r]
        vector<Interval> intervals;

        long long ans = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int tz = countr_zero((uint32_t) x);
            lowbit_pos[tz].push_back(i); // 用 tz 代替 x 的 lowbit

            for (auto& p : intervals) {
                p.g = gcd(p.g, x);
            }
            intervals.emplace_back(x, i - 1, i);

            // 去重（合并 g 相同的区间）
            int idx = 1;
            for (int j = 1; j < intervals.size(); j++) {
                if (intervals[j].g != intervals[j - 1].g) {
                    intervals[idx++] = intervals[j];
                } else {
                    intervals[idx - 1].r = intervals[j].r;
                }
            }
            intervals.resize(idx);

            // 此时我们将区间 [0,i] 划分成了 len(intervals) 个左开右闭区间
            // 对于 intervals 中的 (l,r]，对于任意 j∈(l,r]，gcd(区间[j,i]) 的计算结果均为 g
            for (auto& [g, l, r] : intervals) {
                // 不做任何操作
                ans = max(ans, 1LL * g * (i - l));

                // 看看能否乘 2
                int tz = countr_zero((uint32_t) g);
                auto& pos = lowbit_pos[tz];
                int min_l = pos.size() > k ? max(l, pos[pos.size() - k - 1]) : l;
                if (min_l < r) { // 可以乘 2
                    ans = max(ans, 1LL * g * 2 * (i - min_l));
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxGCDScore(nums []int, k int) int64 {
	mx := bits.Len(uint(slices.Max(nums)))
	lowbitPos := make([][]int, mx)

	ans := 0
	type interval struct{ g, l, r int } // 左开右闭 (l,r]
	intervals := []interval{}
	for i, x := range nums {
		tz := bits.TrailingZeros(uint(x))
		lowbitPos[tz] = append(lowbitPos[tz], i) // 用 tz 代替 x 的 lowbit

		for j, p := range intervals {
			intervals[j].g = gcd(p.g, x)
		}
		intervals = append(intervals, interval{x, i - 1, i})

		// 去重（合并 g 相同的区间）
		idx := 1
		for j := 1; j < len(intervals); j++ {
			if intervals[j].g != intervals[j-1].g {
				intervals[idx] = intervals[j]
				idx++
			} else {
				intervals[idx-1].r = intervals[j].r
			}
		}
		intervals = intervals[:idx]

		// 此时我们将区间 [0,i] 划分成了 len(intervals) 个左闭右开区间
		// 对于任意 p∈intervals，任意 j∈(p.l,p.r]，gcd(区间[j,i]) 的计算结果均为 p.g
		for _, p := range intervals {
			// 不做任何操作
			ans = max(ans, p.g*(i-p.l))
			// 看看能否乘 2
			tz := bits.TrailingZeros(uint(p.g))
			pos := lowbitPos[tz]
			minL := p.l
			if len(pos) > k {
				minL = max(minL, pos[len(pos)-k-1])
			}
			if minL < p.r { // 可以乘 2
				ans = max(ans, p.g*2*(i-minL))
			}
		}
	}
	return int64(ans)
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。外层循环每次会增加一个区间，这个区间在整个算法过程中，要么合并到其他区间中（消失啦），要么其 GCD 一共减少 $\mathcal{O}(\log U)$ 次，所以每个区间的 GCD 的计算过程会贡献 $\mathcal{O}(\log U)$ 个循环次数，所以**总的循环次数**是 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(n+\log U)$。**注**：位置列表改成队列的话可以做到 $\mathcal{O}(\min(n+\log U, k\log U))$。

更多相似题目，见下面位运算题单的「**LogTrick**」。

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
