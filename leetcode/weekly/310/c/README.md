本题 [视频讲解](https://www.bilibili.com/video/BV1it4y1L7kL) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---
 
按照 $\textit{left}$ 排序后，用最小堆模拟，堆顶存储每个组最后一个区间的 $\textit{right}$。

遍历区间：

- 如果当前的 $\textit{left}$ 大于堆顶，则可以接在这个组的末尾，更新堆顶为 $\textit{right}$；
- 否则需要创建一个新的组。

```py [sol1-Python3]
class Solution:
    def minGroups(self, intervals: List[List[int]]) -> int:
        intervals.sort()
        h = []
        for left, right in intervals:
            if h and left > h[0]: heapreplace(h, right)
            else: heappush(h, right)
        return len(h)
```

```go [sol1-Go]
func minGroups(intervals [][]int) int {
	h := hp{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, p := range intervals {
		if h.Len() == 0 || p[0] <= h.IntSlice[0] {
			heap.Push(&h, p[1])
		} else {
			h.IntSlice[0] = p[1]
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

另外一种思路是把区间内的数都加一，最后答案为所有数的最大值。

这可以用差分数组实现，下面代码用的平衡树，方便从小到大计算。

```cpp [sol1-C++]
class Solution {
public:
    int minGroups(vector<vector<int>> &intervals) {
        map<int, int> diff;
        for (auto &p : intervals)
            ++diff[p[0]], --diff[p[1] + 1];
        int ans = 0, sum = 0;
        for (auto &[_, d] : diff)
            ans = max(ans, sum += d);
        return ans;
    }
};
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

#### 思考题

从所有满足 $1\le\textit{left}\le\textit{right}\le m$ 的一共 $\dfrac{m(m+1)}{2}$ 个区间中，随机选择 $n$ 个区间作为本题的输入，得到的答案的期望是多少？