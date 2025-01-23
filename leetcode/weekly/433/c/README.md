## 一、寻找子问题

原问题是涂下标在 $[0,n-1]$ 内的房子的最低成本。

由于有等距异色的要求，考虑**左右同时开工**，涂 $i$ 和 $n-1-i$ 这两个房子。

枚举房子 $i$ 和 $n-1-i$ 的颜色 $j$ 和 $k$，问题变成涂剩余房子的最低成本。这样就找到了子问题。

## 二、状态定义与状态转移方程

为方便把记忆化搜索翻译成递推，从里向外涂色。

定义 $\textit{dfs}(i,\textit{preJ},\textit{preK})$ 表示涂前后各 $i+1$ 个房子的最低成本，其中下标 $i+1$ 房子的颜色为 $\textit{preJ}$，下标 $n-2-i$ 房子的颜色为 $\textit{preK}$。

枚举房子 $i$ 和 $n-1-i$ 的颜色 $j$ 和 $k$，问题变成涂前后各 $i$ 个房子的最低成本，即 $\textit{dfs}(i-1,j,k)$。

累加成本，取最小值，得

$$
\textit{dfs}(i,\textit{preJ},\textit{preK}) = \min\limits_{j,k} \textit{dfs}(i-1,j,k) + \textit{cost}[i][j] + \textit{cost}[n-1-i][k]
$$

其中 $j\ne \textit{preJ}$ 且 $k\ne \textit{preK}$ 且 $j\ne k$。

**递归边界**：$\textit{dfs}(-1,\textit{preJ},\textit{preK})=0$。

**递归入口**：$\textit{dfs}(n/2-1,3,3)$，这是原问题，也是答案。其中 $3$ 保证一开始的涂色不受约束。注意题目保证 $n$ 是偶数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17RwBeqErJ/?t=30m32s)，欢迎点赞关注~

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,\textit{preJ},\textit{preK})$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def minCost(self, n: int, cost: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, pre_j: int, pre_k: int) -> int:
            if i < 0:
                return 0
            res = inf
            for j, c1 in enumerate(cost[i]):
                if j == pre_j:
                    continue
                for k, c2 in enumerate(cost[-1 - i]):
                    if k != pre_k and k != j:
                        res = min(res, dfs(i - 1, j, k) + c1 + c2)
            return res
        return dfs(n // 2 - 1, 3, 3)
```

```java [sol-Java]
class Solution {
    public long minCost(int n, int[][] cost) {
        long[][][] memo = new long[n / 2][4][4];
        for (long[][] mat : memo) {
            for (long[] arr : mat) {
                Arrays.fill(arr, -1); // -1 表示没有计算过
            }
        }
        return dfs(n / 2 - 1, 3, 3, cost, memo);
    }

    private long dfs(int i, int preJ, int preK, int[][] cost, long[][][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][preJ][preK] != -1) { // 之前计算过
            return memo[i][preJ][preK];
        }
        long res = Long.MAX_VALUE;
        for (int j = 0; j < 3; j++) {
            if (j == preJ) {
                continue;
            }
            for (int k = 0; k < 3; k++) {
                if (k != preK && k != j) {
                    res = Math.min(res, dfs(i - 1, j, k, cost, memo) + cost[i][j] + cost[cost.length - 1 - i][k]);
                }
            }
        }
        return memo[i][preJ][preK] = res; // 记忆化
    }
}
```
    
```cpp [sol-C++]
class Solution {
public:
    long long minCost(int n, vector<vector<int>>& cost) {
        long long memo[n / 2][4][4];
        memset(memo, -1, sizeof(memo)); // -1 表示没有计算过
        auto dfs = [&](this auto&& dfs, int i, int pre_j, int pre_k) -> long long {
            if (i < 0) {
                return 0;
            }
            long long& res = memo[i][pre_j][pre_k]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = LLONG_MAX;
            for (int j = 0; j < 3; j++) {
                if (j == pre_j) {
                    continue;
                }
                for (int k = 0; k < 3; k++) {
                    if (k != pre_k && k != j) {
                        res = min(res, dfs(i - 1, j, k) + cost[i][j] + cost[n - 1 - i][k]);
                    }
                }
            }
            return res;
        };
        return dfs(n / 2 - 1, 3, 3);
    }
};
```

```go [sol-Go]
func minCost(n int, cost [][]int) int64 {
	memo := make([][4][4]int, n/2)
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, preJ, preK int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i][preJ][preK]
		if *p != -1 { // 之前计算过
			return *p
		}
		res = math.MaxInt
		for j, c1 := range cost[i] {
			if j == preJ {
				continue
			}
			for k, c2 := range cost[n-1-i] {
				if k != preK && k != j {
					res = min(res, dfs(i-1, j, k)+c1+c2)
				}
			}
		}
		*p = res // 记忆化
		return
	}
	return int64(dfs(n/2-1, 3, 3))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^4)$，其中 $k=3$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nk^2)$，单个状态的计算时间为 $\mathcal{O}(k^2)$，所以总的时间复杂度为 $\mathcal{O}(nk^4)$。
- 空间复杂度：$\mathcal{O}(nk^2)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

$f[i+1][\textit{preJ}][\textit{preK}]$ 的定义和 $\textit{dfs}(i,\textit{preJ},\textit{preK})$ 的定义是一样的，都表示涂前后各 $i+1$ 个房子的最低成本，其中下标 $i+1$ 房子的颜色为 $\textit{preJ}$，下标 $n-2-i$ 房子的颜色为 $\textit{preK}$。这里写 $f[i+1]$ 是为了把 $\textit{dfs}(-1,\textit{preJ},\textit{preK})$ 这个状态也翻译过来，这样我们可以把 $f[0]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][\textit{preJ}][\textit{preK}] = \min\limits_{j,k} f[i][j][k] + \textit{cost}[i][j] + \textit{cost}[n-1-i][k]
$$

其中 $j\ne \textit{preJ}$ 且 $k\ne \textit{preK}$ 且 $j\ne k$。

初始值 $f[0][\textit{preJ}][\textit{preK}]=0$。

答案可以枚举最里面一对房子的颜色，取最小值，即 $f[n/2][j][k]$ 的最小值。

```py [sol-Python3]
class Solution:
    def minCost(self, n: int, cost: List[List[int]]) -> int:
        f = [[[0] * 3 for _ in range(3)] for _ in range(n // 2 + 1)]
        for i, row in enumerate(cost[:n // 2]):
            row2 = cost[-1 - i]
            for pre_j in range(3):
                for pre_k in range(3):
                    res = inf
                    for j, c1 in enumerate(row):
                        if j == pre_j:
                            continue
                        for k, c2 in enumerate(row2):
                            if k != pre_k and k != j:
                                res = min(res, f[i][j][k] + c1 + c2)
                    f[i + 1][pre_j][pre_k] = res
        # 枚举所有初始颜色，取最小值
        return min(map(min, f[-1]))
```

```java [sol-Java]
class Solution {
    public long minCost(int n, int[][] cost) {
        long[][][] f = new long[n / 2 + 1][3][3];
        for (int i = 0; i < n / 2; i++) {
            int[] row = cost[i];
            int[] row2 = cost[n - 1 - i];
            for (int preJ = 0; preJ < 3; preJ++) {
                for (int preK = 0; preK < 3; preK++) {
                    long res = Long.MAX_VALUE;
                    for (int j = 0; j < 3; j++) {
                        if (j == preJ) {
                            continue;
                        }
                        for (int k = 0; k < 3; k++) {
                            if (k != preK && k != j) {
                                res = Math.min(res, f[i][j][k] + row[j] + row2[k]);
                            }
                        }
                    }
                    f[i + 1][preJ][preK] = res;
                }
            }
        }
        // 枚举所有初始颜色，取最小值
        return Arrays.stream(f[n / 2]).flatMapToLong(Arrays::stream).min().getAsLong();
    }
}
```

```cpp [sol-C++]
#include <ranges>
class Solution {
public:
    long long minCost(int n, vector<vector<int>>& cost) {
        vector<array<array<long long, 3>, 3>> f(n / 2 + 1);
        for (int i = 0; i < n / 2; i++) {
            auto& row = cost[i];
            auto& row2 = cost[n - 1 - i];
            for (int pre_j = 0; pre_j < 3; pre_j++) {
                for (int pre_k = 0; pre_k < 3; pre_k++) {
                    long long res = LLONG_MAX;
                    for (int j = 0; j < 3; j++) {
                        if (j == pre_j) {
                            continue;
                        }
                        for (int k = 0; k < 3; k++) {
                            if (k != pre_k && k != j) {
                                res = min(res, f[i][j][k] + row[j] + row2[k]);
                            }
                        }
                    }
                    f[i + 1][pre_j][pre_k] = res;
                }
            }
        }
        // 枚举所有初始颜色，取最小值
        return ranges::min(ranges::transform_view(f.back(), ranges::min));
    }
};
```

```go [sol-Go]
func minCost(n int, cost [][]int) int64 {
	f := make([][3][3]int, n/2+1)
	for i, row := range cost[:n/2] {
		row2 := cost[n-1-i]
		for preJ := range 3 {
			for preK := range 3 {
				res := math.MaxInt
				for j, c1 := range row {
					if j == preJ {
						continue
					}
					for k, c2 := range row2 {
						if k != preK && k != j {
							res = min(res, f[i][j][k]+c1+c2)
						}
					}
				}
				f[i+1][preJ][preK] = res
			}
		}
	}

	// 枚举所有初始颜色，取最小值
	res := math.MaxInt
	for _, row := range f[n/2] {
		res = min(res, slices.Min(row[:]))
	}
	return int64(res)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^4)$，其中 $k=3$。
- 空间复杂度：$\mathcal{O}(nk^2)$。

注：利用滚动数组，可以把空间复杂度优化至 $\mathcal{O}(k^2)$。

更多相似题目，见下面动态规划题单中的「**§7.5 多维 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
