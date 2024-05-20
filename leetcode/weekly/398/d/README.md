## 方法一：记忆化搜索

根据题意，我们需要三个参数来表示当前的状态：

- $i$：当前位于台阶 $i$。
- $j$：已经使用了 $j$ 次第二种操作。
- $\textit{preDown}$：上一次操作是否使用了操作一。

将其定义为 $\textit{dfs}(i,j,\textit{preDown})$，表示在该状态下，有多少种方案可以到达台阶 $k$。

枚举当前使用哪个操作：

- 使用操作一：前提是 $\textit{preDown}=\texttt{false}$ 且 $i>0$。使用操作一后，要解决的子问题是 $\textit{dfs}(i-1,j,\texttt{true})$，加入返回值中。
- 使用操作二：要解决的子问题是 $\textit{dfs}(i+2^j,j+1,\texttt{false})$，加入返回值中。
- 此外，如果 $i=k$，可以把返回值加一。

递归边界：如果 $i>k+1$，由于操作一不能连续使用，无法到达 $k$，返回 $0$。

递归入口：$\textit{dfs}(1,0,\texttt{false})$，即答案。

这个递归过程可以加上记忆化，原理请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)。

[视频讲解](https://www.bilibili.com/video/BV19D421G7mw/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def waysToReachStair(self, k: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, pre_down: bool) -> int:
            if i > k + 1:
                return 0
            res = 1 if i == k else 0
            res += dfs(i + (1 << j), j + 1, False)  # 操作二
            if i and not pre_down:
                res += dfs(i - 1, j, True)  # 操作一
            return res
        return dfs(1, 0, False)
```

```java [sol-Java]
class Solution {
    public int waysToReachStair(int k) {
        return dfs(1, 0, 0, k, new HashMap<>());
    }

    private int dfs(int i, int j, int preDown, int k, Map<Long, Integer> memo) {
        if (i > k + 1) {
            return 0;
        }
        long p = ((long) i << 32) | j << 1 | preDown; // 用一个 long 表示状态
        if (memo.containsKey(p)) { // 之前算过了
            return memo.get(p);
        }
        int res = i == k ? 1 : 0;
        res += dfs(i + (1 << j), j + 1, 0, k, memo); // 操作二
        if (preDown == 0 && i > 0) {
            res += dfs(i - 1, j, 1, k, memo); // 操作一
        }
        memo.put(p, res); // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    unordered_map<long long, int> memo;

    int dfs(int i, int j, bool preDown, int k) {
        if (i > k + 1) {
            return 0;
        }
        long long p = (long long) i << 32 | j << 1 | preDown; // 用一个 long long 表示状态
        if (memo.contains(p)) { // 之前算过了
            return memo[p];
        }
        int res = i == k;
        res += dfs(i + (1 << j), j + 1, false, k); // 操作二
        if (i && !preDown) {
            res += dfs(i - 1, j, true, k); // 操作一
        }
        return memo[p] = res; // 记忆化
    };

public:
    int waysToReachStair(int k) {
        return dfs(1, 0, false, k);
    }
};
```

```go [sol-Go]
func waysToReachStair(k int) int {
	type args struct {
		i, j    int
		preDown bool
	}
	memo := map[args]int{}
	var dfs func(int, int, bool) int
	dfs = func(i, j int, preDown bool) int {
		if i > k+1 {
			return 0
		}
		p := args{i, j, preDown}
		if v, ok := memo[p]; ok { // 之前算过了
			return v
		}
		res := dfs(i+1<<j, j+1, false) // 操作二
		if !preDown && i > 0 {
			res += dfs(i-1, j, true) // 操作一
		}
		if i == k {
			res++
		}
		memo[p] = res // 记忆化
		return res
	}
	return dfs(1, 0, false)
}
```

#### 复杂度分析

有多少个状态？

对于 $j$ 来说，如果 $i+2^j > k + 1$，就不会再递归了。放缩一下，以「$2^j > k + 1$ 不再递归」来估计，只有 $\mathcal{O}(\log k)$ 个不同的 $j$。

对于固定的 $j$ 和 $\textit{preDown}$，有多少个不同的 $i$？

考虑 $j$ 对于 $i$ 的贡献，由于 $j$ 是一点一点增大的，所以 $i$ 一定包含 $2^0,2^1,2^2,\cdots,2^{j-1}$。从 $i$ 中减掉这些数的和，剩下的 $|i-1|$ 就是使用操作一的次数（注意 $i$ 从 $1$ 开始）。由于操作一不能连续使用，所以操作一的次数不超过 $j+1$。所以对于固定的 $j$ 和 $\textit{preDown}$，有 $\mathcal{O}(j)$ 个不同的 $i$。

综上所述，总共有 $\mathcal{O}(1+2+\cdots + \log k) = \mathcal{O}(\log^2 k)$ 个状态。

- 时间复杂度：$\mathcal{O}(\log^2 k)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\log^2 k)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\log^2 k)$。
- 空间复杂度：$\mathcal{O}(\log^2 k)$。保存多少状态，就需要多少空间。

上面的状态个数分析，也引出了下面的组合数学做法。

## 方法二：组合数学

假设使用了 $m$ 次操作一，$j$ 次操作二，那么有

$$
1 + 2^0 + 2^1 + 2^2 + \cdots + 2^{j-1} - m = k
$$

即

$$
m =  2^j - k
$$

注意上式当 $j=0$ 时也是成立的。

由于操作一不能连续使用，我们需要在这 $j$ 次操作二前后，以及相邻两次操作二的空隙中，插入 $m$ 个操作一，所以方案数等于从 $j+1$ 个物品中选出 $m$ 个物品的方案数，即

$$
\binom {j+1} m
$$

枚举 $j$，最终答案为

$$
\sum_{j=0}^{29} \binom {j+1} m
$$

其中 $0\le m \le j+1$。根据题目的数据范围，$j$ 至多枚举到 $29$。

预处理组合数后可以快速计算。如何递推计算组合数请看 [视频讲解](https://www.bilibili.com/video/BV19D421G7mw/) 第四题。

```py [sol-Python3]
class Solution:
    def waysToReachStair(self, k: int) -> int:
        ans = 0
        for j in range(30):
            m = (1 << j) - k
            if 0 <= m <= j + 1:
                ans += comb(j + 1, m)
        return ans
```

```py [sol-Python3 写法二]
comb = cache(comb)  # 记忆化

class Solution:
    def waysToReachStair(self, k: int) -> int:
        ans = 0
        for j in range(30):
            m = (1 << j) - k
            if 0 <= m <= j + 1:
                ans += comb(j + 1, m)
        return ans
```

```java [sol-Java]
public class Solution {
    private static final int MX = 31;
    private static final int[][] c = new int[MX][MX];

    static {
        for (int i = 0; i < MX; i++) {
            c[i][0] = c[i][i] = 1;
            for (int j = 1; j < i; j++) {
                c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
            }
        }
    }

    public int waysToReachStair(int k) {
        int ans = 0;
        for (int j = 0; j < 30; j++) {
            int m = (1 << j) - k;
            if (0 <= m && m <= j + 1) {
                ans += c[j + 1][m];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
int c[31][31];
auto init = [] {
    for (int i = 0; i < 31; i++) {
        c[i][0] = c[i][i] = 1;
        for (int j = 1; j < i; j++) {
            c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
        }
    }
    return 0;
}();

class Solution {
public:
    int waysToReachStair(int k) {
        int ans = 0;
        for (int j = 0; j < 30; j++) {
            int m = (1 << j) - k;
            if (0 <= m && m <= j + 1) {
                ans += c[j + 1][m];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 31
var c [mx][mx]int
func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := 0; j < 30; j++ {
		m := 1<<j - k
		if 0 <= m && m <= j+1 {
			ans += c[j+1][m]
		}
	}
	return
}
```

## 优化

只需要从首个满足 $2^j \ge k$ 的 $j$ 开始枚举，相当于 $j$ 至少是 $\max(k-1,0)$ 的二进制长度（$0$ 的二进制长度为 $0$）。

当 $2^j-k > j + 1$ 时，停止循环。

经测试，$k=1$ 时要循环 $3$ 次，$k=0,2,4$ 时要循环 $2$ 次，其余 $k$ 至多循环 $1$ 次。（几乎所有 $k$ 都不需要循环，返回 $0$。）

```py [sol-Python3]
class Solution:
    def waysToReachStair(self, k: int) -> int:
        ans = 0
        for j in count(max(k - 1, 0).bit_length()):
            m = (1 << j) - k
            if m > j + 1:
                break
            ans += comb(j + 1, m)
        return ans
```

```py [sol-Python3 写法二]
comb = cache(comb)  # 记忆化

class Solution:
    def waysToReachStair(self, k: int) -> int:
        ans = 0
        for j in count(max(k - 1, 0).bit_length()):
            m = (1 << j) - k
            if m > j + 1:
                break
            ans += comb(j + 1, m)
        return ans
```

```java [sol-Java]
public class Solution {
    private static final int MX = 31;
    private static final int[][] c = new int[MX][MX];

    static {
        for (int i = 0; i < MX; i++) {
            c[i][0] = c[i][i] = 1;
            for (int j = 1; j < i; j++) {
                c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
            }
        }
    }

    public int waysToReachStair(int k) {
        int ans = 0;
        for (int j = 32 - Integer.numberOfLeadingZeros(Math.max(k - 1, 0)); (1 << j) - k <= j + 1; j++) {
            ans += c[j + 1][(1 << j) - k];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
int c[31][31];
auto init = [] {
    for (int i = 0; i < 31; i++) {
        c[i][0] = c[i][i] = 1;
        for (int j = 1; j < i; j++) {
            c[i][j] = c[i - 1][j - 1] + c[i - 1][j];
        }
    }
    return 0;
}();

class Solution {
public:
    int waysToReachStair(int k) {
        int ans = 0;
        for (int j = k > 1 ? 32 - __builtin_clz(k - 1) : 0; (1 << j) - k <= j + 1; j++) {
            ans += c[j + 1][(1 << j) - k];
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 31
var c [mx][mx]int
func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。预处理的时间和空间不计入。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

为什么返回值是 `int` 而不是 `long long`？答案最大是多少？

见本题视频讲解。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
