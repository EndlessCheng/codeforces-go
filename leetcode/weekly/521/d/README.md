本题是 [1235. 规划兼职工作](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/) 的变形题，请先完成那题。

下面接着 [我的题解](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/solutions/1913089/dong-tai-gui-hua-er-fen-cha-zhao-you-hua-zkcg/) 继续讲。

类似 1235 题，定义 $f[i]$ 表示最后一场会议是 $\textit{meetings}[i]$ 时，能得到的最大收益。

**注**：本题需要知道倒数第二场会议的位置，从而计算空闲时间的收益。所以本题是**相邻相关 DP**，故定义成 $\textit{meetings}[i]$ **一定要选**，从而方便转移。顺带一提，相邻相关 DP 的代表是 [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)。

枚举倒数第二场会议为 $\textit{meetings}[j]$，满足 $\textit{end}_j \le \textit{start}_i$。问题变成最后一场会议是 $\textit{meetings}[j]$ 时，能得到的最大收益，即 $f[j]$。加上 $j$ 和 $i$ 之间的空闲时间 $\textit{start}_i - \textit{end}_j$，以及会议 $i$ 的收益 $\textit{revenue}_i$，得到状态转移方程

$$
f[i] = \max_{\textit{end}_j \le \textit{start}_i} f[j] + \textit{start}_i - \textit{end}_j + \textit{revenue}_i
$$

提出与 $j$ 无关的项，上式变形为

$$
f[i] = \textit{start}_i + \textit{revenue}_i + \max_{\textit{end}_j \le \textit{start}_i} f[j]  - \textit{end}_j
$$

如果把会议按照右端点从小到大排序，那么上式要计算的是 $f[j]  - \textit{end}_j$ 的**前缀最大值**。

于是定义 

$$
\textit{preMax}[i+1] = \max_{j=0}^{i} f[j]  - \textit{end}_j
$$

设 $k$ 是最大的满足 $\textit{end}_j \le \textit{start}_i$ 的下标 $j$，那么状态转移方程为

$$
f[i] = \textit{preMax}[k+1] + \textit{start}_i + \textit{revenue}_i
$$

$k$ 可以在 $\textit{meetings}$ 数组上 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 求出。

初始值：如果 $\textit{meetings}[i]$ 左边没有会议，那么 $f[i] = \textit{revenue}_i$。

答案：$\max(f)$。

[本题视频讲解](https://www.bilibili.com/video/BV12gah6UE9b/?t=26m2s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxEarnings(self, meetings: list[list[int]]) -> int:
        # 按结束时间从小到大排序
        meetings.sort(key=lambda m: m[1])
        end0 = meetings[0][1]

        # pre_max[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
        pre_max = [-inf] * (len(meetings) + 1)
        ans = 0
        for i, (start, end, revenue) in enumerate(meetings):
            f = revenue
            if start >= end0:  # 左边有会议
                j = bisect_right(meetings, start, hi=i, key=lambda m: m[1])  # hi=i 表示二分上界为 i（默认为 n）
                # 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
                f += pre_max[j] + start
            ans = max(ans, f)
            pre_max[i + 1] = max(pre_max[i], f - end)

        return ans
```

```java [sol-Java]
class Solution {
    public long maxEarnings(int[][] meetings) {
        // 按结束时间从小到大排序
        Arrays.sort(meetings, (a, b) -> a[1] - b[1]);
        int end0 = meetings[0][1];

        // preMax[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
        int n = meetings.length;
        long[] preMax = new long[n + 1];
        preMax[0] = Long.MIN_VALUE;
        long ans = 0;

        for (int i = 0; i < n; i++) {
            int[] m = meetings[i];
            int start = m[0], end = m[1], revenue = m[2];

            long f = revenue;
            if (start >= end0) { // 左边有会议
                int j = search(meetings, i, start);
                f += preMax[j + 1] + start;
            }
            ans = Math.max(ans, f);

            preMax[i + 1] = Math.max(preMax[i], f - end);
        }

        return ans;
    }

    // 返回满足 end[j] <= upper 的最大 j
    private int search(int[][] meetings, int right, int upper) {
        int left = -1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (meetings[mid][1] <= upper) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxEarnings(vector<vector<int>>& meetings) {
        // 按结束时间从小到大排序
        ranges::sort(meetings, {}, [](auto& m) { return m[1]; });
        int end0 = meetings[0][1];

        int n = meetings.size();
        // pre_max[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
        vector<long long> pre_max(n + 1);
        pre_max[0] = LLONG_MIN;
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            auto& m = meetings[i];
            int start = m[0], end = m[1], revenue = m[2];

            long long f = revenue;
            if (start >= end0) { // 左边有会议
                int j = ranges::upper_bound(meetings, start, {}, [](auto& m) { return m[1]; }) - meetings.begin();
                // 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
                f += pre_max[j] + start;
            }
            ans = max(ans, f);

            pre_max[i + 1] = max(pre_max[i], f - end);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxEarnings(meetings [][]int) int64 {
	// 按结束时间从小到大排序
	slices.SortFunc(meetings, func(a, b []int) int { return a[1] - b[1] })
	end0 := meetings[0][1]

	// preMax[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
	preMax := make([]int, len(meetings)+1)
	preMax[0] = math.MinInt
	ans := 0
	for i, m := range meetings {
		start, end, revenue := m[0], m[1], m[2]

		f := revenue
		if start >= end0 { // 左边有会议
			j := sort.Search(i, func(j int) bool { return meetings[j][1] > start })
			// 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
			f += preMax[j] + start
		}
		ans = max(ans, f)

		preMax[i+1] = max(preMax[i], f-end)
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{meetings}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§7.2 不相交区间**」。

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
