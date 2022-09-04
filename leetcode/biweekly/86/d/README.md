下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

前置题目：[239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)。

本题的一种做法是二分答案，这样就转换成了 239 题。

但实际上不用二分，在 239 这题的基础上，把固定大小的滑动窗口改为不固定大小的双指针，具体见代码注释。

更多有关单调队列的题目见我的算法竞赛模板库中的 [monotone_queue.go](https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/monotone_queue.go)。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{chargeTimes}$ 的长度。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def maximumRobots(self, chargeTimes: List[int], runningCosts: List[int], budget: int) -> int:
        ans = s = left = 0
        q = deque()
        # 枚举区间右端点 right，计算区间左端点 left 的最小值
        for right, (t, c) in enumerate(zip(chargeTimes, runningCosts)):
            # 及时清除队列中的无用数据，保证队列的单调性
            while q and t >= chargeTimes[q[-1]]:
                q.pop()
            q.append(right)
            s += c
            # 如果左端点 left 不满足要求，就不断右移 left
            while q and chargeTimes[q[0]] + (right - left + 1) * s > budget:
                # 及时清除队列中的无用数据，保证队列的单调性
                if q[0] == left:
                    q.popleft()
                s -= runningCosts[left]
                left += 1
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol1-Java]
class Solution {
    public int maximumRobots(int[] chargeTimes, int[] runningCosts, long budget) {
        var ans = 0;
        var q = new ArrayDeque<Integer>();
        var sum = 0L;
        // 枚举区间右端点 right，计算区间左端点 left 的最小值
        for (int left = 0, right = 0; right < chargeTimes.length; ++right) {
            // 及时清除队列中的无用数据，保证队列的单调性
            while (!q.isEmpty() && chargeTimes[right] >= chargeTimes[q.peekLast()])
                q.pollLast();
            q.addLast(right);
            sum += runningCosts[right];
            // 如果左端点 left 不满足要求，就不断右移 left
            while (!q.isEmpty() && chargeTimes[q.peekFirst()] + (right - left + 1) * sum > budget) {
                // 及时清除队列中的无用数据，保证队列的单调性
                if (q.peekFirst() == left) q.pollFirst();
                sum -= runningCosts[left++];
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maximumRobots(vector<int> &chargeTimes, vector<int> &runningCosts, long long budget) {
        int ans = 0;
        deque<int> q;
        long sum = 0L;
        // 枚举区间右端点 right，计算区间左端点 left 的最小值
        for (int left = 0, right = 0; right < chargeTimes.size(); ++right) {
            // 及时清除队列中的无用数据，保证队列的单调性
            while (!q.empty() && chargeTimes[right] >= chargeTimes[q.back()])
                q.pop_back();
            q.push_back(right);
            sum += runningCosts[right];
            // 如果左端点 left 不满足要求，就不断右移 left
            while (!q.empty() && chargeTimes[q.front()] + (right - left + 1) * sum > budget) {
                // 及时清除队列中的无用数据，保证队列的单调性
                if (q.front() == left) q.pop_front();
                sum -= runningCosts[left++];
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumRobots(chargeTimes, runningCosts []int, budget int64) (ans int) {
	sum, left, q := int64(0), 0, []int{}
	// 枚举区间右端点 right，计算区间左端点 left 的最小值
	for right, t := range chargeTimes {
		// 及时清除队列中的无用数据，保证队列的单调性
		for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, right)
		sum += int64(runningCosts[right])
		// 如果左端点 left 不满足要求，就不断右移 left
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
			// 及时清除队列中的无用数据，保证队列的单调性
			if q[0] == left {
				q = q[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 思考题

把「子数组」改成「子序列」要怎么做？
