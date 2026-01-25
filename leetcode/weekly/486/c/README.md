跑三次 DFS（或者 BFS），计算从 $x,y,z$ 出发，到每个点的距离。

然后枚举点 $0,1,2,\ldots,n-1$，判断 $\textit{dx},\textit{dy},\textit{dz}$ 是否构成勾股数元组，也就是把 $\textit{dx},\textit{dy},\textit{dz}$ 从小到大排序，得到 $a\le b\le c$，判断 $a^2+b^2=c^2$ 是否成立。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def specialNodes(self, n: int, edges: List[List[int]], x: int, y: int, z: int) -> int:
        g = [[] for _ in range(n)]
        for v, w in edges:
            g[v].append(w)
            g[w].append(v)

        def calc_dis(start: int) -> List[int]:
            dis = [0] * n
            def dfs(v: int, fa: int) -> None:
                for w in g[v]:
                    if w != fa:
                        dis[w] = dis[v] + 1
                        dfs(w, v)
            dfs(start, -1)
            return dis

        dx = calc_dis(x)
        dy = calc_dis(y)
        dz = calc_dis(z)

        ans = 0
        for t in zip(dx, dy, dz):
            a, b, c = sorted(t)
            if a * a + b * b == c * c:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int specialNodes(int n, int[][] edges, int x, int y, int z) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int v = e[0];
            int w = e[1];
            g[v].add(w);
            g[w].add(v);
        }

        int[] dx = calcDis(x, g);
        int[] dy = calcDis(y, g);
        int[] dz = calcDis(z, g);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            int[] a = new int[]{dx[i], dy[i], dz[i]};
            Arrays.sort(a);
            if ((long) a[0] * a[0] + (long) a[1] * a[1] == (long) a[2] * a[2]) {
                ans++;
            }
        }
        return ans;
    }

    private int[] calcDis(int start, List<Integer>[] g) {
        int[] dis = new int[g.length];
        dfs(start, -1, g, dis);
        return dis;
    }

    private void dfs(int v, int fa, List<Integer>[] g, int[] dis) {
        for (int w : g[v]) {
            if (w != fa) {
                dis[w] = dis[v] + 1;
                dfs(w, v, g, dis);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int specialNodes(int n, vector<vector<int>>& edges, int x, int y, int z) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            int v = e[0], w = e[1];
            g[v].push_back(w);
            g[w].push_back(v);
        }

        auto calc_dis = [&](int start) -> vector<int> {
            vector<int> dis(n);
            auto dfs = [&](this auto&& dfs, int v, int fa) -> void {
                for (int w : g[v]) {
                    if (w != fa) {
                        dis[w] = dis[v] + 1;
                        dfs(w, v);
                    }
                }
            };
            dfs(start, -1);
            return dis;
        };

        auto dx = calc_dis(x);
        auto dy = calc_dis(y);
        auto dz = calc_dis(z);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            vector<int> a = {dx[i], dy[i], dz[i]};
            ranges::sort(a);
            if (1LL * a[0] * a[0] + 1LL * a[1] * a[1] == 1LL * a[2] * a[2]) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func specialNodes(n int, edges [][]int, x, y, z int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	calcDis := func(start int) []int {
		dis := make([]int, n)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					dis[w] = dis[v] + 1
					dfs(w, v)
				}
			}
		}
		dfs(start, -1)
		return dis
	}

	dx := calcDis(x)
	dy := calcDis(y)
	dz := calcDis(z)

	for i := range n {
		a := []int{dx[i], dy[i], dz[i]}
		slices.Sort(a)
		if a[0]*a[0]+a[1]*a[1] == a[2]*a[2] {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
