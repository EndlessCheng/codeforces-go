下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def bestTower(self, towers: List[List[int]], center: List[int], radius: int) -> List[int]:
        cx, cy = center
        ans = (1, -1, -1)
        for x, y, q in towers:
            if abs(x - cx) + abs(y - cy) <= radius:
                ans = min(ans, (-q, x, y))  # 加个负号，变成求 q 的最大值
        return [ans[1], ans[2]]
```

```java [sol-Java]
class Solution {
    public int[] bestTower(int[][] towers, int[] center, int radius) {
        int cx = center[0];
        int cy = center[1];
        int maxQ = -1;
        int minX = -1;
        int minY = -1;
        for (int[] t : towers) {
            int x = t[0];
            int y = t[1];
            int q = t[2];
            if (Math.abs(x - cx) + Math.abs(y - cy) <= radius
                    && (q > maxQ || q == maxQ && (x < minX || x == minX && y < minY))) {
                maxQ = q;
                minX = x;
                minY = y;
            }
        }
        return new int[]{minX, minY};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> bestTower(vector<vector<int>>& towers, vector<int>& center, int radius) {
        int cx = center[0], cy = center[1];
        auto ans = tuple{1, -1, -1};
        for (auto& t : towers) {
            int x = t[0], y = t[1], q = t[2];
            if (abs(x - cx) + abs(y - cy) <= radius) {
                ans = min(ans, tuple{-q, x, y}); // 加个负号，变成求 q 的最大值
            }
        }
        return {get<1>(ans), get<2>(ans)};
    }
};
```

```go [sol-Go]
func bestTower(towers [][]int, center []int, radius int) []int {
	cx, cy := center[0], center[1]
	maxQ, minX, minY := -1, -1, -1
	for _, t := range towers {
		x, y, q := t[0], t[1], t[2]
		if abs(x-cx)+abs(y-cy) <= radius &&
			(q > maxQ || q == maxQ && (x < minX || x == minX && y < minY)) {
			maxQ, minX, minY = q, x, y
		}
	}
	return []int{minX, minY}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{towers}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
