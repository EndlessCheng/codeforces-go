本题可以上下左右走，是个**有环图**，求最短路长度可以用 Dijkstra 算法。

在标准 Dijkstra 算法的基础上，额外添加参数 $k$，表示当前行动编号的奇偶性。$k=0$ 表示当前行动编号是偶数，$k=1$ 表示当前行动编号是奇数。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 奇数下标 1,3 对应向右或向下
# 偶数下标 0,2 对应向左或向上
DIRS = (0, -1), (0, 1), (-1, 0), (1, 0)

class Solution:
    def minCost(self, m: int, n: int, penalty: list[list[int]]) -> int:
        dis = [[[inf] * 2 for _ in range(n)] for _ in range(m)]

        # 支付 1 的入口代价
        dis[0][0][1] = 1
        h = [(1, 0, 0, 1)]  # (dis, x, y, k)

        while True:
            d, i, j, k = heappop(h)
            if i == m - 1 and j == n - 1:
                return d
            if d > dis[i][j][k]:
                continue
            p = penalty[i][j]

            # 原地不动
            new_dis = d + p
            if new_dis < dis[i][j][k ^ 1]:
                dis[i][j][k ^ 1] = new_dis
                heappush(h, (new_dis, i, j, k ^ 1))  # k^1 切换行动编号的奇偶性

            # 移动一步
            for idx, (dx, dy) in enumerate(DIRS):
                x, y = i + dx, j + dy
                if 0 <= x < m and 0 <= y < n:
                    # 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
                    new_dis = d + (x + 1) * (y + 1) + (idx % 2 ^ k) * p
                    if new_dis < dis[x][y][k ^ 1]:
                        dis[x][y][k ^ 1] = new_dis
                        heappush(h, (new_dis, x, y, k ^ 1))  # k^1 切换行动编号的奇偶性
```

```java [sol-Java]
class Solution {
    // 奇数下标 1,3 对应向右或向下
    // 偶数下标 0,2 对应向左或向上
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

    private record Node(long d, int i, int j, int k) {
    }

    public long minCost(int m, int n, int[][] penalty) {
        long[][][] dis = new long[m][n][2];
        for (long[][] mat : dis) {
            for (long[] row : mat) {
                Arrays.fill(row, Long.MAX_VALUE);
            }
        }

        PriorityQueue<Node> pq = new PriorityQueue<>((a, b) -> Long.compare(a.d, b.d));

        // 支付 1 的入口代价
        dis[0][0][1] = 1;
        pq.offer(new Node(1, 0, 0, 1));

        while (true) {
            Node top = pq.poll();
            long d = top.d;
            int i = top.i;
            int j = top.j;
            int k = top.k;
            if (i == m - 1 && j == n - 1) {
                return d;
            }
            if (d > dis[i][j][k]) {
                continue;
            }
            int p = penalty[i][j];

            // 原地不动
            long newDis = d + p;
            if (newDis < dis[i][j][k ^ 1]) {
                dis[i][j][k ^ 1] = newDis;
                pq.offer(new Node(newDis, i, j, k ^ 1)); // k^1 切换行动编号的奇偶性
            }

            // 移动一步
            for (int idx = 0; idx < 4; idx++) {
                int x = i + DIRS[idx][0];
                int y = j + DIRS[idx][1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    // 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
                    newDis = d + (x + 1) * (y + 1) + (idx % 2 ^ k) * p;
                    if (newDis < dis[x][y][k ^ 1]) {
                        dis[x][y][k ^ 1] = newDis;
                        pq.offer(new Node(newDis, x, y, k ^ 1)); // k^1 切换行动编号的奇偶性
                    }
                }
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    // 奇数下标 1,3 对应向右或向下
    // 偶数下标 0,2 对应向左或向上
    static constexpr int dirs[4][2] = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

public:
    long long minCost(int m, int n, vector<vector<int>>& penalty) {
        vector dis(m, vector<array<long long, 2>>(n, {LLONG_MAX, LLONG_MAX}));
        priority_queue<tuple<long long, int, int, int>, vector<tuple<long long, int, int, int>>, greater<>> pq;

        // 支付 1 的入口代价
        dis[0][0][1] = 1;
        pq.emplace(1, 0, 0, 1);

        while (true) {
            auto [d, i, j, k] = pq.top();
            pq.pop();
            if (i == m - 1 && j == n - 1) {
                return d;
            }
            if (d > dis[i][j][k]) {
                continue;
            }
            int p = penalty[i][j];

            // 原地不动
            auto new_dis = d + p;
            if (new_dis < dis[i][j][k ^ 1]) {
                dis[i][j][k ^ 1] = new_dis;
                pq.emplace(new_dis, i, j, k ^ 1); // k^1 切换行动编号的奇偶性
            }

            // 移动一步
            for (int idx = 0; idx < 4; idx++) {
                int x = i + dirs[idx][0], y = j + dirs[idx][1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    // 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
                    new_dis = d + (x + 1) * (y + 1) + (idx % 2 ^ k) * p;
                    if (new_dis < dis[x][y][k ^ 1]) {
                        dis[x][y][k ^ 1] = new_dis;
                        pq.emplace(new_dis, x, y, k ^ 1); // k^1 切换行动编号的奇偶性
                    }
                }
            }
        }
    }
};
```

```go [sol-Go]
// 奇数下标 1,3 对应向右或向下
// 偶数下标 0,2 对应向左或向上
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func minCost(m, n int, penalty [][]int) int64 {
	dis := make([][][2]int, m)
	for i := range dis {
		dis[i] = make([][2]int, n)
		for j := range dis[i] {
			dis[i][j] = [2]int{math.MaxInt, math.MaxInt}
		}
	}

	// 支付 1 的入口代价
	dis[0][0][1] = 1
	h := &hp{{1, 0, 0, 1}}

	for {
		top := heap.Pop(h).(tuple)
		d, i, j, k := top.dis, top.x, top.y, top.k
		if i == m-1 && j == n-1 {
			return int64(d)
		}
		if d > dis[i][j][k] {
			continue
		}
		p := penalty[i][j]

		// 原地不动
		newDis := d + p
		if newDis < dis[i][j][k^1] {
			dis[i][j][k^1] = newDis
			heap.Push(h, tuple{newDis, i, j, k ^ 1}) // k^1 切换行动编号的奇偶性
		}

		// 移动一步
		for idx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				// 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
				newDis = d + (x+1)*(y+1) + (idx%2^k)*p
				if newDis < dis[x][y][k^1] {
					dis[x][y][k^1] = newDis
					heap.Push(h, tuple{newDis, x, y, k ^ 1}) // k^1 切换行动编号的奇偶性
				}
			}
		}
	}
}

type tuple struct{ dis, x, y, k int }
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log(mn))$。
- 空间复杂度：$\mathcal{O}(mn)$。

## 专题训练

见下面图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
