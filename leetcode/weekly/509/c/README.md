由于题目要求 $k>1$，我们先特判 $\textit{nums}$ 只包含 $1$ 的情况，此时最优解是只选一个 $1$，分数差为 $-1$，最小 $k$ 为 $2$。

否则，选 $\textit{nums}[i]$ 的因子作为 $k$ 是最优的，直接枚举。我们可以先收集所有因子，然后去重，再枚举因子。这样可以避免重复计算。

对于一个固定的 $k$，设 $x=\textit{nums}[i]$，如果 $x$ 不是 $k$ 的倍数，则视作 $-x$（减去 Bob 的分数）。问题变成 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，请看 [我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 预处理每个数的因子
MX = 1_000_001
divisors = [[] for _ in range(MX)]
for i in range(2, MX):  # 本题 k > 1
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子


class Solution:
    # 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    def maxSubArray(self, nums: list[int], k: int) -> int:
        ans = -inf
        f = 0
        for x in nums:
            f = max(f, 0) + (-x if x % k else x)
            ans = max(ans, f)
        return ans

    def divisibleGame(self, nums: list[int]) -> int:
        MOD = 1_000_000_007
        # 收集所有因子
        all_divisors = []
        for x in nums:
            all_divisors += divisors[x]

        if not all_divisors:
            # 每个数都是 1
            # 最优是只选一个 1（分数差为 -1），最小 k 为 2
            return MOD - 2

        # 排序去重
        all_divisors = sorted(set(all_divisors))

        max_diff = -inf
        best_k = 0
        # 枚举因子作为 k，计算最大子数组和
        for k in all_divisors:
            diff = self.maxSubArray(nums, k)
            if diff > max_diff:
                max_diff = diff
                best_k = k

        return max_diff * best_k % MOD
```

```java [sol-Java]
class Solution {
    // 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    private int maxSubArray(int[] nums, int k) {
        int ans = Integer.MIN_VALUE;
        int f = 0;
        for (int x : nums) {
            f = Math.max(f, 0) + (x % k == 0 ? x : -x);
            ans = Math.max(ans, f);
        }
        return ans;
    }

    public int divisibleGame(int[] nums) {
        final int MOD = 1_000_000_007;

        // 收集所有因子
        // 预处理超时了，改成不预处理的写法
        List<Integer> allDivisors = new ArrayList<>();
        for (int x : nums) {
            for (int d = 2; d * d <= x; d++) {
                if (x % d == 0) {
                    allDivisors.add(d);
                    if (d != x / d) {
                        allDivisors.add(x / d);
                    }
                }
            }
            if (x > 1) {
                allDivisors.add(x);
            }
        }

        if (allDivisors.isEmpty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2
            return MOD - 2;
        }

        Collections.sort(allDivisors);

        int maxDiff = Integer.MIN_VALUE;
        int bestK = 0;
        int preK = 0;
        for (int k : allDivisors) {
            if (k == preK) {
                continue;
            }
            int diff = maxSubArray(nums, k);
            if (diff > maxDiff) {
                maxDiff = diff;
                bestK = k;
            }
            preK = k;
        }

        // 保证结果非负
        return (int) (((long) maxDiff * bestK % MOD + MOD) % MOD);
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'001;
vector<int> divisors[MX];

int init = [] {
    // 本题 k > 1
    for (int i = 2; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

class Solution {
    // 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
    int maxSubArray(vector<int>& nums, int k) {
        int ans = INT_MIN;
        int f = 0;
        for (int x : nums) {
            f = max(f, 0) + (x % k ? -x : x);
            ans = max(ans, f);
        }
        return ans;
    }

public:
    int divisibleGame(vector<int>& nums) {
        constexpr int MOD = 1'000'000'007;

        // 收集所有因子
        vector<int> all_divisors;
        for (int x : nums) {
            all_divisors.insert(all_divisors.end(), divisors[x].begin(), divisors[x].end());
        }

        if (all_divisors.empty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2
            return MOD - 2;
        }

        // 排序去重
        ranges::sort(all_divisors);
        all_divisors.erase(ranges::unique(all_divisors).begin(), all_divisors.end());

        int max_diff = INT_MIN;
        int best_k = 0;
        // 枚举因子作为 k，计算最大子数组和
        for (int k : all_divisors) {
            int diff = maxSubArray(nums, k);
            if (diff > max_diff) {
                max_diff = diff;
                best_k = k;
            }
        }

        // 保证结果非负
        return (1LL * max_diff * best_k % MOD + MOD) % MOD;
    }
};
```

```go [sol-Go]
const mx = 1_000_001
var divisors [mx][]int32

func init() {
	// 本题 k > 1
	for i := int32(2); i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

// 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
func maxSubArray(nums []int, k int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		if x%k != 0 {
			x = -x
		}
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007
	// 收集所有因子
	allDivisors := []int32{}
	for _, x := range nums {
		allDivisors = append(allDivisors, divisors[x]...)
	}

	if len(allDivisors) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	// 排序去重
	slices.Sort(allDivisors)
	allDivisors = slices.Compact(allDivisors)

	maxDiff, bestK := math.MinInt, 0
	// 枚举因子作为 k，计算最大子数组和
	for _, d := range allDivisors {
		k := int(d)
		diff := maxSubArray(nums, k)
		if diff > maxDiff {
			maxDiff, bestK = diff, k
		}
	}

	// 保证结果非负
	return (maxDiff*bestK%mod + mod) % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(Dn)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D$ 是 $\textit{nums}$ 中的不同因子个数。由于我们去重了，最坏情况是 $\textit{nums}$ 的所有元素互不相同。虽然一个数的因子个数可能很多，但在 $n$ 个不同元素的情况下，平均每个数有 $\mathcal{O}(\log U)$ 个因子（$U=\max(\textit{nums})$），所以 $D = \mathcal{O}(n\log U)$。这是一个比较松的估计，这里面有相同的因子，实际不同因子个数可能更少。
- 空间复杂度：$\mathcal{O}(D)$。

## 专题训练

1. 数学题单的「**§1.5 因子**」。
2. 动态规划题单的「**§1.3 最大子数组和**」。

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
