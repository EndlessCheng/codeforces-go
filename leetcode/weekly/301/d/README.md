下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

考虑以 $x$ 结尾的理想数组的个数。

#### 提示 2-1

理想数组在哪些位置上发生了**变化**（$\textit{arr}[i-1]<\textit{arr}[i]$）？这些位置多吗？

#### 提示 2-2

这些位置不多，假设每次都扩大一倍，那么至多可以扩大 $O(\log x$) 次。

#### 提示 3-1

每次扩大的倍率，乘起来等于什么？

#### 提示 3-2

恰好等于 $x$。

#### 提示 3-3

对 $x$ 分解质因数。

#### 提示 4-1

把这些质因数（倍率）分配到数组的 $n$ 个位置中的某些位置上（假设数组前面还有一个数字 $1$，这样就有 $n$ 个位置）。

注意倍率可以叠加，也就是一个位置可以放多个倍率。

#### 提示 4-2

举例说明。对于 $x=8$ 来说，分解出 $k=3$ 个相同的 $2$，那么原问题变成了一个经典的组合问题：

> 把 $k=3$ 个无区别的小球放到 $n$ 个有区别的盒子中，允许盒子为空，有多少种不同的放法？

此问题可以采用隔板法来解决：把 $n$ 个盒子当做 $n-1$ 个隔板，隔板加上球总共有 $n-1+k$ 个位置，从中选择 $n-1$ 个位置放隔板，这样就把 $k$ 个球划分成了 $n$ 部分（允许部分是空），放入对应的盒子中。

因此方案数为 $C(n+k-1,n-1)=C(n+k-1,k)$。

#### 提示 4-3

对于多个不同的质因数，互相之间无影响，可以采用乘法原理计算。

枚举所有 $[1,\textit{maxValue}]$ 的 $x$，计算对应的组合数，累加即为答案。

（其他语言补充中）

```py [sol1-Python3]
MX_K = 13
MOD, MX = 10 ** 9 + 7, 10 ** 4 + MX_K

ks = [[] for _ in range(MX)]  # ks[x] 为 x 的分解质因数的幂次列表
for i in range(2, MX):
    p, x = 2, i
    while p * p <= x:
        if x % p == 0:
            k = 0
            while x % p == 0:
                k += 1
                x //= p
            ks[i].append(k)
        p += 1
    if x > 1:
        ks[i].append(1)

c = [[0] * (MX_K + 1) for _ in range(MX)]  # 组合数
c[0][0] = c[1][0] = c[1][1] = 1
for i in range(2, MX):
    c[i][0] = 1
    for j in range(1, min(i, MX_K) + 1):
        c[i][j] = (c[i - 1][j] + c[i - 1][j - 1]) % MOD

class Solution:
    def idealArrays(self, n: int, maxValue: int) -> int:
        ans = 0
        for x in range(1, maxValue + 1):
            mul = 1
            for k in ks[x]:
                mul = mul * c[n + k - 1][k] % MOD
            ans += mul
        return ans % MOD
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7, MX_K = 13, MX = (int) 1e4 + MX_K;
    static List[] ks = new List[MX]; // ks[x] 为 x 的分解质因数的幂次列表
    static int[][] c = new int[MX][MX_K + 1]; // 组合数

    static {
        for (var i = 1; i < MX; i++) {
            ks[i] = new ArrayList<Integer>();
            var x = i;
            for (var p = 2; p * p <= x; ++p) {
                if (x % p == 0) {
                    var k = 0;
                    for (; x % p == 0; x /= p) ++k;
                    ks[i].add(k);
                }
            }
            if (x > 1) ks[i].add(1);
        }

        c[0][0] = c[1][0] = c[1][1] = 1;
        for (int i = 2; i < MX; ++i) {
            c[i][0] = 1;
            for (int j = 1; j <= Math.min(i, MX_K); ++j)
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
const int MOD = 1e9 + 7, MX_K = 13, MX = 1e4 + MX_K;
vector<int> ks[MX]; // ks[x] 为 x 的分解质因数的幂次列表
int c[MX][MX_K + 1]; // 组合数

int init = []() {
    for (int i = 2; i < MX; ++i) {
        int x = i;
        for (int p = 2; p * p <= x; ++p) {
            if (x % p == 0) {
                int k = 0;
                for (; x % p == 0; x /= p) ++k;
                ks[i].push_back(k);
            }
        }
        if (x > 1) ks[i].push_back(1);
    }

    c[0][0] = c[1][0] = c[1][1] = 1;
    for (int i = 2; i < MX; ++i) {
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
            for (int k : ks[x]) mul = mul * c[n + k - 1][k] % MOD;
            ans += mul;
        }
        return ans % MOD;
    }
};
```

```go [sol1-Go]
const mod, mx int = 1e9 + 7, 1e4 + 20

var ks [mx][]int
var F, invF [mx]int

func init() {
	for i := 2; i < mx; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			k := 0
			for ; x%p == 0; x /= p {
				k++
			}
			if k > 0 {
				ks[i] = append(ks[i], k)
			}
		}
		if x > 1 {
			ks[i] = append(ks[i], 1)
		}
	}
	F[0] = 1
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx-1] = pow(F[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for m := 1; m <= maxValue; m++ {
		mul := 1
		for _, k := range ks[m] {
			comb := F[n+k-1] * invF[k] % mod * invF[n-1] % mod
			mul = mul * comb % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

