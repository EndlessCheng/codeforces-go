## 方法一：暴力枚举

枚举 $x$ 和 $y$。

把 $\textit{nums}[i]$ 变成满足 $\textit{nums}[i]\bmod k = x$ 的数，也就是变成一个与 $x$ 关于模 $k$ [同余](https://leetcode.cn/circle/discuss/mDfnkW/) 的数。

操作次数等价于：

- 在一个长为 $k$ 的环上，计算 $a = \textit{nums}[i]\bmod k$ 到 $x$ 的最短距离。

例如 $k=12$，类似生活中的圆形挂钟（闹钟），把小时从 $17$ 点（即 $5$ 点位置）拨到 $3$ 点位置，需要往回拨 $2$ 圈。

设 $d = |a - x|$，分类讨论：

- 不经过 $0$，直接从 $a$ 走到 $x$，需要 $d$ 步。
- 经过 $0$，绕一圈从 $a$ 走到 $x$，需要 $k-d$ 步。

取二者的最小值 $\min(d, k-d)$，即为 $a$ 到 $x$ 的最短距离。

[本题视频讲解](https://www.bilibili.com/video/BV1iuG76VEXy/?t=4m4s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int], k: int) -> int:
        ans = inf

        for x in range(k):
            for y in range(k):
                if y == x:
                    continue
                target = [x, y]
                s = 0
                for i, v in enumerate(nums):
                    d = abs(v % k - target[i % 2])
                    s += min(d, k - d)  # 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                ans = min(ans, s)

        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;

        for (int x = 0; x < k; x++) {
            for (int y = 0; y < k; y++) {
                if (y == x) {
                    continue;
                }
                int[] target = new int[]{x, y};
                int sum = 0;
                for (int i = 0; i < nums.length; i++) {
                    int d = Math.abs(nums[i] % k - target[i % 2]);
                    sum += Math.min(d, k - d); // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                }
                ans = Math.min(ans, sum);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int ans = INT_MAX;

        for (int x = 0; x < k; x++) {
            for (int y = 0; y < k; y++) {
                if (y == x) {
                    continue;
                }
                int target[2] = {x, y};
                int sum = 0;
                for (int i = 0; i < nums.size(); i++) {
                    int d = abs(nums[i] % k - target[i % 2]);
                    sum += min(d, k - d); // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                }
                ans = min(ans, sum);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
	ans := math.MaxInt
	for x := range k {
		for y := range k {
			if y == x {
				continue
			}
			target := [2]int{x, y}
			sum := 0
			for i, v := range nums {
				d := abs(v%k - target[i%2])
				sum += min(d, k-d) // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
			}
			ans = min(ans, sum)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：贪心 + 维护最小次小操作次数

如果不考虑 $x\ne y$ 的要求，偶数下标和奇数下标是互相独立的，可以分别计算。

以偶数为例，把偶数下标的数模 $k$ 后记在数组 $a$ 中。想象有一个长为 $k$ 的环，上面有一些人，坐标记在 $a$ 中。把这些人集中到环上的哪个位置，可以让总移动距离之和最小？

**核心思路**：如果 $x$ 是集中位置，那么在 $x$ 右手边的到 $x$ 距离在 $k/2$ 以内的数，从右边移动到 $x$；其余数从左边移动到 $x$。

**注**：这意味着，在 $x$ 对面的那段弧是没人经过的，那么断开这段弧，变成一个普通数组问题。根据 [中位数贪心](https://zhuanlan.zhihu.com/p/1922938031687595039)，可以都变成 $a$ 中的某个元素。

枚举 $a[i]$ 作为集中的位置，那么在 $[a[i], a[i]+k/2]$ 中的数，**减少**至 $a[i]$ 更好（相比增加）；在 $[a[i]+k/2 +1, a[i] + k]$ 中的数，**增加**至 $a[i]+k$ 更好（相比减少）。**注**：在模 $k$ 意义下，$x$ 和 $x+k$ 是同一个位置。

为方便计算，把 $a$ 复制一份（每个数都加 $k$），拼在 $a$ 的后面，得到数组 $b$。在 $b$
中二分查找 $\le a[i]+k/2$ 的最后一个数（或者 $\ge a[i]+k/2 +1$ 的第一个数），可以快速找到哪些数要减少，哪些数要增加。关于二分查找的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

然后利用前缀和，可以快速求出操作次数。计算方法见 [图解](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。

计算过程中，额外记录最小操作次数对应的 $a[i]$。

回到原问题，如何解决 $x = y$ 的情况？

我们可以额外求出偶数下标的**次小**操作次数，以及奇数下标的**次小**操作次数。如果 $x=y$，那么退而求其次，计算偶数下标的最小操作次数 + 奇数下标的次小操作次数，以及偶数下标的次小操作次数 +
奇数下标的最小操作次数，二者的最小值即为答案。

如何计算次小操作次数？这可能来自某个 $a[i]$ 的操作次数，也可能来自 $x-1$ 或 $x+1$，其中 $x$ 是最小操作次数对应的 $a[i]$。

```py [sol-Python3]
class Solution:
    def calc(self, a: list[int], k: int) -> tuple[int, int, int]:
        n = len(a)
        a.sort()
        a += [x + k for x in a]

        pre = list(accumulate(a, initial=0))

        # 都变成 target 的最小操作次数
        def calc_op(target: int) -> int:
            i = bisect_left(a, target, 0, n)
            j = bisect_right(a, target + k // 2, i, i + n)
            s1 = (pre[j] - pre[i]) - (j - i) * target  # [i, j) 中的数都减小到 target
            s2 = (n - j + i) * (target + k) - (pre[i + n] - pre[j])  # [j, i+n) 中的数都增大到 target+k
            return s1 + s2 

        mn = mn2 = inf
        best_x = 0

        for x in set(a[:n]):
            op = calc_op(x)
            # 维护最小次小操作次数
            if op < mn:
                mn2 = mn
                mn = op
                best_x = x
            elif op < mn2:
                mn2 = op

        # 还可以都变成 best_x-1 或者 best_x+1
        mn2 = min(mn2, calc_op((best_x - 1) % k), calc_op((best_x + 1) % k))

        return mn, mn2, best_x

    def minOperations(self, nums: list[int], k: int) -> int:
        if len(nums) == 1:
            return 0

        a = [[], []]
        for i, x in enumerate(nums):
            a[i % 2].append(x % k)

        min1x, min2x, best_x = self.calc(a[0], k)
        min1y, min2y, best_y = self.calc(a[1], k)

        if best_x != best_y:
            return min1x + min1y
        return min(min1x + min2y, min2x + min1y)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        if (nums.length == 1) {
            return 0;
        }

        List<Integer>[] a = new ArrayList[2];
        Arrays.setAll(a, _ -> new ArrayList<>());
        for (int i = 0; i < nums.length; i++) {
            a[i % 2].add(nums[i] % k);
        }

        int[] resX = calc(a[0], k);
        int[] resY = calc(a[1], k);

        int min1X = resX[0];
        int min2X = resX[1];
        int bestX = resX[2];

        int min1Y = resY[0];
        int min2Y = resY[1];
        int bestY = resY[2];

        if (bestX != bestY) {
            return min1X + min1Y;
        }
        return Math.min(min1X + min2Y, min2X + min1Y);
    }

    private int[] calc(List<Integer> a, int k) {
        int n = a.size();
        Collections.sort(a);

        int[] b = new int[n * 2];
        for (int i = 0; i < n; i++) {
            b[i] = a.get(i);
            b[n + i] = b[i] + k;
        }

        int[] sum = new int[n * 2 + 1];
        for (int i = 0; i < n * 2; i++) {
            sum[i + 1] = sum[i] + b[i];
        }

        int mn = Integer.MAX_VALUE;
        int mn2 = Integer.MAX_VALUE;
        int bestX = 0;

        for (int i = 0; i < n; i++) {
            int x = b[i];
            if (i > 0 && b[i] == b[i - 1]) { // 优化：相同的值无需重复计算
                continue;
            }

            int op = calcOp(b, sum, n, k, x);
            // 维护最小次小操作次数
            if (op < mn) {
                mn2 = mn;
                mn = op;
                bestX = x;
            } else if (op < mn2) {
                mn2 = op;
            }
        }

        // 还可以都变成 bestX-1 或者 bestX+1
        int op1 = calcOp(b, sum, n, k, (bestX - 1 + k) % k);
        int op2 = calcOp(b, sum, n, k, (bestX + 1) % k);
        mn2 = Math.min(mn2, Math.min(op1, op2));

        return new int[]{mn, mn2, bestX};
    }

    // 都变成 target 的最小操作次数
    private int calcOp(int[] a, int[] sum, int n, int k, int target) {
        int i = lowerBound(a, -1, n, target);
        int j = lowerBound(a, i - 1, i + n, target + k / 2 + 1);
        return (sum[j] - sum[i]) - (j - i) * target + // [i, j) 中的数都减小到 target
               (n - j + i) * (target + k) - (sum[i + n] - sum[j]); // [j, i+n) 中的数都增大到 target+k
    }

    private int lowerBound(int[] a, int left, int right, int target) {
        while (left + 1 < right) { // 开区间不为空
            int mid = (left + right) >>> 1; // 比 /2 快
            if (a[mid] >= target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
    tuple<int, int, int> calc(vector<int>& a, int k) {
        int n = a.size();
        ranges::sort(a);

        a.resize(n * 2);
        for (int i = 0; i < n; i++) {
            a[i + n] = a[i] + k;
        }

        vector<int> sum(n * 2 + 1);
        partial_sum(a.begin(), a.end(), sum.begin() + 1);

        // 都变成 target 的最小操作次数
        auto calc_op = [&](int target) -> int {
            int i = lower_bound(a.begin(), a.begin() + n, target) - a.begin();
            int j = upper_bound(a.begin() + i, a.begin() + i + n, target + k / 2) - a.begin();
            return (sum[j] - sum[i]) - (j - i) * target + // [i, j) 中的数都减小到 target
                   (n - j + i) * (target + k) - (sum[i + n] - sum[j]); // [j, i+n) 中的数都增大到 target+k
        };

        int mn = INT_MAX, mn2 = INT_MAX, best_x = 0;

        for (int i = 0; i < n; i++) {
            int x = a[i];
            if (i > 0 && a[i] == a[i - 1]) { // 优化：相同的值无需重复计算
                continue;
            }

            int op = calc_op(x);
            // 维护最小次小操作次数
            if (op < mn) {
                mn2 = mn;
                mn = op;
                best_x = x;
            } else if (op < mn2) {
                mn2 = op;
            }
        }

        // 还可以都变成 best_x-1 或者 best_x+1
        mn2 = min({mn2, calc_op((best_x - 1 + k) % k), calc_op((best_x + 1) % k)});

        return {mn, mn2, best_x};
    }

public:
    int minOperations(vector<int>& nums, int k) {
        if (nums.size() == 1) {
            return 0;
        }

        vector<int> a[2];
        for (int i = 0; i < nums.size(); i++) {
            a[i % 2].push_back(nums[i] % k);
        }

        auto [min1_x, min2_x, best_x] = calc(a[0], k);
        auto [min1_y, min2_y, best_y] = calc(a[1], k);

        if (best_x != best_y) {
            return min1_x + min1_y;
        }
        return min(min1_x + min2_y, min2_x + min1_y);
    }
};
```

```go [sol-Go]
func calc(a []int, k int) (int, int, int) {
	n := len(a)
	slices.Sort(a)
	for _, x := range a {
		a = append(a, x+k)
	}

	sum := make([]int, n*2+1)
	for i, x := range a {
		sum[i+1] = sum[i] + x
	}

	// 都变成 target 的最小操作次数
	calcOp := func(target int) int {
		i := sort.SearchInts(a[:n], target)
		j := i + sort.SearchInts(a[i:i+n], target+k/2+1)
		return (sum[j] - sum[i]) - (j-i)*target + // [i, j) 中的数都减小到 target
			(n-j+i)*(target+k) - (sum[i+n] - sum[j]) // [j, i+n) 中的数都增大到 target+k
	}

	mn, mn2, bestX := math.MaxInt, math.MaxInt, 0
	for i, x := range a[:n] {
		if i > 0 && a[i] == a[i-1] { // 优化：相同的值无需重复计算
			continue
		}
		op := calcOp(x)
		// 维护最小次小操作次数
		if op < mn {
			mn2 = mn
			mn, bestX = op, x
		} else if op < mn2 {
			mn2 = op
		}
	}

	// 还可以都变成 bestX-1 或者 bestX+1
	mn2 = min(mn2, calcOp((bestX-1+k)%k), calcOp((bestX+1)%k))

	return mn, mn2, bestX
}

func minOperations(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	a := [2][]int{}
	for i, x := range nums {
		a[i%2] = append(a[i%2], x%k)
	}

	min1x, min2x, bestX := calc(a[0], k)
	min1y, min2y, bestY := calc(a[1], k)

	if bestX != bestY {
		return min1x + min1y
	}
	return min(min1x+min2y, min2x+min1y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面贪心题单的「**§4.5 中位数贪心**」。

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
