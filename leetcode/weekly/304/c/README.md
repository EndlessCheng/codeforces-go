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
