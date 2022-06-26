下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

序列计数问题，先往 DP 上想。

#### 提示 2

只定义 $f[i]$ 一个维度够吗（$i$ 表示序列长度）？

不够，没法准确统计满足题目第二个要求的序列个数。

注意到序列的元素只有 $1$ 到 $6$，我们可以把元素值也作为状态。

#### 提示 3

定义 $f[i][\textit{last}][\textit{last}_2]$ 表示序列长度为 $i$，最后一个元素是 $\textit{last}$，倒数第二个元素是 $\textit{last}_2$ 的序列数目。

通过枚举 $\textit{last}$ 和 $\textit{last}_2$，我们可以计算出 $f[i+1][j][\textit{last}]$，需满足

- $\text{GCD}(j,\textit{last})=1$
- $j\ne \textit{last}$
- $j\ne \textit{last}_2$

累加这些 $f[i][\textit{last}][\textit{last}_2]$，即得到 $f[i+1][j][\textit{last}]$。

答案为 $\sum\limits_{i}\sum\limits_{j} f[n][i][j]$。

代码实现时，可以将这些值在外部预计算出来，避免每次都重复计算一遍。

#### 复杂度分析

- 时间复杂度：$O(nm^3)$，这里 $m=6$。求 $\text{GCD}$ 的时间忽略不计（也可以打表预处理出互质的数字）。
- 空间复杂度：$O(nm^2)$。

```Python [sol1-Python3]
MOD, MX = 10 ** 9 + 7, 10 ** 4
f = [[[0] * 6 for _ in range(6)] for _ in range(MX + 1)]
f[2] = [[int(j != i and gcd(j + 1, i + 1) == 1) for j in range(6)] for i in range(6)]
for i in range(2, MX):
    for j in range(6):
        for last in range(6):
            if last != j and gcd(last + 1, j + 1) == 1:
                f[i + 1][j][last] = sum(f[i][last][last2] for last2 in range(6) if last2 != j) % MOD

class Solution:
    def distinctSequences(self, n: int) -> int:
        return sum(sum(row) for row in f[n]) % MOD if n > 1 else 6
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7, MX = (int) 1e4;
    static final int[][][] f = new int[MX + 1][6][6];

    static {
        for (var i = 0; i < 6; ++i)
            for (var j = 0; j < 6; ++j)
                if (j != i && gcd(j + 1, i + 1) == 1)
                    f[2][i][j] = 1;
        for (var i = 2; i < MX; ++i)
            for (var j = 0; j < 6; ++j)
                for (var last = 0; last < 6; ++last)
                    if (last != j && gcd(last + 1, j + 1) == 1)
                        for (var last2 = 0; last2 < 6; ++last2)
                            if (last2 != j)
                                f[i + 1][j][last] = (f[i + 1][j][last] + f[i][last][last2]) % MOD;
    }

    public int distinctSequences(int n) {
        if (n == 1) return 6;
        var ans = 0;
        for (var i = 0; i < 6; ++i)
            for (var j = 0; j < 6; ++j)
                ans = (ans + f[n][i][j]) % MOD;
        return ans;
    }

    static int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

```C++ [sol1-C++]
const int MOD = 1e9 + 7, MX = 1e4;
int f[MX + 1][6][6];
int init = []() {
    for (int i = 0; i < 6; ++i)
        for (int j = 0; j < 6; ++j)
            f[2][i][j] = j != i && gcd(j + 1, i + 1) == 1;
    for (int i = 2; i < MX; ++i)
        for (int j = 0; j < 6; ++j)
            for (int last = 0; last < 6; ++last)
                if (last != j && gcd(last + 1, j + 1) == 1)
                    for (int last2 = 0; last2 < 6; ++last2)
                        if (last2 != j)
                            f[i + 1][j][last] = (f[i + 1][j][last] + f[i][last][last2]) % MOD;
    return 0;
}();

class Solution {
public:
    int distinctSequences(int n) {
        if (n == 1) return 6;
        int ans = 0;
        for (int i = 0; i < 6; ++i)
            for (int j = 0; j < 6; ++j)
                ans = (ans + f[n][i][j]) % MOD;
        return ans;
    }
};
```

```go [sol1-Go]
const mod int = 1e9 + 7
var f = [1e4 + 1][6][6]int{}

func init() {
	for last := 0; last < 6; last++ {
		for last2 := 0; last2 < 6; last2++ {
			if last2 != last && gcd(last2+1, last+1) == 1 {
				f[2][last][last2] = 1
			}
		}
	}
	for i := 2; i < 1e4; i++ {
		for j := 0; j < 6; j++ {
			for last := 0; last < 6; last++ {
				if last != j && gcd(last+1, j+1) == 1 {
					for last2 := 0; last2 < 6; last2++ {
						if last2 != j {
							f[i+1][j][last] = (f[i+1][j][last] + f[i][last][last2]) % mod
						}
					}
				}
			}
		}
	}
}

func distinctSequences(n int) (ans int) {
	if n == 1 {
		return 6
	}
	for _, row := range f[n] {
		for _, v := range row {
			ans = (ans + v) % mod
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

附：Python 记忆化写法

```py
@cache
def f(n: int, last: int, last2: int) -> int:
    if n == 0: return 1
    res = 0
    for j in range(1, 7):
        if j != last and j != last2 and gcd(j, last) == 1:
            res += f(n - 1, j, last)
    return res % (10 ** 9 + 7)

class Solution:
    def distinctSequences(self, n: int) -> int:
        return f(n, 7, 7)  # 7 与 [1,6] 内的数字都不同且互质
```
