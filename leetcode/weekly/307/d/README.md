下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

记 $\textit{nums}$ 中所有非负数的和为 $\textit{sum}$。

任意一个子序列的和，都等价于从 $\textit{sum}$ 中减去某些非负数，或者加上某些负数得到。

#### 提示 2

将 $\textit{nums}$ 所有数取绝对值。然后按照从小到大的顺序取出 $\textit{sum}$ 需要减去的子序列，如何做到？

#### 提示 3-1

将 $\textit{nums}$ 所有数取绝对值后排序，然后用最大堆来实现。

#### 提示 3-2

具体来说，最大堆维护子序列的和，以及（后续需要减去的）数字的下标 $i$。

初始时，将 $\textit{sum}$ 和下标 $0$ 入堆。

每次弹出堆顶时，我们选择减去 $\textit{nums}[i]$，并考虑是否保留 $\textit{nums}[i-1]$，从而满足子序列每个元素「选或不选」的要求。

循环 $k-1$ 次后，堆顶的和就是答案。

#### 复杂度分析

- 时间复杂度：$O(n\log n + k\log k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(k)$。

```py [sol1-Python3]
class Solution:
    def kSum(self, nums: List[int], k: int) -> int:
        sum = 0
        for i, x in enumerate(nums):
            if x >= 0:
                sum += x
            else:
                nums[i] = -x
        nums.sort()
        h = [(-sum, 0)]  # 取负号变成最大堆
        while k > 1:
            k -= 1
            s, i = heappop(h)
            if i < len(nums):
                heappush(h, (s + nums[i], i + 1))  # 保留 nums[i-1]
                if i:
                    heappush(h, (s + nums[i] - nums[i - 1], i + 1))  # 不保留 nums[i-1]
        return -h[0][0]
```

```java [sol1-Java]
class Solution {
    public long kSum(int[] nums, int k) {
        var sum = 0L;
        for (var i = 0; i < nums.length; i++)
            if (nums[i] >= 0) sum += nums[i];
            else nums[i] = -nums[i];
        Arrays.sort(nums);
        var pq = new PriorityQueue<Pair<Long, Integer>>((a, b) -> Long.compare(b.getKey(), a.getKey()));
        pq.offer(new Pair<>(sum, 0));
        while (--k > 0) {
            var p = pq.poll();
            var s = p.getKey();
            var i = p.getValue();
            if (i < nums.length) {
                pq.offer(new Pair<>(s - nums[i], i + 1)); // 保留 nums[i-1]
                if (i > 0) pq.offer(new Pair<>(s - nums[i] + nums[i - 1], i + 1)); // 不保留 nums[i-1]
            }
        }
        return pq.peek().getKey();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long kSum(vector<int> &nums, int k) {
        long sum = 0L;
        for (int &x : nums)
            if (x >= 0) sum += x;
            else x = -x;
        sort(nums.begin(), nums.end());
        priority_queue<pair<long, int>> pq;
        pq.emplace(sum, 0);
        while (--k) {
            auto[sum, i] = pq.top();
            pq.pop();
            if (i < nums.size()) {
                pq.emplace(sum - nums[i], i + 1); // 保留 nums[i-1]
                if (i) pq.emplace(sum - nums[i] + nums[i - 1], i + 1); // 不保留 nums[i-1]
            }
        }
        return pq.top().first;
    }
};
```

```go [sol1-Go]
func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	sort.Ints(nums)
	h := &hp{{sum, 0}}
	for ; k > 1; k-- {
		p := heap.Pop(h).(pair)
		if p.i < n {
			heap.Push(h, pair{p.sum - nums[p.i], p.i + 1}) // 保留 nums[p.i-1]
			if p.i > 0 {
				heap.Push(h, pair{p.sum - nums[p.i] + nums[p.i-1], p.i + 1}) // 不保留 nums[p.i-1]
			}
		}
	}
	return int64((*h)[0].sum)
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].sum > h[j].sum }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```
