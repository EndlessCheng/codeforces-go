请看 [视频讲解](https://www.bilibili.com/video/BV1um4y1M7Rv/) 第四题。

## 提示 1

思考：$k=1$ 要怎么做？

从出现次数最多的字符中选一个字母。

## 提示 2

思考：$k=2$ 要怎么做？

从出现次数最多的开始选，如果有多个出现次数最多的呢？

例如 $s=\texttt{aaaabbbbcccc},\ k=2$，那么需要从 $3$ 种字母中选 $k$ 种，每种都有 $4$ 个字符可以选（题目说相同字符组成的子序列也算不同的），所以方案数为

$$
4^k \cdot C_3^k
$$

## 提示 3

统计每个字符**出现次数的个数**，然后从大到小遍历次数 $c$ 及其个数 $\textit{num}$。

- 如果 $\textit{num}<k$，那么这 $c$ 种字符每种选一个，方案数为 $c^{\textit{num}}$，然后将 $k$ 减去 $\textit{num}$。
- 如果 $\textit{num}\ge k$，根据上面的讨论，方案数为 $c^k\cdot C_{\textit{num}}^k$。

所有方案数相乘即为答案。

如果 $k$ 太大（循环中没有出现 $\textit{num}\ge k$），那么不存在合法子序列，返回 $0$。

有关**模运算**的小知识见文末的讲解。

代码中用到了**快速幂**，请看 [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)。

```py [sol-Python3]
class Solution:
    def countKSubsequencesWithMaxBeauty(self, s: str, k: int) -> int:
        MOD = 10 ** 9 + 7
        ans = 1
        cnt = Counter(Counter(s).values())
        for c, num in sorted(cnt.items(), reverse=True):
            if num >= k:
                return ans * pow(c, k, MOD) * comb(num, k) % MOD
            ans *= pow(c, num, MOD)
            k -= num
        return 0  # k 太大，无法选 k 个不一样的字符
```

```java [sol-Java]
class Solution {
    private static final long MOD = (long) 1e9 + 7;

    public int countKSubsequencesWithMaxBeauty(String s, int k) {
        var cnt = new int[26];
        for (char c : s.toCharArray())
            cnt[c - 'a']++;
        var cc = new TreeMap<Integer, Integer>();
        for (int c : cnt)
            if (c > 0)
                cc.merge(c, 1, Integer::sum);

        long ans = 1;
        for (var e : cc.descendingMap().entrySet()) {
            int c = e.getKey(), num = e.getValue();
            if (num >= k)
                return (int) (ans * pow(c, k) % MOD * comb(num, k) % MOD);
            ans = ans * pow(c, num) % MOD;
            k -= num;
        }
        return 0; // k 太大，无法选 k 个不一样的字符
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0)
                res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }

    // 适用于 n 和 k 都比较小的场景（本题至多 26）
    private long comb(long n, int k) {
        long res = n;
        for (int i = 2; i <= k; i++)
            res = res * --n / i; // n,n-1,n-2,... 中的前 i 个数至少有一个因子 i
        return res % MOD;
    }
}
```

```cpp [sol-C++]
class Solution {
    const long long MOD = 1e9 + 7;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }

    // 适用于 n 和 k 都比较小的场景（本题至多 26）
    long long comb(long long n, int k) {
        auto res = n;
        for (int i = 2; i <= k; i++)
            res = res * --n / i; // n,n-1,n-2,... 中的前 i 个数至少有一个因子 i
        return res % MOD;
    }

public:
    int countKSubsequencesWithMaxBeauty(string s, int k) {
        int cnt[26]{};
        for (char c: s)
            cnt[c - 'a']++;
        map<int, int> cc;
        for (int c: cnt)
            if (c) cc[-c]++; // -c 方便从大到小遍历

        long long ans = 1;
        for (auto [c, num]: cc) {
            if (num >= k)
                return ans * pow(-c, k) % MOD * comb(num, k) % MOD;
            ans = ans * pow(-c, num) % MOD;
            k -= num;
        }
        return 0; // k 太大，无法选 k 个不一样的字符
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	cc := map[int]int{}
	for _, c := range cnt {
		if c > 0 {
			cc[c]++
		}
	}

	type KV struct{ cnt, num int }
	kv := make([]KV, 0, len(cc))
	for k, v := range cc {
		kv = append(kv, KV{k, v})
	}
	sort.Slice(kv, func(i, j int) bool { return kv[i].cnt > kv[j].cnt })

	ans := 1
	for _, p := range kv {
		if p.num >= k {
			return ans * pow(p.cnt, k) % mod * comb(p.num, k) % mod
		}
		ans = ans * pow(p.cnt, p.num) % mod
		k -= p.num
	}
	return 0 // k 太大，无法选 k 个不一样的字符
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

// 适用于 n 和 k 都比较小的场景（本题至多 26）
func comb(n, k int) int {
	res := n
	for i := 2; i <= k; i++ {
		res = res * (n - i + 1) / i // n,n-1,n-2,... 中的前 i 个数至少有一个因子 i
	}
	return res % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。时间主要用在遍历字符串 $s$ 上了。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

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
