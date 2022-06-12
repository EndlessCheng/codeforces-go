双指针使用前提：

1. 子数组（连续）；
2. 有单调性。本题元素均为正数，这意味着只要某个子数组满足题目要求，在该子数组内的更短的子数组同样也满足题目要求。

```Python [sol1-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans = s = left = 0
        for right, num in enumerate(nums):
            s += num
            while s * (right - left + 1) >= k:
                s -= nums[left]
                left += 1
            ans += right - left + 1
        return ans
```

```java [sol1-Java]
class Solution {
    public long countSubarrays(int[] nums, long k) {
        long ans = 0L, sum = 0L;
        for (int left = 0, right = 0; right < nums.length; right++) {
            sum += nums[right];
            while (sum * (right - left + 1) >= k)
                sum -= nums[left++];
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long countSubarrays(vector<int> &nums, long long k) {
        long ans = 0L, sum = 0L;
        for (int left = 0, right = 0; right < nums.size(); ++right) {
            sum += nums[right];
            while (sum * (right - left + 1) >= k)
                sum -= nums[left++];
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countSubarrays(nums []int, k int64) (ans int64) {
	sum, left := int64(0), 0
	for right, num := range nums {
		sum += int64(num)
		for sum*int64(right-left+1) >= k {
			sum -= int64(nums[left])
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
```

