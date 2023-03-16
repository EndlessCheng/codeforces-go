### 前置知识：二分查找

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 思路

贪心：由于元素和有上限，为了能让子序列尽量长，子序列中的**元素值越小越好**。

注意「元素和」以及「子序列的长度」都与子序列元素的**下标无关**，换句话说元素在 $\textit{nums}$ 中的位置与答案是无关的，那么就可以排序了。

把 $\textit{nums}$ 从小到大排序后，再从小到大选择尽量多的元素（相当于选择一个前缀），使这些元素的和不超过询问值。

### 时间优化

既然求的是前缀的元素和（前缀和），那么干脆把每个前缀和都算出来。

做法是递推：前 $i$ 个数的元素和，等于前 $i-1$ 个数的元素和，加上第 $i$ 个数的值。

例如 $[4,5,2,1]$ 排序后为 $[1,2,4,5]$，从左到右递推计算前缀和，得到 $[1,3,7,12]$。

由于 $\textit{nums}[i]$ 都是正整数，前缀和是**严格单调递增**的，这样就能在前缀和上使用二分查找：找到大于 $\textit{queries}[i]$ 的第一个数的下标，由于下标是从 $0$ 开始的，这个数的下标正好就是前缀和小于等于 $\textit{queries}[i]$ 的最长前缀的长度。

例如在 $[1,3,7,12]$ 二分查找大于 $3$ 的第一个数（$7$），得到下标 $2$，这正好就是前缀和小于等于 $3$ 的最长前缀长度。对应到 $\textit{nums}$ 中，就是选择了 $2$ 个数（$1$ 和 $2$）作为子序列中的元素。

### 空间优化

前缀和可以直接记到 $\textit{nums}$ 中。

答案可以直接记到 $\textit{queries}$ 中。

### 视频讲解

本题 [视频讲解](https://www.bilibili.com/video/BV1mG411V7fj)

```py [sol1-Python3]
class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        nums.sort()
        for i in range(1, len(nums)):
            nums[i] += nums[i - 1]  # 原地求前缀和
        for i, q in enumerate(queries):
            queries[i] = bisect_right(nums, q)  # 复用 queries 作为答案
        return queries
```

```java [sol1-Java]
class Solution {
    public int[] answerQueries(int[] nums, int[] queries) {
        Arrays.sort(nums);
        for (int i = 1; i < nums.length; ++i)
            nums[i] += nums[i - 1]; // 原地求前缀和
        for (int i = 0; i < queries.length; i++)
            queries[i] = upperBound(nums, queries[i]); // 复用 queries 作为答案
        return queries;
    }

    // https://www.bilibili.com/video/BV1AP41137w7/
    // 返回 nums 中第一个大于 target 的数的下标（注意是大于，不是大于等于）
    // 如果这样的数不存在，则返回 nums.length
    // 时间复杂度 O(log nums.length)
    // 采用开区间写法实现
    private int upperBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = left + (right - left) / 2;
            if (nums[mid] > target)
                right = mid; // 范围缩小到 (left, mid)
            else
                left = mid; // 范围缩小到 (mid, right)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> answerQueries(vector<int> &nums, vector<int> &queries) {
        sort(nums.begin(), nums.end());
        for (int i = 1; i < nums.size(); ++i)
            nums[i] += nums[i - 1]; // 原地求前缀和
        for (int &q : queries) // 复用 queries 作为答案
            q = upper_bound(nums.begin(), nums.end(), q) - nums.begin();
        return queries;
    }
};
```

```go [sol1-Go]
func answerQueries(nums, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1] // 原地求前缀和
	}
	for i, q := range queries {
		queries[i] = sort.SearchInts(nums, q+1) // 复用 queries 作为答案
	}
	return queries
}
```

### 复杂度分析

- 时间复杂度：$O((n+m)\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{queries}$ 的长度。排序为 $O(n\log n)$，$m$ 次二分查找为 $O(m\log n)$。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。

### 思考题

把子序列改成子数组要怎么做？

---

附：我的 [每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)，已分类整理好。

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~
