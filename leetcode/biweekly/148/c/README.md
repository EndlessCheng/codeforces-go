## 引入

从特殊到一般。想一想，如果树是一条**链**，要怎么做？

链等同于数组，问题相当于在一个数组中找没有重复元素的连续子数组，比如 [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)。

这启发我们用 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 解决。

## 初步思路

数组上的滑动窗口，可以外层循环枚举子数组右端点，内层循环维护子数组左端点。

在树上滑窗，我们可以枚举路径最下面的节点（类似子数组右端点），同时维护路径最上面的节点（类似子数组左端点）：

- 如果发现路径中有重复颜色，那么向下移动路径最上面的节点。

但在树上这么做的话，对于链+星的「扫帚型」树，会跑到 $\mathcal{O}(n^2)$。

怎么优化？

## 更快的思路

记录每种颜色 $\textit{color}$ 最近一次出现的位置（深度）$\textit{lastDepth}[\textit{color}]$，那么路径最上面节点的深度，就是路径中所有节点的 $\textit{lastDepth}[\textit{color}]$ 的最大值加一。这样就可以 $\mathcal{O}(1)$ 维护路径最上面节点的深度了。注意这里只需要维护深度。

把路径最上面节点的深度记作 $\textit{topDepth}$。根节点 $0$ 的深度为 $0$。

## 实现细节

对于从根到当前节点的路径，我们用一个栈 $\textit{dis}$ 维护根到各个节点的距离。那么：

- 路径长度：根到当前节点的距离，减去根到路径最上面节点的距离。前者是 $\textit{dis}$ 的栈顶，后者是 $\textit{dis}[\textit{topDepth}]$。
- 路径节点个数：当前节点的深度加一，减去 $\textit{topDepth}$。前者是当前 $\textit{dis}$ 的大小。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1xBwBeEEie/?t=7m25s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestSpecialPath(self, edges: List[List[int]], nums: List[int]) -> List[int]:
        g = [[] for _ in nums]
        for x, y, w in edges:
            g[x].append((y, w))
            g[y].append((x, w))

        ans = (-1, 0)
        dis = [0]
        last_depth = {}  # 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了

        def dfs(x: int, fa: int, top_depth: int) -> None:
            color = nums[x]
            old_depth = last_depth.get(color, 0)
            top_depth = max(top_depth, old_depth)

            nonlocal ans
            # 把 len(dis) - top_depth 取反，这样 max 算的是最小值
            ans = max(ans, (dis[-1] - dis[top_depth], top_depth - len(dis)))

            last_depth[color] = len(dis)
            for y, w in g[x]:
                if y != fa:  # 避免访问父节点
                    dis.append(dis[-1] + w)
                    dfs(y, x, top_depth)
                    dis.pop()  # 恢复现场
            last_depth[color] = old_depth  # 恢复现场

        dfs(0, -1, 0)
        return [ans[0], -ans[1]]
```

```java [sol-Java]
class Solution {
    private int maxLen = -1;
    private int minNodes = 0;

    public int[] longestSpecialPath(int[][] edges, int[] nums) {
        List<int[]>[] g = new ArrayList[nums.length];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            int w = e[2];
            g[x].add(new int[]{y, w});
            g[y].add(new int[]{x, w});
        }

        List<Integer> dis = new ArrayList<>();
        dis.add(0);
        // 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了
        Map<Integer, Integer> lastDepth = new HashMap<>();
        dfs(0, -1, 0, g, nums, dis, lastDepth);
        return new int[]{maxLen, minNodes};
    }

    private void dfs(int x, int fa, int topDepth, List<int[]>[] g, int[] nums, List<Integer> dis, Map<Integer, Integer> lastDepth) {
        int color = nums[x];
        int oldDepth = lastDepth.getOrDefault(color, 0);
        topDepth = Math.max(topDepth, oldDepth);

        int disX = dis.get(dis.size() - 1);
        int len = disX - dis.get(topDepth);
        int nodes = dis.size() - topDepth;
        if (len > maxLen || len == maxLen && nodes < minNodes) {
            maxLen = len;
            minNodes = nodes;
        }

        lastDepth.put(color, dis.size());
        for (int[] e : g[x]) {
            int y = e[0];
            if (y != fa) { // 避免访问父节点
                dis.add(disX + e[1]);
                dfs(y, x, topDepth, g, nums, dis, lastDepth);
                dis.remove(dis.size() - 1); // 恢复现场
            }
        }
        lastDepth.put(color, oldDepth); // 恢复现场
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> nums;
    vector<vector<pair<int, int>>> g;
    pair<int, int> ans = {-1, 0};
    vector<int> dis = {0};
    unordered_map<int, int> last_depth; // 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了

    void dfs(int x, int fa, int top_depth) {
        int color = nums[x];
        int old_depth = last_depth[color];
        top_depth = max(top_depth, old_depth);

        // 把 dis.size() - top_depth 取反，这样 max 算的是最小值
        ans = max(ans, pair(dis.back() - dis[top_depth], top_depth - (int) dis.size()));

        last_depth[color] = dis.size();
        for (auto& [y, w] : g[x]) {
            if (y != fa) { // 避免访问父节点
                dis.push_back(dis.back() + w);
                dfs(y, x, top_depth);
                dis.pop_back(); // 恢复现场
            }
        }
        last_depth[color] = old_depth; // 恢复现场
    }

public:
    vector<int> longestSpecialPath(vector<vector<int>>& edges, vector<int>& nums) {
        g.resize(nums.size());
        for (auto& e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }
        this->nums = nums;
        dfs(0, -1, 0);
        return {ans.first, -ans.second};
    }
};
```

```cpp [sol-C++ lambda 递归]
class Solution {
public:
    vector<int> longestSpecialPath(vector<vector<int>>& edges, vector<int>& nums) {
        vector<vector<pair<int, int>>> g(nums.size());
        for (auto& e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }

        pair<int, int> ans = {-1, 0};
        vector<int> dis = {0};
        unordered_map<int, int> last_depth; // 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了

        auto dfs = [&](this auto&& dfs, int x, int fa, int top_depth) -> void {
            int color = nums[x];
            int old_depth = last_depth[color];
            top_depth = max(top_depth, old_depth);

            // 把 dis.size() - top_depth 取反，这样 max 算的是最小值
            ans = max(ans, pair(dis.back() - dis[top_depth], top_depth - (int) dis.size()));

            last_depth[color] = dis.size();
            for (auto& [y, w] : g[x]) {
                if (y != fa) { // 避免访问父节点
                    dis.push_back(dis.back() + w);
                    dfs(y, x, top_depth);
                    dis.pop_back(); // 恢复现场
                }
            }
            last_depth[color] = old_depth; // 恢复现场
        };

        dfs(0, -1, 0);
        return {ans.first, -ans.second};
    }
};
```

```go [sol-Go]
func longestSpecialPath(edges [][]int, nums []int) []int {
	type edge struct{ to, weight int }
	g := make([][]edge, len(nums))
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	maxLen := -1
	minNodes := 0
	dis := []int{0}
	// 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了
	lastDepth := map[int]int{}

	var dfs func(int, int, int)
	dfs = func(x, fa, topDepth int) {
		color := nums[x]
		oldDepth := lastDepth[color]
		topDepth = max(topDepth, oldDepth)

		length := dis[len(dis)-1] - dis[topDepth]
		nodes := len(dis) - topDepth
		if length > maxLen || length == maxLen && nodes < minNodes {
			maxLen = length
			minNodes = nodes
		}

		lastDepth[color] = len(dis)
		for _, e := range g[x] {
			y := e.to
			if y != fa { // 避免访问父节点
				dis = append(dis, dis[len(dis)-1]+e.weight)
				dfs(y, x, topDepth)
				dis = dis[:len(dis)-1] // 恢复现场
			}
		}
		lastDepth[color] = oldDepth // 恢复现场
	}

	dfs(0, -1, 0)
	return []int{maxLen, minNodes}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个节点恰好访问一次。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面 11 题单中的「**三、一般树**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
