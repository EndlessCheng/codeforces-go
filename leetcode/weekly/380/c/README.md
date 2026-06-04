## 方法一：二分答案 + 数位 DP

$\textit{num}$ 越大，价值和也越大；$\textit{num}$ 越小，价值和也越小。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定 $\textit{num}$，计算从 $1$ 到 $\textit{num}$ 的价值和，判断其是否 $\le k$。

思路和 [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/) 是一样的，请看 [我的题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)。

设 $\textit{num}$ 的二进制长度为 $n$，我们只需要在 233 题的基础上，额外在 $d=1$ 时判断当前下标是否为 $x$ 的倍数，如果是，把 $\textit{cnt}_1$ 加一。

最后还剩一个问题：二分的上界取多少合适？

在 $1$ 到 $\textit{num}$ 中的数字 $v$，如果 $v$ 除以 $2^{x-1}$ 后如果是奇数（$v$ 右移 $x-1$ 位后最低位是 $1$），就说明 $v$ 至少包含一个我们需要的 $1$。又由于每两个连续数字中有一个是奇数，所以每连续 $2^x$ 个数中，必然有一个数至少包含一个我们需要的 $1$。所以，如果要保证价值和至少是 $k+1$，可以取 $\textit{num}$ 为

$$
(k+1) \cdot 2^x
$$

此时一定不满足二分判定，适合作为上界。

代码中用到的位运算技巧见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def findMaximumNumber(self, k: int, x: int) -> int:
        def count(num: int) -> int:
            @cache
            def dfs(i: int, cnt1: int, is_limit: bool) -> int:
                if i == 0:
                    return cnt1
                res = 0
                up = num >> (i - 1) & 1 if is_limit else 1
                for d in range(up + 1):  # 枚举要填入的数字 d
                    res += dfs(i - 1, cnt1 + (d == 1 and i % x == 0), is_limit and d == up)
                return res
            return dfs(num.bit_length(), 0, True)

        # <= k 转换成 >= k+1 的数再减一
        # 原理见 https://www.bilibili.com/video/BV1AP41137w7/
        return bisect_left(range((k + 1) << x), k + 1, key=count) - 1
```

```java [sol-Java]
class Solution {
    public long findMaximumNumber(long k, int x) {
        this.x = x;
        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        long left = 0;
        long right = (k + 1) << x;
        while (left + 1 < right) {
            long mid = (left + right) >>> 1;
            if (countDigitOne(mid) <= k) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private int x;
    private long num;
    private long memo[][];

    private long countDigitOne(long num) {
        this.num = num;
        int m = 64 - Long.numberOfLeadingZeros(num);
        memo = new long[m][m + 1];
        for (long[] row : memo) {
            Arrays.fill(row, -1);
        }
        return dfs(m - 1, 0, true);
    }

    private long dfs(int i, int cnt1, boolean isLimit) {
        if (i < 0) {
            return cnt1;
        }
        if (!isLimit && memo[i][cnt1] != -1) {
            return memo[i][cnt1];
        }
        int up = isLimit ? (int) (num >> i & 1) : 1;
        long res = 0;
        for (int d = 0; d <= up; d++) { // 枚举要填入的数字 d
            res += dfs(i - 1, cnt1 + (d == 1 && (i + 1) % x == 0 ? 1 : 0), isLimit && d == up);
        }
        if (!isLimit) {
            memo[i][cnt1] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findMaximumNumber(long long k, int x) {
        auto check = [&](long long num) -> bool {
            int m = __lg(num) + 1;
            vector<vector<long long>> memo(m, vector<long long>(m + 1, -1));
            auto dfs = [&](auto&& dfs, int i, int cnt1, bool is_limit) -> long long {
                if (i < 0) {
                    return cnt1;
                }
                if (!is_limit && memo[i][cnt1] >= 0) {
                    return memo[i][cnt1];
                }
                int up = is_limit ? num >> i & 1 : 1;
                long long res = 0;
                for (int d = 0; d <= up; d++) { // 枚举要填入的数字 d
                    res += dfs(dfs, i - 1, cnt1 + (d == 1 && (i + 1) % x == 0), is_limit && d == up);
                }
                if (!is_limit) {
                    memo[i][cnt1] = res;
                }
                return res;
            };
            return dfs(dfs, m - 1, 0, true) <= k;
        };

        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        long long left = 0, right = (k + 1) << x;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func findMaximumNumber(k int64, x int) int64 {
    ans := sort.Search(int(k+1)<<x, func(num int) bool {
        num++
        n := bits.Len(uint(num))
        memo := make([][]int, n)
        for i := range memo {
            memo[i] = make([]int, n+1)
            for j := range memo[i] {
                memo[i][j] = -1
            }
        }
        var dfs func(int, int, bool) int
        dfs = func(i, cnt1 int, limitHigh bool) (res int) {
            if i < 0 {
                return cnt1
            }
            if !limitHigh {
                p := &memo[i][cnt1]
                if *p >= 0 {
                    return *p
                }
                defer func() { *p = res }()
            }
            up := 1
            if limitHigh {
                up = num >> i & 1
            }
            for d := 0; d <= up; d++ {
                c := cnt1
                if d == 1 && (i+1)%x == 0 {
                    c++
                }
                res += dfs(i-1, c, limitHigh && d == up)
            }
            return
        }
        return dfs(n-1, 0, true) > int(k)
    })
    return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((x + \log k)^3)$。$\textit{num}$ 的二进制长度为 $\mathcal{O}(x + \log k)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}((x + \log k)^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}((x + \log k)^2)$。再算上 $\mathcal{O}(x + \log k)$ 的二分次数，总的时间复杂度为 $\mathcal{O}((x + \log k)^3)$。
- 空间复杂度：$\mathcal{O}((x + \log k)^2)$。

## 方法二：二分答案 + 数学公式

做法同 [233 题我的题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/) 的方法二。

```py [sol-Python3]
class Solution:
    def findMaximumNumber(self, k: int, x: int) -> int:
        # 统计 [1,num] 中的第 i=x,2x,3x,... 个比特位上的 1 的个数
        def count(num: int) -> int:
            res = 0
            i = x - 1  # 注意比特位从 0 开始，不是从 1 开始，所以要减一
            while num >> i:
                max_prefix = num >> (i + 1)
                # 1. prefix < max_prefix 时，低位不受约束
                # i 位填 1，suffix 随便填
                res += max_prefix << i
                if num >> i & 1:
                    # 2. prefix = max_prefix 且 i 位可以填 1
                    # i 位填 1，suffix 可以填 [0, max_suffix] 中的任意整数
                    max_suffix = num & ((1 << i) - 1)
                    res += max_suffix + 1
                i += x
            return res

        return bisect_left(range((k + 1) << x), k + 1, key=count) - 1
```

```java [sol-Java]
class Solution {
    public long findMaximumNumber(long k, int x) {
        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        long left = 0;
        long right = (k + 1) << x;
        while (left + 1 < right) {
            long mid = (left + right) >>> 1;
            if (countDigitOne(mid, x) <= k) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    // 统计 [1,num] 中的第 i=x,2x,3x,... 个比特位上的 1 的个数
    private long countDigitOne(long num, int x) {
        long res = 0;
        // 注意比特位从 0 开始，不是从 1 开始，所以 i 从 x - 1 开始
        for (int i = x - 1; (num >> i) > 0; i += x) {
			long maxPrefix = num >> (i + 1);
			// 1. prefix < maxPrefix 时，低位不受约束
			// i 位填 1，suffix 随便填
			res += maxPrefix << i;
			if ((num >> i & 1) > 0) {
				// 2. prefix = maxPrefix 且 i 位可以填 1
				// i 位填 1，suffix 可以填 [0, maxSuffix] 中的任意整数
				long maxSuffix = num & ((1L << i) - 1);
				res += maxSuffix + 1;
			}
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findMaximumNumber(long long k, int x) {
        // 统计 [1,num] 中的第 x,2x,3x,... 个比特位上的 1 的个数
        auto check = [&](long long num) -> bool {
            long long res = 0;
            // 注意比特位从 0 开始，不是从 1 开始，所以 i 从 x - 1 开始
            for (int i = x - 1; num >> i; i += x) {
                long long max_prefix = num >> (i + 1);
                // 1. prefix < max_prefix 时，低位不受约束
                // i 位填 1，suffix 随便填
                res += max_prefix << i;
                if (num >> i & 1) {
                    // 2. prefix = max_prefix 且 i 位可以填 1
                    // i 位填 1，suffix 可以填 [0, max_suffix] 中的任意整数
                    long long max_suffix = num & ((1LL << i) - 1);
                    res += max_suffix + 1;
                }
            }
            return res <= k;
        };

        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        long long left = 0, right = (k + 1) << x;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		// 统计 [1,num] 中的第 x,2x,3x,... 个比特位上的 1 的个数
		// 注意比特位从 0 开始，不是从 1 开始，所以 i 从 x - 1 开始
		res := 0
		for i := x - 1; num>>i > 0; i += x {
			maxPrefix := num >> (i + 1)
			// 1. prefix < maxPrefix 时，低位不受约束
			// i 位填 1，suffix 随便填
			res += maxPrefix << i
			if num>>i&1 > 0 {
				// 2. prefix = maxPrefix 且 i 位可以填 1
				// i 位填 1，suffix 可以填 [0, maxSuffix] 中的任意整数
				maxSuffix := num & (1<<i - 1)
				res += maxSuffix + 1
			}
		}
		return res > int(k)
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(\dfrac{\log^2 k}{x}\right)$。$\textit{num}$ 的二进制长度为 $m=\mathcal{O}(x + \log k)$，我们统计了其中的 $\mathcal{O}(m/x)$ 个比特位，所以每次二分需要 $\mathcal{O}\left(\dfrac{\log k}{x}\right)$ 的时间。再算上 $\mathcal{O}(x + \log k)$ 的二分次数，总的时间复杂度为 $\mathcal{O}\left(\dfrac{\log^2 k}{x}\right)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：试填法（逐位构造）

为方便计算，计算小于 $\textit{num}$ 的价值和，最后返回的时候把 $\textit{num}$ 减一，表示小于等于 $\textit{num}$。

从高到低构建 $\textit{num}$ 的每个比特位。设当前枚举到 $\textit{num}$ 的从低到高的第 $i$ 个比特位（$i$ 从 $0$ 开始），设 $i$ 左边有 $\textit{pre}_1$ 个编号是 $x$ 的倍数且填了 $1$ 的比特位。

如果 $\textit{num}$ 的第 $i$ 个比特位填 $1$，那么价值和会增加多少呢？例如 $i=2$，那么：

- 我们新增了 $4$ 个小于 $\textit{num}$ 的二进制数，它们在 $i=2$ 这个比特位上是 $0$，并且分别以 $00,01,10,11$ 结尾，并且这 $4$ 个二进制数在 $i$ 左边都有 $\textit{pre}_1$ 个编号是 $x$ 的倍数且填了 $1$ 的比特位。这一部分产生了 $4\cdot \textit{pre}_1$ 的价值和。
- $i$ 右边有 $\left\lfloor\dfrac{i}{x}\right\rfloor$ 个编号是 $x$ 的倍数的比特位，这些比特位一共有 $\left\lfloor\dfrac{i}{x}\right\rfloor\cdot 2^{i-1}$ 个 $1$。例如 $x=1$ 时，$00,01,10,11$ 中一共有 $4$ 个 $1$。 

两者相加，在 $\textit{num}$ 的从低到高的第 $i$ 个比特位上填 $1$，会让 $1$ 到 $\textit{num}-1$ 的价值和增加

$$
\textit{cnt} = \textit{pre}_1\cdot 2^i + \left\lfloor\dfrac{i}{x}\right\rfloor\cdot 2^{i-1}
$$

如果 $\textit{cnt}\le k$，那么这个比特位可以填 $1$。由于我们是从高到低考虑的，能填 $1$ 就填 $1$，这会让答案尽量大。然后把 $k$ 减少 $\textit{cnt}$。

最后还剩一个问题：$i$ 从哪个比特位开始枚举？

由方法一可知，$1$ 到 $(k+1) \cdot 2^x$ 的价值和至少是 $k+1$，所以我们至多考虑 $1$ 到 $(k+1) \cdot 2^x - 1$ 的价值和，所以 $i$ 应当初始化为 $(k+1) \cdot 2^x$ 的最高位，即 $(k+1) \cdot 2^x$ 的二进制长度减一。

```py [sol-Python3]
class Solution:
    def findMaximumNumber(self, k: int, x: int) -> int:
        num = pre1 = 0
        for i in range(((k + 1) << x).bit_length() - 1, -1, -1):
            cnt = (pre1 << i) + (i // x << i >> 1)
            if cnt <= k:
                k -= cnt
                num |= 1 << i
                pre1 += (i + 1) % x == 0
        return num - 1
```

```java [sol-Java]
class Solution {
    public long findMaximumNumber(long k, int x) {
        long num = 0;
        long pre1 = 0;
        for (long i = 63 - Long.numberOfLeadingZeros((k + 1) << x); i >= 0; i--) {
            long cnt = (pre1 << i) + (i / x << i >> 1);
            if (cnt > k) {
                continue;
            }
            k -= cnt;
            num |= 1L << i;
            if ((i + 1) % x == 0) {
                pre1++;
            }
        }
        return num - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findMaximumNumber(long long k, int x) {
        long long num = 0, pre1 = 0;
        for (long long i = __lg((k + 1) << x); i >= 0; i--) {
            long long cnt = (pre1 << i) + (i / x << i >> 1);
            if (cnt <= k) {
                k -= cnt;
                num |= 1LL << i;
                pre1 += (i + 1) % x == 0;
            }
        }
        return num - 1;
    }
};
```

```go [sol-Go]
func findMaximumNumber(K int64, x int) int64 {
    k := int(K)
    num, pre1 := 0, 0
    for i := bits.Len(uint((k+1)<<x)) - 1; i >= 0; i-- {
        cnt := pre1<<i + i/x<<i>>1
        if cnt > k {
            continue
        }
        k -= cnt
        num |= 1 << i
        if (i+1)%x == 0 {
            pre1++
        }
    }
    return int64(num - 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(x+\log k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 动态规划题单的「**十、数位 DP**」。
2. 思维题单的「**§5.5 贡献法**」。
3. 位运算题单的「**五、试填法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
