题目要计算从左上角到右下角的最短路，这可以用 Dijkstra 算法解决。[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)

设从起点走到 $(i,j)$ 的最短路为 $\textit{dis}[i][j]$。

那么从 $(i,j)$ 走到相邻格子 $(x,y)$，到达 $(x,y)$ 的时间为

$$
\max(\textit{dis}[i][j], \textit{moveTime}[x][y]) + \textit{time}
$$

对于周赛第二题来说，$\textit{time}$ 恒为 $1$。

对于本题，由于每走一步 $\textit{time}$ 都会在 $1,2$ 之间变化，联系国际象棋棋盘，$(i+j)$ 的奇偶性就决定了 $\textit{time}$，即

$$
\textit{time} = (i+j)\bmod 2 + 1
$$

由于一定可以从起点走到终点，我们可以在循环中判断，只要出堆的点是终点，就立刻返回 $\textit{dis}[n-1][m-1]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hcS1YCETs/)，欢迎点赞关注~

### 答疑

**问**：为什么代码要判断 `d > dis[i][j]`？可以不写 `continue` 吗？

**答**：对于同一个点 $(i,j)$，例如先入堆一个比较大的 $\textit{dis}[i][j]=10$，后面又把 $\textit{dis}[i][j]$ 更新成 $5$。之后这个 $5$ 会先出堆，然后再把 $10$ 出堆。$10$ 出堆时候是没有必要去更新周围邻居的最短路的，因为 $5$ 出堆之后，就已经把邻居的最短路更新过了，用 $10$ 是无法把邻居的最短路变得更短的，所以直接 `continue`。本题由于只有 $4$ 个邻居，写不写其实无所谓。但如果是一般图，不写这个复杂度就不对了，可能会超时。

```py [sol-Python3]
class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        n, m = len(moveTime), len(moveTime[0])
        dis = [[inf] * m for _ in range(n)]
        dis[0][0] = 0
        h = [(0, 0, 0)]
        while True:
            d, i, j = heappop(h)
            if i == n - 1 and j == m - 1:
                return d
            if d > dis[i][j]:
                continue
            for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):  # 枚举周围四个格子
                if 0 <= x < n and 0 <= y < m:
                    new_dis = max(d, moveTime[x][y]) + (i + j) % 2 + 1
                    if new_dis < dis[x][y]:
                        dis[x][y] = new_dis
                        heappush(h, (new_dis, x, y))
```

```java [sol-Java]
class Solution {
    private final static int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int minTimeToReach(int[][] moveTime) {
        int n = moveTime.length, m = moveTime[0].length;
        int[][] dis = new int[n][m];
        for (int[] row : dis) {
            Arrays.fill(row, Integer.MAX_VALUE);
        }
        dis[0][0] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        pq.add(new int[]{0, 0, 0});
        for (;;) {
            int[] p = pq.poll();
            int d = p[0], i = p[1], j = p[2];
            if (i == n - 1 && j == m - 1) {
                return d;
            }
            if (d > dis[i][j]) {
                continue;
            }
            for (int[] q : DIRS) {
                int x = i + q[0], y = j + q[1];
                if (0 <= x && x < n && 0 <= y && y < m) {
                    int newDis = Math.max(d, moveTime[x][y]) + (i + j) % 2 + 1;
                    if (newDis < dis[x][y]) {
                        dis[x][y] = newDis;
                        pq.add(new int[]{newDis, x, y});
                    }
                }
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int minTimeToReach(vector<vector<int>>& moveTime) {
        int n = moveTime.size(), m = moveTime[0].size();
        vector<vector<int>> dis(n, vector<int>(m, INT_MAX));
        dis[0][0] = 0;
        priority_queue<tuple<int, int, int>, vector<tuple<int, int, int>>, greater<>> pq;
        pq.emplace(0, 0, 0);
        for (;;) {
            auto [d, i, j] = pq.top();
            pq.pop();
            if (i == n - 1 && j == m - 1) {
                return d;
            }
            if (d > dis[i][j]) {
                continue;
            }
            for (auto& q : dirs) {
                int x = i + q[0], y = j + q[1];
                if (0 <= x && x < n && 0 <= y && y < m) {
                    int new_dis = max(d, moveTime[x][y]) + (i + j) % 2 + 1;
                    if (new_dis < dis[x][y]) {
                        dis[x][y] = new_dis;
                        pq.emplace(new_dis, x, y);
                    }
                }
            }
        }
    }
};
```

```go [sol-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minTimeToReach(moveTime [][]int) (ans int) {
	n, m := len(moveTime), len(moveTime[0])
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	h := hp{{}}
	for {
		top := heap.Pop(&h).(tuple)
		i, j := top.x, top.y
		if i == n-1 && j == m-1 {
			return top.dis
		}
		if top.dis > dis[i][j] {
			continue
		}
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < m {
				newD := max(top.dis, moveTime[x][y]) + (i+j)%2 + 1
				if newD < dis[x][y] {
					dis[x][y] = newD
					heap.Push(&h, tuple{newD, x, y})
				}
			}
		}
	}
}

type tuple struct{ dis, x, y int }
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\log (nm))$，其中 $n$ 和 $m$ 分别为 $\textit{moveTime}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(nm)$。

**相似题目**：[2577. 在网格图中访问一个格子的最少时间](https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/)

更多相似题目，见下面网格图题单中的「**BFS**」以及图论题单中的「**单源最短路：Dijkstra**」。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
