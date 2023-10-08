请看 [视频讲解](https://www.bilibili.com/video/BV1e84y117R9/) 第二题。

一颗处理器完成它的 $4$ 个任务，完成的时间取决于这 $4$ 个任务中的 $\textit{tasks}$ 的最大值。

直觉上来说，最早空闲时间越大的处理器，处理 $\textit{tasks}$ 越小的任务，那么完成时间越早。

证明：对于两个最早空闲时间分别为 $p_1$ 和 $p_2$ 的处理器，不妨设 $p_1 \le p_2$。完成的 $4$ 个任务中的最大值分别为 $t_1$ 和 $t_2$，不妨设 $t_1 \le t_2$。

如果 $t_1$ 给 $p_1$，$t_2$ 给 $p_2$，那么最后完成时间为

$$
\max(p_1+t_1, p_2+t_2) = p_2+t_2
$$

如果 $t_1$ 给 $p_2$，$t_2$ 给 $p_1$，那么最后完成时间为

$$
\max(p_1+t_2, p_2+t_1) \le \max(p_2+t_2, p_2+t_2) = p_2+t_2
$$

上式表明，最早空闲时间越大的处理器，处理 $\textit{tasks}$ 越小的任务，那么完成时间不会变的更晚。

我们可以把 $\textit{processorTime}$ 从小到大排序，$\textit{tasks}$ 从大到小排序，那么答案就是

$$
\textit{processorTime}[i] + \textit{tasks}[4i]
$$

的最大值。

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
        sort(processorTime.begin(), processorTime.end());
        sort(tasks.begin(), tasks.end(), greater<int>());
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
	sort.Ints(processorTime)
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for i, p := range processorTime {
		ans = max(ans, p+tasks[i*4])
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{processorTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。Python 忽略切片开销。
