一颗处理器完成它的 $4$ 个任务，完成的时间取决于这 $4$ 个任务中的 $\textit{tasks}$ 的最大值。

直觉上来说，处理器的空闲时间越早，应当处理 $\textit{tasks}$ 越大的任务；处理器的空闲时间越晚，应当处理 $\textit{tasks}$ 越小的任务。

采用**交换论证法**（exchange argument）证明。

对于两个最早空闲时间分别为 $p_1$ 和 $p_2$ 的处理器，不妨设 $p_1 \le p_2$。完成的 $4$ 个任务中的最大值分别为 $t_1$ 和 $t_2$，不妨设 $t_1 \le t_2$。

如果 $t_1$ 给 $p_1$，$t_2$ 给 $p_2$，那么最后完成时间为

$$
\max(p_1+t_1, p_2+t_2) = p_2+t_2
$$

如果 $t_1$ 给 $p_2$，$t_2$ 给 $p_1$，那么最后完成时间为

$$
\max(p_1+t_2, p_2+t_1) \le \max(p_2+t_2, p_2+t_2) = p_2+t_2
$$

上式表明，处理器的空闲时间越早，应当处理 $\textit{tasks}$ 越大的任务；处理器的空闲时间越晚，应当处理 $\textit{tasks}$ 越小的任务。

我们可以把 $\textit{processorTime}$ 从小到大排序，$\textit{tasks}$ 从大到小排序，那么答案就是

$$
\textit{processorTime}[i] + \textit{tasks}[4i]
$$

的最大值。

请看 [视频讲解](https://www.bilibili.com/video/BV1e84y117R9/) 第二题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minProcessingTime(self, processorTime: List[int], tasks: List[int]) -> int:
        processorTime.sort()
        tasks.sort(reverse=True)
        return max(p + t for p, t in zip(processorTime, tasks[::4]))
```

```java [sol-Java]
class Solution {
    public int minProcessingTime(List<Integer> processorTime, List<Integer> tasks) {
        Collections.sort(processorTime);
        tasks.sort(Collections.reverseOrder());
        int ans = 0;
        for (int i = 0; i < processorTime.size(); i++) {
            ans = Math.max(ans, processorTime.get(i) + tasks.get(i * 4));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minProcessingTime(vector<int>& processorTime, vector<int>& tasks) {
        ranges::sort(processorTime);
        ranges::sort(tasks, greater());
        int ans = 0;
        for (int i = 0; i < processorTime.size(); i++) {
            ans = max(ans, processorTime[i] + tasks[i * 4]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minProcessingTime(processorTime, tasks []int) (ans int) {
	slices.Sort(processorTime)
	slices.SortFunc(tasks, func(a, b int) int { return b - a })
	for i, p := range processorTime {
		ans = max(ans, p+tasks[i*4])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{processorTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。Python 忽略切片开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
