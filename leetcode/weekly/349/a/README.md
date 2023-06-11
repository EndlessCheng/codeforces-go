### 思路

由于数组元素各不相同，所以答案一定在前三个数中，不妨取前三个数的中间值。

```py [sol-Python3]
class Solution:
    def findNonMinOrMax(self, nums: List[int]) -> int:
        return sorted(nums[:3])[1] if len(nums) > 2 else -1
```

```java [sol-Java]
class Solution {
    public int findNonMinOrMax(int[] nums) {
        if (nums.length < 3) return -1;
        Arrays.sort(nums, 0, 3); // 只对前三个数排序
        return nums[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findNonMinOrMax(vector<int> &nums) {
        if (nums.size() < 3) return -1;
        sort(nums.begin(), nums.begin() + 3); // 只对前三个数排序
        return nums[1];
    }
};
```

```go [sol-Go]
func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums[:3]) // 只对前三个数排序
	return nums[1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
