前置题目：[209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)，视频讲解：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

本题重复元素不能计入窗口元素和。

我们可以用一个哈希表 $\textit{cnt}$ 维护窗口内每个元素的出现次数。

- 元素 $x$ 进入窗口后，如果 $x$ 在窗口内的出现次数等于 $1$，把 $x$ 加到窗口元素和中。
- 元素 $x$ 离开窗口后，如果 $x$ 在窗口内的出现次数等于 $0$，把 $x$ 从窗口元素和中减掉。

```py [sol-Python3]
class Solution:
    def minLength(self, nums: List[int], k: int) -> int:
        cnt = defaultdict(int)
        s = left = 0
        ans = inf

        for i, x in enumerate(nums):
            # 1. 入
            cnt[x] += 1
            if cnt[x] == 1:
                s += x

            while s >= k:
                # 2. 更新答案
                ans = min(ans, i - left + 1)

                # 3. 出
                out = nums[left]
                cnt[out] -= 1
                if cnt[out] == 0:
                    s -= out
                left += 1

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minLength(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>(); // 更快的写法见【Java 数组】
        int sum = 0;
        int left = 0;
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < nums.length; i++) {
            // 1. 入
            int x = nums[i];
            int c = cnt.merge(x, 1, Integer::sum);
            if (c == 1) {
                sum += x;
            }

            while (sum >= k) {
                // 2. 更新答案
                ans = Math.min(ans, i - left + 1);

                // 3. 出
                int out = nums[left];
                c = cnt.merge(out, -1, Integer::sum);
                if (c == 0) {
                    sum -= out;
                }
                left++;
            }
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minLength(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        int sum = 0;
        int left = 0;
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < nums.length; i++) {
            // 1. 入
            int x = nums[i];
            cnt[x]++;
            if (cnt[x] == 1) {
                sum += x;
            }

            while (sum >= k) {
                // 2. 更新答案
                ans = Math.min(ans, i - left + 1);

                // 3. 出
                int out = nums[left];
                cnt[out]--;
                if (cnt[out] == 0) {
                    sum -= out;
                }
                left++;
            }
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(vector<int>& nums, int k) {
        unordered_map<int, int> cnt;
        int sum = 0;
        int left = 0;
        int ans = INT_MAX;

        for (int i = 0; i < nums.size(); i++) {
            // 1. 入
            int x = nums[i];
            cnt[x]++;
            if (cnt[x] == 1) {
                sum += x;
            }

            while (sum >= k) {
                // 2. 更新答案
                ans = min(ans, i - left + 1);

                // 3. 出
                int out = nums[left];
                cnt[out]--;
                if (cnt[out] == 0) {
                    sum -= out;
                }
                left++;
            }
        }

        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minLength(nums []int, k int) int {
	cnt := map[int]int{}
	sum := 0
	left := 0
	ans := math.MaxInt

	for i, x := range nums {
		// 1. 入
		cnt[x]++
		if cnt[x] == 1 {
			sum += x
		}

		for sum >= k {
			// 2. 更新答案
			ans = min(ans, i-left+1)

			// 3. 出
			out := nums[left]
			cnt[out]--
			if cnt[out] == 0 {
				sum -= out
			}
			left++
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面滑动窗口题单的「**§2.2 越长越合法/求最短/最小**」。

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
