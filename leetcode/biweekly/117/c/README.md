正难则反。总共有 $26^n$ 个字符串，减去不含 `leet` 的字符串个数，就得到了答案。

不含 `leet` 的字符串，需要至少满足如下三个条件中的一个：

1. 不含 `l`。
2. 不含 `t`。
3. 不含 `e` 或者恰好包含一个 `e`。

分类讨论。

#### 至少满足一个条件

1. 不含 `l`：每个位置可以填 $25$ 种字母，方案数为 $25^n$。
2. 不含 `t`：同上，方案数为 $25^n$。
3. 不含 `e` 或者恰好包含一个 `e`：不含 `e` 同上，方案数为 $25^n$；恰好包含一个 `e`，先从 $n$ 个位置中选一个填 `e`，然后剩下 $n-1$ 个位置不能包含 `e`，方案数为 $n\cdot 25^{n-1}$。加起来就是 $25^n + n\cdot 25^{n-1}$。

直接加起来，就是 $(3\cdot 25+n)\cdot 25^{n-1}$，但这样就重复统计了「至少满足两个条件」的情况，要减去。

#### 至少满足两个条件

1. 不含 `l` 和 `t`：每个位置可以填 $24$ 种字母，方案数为 $24^n$。
2. 不含 `l` 且 `e` 的个数不足两个：同「满足一个条件」中 3 的分析，额外不能填 `l`，方案数为 $24^n + n\cdot 24^{n-1}$。
3. 不含 `t` 且 `e` 的个数不足两个：同上，方案数为 $24^n + n\cdot 24^{n-1}$。

直接加起来，就是 $(3\cdot 24+2n)\cdot 24^{n-1}$，但这样就重复统计了「满足三个条件」的情况，要减去。

#### 满足三个条件

同「满足一个条件」中 3 的分析，额外不能填 `l` 和 `t`，方案数为 $23^n + n\cdot 23^{n-1}$。

#### 总结

不含 `leet` 的字符串的个数为「至少满足一个条件」减去「至少满足两个条件」加上「满足三个条件」，这就是**容斥原理**。

最后用 $26^n$ 减去不含 `leet` 的字符串的个数，得到答案：

$$
26^n - (3\cdot 25+n)\cdot 25^{n-1} + (3\cdot 24+2n)\cdot 24^{n-1} - (23+n)\cdot 23^{n-1}
$$

其中 $x^n$ 可以用快速幂计算，具体请看 [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)。

关于取模的知识点，见文末的「算法小课堂」。

```py [sol-Python3]
class Solution:
    def stringCount(self, n: int) -> int:
        MOD = 10 ** 9 + 7
        return (pow(26, n, MOD)
              - pow(25, n - 1, MOD) * (75 + n)
              + pow(24, n - 1, MOD) * (72 + n * 2)
              - pow(23, n - 1, MOD) * (23 + n)) % MOD
```

```java [sol-Java]
class Solution {
    private static final long MOD = (long) 1e9 + 7;

    public int stringCount(int n) {
        return (int) (((pow(26, n)
                      - pow(25, n - 1) * (75 + n)
                      + pow(24, n - 1) * (72 + n * 2)
                      - pow(23, n - 1) * (23 + n)) % MOD + MOD) % MOD); // 保证结果非负
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const long long MOD = 1e9 + 7;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    int stringCount(int n) {
        return ((pow(26, n)
               - pow(25, n - 1) * (75 + n)
               + pow(24, n - 1) * (72 + n * 2)
               - pow(23, n - 1) * (23 + n)) % MOD + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func stringCount(n int) (ans int) {
	return ((pow(26, n)-
		     pow(25, n-1)*(75+n)+
		     pow(24, n-1)*(72+n*2)-
		     pow(23, n-1)*(23+n))%mod + mod) % mod // 保证结果非负
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

#### 附：记忆化搜索

看成是「至少装满型」分组背包：有 $n$ 组物品，每组都可以从 `a` 到 `z` 中选一个，求**至少**有 $1$ 个 `l`、$1$ 个 `t` 和 $2$ 个 `e` 的方案数。

```py [sol-Python3]
@cache
def dfs(i: int, L: int, t: int, e: int) -> int:
    if i == 0:
        return 1 if L == t == e == 0 else 0
    res = dfs(i - 1, 0, t, e)  # 选 l
    res += dfs(i - 1, L, 0, e)  # 选 t
    res += dfs(i - 1, L, t, max(e - 1, 0))  # 选 e
    res += dfs(i - 1, L, t, e) * 23  # 其它字母
    return res % (10 ** 9 + 7)

class Solution:
    def stringCount(self, n: int) -> int:
        return dfs(n, 1, 1, 2)
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般涉及到取模的题目，会用到如下两个恒等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。
