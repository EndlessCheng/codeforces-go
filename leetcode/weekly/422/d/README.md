## 公式推导

设 $\textit{num}$ 中所有数字之和为 $\textit{total}$。

如果 $\textit{total}$ 是奇数，那么无法把 $\textit{num}$ 中的数字分成两个和相等的集合，返回 $0$。

否则，问题相当于把 $\textit{num}$ **均分**成两个数字之和均为 $\dfrac{\textit{total}}{2}$ 的多重集。

例如其中一个多重集为 $\{1,1,2,2,2\}$，那么 $5$ 个数有 $5!$ 个排列，其中 $2$ 个 $1$ 的排列个数 $2!$ 是重复的，要除掉；另外 $3$ 个 $2$ 的排列个数 $3!$ 是重复的，要除掉。所以这个多重集的排列数为 $\dfrac{5!}{2!3!}$。

设 $\textit{num}$ 中数字 $i$ 的出现次数为 $\textit{cnt}[i]$。

设有 $k_i$ 个数字 $i$ 分给第一个多重集，那么剩余的 $\textit{cnt}[i] - k_i$ 个数字 $i$ 分给第二个多重集。

设 $n$ 是 $\textit{num}$ 的长度。

第一个多重集的大小为 $\left\lfloor\dfrac{n}{2}\right\rfloor$，排列数为

$$
\dfrac{\left\lfloor\dfrac{n}{2}\right\rfloor!}{\prod\limits_{i=0}^{i=9}k_i!}
$$

第二个多重集的大小为 $\left\lceil\dfrac{n}{2}\right\rceil$，排列数为

$$
\dfrac{\left\lceil\dfrac{n}{2}\right\rceil!}{\prod\limits_{i=0}^{i=9}(\textit{cnt}[i]-k_i)!}
$$

二者相乘，总的排列数为

$$
\dfrac{\left\lfloor\dfrac{n}{2}\right\rfloor!\left\lceil\dfrac{n}{2}\right\rceil!}{\left(\prod\limits_{i=0}^{i=9}k_i!\right)\left(\prod\limits_{i=0}^{i=9}(\textit{cnt}[i]-k_i)!\right)}
$$

由于分子可以直接计算，所以下面只计算

$$
f_9(k_0,k_1,\ldots,k_9) = \dfrac{1}{\left(\prod\limits_{i=0}^{i=9}k_i!\right)\left(\prod\limits_{i=0}^{i=9}(\textit{cnt}[i]-k_i)!\right)}
$$

如果只枚举 $k_9$ 的话，有

$$
\sum_{k_9=0}^{\textit{cnt}[9]} f_9(k_0,k_1,\ldots,k_9) =  \sum_{k_9=0}^{\textit{cnt}[9]} f_8(k_0,k_1,\ldots,k_8)\cdot \dfrac{1}{k_9!(\textit{cnt}[9]-k_9)!}
$$

其中 $f_8(k_0,k_1,\ldots,k_8) = \dfrac{1}{\left(\prod\limits_{i=0}^{i=8}k_i!\right)\left(\prod\limits_{i=0}^{i=8}(\textit{cnt}[i]-k_i)!\right)}$，这又可以通过枚举 $k_8$ 计算，转换成计算 $f_7(k_0,k_1,\ldots,k_7)$ 的子问题。

## 状态定义与状态转移方程

对于每个 $i=0,1,2,\ldots,9$，我们需要枚举分配多少个数字 $i$ 给第一个多重集。此外有如下约束：

- 所有数字分配完毕时，第一个多重集的大小必须恰好等于 $\left\lfloor\dfrac{n}{2}\right\rfloor$。此时第二个多重集的大小一定等于 $\left\lceil\dfrac{n}{2}\right\rceil$。
- 所有数字分配完毕时，第一个多重集的数字之和，必须等于第二个多重集的数字之和。这等价于第一个多重集的数字之和等于 $\dfrac{\textit{total}}{2}$。

为此，我们需要在记忆化搜索/递推的过程中，维护如下变量：

- 剩余要分配的数字是 $[0,i]$，当前要分配的数字是 $i$。
- 第一个多重集还剩下 $\textit{left}_1$ 个数字需要分配。
- 第一个多重集还剩下 $\textit{leftS}$ 的元素和需要分配。

所以，定义 $\textit{dfs}(i,\textit{left}_1,\textit{leftS})$ 表示在剩余要分配的数字是 $[0,i]$，第一个多重集还剩下 $\textit{left}_1$ 个数字需要分配，第一个多重集还剩下 $\textit{leftS}$ 的元素和需要分配的情况下，下式的结果：

$$
\sum_{k_i=0}^{\textit{cnt}[i]} f_i(k_0,k_1,\ldots,k_i)
$$

枚举数字 $i$ 分出 $k$ 个数给第一个多重集，要解决的问题变为：

- 剩余要分配的数字是 $[0,i-1]$。
- 第一个多重集还剩下 $\textit{left}_1 - k$ 个数字需要分配。
- 第一个多重集还剩下 $\textit{leftS} - k\cdot i$ 的元素和需要分配。
- 计算的式子为 $\sum\limits_{k_{i-1}=0}^{\textit{cnt}[i-1]} f_{i-1}(k_0,k_1,\ldots,k_{i-1})$。

即 $\textit{dfs}(i-1,\textit{left}_1 - k, \textit{leftS} - k\cdot i)$。

累加得

$$
\textit{dfs}(i,\textit{left}_1,\textit{leftS}) = \sum_{k=0}^{\textit{cnt}[i]}  \textit{dfs}(i-1,\textit{left}_1 - k, \textit{leftS} - k\cdot i)\cdot \dfrac{1}{k!(\textit{cnt}[i]-k)!}
$$

由于 $\textit{left}_1+\textit{left}_2 = \sum\limits_{j=0}^{i} \textit{cnt}[j]$ 恒成立，所以第二个多重集的大小 $\textit{left}_2$ 可以省略。

注意枚举 $k$ 的时候，还要满足 $k\le \textit{left}_1$ 且 $\textit{cnt}[i]-k \le \textit{left}_2$，所以 $k$ 的实际范围为

$$
[\max(\textit{cnt}[i]-\textit{left}_2,0), \min(\textit{cnt}[i], \textit{left}_1)]
$$

递归边界：$\textit{dfs}(-1,0,0)=1$，其余 $\textit{dfs}(-1,0,\textit{leftS})=0$。

递归入口：$\textit{dfs}(9,n_1,\textit{total}/2)$，其中 $n_1= \left\lfloor\dfrac{n}{2}\right\rfloor$。

最终答案为

$$
n_1!\times (n-n_1)!\times \textit{dfs}(9,n_1,\textit{total}/2)
$$

关于取模的知识点，以及逆元的知识点，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 记忆化搜索

```py [sol-Python3]
MOD = 1_000_000_007
MX = 41

fac = [0] * MX  # f[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

class Solution:
    def countBalancedPermutations(self, num: str) -> int:
        cnt = [0] * 10
        total = 0
        for c in map(int, num):
            cnt[c] += 1
            total += c

        if total % 2:
            return 0

        pre = list(accumulate(cnt))

        @cache
        def dfs(i: int, left1: int, left_s: int) -> int:
            if i < 0:
                return 1 if left_s == 0 else 0
            res = 0
            c = cnt[i]
            left2 = pre[i] - left1
            for k in range(max(c - left2, 0), min(c, left1) + 1):
                if k * i > left_s:
                    break
                r = dfs(i - 1, left1 - k, left_s - k * i)
                res += r * inv_f[k] * inv_f[c - k]
            return res % MOD

        n = len(num)
        n1 = n // 2
        return fac[n1] * fac[n - n1] * dfs(9, n1, total // 2) % MOD
```

```py [sol-Python3 写法二]
MOD = 1_000_000_007
MX = 41

fac = [0] * MX  # f[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

class Solution:
    def countBalancedPermutations(self, num: str) -> int:
        cnt = [0] * 10
        total = 0
        for c in map(int, num):
            cnt[c] += 1
            total += c

        if total % 2:
            return 0

        pre = list(accumulate(cnt))

        @cache
        def dfs(i: int, left1: int, left_s: int) -> int:
            if i == 0:
                return inv_f[left1] * inv_f[cnt[0] - left1] % MOD if left_s == 0 and left1 <= cnt[0] else 0
            res = 0
            c = cnt[i]
            left2 = pre[i] - left1
            for k in range(max(c - left2, 0), min(c, left1, left_s // i) + 1):
                r = dfs(i - 1, left1 - k, left_s - k * i)
                res += r * inv_f[k] * inv_f[c - k]
            return res % MOD

        n = len(num)
        n1 = n // 2
        return fac[n1] * fac[n - n1] * dfs(9, n1, total // 2) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 41;

    private static final long[] F = new long[MX]; // f[i] = i!
    private static final long[] INV_F = new long[MX]; // inv_f[i] = i!^-1

    static {
        F[0] = 1;
        for (int i = 1; i < MX; i++) {
            F[i] = F[i - 1] * i % MOD;
        }
        INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_F[i - 1] = INV_F[i] * i % MOD;
        }
    }

    public int countBalancedPermutations(String num) {
        int[] cnt = new int[10];
        int total = 0;
        for (char c : num.toCharArray()) {
            cnt[c - '0']++;
            total += c - '0';
        }

        if (total % 2 > 0) {
            return 0;
        }

        for (int i = 1; i < 10; i++) {
            cnt[i] += cnt[i - 1];
        }

        int n = num.length();
        int n1 = n / 2;
        int[][][] memo = new int[10][n1 + 1][total / 2 + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }
        return (int) (F[n1] * F[n - n1] % MOD * dfs(9, n1, total / 2, cnt, memo) % MOD);
    }

    private int dfs(int i, int left1, int leftS, int[] cnt, int[][][] memo) {
        if (i < 0) {
            return leftS == 0 ? 1 : 0;
        }
        if (memo[i][left1][leftS] != -1) {
            return memo[i][left1][leftS];
        }
        long res = 0;
        int c = cnt[i] - (i > 0 ? cnt[i - 1] : 0);
        int left2 = cnt[i] - left1;
        for (int k = Math.max(c - left2, 0); k <= Math.min(c, left1) && k * i <= leftS; k++) {
            long r = dfs(i - 1, left1 - k, leftS - k * i, cnt, memo);
            res = (res + r * INV_F[k] % MOD * INV_F[c - k]) % MOD;
        }
        return memo[i][left1][leftS] = (int) res;
    }

    private static long pow(long x, int n) {
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
const int MOD = 1'000'000'007;
const int MX = 41;

long long F[MX]; // F[i] = i!
long long INV_F[MX]; // INV_F[i] = i!^-1

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

auto init = [] {
    F[0] = 1;
    for (int i = 1; i < MX; i++) {
        F[i] = F[i - 1] * i % MOD;
    }
    INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
    for (int i = MX - 1; i; i--) {
        INV_F[i - 1] = INV_F[i] * i % MOD;
    }
    return 0;
}();

class Solution {
public:
    int countBalancedPermutations(string num) {
        int cnt[10];
        int total = 0;
        for (char c : num) {
            cnt[c - '0']++;
            total += c - '0';
        }

        if (total % 2) {
            return 0;
        }

        // 原地求前缀和
        partial_sum(cnt, cnt + 10, cnt);

        int n = num.size(), n1 = n / 2;
        vector<vector<vector<int>>> memo(10, vector(n1 + 1, vector<int>(total / 2 + 1, -1))); // -1 表示没有计算过
        auto dfs = [&](auto& dfs, int i, int left1, int left_s) -> int {
            if (i < 0) {
                return left_s == 0;
            }
            int& res = memo[i][left1][left_s]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = 0;
            int c = cnt[i] - (i ? cnt[i - 1] : 0);
            int left2 = cnt[i] - left1;
            for (int k = max(c - left2, 0); k <= min(c, left1) && k * i <= left_s; k++) {
                int r = dfs(dfs, i - 1, left1 - k, left_s - k * i);
                res = (res + r * INV_F[k] % MOD * INV_F[c - k]) % MOD;
            }
            return res;
        };
        return F[n1] * F[n - n1] % MOD * dfs(dfs, 9, n1, total / 2) % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 40

var fac, invF [mx + 1]int

func init() {
	fac[0] = 1
	for i := 1; i <= mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF[mx] = pow(fac[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func countBalancedPermutations(num string) int {
	cnt := [10]int{}
	total := 0
	for _, c := range num {
		cnt[c-'0']++
		total += int(c - '0')
	}

	if total%2 > 0 {
		return 0
	}

	for i := 1; i < 10; i++ {
		cnt[i] += cnt[i-1]
	}

	n := len(num)
	n1 := n / 2
	memo := [10][][]int{}
	for i := range memo {
		memo[i] = make([][]int, n1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, total/2+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, left1, leftS int) int {
		if i < 0 {
			if leftS > 0 {
				return 0
			}
			return 1
		}
		p := &memo[i][left1][leftS]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		c := cnt[i]
		if i > 0 {
			c -= cnt[i-1]
		}
		left2 := cnt[i] - left1
		for k := max(c-left2, 0); k <= min(c, left1) && k*i <= leftS; k++ {
			r := dfs(i-1, left1-k, leftS-k*i)
			res = (res + r*invF[k]%mod*invF[c-k]) % mod
		}
		*p = res // 记忆化
		return res
	}
	return fac[n1] * fac[n-n1] % mod * dfs(9, n1, total/2) % mod
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

- 时间复杂度：$\mathcal{O}(n^2S)$，其中 $n$ 为 $\textit{num}$ 的长度，$S$ 为 $\textit{num}$ 的数字和的一半，这不超过 $9n/2$。注意本题要把状态 $i$ 和枚举 $k$ 结合起来看，这二者一起是 $\mathcal{O}(n)$ 的。 
- 空间复杂度：$\mathcal{O}(DnS)$，其中 $D=10$。保存多少状态，就需要多少空间。

## 1:1 翻译成递推 + 空间优化

同上，定义 $f[i+1][\textit{left}_1][\textit{leftS}]$ 表示在剩余要分配的数字是 $[0,i]$，第一个多重集还剩下 $\textit{left}_1$ 个数字需要分配，第一个多重集还剩下 $\textit{leftS}$ 的元素和需要分配的情况下，下式的结果：

$$
\sum_{k_i=0}^{\textit{cnt}[i]} f_i(k_0,k_1,\ldots,k_i)
$$

递推式

$$
f[i+1][\textit{left}_1][\textit{leftS}] = \sum_{k=0}^{\textit{cnt}[i]} f[i][\textit{left}_1 - k][\textit{leftS} - k\cdot i]\cdot \dfrac{1}{k!(\textit{cnt}[i]-k)!}
$$

初始值：$f[0][0][0] = 1$，其余 $f[0][0][\textit{leftS}]=0$。

最终答案为

$$
n_1!\times (n-n_1)!\times f[10][n_1][\textit{total}/2]
$$

代码实现时，类似 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/)，去掉第一个维度，倒序循环 $\textit{left}_1$ 和 $\textit{leftS}$。

⚠**注意**：递推会计算一些无效状态，不一定比记忆化搜索快。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 41

fac = [0] * MX  # f[i] = i!
fac[0] = 1
for i in range(1, MX):
    fac[i] = fac[i - 1] * i % MOD

inv_f = [0] * MX  # inv_f[i] = i!^-1
inv_f[-1] = pow(fac[-1], -1, MOD)
for i in range(MX - 1, 0, -1):
    inv_f[i - 1] = inv_f[i] * i % MOD

class Solution:
    def countBalancedPermutations(self, num: str) -> int:
        cnt = [0] * 10
        total = 0
        for c in map(int, num):
            cnt[c] += 1
            total += c

        if total % 2:
            return 0

        n = len(num)
        n1 = n // 2
        f = [[0] * (total // 2 + 1) for _ in range(n1 + 1)]
        f[0][0] = 1
        sc = s = 0
        for i, c in enumerate(cnt):
            sc += c
            s += c * i
            # 保证 left2 <= n-n1，即 left1 >= sc-(n-n1)
            for left1 in range(min(sc, n1), max(sc - (n - n1) - 1, -1), -1):
                left2 = sc - left1
                # 保证分给第二个集合的元素和 <= total/2，即 left_s >= s-total/2
                for left_s in range(min(s, total // 2), max(s - total // 2 - 1, -1), -1):
                    res = 0
                    for k in range(max(c - left2, 0), min(c, left1) + 1):
                        if k * i > left_s:
                            break
                        res += f[left1 - k][left_s - k * i] * inv_f[k] * inv_f[c - k]
                    f[left1][left_s] = res % MOD
        return fac[n1] * fac[n - n1] * f[n1][total // 2] % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 41;

    private static final long[] F = new long[MX]; // f[i] = i!
    private static final long[] INV_F = new long[MX]; // inv_f[i] = i!^-1

    static {
        F[0] = 1;
        for (int i = 1; i < MX; i++) {
            F[i] = F[i - 1] * i % MOD;
        }
        INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
        for (int i = MX - 1; i > 0; i--) {
            INV_F[i - 1] = INV_F[i] * i % MOD;
        }
    }

    public int countBalancedPermutations(String num) {
        int[] cnt = new int[10];
        int total = 0;
        for (char c : num.toCharArray()) {
            cnt[c - '0']++;
            total += c - '0';
        }

        if (total % 2 > 0) {
            return 0;
        }

        int n = num.length();
        int n1 = n / 2;
        int[][] f = new int[n1 + 1][total / 2 + 1];
        f[0][0] = 1;
        int sc = 0;
        int s = 0;
        for (int i = 0; i < 10; i++) {
            int c = cnt[i];
            sc += c;
            s += c * i;
            // 保证 left2 <= n-n1，即 left1 >= sc-(n-n1)
            for (int left1 = Math.min(sc, n1); left1 >= Math.max(sc - (n - n1), 0); left1--) {
                int left2 = sc - left1;
                // 保证分给第二个集合的元素和 <= total/2，即 leftS >= s-total/2
                for (int leftS = Math.min(s, total / 2); leftS >= Math.max(s - total / 2, 0); leftS--) {
                    long res = 0;
                    for (int k = Math.max(c - left2, 0); k <= Math.min(c, left1) && k * i <= leftS; k++) {
                        res = (res + f[left1 - k][leftS - k * i] * INV_F[k] % MOD * INV_F[c - k]) % MOD;
                    }
                    f[left1][leftS] = (int) res;
                }
            }
        }
        return (int) (F[n1] * F[n - n1] % MOD * f[n1][total / 2] % MOD);
    }

    private static long pow(long x, int n) {
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
const int MOD = 1'000'000'007;
const int MX = 41;

long long F[MX]; // F[i] = i!
long long INV_F[MX]; // INV_F[i] = i!^-1

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

auto init = [] {
    F[0] = 1;
    for (int i = 1; i < MX; i++) {
        F[i] = F[i - 1] * i % MOD;
    }
    INV_F[MX - 1] = pow(F[MX - 1], MOD - 2);
    for (int i = MX - 1; i; i--) {
        INV_F[i - 1] = INV_F[i] * i % MOD;
    }
    return 0;
}();

class Solution {
public:
    int countBalancedPermutations(string num) {
        int cnt[10];
        int total = 0;
        for (char c : num) {
            cnt[c - '0']++;
            total += c - '0';
        }

        if (total % 2) {
            return 0;
        }

        int n = num.size();
        int n1 = n / 2;
        vector<vector<int>> f(n1 + 1, vector<int>(total / 2 + 1));
        f[0][0] = 1;
        int sc = 0, s = 0;
        for (int i = 0; i < 10; i++) {
            int c = cnt[i];
            sc += c;
            s += c * i;
            // 保证 left2 <= n-n1，即 left1 >= sc-(n-n1)
            for (int left1 = min(sc, n1); left1 >= max(sc - (n - n1), 0); left1--) {
                int left2 = sc - left1;
                // 保证分给第二个集合的元素和 <= total/2，即 leftS >= s-total/2
                for (int left_s = min(s, total / 2); left_s >= max(s - total / 2, 0); left_s--) {
                    int res = 0;
                    for (int k = max(c - left2, 0); k <= min(c, left1) && k * i <= left_s; k++) {
                        res = (res + f[left1 - k][left_s - k * i] * INV_F[k] % MOD * INV_F[c - k]) % MOD;
                    }
                    f[left1][left_s] = res;
                }
            }
        }
        return F[n1] * F[n - n1] % MOD * f[n1][total / 2] % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 40

var fac, invF [mx + 1]int

func init() {
	fac[0] = 1
	for i := 1; i <= mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF[mx] = pow(fac[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func countBalancedPermutations(num string) int {
	cnt := [10]int{}
	total := 0
	for _, c := range num {
		cnt[c-'0']++
		total += int(c - '0')
	}

	if total%2 > 0 {
		return 0
	}

	n := len(num)
	n1 := n / 2
	f := make([][]int, n1+1)
	for i := range f {
		f[i] = make([]int, total/2+1)
	}
	f[0][0] = 1
	sc := 0
	s := 0
	for i, c := range cnt {
		sc += c
		s += c * i
		// 保证 left2 <= n-n1，即 left1 >= sc-(n-n1)
		for left1 := min(sc, n1); left1 >= max(sc-(n-n1), 0); left1-- {
			left2 := sc - left1
			// 保证分给第二个集合的元素和 <= total/2，即 leftS >= s-total/2
			for leftS := min(s, total/2); leftS >= max(s-total/2, 0); leftS-- {
				res := 0
				for k := max(c-left2, 0); k <= min(c, left1) && k*i <= leftS; k++ {
					res = (res + f[left1-k][leftS-k*i]*invF[k]%mod*invF[c-k]) % mod
				}
				f[left1][leftS] = res
			}
		}
	}
	return fac[n1] * fac[n-n1] % mod * f[n1][total/2] % mod
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

- 时间复杂度：$\mathcal{O}(n^2S)$，其中 $n$ 为 $\textit{num}$ 的长度，$S$ 为 $\textit{num}$ 的数字和的一半，这不超过 $9n/2$。注意把 $i$ 和 $k$ 结合起来看，这二者一起是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(nS)$。保存多少状态，就需要多少空间。

更多相似题目，见下面动态规划题单中的「**§7.5 多维 DP**」和数学题单中的「**§2.2 组合计数**」

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
