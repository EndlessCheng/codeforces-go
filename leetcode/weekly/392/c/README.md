把 $\textit{nums}$ 从小到大排序后，中位数为 $\textit{nums}[m]$，其中 $m=\left\lfloor\dfrac{n}{2}\right\rfloor$，$n$ 为 $\textit{nums}$ 的长度。

我们需要把中位数左边的数都变成 $\le k$ 的，右边的数都变成 $\ge k$ 的。

分类讨论：

- 如果 $\textit{nums}[m] > k$，要把下标在 $[0,m]$ 中的大于 $k$ 的数都变成 $k$，由于下标在 $[m+1,n-1]$ 中的数已经大于 $k$（因为数组是有序的），所以下标在 $[m+1,n-1]$ 中的数无需操作。
- 如果 $\textit{nums}[m] < k$，要把下标在 $[m,n-1]$ 中的小于 $k$ 的数都变成 $k$，由于下标在 $[0,m-1]$ 中的数已经小于 $k$（因为数组是有序的），所以下标在 $[0,m-1]$ 中的数无需操作。

累加元素的变化量，即为答案。具体请看 [视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minOperationsToMakeMedianK(self, nums: List[int], k: int) -> int:
        nums.sort()
        m = len(nums) // 2
        ans = 0
        if nums[m] > k:
            for i in range(m, -1, -1):
                if nums[i] <= k:
                    break
                ans += nums[i] - k
        else:
            for i in range(m, len(nums)):
                if nums[i] >= k:
                    break
                ans += k - nums[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long minOperationsToMakeMedianK(int[] nums, int k) {
        Arrays.sort(nums);
        long ans = 0;
        int m = nums.length / 2;
        if (nums[m] > k) {
            for (int i = m; i >= 0 && nums[i] > k; i--) {
                ans += nums[i] - k;
            }
        } else {
            for (int i = m; i < nums.length && nums[i] < k; i++) {
                ans += k - nums[i];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minOperationsToMakeMedianK(vector<int> &nums, int k) {
        ranges::sort(nums);
        long long ans = 0;
        int m = nums.size() / 2;
        if (nums[m] > k) {
            for (int i = m; i >= 0 && nums[i] > k; i--) {
                ans += nums[i] - k;
            }
        } else {
            for (int i = m; i < nums.size() && nums[i] < k; i++) {
                ans += k - nums[i];
            }
        }
        return ans;
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long minOperationsToMakeMedianK(vector<int> &nums, int k) {
        int m = nums.size() / 2;
        ranges::nth_element(nums, nums.begin() + m);
        long long ans = 0;
        if (nums[m] > k) {
            for (int i = 0; i <= m; i++) {
                ans += max(nums[i] - k, 0);
            }
        } else {
            for (int i = m; i < nums.size(); i++) {
                ans += max(k - nums[i], 0);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperationsToMakeMedianK(nums []int, k int) (ans int64) {
	slices.Sort(nums)
	m := len(nums) / 2
	if nums[m] > k {
		for i := m; i >= 0 && nums[i] > k; i-- {
			ans += int64(nums[i] - k)
		}
	} else {
		for i := m; i < len(nums) && nums[i] < k; i++ {
			ans += int64(k - nums[i])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。用快速选择算法可以做到期望 $\mathcal{O}(n)$ 的时间复杂度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 思考题

改成输入 $10^5$ 个询问，每个询问包含一个 $k$，如何高效地回答每个询问？

欢迎在评论区发表你的思路。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
