竖切的时候，答案与纵坐标，也就是矩形的高无关，我们可以把每个矩形视作一个区间 $[\textit{start}_x, \textit{end}_x]$。

问题变成把这 $n$ 个区间 [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) 后，区间的个数是否 $\ge 3$。如何合并区间？请看 [我的题解](https://leetcode.cn/problems/merge-intervals/solutions/2798138/jian-dan-zuo-fa-yi-ji-wei-shi-yao-yao-zh-f2b3/comments/2439822/)。

横切同理，把每个矩形视作一个区间 $[\textit{start}_y, \textit{end}_y]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ifkqYjEvc/?t=12m20s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def check(self, intervals: List[Tuple[int, int]]) -> bool:
        intervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序
        cnt = max_r = 0
        for l, r in intervals:
            if l >= max_r:  # 新区间
                cnt += 1
            if r > max_r:
                max_r = r  # 更新右端点最大值（手写 if 效率更高）
        return cnt >= 3  # 也可以在循环中提前退出，但是慢一些

    def checkValidCuts(self, _: int, rectangles: List[List[int]]) -> bool:
        return self.check([(sx, ex) for sx, _, ex, _ in rectangles]) or \
               self.check([(sy, ey) for _, sy, _, ey in rectangles])
```

```java [sol-Java]
class Solution {
    boolean checkValidCuts(int n, int[][] rectangles) {
        int m = rectangles.length;
        int[][] a = new int[m][2];
        int[][] b = new int[m][2];
        for (int i = 0; i < m; i++) {
            int[] rect = rectangles[i];
            a[i][0] = rect[0];
            a[i][1] = rect[2];
            b[i][0] = rect[1];
            b[i][1] = rect[3];
        }
        return check(a) || check(b);
    }

    private boolean check(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> a[0] - b[0]); // 按照左端点从小到大排序
        int cnt = 0;
        int maxR = 0;
        for (int[] interval : intervals) {
            if (interval[0] >= maxR) { // 新区间
                cnt++;
            }
            maxR = Math.max(maxR, interval[1]); // 更新右端点最大值
        }
        return cnt >= 3;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool check(vector<pair<int, int>>& intervals) {
        ranges::sort(intervals, {}, [](auto& a) { return a.first; });// 按照左端点从小到大排序
        int cnt = 0, max_r = 0;
        for (auto& [l, r] : intervals) {
            if (l >= max_r) { // 新区间
                cnt++;
            }
            max_r = max(max_r, r); // 更新右端点最大值
        }
        return cnt >= 3; // 也可以在循环中提前退出
    }

    bool checkValidCuts(int, vector<vector<int>>& rectangles) {
        vector<pair<int, int>> a, b;
        for (auto& rect : rectangles) {
            a.emplace_back(rect[0], rect[2]);
            b.emplace_back(rect[1], rect[3]);
        }
        return check(a) || check(b);
    }
};
```

```go [sol-Go]
type pair struct{ l, r int }

func check(intervals []pair) bool {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(a, b pair) int { return a.l - b.l })
	cnt, maxR := 0, 0
	for _, p := range intervals {
		if p.l >= maxR { // 新区间
			cnt++
		}
		maxR = max(maxR, p.r) // 更新右端点最大值
	}
	return cnt >= 3 // 也可以在循环中提前退出
}

func checkValidCuts(_ int, rectangles [][]int) bool {
	a := make([]pair, len(rectangles))
	b := make([]pair, len(rectangles))
	for i, rect := range rectangles {
		a[i] = pair{rect[0], rect[2]}
		b[i] = pair{rect[1], rect[3]}
	}
	return check(a) || check(b)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 是 $\textit{rectangles}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(m)$。

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
