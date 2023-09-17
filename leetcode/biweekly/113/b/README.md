证明过程请看 [视频讲解](https://b23.tv/MRNaADm) 第二题。

假设 $x$ 出现次数最多，其出现次数为 $\textit{maxCnt}$。

分类讨论：

- 如果 $\textit{maxCnt}\cdot 2 > n$，其余所有 $n-\textit{maxCnt}$ 个数都要与 $x$ 消除，所以最后剩下 $\textit{maxCnt}\cdot 2 - n$ 个数。
- 如果 $\textit{maxCnt}\cdot 2 \le n$ 且 $n$ 是偶数，那么可以把其余数消除至剩下 $\textit{maxCnt}$ 个数，然后再和 $x$ 消除，最后剩下 $0$ 个数。
- 如果 $\textit{maxCnt}\cdot 2 \le n$ 且 $n$ 是奇数，同上，最后剩下 $1$ 个数。

所以本题核心是计算 $\textit{maxCnt}$，这可以遍历一遍 $\textit{nums}$ 算出来。

但我们还可以更快！

由于 $\textit{nums}$ 是有序的，如果 $\textit{maxCnt}$ 超过数组长度的一半，那么 $\textit{nums}[n/2]$ 一定是出现次数最多的那个数！

按照 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) 的做法，可以用二分查找在 $\mathcal{O}(\log n)$ 的时间计算 $\textit{nums}[n/2]$ 第一次和最后一次出现的位置，从而算出 $\textit{maxCnt}$。

关于二分的原理（34 题）请看视频讲解[【基础算法精讲 04】](https://b23.tv/CBJnyNJ)。

```py [sol-Python3]
class Solution:
    def minLengthAfterRemovals(self, nums: List[int]) -> int:
        n = len(nums)
        x = nums[n // 2]
        max_cnt = bisect_right(nums, x) - bisect_left(nums, x)
        return max(max_cnt * 2 - n, n % 2)
```

```java [sol-Java]
class Solution {
    public int minLengthAfterRemovals(List<Integer> nums) {
        int n = nums.size();
        int x = nums.get(n / 2);
        int maxCnt = lowerBound(nums, x + 1) - lowerBound(nums, x);
        return Math.max(maxCnt * 2 - n, n % 2);
    }

    // 开区间写法
    private int lowerBound(List<Integer> nums, int target) {
        int left = -1, right = nums.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums.get(mid) < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLengthAfterRemovals(vector<int> &nums) {
        int n = nums.size();
        int x = nums[n / 2];
        int max_cnt = upper_bound(nums.begin(), nums.end(), x) -
                      lower_bound(nums.begin(), nums.end(), x);
        return max(max_cnt * 2 - n, n % 2);
    }
};
```

```go [sol-Go]
func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	x := nums[n/2]
	maxCnt := sort.SearchInts(nums, x+1) - sort.SearchInts(nums, x)
	return max(maxCnt*2-n, n%2)
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var minLengthAfterRemovals = function (nums) {
    const n = nums.length;
    const x = nums[n >> 1];
    const maxCnt = lowerBound(nums, x + 1) - lowerBound(nums, x);
    return Math.max(maxCnt * 2 - n, n % 2);
};

var lowerBound = function (nums, target) {
    let left = -1, right = nums.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        // 循环不变量：
        // nums[left] < target
        // nums[right] >= target
        const mid = left + ((right - left) >> 1);
        if (nums[mid] < target)
            left = mid; // 范围缩小到 (mid, right)
        else
            right = mid; // 范围缩小到 (left, mid)
    }
    return right; // 或者 left+1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
