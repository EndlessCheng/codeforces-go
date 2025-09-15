示例 1 的 $\textit{nums}=[4,3,2,4]$，$k=5$。

来看看 $x=3$ 时，问题相当于什么。

- 对于 $\le 3$ 的数，考虑**选或不选**。
- 对于 $> 3$ 的数，考虑**枚举选多少个**。这些数都替换成了 $3$，相当于枚举选多少个 $3$。

当 $x=3$ 时，我们可以枚举选 $0,1$ 个 $3$（不能选 $2$ 个 $3$，超过 $k=5$ 了）：

- 从大于 $3$ 的数中选 $0$ 个，问题变成：从 $\le 3$ 的数，即从 $[3,2]$ 中选择一些数，元素和能否恰好等于 $k=5$？
- 从大于 $3$ 的数中选 $1$ 个，将其替换为 $3$，问题变成：从 $\le 3$ 的数，即从 $[3,2]$ 中选择一些数，元素和能否恰好等于 $k-3=2$？

这是标准的 **0-1 背包**，原理见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)，包含**为什么要倒序循环**的讲解。

如果对于每个 $x$，都重新算一遍 $\mathcal{O}(nk)$ 的 0-1 背包，时间复杂度为 $\mathcal{O}(n^2k)$，太慢了。

我们可以把 $\textit{nums}$ 从小到大排序，然后写一个双指针，外层循环枚举 $x=1,2,3,\dots,n$，内层循环在 $[1,x-1]$ 的基础上，**增量地**考虑所有满足 $\textit{nums}[i] = x$ 的 $\textit{nums}[i]$ 的 0-1 背包。

然后（从大于 $x$ 的数中）**枚举**选 $j$ 个 $x$。

由于剩下的大于 $x$ 的数有 $n-i$ 个，且必须满足 $k-jx\ge 0$，所以 $j$ 至多枚举到

$$
\min\left(n-i,\left\lfloor\dfrac{k}{x}\right\rfloor\right)
$$

如果我们可以从 $\le x$ 的数中得到元素和 $k-jx$，那么 $\textit{ans}[x-1] = \texttt{true}$。（注意 $\textit{ans}$ 的下标从 $0$ 开始）

[本题视频讲解](https://www.bilibili.com/video/BV1TBpczdE8P/?t=3m48s)，包含 bitset 优化的讲解，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def subsequenceSumAfterCapping(self, nums: List[int], k: int) -> List[bool]:
        nums.sort()

        n = len(nums)
        ans = [False] * n
        f = [False] * (k + 1)
        f[0] = True  # 不选元素，和为 0

        i = 0
        for x in range(1, n + 1):
            # 增量地考虑所有恰好等于 x 的数
            # 小于 x 的数在之前的循环中已计算完毕，无需重复计算
            while i < n and nums[i] == x:
                for j in range(k, nums[i] - 1, -1):
                    f[j] = f[j] or f[j - nums[i]]  # 0-1 背包：不选 or 选
                i += 1

            # 枚举（从大于 x 的数中）选了 j 个 x
            for j in range(min(n - i, k // x) + 1):
                if f[k - j * x]:
                    ans[x - 1] = True
                    break
        return ans
```

```java [sol-Java]
class Solution {
    public boolean[] subsequenceSumAfterCapping(int[] nums, int k) {
        Arrays.sort(nums);

        int n = nums.length;
        boolean[] ans = new boolean[n];
        boolean[] f = new boolean[k + 1];
        f[0] = true; // 不选元素，和为 0

        int i = 0;
        for (int x = 1; x <= n; x++) {
            // 增量地考虑所有恰好等于 x 的数
            // 小于 x 的数在之前的循环中已计算完毕，无需重复计算
            while (i < n && nums[i] == x) {
                for (int j = k; j >= nums[i]; j--) {
                    f[j] = f[j] || f[j - nums[i]]; // 0-1 背包：不选 or 选
                }
                i++;
            }

            // 枚举（从大于 x 的数中）选了 j 个 x
            for (int j = 0; j <= Math.min(n - i, k / x); j++) {
                if (f[k - j * x]) {
                    ans[x - 1] = true;
                    break;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> subsequenceSumAfterCapping(vector<int>& nums, int k) {
        ranges::sort(nums);

        int n = nums.size();
        vector<bool> ans(n);
        vector<int8_t> f(k + 1);
        f[0] = true; // 不选元素，和为 0

        int i = 0;
        for (int x = 1; x <= n; x++) {
            // 增量地考虑所有恰好等于 x 的数
            // 小于 x 的数在之前的循环中已计算完毕，无需重复计算
            while (i < n && nums[i] == x) {
                for (int j = k; j >= nums[i]; j--) {
                    f[j] = f[j] || f[j - nums[i]]; // 0-1 背包：不选 or 选
                }
                i++;
            }

            // 枚举（从大于 x 的数中）选了 j 个 x
            for (int j = 0; j <= min(n - i, k / x); j++) {
                if (f[k - j * x]) {
                    ans[x - 1] = true;
                    break;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func subsequenceSumAfterCapping(nums []int, k int) []bool {
	slices.Sort(nums)

	n := len(nums)
	ans := make([]bool, n)
	f := make([]bool, k+1)
	f[0] = true // 不选元素，和为 0

	i := 0
	for x := 1; x <= n; x++ {
		// 增量地考虑所有恰好等于 x 的数
		// 小于 x 的数在之前的循环中已计算完毕，无需重复计算
		for i < n && nums[i] == x {
			for j := k; j >= nums[i]; j-- {
				f[j] = f[j] || f[j-nums[i]] // 0-1 背包：不选 or 选
			}
			i++
		}

		// 枚举（从大于 x 的数中）选了 j 个 x
		for j := range min(n-i, k/x) + 1 {
			if f[k-j*x] {
				ans[x-1] = true
				break
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk + n\log n + \min(n^2,k\log n))$，其中 $n$ 是 $\textit{nums}$ 的长度。计算 0-1 背包的双指针是 $\mathcal{O}(nk)$。排序是 $\mathcal{O}(n\log n)$。枚举 $j$ 的循环，如果 $k$ 很大，循环次数是 $\mathcal{O}(n^2)$；如果 $n$ 很大，循环次数是 $\dfrac{k}{1} + \dfrac{k}{2} + \dots + \dfrac{k}{n}$，由调和级数可知，循环次数为 $\mathcal{O}(k\log n)$，二者取最小值。
- 空间复杂度：$\mathcal{O}(k)$。忽略排序的栈开销。返回值不计入。

## 附：bitset 优化

把布尔数组用二进制表示，用位运算加速状态转移。

```py [sol-Python3]
class Solution:
    def subsequenceSumAfterCapping(self, nums: List[int], k: int) -> List[bool]:
        nums.sort()

        n = len(nums)
        ans = [False] * n
        f = 1
        u = (1 << (k + 1)) - 1

        i = 0
        for x in range(1, n + 1):
            # 增量地考虑所有恰好等于 x 的数
            while i < n and nums[i] == x:
                f |= (f << nums[i]) & u  # 保证 f 的二进制长度 <= k+1
                i += 1

            if f >> k & 1:  # 等价于优化前的 f[k]
                ans[x - 1:] = [True] * (n - x + 1)  # 后面都是 True
                break

            # 枚举（从大于 x 的数中）选了 j 个 x
            for j in range(min(n - i, k // x) + 1):
                if f >> (k - j * x) & 1:  # 等价于优化前的 f[k - j * x]
                    ans[x - 1] = True
                    break
        return ans
```

```java [sol-Java]
class Solution {
    public boolean[] subsequenceSumAfterCapping(int[] nums, int k) {
        Arrays.sort(nums);

        int n = nums.length;
        boolean[] ans = new boolean[n];
        BigInteger f = BigInteger.ONE;
        BigInteger u = BigInteger.ONE.shiftLeft(k + 1).subtract(BigInteger.ONE); // (1 << (k + 1)) - 1

        int i = 0;
        for (int x = 1; x <= n; x++) {
            // 增量地考虑所有恰好等于 x 的数
            while (i < n && nums[i] == x) {
                f = f.or(f.shiftLeft(nums[i])).and(u); // 保证 f 的二进制长度 <= k+1
                i++;
            }

            // 枚举（从大于 x 的数中）选了 j 个 x
            for (int j = 0; j <= Math.min(n - i, k / x); j++) {
                if (f.testBit(k - j * x)) {
                    ans[x - 1] = true;
                    break;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> subsequenceSumAfterCapping(vector<int>& nums, int k) {
        ranges::sort(nums);

        int n = nums.size();
        vector<bool> ans(n);
        bitset<4001> f = 1;

        int i = 0;
        for (int x = 1; x <= n; x++) {
            // 增量地考虑所有恰好等于 x 的数
            while (i < n && nums[i] == x) {
                f |= f << nums[i];
                i++;
            }

            // 枚举（从大于 x 的数中）选了 j 个 x
            for (int j = 0; j <= min(n - i, k / x); j++) {
                if (f[k - j * x]) {
                    ans[x - 1] = true;
                    break;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func subsequenceSumAfterCapping(nums []int, k int) []bool {
	slices.Sort(nums)

	n := len(nums)
	ans := make([]bool, n)
	f := big.NewInt(1)
	u := new(big.Int).Lsh(big.NewInt(1), uint(k+1))
	u.Sub(u, big.NewInt(1))

	i := 0
	for x := 1; x <= n; x++ {
		// 增量地考虑所有恰好等于 x 的数
		for i < n && nums[i] == x {
			shifted := new(big.Int).Lsh(f, uint(nums[i]))
			f.Or(f, shifted).And(f, u) // And(f, u) 保证 f 的二进制长度 <= k+1
			i++
		}

		// 枚举（从大于 x 的数中）选了 j 个 x
		for j := range min(n-i, k/x) + 1 {
			if f.Bit(k-j*x) > 0 {
				ans[x-1] = true
				break
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk/w + n\log n + \min(n^2,k\log n))$，其中 $n$ 是 $\textit{nums}$ 的长度，$w=32$ 或 $64$。计算 0-1 背包的双指针是 $\mathcal{O}(nk/w)$。排序是 $\mathcal{O}(n\log n)$。枚举 $j$ 的循环，如果 $k$ 很大，循环次数是 $\mathcal{O}(n^2)$；如果 $n$ 很大，循环次数是 $\dfrac{k}{1} + \dfrac{k}{2} + \dots + \dfrac{k}{n}$，由调和级数可知，循环次数为 $\mathcal{O}(k\log n)$，二者取最小值。
- 空间复杂度：$\mathcal{O}(k/w)$。忽略排序的栈开销。返回值不计入。

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」。

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
