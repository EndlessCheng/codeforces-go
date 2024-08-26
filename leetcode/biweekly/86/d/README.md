**前置知识**：[239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) 以及 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

本题的一种做法是二分答案，这样就转换成了 239 题。

但实际上不用二分，在 239 这题的基础上，把定长滑窗改为**不定长滑窗**，具体见代码注释。

本题 [视频讲解](https://www.bilibili.com/video/BV1na41137jv)。

```py [sol-Python3]
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

```java [sol-Java]
class Solution {
    public int maximumRobots(int[] chargeTimes, int[] runningCosts, long budget) {
        int ans = 0;
        long sum = 0;
        Deque<Integer> q = new ArrayDeque<>();
        // 枚举区间右端点 right，计算区间左端点 left 的最小值
        for (int left = 0, right = 0; right < chargeTimes.length; right++) {
            // 及时清除队列中的无用数据，保证队列的单调性
            while (!q.isEmpty() && chargeTimes[right] >= chargeTimes[q.peekLast()]) {
                q.pollLast();
            }
            q.addLast(right);
            sum += runningCosts[right];
            // 如果左端点 left 不满足要求，就不断右移 left
            while (!q.isEmpty() && chargeTimes[q.peekFirst()] + (right - left + 1) * sum > budget) {
                // 及时清除队列中的无用数据，保证队列的单调性
                if (q.peekFirst() == left) {
                    q.pollFirst();
                }
                sum -= runningCosts[left++];
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumRobots(vector<int>& chargeTimes, vector<int>& runningCosts, long long budget) {
        int ans = 0;
        long long sum = 0;
        deque<int> q;
        // 枚举区间右端点 right，计算区间左端点 left 的最小值
        for (int left = 0, right = 0; right < chargeTimes.size(); right++) {
            // 及时清除队列中的无用数据，保证队列的单调性
            while (!q.empty() && chargeTimes[right] >= chargeTimes[q.back()]) {
                q.pop_back();
            }
            q.push_back(right);
            sum += runningCosts[right];
            // 如果左端点 left 不满足要求，就不断右移 left
            while (!q.empty() && chargeTimes[q.front()] + (right - left + 1) * sum > budget) {
                // 及时清除队列中的无用数据，保证队列的单调性
                if (q.front() == left) {
                    q.pop_front();
                }
                sum -= runningCosts[left++];
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumRobots(chargeTimes, runningCosts []int, budget int64) (ans int) {
	sum := int64(0)
	left := 0
	q := []int{}
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
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{chargeTimes}$ 的长度。虽然有二重循环，但是每个元素至多出队一次，以及 $\textit{left}$ 最多增加 $n$ 次。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

把「子数组」改成「子序列」要怎么做？

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
