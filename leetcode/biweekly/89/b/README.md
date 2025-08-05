## 方法一：暴力

把 $n$ 分解为若干不同的 $2$ 的幂之和，把这些 $2$ 的幂添加到数组 $\textit{powers}$ 中。示例 1 的 $n=15=1+2+4+8$，分解得到的 $\textit{powers}=[1,2,4,8]$。

如何快速分解？这可以用 $\text{lowbit}$ 计算，详见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

对于每个询问，直接遍历 $\textit{powers}$ 的子数组，计算元素积。

注意取模，为什么可以在中途取模，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        MOD = 1_000_000_007
        # 例如二进制 1100 分解为 100 + 1000
        # 第一轮循环 lowbit(1100) = 100，然后 1100 ^ 100 = 1000
        # 第二轮循环 lowbit(1000) = 1000，然后 1000 ^ 1000 = 0，循环结束
        powers = []
        while n:
            lowbit = n & -n
            powers.append(lowbit)
            n ^= lowbit
        return [reduce(mul, powers[l: r + 1]) % MOD for l, r in queries]
```

```java [sol-Java]
class Solution {
    public int[] productQueries(int n, int[][] queries) {
        final int MOD = 1_000_000_007;
        // 例如二进制 1100 分解为 100 + 1000
        // 第一轮循环 lowbit(1100) = 100，然后 1100 ^ 100 = 1000
        // 第二轮循环 lowbit(1000) = 1000，然后 1000 ^ 1000 = 0，循环结束
        List<Integer> powers = new ArrayList<>();
        while (n > 0) {
            int lowbit = n & -n;
            powers.add(lowbit);
            n ^= lowbit;
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            long mul = 1;
            for (int j = q[0]; j <= q[1]; j++) {
                mul = mul * powers.get(j) % MOD;
            }
            ans[i] = (int) mul;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> productQueries(int n, vector<vector<int>>& queries) {
        const int MOD = 1'000'000'007;
        // 例如二进制 1100 分解为 100 + 1000
        // 第一轮循环 lowbit(1100) = 100，然后 1100 ^ 100 = 1000
        // 第二轮循环 lowbit(1000) = 1000，然后 1000 ^ 1000 = 0，循环结束
        vector<int> powers;
        while (n) {
            int lowbit = n & -n;
            powers.push_back(lowbit);
            n ^= lowbit;
        }

        vector<int> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            long long mul = 1;
            for (int j = q[0]; j <= q[1]; j++) {
                mul = mul * powers[j] % MOD;
            }
            ans.push_back(mul);
        }
        return ans;
    }
};
```

```go [sol-Go]
func productQueries(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	// 例如二进制 1100 分解为 100 + 1000
	// 第一轮循环 lowbit(1100) = 100，然后 1100 ^ 100 = 1000
	// 第二轮循环 lowbit(1000) = 1000，然后 1000 ^ 1000 = 0，循环结束
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		mul := 1
		for _, x := range powers[q[0] : q[1]+1] {
			mul = mul * x % mod
		}
		ans[i] = mul
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log n)$，其中 $m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(\log n)$。返回值不计入。

## 方法二：预处理

注意 $\textit{powers}$ 的长度等于 $n$ 的二进制中的 $1$ 的个数，这不超过 $n$ 的二进制长度，即 $\mathcal{O}(\log n)$，所以只有 $\mathcal{O}(\log^2 n)$ 个不同的询问，远小于 $10^5$。

我们可以先预处理 $\textit{powers}$ 每个子数组的元素积，得到一个表，然后 $\mathcal{O}(1)$ 查表，回答每个询问。

```py [sol-Python3]
class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        MOD = 1_000_000_007
        powers = []
        while n:
            lowbit = n & -n
            powers.append(lowbit)
            n ^= lowbit

        m = len(powers)
        res = [[0] * m for _ in range(m)]
        for i, x in enumerate(powers):
            res[i][i] = x
            for j in range(i + 1, m):
                res[i][j] = res[i][j - 1] * powers[j] % MOD

        return [res[l][r] for l, r in queries]
```

```java [sol-Java]
class Solution {
    public int[] productQueries(int n, int[][] queries) {
        final int MOD = 1_000_000_007;
        List<Integer> powers = new ArrayList<>();
        while (n > 0) {
            int lowbit = n & -n;
            powers.add(lowbit);
            n ^= lowbit;
        }

        int m = powers.size();
        int[][] res = new int[m][m];
        for (int i = 0; i < m; i++) {
            res[i][i] = powers.get(i);
            for (int j = i + 1; j < m; j++) {
                res[i][j] = (int) ((long) res[i][j - 1] * powers.get(j) % MOD);
            }
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            ans[i] = res[q[0]][q[1]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> productQueries(int n, vector<vector<int>>& queries) {
        const int MOD = 1'000'000'007;
        vector<int> powers;
        while (n) {
            int lowbit = n & -n;
            powers.push_back(lowbit);
            n ^= lowbit;
        }

        int m = powers.size();
        vector res(m, vector<int>(m));
        for (int i = 0; i < m; i++) {
            res[i][i] = powers[i];
            for (int j = i + 1; j < m; j++) {
                res[i][j] = 1LL * res[i][j - 1] * powers[j] % MOD;
            }
        }

        vector<int> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            ans.push_back(res[q[0]][q[1]]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func productQueries(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	m := len(powers)
	res := make([][]int, m)
	for i, x := range powers {
		res[i] = make([]int, m)
		res[i][i] = x
		for j := i + 1; j < m; j++ {
			res[i][j] = res[i][j-1] * powers[j] % mod
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + \log^2 n)$，其中 $m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(\log^2 n)$。返回值不计入。

## 方法三：前缀和

比如 $\textit{powers}=[2^1,2^3,2^4,2^6]$，对应的幂次数组 $e=[1,3,4,6]$。计算 $\textit{powers}$ 子数组的元素积，等价于先计算 $e$ 中对应的子数组和，然后再计算 $2$ 的幂。比如子数组 $[2^3,2^4,2^6]$ 的元素积为 $2^3\times 2^4\times 2^6 = 2^{3+4+6}$，我们可以先计算幂次 $3+4+6=13$，再计算 $2^{13}$。

$e$ 的子数组和可以用**前缀和**快速计算，原理见 [前缀和讲解](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

为了快速计算 $2$ 的幂，可以预处理 $\textit{pow}_2$ 数组，其中 $\textit{pow}_2[i] = 2^i\bmod (10^9+7)$。

> **注**：由于本题 $n < 2^{30}$，幂次和不超过 $1+2+\dots+29 = 435$。

### 答疑

**问**：能不能直接计算前缀积？子数组元素积等于两个前缀积的商？

**答**：取模后不能直接相除，例如 $\dfrac{24}{8}\bmod 5 \ne \dfrac{24\bmod 5}{8\bmod 5} = \dfrac{4}{3}$，后者甚至不是一个整数。对于本题，可以计算分母的逆元，具体见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。但计算逆元时间复杂度要额外乘个 $\log$，效率不如预处理 $\textit{pow}_2$ 的写法。

```py [sol-Python3]
MOD = 1_000_000_007
MX = 436
pow2 = [0] * MX
pow2[0] = 1
for i in range(1, MX):
    pow2[i] = pow2[i - 1] * 2 % 1_000_000_007

class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        s = [0]
        while n:
            lowbit = n & -n
            e = lowbit.bit_length() - 1
            # 直接计算 e 的前缀和
            s.append(s[-1] + e)
            n ^= lowbit
        return [pow2[s[r + 1] - s[l]] for l, r in queries]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 436;
    private static final int[] pow2 = new int[MX];

    static {
        pow2[0] = 1;
        for (int i = 1; i < MX; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }
    }

    public int[] productQueries(int n, int[][] queries) {
        List<Integer> s = new ArrayList<>();
        s.add(0);
        for (; n > 0; n &= n - 1) { // n &= n-1 去掉 n 的最低比特 1
            int e = Integer.numberOfTrailingZeros(n);
            // 直接计算 e 的前缀和
            s.add(s.getLast() + e);
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            int sumE = s.get(q[1] + 1) - s.get(q[0]);
            ans[i] = pow2[sumE];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 436;
int pow2[MX] = {1};

int init = [] {
    for (int i = 1; i < MX; i++) {
        pow2[i] = pow2[i - 1] * 2 % MOD;
    }
    return 0;
}();

class Solution {
public:
    vector<int> productQueries(int n, vector<vector<int>>& queries) {
        vector<int> s = {0};
        for (; n > 0; n &= n - 1) { // n &= n-1 去掉 n 的最低比特 1
            int e = countr_zero((uint32_t) n);
            // 直接计算 e 的前缀和
            s.push_back(s.back() + e);
        }

        vector<int> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            int sum_e = s[q[1] + 1] - s[q[0]];
            ans.push_back(pow2[sum_e]);
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 436

var pow2 = [mx]int{1}

func init() {
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func productQueries(n int, queries [][]int) []int {
	s := []int{0}
	for ; n > 0; n &= n - 1 { // n &= n-1 去掉 n 的最低比特 1
		e := bits.TrailingZeros(uint(n))
		// 直接计算 e 的前缀和
		s = append(s, s[len(s)-1]+e)
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		sumE := s[q[1]+1] - s[q[0]]
		ans[i] = pow2[sumE]
	}
	return ans
}
```

#### 复杂度分析

预处理 $\textit{pow}_2$ 的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(m + \log n)$，其中 $m$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(\log n)$。返回值不计入。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
