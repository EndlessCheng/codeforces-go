**前置题目**：[239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)，视频讲解请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

题目要求机器人连续运行，看成一个连续子数组，题目要求计算最长子数组长度。

枚举子数组右端点 $\textit{right}$，我们需要知道此时左端点 $\textit{left}$ 的最小值，这样子数组尽量长。

由于有 $\textit{budget}$ 的限制，所以 $\textit{right}$ 越大，$\textit{left}$ 也越大，有单调性，可以用**滑动窗口**解决。

本题的一种做法是二分答案，这样就转换成了固定长度的 239 题。

但实际上不用二分，在 239 题的基础上，把定长滑窗改为不定长滑窗，套路如下：

1. **入**：$\textit{chargeTimes}[\textit{right}]$ 进入窗口时，弹出队尾的 $\le \textit{chargeTimes}[\textit{right}]$ 的元素。
2. **出**：如果总开销超过 $\textit{budget}$，则不断移出左端点，直到总开销不超过 $\textit{budget}$。特别地，如果左端点恰好等于队首，则弹出队首。
3. **更新答案**：用窗口长度 $\textit{right}-\textit{left}+1$ 更新答案的最大值。

⚠**注意**：为了方便判断队首是否要出队，单调队列中保存的是下标。

```py [sol-Python3]
class Solution:
    def maximumRobots(self, chargeTimes: List[int], runningCosts: List[int], budget: int) -> int:
        ans = s = left = 0
        q = deque()
        for right, (t, c) in enumerate(zip(chargeTimes, runningCosts)):
            # 1. 入
            while q and t >= chargeTimes[q[-1]]:
                q.pop()
            q.append(right)
            s += c  # 维护 sum(runningCosts)

            # 2. 出
            while q and chargeTimes[q[0]] + (right - left + 1) * s > budget:
                if q[0] == left:
                    q.popleft()
                s -= runningCosts[left]  # 维护 sum(runningCosts)
                left += 1

            # 3. 更新答案
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumRobots(int[] chargeTimes, int[] runningCosts, long budget) {
        int ans = 0;
        int left = 0;
        long sum = 0;
        Deque<Integer> q = new ArrayDeque<>();
        for (int right = 0; right < chargeTimes.length; right++) {
            // 1. 入
            while (!q.isEmpty() && chargeTimes[right] >= chargeTimes[q.peekLast()]) {
                q.pollLast();
            }
            q.addLast(right);
            sum += runningCosts[right];

            // 2. 出
            while (!q.isEmpty() && chargeTimes[q.peekFirst()] + (right - left + 1) * sum > budget) {
                if (q.peekFirst() == left) {
                    q.pollFirst();
                }
                sum -= runningCosts[left++];
            }

            // 3. 更新答案
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int maximumRobots(int[] chargeTimes, int[] runningCosts, long budget) {
        int ans = 0;
        int left = 0;
        long sum = 0;
        int[] q = new int[chargeTimes.length];
        int head = 0; // 队头
        int tail = 0; // 队尾+1
        for (int right = 0; right < chargeTimes.length; right++) {
            // 1. 入
            while (head < tail && chargeTimes[right] >= chargeTimes[q[tail - 1]]) {
                tail--;
            }
            q[tail++] = right;
            sum += runningCosts[right];

            // 2. 出
            while (head < tail && chargeTimes[q[head]] + (right - left + 1) * sum > budget) {
                if (q[head] == left) {
                    head++;
                }
                sum -= runningCosts[left++];
            }

            // 3. 更新答案
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
        int ans = 0, left = 0;
        long long sum = 0;
        deque<int> q;
        for (int right = 0; right < chargeTimes.size(); right++) {
            // 1. 入
            while (!q.empty() && chargeTimes[right] >= chargeTimes[q.back()]) {
                q.pop_back();
            }
            q.push_back(right);
            sum += runningCosts[right];

            // 2. 出
            while (!q.empty() && chargeTimes[q.front()] + (right - left + 1) * sum > budget) {
                if (q.front() == left) {
                    q.pop_front();
                }
                sum -= runningCosts[left++];
            }

            // 3. 更新答案
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maximumRobots(int* chargeTimes, int chargeTimesSize, int* runningCosts, int runningCostsSize, long long budget) {
    int ans = 0, left = 0;
    long long sum = 0;
    int* q = malloc(chargeTimesSize * sizeof(int));
    int head = 0, tail = 0; // 队头和队尾
    for (int right = 0; right < chargeTimesSize; right++) {
        // 1. 入
        while (head < tail && chargeTimes[right] >= chargeTimes[q[tail - 1]]) {
            tail--;
        }
        q[tail++] = right;
        sum += runningCosts[right];

        // 2. 出
        while (head < tail && chargeTimes[q[head]] + (right - left + 1) * sum > budget) {
            if (q[head] == left) {
                head++;
            }
            sum -= runningCosts[left++];
        }

        // 3. 更新答案
        ans = MAX(ans, right - left + 1);
    }
    free(q);
    return ans;
}
```

```go [sol-Go]
func maximumRobots(chargeTimes, runningCosts []int, budget int64) (ans int) {
    q := []int{}
    sum := int64(0)
    left := 0
    for right, t := range chargeTimes {
        // 1. 入
        for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, right)
        sum += int64(runningCosts[right])

        // 2. 出
        for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
            if q[0] == left {
                q = q[1:]
            }
            sum -= int64(runningCosts[left])
            left++
        }

        // 3. 更新答案
        ans = max(ans, right-left+1)
    }
    return
}
```

```js [sol-JavaScript]
var maximumRobots = function(chargeTimes, runningCosts, budget) {
    let ans = 0, left = 0, sum = 0;
    const q = Array(chargeTimes.length);
    let head = 0, tail = 0; // 队头和队尾
    for (let right = 0; right < chargeTimes.length; right++) {
        // 1. 入
        while (head < tail && chargeTimes[right] >= chargeTimes[q[tail - 1]]) {
            tail--;
        }
        q[tail++] = right;
        sum += runningCosts[right];

        // 2. 出
        while (head < tail && chargeTimes[q[head]] + (right - left + 1) * sum > budget) {
            if (q[head] === left) {
                head++;
            }
            sum -= runningCosts[left++];
        }

        // 3. 更新答案
        ans = Math.max(ans, right - left + 1);
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::VecDeque;

impl Solution {
    pub fn maximum_robots(charge_times: Vec<i32>, running_costs: Vec<i32>, budget: i64) -> i32 {
        let mut ans = 0;
        let mut left = 0;
        let mut sum = 0i64;
        let mut q = VecDeque::new();
        for right in 0..charge_times.len() {
            // 1. 入
            while !q.is_empty() && charge_times[right] >= charge_times[*q.back().unwrap()] {
                q.pop_back();
            }
            q.push_back(right);
            sum += running_costs[right] as i64;

            // 2. 出
            while !q.is_empty() && charge_times[*q.front().unwrap()] as i64 + (right - left + 1) as i64 * sum > budget {
                if *q.front().unwrap() == left {
                    q.pop_front();
                }
                sum -= running_costs[left] as i64;
                left += 1;
            }

            // 3. 更新答案
            ans = ans.max(right - left + 1);
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{chargeTimes}$ 的长度。虽然有二重循环，但是每个元素至多出队一次，以及 $\textit{left}$ 最多增加 $n$ 次。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

把「子数组」改成「子序列」要怎么做？

欢迎在评论区发表的你的思路/代码。

更多相似题目，见下面数据结构题单中的「**§4.3 单调队列**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
