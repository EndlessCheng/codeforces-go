## O(n^2) 做法

本题是标准的划分型 DP，见 [DP 题单](https://leetcode.cn/circle/discuss/tXLS3i/) 的「§5.2 最优划分」。

一般定义 $f[i+1]$ 表示前缀 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 在题目约束下，分割出的最少（最多）子数组个数，本题是定义成分割方案数。这里 $i+1$ 是为了把 $f[0]$ 当作初始值。

枚举最后一个子数组的左端点 $j$，那么问题变成前缀 $\textit{nums}[0]$ 到 $\textit{nums}[j-1]$ 在题目约束下的分割方案数，即 $f[j]$。

当子数组右端点 $i$ 固定时，由于子数组越长，最大值越大，最小值越小，最大最小的差值越可能大于 $k$。所以符合要求的左端点 $j$ 一定在一个**连续区间** $[L,i]$ 中。累加 $f[j]$ 得

$$
f[i+1] = \sum_{j=L}^{i} f[j]
$$

初始值 $f[0] = 1$，空子数组算一个方案。也可以从递归的角度理解，递归到空子数组，就表示我们找到了一个合法分割方案。

答案为 $f[n]$。

## O(n) 做法

由于 $i$ 越大，$L$ 也越大，可以用 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

同时，我们需要计算 [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) 和滑动窗口最小值，这可以用 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)解决。

维护窗口中的 $\displaystyle\sum_{j=L}^{i} f[j]$，记作 $\textit{sumF}$，转移方程优化成

$$
f[i+1] = \textit{sumF}
$$

注意取模。关于模运算的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV113T9zFEjQ/?t=13m17s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countPartitions(self, nums: List[int], k: int) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        min_q = deque()
        max_q = deque()
        f = [0] * (n + 1)
        f[0] = 1
        sum_f = 0  # 窗口中的 f[i] 之和
        left = 0

        for i, x in enumerate(nums):
            # 1. 入
            sum_f += f[i]

            while min_q and x <= nums[min_q[-1]]:
                min_q.pop()
            min_q.append(i)

            while max_q and x >= nums[max_q[-1]]:
                max_q.pop()
            max_q.append(i)

            # 2. 出
            while nums[max_q[0]] - nums[min_q[0]] > k:
                sum_f -= f[left]
                left += 1
                if min_q[0] < left:
                    min_q.popleft()
                if max_q[0] < left:
                    max_q.popleft()

            # 3. 更新答案
            f[i + 1] = sum_f % MOD

        return f[n]
```

```java [sol-Java]
class Solution {
    public int countPartitions(int[] nums, int k) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        Deque<Integer> minQ = new ArrayDeque<>();
        Deque<Integer> maxQ = new ArrayDeque<>();
        int[] f = new int[n + 1];
        f[0] = 1;
        long sumF = 0; // 窗口中的 f[i] 之和
        int left = 0;

        for (int i = 0; i < n; i++) {
            // 1. 入
            sumF += f[i];

            int x = nums[i];
            while (!minQ.isEmpty() && x <= nums[minQ.peekLast()]) {
                minQ.pollLast();
            }
            minQ.addLast(i);

            while (!maxQ.isEmpty() && x >= nums[maxQ.peekLast()]) {
                maxQ.pollLast();
            }
            maxQ.addLast(i);

            // 2. 出
            while (nums[maxQ.peekFirst()] - nums[minQ.peekFirst()] > k) {
                sumF -= f[left];
                left++;
                if (minQ.peekFirst() < left) {
                    minQ.pollFirst();
                }
                if (maxQ.peekFirst() < left) {
                    maxQ.pollFirst();
                }
            }

            // 3. 更新答案
            f[i + 1] = (int) (sumF % MOD);
        }

        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPartitions(vector<int>& nums, int k) {
        const int MOD = 1'000'000'007;
        int n = nums.size();
        deque<int> min_q, max_q;
        vector<int> f(n + 1);
        f[0] = 1;
        long long sum_f = 0; // 窗口中的 f[i] 之和
        int left = 0;

        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 1. 入
            sum_f += f[i];

            while (!min_q.empty() && x <= nums[min_q.back()]) {
                min_q.pop_back();
            }
            min_q.push_back(i);

            while (!max_q.empty() && x >= nums[max_q.back()]) {
                max_q.pop_back();
            }
            max_q.push_back(i);

            // 2. 出
            while (nums[max_q.front()] - nums[min_q.front()] > k) {
                sum_f -= f[left];
                left++;
                if (min_q.front() < left) {
                    min_q.pop_front();
                }
                if (max_q.front() < left) {
                    max_q.pop_front();
                }
            }

            // 3. 更新答案
            f[i + 1] = sum_f % MOD;
        }

        return f[n];
    }
};
```

```go [sol-Go]
func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(nums)
	var minQ, maxQ []int
	f := make([]int, n+1)
	f[0] = 1
	sumF := 0 // 窗口中的 f[i] 之和
	left := 0

	for i, x := range nums {
		// 1. 入
		sumF += f[i]

		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		// 2. 出
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			sumF -= f[left]
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		f[i+1] = sumF % mod
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个下标至多入队出队各两次。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单的「**§5.2 最优划分**」和「**§11.3 单调队列优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
