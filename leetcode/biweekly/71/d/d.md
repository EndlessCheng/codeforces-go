将 $\textit{nums}$ 拆分成两部分，左半部分的最小和（前缀最小和）减右半部分的最大和（后缀最大和）即为两部分和的最小差值，枚举拆分位置（保证左右两部分至少有 $n$ 个元素），所有差值的最小值就是答案。

一个 $n=4$ 的例子：

![](https://pic.leetcode-cn.com/1644495549-IzYFpw-LC2163.drawio.png)

我们需要计算出 $\textit{nums}$ 的前缀最小和 $\textit{preMin}[i]$，即前 $i$ 个元素中的最小的 $n$ 个元素之和；以及后缀最大和 $\textit{sufMax}[i]$，即后 $i$ 个元素中的最大的 $n$ 个元素之和。答案即为 $\textit{preMin}[i]-\textit{sufMax}[i+1]$ 中的最小值。

计算前缀最小和时，可以维护一个包含 $n$ 个元素的最大堆，我们不断向右遍历 $\textit{nums}$ 中的元素 $v$，计算前缀最小和，若 $v$ 比堆顶元素小，则弹出堆顶元素，并将 $v$ 入堆。

计算后缀最大和，则需要维护一个包含 $n$ 个元素的最小堆，我们不断向左遍历 $\textit{nums}$ 中的元素 $v$，计算后缀最大和，若 $v$ 比堆顶元素大，则弹出堆顶元素，并将 $v$ 入堆。

代码实现时，可以先计算出后缀最大和，然后在计算前缀最小和的同时计算出答案。

```py [sol-Python3]
class Solution:
    def minimumDifference(self, nums: List[int]) -> int:
        m = len(nums)
        n = m // 3
        min_pq = nums[-n:]
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

```java [sol-Java]
class Solution {
    public long minimumDifference(int[] nums) {
        int m = nums.length;
        int n = m / 3;
        PriorityQueue<Integer> minPQ = new PriorityQueue<>();
        long sum = 0;
        for (int i = m - n; i < m; i++) {
            minPQ.add(nums[i]);
            sum += nums[i];
        }

        long[] sufMax = new long[m - n + 1]; // 后缀最大和
        sufMax[m - n] = sum;
        for (int i = m - n - 1; i >= n; i--) {
            minPQ.add(nums[i]);
            sum += nums[i] - minPQ.poll();
            sufMax[i] = sum;
        }

        PriorityQueue<Integer> maxPQ = new PriorityQueue<>(Collections.reverseOrder());
        long preMin = 0; // 前缀最小和
        for (int i = 0; i < n; ++i) {
            maxPQ.add(nums[i]);
            preMin += nums[i];
        }

        long ans = preMin - sufMax[n];
        for (int i = n; i < m - n; i++) {
            maxPQ.add(nums[i]);
            preMin += nums[i] - maxPQ.poll();
            ans = Math.min(ans, preMin - sufMax[i + 1]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumDifference(vector<int> &nums) {
        int m = nums.size(), n = m / 3;
        priority_queue<int, vector<int>, greater<>> minPQ;
        long long sum = 0;
        for (int i = m - n; i < m; i++) {
            minPQ.push(nums[i]);
            sum += nums[i];
        }

        vector<long long> sufMax(m - n + 1); // 后缀最大和
        sufMax[m - n] = sum;
        for (int i = m - n - 1; i >= n; i--) {
            minPQ.push(nums[i]);
            sum += nums[i] - minPQ.top();
            minPQ.pop();
            sufMax[i] = sum;
        }

        priority_queue<int> maxPQ;
        long long preMin = 0; // 前缀最小和
        for (int i = 0; i < n; i++) {
            maxPQ.push(nums[i]);
            preMin += nums[i];
        }

        long long ans = preMin - sufMax[n];
        for (int i = n; i < m - n; i++) {
            maxPQ.push(nums[i]);
            preMin += nums[i] - maxPQ.top();
            maxPQ.pop();
            ans = min(ans, preMin - sufMax[i + 1]);
        }
        return ans;
    }
};
```

```go [sol-Go]
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
func (minHeap) Push(any) {}
func (minHeap) Pop() (_ any) { return }

type maxHeap struct{ sort.IntSlice }
func (h maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (maxHeap) Push(any) {}
func (maxHeap) Pop() (_ any) { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
