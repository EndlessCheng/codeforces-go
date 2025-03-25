## 前言：关于图的性质

「每个节点至多有一条出边」意味着，对于图中任意一个大小为 $m$ 的连通块，有 $m$ 个点，每个点至多出去一条边，所以连通块至多有 $m$ 条边。

我们知道，$m$ 个点 $m-1$ 条边的连通图是一棵树，在树上增加一条有向边，至多会形成一个环。（这样的图叫做**内向基环树**）

## 思路

![lc2360.png](https://pic.leetcode.cn/1742870244-vqHncL-lc2360.png)

假设你在跑步，操场跑道是 $3\to 2\to 4\to 3$（上图红线）。

你想知道跑一圈要多久。你从节点 $0$ 开始跑，跑到节点 $3$ 的时候，记录当前时间为 $t_1=2$，再次跑到节点 $3$ 的时候，记录当前时间为 $t_2=5$，那么跑一圈就需要 $t_2-t_1=5-2=3$ 个单位时间。

如果每访问一个节点，计时器就加一，那么 $t_2-t_1$ 就是跑道长度，即环长。

## 算法

初始时间为 $\textit{curTime}=1$。遍历图，每访问到一个新的节点 $x$，就记录**首次访问时间** $\textit{visTime}[x]=\textit{curTime}$，然后将 $\textit{curTime}$ 加一。

假设我们从节点 $i$ 开始。首先记录开始时间 $\textit{startTime}=\textit{curTime}$，然后继续走，如果走到死路，或者找到了一个之前访问过的点 $x$，则退出循环。

退出循环后，分类讨论：

- 如果 $\textit{visTime}[x] < \textit{startTime}$，说明 $x$ 不是在本轮循环中访问的。例如上图从节点 $0$ 开始，访问节点 $0,3,2,4$。然后接着从节点 $1$ 开始，访问节点 $3$，发现 $\textit{visTime}[3]$ 比访问节点 $1$ 的时间还要早，那么包含节点 $3$ 的环长我们之前已经计算过了，无需再次计算。
- 如果 $\textit{visTime}[x] \ge \textit{startTime}$，说明 $x$ 是在本轮循环中访问的，且被访问了两次。这只有一种可能，就是 $x$ 在环上。根据前后两次访问 $x$ 的时间差，就能算出环长，即 $\textit{curTime}-\textit{visTime}[x]$。

> 注：本题保证每个连通块至多有一个环，所以可以根据时间差算出环长。如果没有这个保证，时间差算出的可能不是最长环。一般图的最长环是 NP-hard 问题。

取所有环长的最大值作为答案。如果图中无环，则返回 $-1$。

```py [sol-Python3]
class Solution:
    def longestCycle(self, edges: List[int]) -> int:
        n = len(edges)
        ans = -1
        cur_time = 1  # 当前时间
        vis_time = [0] * n  # 首次访问 x 的时间
        for x in range(n):
            start_time = cur_time  # 本轮循环的开始时间
            while x != -1 and vis_time[x] == 0:  # 没有访问过 x
                vis_time[x] = cur_time  # 记录访问 x 的时间
                cur_time += 1
                x = edges[x]  # 访问下一个节点
            if x != -1 and vis_time[x] >= start_time:  # x 在本轮循环中访问了两次，说明 x 在环上
                ans = max(ans, cur_time - vis_time[x])  # 前后两次访问 x 的时间差，即为环长
        return ans  # 如果没有找到环，返回的是 ans 的初始值 -1
```

```java [sol-Java]
class Solution {
    public int longestCycle(int[] edges) {
        int n = edges.length;
        int ans = -1;
        int curTime = 1; // 当前时间
        int[] visTime = new int[n]; // 首次访问 x 的时间
        for (int i = 0; i < n; i++) {
            int x = i;
            int startTime = curTime; // 本轮循环的开始时间
            while (x != -1 && visTime[x] == 0) { // 没有访问过 x
                visTime[x] = curTime++; // 记录访问 x 的时间
                x = edges[x]; // 访问下一个节点
            }
            if (x != -1 && visTime[x] >= startTime) { // x 在本轮循环中访问了两次，说明 x 在环上
                ans = Math.max(ans, curTime - visTime[x]); // 前后两次访问 x 的时间差，即为环长
            }
        }
        return ans; // 如果没有找到环，返回的是 ans 的初始值 -1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCycle(vector<int>& edges) {
        int n = edges.size();
        int ans = -1;
        int cur_time = 1; // 当前时间
        vector<int> vis_time(n); // 首次访问 x 的时间
        for (int i = 0; i < n; i++) {
            int x = i;
            int start_time = cur_time; // 本轮循环的开始时间
            while (x != -1 && vis_time[x] == 0) { // 没有访问过 x
                vis_time[x] = cur_time++; // 记录访问 x 的时间
                x = edges[x]; // 访问下一个节点
            }
            if (x != -1 && vis_time[x] >= start_time) { // x 在本轮循环中访问了两次，说明 x 在环上
                ans = max(ans, cur_time - vis_time[x]); // 前后两次访问 x 的时间差，即为环长
            }
        }
        return ans; // 如果没有找到环，返回的是 ans 的初始值 -1
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int longestCycle(int* edges, int edgesSize) {
    int ans = -1;
    int cur_time = 1; // 当前时间
    int* vis_time = calloc(edgesSize, sizeof(int)); // 首次访问 x 的时间
    for (int i = 0; i < edgesSize; i++) {
        int x = i;
        int start_time = cur_time; // 本轮循环的开始时间
        while (x != -1 && vis_time[x] == 0) { // 没有访问过 x
            vis_time[x] = cur_time++; // 记录访问 x 的时间
            x = edges[x]; // 访问下一个节点
        }
        if (x != -1 && vis_time[x] >= start_time) { // x 在本轮循环中访问了两次，说明 x 在环上
            ans = MAX(ans, cur_time - vis_time[x]); // 前后两次访问 x 的时间差，即为环长
        }
    }
    free(vis_time);
    return ans; // 如果没有找到环，返回的是 ans 的初始值 -1
}
```

```go [sol-Go]
func longestCycle(edges []int) int {
    ans := -1
    curTime := 1 // 当前时间
    visTime := make([]int, len(edges)) // 首次访问 x 的时间
    for x := range edges {
        startTime := curTime // 本轮循环的开始时间
        for x != -1 && visTime[x] == 0 { // 没有访问过 x
            visTime[x] = curTime // 记录访问 x 的时间
            curTime++
            x = edges[x] // 访问下一个节点
        }
        if x != -1 && visTime[x] >= startTime { // x 在本轮循环中访问了两次，说明 x 在环上
            ans = max(ans, curTime-visTime[x]) // 前后两次访问 x 的时间差，即为环长
        }
    }
    return ans // 如果没有找到环，返回的是 ans 的初始值 -1
}
```

```js [sol-JavaScript]
var longestCycle = function(edges) {
    const n = edges.length;    
    const visTime = Array(n).fill(0); // 首次访问 x 的时间
    let curTime = 1; // 当前时间
    let ans = -1;
    for (let i = 0; i < n; i++) {
        let x = i;
        let startTime = curTime; // 本轮循环的开始时间
        while (x !== -1 && visTime[x] === 0) { // 没有访问过 x
            visTime[x] = curTime++; // 记录访问 x 的时间
            x = edges[x]; // 访问下一个节点
        }
        if (x !== -1 && visTime[x] >= startTime) { // x 在本轮循环中访问了两次，说明 x 在环上
            ans = Math.max(ans, curTime - visTime[x]); // 前后两次访问 x 的时间差，即为环长
        }
    }
    return ans; // 如果没有找到环，返回的是 ans 的初始值 -1
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_cycle(edges: Vec<i32>) -> i32 {
        let n = edges.len();
        let mut ans = -1;
        let mut cur_time = 1; // 当前时间
        let mut vis_time = vec![0; n]; // 首次访问 x 的时间
        for mut x in 0..n {
            let start_time = cur_time; // 本轮循环的开始时间
            while x < n && vis_time[x] == 0 { // 没有访问过 x
                vis_time[x] = cur_time; // 记录访问 x 的时间
                cur_time += 1;
                x = edges[x] as usize; // 访问下一个节点，如果是 -1 则转换成 MAX
            }
            if x < n && vis_time[x] >= start_time { // x 在本轮循环中访问了两次，说明 x 在环上
                ans = ans.max(cur_time - vis_time[x]); // 前后两次访问 x 的时间差，即为环长
            }
        }
        ans // 如果没有找到环，返回的是 ans 的初始值 -1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{edges}$ 的长度。虽然写了个二重循环，但每个节点只会记录一次访问时间，所以二重循环的总循环次数是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

1. 改成返回最长环上的所有节点（返回一个列表），要怎么做？欢迎在评论区分享你的思路/代码。
2. 改成最短环呢？对于一般图，如何计算最短环？见 [2608. 图中的最短环](https://leetcode.cn/problems/shortest-cycle-in-a-graph/)。

更多相似题目，见下面图论题单中的「**§2.3 基环树**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
