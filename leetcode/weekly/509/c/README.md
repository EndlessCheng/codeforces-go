由于题目要求 $k>1$，我们先特判 $\textit{nums}$ 只包含 $1$ 的情况，此时最优解是只选一个 $1$，分数差为 $-1$，最小 $k$ 为 $2$（见示例 3）。

否则，选 $\textit{nums}[i]$ 的因子作为 $k$ 更好（否则分数差一定是负数）。我们可以先收集所有因子，去重，再枚举因子，这样可以避免重复计算。

进一步地，如果一个数是 $x$ 的倍数，那么也是 $x$ 的质因子的倍数，所以我们**只需枚举质因子**，这样倍数更多。即使倍数相同，质因子作为 $k$ 也更小，符合题目要求。

对于一个固定的 $k$，设 $x=\textit{nums}[i]$，如果 $x$ 不是 $k$ 的倍数，则视作 $-x$（减去 Bob 的分数）。问题变成 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，请看 [我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。

[本题视频讲解](https://www.bilibili.com/video/BV1ioTC6BECj/?t=13m11s)，欢迎点赞关注~

```py [sol-Python3]
# 预处理每个数的质因子
MX = 1_000_001
prime_divisors = [[] for _ in range(MX)]
for i in range(2, MX):
    if not prime_divisors[i]:  # i 是质数
        for j in range(i, MX, i):  # 枚举 i 的倍数 j
            prime_divisors[j].append(i)  # i 是 j 的质因子


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
        # 收集所有质因子
        all_prime_divisors = []
        for x in nums:
            all_prime_divisors += prime_divisors[x]

        if not all_prime_divisors:
            # 每个数都是 1
            # 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2

        # 排序去重
        all_prime_divisors = sorted(set(all_prime_divisors))

        max_diff = -inf
        best_k = 0
        # 枚举质因子作为 k，计算最大子数组和
        for k in all_prime_divisors:
            diff = self.maxSubArray(nums, k)
            if diff > max_diff:
                max_diff = diff
                best_k = k

        return max_diff * best_k % MOD
```

```java [sol-Java]
class Solution {
    public int divisibleGame(int[] nums) {
        final int MOD = 1_000_000_007;

        // 收集所有质因子
        // 预处理有些慢，改成不预处理的写法
        List<Integer> allPrimeDivisors = new ArrayList<>();
        for (int x : nums) {
            for (int p = 2; p * p <= x; p++) {
                if (x % p == 0) {
                    allPrimeDivisors.add(p);
                    do {
                        x /= p;
                    } while (x % p == 0);
                }
            }
            if (x > 1) {
                allPrimeDivisors.add(x);
            }
        }

        if (allPrimeDivisors.isEmpty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        Collections.sort(allPrimeDivisors);

        int maxDiff = Integer.MIN_VALUE;
        int bestK = 0;
        int preK = 0;
        // 枚举质因子作为 k，计算最大子数组和
        for (int k : allPrimeDivisors) {
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

        return (int) ((long) maxDiff * bestK % MOD);
    }

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
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'001;
vector<int> prime_divisors[MX];

// 预处理每个数的质因子
int init = [] {
    for (int i = 2; i < MX; i++) {
        if (prime_divisors[i].empty()) { // i 是质数
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                prime_divisors[j].push_back(i); // i 是 j 的因子
            }
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

        // 收集所有质因子
        vector<int> all_prime_divisors;
        for (int x : nums) {
            auto& pd = prime_divisors[x];
            all_prime_divisors.insert(all_prime_divisors.end(), pd.begin(), pd.end());
        }

        if (all_prime_divisors.empty()) {
            // 每个数都是 1
            // 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
            return MOD - 2;
        }

        // 排序去重
        ranges::sort(all_prime_divisors);
        all_prime_divisors.erase(ranges::unique(all_prime_divisors).begin(), all_prime_divisors.end());

        int max_diff = INT_MIN;
        int best_k = 0;
        // 枚举质因子作为 k，计算最大子数组和
        for (int k : all_prime_divisors) {
            int diff = maxSubArray(nums, k);
            if (diff > max_diff) {
                max_diff = diff;
                best_k = k;
            }
        }

        return 1LL * max_diff * best_k % MOD;
    }
};
```

```go [sol-Go]
const mx = 1_000_001
var primeDivisors [mx][]int32

// 预处理每个数的质因子
func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
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
	// 收集所有质因子
	allPrimeDivisors := []int32{}
	for _, x := range nums {
		allPrimeDivisors = append(allPrimeDivisors, primeDivisors[x]...)
	}

	if len(allPrimeDivisors) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2（见示例 3）
		return mod - 2
	}

	// 排序去重
	slices.Sort(allPrimeDivisors)
	allPrimeDivisors = slices.Compact(allPrimeDivisors)

	maxDiff, bestK := math.MinInt, 0
	// 枚举质因子作为 k，计算最大子数组和
	for _, d := range allPrimeDivisors {
		k := int(d)
		diff := maxSubArray(nums, k)
		if diff > maxDiff {
			maxDiff, bestK = diff, k
		}
	}

	return maxDiff * bestK % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(Dn)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D$ 是 $\textit{nums}$ 中的不同质因子个数。由于我们去重了，最坏情况是 $\textit{nums}$ 的所有元素互不相同。在 $n$ 个不同元素的情况下，平均每个数有 $\mathcal{O}(\log\log U)$ 个质因子（$U=\max(\textit{nums})$），所以 $D = \mathcal{O}(n\log\log U)$。
- 空间复杂度：$\mathcal{O}(D)$。

## 专题训练

1. 数学题单的「**§1.3 质因数分解**」和「**§1.5 因子**」。
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
