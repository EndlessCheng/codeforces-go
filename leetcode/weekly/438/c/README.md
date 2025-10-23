## 公式推导

先不考虑取模，考察最终的元素和是由哪些元素相加得到的。

把字符串视作一个数字数组。例如数组为 $[a,b,c]$，那么操作一次后变成 $[a+b, b+c]$，再操作一次（假设操作到只剩一个数），得到 $a+2b+c$。

又例如数组为 $[a,b,c,d]$，那么操作一次后变成 $[a+b,b+c,c+d]$，再操作一次，变成 $[a+2b+c,b+2c+d]$，再操作一次，得到 $a+3b+3c+d$。

又例如数组为 $[a,b,c,d,e]$，倒数第二次操作后变成 $[a+3b+3c+d, b+3c+3d+e]$，相加得到 $a+4b+6c+4d+e$。

这和 [118. 杨辉三角](https://leetcode.cn/problems/pascals-triangle/) 的计算过程是一样的，比如 $1,4,6,4,1$ 可以理解成 $1,3,3,1,0$ 加上 $0,1,3,3,1$。所以（根据数学归纳法）系数就是组合数。

也可以从 [62. 不同路径](https://leetcode.cn/problems/unique-paths/) 的角度理解。比如第一排的 $2$，在最终结果中的系数是 $4$，这等于从第一排的 $2$ 移动到底部的**路径数**。

![lc2221.png](https://pic.leetcode.cn/1757988106-DlcZTQ-lc2221.png){:width=260}

设 $m$ 是数组长度。根据 [62 我的题解](https://leetcode.cn/problems/unique-paths/solutions/3062432/liang-chong-fang-fa-dong-tai-gui-hua-zu-o5k32/) 的方法二，我们要走 $m-1$ 步，其中恰好有 $i$ 步要往左下走，恰好有 $m-1-i$ 步要往右下走，问题变成从 $m-1$ 步中选择 $i$ 步往左下走的方案数。所以在最终结果中，$a[i]$ 的系数是组合数

$$
\dbinom {m-1} {i}
$$

本题相当于计算 $s[0]$ 到 $s[n-2]$ 和 $s[1]$ 到 $s[n-1]$ 这两个数组的如下结果：（把数组记作 $a$，长度为 $m=n-1$）

$$
\sum_{i=0}^{m-1} \dbinom {m-1} {i} \cdot a[i]
$$

也可以计算两个数组的结果的差值

$$
\sum_{i=0}^{n-2} \dbinom {n-2} {i} \cdot (s[i] - s[i+1])
$$

判断上式模 $10$ 的结果是否为 $0$。

## 方法一：提取因子 2 和 5 + 欧拉定理

关于组合数取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

本题由于模数不是质数，计算逆元无法用费马小定理。怎么办？

把每个数中的因子 $2$ 和 $5$ 单独提取出来，单独统计因子个数。一个数去掉其中所有因子 $2$ 和 $5$ 之后，和 $10$ 互质，这样可以用**欧拉定理**计算整数 $a$ 在模 $10$ 下的逆元，即 $a^{\varphi(10)-1} = a^3$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hiAUeWEUG/?t=3m44s)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
MOD = 10
MX = 100_000

f = [0] * (MX + 1)
inv_f = [0] * (MX + 1)
p2 = [0] * (MX + 1)
p5 = [0] * (MX + 1)

f[0] = inv_f[0] = 1
for i in range(1, MX + 1):
    x = i
    # 计算 2 的幂次
    e2 = (x & -x).bit_length() - 1
    x >>= e2
    # 计算 5 的幂次
    e5 = 0
    while x % 5 == 0:
        e5 += 1
        x //= 5
    f[i] = f[i - 1] * x % MOD
    inv_f[i] = pow(f[i], 3, MOD)  # 欧拉定理求逆元
    p2[i] = p2[i - 1] + e2
    p5[i] = p5[i - 1] + e5

def comb(n: int, k: int) -> int:
    # 由于每项都 < 10，所以无需中途取模
    return f[n] * inv_f[k] * inv_f[n - k] * \
        pow(2, p2[n] - p2[k] - p2[n - k], MOD) * \
        pow(5, p5[n] - p5[k] - p5[n - k], MOD)

class Solution:
    def hasSameDigits(self, s: str) -> bool:
        n = len(s)
        s = map(ord, s)
        return sum(comb(n - 2, i) * (x - y) for i, (x, y) in enumerate(pairwise(s))) % MOD == 0
```

```java [sol-Java]
class Solution {
    private static final int MOD = 10;
    private static final int MX = 100_000;

    private static final int[] f = new int[MX + 1];
    private static final int[] invF = new int[MX + 1];
    private static final int[] p2 = new int[MX + 1];
    private static final int[] p5 = new int[MX + 1];

    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        f[0] = invF[0] = 1;
        for (int i = 1; i <= MX; i++) {
            int x = i;
            // 计算 2 的幂次
            int e2 = Integer.numberOfTrailingZeros(x);
            x >>= e2;
            // 计算 5 的幂次
            int e5 = 0;
            while (x % 5 == 0) {
                e5++;
                x /= 5;
            }
            f[i] = f[i - 1] * x % MOD;
            invF[i] = pow(f[i], 3); // 欧拉定理求逆元
            p2[i] = p2[i - 1] + e2;
            p5[i] = p5[i - 1] + e5;
        }
    }

    private int pow(int x, int n) {
        int res = 1;
        while (n > 0) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
            n /= 2;
        }
        return res;
    }

    private int comb(int n, int k) {
        // 由于每项都 < 10，所以无需中途取模
        return f[n] * invF[k] * invF[n - k] *
                pow(2, p2[n] - p2[k] - p2[n - k]) *
                pow(5, p5[n] - p5[k] - p5[n - k]) % MOD;
    }

    public boolean hasSameDigits(String S) {
        init();
        char[] s = S.toCharArray();
        int diff = 0;
        for (int i = 0; i < s.length - 1; i++) {
            diff += comb(s.length - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % MOD == 0;
    }
}
```

```cpp [sol-C++]
const int MOD = 10;
const int MX = 100'000;
array<int, MX + 1> f, inv_f, p2, p5;

int qpow(int x, int n) {
    int res = 1;
    while (n > 0) {
        if (n % 2 > 0) {
            res = res * x % MOD;
        }
        x = x * x % MOD;
        n /= 2;
    }
    return res;
}

auto init = []() {
    f[0] = inv_f[0] = 1;
    for (int i = 1; i <= MX; i++) {
        int x = i;
        // 计算 2 的幂次
        int e2 = countr_zero((unsigned) x);
        x >>= e2;
        // 计算 5 的幂次
        int e5 = 0;
        while (x % 5 == 0) {
            e5++;
            x /= 5;
        }
        f[i] = f[i - 1] * x % MOD;
        inv_f[i] = qpow(f[i], 3); // 欧拉定理求逆元
        p2[i] = p2[i - 1] + e2;
        p5[i] = p5[i - 1] + e5;
    }
    return 0;
}();

int comb(int n, int k) {
    // 由于每项都 < 10，所以无需中途取模
    return f[n] * inv_f[k] * inv_f[n - k] *
           qpow(2, p2[n] - p2[k] - p2[n - k]) *
           qpow(5, p5[n] - p5[k] - p5[n - k]) % MOD;
}

class Solution {
public:
    bool hasSameDigits(string s) {
        int diff = 0;
        for (int i = 0; i + 1 < s.size(); i++) {
            diff += comb(s.size() - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % MOD == 0;
    }
};
```

```go [sol-Go]
const mod = 10

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

const mx = 100_000

var f, invF, p2, p5 [mx + 1]int

func init() {
	f[0] = 1
	invF[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		// 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2
		// 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}
		f[i] = f[i-1] * x % mod
		invF[i] = pow(f[i], 3) // 欧拉定理求逆元
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}
}

func comb(n, k int) int {
	// 由于每项都 < 10，所以无需中途取模
	return f[n] * invF[k] * invF[n-k] *
		pow(2, p2[n]-p2[k]-p2[n-k]) *
		pow(5, p5[n]-p5[k]-p5[n-k])
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%mod == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $s$ 的长度，$\log U$ 是计算快速幂的时间。
- 空间复杂度：$\mathcal{O}(1)$。

### 优化

预处理 $2$ 的幂模 $10$ 和 $5$ 的幂模 $10$。

由于 $2^i\bmod 10\ (i>0)$ 按照 $2,4,8,6$ 的周期循环，只需预处理一个长为 $4$ 的数组。

当 $i>0$ 时，$5^i\bmod 10 = 5$ 恒成立，所以无需预处理。

```py [sol-Python3]
MOD = 10
MX = 100_000
POW2 = (2, 4, 8, 6)

f = [0] * (MX + 1)
inv_f = [0] * (MX + 1)
p2 = [0] * (MX + 1)
p5 = [0] * (MX + 1)

f[0] = inv_f[0] = 1
for i in range(1, MX + 1):
    x = i
    # 计算 2 的幂次
    e2 = (x & -x).bit_length() - 1
    x >>= e2
    # 计算 5 的幂次
    e5 = 0
    while x % 5 == 0:
        e5 += 1
        x //= 5
    f[i] = f[i - 1] * x % MOD
    inv_f[i] = pow(f[i], 3, MOD)  # 欧拉定理求逆元
    p2[i] = p2[i - 1] + e2
    p5[i] = p5[i - 1] + e5

def comb(n: int, k: int) -> int:
    e2 = p2[n] - p2[k] - p2[n - k]
    return f[n] * inv_f[k] * inv_f[n - k] * \
        (POW2[(e2 - 1) % 4] if e2 else 1) * \
        (5 if p5[n] - p5[k] - p5[n - k] else 1)

class Solution:
    def hasSameDigits(self, s: str) -> bool:
        n = len(s)
        s = map(ord, s)
        return sum(comb(n - 2, i) * (x - y) for i, (x, y) in enumerate(pairwise(s))) % MOD == 0
```

```java [sol-Java]
class Solution {
    private static final int MOD = 10;
    private static final int MX = 100_000;
    private static final int[] POW2 = new int[]{2, 4, 8, 6};

    private static final int[] f = new int[MX + 1];
    private static final int[] invF = new int[MX + 1];
    private static final int[] p2 = new int[MX + 1];
    private static final int[] p5 = new int[MX + 1];

    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        f[0] = invF[0] = 1;
        for (int i = 1; i <= MX; i++) {
            int x = i;
            // 计算 2 的幂次
            int e2 = Integer.numberOfTrailingZeros(x);
            x >>= e2;
            // 计算 5 的幂次
            int e5 = 0;
            while (x % 5 == 0) {
                e5++;
                x /= 5;
            }
            f[i] = f[i - 1] * x % MOD;
            invF[i] = pow(f[i], 3); // 欧拉定理求逆元
            p2[i] = p2[i - 1] + e2;
            p5[i] = p5[i - 1] + e5;
        }
    }

    private static int pow(int x, int n) {
        int res = 1;
        while (n > 0) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
            n /= 2;
        }
        return res;
    }

    private int comb(int n, int k) {
        int e2 = p2[n] - p2[k] - p2[n - k];
        return f[n] * invF[k] * invF[n - k] *
                (e2 > 0 ? POW2[(e2 - 1) % 4] : 1) *
                (p5[n] - p5[k] - p5[n - k] > 0 ? 5 : 1) % MOD;
    }

    public boolean hasSameDigits(String S) {
        init();
        char[] s = S.toCharArray();
        int diff = 0;
        for (int i = 0; i < s.length - 1; i++) {
            diff += comb(s.length - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % MOD == 0;
    }
}
```

```cpp [sol-C++]
const int MOD = 10;
const int MX = 100'000;
const int POW2[4] = {2, 4, 8, 6};
array<int, MX + 1> f, inv_f, p2, p5;

int qpow(int x, int n) {
    int res = 1;
    while (n > 0) {
        if (n % 2 > 0) {
            res = res * x % MOD;
        }
        x = x * x % MOD;
        n /= 2;
    }
    return res;
}

auto init = []() {
    f[0] = inv_f[0] = 1;
    for (int i = 1; i <= MX; i++) {
        int x = i;
        // 计算 2 的幂次
        int e2 = countr_zero((unsigned) x);
        x >>= e2;
        // 计算 5 的幂次
        int e5 = 0;
        while (x % 5 == 0) {
            e5++;
            x /= 5;
        }
        f[i] = f[i - 1] * x % MOD;
        inv_f[i] = qpow(f[i], 3); // 欧拉定理求逆元
        p2[i] = p2[i - 1] + e2;
        p5[i] = p5[i - 1] + e5;
    }
    return 0;
}();

int comb(int n, int k) {
    int e2 = p2[n] - p2[k] - p2[n - k];
    return f[n] * inv_f[k] * inv_f[n - k] *
           (e2 ? POW2[(e2 - 1) % 4] : 1) *
           (p5[n] - p5[k] - p5[n - k] ? 5 : 1) % MOD;
}

class Solution {
public:
    bool hasSameDigits(string s) {
        int diff = 0;
        for (int i = 0; i + 1 < s.size(); i++) {
            diff += comb(s.size() - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % MOD == 0;
    }
};
```

```go [sol-Go]
const mod = 10

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

const mx = 100_000

var f, invF, p2, p5 [mx + 1]int

func init() {
	f[0] = 1
	invF[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		// 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2
		// 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}
		f[i] = f[i-1] * x % mod
		invF[i] = pow(f[i], 3) // 欧拉定理求逆元
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}
}

var pow2 = [4]int{2, 4, 8, 6}

func comb(n, k int) int {
	res := f[n] * invF[k] * invF[n-k]
	e2 := p2[n] - p2[k] - p2[n-k]
	if e2 > 0 {
		res *= pow2[(e2-1)%4]
	}
	if p5[n]-p5[k]-p5[n-k] > 0 {
		res *= 5
	}
	return res
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%mod == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：扩展 Lucas

证明见 [Lucas 定理](https://oi-wiki.org/math/number-theory/lucas/)。

Lucas 定理可以帮助我们快速计算 $r_2 = \dbinom {n} {k}\bmod 2$ 和 $r_5 = \dbinom {n} {k}\bmod 5$。

根据 [中国剩余定理](https://oi-wiki.org/math/number-theory/crt/) 的计算方法，知道了 $r_2$ 和 $r_5$，那么有

$$
\dbinom {n} {k}\bmod 10 = (5r_2+6r_5)\bmod 10
$$

```py [sol-Python3]
# 预处理组合数
MX = 5
c = [[0] * MX for i in range(MX)]
for i in range(MX):
    c[i][0] = c[i][i] = 1
    for j in range(1, i):
        c[i][j] = c[i - 1][j - 1] + c[i - 1][j]

# 计算 C(n, k) % p，要求 p 是质数
def lucas(n: int, k: int, p: int) -> int:
    if k == 0:
        return 1
    return c[n % p][k % p] * lucas(n // p, k // p, p) % p

def comb(n: int, k: int) -> int:
    # 结果至多为 5 + 4 * 6 = 29，无需中途取模
    return lucas(n, k, 2) * 5 + lucas(n, k, 5) * 6

class Solution:
    def hasSameDigits(self, s: str) -> bool:
        n = len(s)
        s = map(ord, s)
        return sum(comb(n - 2, i) * (x - y) for i, (x, y) in enumerate(pairwise(s))) % 10 == 0
```

```java [sol-Java]
class Solution {
    private static final int MX = 5;
    private static final int[][] c = new int[MX][MX];

    static {
        // 预处理组合数
        for (int i = 0; i < MX; i++) {
            c[i][0] = c[i][i] = 1;
            for (int j = 1; j < i; j++) {
                c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
            }
        }
    }

    public boolean hasSameDigits(String S) {
        char[] s = S.toCharArray();
        int diff = 0;
        for (int i = 0; i < s.length - 1; i++) {
            diff += comb(s.length - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % 10 == 0;
    }

    private int comb(int n, int k) {
        // 结果至多为 5 + 4 * 6 = 29，无需中途取模
        return lucas(n, k, 2) * 5 + lucas(n, k, 5) * 6;
    }

    // 计算 C(n, k) % p，要求 p 是质数
    private int lucas(int n, int k, int p) {
        if (k == 0) {
            return 1;
        }
        return c[n % p][k % p] * lucas(n / p, k / p, p) % p;
    }
}
```

```cpp [sol-C++]
const int MX = 5;
int c[MX][MX];

// 预处理组合数
auto init = []() {
    for (int i = 0; i < MX; i++) {
        c[i][0] = c[i][i] = 1;
        for (int j = 1; j < i; j++) {
            c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
        }
    }
    return 0;
}();

// 计算 C(n, k) % p，要求 p 是质数
int lucas(int n, int k, int p) {
    if (k == 0) {
        return 1;
    }
    return c[n % p][k % p] * lucas(n / p, k / p, p) % p;
};

int comb(int n, int k) {
    // 结果至多为 5 + 4 * 6 = 29，无需中途取模
    return lucas(n, k, 2) * 5 + lucas(n, k, 5) * 6;
}

class Solution {
public:
    bool hasSameDigits(string s) {
        int diff = 0;
        for (int i = 0; i + 1 < s.size(); i++) {
            diff += comb(s.size() - 2, i) * (s[i] - s[i + 1]);
        }
        return diff % 10 == 0;
    }
};
```

```go [sol-Go]
const mx = 5

var c [mx][mx]int

func init() {
    // 预处理组合数
	for i := range mx {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

// 计算 C(n, k) % p，要求 p 是质数
func lucas(n, k, p int) int {
	if k == 0 {
		return 1
	}
	return c[n%p][k%p] * lucas(n/p, k/p, p) % p
}

func comb(n, k int) int {
	// 结果至多为 5 + 4 * 6 = 29，无需中途取模
	return lucas(n, k, 2)*5 + lucas(n, k, 5)*6
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%10 == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。计算 $\texttt{lucas}$ 的时间为 $\mathcal{O}(\log_{p} n)$。
- 空间复杂度：$\mathcal{O}(\log_{p} n)$。递归需要 $\mathcal{O}(\log_{p} n)$ 的栈空间。

## 思考题（Hack）

方法一 Java 和 C++ 的 $\texttt{comb}$ 的返回值需要取模，否则会导致 $\textit{diff}$ 累加到一定程度后溢出。

如何构造一个 $s$，使得在 $\texttt{comb}$ 不取模的情况下，$\textit{diff}$ 既能够溢出，同时溢出后还是 $10$ 的倍数？

见 [Issue #27359](https://github.com/LeetCode-Feedback/LeetCode-Feedback/issues/27359) 的 Additional context。

## 相似题目

[2221. 数组的三角和](https://leetcode.cn/problems/find-triangular-sum-of-an-array/)

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
