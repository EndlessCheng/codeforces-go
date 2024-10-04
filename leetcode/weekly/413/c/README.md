![lc3276.png](https://pic.leetcode.cn/1725186494-jnwWda-lc3276.png)

看示例 1，$\textit{grid}$ 的最大值 $\textit{mx}=4$。

讨论 $4$ **选或不选**：

- 不选 $4$，问题变成在 $[1,3]$ 中选数字，每行至多一个数且没有相同元素时，所选元素之和的最大值。注意题目要求不能选相同的数，选了 $4$ 之后，就不能再选 $4$ 了。
- 选 $4$，**枚举**选哪一行的 $4$（本例中只能选第二行），问题变成在 $[1,3]$ 中选数字，第二行不能选数字，每行至多一个数且没有相同元素时，所选元素之和的最大值。

于是，我们需要在递归的过程中，跟踪两个信息：

- $i$：当前要在 $[1,i]$ 中选数字。
- $j$：已选的行号集合为 $j$。后续所选数字所在的行号不能在 $j$ 中。

因此，定义状态为 $\textit{dfs}(i,j)$，表示在 $[1,i]$ 中选数字，所选数字所处的行号不能在集合 $j$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值。

讨论元素 $i$ **选或不选**：

- 不选 $i$，问题变成在 $[1,i-1]$ 中选数字，所选数字所处的行号不能在集合 $j$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值，即 $\textit{dfs}(i-1,j)$。
- 选 $i$，**枚举**选第 $k$ 行的 $i$（提前预处理 $i$ 所处的行号），问题变成在 $[1,i-1]$ 中选数字（注意 $i$ 只能选一次），所选数字所处的行号不能在集合 $j\cup \{k\}$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值，即 $\textit{dfs}(i-1,j\cup \{k\})$。

两种情况取最大值，即为 $\textit{dfs}(i,j)$。

递归边界：$\textit{dfs}(0,j) = 0$。

递归入口：$\textit{dfs}(\textit{mx},\varnothing)$。一开始没有选行号。

代码实现时，可以只考虑在 $\textit{grid}$ 中的元素。

代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/)（第三题），欢迎点赞关注~

## 一、记忆化搜索（优化前）

```py [sol-Python3]
class Solution:
    def maxScore(self, grid: List[List[int]]) -> int:
        pos = defaultdict(list)
        for i, row in enumerate(grid):
            for x in set(row):  # 去重
                pos[x].append(i)
        all_nums = list(pos)  # 只考虑在 grid 中的元素

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i < 0:
                return 0
            # 不选 x
            res = dfs(i - 1, j)
            # 枚举选第 k 行的 x
            for k in pos[all_nums[i]]:
                if (j >> k & 1) == 0:
                    res = max(res, dfs(i - 1, j | 1 << k) + all_nums[i])
            return res
        return dfs(len(all_nums) - 1, 0)
```

```java [sol-Java]
class Solution {
    public int maxScore(List<List<Integer>> grid) {
        Map<Integer, Integer> pos = new HashMap<>();
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid.get(i)) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos.merge(x, 1 << i, (a, b) -> a | b);
            }
        }

        // 只考虑在 grid 中的元素
        List<Integer> allNums = new ArrayList<>(pos.keySet());
        int n = allNums.size();
        int[][] memo = new int[n][1 << m];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, 0, pos, allNums, memo);
    }

    private int dfs(int i, int j, Map<Integer, Integer> pos, List<Integer> allNums, int[][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        // 不选 x
        int res = dfs(i - 1, j, pos, allNums, memo);
        // 枚举选第 k 行的 x
        int x = allNums.get(i);
        for (int t = pos.get(x), lb; t > 0; t ^= lb) {
            lb = t & -t; // lb = 1<<k，其中 k 是行号
            if ((j & lb) == 0) { // 没选过第 k 行的数
                res = Math.max(res, dfs(i - 1, j | lb, pos, allNums, memo) + x);
            }
        }
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(vector<vector<int>>& grid) {
        unordered_map<int, int> pos;
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid[i]) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos[x] |= 1 << i;
            }
        }

        // 只考虑在 grid 中的元素
        vector<int> all_nums;
        for (auto& [x, _] : pos) {
            all_nums.push_back(x);
        }

        int n = all_nums.size();
        vector<vector<int>> memo(n, vector<int>(1 << m, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            // 不选 x
            res = dfs(dfs, i - 1, j);
            // 枚举选第 k 行的 x
            int x = all_nums[i];
            for (int t = pos[x], lb; t; t ^= lb) {
                lb = t & -t; // lb = 1<<k，其中 k 是行号
                if ((j & lb) == 0) { // 没选过第 k 行的数
                    res = max(res, dfs(dfs, i - 1, j | lb) + x);
                }
            }
            return res;
        };
        return dfs(dfs, n - 1, 0);
    }
};
```

```go [sol-Go]
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	// 只考虑在 grid 中的元素
	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		// 不选 x
		res := dfs(i-1, j)
		// 枚举选第 k 行的 x
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, 0)
}
```

## 二、记忆化搜索（优化）

$\textit{grid}$ 的最大值是肯定要选的，不会出现不选的情况。

进一步地，我们从大到小递归所有元素，如果当前元素可以选（有之前没选过的行），那么就**不用考虑不选的情况**了。

```py [sol-Python3]
class Solution:
    def maxScore(self, grid: List[List[int]]) -> int:
        pos = defaultdict(list)
        for i, row in enumerate(grid):
            for x in set(row):  # 去重
                pos[x].append(i)
        all_nums = sorted(pos)  # 下面从大到小递归

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i < 0:
                return 0
            res = 0  # 如果循环结束后 res > 0，就不再递归不选的情况
            for k in pos[all_nums[i]]:
                if (j >> k & 1) == 0:
                    res = max(res, dfs(i - 1, j | 1 << k) + all_nums[i])
            return res if res else dfs(i - 1, j)
        return dfs(len(all_nums) - 1, 0)
```

```java [sol-Java]
class Solution {
    public int maxScore(List<List<Integer>> grid) {
        Map<Integer, Integer> pos = new HashMap<>();
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid.get(i)) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos.merge(x, 1 << i, (a, b) -> a | b);
            }
        }

        List<Integer> allNums = new ArrayList<>(pos.keySet());
        Collections.sort(allNums); // 下面从大到小递归
        int n = allNums.size();
        int[][] memo = new int[n][1 << m];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, 0, pos, allNums, memo);
    }

    private int dfs(int i, int j, Map<Integer, Integer> pos, List<Integer> allNums, int[][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        // 枚举选第 k 行的 x
        // 如果循环结束后 res > 0，就不再递归不选的情况
        int res = 0;
        int x = allNums.get(i);
        for (int t = pos.get(x), lb; t > 0; t ^= lb) {
            lb = t & -t; // lb = 1<<k，其中 k 是行号
            if ((j & lb) == 0) { // 没选过第 k 行的数
                res = Math.max(res, dfs(i - 1, j | lb, pos, allNums, memo) + x);
            }
        }
        if (res == 0) {
            // 不选 x
            res = dfs(i - 1, j, pos, allNums, memo);
        }
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(vector<vector<int>>& grid) {
        unordered_map<int, int> pos;
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid[i]) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos[x] |= 1 << i;
            }
        }

        vector<int> all_nums;
        for (auto& [x, _] : pos) {
            all_nums.push_back(x);
        }
        ranges::sort(all_nums); // 下面从大到小递归

        int n = all_nums.size();
        vector<vector<int>> memo(n, vector<int>(1 << m, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            // 枚举选第 k 行的 x
            int x = all_nums[i];
            for (int t = pos[x], lb; t; t ^= lb) {
                lb = t & -t; // lb = 1<<k，其中 k 是行号
                if ((j & lb) == 0) { // 没选过第 k 行的数
                    res = max(res, dfs(dfs, i - 1, j | lb) + x);
                }
            }
            if (res == -1) {
                // 不选 x
                res = dfs(dfs, i - 1, j);
            }
            return res;
        };
        return dfs(dfs, n - 1, 0);
    }
};
```

```go [sol-Go]
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}
	slices.Sort(allNums) // 下面从大到小递归

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		// 枚举选第 k 行的 x
		// 如果循环结束后 res > 0，就不再递归不选的情况
		res := 0
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		if res == 0 {
			// 不选 x
			res = dfs(i-1, j)
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, 0)
}
```

## 三、递推

```py [sol-Python3]
class Solution:
    def maxScore(self, grid: List[List[int]]) -> int:
        pos = defaultdict(list)
        for i, row in enumerate(grid):
            for x in set(row):  # 去重
                pos[x].append(i)
        all_nums = list(pos)

        u = 1 << len(grid)
        f = [[0] * u for _ in range(len(all_nums) + 1)]
        for i, x in enumerate(all_nums):
            for j in range(u):
                f[i + 1][j] = f[i][j]  # 不选 x
                for k in pos[x]:
                    if (j >> k & 1) == 0:  # 没选过第 k 行的数
                        f[i + 1][j] = max(f[i + 1][j], f[i][j | 1 << k] + x)  # 选第 k 行的 x
        return f[-1][0]
```

```java [sol-Java]
class Solution {
    public int maxScore(List<List<Integer>> grid) {
        Map<Integer, Integer> pos = new HashMap<>();
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid.get(i)) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos.merge(x, 1 << i, (a, b) -> a | b);
            }
        }

        List<Integer> allNums = new ArrayList<>(pos.keySet());
        int u = 1 << m;
        int[][] f = new int[allNums.size() + 1][u];
        for (int i = 0; i < allNums.size(); i++) {
            int x = allNums.get(i);
            int posMask = pos.get(x);
            for (int j = 0; j < u; j++) {
                f[i + 1][j] = f[i][j]; // 不选 x
                for (int t = posMask, lb; t > 0; t ^= lb) {
                    lb = t & -t; // lb = 1<<k，其中 k 是行号
                    if ((j & lb) == 0) { // 没选过第 k 行的数
                        f[i + 1][j] = Math.max(f[i + 1][j], f[i][j | lb] + x); // 选第 k 行的 x
                    }
                }
            }
        }
        return f[allNums.size()][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(vector<vector<int>>& grid) {
        unordered_map<int, int> pos;
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid[i]) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos[x] |= 1 << i;
            }
        }

        vector<int> all_nums;
        for (auto& [x, _] : pos) {
            all_nums.push_back(x);
        }

        int u = 1 << m;
        vector<vector<int>> f(all_nums.size() + 1, vector<int>(u));
        for (int i = 0; i < all_nums.size(); i++) {
            int x = all_nums[i];
            for (int j = 0; j < u; j++) {
                f[i + 1][j] = f[i][j]; // 不选 x
                for (int t = pos[x], lb; t; t ^= lb) {
                    lb = t & -t; // lb = 1<<k，其中 k 是行号
                    if ((j & lb) == 0) { // 没选过第 k 行的数
                        f[i + 1][j] = max(f[i + 1][j], f[i][j | lb] + x); // 选第 k 行的 x
                    }
                }
            }
        }
        return f.back()[0];
    }
};
```

```go [sol-Go]
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	f := make([][]int, len(allNums)+1)
	for i := range f {
		f[i] = make([]int, 1<<len(grid))
	}
	for i, x := range allNums {
		for j, v := range f[i] {
			f[i+1][j] = v // 不选 x
			for t, lb := pos[x], 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[i+1][j] = max(f[i+1][j], f[i][j|lb]+x) // 选第 k 行的 x
				}
			}
		}
	}
	return f[len(allNums)][0]
}
```

## 四、空间压缩

去掉 $f$ 的第一个维度。

```py [sol-Python3]
class Solution:
    def maxScore(self, grid: List[List[int]]) -> int:
        pos = defaultdict(list)
        for i, row in enumerate(grid):
            for x in set(row):  # 去重
                pos[x].append(i)

        u = 1 << len(grid)
        f = [0] * u
        for x, ps in pos.items():
            for j in range(u):
                for k in ps:
                    if (j >> k & 1) == 0:
                        f[j] = max(f[j], f[j | 1 << k] + x)
        return f[0]
```

```java [sol-Java]
class Solution {
    public int maxScore(List<List<Integer>> grid) {
        Map<Integer, Integer> pos = new HashMap<>();
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid.get(i)) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos.merge(x, 1 << i, (a, b) -> a | b);
            }
        }

        int u = 1 << m;
        int[] f = new int[u];
        for (Map.Entry<Integer, Integer> e : pos.entrySet()) {
            int x = e.getKey();
            int posMask = e.getValue();
            for (int j = 0; j < u; j++) {
                for (int t = posMask, lb; t > 0; t ^= lb) {
                    lb = t & -t; // lb = 1<<k，其中 k 是行号
                    if ((j & lb) == 0) { // 没选过第 k 行的数
                        f[j] = Math.max(f[j], f[j | lb] + x);
                    }
                }
            }
        }
        return f[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(vector<vector<int>>& grid) {
        unordered_map<int, int> pos;
        int m = grid.size();
        for (int i = 0; i < m; i++) {
            for (int x : grid[i]) {
                // 保存所有包含 x 的行号（压缩成二进制数）
                pos[x] |= 1 << i;
            }
        }

        int u = 1 << m;
        vector<int> f(u);
        for (auto& [x, pos_mask] : pos) {
            for (int j = 0; j < u; j++) {
                for (int t = pos_mask, lb; t; t ^= lb) {
                    lb = t & -t; // lb = 1<<k，其中 k 是行号
                    if ((j & lb) == 0) { // 没选过第 k 行的数
                        f[j] = max(f[j], f[j | lb] + x);
                    }
                }
            }
        }
        return f[0];
    }
};
```

```go [sol-Go]
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	f := make([]int, 1<<len(grid))
	for x, posMask := range pos {
		for j := range f {
			for t, lb := posMask, 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[j] = max(f[j], f[j|lb]+x)
				}
			}
		}
	}
	return f[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn2^m)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn + 2^m)$ 或 $\mathcal{O}(k + 2^m)$，其中 $k$ 是 $\textit{grid}$ 中的不同元素的个数。

## 五、费用流（选读）

和 [3257. 放三个车的价值之和最大 II](https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/) 一样，本题也可以抽象成图论模型，用**费用流**解决。

- 创建一个**二分图**，左部为元素值，右部为行号。
- 把第 $i$ 个元素看作节点 $i$，第 $j$ 行看作节点 $k+j$。其中 $k$ 是 $\textit{grid}$ 中的不同元素个数。
- 在第 $i$ 个元素到其所在的第 $j$ 行之间连边，容量为 $1$，费用为 $0$。
- 从超级源点 $S=k+m$ 向所有元素节点 $0,1,2,\cdots,k-1$ 连边，容量为 $1$，费用为节点值取反。取反是为了方便套最小费用流模板。
- 从所有行节点 $k,k+1,k+2,\cdots,k+m-1$ 向超级汇点 $T=k+m+1$ 连边，容量为 $1$，费用为 $0$。

这样建图可以保证同样的元素至多选一个，且每一行至多选一个元素。

计算从 $S$ 到 $T$ 的最小费用流，取相反数，即为答案。

```go
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	// 建图
	k := len(pos)
	m := len(grid)
	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, k+m+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	S := k + m
	T := k + m + 1
	i := 0
	for x, posMask := range pos {
		for t := uint(posMask); t > 0; t &= t - 1 {
			j := bits.TrailingZeros(t)
			addEdge(i, k+j, 1, 0)
		}
		addEdge(S, i, 1, -x)
		i++
	}
	for j := range grid {
		addEdge(k+j, T, 1, 0)
	}

	// 下面是费用流模板
	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[T] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[T] * minF
	}
	return -minCost
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m^2n)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于二分图的特殊性，算法跑至多 $m$ 次 $\mathcal{O}(mn)$ 的 SPFA 就结束了。
- 空间复杂度：$\mathcal{O}(mn)$。

## 相似题目

- [1434. 每个人戴不同帽子的方案数](https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/)

更多相似题目，见下面动态规划题单中的「**状压 DP**」和图论题单中的「**二分图**」和「**网络流**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
