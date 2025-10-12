根据题意：

- $\textit{time}$ 参数是严格递增的，我们可以用一个 $\textit{times}$ 数组记录 $\textit{time}$。这样 $\textit{times}$ 是严格递增数组，可以在上面**二分查找**，原理请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。
- $\texttt{totalScore}$ 本质上计算的是关于 $\textit{score}$ 的子数组和，我们可以维护 $\textit{score}$ 的**前缀和**数组 $\textit{preSum}$，从而 $\mathcal{O}(1)$ 计算子数组和。原理请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

二分查找 $\textit{times}$ 在 $[\textit{startTime},\textit{endTime}]$ 中的元素的下标范围 $[\textit{left},\textit{right}-1]$。

其中：

- $\textit{left}$ 是 $\textit{times}$ 第一个 $\ge \textit{startTime}$ 的元素下标。
- $\textit{right}-1$ 是 $\textit{times}$ 最后一个 $\le \textit{startTime}$ 的元素下标。
- $\textit{right}$ 是 $\textit{times}$ 第一个 $> \textit{startTime}$ 的元素下标，也是第一个 $\ge \textit{startTime}+1$ 的元素下标。

知道了下标范围，根据 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 中的公式，对应的子数组和为

$$
\textit{preSum}[\textit{right}] - \textit{preSum}[\textit{left}]
$$

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class ExamTracker:
    def __init__(self):
        self.times = []
        self.pre_sum = [0]  # 原理见题目 303. 区域和检索 - 数组不可变

    def record(self, time: int, score: int) -> None:
        self.times.append(time)
        self.pre_sum.append(self.pre_sum[-1] + score)

    def totalScore(self, startTime: int, endTime: int) -> int:
        left = bisect_left(self.times, startTime)
        right = bisect_right(self.times, endTime)
        return self.pre_sum[right] - self.pre_sum[left]
```

```java [sol-Java]
class ExamTracker {
    private final List<Integer> times = new ArrayList<>();
    private final List<Long> preSum = new ArrayList<>();

    public ExamTracker() {
        preSum.add(0L); // 原理见题目 303. 区域和检索 - 数组不可变
    }

    public void record(int time, int score) {
        times.add(time);
        preSum.add(preSum.getLast() + score);
    }

    public long totalScore(int startTime, int endTime) {
        // times 没有重复元素，可以用库函数二分（有重复元素则不行）
        int left = Collections.binarySearch(times, startTime);
        if (left < 0) left = ~left;
        int right = Collections.binarySearch(times, endTime + 1);
        if (right < 0) right = ~right;
        return preSum.get(right) - preSum.get(left);
    }
}
```

```cpp [sol-C++]
class ExamTracker {
    vector<int> times;
    vector<long long> pre_sum = {0}; // 原理见题目 303. 区域和检索 - 数组不可变

public:
    void record(int time, int score) {
        times.push_back(time);
        pre_sum.push_back(pre_sum.back() + score);
    }

    long long totalScore(int startTime, int endTime) {
        int left = ranges::lower_bound(times, startTime) - times.begin();
        int right = ranges::upper_bound(times, endTime) - times.begin();
        return pre_sum[right] - pre_sum[left];
    }
};
```

```go [sol-Go]
type ExamTracker struct {
	times  []int
	preSum []int64
}

func Constructor() ExamTracker {
	// preSum 为什么加个 0，见题目 303. 区域和检索 - 数组不可变
	return ExamTracker{[]int{}, []int64{0}}
}

func (e *ExamTracker) Record(time, score int) {
	e.times = append(e.times, time)
	e.preSum = append(e.preSum, e.preSum[len(e.preSum)-1]+int64(score))
}

func (e *ExamTracker) TotalScore(startTime, endTime int) int64 {
	left := sort.SearchInts(e.times, startTime)
	right := sort.SearchInts(e.times, endTime+1) // 也可以在 e.times[left:] 中二分
	return e.preSum[right] - e.preSum[left]
}
```

#### 复杂度分析

- 时间复杂度：
  - 初始化：$\mathcal{O}(1)$。
  - $\texttt{record}$：均摊 $\mathcal{O}(1)$。
  - $\texttt{totalScore}$：$\mathcal{O}(\log q)$，其中 $q$ 是 $\texttt{record}$ 的调用次数。
- 空间复杂度：$\mathcal{O}(q)$。

## 专题训练

1. 下面二分题单的「**一、二分查找**」。
2. 下面数据结构题单的「**一、前缀和**」。

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
