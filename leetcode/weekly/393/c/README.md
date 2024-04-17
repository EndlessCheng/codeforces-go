由于 $k$ 越大答案越大，有单调性，可以**二分答案**。不了解二分答案的同学请看 [一图掌握二分答案！四种写法！](https://leetcode.cn/problems/h-index-ii/solution/tu-jie-yi-tu-zhang-wo-er-fen-da-an-si-ch-d15k/)

问题变成：

- 可以组合出多少个不超过 $m$ 的金额？我们可以比较个数与 $k$ 的大小，来缩小二分区间。

对于面额为 $x=\textit{coins}[i]$ 的硬币，我们可以用它组合出 $\left\lfloor\dfrac{m}{x}\right\rfloor$ 个不同的金额。

比如 $\textit{coins}=[4,6],\ m=13$，用 $4$ 可以组合出 $4,8,12$ 共 $3$ 个不同的金额，用 $6$ 可以组合出 $6,12$ 共 $2$ 个不同的金额。其中 $12$ 是重复的，需要去掉。所以一共可以组合出 $3+2-1=4$ 个不同的不超过 $m=13$ 的金额。

一般地，如果只有两种面额为 $x$ 和 $y$ 的硬币，则可以组合出

$$
\left\lfloor\dfrac{m}{x}\right\rfloor + \left\lfloor\dfrac{m}{y}\right\rfloor - \left\lfloor\dfrac{m}{\texttt{LCM}(x,y)}\right\rfloor
$$

个不同的不超过 $m$ 的金额。其中 $\texttt{LCM}(x,y)$ 是 $x$ 和 $y$ 的最小公倍数。这是**容斥原理**在 $n=2$ 的情况。

对于更一般的容斥原理，请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/) 第三题，欢迎点赞关注！

我们需要枚举 $\textit{coins}$ 的所有**非空子集**，设子集大小为 $k$，子集元素的最小公倍数为 $\textit{lcm}$，那么这个子集对个数的贡献为

$$
(-1)^{k-1} \left\lfloor\dfrac{m}{\textit{lcm}}\right\rfloor
$$

累加所有非空子集的贡献，即为不同的不超过 $m$ 的金额个数。

- 开区间二分下界：$k-1$ 一定无法满足要求。
- 开区间二分上界：$\min(\textit{coins})\cdot k$ 一定可以满足要求。

代码中使用二进制来操作集合，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 答疑

**问**：为什么二分出来的答案，一定是一个可以组合出的金额？

**答**：反证法。如果答案 $m$ 不是任何 $\textit{coins}[i]$ 的倍数，那么 $\le m$ 的金额个数和 $\le m-1$ 的金额个数是一样的。也就是说，对于 $m-1$，我们同样可以组合出 $k$ 个不同的金额，说明 $m-1$ 同样可以满足要求，即 `check(m - 1) == true`，这与循环不变量相矛盾。

## 优化前

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        def check(m: int) -> bool:
            cnt = 0
            for i in range(1, 1 << len(coins)):  # 枚举所有非空子集
                lcm_res = 1  # 计算子集 LCM
                for j, x in enumerate(coins):
                    if i >> j & 1:
                        lcm_res = lcm(lcm_res, x)
                        if lcm_res > m:  # 太大了
                            break
                else:  # 中途没有 break
                    cnt += m // lcm_res if i.bit_count() % 2 else -(m // lcm_res)
            return cnt >= k
        return bisect_left(range(min(coins) * k), True, k, key=check)
```

```java [sol-Java]
class Solution {
    public long findKthSmallest(int[] coins, int k) {
        int mn = Integer.MAX_VALUE;
        for (int x : coins) {
            mn = Math.min(mn, x);
        }
        long left = k - 1, right = (long) mn * k;
        while (left + 1 < right) {
            long mid = (left + right) / 2;
            if (check(mid, coins, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long m, int[] coins, int k) {
        long cnt = 0;
        next:
        for (int i = 1; i < (1 << coins.length); i++) { // 枚举所有非空子集
            long lcmRes = 1; // 计算子集 LCM
            for (int j = 0; j < coins.length; j++) {
                if ((i >> j & 1) == 1) {
                    lcmRes = lcm(lcmRes, coins[j]);
                    if (lcmRes > m) { // 太大了
                        continue next;
                    }
                }
            }
            cnt += Integer.bitCount(i) % 2 == 1 ? m / lcmRes : -m / lcmRes;
        }
        return cnt >= k;
    }

    private long lcm(long a, long b) {
        return a * b / gcd(a, b);
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < (1 << coins.size()); i++) { // 枚举所有非空子集
                long long lcm_res = 1; // 计算子集 LCM
                for (int j = 0; j < coins.size(); j++) {
                    if (i >> j & 1) {
                        lcm_res = lcm(lcm_res, coins[j]);
                        if (lcm_res > m) { // 太大了
                            break;
                        }
                    }
                }
                cnt += __builtin_popcount(i) % 2 ? m / lcm_res : -m / lcm_res;
            }
            return cnt >= k;
        };

        long long left = k - 1, right = (long long) ranges::min(coins) * k;
        while (left + 1 < right) {
            long long mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func findKthSmallest(coins []int, k int) int64 {
	ans := sort.Search(slices.Min(coins)*k, func(m int) bool {
		cnt := 0
	next:
		for i := uint(1); i < 1<<len(coins); i++ { // 枚举所有非空子集
			lcmRes := 1 // 计算子集 LCM
			for j := i; j > 0; j &= j - 1 {
				lcmRes = lcm(lcmRes, coins[bits.TrailingZeros(j)])
				if lcmRes > m { // 太大了
					continue next
				}
			}
			c := m / lcmRes
			if bits.OnesCount(i)%2 == 0 {
				c = -c
			}
			cnt += c
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n\log (mk)\log M)$，其中 $n$ 为 $\textit{coins}$ 的长度，$m=\min(coins),\ M=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化一

预处理每个子集的 $\texttt{LCM}$。

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        subset_lcm = [1] * (1 << len(coins))
        for i, x in enumerate(coins):
            bit = 1 << i
            for mask in range(bit):
                # 刷表法 DP，在集合 mask 的基础上添加元素 i
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], x)

        def check(m: int) -> bool:
            cnt = 0
            for i in range(1, len(subset_lcm)):  # 枚举所有非空子集
                cnt += m // subset_lcm[i] if i.bit_count() % 2 else -(m // subset_lcm[i])
            return cnt >= k
        return bisect_left(range(min(coins) * k), True, k, key=check)
```

```java [sol-Java]
class Solution {
    public long findKthSmallest(int[] coins, int k) {
        long[] subsetLcm = new long[1 << coins.length];
        subsetLcm[0] = 1;
        for (int i = 0; i < coins.length; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                //刷表法 DP，在集合 mask 的基础上添加元素 i
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }

        int mn = Integer.MAX_VALUE;
        for (int x : coins) {
            mn = Math.min(mn, x);
        }
        long left = k - 1, right = (long) mn * k;
        while (left + 1 < right) {
            long mid = (left + right) / 2;
            if (check(mid, subsetLcm, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long m, long[] subsetLcm, int k) {
        long cnt = 0;
        for (int i = 1; i < subsetLcm.length; i++) { // 枚举所有非空子集
            cnt += Integer.bitCount(i) % 2 == 1 ? m / subsetLcm[i] : -m / subsetLcm[i];
        }
        return cnt >= k;
    }

    private long lcm(long a, long b) {
        return a * b / gcd(a, b);
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public long findKthSmallest(int[] coins, int k) {
        long[] subsetLcm = new long[1 << coins.length];
        subsetLcm[0] = 1;
        for (int i = 0; i < coins.length; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                //刷表法 DP，在集合 mask 的基础上添加元素 i
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }
        for (int i = 1; i < subsetLcm.length; i++) {
            if (Integer.bitCount(i) % 2 == 0) {
                subsetLcm[i] *= -1;
            }
        }

        int mn = Integer.MAX_VALUE;
        for (int x : coins) {
            mn = Math.min(mn, x);
        }
        long left = k - 1, right = (long) mn * k;
        while (left + 1 < right) {
            long mid = (left + right) / 2;
            if (check(mid, subsetLcm, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long m, long[] subsetLcm, int k) {
        long cnt = 0;
        for (int i = 1; i < subsetLcm.length; i++) { // 枚举所有非空子集
            cnt += m / subsetLcm[i];
        }
        return cnt >= k;
    }

    private long lcm(long a, long b) {
        return a * b / gcd(a, b);
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        vector<long long> subset_lcm(1 << coins.size());
        subset_lcm[0] = 1;
        for (int i = 0; i < coins.size(); i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                // 刷表法 DP，在集合 mask 的基础上添加元素 i
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], coins[i]);
            }
        }

        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < subset_lcm.size(); i++) { // 枚举所有非空子集
                cnt += __builtin_popcount(i) % 2 ? m / subset_lcm[i] : -m / subset_lcm[i];
            }
            return cnt >= k;
        };

        long long left = k - 1, right = (long long) ranges::min(coins) * k;
        while (left + 1 < right) {
            long long mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func findKthSmallest(coins []int, k int) int64 {
	subsetLcm := make([]int, 1<<len(coins))
	subsetLcm[0] = 1
	for i, x := range coins {
		bit := 1 << i
		for mask, l := range subsetLcm[:bit] {
			// 刷表法 DP，在集合 mask 的基础上添加元素 i
			subsetLcm[bit|mask] = lcm(l, x)
		}
	}

	ans := sort.Search(slices.Min(coins)*k, func(m int) bool {
		cnt := 0
		for i := uint(1); i < 1<<len(coins); i++ { // 枚举所有非空子集
			c := m / subsetLcm[i]
			if bits.OnesCount(i)%2 == 0 {
				c = -c
			}
			cnt += c
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n(\log M + \log (mk))$，其中 $n$ 为 $\textit{coins}$ 的长度，$m=\min(coins),\ M=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 优化二

进一步地，如果 $\textit{coins}$ 中有 $2$，那么所有 $2$ 的倍数我们都可以生成，所有 $\textit{coins}$ 中的其它 $2$ 的倍数都无需考虑。按照这个方法可以去掉 $\textit{coins}$ 中的一些元素。

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        coins.sort()
        a = []
        for x in coins:
            if all(x % y for y in a):
                a.append(x)

        subset_lcm = [1] * (1 << len(a))
        for i, x in enumerate(a):
            bit = 1 << i
            for mask in range(bit):
                # 刷表法 DP，在集合 mask 的基础上添加元素 i
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], x)

        def check(m: int) -> bool:
            cnt = 0
            for i in range(1, len(subset_lcm)):  # 枚举所有非空子集
                cnt += m // subset_lcm[i] if i.bit_count() % 2 else -(m // subset_lcm[i])
            return cnt >= k
        return bisect_left(range(a[0] * k), True, k, key=check)
```

```java [sol-Java]
class Solution {
    public long findKthSmallest(int[] coins, int k) {
        Arrays.sort(coins);
        int n = 0;
        next:
        for (int x : coins) {
            for (int i = 0; i < n; i++) {
                if (x % coins[i] == 0) {
                    continue next;
                }
            }
            coins[n++] = x;
        }

        long[] subsetLcm = new long[1 << n];
        subsetLcm[0] = 1;
        for (int i = 0; i < n; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                //刷表法 DP，在集合 mask 的基础上添加元素 i
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }
        for (int i = 1; i < subsetLcm.length; i++) {
            if (Integer.bitCount(i) % 2 == 0) {
                subsetLcm[i] *= -1;
            }
        }

        long left = k - 1, right = (long) coins[0] * k;
        while (left + 1 < right) {
            long mid = (left + right) / 2;
            if (check(mid, subsetLcm, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long m, long[] subsetLcm, int k) {
        long cnt = 0;
        for (int i = 1; i < subsetLcm.length; i++) { // 枚举所有非空子集
            cnt += m / subsetLcm[i];
        }
        return cnt >= k;
    }

    private long lcm(long a, long b) {
        return a * b / gcd(a, b);
    }

    private long gcd(long a, long b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        ranges::sort(coins);
        int n = 0;
        for (int x : coins) {
            bool ok = true;
            for (int i = 0; i < n; i++) {
                if (x % coins[i] == 0) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                coins[n++] = x;
            }
        }

        vector<long long> subset_lcm(1 << n);
        subset_lcm[0] = 1;
        for (int i = 0; i < n; i++) {
            int bit = 1 << i;
            for (int mask = 0; mask < bit; mask++) {
                // 刷表法 DP，在集合 mask 的基础上添加元素 i
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], coins[i]);
            }
        }

        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < subset_lcm.size(); i++) { // 枚举所有非空子集
                cnt += __builtin_popcount(i) % 2 ? m / subset_lcm[i] : -m / subset_lcm[i];
            }
            return cnt >= k;
        };

        long long left = k - 1, right = (long long) coins[0] * k;
        while (left + 1 < right) {
            long long mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func findKthSmallest(coins []int, k int) int64 {
	slices.Sort(coins)
	a := coins[:0]
next:
	for _, x := range coins {
		for _, y := range a {
			if x%y == 0 {
				continue next
			}
		}
		a = append(a, x)
	}

	subsetLcm := make([]int, 1<<len(a))
	subsetLcm[0] = 1
	for i, x := range a {
		bit := 1 << i
		for mask, l := range subsetLcm[:bit] {
			subsetLcm[bit|mask] = lcm(l, x)
		}
	}
	for i := range subsetLcm {
		if bits.OnesCount(uint(i))%2 == 0 {
			subsetLcm[i] *= -1
		}
	}

	ans := sort.Search(a[0]*k, func(m int) bool {
		cnt := 0
		for _, l := range subsetLcm[1:] {
			cnt += m / l
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + 2^t(\log M + \log (mk))$，其中 $t=\min(n, M/2)$，$n$ 为 $\textit{coins}$ 的长度，$m=\min(coins),\ M=\max(\textit{coins})$。最坏情况下 $\textit{coins}$ 可以包含 $[13,25]$ 内的所有数字。
- 空间复杂度：$\mathcal{O}(2^t)$。

## 相似题目

- [1201. 丑数 III](https://leetcode.cn/problems/ugly-number-iii/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
