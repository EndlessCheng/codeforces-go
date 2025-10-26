## 方法一：动态规划

### 寻找子问题

比如 $n-1$ 行选了 $6$，问题变成在已选元素的 GCD 为 $6$ 的情况下，从 $[0,n-2]$ 每行选一个数，最终 GCD 等于 $1$ 的方案数。

如果 $n-2$ 行选了 $4$，问题变成在已选元素的 GCD 为 $\gcd(6,4)=2$ 的情况下，从 $[0,n-3]$ 每行选一个数，最终 GCD 等于 $1$ 的方案数。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

### 状态定义与状态转移方程

定义 $\textit{dfs}(i,g)$ 表示在已选元素的 GCD 为 $g$ 的情况下，从 $[0,i]$ 每行选一个数，最终 GCD 等于 $1$ 的方案数。

枚举从 $\textit{mat}[i]$ 这一行选 $x$，问题变成在已选元素的 GCD 为 $\gcd(g,x)$ 的情况下，从 $[0,i-1]$ 每行选一个数，最终 GCD 等于 $1$ 的方案数，即 $\textit{dfs}(i-1, \gcd(g,x))$。

这一行要从 $n$ 个元素中选一个，这 $n$ 种情况互斥，那么根据**加法原理**，有

$$
\textit{dfs}(i,g) = \sum_{j=0}^{n-1} \textit{dfs}(i-1, \gcd(g,\textit{mat}[i][j]))
$$

**递归边界**：$\textit{dfs}(-1,1)=1$，其余 $\textit{dfs}(-1,g)=0$。

**递归入口**：$\textit{dfs}(m-1,0)$，这是原问题，也是答案。$g$ 初始化成 $0$ 是因为 $0$ 和任何数 $x$ 的 GCD 都等于 $x$。

### 递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

⚠**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,g)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1zxxNzcERu/?t=13m51s)，两种方法都讲了，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countCoprime(self, mat: List[List[int]]) -> int:
        MOD = 1_000_000_007

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, g: int) -> int:
            if i < 0:
                return 1 if g == 1 else 0
            return sum(dfs(i - 1, gcd(g, x)) for x in mat[i]) % MOD

        return dfs(len(mat) - 1, 0)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countCoprime(int[][] mat) {
        int m = mat.length;
        int mx = 0;
        for (int[] row : mat) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }

        int[][] memo = new int[m][mx + 1];
        for (int i = 0; i < m; i++) {
            Arrays.fill(memo[i], -1); // -1 表示没有计算过
        }
        return dfs(m - 1, 0, mat, memo);
    }

    private int dfs(int i, int g, int[][] mat, int[][] memo) {
        if (i < 0) {
            return g == 1 ? 1 : 0;
        }
        if (memo[i][g] != -1) { // 之前计算过
            return memo[i][g];
        }
        long res = 0;
        for (int x : mat[i]) {
            res += dfs(i - 1, gcd(g, x), mat, memo);
        }
        return memo[i][g] = (int) (res % MOD); // 记忆化
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countCoprime(vector<vector<int>>& mat) {
        constexpr int MOD = 1'000'000'007;
        int m = mat.size();
        int mx = 0;
        for (auto& row : mat) {
            mx = max(mx, ranges::max(row));
        }

        vector memo(m, vector<int>(mx + 1, -1));
        auto dfs = [&](this auto&& dfs, int i, int g) -> int {
            if (i < 0) {
                return g == 1;
            }
            int& res = memo[i][g]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            res = 0;
            for (int x : mat[i]) {
                res = (res + dfs(i - 1, gcd(g, x))) % MOD;
            }
            return res;
        };
        return dfs(m - 1, 0);
    }
};
```

```go [sol-Go]
func countCoprime(mat [][]int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range mat {
		mx = max(mx, slices.Max(row))
	}

	m := len(mat)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, mx+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, g int) (res int) {
		if i < 0 {
			if g == 1 {
				return 1
			}
			return
		}
		p := &memo[i][g]
		if *p != -1 { // 之前计算过
			return *p
		}
		for _, x := range mat[i] {
			res += dfs(i-1, gcd(g, x))
		}
		res %= mod
		*p = res // 记忆化
		return
	}
	return dfs(m-1, 0)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU\log U)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数，$U$ 是 $\textit{mat}[i][j]$ 的最大值。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mU)$，单个状态的计算时间为 $\mathcal{O}(n\log U)$，所以总的时间复杂度为 $\mathcal{O}(mnU\log U)$。
- 空间复杂度：$\mathcal{O}(mU)$。保存多少状态，就需要多少空间。

### 1:1 翻译成递推

由于递推会计算很多无效状态，而记忆化搜索不会，所以递推可能比记忆化搜索慢。

```py [sol-Python3]
class Solution:
    def countCoprime(self, mat: List[List[int]]) -> int:
        MOD = 1_000_000_007
        m = len(mat)
        mx = max(map(max, mat))
        f = [[0] * (mx + 1) for _ in range(m + 1)]
        f[0][1] = 1
        for i, row in enumerate(mat):
            for g in range(mx + 1):
                f[i + 1][g] = sum(f[i][gcd(g, x)] for x in row) % MOD
        return f[m][0]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countCoprime(int[][] mat) {
        int m = mat.length;
        int mx = 0;
        for (int[] row : mat) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }

        int[][] f = new int[m + 1][mx + 1];
        f[0][1] = 1;
        for (int i = 0; i < m; i++) {
            for (int g = 0; g <= mx; g++) {
                long res = 0;
                for (int x : mat[i]) {
                    res += f[i][gcd(g, x)];
                }
                f[i + 1][g] = (int) (res % MOD);
            }
        }
        return f[m][0];
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countCoprime(vector<vector<int>>& mat) {
        constexpr int MOD = 1'000'000'007;
        int m = mat.size();
        int mx = 0;
        for (auto& row : mat) {
            mx = max(mx, ranges::max(row));
        }

        vector f(m + 1, vector<int>(mx + 1));
        f[0][1] = 1;
        for (int i = 0; i < m; i++) {
            for (int g = 0; g <= mx; g++) {
                long long res = 0;
                for (int x : mat[i]) {
                    res += f[i][gcd(g, x)];
                }
                f[i + 1][g] = res % MOD;
            }
        }
        return f[m][0];
    }
};
```

```go [sol-Go]
func countCoprime(mat [][]int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range mat {
		mx = max(mx, slices.Max(row))
	}

	m := len(mat)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, mx+1)
	}
	f[0][1] = 1
	for i, row := range mat {
		for g := 0; g <= mx; g++ {
			res := 0
			for _, x := range row {
				res += f[i][gcd(g, x)]
			}
			f[i+1][g] = res % mod
		}
	}
	return f[m][0]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU\log U)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数，$U$ 是 $\textit{mat}[i][j]$ 的最大值。
- 空间复杂度：$\mathcal{O}(mU)$。

**注**：还可以在类外预处理所有 $\gcd(i,j)$ 的值，加快速度。

## 方法二：倍数容斥

**前置题目**：[3312. 查询排序后的最大公约数](https://leetcode.cn/problems/sorted-gcd-pair-queries/)，[我的题解](https://leetcode.cn/problems/sorted-gcd-pair-queries/solutions/2940415/mei-ju-rong-chi-qian-zhui-he-er-fen-pyth-ujis/)。

对于本题，枚举 $i$，我们需要计算从每行选一个 $i$ 的倍数的方案数。其余逻辑同 3312 题。

**小优化**：如果遍历过程中发现某一行没有 $i$ 的倍数，可以提前跳出循环。

### 优化前

```py [sol-Python3]
class Solution:
    def countCoprime(self, mat: List[List[int]]) -> int:
        MOD = 1_000_000_007
        mx = max(map(max, mat))

        cnt_gcd = [0] * (mx + 1)
        for i in range(mx, 0, -1):
            # 每行选一个 i 的倍数的方案数
            res = 1
            for row in mat:
                cnt = 0
                for x in row:
                    if x % i == 0:
                        cnt += 1
                if cnt == 0:
                    res = 0
                    break
                res = res * cnt % MOD  # 乘法原理

            for j in range(i, mx + 1, i):
                res -= cnt_gcd[j]
            cnt_gcd[i] = res % MOD

        return cnt_gcd[1]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countCoprime(int[][] mat) {
        int mx = 0;
        for (int[] row : mat) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }

        int[] cntGcd = new int[mx + 1];
        for (int i = mx; i > 0; i--) {
            // 每行选一个 i 的倍数的方案数
            long res = 1;
            for (int[] row : mat) {
                int cnt = 0;
                for (int x : row) {
                    if (x % i == 0) {
                        cnt++;
                    }
                }
                if (cnt == 0) {
                    res = 0;
                    break;
                }
                res = res * cnt % MOD; // 乘法原理
            }

            for (int j = i; j <= mx; j += i) {
                res -= cntGcd[j]; // 注意这里有减法，可能导致 res 是负数
            }

            cntGcd[i] = (int) (res % MOD);
        }
        return (cntGcd[1] + MOD) % MOD; // 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countCoprime(vector<vector<int>>& mat) {
        constexpr int MOD = 1'000'000'007;
        int mx = 0;
        for (auto& row : mat) {
            mx = max(mx, ranges::max(row));
        }

        vector<int> cnt_gcd(mx + 1);
        for (int i = mx; i > 0; i--) {
            // 每行选一个 i 的倍数的方案数
            long long res = 1;
            for (auto& row : mat) {
                int cnt = 0;
                for (int x : row) {
                    if (x % i == 0) {
                        cnt++;
                    }
                }
                if (cnt == 0) {
                    res = 0;
                    break;
                }
                res = res * cnt % MOD; // 乘法原理
            }

            for (int j = i; j <= mx; j += i) {
                res -= cnt_gcd[j]; // 注意这里有减法，可能导致 res 是负数
            }

            cnt_gcd[i] = res % MOD;
        }
        return (cnt_gcd[1] + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
func countCoprime(mat [][]int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range mat {
		mx = max(mx, slices.Max(row))
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		// 每行选一个 i 的倍数的方案数
		res := 1
		for _, row := range mat {
			cnt := 0
			for _, x := range row {
				if x%i == 0 {
					cnt++
				}
			}
			if cnt == 0 {
				res = 0
				break
			}
			res = res * cnt % mod // 乘法原理
		}

		for j := i; j <= mx; j += i {
			res -= cntGcd[j] // 注意这里有减法，可能导致 res 是负数
		}

		cntGcd[i] = res % mod
	}
	return (cntGcd[1] + mod) % mod // 保证结果非负
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU+U\log U)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数，$U$ 是 $\textit{mat}[i][j]$ 的最大值。
- 空间复杂度：$\mathcal{O}(U)$。

### 优化

`if x % i == 0: cnt += 1` 本质是在统计 $x$ 的因子 $i$ 的个数。

所以对于每一行，可以先**预处理**每个数的因子的出现次数，避免反复遍历 $\textit{mat}$。

```py [sol-Python3]
MX = 151
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):
        divisors[j].append(i)

class Solution:
    def countCoprime(self, mat: List[List[int]]) -> int:
        MOD = 1_000_000_007
        # 预处理每行的因子个数
        divisor_cnt = []
        mx = 0
        for row in mat:
            row_max = max(row)
            mx = max(mx, row_max)
            cnt = [0] * (row_max + 1)
            for x in row:
                for d in divisors[x]:
                    cnt[d] += 1
            divisor_cnt.append(cnt)

        cnt_gcd = [0] * (mx + 1)
        for i in range(mx, 0, -1):
            # 每行选一个 i 的倍数的方案数
            res = 1
            for cnt in divisor_cnt:
                if i >= len(cnt) or cnt[i] == 0:
                    res = 0
                    break
                res = res * cnt[i] % MOD  # 乘法原理

            for j in range(i, mx + 1, i):
                res -= cnt_gcd[j]
            cnt_gcd[i] = res % MOD

        return cnt_gcd[1]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 151;
    private static final List<Integer>[] divisors = new ArrayList[MX];

    static {
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                divisors[j].add(i); // i 是 j 的因子
            }
        }
    }

    public int countCoprime(int[][] mat) {
        // 预处理每行的因子个数
        int m = mat.length;
        int[][] divisorCnt = new int[m][];
        int mx = 0;
        for (int i = 0; i < m; i++) {
            int[] row = mat[i];
            int rowMax = 0;
            for (int x : row) {
                rowMax = Math.max(rowMax, x);
            }
            mx = Math.max(mx, rowMax);

            int[] cnt = new int[rowMax + 1];
            for (int x : mat[i]) {
                for (int d : divisors[x]) {
                    cnt[d]++;
                }
            }
            divisorCnt[i] = cnt;
        }

        int[] cntGcd = new int[mx + 1];
        for (int i = mx; i > 0; i--) {
            // 每行选一个 i 的倍数的方案数
            long res = 1;
            for (int[] cnt : divisorCnt) {
                if (i >= cnt.length || cnt[i] == 0) {
                    res = 0;
                    break;
                }
                res = res * cnt[i] % MOD; // 乘法原理
            }

            for (int j = i; j <= mx; j += i) {
                res -= cntGcd[j]; // 注意这里有减法，可能导致 res 是负数
            }

            cntGcd[i] = (int) (res % MOD);
        }
        return (cntGcd[1] + MOD) % MOD; // 保证结果非负
    }
}
```

```cpp [sol-C++]
constexpr int MX = 151;
vector<int> divisors[MX];

int init = [] {
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) {
            divisors[j].push_back(i);
        }
    }
    return 0;
}();

class Solution {
public:
    int countCoprime(vector<vector<int>>& mat) {
        constexpr int MOD = 1'000'000'007;
        // 预处理每行的因子个数
        int m = mat.size();
        vector<vector<int>> divisor_cnt(m);
        int mx = 0;
        for (int i = 0; i < m; i++) {
            int row_max = ranges::max(mat[i]);
            mx = max(mx, row_max);
            divisor_cnt[i].resize(row_max + 1);
            for (int x : mat[i]) {
                for (int d : divisors[x]) {
                    divisor_cnt[i][d]++;
                }
            }
        }

        vector<int> cnt_gcd(mx + 1);
        for (int i = mx; i > 0; i--) {
            // 每行选一个 i 的倍数的方案数
            long long res = 1;
            for (auto& cnt : divisor_cnt) {
                if (i >= cnt.size() || cnt[i] == 0) {
                    res = 0;
                    break;
                }
                res = res * cnt[i] % MOD; // 乘法原理
            }

            for (int j = i; j <= mx; j += i) {
                res -= cnt_gcd[j]; // 注意这里有减法，可能导致 res 是负数
            }

            cnt_gcd[i] = res % MOD;
        }
        return (cnt_gcd[1] + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
const maxVal = 151
var divisors [maxVal][]int

func init() {
	for i := 1; i < maxVal; i++ {
		for j := i; j < maxVal; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func countCoprime(mat [][]int) int {
	const mod = 1_000_000_007
	// 预处理每行的因子个数
	divisorCnt := make([][]int, len(mat))
	mx := 0
	for i, row := range mat {
		rowMax := slices.Max(row)
		mx = max(mx, rowMax)
		divisorCnt[i] = make([]int, rowMax+1)
		for _, x := range row {
			for _, d := range divisors[x] {
				divisorCnt[i][d]++
			}
		}
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		// 每行选一个 i 的倍数的方案数
		res := 1
		for _, cnt := range divisorCnt {
			if i >= len(cnt) || cnt[i] == 0 {
				res = 0
				break
			}
			res = res * cnt[i] % mod // 乘法原理
		}

		for j := i; j <= mx; j += i {
			res -= cntGcd[j] // 注意这里有减法，可能导致 res 是负数
		}

		cntGcd[i] = res % mod
	}
	return (cntGcd[1] + mod) % mod // 保证结果非负
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(mn\cdot d(U) + mU + U\log U)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数，$U$ 是 $\textit{mat}[i][j]$ 的最大值，$d(U)\le 16$ 是单个数的最大因子个数。
- 空间复杂度：$\mathcal{O}(mU)$。

**注**：在统计因子时，可以先统计 $\textit{row}$ 每个元素的出现次数，这样每行可以做到 $\mathcal{O}(n + U\log U)$ 时间，总体时间复杂度为 $\mathcal{O}(m(n + U\log U))$。对于本题的数据范围，这个优化不明显。

## 专题训练

1. 下面动态规划题单的「**§7.6 多维 DP**」。
2. 下面数学题单的「**§1.6 最大公约数（GCD）**」。

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
