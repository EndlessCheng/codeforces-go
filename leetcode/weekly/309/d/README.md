下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

用两个小顶堆模拟：

- $\textit{idle}$ 维护在 $\textit{start}_i$ 时刻空闲的会议室；
- $\textit{using}$ 维护在 $\textit{start}_i$ 时刻使用中的会议室。

对 $\textit{meetings}$ 按照开始时间排序，然后遍历 $\textit{meetings}$，按照题目要求模拟即可，具体模拟方式见代码。

#### 复杂度分析

- 时间复杂度：$O(n+m(\log n + \log m))$，其中 $m$ 为 $\textit{meetings}$ 的长度。注意每个会议至多入堆出堆各一次。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        cnt = [0] * n
        idle, using = list(range(n)), []
        meetings.sort(key=lambda m: m[0])
        for st, end in meetings:
            while using and using[0][0] <= st:
                heappush(idle, heappop(using)[1])  # 维护在 st 时刻空闲的会议室
            if len(idle) == 0:
                e, i = heappop(using)  # 没有可用的会议室，那么弹出一个最早结束的会议室
                end += e - st  # 更新当前会议的结束时间
            else:
                i = heappop(idle)
            cnt[i] += 1
            heappush(using, (end, i))  # 使用一个会议室
        ans = 0
        for i, c in enumerate(cnt):
            if c > cnt[ans]:
                ans = i
        return ans
```

```java [sol1-Java]
class Solution {
    public int mostBooked(int n, int[][] meetings) {
        var cnt = new int[n];
        var idle = new PriorityQueue<Integer>((a, b) -> Integer.compare(a, b));
        for (var i = 0; i < n; ++i) idle.offer(i);
        var using = new PriorityQueue<Pair<Long, Integer>>((a, b) -> !Objects.equals(a.getKey(), b.getKey()) ? Long.compare(a.getKey(), b.getKey()) : Integer.compare(a.getValue(), b.getValue()));
        Arrays.sort(meetings, (a, b) -> Integer.compare(a[0], b[0]));
        for (var m : meetings) {
            long st = m[0], end = m[1];
            while (!using.isEmpty() && using.peek().getKey() <= st) {
                idle.offer(using.poll().getValue()); // 维护在 st 时刻空闲的会议室
            }
            int id;
            if (idle.isEmpty()) {
                var p = using.poll(); // 没有可用的会议室，那么弹出一个最早结束的会议室
                end += p.getKey() - st; // 更新当前会议的结束时间
                id = p.getValue();
            } else id = idle.poll();
            ++cnt[id];
            using.offer(new Pair<>(end, id)); // 使用一个会议室
        }
        var ans = 0;
        for (var i = 0; i < n; ++i) if (cnt[i] > cnt[ans]) ans = i;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int mostBooked(int n, vector<vector<int>> &meetings) {
        int cnt[n]; memset(cnt, 0, sizeof(cnt));
        priority_queue<int, vector<int>, greater<>> idle;
        for (int i = 0; i < n; ++i) idle.push(i);
        priority_queue<pair<long, int>, vector<pair<long, int>>, greater<>> using_;
        sort(meetings.begin(), meetings.end(), [](auto &a, auto &b) { return a[0] < b[0]; });
        for (auto &m : meetings) {
            long st = m[0], end = m[1], id;
            while (!using_.empty() && using_.top().first <= st) {
                idle.push(using_.top().second); // 维护在 st 时刻空闲的会议室
                using_.pop();
            }
            if (idle.empty()) {
                auto[e, i] = using_.top(); // 没有可用的会议室，那么弹出一个最早结束的会议室
                using_.pop();
                end += e - st; // 更新当前会议的结束时间
                id = i;
            } else {
                id = idle.top();
                idle.pop();
            }
            ++cnt[id];
            using_.emplace(end, id); // 使用一个会议室
        }
        int ans = 0;
        for (int i = 0; i < n; ++i) if (cnt[i] > cnt[ans]) ans = i;
        return ans;
    }
};
```

```go [sol1-Go]
func mostBooked(n int, meetings [][]int) (ans int) {
	cnt := make([]int, n)
	idle := hp{make([]int, n)}
	for i := 0; i < n; i++ {
		idle.IntSlice[i] = i
	}
	using := hp2{}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	for _, m := range meetings {
		st, end := m[0], m[1]
		for len(using) > 0 && using[0].end <= st {
			heap.Push(&idle, heap.Pop(&using).(pair).i) // 维护在 st 时刻空闲的会议室
		}
		var i int
		if idle.Len() == 0 {
			p := heap.Pop(&using).(pair) // 没有可用的会议室，那么弹出一个最早结束的会议室
			end += p.end - st // 更新当前会议的结束时间
			i = p.i
		} else {
			i = heap.Pop(&idle).(int)
		}
		cnt[i]++
		heap.Push(&using, pair{end, i}) // 使用一个会议室
	}
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type pair struct{ end, i int }
type hp2 []pair
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { a, b := h[i], h[j]; return a.end < b.end || a.end == b.end && a.i < b.i }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```
