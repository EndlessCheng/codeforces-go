## 等价转化

设 $\textit{nums}$ 所有元素的按位或等于 $\textit{or}$。

本题要移除 $\textit{nums}$ 的一个子序列，要求剩余元素的按位或严格小于 $\textit{or}$。

由于**剩余元素也是 $\textit{nums}$ 的一个子序列**，所以问题等价于：

- 计算 $\textit{nums}$ 的子序列 $b$ 的个数，满足 $b$ 的按位或严格小于 $\textit{or}$。

## 正难则反

用所有子序列的个数 $2^n$，减去按位或**恰好等于** $\textit{or}$ 的子序列个数，即为按位或严格小于 $\textit{or}$ 的子序列个数。

注意我们不可能得到比 $\textit{or}$ 还大的按位或，所以这里是恰好等于。

如何计算按位或恰好等于 $\textit{or}$ 的子序列个数？

## 容斥原理

直接计算有多少个子序列的按位或恰好等于 $\textit{or}$ 是困难的。

在组合数学中，至多型问题（或者至少型问题）往往比恰好型问题更好算。

为方便描述，下面把二进制数视作集合，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

例如 $\textit{or}=11$（二进制，下同）：

- 首先计算从 $\textit{nums}$ 中选择一个子序列 $b$，满足 $b$ 的按位或是 $11$ 的**子集**的方案数。哪些数可以在 $b$ 中？如果一个数 $\textit{nums}[i]$ 不是 $11$ 的子集，那么参与或运算后，结果必然不是 $11$ 的子集。所以**只有是 $11$ 的子集的数，才能在子序列 $b$ 中**。设 $\textit{nums}$ 有 $f[11]$ 个数是 $11$ 的子集，这些数选或不选都可以，有 $2^{f[11]}$ 种方案。
- 这 $2^{f[11]}$ 种方案中，有些子序列的按位或不是恰好等于 $11$，有可能是 $10$、$01$ 或者 $00$，需要减掉。
- 减去 $b$ 的按位或是 $10$ 的子集的方案数，即 $2^{f[10]}$。
- 减去 $b$ 的按位或是 $01$ 的子集的方案数，即 $2^{f[01]}$。
- 其中 $b$ 的按位或是 $00$ 的子集的方案数多减了一次，加回来，即加上 $2^{f[00]}$。

最终得到 $b$ 的按位或恰好等于 $11$ 方案数

$$
2^{f[11]} - 2^{f[10]} - 2^{f[01]} + 2^{f[00]}
$$

一般地，根据**容斥原理**，$b$ 的按位或恰好等于 $\textit{or}$ 的方案数为

$$
\sum_{S \subseteq \textit{or}} (-1)^{|\complement_{\textit{or}}S|}2^{f[S]}
$$

其中 $\complement_{\textit{or}}S$ 表示 $S$ 关于 $\textit{or}$ 的补集，$|\complement_{\textit{or}}S|$ 即这个补集的大小（元素个数）。

通俗地说：

- 如果 $S$ 与 $\textit{or}$ 相差偶数个数，那么加上 $2^{f[S]}$。
- 如果 $S$ 与 $\textit{or}$ 相差奇数个数，那么减去 $2^{f[S]}$。

最后，剩下的问题是，如何计算 $f[S]$？也就是在 $\textit{nums}$ 中，有多少个数是 $S$ 的子集？注意，现在要算的是元素个数，不是子序列的个数。

## SOS DP

原理见 [3670 我的题解](https://leetcode.cn/problems/maximum-product-of-two-integers-with-no-common-bits/solutions/3768219/mo-ban-gao-wei-qian-zhui-he-sos-dppython-78fz/) 的方法二，状态定义改成求个数。

初始值：$f[S]$ 等于「$\textit{nums}$ 中的恰好等于 $S$ 的元素个数」。

计算 SOS DP 后，$f[S]$ 等于「$\textit{nums}$ 中的是 $S$ 的子集的元素个数」。

代码实现时，注意取模。为什么可以在中途取模？见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1arUKBbEks/)，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007
MAX_N = 100_001

# 预处理 2 的幂
pow2 = [1] * MAX_N
for i in range(1, MAX_N):
    pow2[i] = (pow2[i - 1] * 2) % MOD

class Solution:
    def countEffective(self, nums: List[int]) -> int:
        # 优化：如果 nums 只有一种数字，可以把整个数组去掉，按位或 = 0 < or_all
        if len(set(nums)) == 1:
            return 1

        or_all = reduce(or_, nums)
        w = or_all.bit_length()
        u = 1 << w

        f = [0] * u
        for x in nums:
            f[x] += 1
        for i in range(w):
            bit = 1 << i  # 避免在循环中反复计算 1 << i
            if or_all & bit == 0:  # 优化：or_all 中是 0 但 s 中是 1 的 f[s] 后面容斥用不到，无需计算
                continue
            s = 0
            while s < u:
                s |= bit  # 快速跳到第 i 位是 1 的 s
                f[s] += f[s ^ bit]
                s += 1
        # 计算完毕后，f[s] 表示 nums 中的是 s 的子集的元素个数

        ans = pow2[len(nums)]  # 所有子序列的个数
        # 枚举 or 的所有子集（包括空集）
        sub = or_all
        while True:
            p2 = pow2[f[sub]]
            ans -= -p2 if (or_all ^ sub).bit_count() % 2 else p2
            if sub == 0:
                break
            sub = (sub - 1) & or_all
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MAX_N = 100_001;
    private static final int[] pow2 = new int[MAX_N];
    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理 2 的幂
        pow2[0] = 1;
        for (int i = 1; i < MAX_N; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }
    }

    public int countEffective(int[] nums) {
        init();

        int or = 0;
        for (int x : nums) {
            or |= x;
        }

        int w = 32 - Integer.numberOfLeadingZeros(or);
        int[] f = new int[1 << w];
        for (int x : nums) {
            f[x]++;
        }
        for (int i = 0; i < w; i++) {
            if ((or >> i & 1) == 0) { // 优化：or 中是 0 但 s 中是 1 的 f[s] 后面容斥用不到，无需计算
                continue;
            }
            for (int s = 0; s < (1 << w); s++) {
                s |= 1 << i;
                f[s] += f[s ^ (1 << i)];
            }
        }
        // 计算完毕后，f[s] 表示 nums 中的是 s 的子集的元素个数

        long ans = pow2[nums.length]; // 所有子序列的个数
        // 枚举 or 的所有子集（包括空集）
        int sub = or;
        do {
            int sign = Integer.bitCount(or ^ sub) % 2 > 0 ? -1 : 1;
            ans -= sign * pow2[f[sub]];
            sub = (sub - 1) & or;
        } while (sub != or);
        return (int) ((ans % MOD + MOD) % MOD);
    }
}
```

```cpp [sol-C++]
constexpr static int MOD = 1'000'000'007;
constexpr static int MAX_N = 100'001;
int pow2[MAX_N];

int init = [] {
    // 预处理 2 的幂
    pow2[0] = 1;
    for (int i = 1; i < MAX_N; i++) {
        pow2[i] = pow2[i - 1] * 2 % MOD;
    }
    return 0;
}();

class Solution {
public:
    int countEffective(vector<int>& nums) {
        int or_ = reduce(nums.begin(), nums.end(), 0, bit_or<>());
        int w = bit_width((uint32_t) or_);

        vector<int> f(1 << w);
        for (int x : nums) {
            f[x]++;
        }
        for (int i = 0; i < w; i++) {
            if ((or_ >> i & 1) == 0) { // 优化：or_ 中是 0 但 s 中是 1 的 f[s] 后面容斥用不到，无需计算
                continue;
            }
            for (int s = 0; s < (1 << w); s++) {
                s |= 1 << i;
                f[s] += f[s ^ (1 << i)];
            }
        }
        // 计算完毕后，f[s] 表示 nums 中的是 s 的子集的元素个数

        long long ans = pow2[nums.size()]; // 所有子序列的个数
        // 枚举 or 的所有子集（包括空集）
        int sub = or_;
        do {
            int sign = popcount((uint32_t) or_ ^ sub) % 2 ? -1 : 1;
            ans -= sign * pow2[f[sub]];
            sub = (sub - 1) & or_;
        } while (sub != or_);
        return (ans % MOD + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const maxN = 100_001

var pow2 = [maxN]int{1}

func init() {
	// 预处理 2 的幂
	for i := 1; i < maxN; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func countEffective(nums []int) int {
	or := 0
	for _, x := range nums {
		or |= x
	}

	w := bits.Len(uint(or))
	f := make([]int, 1<<w)
	for _, x := range nums {
		f[x]++
	}
	for i := range w {
		if or>>i&1 == 0 { // 优化：or 中是 0 但 s 中是 1 的 f[s] 后面容斥用不到，无需计算
			continue
		}
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i
			f[s] += f[s^1<<i]
		}
	}
	// 计算完毕后，f[s] 表示 nums 中的是 s 的子集的元素个数

	ans := pow2[len(nums)] // 所有子序列的个数
	// 枚举 or 的所有子集（包括空集）
	for sub, ok := or, true; ok; ok = sub != or {
		sign := 1 - bits.OnesCount(uint(or^sub))%2*2
		ans -= sign * pow2[f[sub]]
		sub = (sub - 1) & or
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

## 专题训练

1. 数学题单的「**§2.4 容斥原理**」。
2. 动态规划题单的「**§9.5 SOS DP**」。

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
