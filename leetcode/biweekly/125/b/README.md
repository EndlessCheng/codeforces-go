注意题目说每次选择**最小**的两个数操作。那么用最小堆模拟即可。

```py [sol-Python3]
class Solution:
    def minOperations(self, h: List[int], k: int) -> int:
        ans = 0
        heapify(h)
        while h[0] < k:
            x = heappop(h)
            heapreplace(h, x * 2 + h[0])
            ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = 0;
        PriorityQueue<Long> pq = new PriorityQueue<>();
        for (int x : nums) {
            pq.offer((long) x);
        }
        while (pq.peek() < k) {
            long x = pq.poll();
            long y = pq.poll();
            pq.offer(x * 2 + y);
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, int k) {
        int ans = 0;
        priority_queue<long long, vector<long long>, greater<>> pq;
        for (int x : nums) {
            pq.push((long long) x);
        }
        while (pq.top() < k) {
            long long x = pq.top(); pq.pop();
            long long y = pq.top(); pq.pop();
            pq.push(x * 2 + y);
            ans++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) (ans int) {
	h := &hp{nums}
	heap.Init(h)
	for h.IntSlice[0] < k {
		x := heap.Pop(h).(int)
		h.IntSlice[0] += x * 2
		heap.Fix(h, 0)
		ans++
	}
	return
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)    {}
func (h *hp) Pop() any { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
