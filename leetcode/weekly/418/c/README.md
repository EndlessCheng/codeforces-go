## 拼图怎么玩？

回想一下怎么玩拼图。

![puzzle-1.jpg](https://pic.leetcode.cn/1728206807-HRnpIR-puzzle-1.jpg)

拼图的最外圈是比较容易完成的：

1. 先找角：找有两条直边的拼图块。
2. 再找边：找有一条直边的拼图块。

我们可以先把第一排拼好，然后再考虑和第一排相关的拼图，依此类推。

## 总体思路

1. 构造答案的第一行。
2. 根据第一行的元素，构造下一行，依此类推，直到最后一行。

## 构造答案的第一行

建图，然后把度数相同的点分到同一组。

分类讨论：

1. 如果最小度数是 $1$，类似示例 2，答案只有一列。选其中一个度数为 $1$ 的点，作为第一行。
2. 如果不存在度数为 $4$ 的点，类似示例 1，答案只有两列。选其中一个度数为 $2$ 的点 $x$，以及 $x$ 的一个度数为 $2$ 的邻居 $y$，作为第一行。
3. 否则，答案至少有三列。从其中一个度数为 $2$ 的点（拼图的角）开始，不断寻找度数等于 $3$ 的点（拼图的边），直到找到度数为 $2$ 的点（拼图的另一个角）为止。把遇到的点按顺序作为第一行。

代码实现时，每种度数只需要知道一个点就够了。

## 构造其余行

设第一行的长度为 $k$，那么答案有 $\dfrac{n}{k}$ 行。

用一个布尔数组 $\textit{vis}$ 标记已经填入的数字。

遍历当前行中的元素 $x$，由于 $x$ 的上左右的数字都被标记了，如果 $x$ 的邻居 $y$ 没有被标记过，那么 $y$ 就在 $x$ 的正下方。把 $y$ 加到下一行中。

如此迭代，循环 $\dfrac{n}{k}-1$ 次后构造完毕。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15y1iYUE2h/)，欢迎点赞关注~

## 写法一：分类讨论

```py [sol-Python3]
class Solution:
    def constructGridLayout(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # 每种度数选一个点
        deg_to_node = [-1] * 5
        for x, e in enumerate(g):
            deg_to_node[len(e)] = x

        if deg_to_node[1] != -1:
            # 答案只有一列
            row = [deg_to_node[1]]
        elif deg_to_node[4] == -1:
            # 答案只有两列
            x = deg_to_node[2]
            for y in g[x]:
                if len(g[y]) == 2:
                    row = [x, y]
                    break
        else:
            # 答案至少有三列
            # 寻找度数为 2333...32 的序列作为第一排
            x = deg_to_node[2]
            row = [x]
            pre = x
            x = g[x][0]
            while len(g[x]) == 3:
                row.append(x)
                for y in g[x]:
                    if y != pre and len(g[y]) < 4:
                        pre = x
                        x = y
                        break
            row.append(x)  # x 的度数是 2

        ans = [[] for _ in range(n // len(row))]
        ans[0] = row
        vis = [False] * n
        for x in row:
            vis[x] = True
        for i in range(1, len(ans)):
            for x in ans[i - 1]:
                for y in g[x]:
                    # x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                    if not vis[y]:
                        vis[y] = True
                        ans[i].append(y)
                        break
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] constructGridLayout(int n, int[][] edges) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // 每种度数选一个点
        int[] degToNode = new int[5];
        Arrays.fill(degToNode, -1);
        for (int x = 0; x < n; x++) {
            degToNode[g[x].size()] = x;
        }

        List<Integer> row = new ArrayList<>();
        if (degToNode[1] != -1) {
            // 答案只有一列
            row.add(degToNode[1]);
        } else if (degToNode[4] == -1) {
            // 答案只有两列
            int x = degToNode[2];
            for (int y : g[x]) {
                if (g[y].size() == 2) {
                    row.add(x);
                    row.add(y);
                    break;
                }
            }
        } else {
            // 答案至少有三列
            // 寻找度数为 2333...32 的序列作为第一排
            int x = degToNode[2];
            row.add(x);
            int pre = x;
            x = g[x].get(0);
            while (g[x].size() == 3) {
                row.add(x);
                for (int y : g[x]) {
                    if (y != pre && g[y].size() < 4) {
                        pre = x;
                        x = y;
                        break;
                    }
                }
            }
            row.add(x); // x 的度数是 2
        }

        int k = row.size();
        int[][] ans = new int[n / k][k];
        boolean[] vis = new boolean[n];
        for (int j = 0; j < k; j++) {
            int x = row.get(j);
            ans[0][j] = x;
            vis[x] = true;
        }
        for (int i = 1; i < ans.length; i++) {
            for (int j = 0; j < k; j++) {
                for (int y : g[ans[i - 1][j]]) {
                    // 上左右的邻居都访问过了，没访问过的邻居只会在下面
                    if (!vis[y]) {
                        vis[y] = true;
                        ans[i][j] = y;
                        break;
                    }
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> constructGridLayout(int n, vector<vector<int>>& edges) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // 每种度数选一个点
        int deg_to_node[5]{-1, -1, -1, -1, -1};
        for (int x = 0; x < n; x++) {
            deg_to_node[g[x].size()] = x;
        }

        vector<int> row;
        if (deg_to_node[1] != -1) {
            // 答案只有一列
            row = {deg_to_node[1]};
        } else if (deg_to_node[4] == -1) {
            // 答案只有两列
            int x = deg_to_node[2];
            for (int y : g[x]) {
                if (g[y].size() == 2) {
                    row = {x, y};
                    break;
                }
            }
        } else {
            // 答案至少有三列
            // 寻找度数为 2333...32 的序列作为第一排
            int x = deg_to_node[2];
            row = {x};
            int pre = x;
            x = g[x][0];
            while (g[x].size() == 3) {
                row.push_back(x);
                for (int y : g[x]) {
                    if (y != pre && g[y].size() < 4) {
                        pre = x;
                        x = y;
                        break;
                    }
                }
            }
            row.push_back(x); // x 的度数是 2
        }

        vector<int> vis(n);
        for (int x : row) {
            vis[x] = true;
        }
        vector<vector<int>> ans(n / row.size());
        ans[0] = move(row);
        for (int i = 1; i < ans.size(); i++) {
            for (int x : ans[i - 1]) {
                for (int y : g[x]) {
                    // x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                    if (!vis[y]) {
                        vis[y] = true;
                        ans[i].push_back(y);
                        break;
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func constructGridLayout(n int, edges [][]int) [][]int {
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    // 每种度数选一个点
    degToNode := [5]int{-1, -1, -1, -1, -1}
    for x, to := range g {
        degToNode[len(to)] = x
    }

    var row []int
    if degToNode[1] != -1 {
        // 答案只有一列
        row = []int{degToNode[1]}
    } else if degToNode[4] == -1 {
        // 答案只有两列
        x := degToNode[2]
        for _, y := range g[x] {
            if len(g[y]) == 2 {
                row = []int{x, y}
                break
            }
        }
    } else {
        // 答案至少有三列
        // 寻找度数为 2333...32 的序列作为第一排
        x := degToNode[2]
        row = []int{x}
        pre := x
        x = g[x][0]
        for len(g[x]) == 3 {
            row = append(row, x)
            for _, y := range g[x] {
                if y != pre && len(g[y]) < 4 {
                    pre = x
                    x = y
                    break
                }
            }
        }
        row = append(row, x) // x 的度数是 2
    }

    k := len(row)
    ans := make([][]int, n/k)
    ans[0] = row
    vis := make([]bool, n)
    for _, x := range row {
        vis[x] = true
    }
    for i := 1; i < len(ans); i++ {
        ans[i] = make([]int, k)
        for j, x := range ans[i-1] {
            for _, y := range g[x] {
                // x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                if !vis[y] {
                    vis[y] = true
                    ans[i][j] = y
                    break
                }
            }
        }
    }
    return ans
}
```

## 写法二：合三为一

上文的三种情况，可以统一起来：

1. 首先找到一个度数最小的点 $x$。
2. 找 $x$ 的度数最小的邻居 $y$。
3. 找 $y$ 的度数最小的邻居 $z$，且 $z$ 之前没有访问过。（或者判断 $z\ne x$ 也可以，但是代码要更长一些。）
4. 依此类推，直到找到一个点，度数和起点是一样的。

```py [sol-Python3]
class Solution:
    def constructGridLayout(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # 找一个度数最小的点
        x = 0
        for i, to in enumerate(g):
            if len(to) < len(g[x]):
                x = i

        row = [x]
        vis = [False] * n
        vis[x] = True
        deg_st = len(g[x])  # 起点的度数
        while True:  # 注意题目保证 n >= 2，可以至少循环一次
            nxt = -1
            for y in g[x]:
                if not vis[y] and (nxt < 0 or len(g[y]) < len(g[nxt])):
                    nxt = y
            x = nxt
            row.append(x)
            vis[x] = True
            if len(g[x]) == deg_st:
                break

        ans = [[] for _ in range(n // len(row))]
        ans[0] = row
        for i in range(1, len(ans)):
            for x in ans[i - 1]:
                for y in g[x]:
                    # x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                    if not vis[y]:
                        vis[y] = True
                        ans[i].append(y)
                        break
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] constructGridLayout(int n, int[][] edges) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // 找一个度数最小的点
        int x = 0;
        for (int i = 0; i < g.length; i++) {
            if (g[i].size() < g[x].size()) {
                x = i;
            }
        }

        List<Integer> row = new ArrayList<>();
        row.add(x);
        int degSt = g[x].size(); // 起点的度数
        int pre = -1;
        do { // 注意题目保证 n >= 2，可以至少循环一次
            int nxt = -1;
            for (int y : g[x]) {
                if (y != pre && (nxt < 0 || g[y].size() < g[nxt].size())) {
                    nxt = y;
                }
            }
            pre = x;
            x = nxt;
            row.add(x);
        } while (g[x].size() > degSt);

        int k = row.size();
        int[][] ans = new int[n / k][k];
        boolean[] vis = new boolean[n];
        for (int j = 0; j < k; j++) {
            x = row.get(j);
            ans[0][j] = x;
            vis[x] = true;
        }
        for (int i = 1; i < ans.length; i++) {
            for (int j = 0; j < k; j++) {
                for (int y : g[ans[i - 1][j]]) {
                    // 上左右的邻居都访问过了，没访问过的邻居只会在下面
                    if (!vis[y]) {
                        vis[y] = true;
                        ans[i][j] = y;
                        break;
                    }
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> constructGridLayout(int n, vector<vector<int>>& edges) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // 找一个度数最小的点
        int x = 0;
        for (int i = 0; i < g.size(); i++) {
            if (g[i].size() < g[x].size()) {
                x = i;
            }
        }

        vector<int> row = {x};
        vector<int> vis(n);
        vis[x] = true;
        int deg_st = g[x].size(); // 起点的度数
        do { // 注意题目保证 n >= 2，可以至少循环一次
            int nxt = -1;
            for (int y : g[x]) {
                if (!vis[y] && (nxt < 0 || g[y].size() < g[nxt].size())) {
                    nxt = y;
                }
            }
            x = nxt;
            row.push_back(x);
            vis[x] = true;
        } while (g[x].size() > deg_st);

        vector<vector<int>> ans(n / row.size());
        ans[0] = move(row);
        for (int i = 1; i < ans.size(); i++) {
            for (int x : ans[i - 1]) {
                for (int y : g[x]) {
                    // x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                    if (!vis[y]) {
                        vis[y] = true;
                        ans[i].push_back(y);
                        break;
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func constructGridLayout(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 找一个度数最小的点
	x := 0
	for i, to := range g {
		if len(to) < len(g[x]) {
			x = i
		}
	}

	row := []int{x}
	vis := make([]bool, n)
	vis[x] = true
	degSt := len(g[x]) // 起点的度数
	for { // 注意题目保证 n >= 2，可以至少循环一次
		nxt := -1
		for _, y := range g[x] {
			if !vis[y] && (nxt < 0 || len(g[y]) < len(g[nxt])) {
				nxt = y
			}
		}
		x = nxt
		row = append(row, x)
		vis[x] = true
		if len(g[x]) == degSt {
			break
		}
	}

	k := len(row)
	ans := make([][]int, n/k)
	ans[0] = row
	for i := 1; i < len(ans); i++ {
		ans[i] = make([]int, k)
		for j, x := range ans[i-1] {
			for _, y := range g[x] {
				// 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
				if !vis[y] {
					vis[y] = true
					ans[i][j] = y
					break
				}
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 是 $\textit{edges}$ 的长度。其实也可以算作 $\mathcal{O}(n)$，因为每个点至多 $4$ 个邻居。
- 空间复杂度：$\mathcal{O}(n+m)$。返回值不计入。

想提升思维/构造能力？见下面贪心与思维题单中的「**六、构造题**」。

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
