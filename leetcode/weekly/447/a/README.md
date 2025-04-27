用哈希表记录每一行和每一列的点。

把所有行和列都排序。

对于每个建筑，在所在行和列二分，如果左右上下都有建筑，答案加一。

**注**：由于我们是用哈希表实现的，所以这个做法的时空复杂度和 $n$ 无关。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def countCoveredBuildings(self, n: int, buildings: List[List[int]]) -> int:
        row = defaultdict(list)
        col = defaultdict(list)
        for x, y in buildings:
            row[x].append(y)
            col[y].append(x)

        for a in row.values():
            a.sort()
        for a in col.values():
            a.sort()

        def is_inner(a: List[int], x: int) -> bool:
            return 0 < bisect_left(a, x) < len(a) - 1  # 左右都有建筑

        ans = 0
        for x, y in buildings:
            if is_inner(row[x], y) and is_inner(col[y], x):
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countCoveredBuildings(int n, int[][] buildings) {
        Map<Integer, List<Integer>> row = new HashMap<>();
        Map<Integer, List<Integer>> col = new HashMap<>();
        for (int[] p : buildings) {
            int x = p[0], y = p[1];
            row.computeIfAbsent(x, k -> new ArrayList<>()).add(y);
            col.computeIfAbsent(y, k -> new ArrayList<>()).add(x);
        }

        for (List<Integer> a : row.values()) {
            Collections.sort(a);
        }
        for (List<Integer> a : col.values()) {
            Collections.sort(a);
        }

        int ans = 0;
        for (int[] p : buildings) {
            int x = p[0], y = p[1];
            if (isInner(row.get(x), y) && isInner(col.get(y), x)) {
                ans++;
            }
        }
        return ans;
    }

    private boolean isInner(List<Integer> a, int x) {
        int i = lowerBound(a, x);
        return 0 < i && i < a.size() - 1; // 左右都有建筑
    }

    private int lowerBound(List<Integer> a, int x) {
        int left = -1, right = a.size();
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (a.get(mid) >= x) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool is_inner(vector<int>& a, int x) {
        int i = ranges::lower_bound(a, x) - a.begin();
        return 0 < i && i + 1 < a.size(); // 左右都有建筑
    }

    int countCoveredBuildings(int _, vector<vector<int>>& buildings) {
        unordered_map<int, vector<int>> row, col;
        for (auto& p : buildings) {
            int x = p[0], y = p[1];
            row[x].push_back(y);
            col[y].push_back(x);
        }

        for (auto& [_, a] : row) {
            ranges::sort(a);
        }
        for (auto& [_, a] : col) {
            ranges::sort(a);
        }

        int ans = 0;
        for (auto& p : buildings) {
            int x = p[0], y = p[1];
            if (is_inner(row[x], y) && is_inner(col[y], x)) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func isInner(a []int, x int) bool {
	i := sort.SearchInts(a, x)
	return 0 < i && i < len(a)-1 // 左右都有建筑
}

func countCoveredBuildings(_ int, buildings [][]int) (ans int) {
	row := map[int][]int{}
	col := map[int][]int{}
	for _, p := range buildings {
		x, y := p[0], p[1]
		row[x] = append(row[x], y)
		col[y] = append(col[y], x)
	}

	for _, a := range row {
		slices.Sort(a)
	}
	for _, a := range col {
		slices.Sort(a)
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		if isInner(row[x], y) && isInner(col[y], x) {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 是 $\textit{buildings}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

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
