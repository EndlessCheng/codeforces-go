[本题视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/)

## 前置知识：动态规划入门

请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://b23.tv/72onpYq)，制作不易，欢迎点赞！

## 写法一：自顶向下（记忆化搜索）

把 `floor(coins[i] / 2)` 看成右移操作。

一个数最多右移多少次，就变成 $0$ 了？在本题的数据范围下，这至多是 $14$ 次。

同时，右移操作是可以叠加的，我们可以记录子树中的节点值右移了多少次。

所以可以定义 $\textit{dfs}(i,j)$ 表示子树 $i$ 在已经右移 $j$ 次的前提下，最多可以得到多少积分。

用「选或不选」来思考，即是否右移：

- 不右移：答案为 $(\textit{coins}[i]>>j)-k$ 加上每个子树 $\textit{ch}$ 的 $\textit{dfs}(ch,j)$。
- 右移：答案为 $\textit{coins}[i]>>(j+1)$ 加上每个子树 $\textit{ch}$ 的 $\textit{dfs}(ch,j+1)$。

这两种情况取最大值。

```py [sol-Python3]
class Solution:
    def maximumPoints(self, edges: List[List[int]], coins: List[int], k: int) -> int:
        g = [[] for _ in coins]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        @cache
        def dfs(i: int, j: int, fa: int) -> int:
            res1 = (coins[i] >> j) - k
            res2 = coins[i] >> (j + 1)
            for ch in g[i]:
                if ch != fa:
                    res1 += dfs(ch, j, i)  # 不右移
                    if j < 13:  # j+1 >= 14 相当于 res2 += 0，无需递归
                        res2 += dfs(ch, j + 1, i)  # 右移
            return max(res1, res2)
        return dfs(0, 0, -1)
```

```java [sol-Java]
class Solution {
    public int maximumPoints(int[][] edges, int[] coins, int k) {
        int n = coins.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        int[][] memo = new int[n][14];
        for (int[] m : memo) {
            Arrays.fill(m, -1); // -1 表示没有计算过
        }
        return dfs(0, 0, -1, memo, g, coins, k);
    }

    private int dfs(int i, int j, int fa, int[][] memo, List<Integer>[] g, int[] coins, int k) {
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res1 = (coins[i] >> j) - k;
        int res2 = coins[i] >> (j + 1);
        for (int ch : g[i]) {
            if (ch == fa) continue;
            res1 += dfs(ch, j, i, memo, g, coins, k); // 不右移
            if (j < 13) { // j+1 >= 14 相当于 res2 += 0，无需递归
                res2 += dfs(ch, j + 1, i, memo, g, coins, k); // 右移
            }
        }
        return memo[i][j] = Math.max(res1, res2); // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    int dfs(int i, int j, int fa, vector<array<int, 14>> &memo, vector<vector<int>> &g, vector<int> &coins, int k) {
        int &res = memo[i][j]; // 注意这里是引用
        if (res != -1) { // 之前计算过
            return res;
        }
        int res1 = (coins[i] >> j) - k;
        int res2 = coins[i] >> (j + 1);
        for (int ch: g[i]) {
            if (ch == fa) continue;
            res1 += dfs(ch, j, i, memo, g, coins, k); // 不右移
            if (j < 13) { // j+1 >= 14 相当于 res2 += 0，无需递归
                res2 += dfs(ch, j + 1, i, memo, g, coins, k); // 右移
            }
        }
        return res = max(res1, res2); // 记忆化
    };

public:
    int maximumPoints(vector<vector<int>> &edges, vector<int> &coins, int k) {
        int n = coins.size();
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        array<int, 14> init_val;
        ranges::fill(init_val, -1); // -1 表示没有计算过
        vector<array<int, 14>> memo(n, init_val);
        return dfs(0, 0, -1, memo, g, coins, k);
    }
};
```

```go [sol-Go]
func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][14]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, fa int) int {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res1 := coins[i]>>j - k
		res2 := coins[i] >> (j + 1)
		for _, ch := range g[i] {
			if ch != fa {
				res1 += dfs(ch, j, i) // 不右移
				if j < 13 { // j+1 >= 14 相当于 res2 += 0 无需递归
					res2 += dfs(ch, j+1, i) // 右移
				}
			}
		}
		*p = max(res1, res2)
		return *p
	}
	return dfs(0, 0, -1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{coins}$ 的长度，$U=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 写法二：自底向上

类似把记忆化搜索翻译成递推的过程，我们也可以从下往上算：

```py [sol-Python3]
class Solution:
    def maximumPoints(self, edges: List[List[int]], coins: List[int], k: int) -> int:
        g = [[] for _ in coins]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int) -> List[int]:
            res1 = [0] * 14
            res2 = [0] * 14
            for y in g[x]:
                if y == fa: continue
                r = dfs(y, x)
                for j, v in enumerate(r):
                    res1[j] += v
                    if j < 13:
                        res2[j] += r[j + 1]
            for j, (r1, r2) in enumerate(zip(res1, res2)):
                res1[j] = max(r1 + (coins[x] >> j) - k, r2 + (coins[x] >> (j + 1)))
            return res1
        return dfs(0, -1)[0]
```

```java [sol-Java]
class Solution {
    public int maximumPoints(int[][] edges, int[] coins, int k) {
        List<Integer>[] g = new ArrayList[coins.length];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return dfs(0, -1, g, coins, k)[0];
    }

    private int[] dfs(int x, int fa, List<Integer>[] g, int[] coins, int k) {
        int[] res1 = new int[14];
        int[] res2 = new int[14];
        for (int y : g[x]) {
            if (y == fa) continue;
            int[] r = dfs(y, x, g, coins, k);
            for (int j = 0; j < r.length; j++) {
                res1[j] += r[j];
                if (j < 13) {
                    res2[j] += r[j + 1];
                }
            }
        }
        for (int j = 0; j < res1.length; j++) {
            res1[j] = Math.max(res1[j] + (coins[x] >> j) - k, res2[j] + (coins[x] >> (j + 1)));
        }
        return res1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumPoints(vector<vector<int>> &edges, vector<int> &coins, int k) {
        vector<vector<int>> g(coins.size());
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        function<array<int, 14>(int, int)> dfs = [&](int x, int fa) -> array<int, 14> {
            array<int, 14> res1{}, res2{};
            for (int y : g[x]) {
                if (y == fa) continue;
                auto r = dfs(y, x);
                for (int j = 0; j < 14; j++) {
                    res1[j] += r[j];
                    if (j < 13) {
                        res2[j] += r[j + 1];
                    }
                }
            }
            for (int j = 0; j < 14; j++) {
                res1[j] = max(res1[j] + (coins[x] >> j) - k, res2[j] + (coins[x] >> (j + 1)));
            }
            return res1;
        };
        return dfs(0, -1)[0];
    }
};
```

```go [sol-Go]
func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) [14]int
	dfs = func(x, fa int) (res1 [14]int) {
		res2 := [14]int{}
		for _, y := range g[x] {
			if y != fa {
				r := dfs(y, x)
				for j, v := range r {
					res1[j] += v
					if j < 13 {
						res2[j] += r[j+1]
					}
				}
			}
		}
		for j := 0; j < 14; j++ {
			res1[j] = max(res1[j]+coins[x]>>j-k, res2[j]+coins[x]>>(j+1))
		}
		return
	}
	return dfs(0, -1)[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{coins}$ 的长度，$U=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(n\log U)$。
