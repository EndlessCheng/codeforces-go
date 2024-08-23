## 转换

由于要计算的是 $\textit{bigNums}$ 数组的元素**乘积**，而这些元素都是 $2$ 的幂，所以先计算出**幂次之和**。

$\textit{bigNums}$ 的幂次数组为

$$
[0] + [1] + [0,1] + [2] + [0,2] + [1,2] + [0,1,2] + [3] + \cdots
$$

其中每个小数组内的是 $1,2,3,4,5,6,7,8,\cdots$ 对应的强数组的幂次。

根据前缀和的思想，计算从 $\textit{from}$ 到 $\textit{to}$ 的幂次之和，等于「前 $\textit{to}+1$ 个幂次之和」减去「前 $\textit{from}$ 个幂次之和」。

## 计算幂次个数

为了计算前 $k$ 个幂次之和，我们首先要计算出，这 $k$ 个幂次是由多少个强数组组成的。

例如前 $n=3$ 个强数组，组成了前 $1+1+2=4$ 个幂次。这也等于 $[1,n]$ 中数字的二进制中的 $1$ 的个数之和。

定义 $\texttt{ones}(n)$ 为 $[0,n-1]$ 中数字的二进制中的 $1$ 的个数之和。（$0$ 中没有 $1$，为方便描述，把左边界设为 $0$。）

如果 $n=2^i$，可以证明：

$$
\texttt{ones}(2^i) = i\cdot 2^{i-1}
$$

证明方法见 [本题视频讲解](https://www.bilibili.com/video/BV1cz421m786/)。

例如 $\texttt{ones}(2^2)$ 为 $[0,3]$ 中数字的二进制中的 $1$ 的个数之和，即 $1+1+2=4$。

我们要找一个最大的 $n$，满足 $\texttt{ones}(n)\le k$。

这可以用**试填法**计算。以 $k=10$ 为例：

- 从 $i=3$ 开始考虑，因为 $i=4$ 的时候 $[0,2^i-1]$ 中的元素个数已经大于 $k$ 了，即使每个数都只算一个 $1$ 也超过了。
- 假设 $n$ 二进制从低到高第 $i=3$ 位填 $1$，即 $n=8$，那么 $\texttt{ones}(n) = 3\cdot 2^{3-1} = 12 > k$，不能填 $1$，只能填 $0$。
- 假设 $n$ 二进制从低到高第 $i=2$ 位填 $1$，即 $n=4$，那么 $\texttt{ones}(n) = 2\cdot 2^{2-1} = 4 \le k$，可以填 $1$，现在 $n=4$。
- 在 $n=4$ 的基础上，假设 $n$ 二进制从低到高第 $i=1$ 位填 $1$，即 $n=6$，我们来计算相比 $n=4$，$1$ 的个数**增加**了多少，即 $[4,6-1]$ 中的 $1$ 的个数。由于第 $2$ 位都是 $1$，所以增加量是之前填的 $1$ 的个数（$1$ 个）乘以 $[4,6-1]$ 中的元素个数（$2$ 个），加上低两位，即 $[0, 2-1]$ 中的 $1$ 的个数（$1$ 个）。现在 $1$ 的个数为 $4 + 1\cdot 2 + 1 = 7\le k$，可以填 $1$，现在 $n=6$。
- 在 $n=6$ 的基础上，假设 $n$ 二进制从低到高第 $i=0$ 位填 $1$，即 $n=7$，我们来计算相比 $n=6$，$1$ 的个数**增加**了多少，即 $[6,7-1]$ 中的 $1$ 的个数。同上，增加量是之前填的 $1$ 的个数（$2$ 个）乘以 $[6,7-1]$ 中的元素个数（$1$ 个），加上低一位，即 $[0, 1-1]$ 中的 $1$ 的个数（$0$ 个）。现在 $1$ 的个数为 $7 + 2\cdot 1 = 9\le k$，可以填 $1$，所以 $n=7$。
- 注意现在算出的 $1$ 的个数是 $9$，相比 $k=10$ 还差一个，这可以用 $n$ 的最低位补充。
- 总结一下，前 $k=10$ 个幂次分别来自 $1,2,3,4,5,6$ 的强数组中的幂次，以及 $7$ 中最低位的幂次。

## 计算幂次之和

知道了 $n$，现在来计算幂次之和。

定义 $\texttt{sumE}(n)$ 为 $[0,n-1]$ 中数字的强数组的幂次之和。（规定 $0$ 的强数组的幂次之和为 $0$。）

如果 $n=2^i$，可以证明：

$$
\texttt{sumE}(2^i) = \dfrac{i(i-1)}{2}\cdot 2^{i-1}
$$

证明方法见 [本题视频讲解](https://www.bilibili.com/video/BV1cz421m786/)。

对于一般的 $n$，计算方式同「计算幂次个数」，如果 $n$ 二进制从低到高第 $i$ 位是 $1$，那么幂次之和的**增加量**，分为如下两部分：

- 之前填的 $1$ 的幂次之和 $\textit{sumI}$ 乘以因为填 $1$ 新增加的元素个数 $2^i$。
- $\texttt{sumE}(2^i)$。

可以在「计算幂次个数」的同时计算 $\texttt{sumE}$。

得到了幂次之和，可以用**快速幂**计算 $2$ 的幂模 $\textit{mod}$，即为答案。原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

代码中用到了一些位运算技巧，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

代码实现时，乘以 $2^{i-1}$ 可以写成 `<< (i - 1)`。为了避免特判 $i=0$ 的情况，`<< (i - 1)` 可以用 `<< i >> 1` 代替

```py [sol-Python3]
class Solution:
    def findProductsOfElements(self, queries: List[List[int]]) -> List[int]:
        def sum_e(k: int) -> int:
            res = n = cnt1 = sum_i = 0
            for i in range((k + 1).bit_length() - 1, -1, -1):
                c = (cnt1 << i) + (i << i >> 1)  # 新增的幂次个数
                if c <= k:
                    k -= c
                    res += (sum_i << i) + ((i * (i - 1) // 2) << i >> 1)
                    sum_i += i  # 之前填的 1 的幂次之和
                    cnt1 += 1  # 之前填的 1 的个数
                    n |= 1 << i  # 填 1
            # 剩余的 k 个幂次，由 n 的低 k 个 1 补充
            for _ in range(k):
                lb = n & -n
                res += lb.bit_length() - 1
                n ^= lb  # 去掉最低位的 1（置为 0）
            return res
        return [pow(2, sum_e(r + 1) - sum_e(l), mod) for l, r, mod in queries]
```

```java [sol-Java]
class Solution {
    public int[] findProductsOfElements(long[][] queries) {
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            long[] q = queries[i];
            long er = sumE(q[1] + 1);
            long el = sumE(q[0]);
            ans[i] = pow(2, er - el, q[2]);
        }
        return ans;
    }

    private long sumE(long k) {
        long res = 0;
        long n = 0;
        long cnt1 = 0; // 之前填的 1 的个数
        long sumI = 0; // 之前填的 1 的幂次之和
        for (long i = 63 - Long.numberOfLeadingZeros(k + 1); i >= 0; i--) {
            long c = (cnt1 << i) + (i << i >> 1); // 新增的幂次个数
            if (c <= k) {
                k -= c;
                res += (sumI << i) + ((i * (i - 1) / 2) << i >> 1);
                sumI += i;
                cnt1++;
                n |= 1L << i; // 填 1
            }
        }
        // 剩余的 k 个幂次，由 n 的低 k 个 1 补充
        while (k-- > 0) {
            res += Long.numberOfTrailingZeros(n);
            n &= n - 1; // 去掉最低位的 1（置为 0）
        }
        return res;
    }

    private int pow(long x, long n, long mod) {
        long res = 1 % mod; // 注意 mod 可能等于 1
        for (; n > 0; n /= 2) {
            if (n % 2 == 1) {
                res = res * x % mod;
            }
            x = x * x % mod;
        }
        return (int) res;
    }
}
```

```cpp [sol-C++]
class Solution {
    int pow(long long x, long long n, long long mod) {
        long long res = 1 % mod; // 注意 mod 可能等于 1
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % mod;
            }
            x = x * x % mod;
        }
        return res;
    }

    long long sum_e(long long k) {
        long long res = 0, n = 0, cnt1 = 0, sum_i = 0;
        for (long long i = __lg(k + 1); i >= 0; i--) {
            long long c = (cnt1 << i) + (i << i >> 1); // 新增的幂次个数
            if (c <= k) {
                k -= c;
                res += (sum_i << i) + ((i * (i - 1) / 2) << i >> 1);
                sum_i += i; // 之前填的 1 的幂次之和
                cnt1++; // 之前填的 1 的个数
                n |= 1LL << i; // 填 1
            }
        }
        // 剩余的 k 个幂次，由 n 的低 k 个 1 补充
        while (k--) {
            res += __builtin_ctzll(n);
            n &= n - 1; // 去掉最低位的 1（置为 0）
        }
        return res;
    }

public:
    vector<int> findProductsOfElements(vector<vector<long long>>& queries) {
        vector<int> ans;
        for (auto& q : queries) {
            auto er = sum_e(q[1] + 1);
            auto el = sum_e(q[0]);
            ans.push_back(pow(2, er - el, q[2]));
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumE(k int) (res int) {
    var n, cnt1, sumI int
    for i := bits.Len(uint(k+1)) - 1; i >= 0; i-- {
        c := cnt1<<i + i<<i>>1 // 新增的幂次个数
        if c <= k {
            k -= c
            res += sumI<<i + i*(i-1)/2<<i>>1
            sumI += i   // 之前填的 1 的幂次之和
            cnt1++      // 之前填的 1 的个数
            n |= 1 << i // 填 1
        }
    }
    // 剩余的 k 个幂次，由 n 的低 k 个 1 补充
    for ; k > 0; k-- {
        res += bits.TrailingZeros(uint(n))
        n &= n - 1 // 去掉最低位的 1（置为 0）
    }
    return
}

func pow(x, n, mod int) int {
    res := 1 % mod // 注意 mod 可能等于 1
    for ; n > 0; n /= 2 {
        if n%2 > 0 {
            res = res * x % mod
        }
        x = x * x % mod
    }
    return res
}

func findProductsOfElements(queries [][]int64) []int {
    ans := make([]int, len(queries))
    for i, q := range queries {
        er := sumE(int(q[1]) + 1)
        el := sumE(int(q[0]))
        ans[i] = pow(2, er-el, int(q[2]))
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q\log r)$，其中 $q$ 为 $\textit{queries}$ 的长度，$r=\max(\textit{to}_i)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 思考题

原题是计算 $\textit{bigNums}$ 子数组的乘积，把乘法改成加法，也就是计算子数组的元素和，要怎么做？

欢迎在评论区分享你的思路/代码。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
