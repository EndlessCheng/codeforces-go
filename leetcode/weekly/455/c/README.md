## 分析

注意点权 $\textit{cost}[i]$ 只能增大，不能减小。

比如一个子树 $x$ 有 $3$ 个叶子，根到叶子的路径和分别为 $2,3,4$。

那么其中的 $2$ 和 $3$ 都要（通过修改叶子的点权）改成 $4$。

修改次数就是不等于最大值的路径和的个数。

现在，从根经过 $x$ 到叶子的路径和都是 $4$。

如果后面要增大这些路径和，只需增大 $\textit{cost}[x]$，就可以统一增大这些条路径和。这样我们可以视作子树 $x$ 只有一个叶子，且路径和为 $4$。

## 思路

$\text{DFS}(x)$ 返回从根经过 $x$ 到叶子的最大路径和，对于 $x$ 的儿子 $y$ 返回的最大路径和，统计其中的最大值，以及等于最大值的个数 $\textit{cnt}$，那么需要修改的节点个数就是 $x$ 的儿子个数减去 $\textit{cnt}$。

如果 $x$ 是叶子（只有一个邻居），只需返回路径和。

## 细节

为避免误把根节点（只有一个邻居的情况）认为是叶子，可以把 $0$ 与 $-1$ 相连，这样可以保证根至少有两个邻居。注意本题 $n\ge 2$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1GCNRzgEYp/?t=24m5s)，欢迎点赞关注~

## 写法一：有递有归

```py [sol-Python3]
class Solution:
    def minIncrease(self, n: int, edges: List[List[int]], cost: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)
        g[0].append(-1)  # 避免误把根节点当作叶子

        def dfs(x: int, fa: int, path_sum: int) -> int:
            path_sum += cost[x]
            if len(g[x]) == 1:
                return path_sum

            # 在根到叶子的 path_sum 中，有 cnt 个 path_sum 等于 max_s
            max_s = cnt = 0
            for y in g[x]:
                if y == fa:
                    continue
                mx = dfs(y, x, path_sum)
                if mx > max_s:
                    max_s = mx
                    cnt = 1
                elif mx == max_s:
                    cnt += 1

            # 其余小于 max_s 的 path_sum，可以通过增大 cost[y] 的值，改成 max_s
            nonlocal ans
            ans += len(g[x]) - 1 - cnt
            return max_s

        ans = 0
        dfs(0, -1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = 0;

    public int minIncrease(int n, int[][] edges, int[] cost) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        g[0].add(-1); // 避免误把根节点当作叶子

        dfs(0, -1, 0, g, cost);
        return ans;
    }

    private long dfs(int x, int fa, long pathSum, List<Integer>[] g, int[] cost) {
        pathSum += cost[x];
        if (g[x].size() == 1) {
            return pathSum;
        }

        // 在根到叶子的 pathSum 中，有 cnt 个 pathSum 等于 maxS
        long maxS = 0;
        int cnt = 0;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            long mx = dfs(y, x, pathSum, g, cost);
            if (mx > maxS) {
                maxS = mx;
                cnt = 1;
            } else if (mx == maxS) {
                cnt++;
            }
        }

        // 其余小于 maxS 的 pathSum，可以通过增大 cost[y] 的值，改成 maxS
        ans += g[x].size() - 1 - cnt;
        return maxS;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minIncrease(int n, vector<vector<int>>& edges, vector<int>& cost) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }
        g[0].push_back(-1); // 避免误把根节点当作叶子

        int ans = 0;
        auto dfs = [&](this auto&& dfs, int x, int fa, long long path_sum) -> long long {
            path_sum += cost[x];
            if (g[x].size() == 1) {
                return path_sum;
            }

            // 在根到叶子的 path_sum 中，有 cnt 个 path_sum 等于 max_s
            long long max_s = 0;
            int cnt = 0;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                long long mx = dfs(y, x, path_sum);
                if (mx > max_s) {
                    max_s = mx;
                    cnt = 1;
                } else if (mx == max_s) {
                    cnt++;
                }
            }

            // 其余小于 max_s 的 path_sum，可以通过增大 cost[y] 的值，改成 max_s
            ans += g[x].size() - 1 - cnt;
            return max_s;
        };
        dfs(0, -1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func minIncrease(n int, edges [][]int, cost []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	g[0] = append(g[0], -1) // 避免误把根节点当作叶子

	var dfs func(int, int, int) int
	dfs = func(x, fa, pathSum int) (maxS int) {
		pathSum += cost[x]
		if len(g[x]) == 1 {
			return pathSum
		}

		cnt := 0 // 在根到叶子的 pathSum 中，有 cnt 个 pathSum 等于 maxS
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			mx := dfs(y, x, pathSum)
			if mx > maxS {
				maxS = mx
				cnt = 1
			} else if mx == maxS {
				cnt++
			}
		}

		// 其余小于 maxS 的 pathSum，可以通过增大 cost[y] 的值，改成 maxS
		ans += len(g[x]) - 1 - cnt
		return maxS
	}
	dfs(0, -1, 0)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：自底向上

在计算节点 $x$ 的 $\textit{maxS}$ 时，由于路径和都会经过 $x$，所以从根到 $x$ 的这一段和都是一样的。如果去掉这一段和，改成自底向上计算，并不会影响等于 $\textit{maxS}$ 的 $\textit{cnt}$ 值，所以也可以自底向上计算。

```py [sol-Python3]
class Solution:
    def minIncrease(self, n: int, edges: List[List[int]], cost: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)
        g[0].append(-1)

        def dfs(x: int, fa: int) -> int:
            max_s = cnt = 0
            for y in g[x]:
                if y == fa:
                    continue
                mx = dfs(y, x)
                if mx > max_s:
                    max_s = mx
                    cnt = 1
                elif mx == max_s:
                    cnt += 1

            nonlocal ans
            ans += len(g[x]) - 1 - cnt
            return max_s + cost[x]

        ans = 0
        dfs(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = 0;

    public int minIncrease(int n, int[][] edges, int[] cost) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        g[0].add(-1);

        dfs(0, -1, g, cost);
        return ans;
    }

    private long dfs(int x, int fa, List<Integer>[] g, int[] cost) {
        long maxS = 0;
        int cnt = 0;
        for (int y : g[x]) {
            if (y == fa) {
                continue;
            }
            long mx = dfs(y, x, g, cost);
            if (mx > maxS) {
                maxS = mx;
                cnt = 1;
            } else if (mx == maxS) {
                cnt++;
            }
        }
        ans += g[x].size() - 1 - cnt;
        return maxS + cost[x];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minIncrease(int n, vector<vector<int>>& edges, vector<int>& cost) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }
        g[0].push_back(-1);

        int ans = 0;
        auto dfs = [&](this auto&& dfs, int x, int fa) -> long long {
            long long max_s = 0;
            int cnt = 0;
            for (int y : g[x]) {
                if (y == fa) {
                    continue;
                }
                long long mx = dfs(y, x);
                if (mx > max_s) {
                    max_s = mx;
                    cnt = 1;
                } else if (mx == max_s) {
                    cnt++;
                }
            }
            ans += g[x].size() - 1 - cnt;
            return max_s + cost[x];
        };
        dfs(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func minIncrease(n int, edges [][]int, cost []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	g[0] = append(g[0], -1)

	var dfs func(int, int) int
	dfs = func(x, fa int) (maxS int) {
		cnt := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			mx := dfs(y, x)
			if mx > maxS {
				maxS = mx
				cnt = 1
			} else if mx == maxS {
				cnt++
			}
		}
		ans += len(g[x]) - 1 - cnt
		return maxS + cost[x]
	}
	dfs(0, -1)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

见下面树题单的「三、一般树」中的前四小节。

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
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
