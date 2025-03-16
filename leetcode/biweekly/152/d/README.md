**前置题目**：[3425. 最长特殊路径](https://leetcode.cn/problems/longest-special-path/)，[我的题解](https://leetcode.cn/problems/longest-special-path/solutions/3051377/shu-shang-hua-chuang-pythonjavacgo-by-en-rh5m/)。

书接上回，仍然用滑动窗口思考。

本题允许窗口中的一个元素（颜色）出现两次，我们记录这个颜色在窗口中的更靠上的深度 $\textit{last}_1$。

不断往下递归（相当于窗口往右扩大），每次访问一个新的节点，设这个节点颜色上一次出现的深度为 $\textit{last}_2$。那么窗口左端点 $\textit{topDepth}$ 必须大于等于

$$
\min(\textit{last}_1,\textit{last}_2) + 1
$$

以保证窗口中只有一个颜色出现两次。用上式更新窗口左端点 $\textit{topDepth}$ 的最大值。为什么是更新最大值而不是直接替换？因为 $\textit{last}_2$ 有可能在窗口外面，此时 $\textit{topDepth}$ 是不变的。

此外，如果 $\textit{last}_2$ 更大，意味着我们找到了一个新的出现两次的颜色，用 $\textit{last}_2$ 替换 $\textit{last}_1$。

**总结**：只需要在 3425 题的基础上，多加一个参数 $\textit{last}_1$，改几处代码，其余逻辑不变。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1i6Q8YUEtN/?t=27m27s)，欢迎点赞关注~

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

        def dfs(x: int, fa: int, top_depth: int, last1: int) -> None:
            color = nums[x]
            last2 = last_depth.get(color, 0)
            top_depth = max(top_depth, min(last1, last2))  # 相较 3425 题，维护窗口左端点的逻辑变了

            nonlocal ans
            ans = max(ans, (dis[-1] - dis[top_depth], top_depth - len(dis)))

            last_depth[color] = len(dis)
            for y, w in g[x]:
                if y != fa:
                    dis.append(dis[-1] + w)
                    dfs(y, x, top_depth, max(last1, last2))  # 相较 3425 题，额外维护 last1
                    dis.pop()
            last_depth[color] = last2

        dfs(0, -1, 0, 0)
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
        dfs(0, -1, 0, 0, g, nums, dis, lastDepth);
        return new int[]{maxLen, minNodes};
    }

    private void dfs(int x, int fa, int topDepth, int last1, List<int[]>[] g, int[] nums, List<Integer> dis, Map<Integer, Integer> lastDepth) {
        int color = nums[x];
        int last2 = lastDepth.getOrDefault(color, 0);
        // 相较 3425 题，维护窗口左端点 topDepth 的逻辑变了
        topDepth = Math.max(topDepth, Math.min(last1, last2));

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
            if (y != fa) {
                dis.add(disX + e[1]);
                // 相较 3425 题，额外维护 last1
                dfs(y, x, topDepth, Math.max(last1, last2), g, nums, dis, lastDepth);
                dis.remove(dis.size() - 1);
            }
        }
        lastDepth.put(color, last2);
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

    // 对于本题，dfs 写外面效率更高（可能是 unordered_map 导致）
    void dfs(int x, int fa, int top_depth, int last1) {
        int color = nums[x];
        int last2 = last_depth[color];
        top_depth = max(top_depth, min(last1, last2)); // 相较 3425 题，维护窗口左端点的逻辑变了

        ans = max(ans, pair(dis.back() - dis[top_depth], top_depth - (int) dis.size()));

        last_depth[color] = dis.size();
        for (auto& [y, w] : g[x]) {
            if (y != fa) {
                dis.push_back(dis.back() + w);
                dfs(y, x, top_depth, max(last1, last2)); // 相较 3425 题，额外维护 last1
                dis.pop_back();
            }
        }
        last_depth[color] = last2;
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
        dfs(0, -1, 0, 0);
        return {ans.first, -ans.second};
    }
};
```

```cpp [sol-C++ lambda 递归（慢）]
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

        auto dfs = [&](this auto&& dfs, int x, int fa, int top_depth, int last1) -> void {
            int color = nums[x];
            int last2 = last_depth[color];
            top_depth = max(top_depth, min(last1, last2)); // 相较 3425 题，维护窗口左端点的逻辑变了
    
            ans = max(ans, pair(dis.back() - dis[top_depth], top_depth - (int) dis.size()));
    
            last_depth[color] = dis.size();
            for (auto& [y, w] : g[x]) {
                if (y != fa) {
                    dis.push_back(dis.back() + w);
                    dfs(y, x, top_depth, max(last1, last2)); // 相较 3425 题，额外维护 last1
                    dis.pop_back();
                }
            }
            last_depth[color] = last2;
        };

        dfs(0, -1, 0, 0);
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
    // 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了，下面不需要再 +1
    lastDepth := map[int]int{}

    var dfs func(int, int, int, int)
    dfs = func(x, fa, topDepth, last1 int) {
        color := nums[x]
        last2 := lastDepth[color]
        topDepth = max(topDepth, min(last1, last2)) // 相较 3425 题，维护窗口左端点的逻辑变了

        length := dis[len(dis)-1] - dis[topDepth]
        nodes := len(dis) - topDepth
        if length > maxLen || length == maxLen && nodes < minNodes {
            maxLen = length
            minNodes = nodes
        }

        lastDepth[color] = len(dis)
        for _, e := range g[x] {
            y := e.to
            if y != fa {
                dis = append(dis, dis[len(dis)-1]+e.weight)
                dfs(y, x, topDepth, max(last1, last2)) // 相较 3425 题，额外维护 last1
                dis = dis[:len(dis)-1]
            }
        }
        lastDepth[color] = last2
    }

    dfs(0, -1, 0, 0)
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
