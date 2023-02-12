### 前置知识：二分查找

有关二分查找的写法，可以看我的 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频。

**视频中详细介绍了如何把「小于」和「小于等于」转换成「大于等于」。**

### 思路

由于排序不会影响数对的个数，为了能够二分，可以先排序。

然后枚举 $\textit{nums}[j]$，二分查找符合要求的 $\textit{nums}[i]$ 的个数。

详细的计算过程请看本题的 [视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)。

```py [sol1-Python3]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        ans = 0
        nums.sort()
        for j, x in enumerate(nums):
            r = bisect_right(nums, upper - x, 0, j)
            l = bisect_left(nums, lower - x, 0, j)
            ans += r - l
        return ans
```

```java [sol1-Java]
class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        long ans = 0;
        Arrays.sort(nums);
        for (int j = 0; j < nums.length; ++j) {
            int r = lowerBound(nums, j, upper - nums[j] + 1);
            int l = lowerBound(nums, j, lower - nums[j]);
            ans += r - l;
        }
        return ans;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int right, int target) {
        int left = -1; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = (left + right) >>> 1;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long countFairPairs(vector<int> &nums, int lower, int upper) {
        long long ans = 0;
        sort(nums.begin(), nums.end());
        for (int j = 0; j < nums.size(); ++j) {
            auto r = upper_bound(nums.begin(), nums.begin() + j, upper - nums[j]);
            auto l = lower_bound(nums.begin(), nums.begin() + j, lower - nums[j]);
            ans += r - l;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countFairPairs(nums []int, lower, upper int) (ans int64) {
	sort.Ints(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序的栈开销，仅用到若干额外变量。
