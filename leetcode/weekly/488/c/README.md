**前置题目**：

1. [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)，视频讲解：[单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。
2. [2762. 不间断子数组](https://leetcode.cn/problems/continuous-subarrays/)，[我的题解](https://leetcode.cn/problems/continuous-subarrays/solutions/2327219/shuang-zhi-zhen-ping-heng-shu-ha-xi-biao-4frl/)。

本题只需把 2762 题的 $\texttt{while}$ 循环条件改成 `(nums[maxQ[0]] - nums[minQ[0]]) * (right - left + 1) > k` 即可。

[本题视频讲解](https://www.bilibili.com/video/BV1idFoz3Efi/?t=6m12s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        min_q = deque()
        max_q = deque()
        ans = left = 0
        for right, x in enumerate(nums):
            # 1. 入：元素进入窗口
            while min_q and x <= nums[min_q[-1]]:
                min_q.pop()
            min_q.append(right)

            while max_q and x >= nums[max_q[-1]]:
                max_q.pop()
            max_q.append(right)

            # 2. 出：如果窗口不满足要求，左端点离开窗口
            # 只需改下面这行代码，其余逻辑和 2762 题完全一致
            while (nums[max_q[0]] - nums[min_q[0]]) * (right - left + 1) > k:
                left += 1
                if min_q[0] < left:
                    min_q.popleft()
                if max_q[0] < left:
                    max_q.popleft()

            # 3. 更新答案
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, long k) {
        Deque<Integer> minQ = new ArrayDeque<>();
        Deque<Integer> maxQ = new ArrayDeque<>();
        long ans = 0;
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            // 1. 入：元素进入窗口
            int x = nums[right];
            while (!minQ.isEmpty() && x <= nums[minQ.peekLast()]) {
                minQ.pollLast();
            }
            minQ.addLast(right);

            while (!maxQ.isEmpty() && x >= nums[maxQ.peekLast()]) {
                maxQ.pollLast();
            }
            maxQ.addLast(right);

		    // 2. 出：如果窗口不满足要求，左端点离开窗口
		    // 只需改下面这行代码，其余逻辑和 2762 题完全一致
            while ((long) (nums[maxQ.peekFirst()] - nums[minQ.peekFirst()]) * (right - left + 1) > k) {
                left++;
                if (minQ.peekFirst() < left) {
                    minQ.pollFirst();
                }
                if (maxQ.peekFirst() < left) {
                    maxQ.pollFirst();
                }
            }

            // 3. 更新答案
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, long long k) {
        deque<int> min_q, max_q;
        long long ans = 0;
        int left = 0;
        for (int right = 0; right < nums.size(); right++) {
            // 1. 入：元素进入窗口
            int x = nums[right];
            while (!min_q.empty() && x <= nums[min_q.back()]) {
                min_q.pop_back();
            }
            min_q.push_back(right);

            while (!max_q.empty() && x >= nums[max_q.back()]) {
                max_q.pop_back();
            }
            max_q.push_back(right);

		    // 2. 出：如果窗口不满足要求，左端点离开窗口
		    // 只需改下面这行代码，其余逻辑和 2762 题完全一致
            while (1LL * (nums[max_q.front()] - nums[min_q.front()]) * (right - left + 1) > k) {
                left++;
                if (min_q.front() < left) {
                    min_q.pop_front();
                }
                if (max_q.front() < left) {
                    max_q.pop_front();
                }
            }

            // 3. 更新答案
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int64) (ans int64) {
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		// 1. 入：元素进入窗口
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// 2. 出：如果窗口不满足要求，左端点离开窗口
		// 只需改下面这行代码，其余逻辑和 2762 题完全一致
		for int64(nums[maxQ[0]]-nums[minQ[0]])*int64(right-left+1) > k {
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		ans += int64(right - left + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入队出队各两次（有两个单调队列），因此循环次数**之和**是 $\mathcal{O}(n)$，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 滑动窗口题单的「**§2.3.1 越短越合法**」。
2. 数据结构题单的「**§4.4 单调队列**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
