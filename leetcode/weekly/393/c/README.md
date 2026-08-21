对于第 $k$ 小/大问题，有如下转化套路：

- 第 $k$ 小等价于：求**最小**的 $m$，满足 $\le m$ 的数**至少**有 $k$ 个。
- 第 $k$ 大等价于：求**最大**的 $m$，满足 $\ge m$ 的数**至少**有 $k$ 个。

对于本题，$m$ 越大 $k$ 越大，$m$ 越小 $k$ 越小，所以可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

猜测上界为 $m$，问题变成：

- 可以用硬币生成多少个不超过 $m$ 的金额？比较个数与 $k$ 的大小，缩小二分区间。

对于面额为 $x=\textit{coins}[i]$ 的硬币，我们可以用它生成 $\left\lfloor\dfrac{m}{x}\right\rfloor$ 个不同的金额。

例如 $\textit{coins}=[4,6],\ m=13$，用 $4$ 可以生成 $4,8,12$ 共 $3$ 个不同的金额，用 $6$ 可以生成 $6,12$ 共 $2$ 个不同的金额。其中 $12$ 是重复的，需要去掉。所以一共可以生成 $3+2-1=4$ 个不同的不超过 $13$ 的金额。

一般地，如果只有两种面额为 $x$ 和 $y$ 的硬币，可以生成

$$
\left\lfloor\dfrac{m}{x}\right\rfloor + \left\lfloor\dfrac{m}{y}\right\rfloor - \left\lfloor\dfrac{m}{\text{lcm}(x,y)}\right\rfloor
$$

个不同的不超过 $m$ 的金额，其中 $\text{lcm}(x,y)$ 是 $x$ 和 $y$ 的最小公倍数。上式是**容斥原理**在 $n=2$ 的情况。

对于更一般的容斥原理，推导过程请看 [本题视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/?t=11m16s)，欢迎点赞关注~

枚举 $\textit{coins}$ 的所有**非空子集** $T$，这个子集对个数的贡献为

$$
(-1)^{|T|-1} \left\lfloor\dfrac{m}{\text{lcm}(T)}\right\rfloor
$$

累加所有非空子集的贡献，即为不同的不超过 $m$ 的金额个数。

- 开区间二分下界：$k-1$，一定无法满足要求。
- 开区间二分上界：$\min(\textit{coins})\cdot k$，一定可以满足要求。

代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

## 答疑

**问**：为什么二分出来的答案，一定是一个可以得到的金额？即某个 $\textit{coins}[i]$ 的倍数。

**答**：反证法。如果答案 $m$ 不是任何 $\textit{coins}[i]$ 的倍数，那么 $\le m$ 的金额个数和 $\le m-1$ 的金额个数是一样的。也就是说，对于 $m-1$，我们同样可以生成 $k$ 个不同的金额，说明 $m-1$ 同样可以满足要求，即 $\texttt{check}(m-1)$ 也是 $\texttt{true}$，这与循环不变量相矛盾。

## 优化前

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        def check(m: int) -> bool:
            cnt = 0
            for i in range(1, 1 << len(coins)):  # 枚举所有非空子集 i
                lcm_res = 1  # 计算子集 lcm
                for j, x in enumerate(coins):
                    if i >> j & 1:  # j 在集合 i 中
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

        long left = k - 1;
        long right = (long) mn * k;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
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
        for (int i = 1; i < (1 << coins.length); i++) { // 枚举所有非空子集 i
            long lcmRes = 1; // 计算子集 lcm
            for (int j = 0; j < coins.length; j++) {
                if ((i >> j & 1) == 1) { // j 在集合 i 中
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

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < (1 << coins.size()); i++) { // 枚举所有非空子集 i
                long long lcm_res = 1; // 计算子集 lcm
                for (int j = 0; j < coins.size(); j++) {
                    if (i >> j & 1) { // j 在集合 i 中
                        lcm_res = lcm(lcm_res, coins[j]);
                        if (lcm_res > m) { // 太大了
                            break;
                        }
                    }
                }
                cnt += popcount(1u * i) % 2 ? m / lcm_res : -m / lcm_res;
            }
            return cnt >= k;
        };

        long long left = k - 1, right = 1LL * ranges::min(coins) * k;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
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
		for i := uint(1); i < 1<<len(coins); i++ { // 枚举所有非空子集 i
			lcmRes := 1 // 计算子集 lcm
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

- 时间复杂度：$\mathcal{O}(n2^n\log (mk)\log M)$，其中 $n$ 是 $\textit{coins}$ 的长度，$m=\min(coins)$，$M=\max(\textit{coins})$。二分 $\mathcal{O}(\log(mk))$ 次，每次 $\mathcal{O}(n2^n\log M)$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化一

预处理每个子集的 $\text{lcm}$。

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        subset_lcm = [1] * (1 << len(coins))
        for i, x in enumerate(coins):
            bit = 1 << i
            for mask in range(bit):
                # 刷表法 DP，在 lcm(mask) 的基础上添加 x
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
                // 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }

        int mn = Integer.MAX_VALUE;
        for (int x : coins) {
            mn = Math.min(mn, x);
        }

        long left = k - 1;
        long right = (long) mn * k;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
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

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
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
                // 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }
        for (int i = 1; i < subsetLcm.length; i++) {
            if (Integer.bitCount(i) % 2 == 0) {
                subsetLcm[i] *= -1; // 避免在 check 中反复计算 Integer.bitCount
            }
        }

        int mn = Integer.MAX_VALUE;
        for (int x : coins) {
            mn = Math.min(mn, x);
        }

        long left = k - 1;
        long right = (long) mn * k;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
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

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
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
                // 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], coins[i]);
            }
        }

        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < subset_lcm.size(); i++) { // 枚举所有非空子集
                cnt += popcount(1u * i) % 2 ? m / subset_lcm[i] : -m / subset_lcm[i];
            }
            return cnt >= k;
        };

        long long left = k - 1, right = 1LL * ranges::min(coins) * k;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
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
			// 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
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

- 时间复杂度：$\mathcal{O}(2^n(\log M + \log (mk))$，其中 $n$ 是 $\textit{coins}$ 的长度，$m=\min(coins)$，$M=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 优化二

如果 $\textit{coins}$ 中有 $2$，那么所有 $2$ 的倍数我们都可以生成，那么 $\textit{coins}$ 中的其它 $2$ 的倍数（$4,6,8,\ldots$）都无需考虑。

一般地，如果一个 $\textit{coins}[i]$ 是另一个 $\textit{coins}[j]$ 的倍数，那么 $\textit{coins}[i]$ 能生成的金额，$\textit{coins}[j]$ 也能生成。所以无需考虑 $\textit{coins}[i]$。

按照这个方法可以去掉 $\textit{coins}$ 中的一些元素，降低时空复杂度。

```py [sol-Python3]
class Solution:
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        a = []
        coins.sort()  # 排序后，能整除 x 的数都在 a 中
        for x in coins:
            if all(x % y for y in a):
                a.append(x)

        subset_lcm = [1] * (1 << len(a))
        for i, x in enumerate(a):
            bit = 1 << i
            for mask in range(bit):
                # 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
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
        Arrays.sort(coins); // 排序后，能整除 x 的数都在 coins 的前 n 个数中
        int n = 0;
        next:
        for (int x : coins) {
            for (int j = 0; j < n; j++) {
                if (x % coins[j] == 0) {
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
                // 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
                subsetLcm[bit | mask] = lcm(subsetLcm[mask], coins[i]);
            }
        }
        for (int i = 1; i < subsetLcm.length; i++) {
            if (Integer.bitCount(i) % 2 == 0) {
                subsetLcm[i] *= -1; // 避免在 check 中反复计算 Integer.bitCount
            }
        }

        long left = k - 1;
        long right = (long) coins[0] * k;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
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

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private long lcm(long a, long b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        ranges::sort(coins); // 排序后，能整除 x 的数都在 coins 的前 n 个数中
        int n = 0;
        for (int x : coins) {
            bool ok = true;
            for (int j = 0; j < n; j++) {
                if (x % coins[j] == 0) {
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
                // 刷表法 DP，在 lcm(mask) 的基础上添加 coins[i]
                subset_lcm[bit | mask] = lcm(subset_lcm[mask], coins[i]);
            }
        }

        auto check = [&](long long m) -> bool {
            long long cnt = 0;
            for (int i = 1; i < subset_lcm.size(); i++) { // 枚举所有非空子集
                cnt += popcount(1u * i) % 2 ? m / subset_lcm[i] : -m / subset_lcm[i];
            }
            return cnt >= k;
        };

        long long left = k - 1, right = 1LL * coins[0] * k;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func findKthSmallest(coins []int, k int) int64 {
	slices.Sort(coins) // 排序后，能整除 x 的数都在 a 中
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

- 时间复杂度：$\mathcal{O}(n^2 + 2^t(\log M + \log (mk))$，其中 $t=\min(n, M/2)$，$n$ 是 $\textit{coins}$ 的长度，$m=\min(coins)$，$M=\max(\textit{coins})$。最坏情况下，$\textit{coins}$ 可以包含 $[13,25]$ 内的所有数字。
- 空间复杂度：$\mathcal{O}(2^t)$。

## 专题训练

见下面二分题单的「**§2.6 第 K 小/大**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
