## 方法一：排序 + 二分查找

首先，区间在 $\textit{intervals}$ 中的顺序不影响答案，我们可以先把区间排序。

> 为什么可以排序？例如 $\textit{intervals} = [A,B,C]$，有三对区间 $(A,B),(A,C),(B,C)$。假如重排成 $[B,C,A]$，那么仍然有三对区间 $(B,C),(B,A),(C,A)$。如果区间 $A$ 和 $B$ 相交，那么区间 $B$ 和 $A$ 也相交。排序不改变相交关系。

按照什么顺序排序呢？以左端点为主，还是以右端点为主？

首先转化一下问题，改成求**不相交**的区间对，这样更好求。因为无论按照左端点还是右端点排序，对于一对不相交的区间 $(i,j)$，满足 $i<j$ 的区间 $i$ 都在区间 $j$ 的左侧，且区间 $i$ 的左右端点都小于 $\textit{start}_j$。

不相交的区间 $(i,j)$ 满足 $\textit{end}_i < \textit{start}_j$。如果按照**右端点**升序排序，我们就可以用 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 快速求出有多少个 $\textit{end}_i < \textit{start}_j$。

最后，用所有区间对的个数 $\dfrac{n(n-1)}{2}$ 减去不相交区间对的个数，即为答案。

[本题视频讲解](https://www.bilibili.com/video/BV1MEeB65EjZ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countIntersectingIntervals(self, intervals: list[list[int]]) -> int:
        n = len(intervals)
        intervals.sort(key=lambda p: p[1])  # 按照右端点升序排序

        ans = n * (n - 1) // 2
        for start, _ in intervals:
            # 设 j 是最小的满足 intervals[j][1] >= start 的下标
            # 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
            ans -= bisect_left(intervals, start, key=lambda p: p[1])
        return ans
```

```java [sol-Java]
class Solution {
    public long countIntersectingIntervals(int[][] intervals) {
        int n = intervals.length;
        Arrays.sort(intervals, (a, b) -> a[1] - b[1]); // 按照右端点升序排序

        long ans = (long) n * (n - 1) / 2;
        for (int[] p : intervals) {
            int start = p[0];
            // 设 j 是最小的满足 intervals[j][1] >= start 的下标
            // 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
            ans -= lowerBound(intervals, start);
        }
        return ans;
    }

    private int lowerBound(int[][] intervals, int target) {
        int left = -1;
        int right = intervals.length;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (intervals[mid][1] >= target) {
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
    long long countIntersectingIntervals(vector<vector<int>>& intervals) {
        int n = intervals.size();
        ranges::sort(intervals, {}, [](auto& p) { return p[1]; });  // 按照右端点升序排序

        long long ans = 1LL * n * (n - 1) / 2;
        for (auto& p : intervals) {
            int start = p[0];
            // 设 j 是最小的满足 intervals[j][1] >= start 的下标
            // 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
            ans -= ranges::lower_bound(intervals, start, {}, [](auto& q) { return q[1]; }) - intervals.begin();
        }
        return ans;
    }
};
```

```go [sol-Go]
func countIntersectingIntervals(intervals [][]int) int64 {
	n := len(intervals)
	slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] }) // 按照右端点升序排序

	ans := n * (n - 1) / 2
	for _, p := range intervals {
		start := p[0]
		// 设 j 是最小的满足 intervals[j][1] >= start 的下标
		// 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
		ans -= sort.Search(n, func(j int) bool { return intervals[j][1] >= start })
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{intervals}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：排序 + 双指针

另一种做法是，把区间按照左端点升序排序。设当前区间为 $[\textit{start}_j,\textit{end}_j]$。由于我们按照左端点排序了，能与当前区间相交的下标小于 $j$ 的区间，当且仅当其右端点 $\ge \textit{start}_j$。于是问题变成：

- 有多少对 $(i,j)$ 满足 $i<j$ 且 $\textit{end}_i \ge \textit{start}_j$？

这个问题的补集更容易计算，也就是计算**不相交**的区间对：

- 有多少对 $(i,j)$ 满足 $i<j$ 且 $\textit{end}_i < \textit{start}_j$？

如果 $\textit{end}_i < \textit{start}_j$ 成立，那么有 $\textit{start}_i \le \textit{end}_i < \textit{start}_j$。由于我们已按照 $\textit{start}$ 升序排序，所以 $i<j$ 一定成立，是冗余的。问题进一步简化成：

- 有多少对 $(i,j)$ 满足 $\textit{end}_i < \textit{start}_j$？

**这个问题与下标无关**，我们可以把 $\textit{intervals}$ 拆成两个**独立**的数组 $\textit{starts}$ 和 $\textit{ends}$，分别保存所有左端点和所有右端点。然后，把两个数组都升序排序，就可以用**双指针**快速求出有多少对 $(i,j)$ 满足 $\textit{end}_i < \textit{start}_j$。

最后，用所有区间对的个数 $\dfrac{n(n-1)}{2}$ 减去不相交区间对的个数，即为答案。

```py [sol-Python3]
class Solution:
    def countIntersectingIntervals(self, intervals: list[list[int]]) -> int:
        n = len(intervals)
        starts = sorted(p[0] for p in intervals)
        ends = sorted(p[1] for p in intervals)

        ans = n * (n - 1) // 2
        # 对于每个左端点 start，右端点 < start 的区间都与之不相交
        j = 0
        for start in starts:
            while j < n and ends[j] < start:
                j += 1
            # [0, j-1] 的区间与当前区间不相交，这有 j 个
            ans -= j
        return ans
```

```java [sol-Java]
class Solution {
    public long countIntersectingIntervals(int[][] intervals) {
        int n = intervals.length;
        int[] starts = new int[n];
        int[] ends = new int[n];
        for (int i = 0; i < n; i++) {
            starts[i] = intervals[i][0];
            ends[i] = intervals[i][1];
        }

        Arrays.sort(starts);
        Arrays.sort(ends);

        long ans = (long) n * (n - 1) / 2;
        // 对于每个左端点 start，右端点 < start 的区间都与之不相交
        int j = 0;
        for (int start : starts) {
            while (j < n && ends[j] < start) {
                j++;
            }
            // [0, j-1] 的区间与当前区间不相交，这有 j 个
            ans -= j;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countIntersectingIntervals(vector<vector<int>>& intervals) {
        int n = intervals.size();
        vector<int> starts(n);
        vector<int> ends(n);
        for (int i = 0; i < n; i++) {
            starts[i] = intervals[i][0];
            ends[i] = intervals[i][1];
        }

        ranges::sort(starts);
        ranges::sort(ends);

        long long ans = 1LL * n * (n - 1) / 2;
        // 对于每个左端点 start，右端点 < start 的区间都与之不相交
        int j = 0;
        for (int start : starts) {
            while (j < n && ends[j] < start) {
                j++;
            }
            // [0, j-1] 的区间与当前区间不相交，这有 j 个
            ans -= j;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countIntersectingIntervals(intervals [][]int) int64 {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, p := range intervals {
		starts[i] = p[0]
		ends[i] = p[1]
	}

	slices.Sort(starts)
	slices.Sort(ends)

	ans := n * (n - 1) / 2
	// 对于每个左端点 start，右端点 < start 的区间都与之不相交
	j := 0
	for _, start := range starts {
		for j < n && ends[j] < start {
			j++
		}
		// [0, j-1] 的区间与当前区间不相交，这有 j 个
		ans -= j
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{intervals}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
