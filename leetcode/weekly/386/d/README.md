[视频讲解](https://www.bilibili.com/video/BV1qx421179t/) 第四题。

## 更形象的题意

题意有点抽象，形象地解释一下：

你有 $n$ 门课程需要考试，第 $i$ 门课程需要用 $\textit{nums}[i]$ 天复习。同一天只能复习一门课程（**慢速**复习）。

在第 $i$ 天，你可以**快速**搞定第 $\textit{changeIndices}[i]$ 门课程的复习。

你可以在任意一天完成一门课程的考试（前提是复习完成）。考试这一天不能复习。

搞定所有课程的复习+考试，至少要多少天？

## 提示 1

答案越大，越能够搞定所有课程，反之越不能。

有单调性，可以**二分答案**。

## 提示 2

如果决定在第 $i$ 天快速复习第 $\textit{changeIndices}[i]$ 门课程，那么在第 $i$ 天前慢速复习这门课程是没有意义的。同理，如果决定慢速复习某门课程，那么也没必要对这门课程使用快速复习。

在 $\textit{nums}[i] > 1$ 的情况下，如果用快速复习+考试，只需要花费 $2$ 天。这比只用慢速复习+考试要更快。

但是！如果一门课程对应的 $\textit{changeIndices}$ 很靠后，可能没有时间快速复习这门课程并完成考试。比如只剩下 $2$ 天，但是还有 $3$ 门课程没有考。这样的课程用慢速复习更合适。

此外，如果一门课程的复习时间很长（$\textit{nums}[i]$ 很大），当我们把后续时间都用在快速复习其它复习时间比较小的课程上，可能就没有时间快速复习 $\textit{nums}[i]$ 很大的课程了（还要留一天来考试）。

**如何权衡哪些课程快速复习，哪些课程慢速复习呢？**

## 提示 3

设二分的答案为 $\textit{mx}$。我们倒着遍历 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数，和 [周赛第三题](https://leetcode.cn/problems/earliest-second-to-mark-indices-i/) 一样，尽量选择靠左的时间来快速复习，这样右边就有更充足的时间用来考试。

用一个数组 $\textit{firstT}$ 记录 $1$ 到 $n$（代码中是 $0$ 到 $n-1$）在 $\textit{changeIndices}$ 中首次出现的下标。初始化可用天数 $\textit{cnt}=0$。

- 设当前天数为 $t$，设 $i = \textit{changeIndices}[t] - 1$。
- 如果 $i$ 不是在 $\textit{changeIndices}$ 中首次出现的数，或者 $\textit{nums}[i]\le 1$，那么把时间留给左边再决定做什么，$\textit{cnt}$ 加一。
- 否则如果 $\textit{cnt}>0$，我们直接快速复习第 $i$ 门课程，并消耗一天用来考试，把 $\textit{cnt}$ 减一。然后把 $\textit{nums}[i]$ 加到一个**小根堆**中。
- 否则如果 $\textit{cnt}=0$，那么尝试在小根堆中「反悔」一个复习时间比 $\textit{nums}[i]$ 小的数。如果堆为空或者堆顶大于等于 $\textit{nums}[i]$ 就不反悔，否则弹出堆顶并把 $\textit{cnt}$ 加二（一天快速复习，一天考试），然后做法同上述 $\textit{cnt}>0$ 的情况。这里从堆中弹出的课程，相当于用更靠左的时间去慢速复习+考试。

遍历结束后，对于每个未快速复习的课程，全部使用慢速复习+考试，将 $\textit{cnt}$ 减去这些课程对应的 $\textit{nums}[i]+1$。如果最终 $\textit{cnt}\ge 0$ 则说明可以在 $\textit{mx}$ 天内搞定所有课程的复习+考试。这一过程可以在遍历中动态维护，具体见代码。

下面代码用的开区间二分（其它写法也可以），原理请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def earliestSecondToMarkIndices(self, nums: List[int], changeIndices: List[int]) -> int:
        n, m = len(nums), len(changeIndices)
        total = n + sum(nums)

        first_t = [-1] * n
        for t in range(m - 1, -1, -1):
            first_t[changeIndices[t] - 1] = t

        def check(mx: int) -> bool:
            cnt = 0
            slow = total  # 慢速复习+考试所需天数
            h = []
            for t in range(mx - 1, -1, -1):
                i = changeIndices[t] - 1
                v = nums[i]
                if v <= 1 or t != first_t[i]:
                    cnt += 1  # 留给左边，用来快速复习/考试
                    continue
                if cnt == 0:
                    if not h or v <= h[0]:
                        cnt += 1  # 留给左边，用来快速复习/考试
                        continue
                    slow += heappop(h) + 1
                    cnt += 2  # 反悔：一天快速复习，一天考试
                slow -= v + 1
                cnt -= 1  # 快速复习，然后消耗一天来考试
                heappush(h, v)
            return cnt >= slow  # 剩余天数搞定慢速复习+考试

        ans = n + bisect_left(range(n, m + 1), True, key=check)
        return -1 if ans > m else ans
```

```java [sol-Java]
class Solution {
    public int earliestSecondToMarkIndices(int[] nums, int[] changeIndices) {
        int n = nums.length;
        int m = changeIndices.length;
        if (n > m) {
            return -1;
        }

        long slow = n; // 慢速复习+考试所需天数
        for (int v : nums) {
            slow += v;
        }

        int[] firstT = new int[n];
        Arrays.fill(firstT, -1);
        for (int t = m - 1; t >= 0; t--) {
            firstT[changeIndices[t] - 1] = t;
        }

        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> a - b);
        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            pq.clear();
            int mid = (left + right) / 2;
            if (check(nums, changeIndices, firstT, pq, slow, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right > m ? -1 : right;
    }

    private boolean check(int[] nums, int[] changeIndices, int[] firstT, PriorityQueue<Integer> pq, long slow, int mx) {
        int cnt = 0;
        for (int t = mx - 1; t >= 0; t--) {
            int i = changeIndices[t] - 1;
            int v = nums[i];
            if (v <= 1 || t != firstT[i]) {
                cnt++; // 留给左边，用来快速复习/考试
                continue;
            }
            if (cnt == 0) {
                if (pq.isEmpty() || v <= pq.peek()) {
                    cnt++; // 留给左边，用来快速复习/考试
                    continue;
                }
                slow += pq.poll() + 1;
                cnt += 2; // 反悔：一天快速复习，一天考试
            }
            slow -= v + 1;
            cnt--; // 快速复习，然后消耗一天来考试
            pq.offer(v);
        }
        return cnt >= slow; // 剩余天数搞定慢速复习+考试
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int earliestSecondToMarkIndices(vector<int> &nums, vector<int> &changeIndices) {
        int n = nums.size();
        int m = changeIndices.size();
        long long total = n + accumulate(nums.begin(), nums.end(), 0LL);

        vector<int> first_t(n, -1);
        for (int t = m - 1; t >= 0; t--) {
            first_t[changeIndices[t] - 1] = t;
        }

        auto check = [&](int mx) -> bool {
            int cnt = 0;
            long long slow = total; // 慢速复习+考试所需天数
            priority_queue<int, vector<int>, greater<>> pq;
            for (int t = mx - 1; t >= 0; t--) {
                int i = changeIndices[t] - 1;
                int v = nums[i];
                if (v <= 1 || t != first_t[i]) {
                    cnt++; // 留给左边，用来快速复习/考试
                    continue;
                }
                if (cnt == 0) {
                    if (pq.empty() || v <= pq.top()) {
                        cnt++; // 留给左边，用来快速复习/考试
                        continue;
                    }
                    slow += pq.top() + 1;
                    pq.pop();
                    cnt += 2; // 反悔：一天快速复习，一天考试
                }
                slow -= v + 1;
                cnt--; // 快速复习，然后消耗一天来考试
                pq.push(v);
            }
            return cnt >= slow; // 剩余天数搞定慢速复习+考试
        };

        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right > m ? -1 : right;
    }
};
```

```go [sol-Go]
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	total := n
	for _, v := range nums {
		total += v // 慢速复习+考试所需天数
	}

	firstT := make([]int, n)
	for t := m - 1; t >= 0; t-- {
		firstT[changeIndices[t]-1] = t + 1
	}

	h := hp{}
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		cnt, slow := 0, total
		h.IntSlice = h.IntSlice[:0]
		for t := mx - 1; t >= 0; t-- {
			i := changeIndices[t] - 1
			v := nums[i]
			if v <= 1 || t != firstT[i]-1 {
				cnt++ // 留给左边，用来快速复习/考试
				continue
			}
			if cnt == 0 {
				if h.Len() == 0 || v <= h.IntSlice[0] {
					cnt++ // 留给左边，用来快速复习/考试
					continue
				}
				slow += heap.Pop(&h).(int) + 1
				cnt += 2 // 反悔：一天快速复习，一天考试
			}
			slow -= v + 1
			cnt-- // 快速复习，然后消耗一天来考试
			heap.Push(&h, v)
		}
		return cnt >= slow // 剩余天数搞定慢速复习+考试
	})
	if ans > m {
		return -1
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m \log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{changeIndices}$ 的长度。二分的时候保证 $n\le m$，时间复杂度以 $m$ 为主。注意堆中至多有 $n$ 个元素。
- 空间复杂度：$\mathcal{O}(n)$。

## 题单：二分答案

- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)

## 题单：基于堆的反悔贪心（反悔堆）

- [LCP 30. 魔塔游戏](https://leetcode.cn/problems/p0NxJO/)
- [1642. 可以到达的最远建筑](https://leetcode.cn/problems/furthest-building-you-can-reach/) 1962
- [630. 课程表 III](https://leetcode.cn/problems/course-schedule-iii/)
- [871. 最低加油次数](https://leetcode.cn/problems/minimum-number-of-refueling-stops/) 2074
- [2813. 子序列最大优雅度](https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence/) 2582
- [2599. 使前缀和数组非负](https://leetcode.cn/problems/make-the-prefix-sum-non-negative/)（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
