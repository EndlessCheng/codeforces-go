# 方法一：遍历

遍历数组，用两个变量统计。

```py [sol-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        less = great = 0
        for x in nums:
            if x < 0:
                less += 1
            elif x > 0:
                great += 1
        return max(less, great)
```

```java [sol-Java]
public class Solution {
    public int maximumCount(int[] nums) {
        int less = 0;
        int great = 0;
        for (int x : nums) {
            if (x < 0) {
                less++;
            } else if (x > 0) {
                great++;
            }
        }
        return Math.max(less, great);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumCount(vector<int> &nums) {
        int less = 0, great = 0;
        for (int x : nums) {
            if (x < 0) {
                less++;
            } else if (x > 0) {
                great++;
            }
        }
        return max(less, great);
    }
};
```

```go [sol-Go]
func maximumCount(nums []int) int {
	less, great := 0, 0
	for _, x := range nums {
		if x < 0 {
			less++
		} else if x > 0 {
			great++
		}
	}
	return max(less, great)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

# 方法二：二分查找

二分查找 $\ge 0$ 和 $> 0$ 的位置，原理见 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        return max(bisect_left(nums, 0), len(nums) - bisect_right(nums, 0))
```

```java [sol-Java]
public class Solution {
    public int maximumCount(int[] nums) {
        int less = lowerBound(nums, 0);
        int great = nums.length - lowerBound(nums, 1);
        return Math.max(less, great);
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumCount(vector<int> &nums) {
        int less = lower_bound(nums.begin(), nums.end(), 0) - nums.begin();
        int great = nums.end() - upper_bound(nums.begin(), nums.end(), 0);
        return max(less, great);
    }
};
```

```go [sol-Go]
func maximumCount(nums []int) int {
	return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}
```

#### 复杂度分析

- 时间复杂度：$O(\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
