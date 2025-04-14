本题 [视频讲解](https://www.bilibili.com/video/BV1aU4y1q7BA?t=7m29s) 已出炉，欢迎点赞三连~

---

#### 提示 1

考虑以 $x$ 结尾的理想数组的个数。

#### 提示 2-1

理想数组在哪些位置上发生了**变化**（$\textit{arr}[i-1]<\textit{arr}[i]$）？这些位置多吗？

#### 提示 2-2

这些位置不多，假设每次都扩大一倍，那么至多可以扩大 $O(\log x$) 次。

#### 提示 3-1

每次扩大的倍率 $\dfrac{\textit{arr}[i]}{\textit{arr}[i-1]}$，乘起来等于什么？

#### 提示 3-2

恰好等于 $x$。

#### 提示 3-3

反过来，这些倍率是 $x$ 的因子。

我们可以先对 $x$ 分解质因数，在质因数的基础上寻找答案。

#### 提示 4-1

把这些质因数（倍率）分配到数组的 $n$ 个位置中的某些位置上，这样就可以直接确定整个理想数组（因为其余位置都和上一个元素相同）。

以 $x=8$ 为例：

- $[2,2,4,4,4,8]$ 对应着 $[\times 2,\rule{0.45cm}{0.15mm},\times 2,\rule{0.45cm}{0.15mm},\rule{0.45cm}{0.15mm},\times 2]$；
- $[1,1,4,4,8,8]$ 对应着 $[\rule{0.45cm}{0.15mm},\rule{0.45cm}{0.15mm},\times 4,\rule{0.45cm}{0.15mm},\times 2,\rule{0.45cm}{0.15mm}]$。

#### 提示 4-2

对于 $x=8$ 来说，分解出 $k=3$ 个**相同**的 $2$，那么以 $8$ 结尾的理想数组的个数，等价于一个经典的组合问题：

> 把 $k$ 个无区别的小球放到 $n$ 个有区别的盒子中，允许盒子为空，一个盒子也可以放多个小球，有多少种不同的放法？

此问题可以采用**隔板法**来解决：把 $n$ 个盒子当做 $n-1$ 个隔板，隔板加上球总共有 $n-1+k$ 个位置，从中选择 $n-1$ 个位置放隔板，$k$ 个位置放球，两个隔板之间的球（球可以有零个，一个，或者多个）放入对应盒子中（最两侧的隔板同理）。

因此方案数为 $C(n+k-1,n-1)=C(n+k-1,k)$。

#### 提示 4-3

对于多个**不同**的质因数，互相之间无影响，可以采用**乘法原理**计算。

枚举所有 $[1,\textit{maxValue}]$ 的 $x$，计算对应的组合数，累加即为答案。

代码实现时，分解出的质因数（只需要每个质因数的个数）和组合数都可以预处理出来。

#### 复杂度分析

预处理之后，`idealArrays` 的时间复杂度：$O(m\log\log m)$（$m$ 表示 $\textit{maxValue}$）。见 [Prime omega function](https://en.wikipedia.org/wiki/Prime_omega_function)。

```py [sol1-Python3]
MOD, MX = 10 ** 9 + 7, 10 ** 4 + 1

ks = [[] for _ in range(MX)]  # ks[x] 为 x 分解质因数后，每个质因数的个数列表
for i in range(2, MX):
    p, x = 2, i
    while p * p <= x:
        if x % p == 0:
            k = 1
            x //= p
            while x % p == 0:
                k += 1
                x //= p
            ks[i].append(k)
        p += 1
    if x > 1: ks[i].append(1)

class Solution:
    def idealArrays(self, n: int, maxValue: int) -> int:
        ans = 0
        for x in range(1, maxValue + 1):
            mul = 1
            for k in ks[x]:
                mul = mul * comb(n + k - 1, k) % MOD
            ans += mul
        return ans % MOD
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7, MX = (int) 1e4 + 1, MX_K = 13; // 至多 13 个质因数
    static List[] ks = new List[MX]; // ks[x] 为 x 分解质因数后，每个质因数的个数列表
    static int[][] c = new int[MX + MX_K][MX_K + 1]; // 组合数

    static {
        for (var i = 1; i < MX; i++) {
            ks[i] = new ArrayList<Integer>();
            var x = i;
            for (var p = 2; p * p <= x; ++p) {
                if (x % p == 0) {
                    var k = 1;
                    for (x /= p; x % p == 0; x /= p) ++k;
                    ks[i].add(k);
                }
            }
            if (x > 1) ks[i].add(1);
        }

        c[0][0] = 1;
        for (var i = 1; i < MX + MX_K; ++i) {
            c[i][0] = 1;
            for (var j = 1; j <= Math.min(i, MX_K); ++j)
                c[i][j] = (c[i - 1][j] + c[i - 1][j - 1]) % MOD;
        }
    }

    public int idealArrays(int n, int maxValue) {
        var ans = 0L;
        for (var x = 1; x <= maxValue; ++x) {
            var mul = 1L;
            for (var k : ks[x]) mul = mul * c[n + (int) k - 1][(int) k] % MOD;
            ans += mul;
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol1-C++]
const int MOD = 1e9 + 7, MX = 1e4 + 1, MX_K = 13; // 至多 13 个质因数
vector<int> ks[MX]; // ks[x] 为 x 分解质因数后，每个质因数的个数列表
int c[MX + MX_K][MX_K + 1]; // 组合数

int init = []() {
    for (int i = 2; i < MX; ++i) {
        int x = i;
        for (int p = 2; p * p <= x; ++p) {
            if (x % p == 0) {
                int k = 1;
                for (x /= p; x % p == 0; x /= p) ++k;
                ks[i].push_back(k);
            }
        }
        if (x > 1) ks[i].push_back(1);
    }

    c[0][0] = 1;
    for (int i = 1; i < MX + MX_K; ++i) {
        c[i][0] = 1;
        for (int j = 1; j <= min(i, MX_K); ++j)
            c[i][j] = (c[i - 1][j] + c[i - 1][j - 1]) % MOD;
    }
    return 0;
}();

class Solution {
public:
    int idealArrays(int n, int maxValue) {
        long ans = 0L;
        for (int x = 1; x <= maxValue; ++x) {
            long mul = 1L;
            for (int k: ks[x]) mul = mul * c[n + k - 1][k] % MOD;
            ans += mul;
        }
        return ans % MOD;
    }
};
```

```go [sol1-Go]
const mod, mx, mxK int = 1e9 + 7, 1e4 + 1, 13 // 至多 13 个质因数

var ks [mx][]int
var c [mx + mxK][mxK + 1]int

func init() {
	for i := 2; i < mx; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				k := 1
				for x /= p; x%p == 0; x /= p {
					k++
				}
				ks[i] = append(ks[i], k)
			}
		}
		if x > 1 {
			ks[i] = append(ks[i], 1)
		}
	}

	c[0][0] = 1
	for i := 1; i < len(c); i++ {
		c[i][0] = 1
		for j := 1; j <= mxK && j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for _, ks := range ks[1 : maxValue+1] {
		mul := 1
		for _, k := range ks {
			mul = mul * c[n+k-1][k] % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}
```

