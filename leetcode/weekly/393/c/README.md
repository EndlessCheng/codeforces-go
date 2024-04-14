由于 $k$ 越大答案越大，有单调性，可以**二分答案**。不了解二分答案的同学请看 [一图掌握二分答案！四种写法！](https://leetcode.cn/problems/h-index-ii/solution/tu-jie-yi-tu-zhang-wo-er-fen-da-an-si-ch-d15k/)

问题变成：

- 可以组合出多少个不超过 $m$ 的金额？我们可以比较个数与 $k$ 的大小，来缩小二分区间。

对于面额为 $x=\textit{coins}[i]$ 的硬币，我们可以用它组合出 $\left\lfloor\dfrac{m}{x}\right\rfloor$ 个不同的金额。

比如 $\textit{coins}=[4,6],\ m=13$，用 $4$ 可以组合出 $4,8,12$ 共 $3$ 个不同的金额，用 $6$ 可以组合出 $6,12$ 共 $2$ 个不同的金额。其中 $12$ 是重复的，需要去掉。所以一共可以组合出 $3+2-1=4$ 个不同的不超过 $m=13$ 的金额。这是**容斥原理**在 $n=2$ 的情况。

一般地，如果只有两种面额为 $x$ 和 $y$ 的硬币，则可以组合出

$$
\left\lfloor\dfrac{m}{x}\right\rfloor + \left\lfloor\dfrac{m}{y}\right\rfloor - \left\lfloor\dfrac{m}{\texttt{LCM}(x,y)}\right\rfloor
$$

个不同的不超过 $m$ 的金额。其中 $\texttt{LCM}(x,y)$ 是 $x$ 和 $y$ 的最小公倍数。 

对于更一般的容斥原理，请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/) 第三题，欢迎点赞关注！

我们需要枚举 $\textit{coins}$ 的所有**非空子集**，设子集元素的最小公倍数为 $\textit{lcm}$，那么这个子集对个数的贡献为

$$
(-1)^{k-1} \left\lfloor\dfrac{m}{\textit{lcm}}\right\rfloor
$$

累加所有非空子集的贡献，即为不同的不超过 $m$ 的金额个数。

- 开区间二分下界：$k-1$ 一定无法满足要求。
- 开区间二分上界：$\min(\textit{coins})\cdot k$ 一定可以满足要求。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

### 答疑

**问**：为什么二分出来的答案，一定是一个可以组合出的金额？

**答**：反证法。如果答案 $m$ 不是一个可以组合出的金额，那么对于 $m-1$，我们同样可以组合出 $k$ 个不同的金额，说明 $m-1$ 同样可以满足要求，这与循环不变量相矛盾。

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
        int mx = Integer.MAX_VALUE;
        for (int x : coins) {
            mx = Math.min(mx, x);
        }
        long left = k - 1, right = (long) mx * k;
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
