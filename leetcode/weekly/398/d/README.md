## 方法一：记忆化搜索

### 寻找子问题

在跳台阶的过程中：

- 为了判断是否到达终点 $k$，我们需要知道当前在第几个台阶。
- 为了计算操作二能往上跳多少个台阶，我们需要知道已经使用了多少次操作二。
- 为了判断能否使用操作一（往下跳），我们需要知道上一次操作是否使用了操作一。

例如现在在第 $5$ 个台阶，且之前使用了 $3$ 次操作二，那么可以：

- 使用操作二：向上跳 $2^3=8$ 个台阶，到达第 $5+8=13$ 个台阶。接下来要解决的问题是：在使用了 $4$ 次操作二，且上一次操作没有使用操作一的情况下，从 $13$ 跳到 $k$ 的方案数。
- 使用操作一（前提是上一次操作没有使用操作一）：向下跳 $1$ 个台阶，到达第 $5-1=4$ 个台阶。接下来要解决的问题是：在使用了 $3$ 次操作二，且上一次操作使用了操作一的情况下，从 $4$ 跳到 $k$ 的方案数。

这些问题都是和原问题相似的子问题，可以用**递归**解决。

### 状态定义

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前位于第 $i$ 个台阶。
- $j$：已经使用了 $j$ 次操作二。
- $\textit{preDown}$：一个布尔值，表示上一次操作是否使用了操作一。

因此，定义状态为 $\textit{dfs}(i,j,\textit{preDown})$，表示在使用了 $j$ 次操作二，且上一次操作使用/未使用操作一的情况下，从 $i$ 跳到 $k$ 的方案数。

### 状态转移方程

枚举当前使用哪个操作：

- 使用操作二：接下来要解决的问题是 $\textit{dfs}(i+2^j,j+1,\texttt{false})$，将其方案数加到返回值中。
- 使用操作一（前提是 $\textit{preDown}=\texttt{false}$）：接下来要解决的问题是 $\textit{dfs}(i-1,j,\texttt{true})$，将其方案数加到返回值中。注意当 $\textit{preDown}=\texttt{false}$ 时，$i>0$ 必然成立。
- 此外，如果 $i=k$（到达终点），则找到了一个方案，把返回值加一。

⚠**注意**：到达第 $k$ 个台阶后，还可以继续操作，重新回到第 $k$ 个台阶。所以 $i=k$ 并不是递归边界。

**递归边界**：如果 $i>k+1$，由于操作一不能连续使用，无法到达 $k$，返回 $0$。

**递归入口**：$\textit{dfs}(1,0,\texttt{false})$，即答案。一开始在第 $1$ 个台阶，没有使用过操作二，也没有使用过操作一。

这个递归过程可以加上记忆化，原理请看 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)。

本题讲解见 [视频](https://www.bilibili.com/video/BV19D421G7mw/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def waysToReachStair(self, k: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, pre_down: bool) -> int:
            if i > k + 1:  # 无法到达终点 k
                return 0
            res = 1 if i == k else 0
            res += dfs(i + (1 << j), j + 1, False)  # 操作二
            if not pre_down:
                res += dfs(i - 1, j, True)  # 操作一
            return res
        return dfs(1, 0, False)
```

```java [sol-Java]
class Solution {
    public int waysToReachStair(int k) {
        return dfs(1, 0, 0, k, new HashMap<>());
    }

    // preDown = 0/1 表示 false/true
    private int dfs(int i, int j, int preDown, int k, Map<Long, Integer> memo) {
        if (i > k + 1) { // 无法到达终点 k
            return 0;
        }
        // 把状态 (i, j, preDown) 压缩成一个 long
        long mask = (long) i << 32 | j << 1 | preDown;
        if (memo.containsKey(mask)) { // 之前算过了
            return memo.get(mask);
        }
        int res = i == k ? 1 : 0;
        res += dfs(i + (1 << j), j + 1, 0, k, memo); // 操作二
        if (preDown == 0) {
            res += dfs(i - 1, j, 1, k, memo); // 操作一
        }
        memo.put(mask, res); // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int waysToReachStair(int k) {
        unordered_map<long long, int> memo;
        auto dfs = [&](auto&& dfs, int i, int j, bool pre_down) -> int {
            if (i > k + 1) { // 无法到达终点 k
                return 0;
            }
            // 把状态 (i, j, pre_down) 压缩成一个 long long
            long long mask = (long long) i << 32 | j << 1 | pre_down;
            if (memo.contains(mask)) { // 之前算过了
                return memo[mask];
            }
            int res = i == k;
            res += dfs(dfs, i + (1 << j), j + 1, false); // 操作二
            if (!pre_down) {
                res += dfs(dfs, i - 1, j, true); // 操作一
            }
            return memo[mask] = res; // 记忆化
        };
        return dfs(dfs, 1, 0, false);
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
		if i > k+1 { // 无法到达终点 k
			return 0
		}
		p := args{i, j, preDown}
		if v, ok := memo[p]; ok { // 之前算过了
			return v
		}
		res := dfs(i+1<<j, j+1, false) // 操作二
		if !preDown {
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

对于 $j$ 来说，如果 $i+2^j > k + 1$，就不会再递归了。放缩一下，以「$2^j > k + 1$ 不再递归」来估计，说明 $j$ 的上界为 $\mathcal{O}(\log k)$，也就是只有 $\mathcal{O}(\log k)$ 个不同的 $j$。

对于固定的 $j$ 和 $\textit{preDown}$，有多少个不同的 $i$？

由于操作一不能连续使用，所以操作一的次数不超过 $j+1$。所以对于固定的 $j$ 和 $\textit{preDown}$，只有 $\mathcal{O}(j)$ 个不同的 $i$。

累加 $j=0,1,2,\cdots$ 时的 $i$ 的个数，总共有 $\mathcal{O}(1+2+\cdots + \log k) = \mathcal{O}(\log^2 k)$ 个状态。

- 时间复杂度：$\mathcal{O}(\log^2 k)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\log^2 k)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\log^2 k)$。
- 空间复杂度：$\mathcal{O}(\log^2 k)$。保存多少状态，就需要多少空间。

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
class Solution {
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
class Solution {
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
        for (int j = k ? __lg(k - 1) + 1 : 0; (1 << j) - k <= j + 1; j++) {
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

- 时间复杂度：$\mathcal{O}(1)$。预处理的时间和空间不计入。$\texttt{waysToReachStair}$ 中的循环次数至多为 $3$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

为什么返回值是 `int` 而不是 `long long`？答案最大是多少？

见本题视频讲解。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
