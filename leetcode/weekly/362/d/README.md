请看 [视频讲解](https://www.bilibili.com/video/BV1U34y1N7Pe/) 第四题。

操作等价于把末尾字母一个一个地移到开头，比如字符串 $\texttt{abcd}$，「把 $\texttt{cd}$ 移到开头」和「先把 $\texttt{d}$ 移到开头，再把 $\texttt{c}$ 移到开头」，都会得到字符串 $\texttt{cdab}$。

所以操作得到的是 $s$ 的**循环同构字符串**，这意味着，只要 $s+s$ 中包含 $t$，就可以从 $s$ 变成 $t$。比如示例 1 的 $s+s=\texttt{abcdabcd}$，其中就包含一个 $\texttt{cdab}$。

计算有多少个 $s$ 的循环同构字符串等于 $t$，记作 $c$。这可以用 KMP 等字符串匹配算法解决，即寻找 $t$ 在 $s+s$（去掉最后一个字符）中的出现次数。例如示例 2 中 $c=3$。

关于 KMP 的原理，请看我在知乎上的 [这篇讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

定义 $f[i][0]$ 表示 $i$ 次操作后等于 $t$ 的方案数，$f[i][1]$ 表示 $i$ 次操作后不等于 $t$ 的方案数。

初始值：

- 如果 $s=t$，那么 $f[0][0]=1,f[0][1]=0$。
- 如果 $s\ne t$，那么 $f[0][0]=0,f[0][1]=1$。

分类讨论（具体请看视频中的画图），如果操作后等于 $t$：

- 如果上一步也是 $t$，我们有 $c-1$ 种操作方案。
- 如果上一步不是 $t$，我们有 $c$ 种操作方案。

如果操作后不等于 $t$：

- 如果上一步是 $t$，我们有 $n-c$ 种操作方案。
- 如果上一步不是 $t$，我们有 $n-1-c$ 种操作方案。

所以有

$$
\begin{aligned}
f[i][0] &= f[i-1][0]\cdot (c-1) + f[i-1][1]\cdot c\\
f[i][1] &= f[i-1][0]\cdot (n-c) + f[i-1][1]\cdot (n-1-c)
\end{aligned}
$$

上式可以改写成如下矩阵形式

$$
\begin{bmatrix}
f[i][0] \\
f[i][1] \\
\end{bmatrix}
=
\begin{bmatrix}
c-1 & c \\
n-c & n-1-c \\
\end{bmatrix}
\cdot
\begin{bmatrix}
f[i-1][0] \\
f[i-1][1] \\
\end{bmatrix}
$$

进而得到

$$
\begin{bmatrix}
f[k][0] \\
f[k][1] \\
\end{bmatrix}
=
\begin{bmatrix}
c-1 & c \\
n-c & n-1-c \\
\end{bmatrix} ^ k
\cdot
\begin{bmatrix}
f[0][0]\\
f[0][1] \\
\end{bmatrix}
$$

利用**矩阵快速幂**（参考 [70. 爬楼梯的官方题解的方法二](https://leetcode.cn/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/)），可以得到 $f[k][0]$，即本题答案。

关于取模的知识点见文末的讲解。

```py [sol-Python3]
class Solution:
    def numberOfWays(self, s, t, k):
        n = len(s)
        c = self.kmp_search(s + s[:-1], t)
        m = [
            [c - 1, c],
            [n - c, n - 1 - c]
        ]
        m = self.pow(m, k)
        return m[0][s != t]

    # KMP 模板
    def calc_pi(self, s: str) -> List[int]:
        pi = [0] * len(s)
        c = 0
        for i in range(1, len(s)):
            v = s[i]
            while c and s[c] != v:
                c = pi[c - 1]
            if s[c] == v:
                c += 1
            pi[i] = c
        return pi

    # KMP 模板
    # 返回 text 中出现了多少次 pattern（允许 pattern 重叠）
    def kmp_search(self, text: str, pattern: str) -> int:
        pi = self.calc_pi(pattern)
        match_cnt = c = 0
        for i, v in enumerate(text):
            while c and pattern[c] != v:
                c = pi[c - 1]
            if pattern[c] == v:
                c += 1
            if c == len(pattern):
                match_cnt += 1
                c = pi[c - 1]
        return match_cnt

    # 矩阵乘法
    def multiply(self, a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
        c = [[0, 0], [0, 0]]
        for i in range(2):
            for j in range(2):
                c[i][j] = (a[i][0] * b[0][j] + a[i][1] * b[1][j]) % (10 ** 9 + 7)
        return c

    # 矩阵快速幂
    def pow(self, a: List[List[int]], n: int) -> List[List[int]]:
        res = [[1, 0], [0, 1]]
        while n:
            if n % 2:
                res = self.multiply(res, a)
            a = self.multiply(a, a)
            n //= 2
        return res
```

```java [sol-Java]
class Solution {
    public int numberOfWays(String s, String t, long k) {
        int n = s.length();
        int c = kmpSearch(s + s.substring(0, n - 1), t);
        long[][] m = {
            {c - 1, c},
            {n - c, n - 1 - c},
        };
        m = pow(m, k);
        return s.equals(t) ? (int) m[0][0] : (int) m[0][1];
    }

    // KMP 模板
    private int[] calcMaxMatch(String s) {
        int[] match = new int[s.length()];
        int c = 0;
        for (int i = 1; i < s.length(); i++) {
            char v = s.charAt(i);
            while (c > 0 && s.charAt(c) != v) {
                c = match[c - 1];
            }
            if (s.charAt(c) == v) {
                c++;
            }
            match[i] = c;
        }
        return match;
    }

    // KMP 模板
    // 返回 text 中出现了多少次 pattern（允许 pattern 重叠）
    private int kmpSearch(String text, String pattern) {
        int[] match = calcMaxMatch(pattern);
        int lenP = pattern.length();
        int matchCnt = 0;
        int c = 0;
        for (int i = 0; i < text.length(); i++) {
            char v = text.charAt(i);
            while (c > 0 && pattern.charAt(c) != v) {
                c = match[c - 1];
            }
            if (pattern.charAt(c) == v) {
                c++;
            }
            if (c == lenP) {
                matchCnt++;
                c = match[c - 1];
            }
        }
        return matchCnt;
    }

    private static final long MOD = (long) 1e9 + 7;

    // 矩阵乘法
    private long[][] multiply(long[][] a, long[][] b) {
        long[][] c = new long[2][2];
        for (int i = 0; i < 2; i++) {
            for (int j = 0; j < 2; j++) {
                c[i][j] = (a[i][0] * b[0][j] + a[i][1] * b[1][j]) % MOD;
            }
        }
        return c;
    }

    // 矩阵快速幂
    private long[][] pow(long[][] a, long n) {
        long[][] res = {{1, 0}, {0, 1}};
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = multiply(res, a);
            }
            a = multiply(a, a);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfWays(string s, string t, long long k) {
        int n = s.length();
        int c = kmp_search(s + s.substr(0, n - 1), t);
        vector<vector<long long>> m = {
            {c - 1, c},
            {n - c, n - 1 - c}
        };
        m = pow(m, k);
        return m[0][s != t];
    }

private:
    // KMP 模板
    vector<int> calc_max_match(string s) {
        vector<int> match(s.length());
        int c = 0;
        for (int i = 1; i < s.length(); i++) {
            char v = s[i];
            while (c && s[c] != v) {
                c = match[c - 1];
            }
            if (s[c] == v) {
                c++;
            }
            match[i] = c;
        }
        return match;
    }

    // KMP 模板
    // 返回 text 中出现了多少次 pattern（允许 pattern 重叠）
    int kmp_search(string text, string pattern) {
        vector<int> match = calc_max_match(pattern);
        int match_cnt = 0, c = 0;
        for (int i = 0; i < text.length(); i++) {
            char v = text[i];
            while (c && pattern[c] != v) {
                c = match[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            if (c == pattern.length()) {
                match_cnt++;
                c = match[c - 1];
            }
        }
        return match_cnt;
    }

    const long long MOD = 1e9 + 7;

    // 矩阵乘法
    vector<vector<long long>> multiply(vector<vector<long long>> &a, vector<vector<long long>> &b) {
        vector<vector<long long>> c(2, vector<long long>(2));
        for (int i = 0; i < 2; i++) {
            for (int j = 0; j < 2; j++) {
                c[i][j] = (a[i][0] * b[0][j] + a[i][1] * b[1][j]) % MOD;
            }
        }
        return c;
    }

    // 矩阵快速幂
    vector<vector<long long>> pow(vector<vector<long long>> &a, long long n) {
        vector<vector<long long>> res = {{1, 0}, {0, 1}};
        for (; n; n /= 2) {
            if (n % 2) {
                res = multiply(res, a);
            }
            a = multiply(a, a);
        }
        return res;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int64) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func numberOfWays(s, t string, k int64) int {
	n := len(s)
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (cnt int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				if i-lenP+1 < n {
					cnt++
				}
				c = match[c-1]
			}
		}
		return
	}
	c := kmpSearch(s+s, t)
	m := matrix{
		{c - 1, c},
		{n - c, n - 1 - c},
	}.pow(k)
	if s == t {
		return m[0][0]
	}
	return m[0][1]
}
```

```js [sol-JavaScript]
var numberOfWays = function (s, t, k) {
    const n = s.length;
    const c = kmpSearch(s + s.substring(0, n - 1), t);
    const m = [
        [BigInt(c - 1), BigInt(c)],
        [BigInt(n - c), BigInt(n - 1 - c)],
    ];
    const res = pow(m, k);
    return s === t ? res[0][0] : res[0][1];
};

// KMP 模板
function calcMaxMatch(s) {
    const match = new Array(s.length).fill(0);
    let c = 0;
    for (let i = 1; i < s.length; i++) {
        const v = s.charAt(i);
        while (c && s.charAt(c) !== v) {
            c = match[c - 1];
        }
        if (s.charAt(c) === v) {
            c++;
        }
        match[i] = c;
    }
    return match;
}

// KMP 模板
// 返回 text 中出现了多少次 pattern（允许 pattern 重叠）
function kmpSearch(text, pattern) {
    const match = calcMaxMatch(pattern);
    let matchCnt = 0;
    let c = 0;
    for (let i = 0; i < text.length; i++) {
        const v = text.charAt(i);
        while (c && pattern.charAt(c) !== v) {
            c = match[c - 1];
        }
        if (pattern.charAt(c) === v) {
            c++;
        }
        if (c === pattern.length) {
            matchCnt++;
            c = match[c - 1];
        }
    }
    return matchCnt;
}

// 矩阵乘法
function multiply(a, b) {
    const c = [[0, 0], [0, 0]]
    for (let i = 0; i < 2; i++) {
        for (let j = 0; j < 2; j++) {
            c[i][j] = (a[i][0] * b[0][j] + a[i][1] * b[1][j]) % BigInt(1e9 + 7);
        }
    }
    return c;
}

// 矩阵快速幂
function pow(a, n) {
    let res = [[BigInt(1), BigInt(0)], [BigInt(0), BigInt(1)]];
    while (n) {
        if (n % 2) {
            res = multiply(res, a);
        }
        a = multiply(a, a);
        n = Math.floor(n / 2);
    }
    return res;
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log k)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 练习：矩阵快速幂

- [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)
- [509. 斐波那契数](https://leetcode.cn/problems/fibonacci-number/)
- [1137. 第 N 个泰波那契数](https://leetcode.cn/problems/n-th-tribonacci-number/)
- [1220. 统计元音字母序列的数目](https://leetcode.cn/problems/count-vowels-permutation/)
- [552. 学生出勤记录 II](https://leetcode.cn/problems/student-attendance-record-ii/)
- [790. 多米诺和托米诺平铺](https://leetcode.cn/problems/domino-and-tromino-tiling/)

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
