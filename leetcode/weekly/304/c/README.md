## 方法一：计算最短路

我们需要知道 $\textit{node}_1$ 到每个点的最短路长度 $\textit{dis}_1[i]$，以及 $\textit{node}_2$ 到每个点的最短路长度 $\textit{dis}_2[i]$。

题目要我们计算的，是 $\max(\textit{dis}_1[i],\textit{dis}_2[i])$ 的最小值对应的节点编号 $i$。若没有这样的节点，返回 $-1$。

求最短路可以用 BFS 做。不过，由于本题输入的是 [内向基环树](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/)（森林），每个连通块至多有一个环，我们可以用一个简单的循环求出 $\textit{dis}_i$。

```py [sol-Python3]
class Solution:
    def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
        n = len(edges)
        def calc_dis(x: int) -> List[int]:
            dis = [n] * n  # 初始化成 n，表示无法到达或者尚未访问的节点
            d = 0
            # 从 x 出发，直到无路可走（x=-1）或者重复访问节点（dis[x]<n）
            while x >= 0 and dis[x] == n:
                dis[x] = d
                d += 1
                x = edges[x]
            return dis

        dis1 = calc_dis(node1)
        dis2 = calc_dis(node2)

        min_dis, ans = n, -1
        for i, (d1, d2) in enumerate(zip(dis1, dis2)):
            d = max(d1, d2)
            if d < min_dis:
                min_dis, ans = d, i
        return ans
```

```java [sol-Java]
class Solution {
    public int closestMeetingNode(int[] edges, int node1, int node2) {
        int[] dis1 = calcDis(edges, node1);
        int[] dis2 = calcDis(edges, node2);

        int n = edges.length;
        int minDis = n;
        int ans = -1;
        for (int i = 0; i < n; i++) {
            int d = Math.max(dis1[i], dis2[i]);
            if (d < minDis) {
                minDis = d;
                ans = i;
            }
        }
        return ans;
    }

    private int[] calcDis(int[] edges, int x) {
        int n = edges.length;
        int[] dis = new int[n];
        Arrays.fill(dis, n); // n 表示无法到达或者尚未访问的节点
        // 从 x 出发，直到无路可走（x=-1）或者重复访问节点（dis[x]<n）
        for (int d = 0; x >= 0 && dis[x] == n; x = edges[x]) {
            dis[x] = d++;
        }
        return dis;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int closestMeetingNode(vector<int>& edges, int node1, int node2) {
        int n = edges.size();
        auto calc_dis = [&](int x) {
            vector<int> dis(n, n); // 初始化成 n，表示无法到达或者尚未访问的节点
            // 从 x 出发，直到无路可走（x=-1）或者重复访问节点（dis[x]<n）
            for (int d = 0; x >= 0 && dis[x] == n; x = edges[x]) {
                dis[x] = d++;
            }
            return dis;
        };

        vector<int> dis1 = calc_dis(node1);
        vector<int> dis2 = calc_dis(node2);

        int min_dis = n, ans = -1;
        for (int i = 0; i < n; i++) {
            int d = max(dis1[i], dis2[i]);
            if (d < min_dis) {
                min_dis = d;
                ans = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func closestMeetingNode(edges []int, node1, node2 int) int {
	n := len(edges)
	calcDis := func(x int) []int {
		dis := make([]int, n)
		for i := range dis {
			dis[i] = n // n 表示无法到达或者尚未访问的节点
		}
		// 从 x 出发，直到无路可走（x=-1）或者重复访问节点（dis[x]<n）
		for d := 0; x >= 0 && dis[x] == n; x = edges[x] {
			dis[x] = d
			d++
		}
		return dis
	}

	dis1 := calcDis(node1)
	dis2 := calcDis(node2)

	minDis, ans := n, -1
	for i, d1 := range dis1 {
		d := max(d1, dis2[i])
		if d < minDis {
			minDis, ans = d, i
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：「我吹过你吹过的晚风」

为方便描述，把 $\textit{node}_1$ 和 $\textit{node}_2$ 分别视作两个人 $x$ 和 $y$。

两个人顺着 $\textit{edges}$ 移动，标记访问过的节点，即更新 $\textit{visX}[x]=\texttt{true}$ 以及 $\textit{visY}[y]=\texttt{true}$。

如果某个时刻，$x$ 到达 $y$ 走过的路，即 $\textit{visY}[x]$ 为 $\texttt{true}$，那么当前 $x$ 所在节点，就是两人都可以到达的节点。

如果某个时刻，$y$ 到达 $x$ 走过的路，即 $\textit{visX}[y]$ 为 $\texttt{true}$，那么当前 $y$ 所在节点，就是两人都可以到达的节点。

如果某个时刻上面两种情况同时发生，返回二者的最小值。

如果两个人都走到死路或者各自走过的路，返回 $-1$。

```py [sol-Python3 集合]
class Solution:
    def closestMeetingNode(self, edges: List[int], x: int, y: int) -> int:
        ans = inf
        vis_x = set()
        vis_y = set()

        while x not in vis_x or y not in vis_y:  # x 或 y 尚未被标记
            vis_x.add(x)
            vis_y.add(y)

            if x in vis_y:  # 我吹过你吹过的晚风
                ans = x
            if y in vis_x:
                ans = min(ans, y)  # 如果有多个答案，返回最小的节点编号
            if ans < inf:
                return ans

            if edges[x] != -1:
                x = edges[x]
            if edges[y] != -1:
                y = edges[y]

        return -1
```

```py [sol-Python3 列表]
class Solution:
    def closestMeetingNode(self, edges: List[int], x: int, y: int) -> int:
        ans = n = len(edges)
        vis_x = [False] * n
        vis_y = [False] * n

        while x >= 0 or y >= 0:
            if x >= 0:
                if vis_y[x]:  # 我吹过你吹过的晚风
                    ans = x
                elif vis_x[x]:
                    x = -1  # 不再重复走
                else:
                    vis_x[x] = True
                    x = edges[x]

            if y >= 0:
                if vis_x[y]:
                    ans = min(ans, y)  # 如果有多个答案，返回最小的节点编号
                elif vis_y[y]:
                    y = -1  # 不再重复走
                else:
                    vis_y[y] = True
                    y = edges[y]

            if ans < n:
                return ans

        return -1
```

```java [sol-Java]
class Solution {
    public int closestMeetingNode(int[] edges, int x, int y) {
        int n = edges.length;
        int ans = n;
        boolean[] visX = new boolean[n];
        boolean[] visY = new boolean[n];

        while (x >= 0 || y >= 0) { // x 或 y 尚未被标记
            if (x >= 0) {
                if (visY[x]) { // 我吹过你吹过的晚风
                    ans = x;
                } else if (visX[x]) {
                    x = -1; // 不再重复走
                } else {
                    visX[x] = true;
                    x = edges[x];
                }
            }

            if (y >= 0) {
                if (visX[y]) {
                    ans = Math.min(ans, y); // 如果有多个答案，返回最小的节点编号
                } else if (visY[y]) {
                    y = -1; // 不再重复走
                } else {
                    visY[y] = true;
                    y = edges[y];
                }
            }

            if (ans < n) {
                return ans;
            }
        }

        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int closestMeetingNode(vector<int>& edges, int x, int y) {
        int n = edges.size();
        int ans = n;
        vector<int8_t> vis_x(n), vis_y(n);

        while (x >= 0 || y >= 0) {
            if (x >= 0) {
                if (vis_y[x]) { // 我吹过你吹过的晚风
                    ans = x;
                } else if (vis_x[x]) {
                    x = -1; // 不再重复走
                } else {
                    vis_x[x] = true;
                    x = edges[x];
                }
            }

            if (y >= 0) {
                if (vis_x[y]) {
                    ans = min(ans, y); // 如果有多个答案，返回最小的节点编号
                } else if (vis_y[y]) {
                    y = -1; // 不再重复走
                } else {
                    vis_y[y] = true;
                    y = edges[y];
                }
            }

            if (ans < n) {
                return ans;
            }
        }

        return -1;
    }
};
```

```go [sol-Go]
func closestMeetingNode(edges []int, x, y int) int {
	n := len(edges)
	ans := n
	visX := make([]bool, n)
	visY := make([]bool, n)

	for x >= 0 || y >= 0 {
		if x >= 0 {
			if visY[x] { // 我吹过你吹过的晚风
				ans = x
			} else if visX[x] {
				x = -1 // 不再重复走
			} else {
				visX[x] = true
				x = edges[x]
			}
		}

		if y >= 0 {
			if visX[y] {
				ans = min(ans, y) // 如果有多个答案，返回最小的节点编号
			} else if visY[y] {
				y = -1 // 不再重复走
			} else {
				visY[y] = true
				y = edges[y]
			}
		}

		if ans < n {
			return ans
		}
	}

	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{edges}$ 的长度。每次循环会有一个新的节点被加到 $\textit{visX}$ 或者 $\textit{visY}$，一共有 $n$ 个节点，所以循环次数至多是 $n$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

1. 如果输入的不止两个节点 $\textit{node}_1$ 和 $\textit{node}_2$，而是一个很长的 $\textit{nodes}$ 列表，要怎么做呢？
2. 如果输入的是 $\textit{queries}$ 询问数组，每个询问包含两个节点 $\textit{node}_1$ 和 $\textit{node}_2$，你需要快速计算 `closestMeetingNode(edges, node1, node2)`，要怎么做呢？

**解答**：见 [视频讲解](https://www.bilibili.com/video/BV1Ba411N78j/?t=22m01s) 第三题。

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
