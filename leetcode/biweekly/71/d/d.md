将 $\textit{nums}$ 拆分成两部分，左半部分的最小和（前缀最小和）减右半部分的最大和（后缀最大和）即为两部分和的最小差值，枚举拆分位置（保证左右两部分至少有 $n$ 个元素），所有差值的最小值就是答案。

一个 $n=4$ 的例子：

![](https://pic.leetcode-cn.com/1644495549-IzYFpw-LC2163.drawio.png)

我们需要计算出 $\textit{nums}$ 的前缀最小和 $\textit{preMin}[i]$，即前 $i$ 个元素中的最小的 $n$ 个元素之和；以及后缀最大和 $\textit{sufMax}[i]$，即后 $i$ 个元素中的最大的 $n$ 个元素之和。答案即为 $\textit{preMin}[i]-\textit{sufMax}[i+1]$ 中的最小值。

计算前缀最小和时，可以维护一个包含 $n$ 个元素的最大堆，我们不断向右遍历 $\textit{nums}$ 中的元素 $v$，计算前缀最小和，若 $v$ 比堆顶元素小，则弹出堆顶元素，并将 $v$ 入堆。

计算后缀最大和，则需要维护一个包含 $n$ 个元素的最小堆，我们不断向左遍历 $\textit{nums}$ 中的元素 $v$，计算后缀最大和，若 $v$ 比堆顶元素大，则弹出堆顶元素，并将 $v$ 入堆。

代码实现时，可以先计算出后缀最大和，然后在计算前缀最小和的同时计算出答案。

```go [sol1-Go]
func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3
	minPQ := minHeap{nums[m-n:]}
	heap.Init(&minPQ)
	sum := 0
	for _, v := range nums[m-n:] {
		sum += v
	}
	sufMax := make([]int, m-n+1) // 后缀最大和
	sufMax[m-n] = sum
	for i := m - n - 1; i >= n; i-- {
		if v := nums[i]; v > minPQ.IntSlice[0] {
			sum += v - minPQ.IntSlice[0]
			minPQ.IntSlice[0] = v
			heap.Fix(&minPQ, 0)
		}
		sufMax[i] = sum
	}

	maxPQ := maxHeap{nums[:n]}
	heap.Init(&maxPQ)
	preMin := 0 // 前缀最小和
	for _, v := range nums[:n] {
		preMin += v
	}
	ans := preMin - sufMax[n]
	for i := n; i < m-n; i++ {
		if v := nums[i]; v < maxPQ.IntSlice[0] {
			preMin += v - maxPQ.IntSlice[0]
			maxPQ.IntSlice[0] = v
			heap.Fix(&maxPQ, 0)
		}
		ans = min(ans, preMin-sufMax[i+1])
	}
	return int64(ans)
}

type minHeap struct{ sort.IntSlice }
func (minHeap) Push(interface{})     {}
func (minHeap) Pop() (_ interface{}) { return }

type maxHeap struct{ sort.IntSlice }
func (h maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (maxHeap) Push(interface{})     {}
func (maxHeap) Pop() (_ interface{}) { return }

func min(a, b int) int { if a > b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    long long minimumDifference(vector<int> &nums) {
        int m = nums.size(), n = m / 3;
        priority_queue<int, vector<int>, greater<int>> minPQ;
        long sum = 0L;
        for (int i = m - n; i < m; ++i) {
            minPQ.push(nums[i]);
            sum += nums[i];
        }
        vector<long> sufMax(m - n + 1); // 后缀最大和
        sufMax[m - n] = sum;
        for (int i = m - n - 1; i >= n; --i) {
            minPQ.push(nums[i]);
            sum += nums[i] - minPQ.top();
            minPQ.pop();
            sufMax[i] = sum;
        }

        priority_queue<int> maxPQ;
        long preMin = 0L; // 前缀最小和
        for (int i = 0; i < n; ++i) {
            maxPQ.push(nums[i]);
            preMin += nums[i];
        }
        long ans = preMin - sufMax[n];
        for (int i = n; i < m - n; ++i) {
            maxPQ.push(nums[i]);
            preMin += nums[i] - maxPQ.top();
            maxPQ.pop();
            ans = min(ans, preMin - sufMax[i + 1]);
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def minimumDifference(self, nums: List[int]) -> int:
        m = len(nums)
        n = m // 3

        min_pq = nums[m - n:]
        heapify(min_pq)
        suf_max = [0] * (m - n + 1)  # 后缀最大和
        suf_max[-1] = s = sum(min_pq)
        for i in range(m - n - 1, n - 1, -1):
            s += nums[i] - heappushpop(min_pq, nums[i])
            suf_max[i] = s

        max_pq = [-v for v in nums[:n]]  # 所有元素取反当最大堆
        heapify(max_pq)
        pre_min = -sum(max_pq)  # 前缀最小和
        ans = pre_min - suf_max[n]
        for i in range(n, m - n):
            pre_min += nums[i] + heappushpop(max_pq, -nums[i])
            ans = min(ans, pre_min - suf_max[i + 1])
        return ans
```

```java [sol1-Java]
class Solution {
    public long minimumDifference(int[] nums) {
        var m = nums.length;
        var n = m / 3;
        var minPQ = new PriorityQueue<Integer>();
        var sum = 0L;
        for (var i = m - n; i < m; i++) {
            minPQ.add(nums[i]);
            sum += nums[i];
        }
        var sufMax = new long[m - n + 1]; // 后缀最大和
        sufMax[m - n] = sum;
        for (var i = m - n - 1; i >= n; --i) {
            minPQ.add(nums[i]);
            sum += nums[i] - minPQ.poll();
            sufMax[i] = sum;
        }

        var maxPQ = new PriorityQueue<Integer>(Collections.reverseOrder());
        var preMin = 0L; // 前缀最小和
        for (var i = 0; i < n; ++i) {
            maxPQ.add(nums[i]);
            preMin += nums[i];
        }
        var ans = preMin - sufMax[n];
        for (var i = n; i < m - n; ++i) {
            maxPQ.add(nums[i]);
            preMin += nums[i] - maxPQ.poll();
            ans = Math.min(ans, preMin - sufMax[i + 1]);
        }
        return ans;
    }
}
```

