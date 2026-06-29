**前置题目**：[56. 合并区间](https://leetcode.cn/problems/merge-intervals/)，[我的题解](https://leetcode.cn/problems/merge-intervals/solutions/2798138/jian-dan-zuo-fa-yi-ji-wei-shi-yao-yao-zh-f2b3/comments/2439822/)。

本题区间不能包含 $[\textit{freeStart}, \textit{freeEnd}]$ 中的整数。

对于一个合并后的区间 $[\ell,r]$，讨论 $[\ell,r]$ 与 $[\textit{freeStart}, \textit{freeEnd}]$ 的位置关系：

- 如果 $r < \textit{freeStart}$ 或者 $\ell > \textit{freeEnd}$，那么两区间不相交，$[\ell,r]$ 不变。
- 否则：
    - 如果 $\ell < \textit{freeStart}$，那么一个剩余区间为 $[\ell, \textit{freeStart}-1]$。
    - 如果 $r > \textit{freeEnd}$，那么另一个剩余区间为 $[\textit{freeEnd}+1, r]$。

[本题视频讲解](https://www.bilibili.com/video/BV1wwTu6EEcG/?t=3m37s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def filterOccupiedIntervals(self, occupiedIntervals: list[list[int]], freeStart: int, freeEnd: int) -> list[list[int]]:
        occupiedIntervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序
        ans = []

        def add(l: int, r: int) -> None:
            if r < freeStart or l > freeEnd:  # 不相交
                ans.append([l, r])
                return
            if l < freeStart:
                ans.append([l, freeStart - 1])  # 余留前缀
            if r > freeEnd:
                ans.append([freeEnd + 1, r])  # 余留后缀

        left, max_r = occupiedIntervals[0]
        for l, r in occupiedIntervals[1:]:  # 从第二个区间开始
            if l - 1 > max_r:  # 发现一个新区间
                add(left, max_r)  # 先把旧的加入答案
                left = l  # 记录新区间左端点
            max_r = max(max_r, r)
        add(left, max_r)

        return ans
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> filterOccupiedIntervals(int[][] occupiedIntervals, int freeStart, int freeEnd) {
        Arrays.sort(occupiedIntervals, (a, b) -> a[0] - b[0]); // 按照左端点从小到大排序
        List<List<Integer>> ans = new ArrayList<>();

        int left = occupiedIntervals[0][0];
        int maxR = occupiedIntervals[0][1];
        for (int i = 1; i < occupiedIntervals.length; i++) { // 从第二个区间开始
            int l = occupiedIntervals[i][0];
            int r = occupiedIntervals[i][1];
            if (l - 1 > maxR) { // 发现一个新区间
                add(ans, left, maxR, freeStart, freeEnd); // 先把旧的加入答案
                left = l; // 记录新区间左端点
            }
            maxR = Math.max(maxR, r);
        }
        add(ans, left, maxR, freeStart, freeEnd);

        return ans;
    }

    private void add(List<List<Integer>> ans, int l, int r, int freeStart, int freeEnd) {
        if (r < freeStart || l > freeEnd) { // 不相交
            ans.add(List.of(l, r));
            return;
        }
        if (l < freeStart) {
            ans.add(List.of(l, freeStart - 1)); // 余留前缀
        }
        if (r > freeEnd) {
            ans.add(List.of(freeEnd + 1, r)); // 余留后缀
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> filterOccupiedIntervals(vector<vector<int>>& occupiedIntervals, int freeStart, int freeEnd) {
        ranges::sort(occupiedIntervals, {}, [](auto& a) { return a[0]; }); // 按照左端点从小到大排序
        vector<vector<int>> ans;

        auto add = [&](int l, int r) -> void {
            if (r < freeStart || l > freeEnd) { // 不相交
                ans.push_back({l, r});
                return;
            }
            if (l < freeStart) {
                ans.push_back({l, freeStart - 1}); // 余留前缀
            }
            if (r > freeEnd) {
                ans.push_back({freeEnd + 1, r}); // 余留后缀
            }
        };

        int left = occupiedIntervals[0][0];
        int max_r = occupiedIntervals[0][1];
        for (int i = 1; i < occupiedIntervals.size(); i++) { // 从第二个区间开始
            int l = occupiedIntervals[i][0];
            int r = occupiedIntervals[i][1];
            if (l - 1 > max_r) { // 发现一个新区间
                add(left, max_r); // 先把旧的加入答案
                left = l; // 记录新区间左端点
            }
            max_r = max(max_r, r);
        }
        add(left, max_r);

        return ans;
    }
};
```

```go [sol-Go]
func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) (ans [][]int) {
	slices.SortFunc(occupiedIntervals, func(a, b []int) int { return a[0] - b[0] }) // 按照左端点从小到大排序

	add := func(l, r int) {
		if r < freeStart || l > freeEnd { // 不相交
			ans = append(ans, []int{l, r})
			return
		}
		if l < freeStart {
			ans = append(ans, []int{l, freeStart - 1}) // 余留前缀
		}
		if r > freeEnd {
			ans = append(ans, []int{freeEnd + 1, r}) // 余留后缀
		}
	}

	left := occupiedIntervals[0][0]
	maxR := occupiedIntervals[0][1]
	for _, p := range occupiedIntervals[1:] { // 从第二个区间开始
		l, r := p[0], p[1]
		if l-1 > maxR { // 发现一个新区间
			add(left, maxR) // 先把旧的加入答案
			left = l // 记录新区间左端点
		}
		maxR = max(maxR, r)
	}
	add(left, maxR)

	return
}
```

## 写法二

```py [sol-Python3]
class Solution:
    def filterOccupiedIntervals(self, occupiedIntervals: list[list[int]], freeStart: int, freeEnd: int) -> list[list[int]]:
        occupiedIntervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序
        ans = []
        n = len(occupiedIntervals)
        left, right = inf, 0

        for i, (l, r) in enumerate(occupiedIntervals):
            left = min(left, l)
            right = max(right, r)
            if i == n - 1 or occupiedIntervals[i + 1][0] - 1 > right:
                if right < freeStart or left > freeEnd:  # 不相交
                    ans.append([left, right])
                else:
                    if left < freeStart:
                        ans.append([left, freeStart - 1])  # 余留前缀
                    if right > freeEnd:
                        ans.append([freeEnd + 1, right])  # 余留后缀
                left = inf

        return ans
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> filterOccupiedIntervals(int[][] occupiedIntervals, int freeStart, int freeEnd) {
        Arrays.sort(occupiedIntervals, (a, b) -> a[0] - b[0]); // 按照左端点从小到大排序
        List<List<Integer>> ans = new ArrayList<>();
        int n = occupiedIntervals.length;
        int left = Integer.MAX_VALUE;
        int right = 0;

        for (int i = 0; i < n; i++) {
            int[] p = occupiedIntervals[i];
            left = Math.min(left, p[0]);
            right = Math.max(right, p[1]);
            if (i == n - 1 || occupiedIntervals[i + 1][0] - 1 > right) {
                if (right < freeStart || left > freeEnd) { // 不相交
                    ans.add(List.of(left, right));
                } else {
                    if (left < freeStart) {
                        ans.add(List.of(left, freeStart - 1)); // 余留前缀
                    }
                    if (right > freeEnd) {
                        ans.add(List.of(freeEnd + 1, right)); // 余留后缀
                    }
                }
                left = Integer.MAX_VALUE;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> filterOccupiedIntervals(vector<vector<int>>& occupiedIntervals, int freeStart, int freeEnd) {
        ranges::sort(occupiedIntervals, {}, [](auto& a) { return a[0]; }); // 按照左端点从小到大排序
        vector<vector<int>> ans;
        int n = occupiedIntervals.size();
        int left = INT_MAX, right = 0;

        for (int i = 0; i < n; i++) {
            auto& p = occupiedIntervals[i];
            left = min(left, p[0]);
            right = max(right, p[1]);
            if (i == n - 1 || occupiedIntervals[i + 1][0] - 1 > right) {
                if (right < freeStart || left > freeEnd) { // 不相交
                    ans.push_back({left, right});
                } else {
                    if (left < freeStart) {
                        ans.push_back({left, freeStart - 1}); // 余留前缀
                    }
                    if (right > freeEnd) {
                        ans.push_back({freeEnd + 1, right}); // 余留后缀
                    }
                }
                left = INT_MAX;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) (ans [][]int) {
	slices.SortFunc(occupiedIntervals, func(a, b []int) int { return a[0] - b[0] }) // 按照左端点从小到大排序

	left, right := math.MaxInt, 0
	for i, p := range occupiedIntervals {
		left = min(left, p[0])
		right = max(right, p[1])
		if i == len(occupiedIntervals)-1 || occupiedIntervals[i+1][0]-1 > right {
			if right < freeStart || left > freeEnd { // 不相交
				ans = append(ans, []int{left, right})
			} else {
				if left < freeStart {
					ans = append(ans, []int{left, freeStart - 1}) // 余留前缀
				}
				if right > freeEnd {
					ans = append(ans, []int{freeEnd + 1, right}) // 余留后缀
				}
			}
			left = math.MaxInt
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{occupiedIntervals}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，返回值不计入。

## 专题训练

见下面贪心题单的「**§2.5 合并区间**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
