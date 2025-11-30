窗口（子数组）越长，所包含的质数的最大值越大，最小值越小，质数极差越大；反之，窗口越短，质数极差越小。

有这样的性质，可以用**不定长滑动窗口**解决，原理请看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

设 $[0,i]$ 中的倒数第二个质数的下标为 $\textit{last}_2$。若不存在，则 $\textit{last}_2=-1$。

设滑动窗口的右端点为 $i$，左端点为 $\textit{left}$。注意这意味着 $[\textit{left},i]$ 是一个合法子数组，且左端点不能再小了。

当窗口右端点**固定**为 $i$ 时，合法子数组**左端点**的范围为 $[\textit{left},\textit{last}_2]$。这有

$$
\textit{last}_2 - \textit{left} + 1
$$

个，所以我们找到了 $\textit{last}_2 - \textit{left} + 1$ 个右端点固定为 $i$ 的合法子数组，加入答案。

**注**：当 $\textit{left} = \textit{last}_2+1$ 时，窗口中至多有一个质数，一定满足要求，所以 $\textit{left} \le  \textit{last}_2+1$ 恒成立，即 $\textit{last}_2 - \textit{left} + 1\ge 0$ 恒成立。另请注意，$\textit{last}_2=-1$ 这一初始值同样满足这一性质。

维护窗口内的质数最大值和最小值可以用两个**单调队列**。原理讲解：[单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

代码实现时，可以用 [埃氏筛](https://www.bilibili.com/video/BV1GCNRzgEYp/)（或者欧拉筛）预处理一个布尔数组，表示哪些数是质数。注意 $1$ 不是质数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1qeNRzjEEk/?t=13m22s)，欢迎点赞关注~

```py [sol-Python3]
MX = 50_001
is_prime = [False] * 2 + [True] * (MX - 2)
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

class Solution:
    def primeSubarray(self, nums: List[int], k: int) -> int:
        min_q = deque()
        max_q = deque()
        last = last2 = -1
        ans = left = 0

        for i, x in enumerate(nums):
            if is_prime[x]:
                # 1. 入
                last2 = last
                last = i

                while min_q and x <= nums[min_q[-1]]:
                    min_q.pop()
                min_q.append(i)

                while max_q and x >= nums[max_q[-1]]:
                    max_q.pop()
                max_q.append(i)

                # 2. 出
                while nums[max_q[0]] - nums[min_q[0]] > k:
                    left += 1
                    if min_q[0] < left:
                        min_q.popleft()
                    if max_q[0] < left:
                        max_q.popleft()

            # 3. 更新答案
            ans += last2 - left + 1

        return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 50_001;
    private static final boolean[] NOT_PRIME = new boolean[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        NOT_PRIME[1] = true; // 1 不是质数
        for (int i = 2; i * i < MX; i++) {
            if (NOT_PRIME[i]) {
                continue;
            }
            for (int j = i * i; j < MX; j += i) {
                NOT_PRIME[j] = true; // j 是质数 i 的倍数
            }
        }
    }

    public int primeSubarray(int[] nums, int k) {
        init();

        Deque<Integer> minQ = new ArrayDeque<>();
        Deque<Integer> maxQ = new ArrayDeque<>();
        int last = -1, last2 = -1;
        int ans = 0, left = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            if (!NOT_PRIME[x]) {
                // 1. 入
                last2 = last;
                last = i;

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
                    left++;
                    if (minQ.peekFirst() < left) {
                        minQ.pollFirst();
                    }
                    if (maxQ.peekFirst() < left) {
                        maxQ.pollFirst();
                    }
                }
            }

            // 3. 更新答案
            ans += last2 - left + 1;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
const int MX = 50'001;
bool np[MX];

auto init = [] {
    np[1] = true; // 1 不是质数
    for (int i = 2; i * i < MX; i++) {
        if (!np[i]) {
            for (int j = i * i; j < MX; j += i) {
                np[j] = true; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int primeSubarray(vector<int>& nums, int k) {
        deque<int> min_q, max_q;
        int last = -1, last2 = -1;
        int ans = 0, left = 0;

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];

            if (!np[x]) {
                // 1. 入
                last2 = last;
                last = i;

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
                    left++;
                    if (min_q.front() < left) {
                        min_q.pop_front();
                    }
                    if (max_q.front() < left) {
                        max_q.pop_front();
                    }
                }
            }

            // 3. 更新答案
            ans += last2 - left + 1;
        }

        return ans;
    }
};
```

```go [sol-Go]
const mx = 50_001
var np = [mx]bool{1: true} // 1 不是质数

func init() {
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubarray(nums []int, k int) (ans int) {
	var minQ, maxQ []int
	last, last2 := -1, -1
	left := 0

	for i, x := range nums {
		if !np[x] {
			// 1. 入
			last2 = last
			last = i

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
				left++
				if minQ[0] < left {
					minQ = minQ[1:]
				}
				if maxQ[0] < left {
					maxQ = maxQ[1:]
				}
			}
		}

		// 3. 更新答案
		ans += last2 - left + 1
	}

	return
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)
- [3578. 统计极差最大为 K 的分割方式数](https://leetcode.cn/problems/count-partitions-with-max-min-difference-at-most-k/)

更多相似题目，见

1. 滑动窗口题单的「**§2.3.2 越短越合法**」。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
