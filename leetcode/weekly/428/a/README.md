遍历的过程中，维护答案 $\textit{idx}$，以及最大时间差 $\textit{maxDiff}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1pnqZYKEqr/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def buttonWithLongestTime(self, events: List[List[int]]) -> int:
        idx, max_diff = events[0]
        for (_, t1), (i, t2) in pairwise(events):
            d = t2 - t1
            if d > max_diff or d == max_diff and i < idx:
                idx, max_diff = i, d
        return idx
```

```java [sol-Java]
class Solution {
    int buttonWithLongestTime(int[][] events) {
        int idx = events[0][0];
        int maxDiff = events[0][1];
        for (int i = 1; i < events.length; i++) {
            int[] p = events[i - 1];
            int[] q = events[i];
            int d = q[1] - p[1];
            if (d > maxDiff || d == maxDiff && q[0] < idx) {
                idx = q[0];
                maxDiff = d;
            }
        }
        return idx;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int buttonWithLongestTime(vector<vector<int>>& events) {
        int idx = events[0][0], max_diff = events[0][1];
        for (int i = 1; i < events.size(); i++) {
            auto& p = events[i - 1];
            auto& q = events[i];
            int d = q[1] - p[1];
            if (d > max_diff || d == max_diff && q[0] < idx) {
                idx = q[0];
                max_diff = d;
            }
        }
        return idx;
    }
};
```

```go [sol-Go]
func buttonWithLongestTime(events [][]int) int {
	idx, maxDiff := events[0][0], events[0][1]
	for i := 1; i < len(events); i++ {
		p, q := events[i-1], events[i]
		d := q[1] - p[1]
		if d > maxDiff || d == maxDiff && q[0] < idx {
			idx, maxDiff = q[0], d
		}
	}
	return idx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{events}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
