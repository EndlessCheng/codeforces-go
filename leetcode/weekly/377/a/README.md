[本题视频讲解](https://www.bilibili.com/video/BV1rG411k72D/)

把 $\textit{nums}$ 排序后，两个两个一对交换，就得到了答案。

```py [sol-Python3]
class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        nums.sort()
        for i in range(1, len(nums), 2):
            nums[i - 1], nums[i] = nums[i], nums[i - 1]
        return nums
```

```java [sol-Java]
class Solution {
    public int[] numberGame(int[] nums) {
        Arrays.sort(nums);
        for (int i = 1; i < nums.length; i += 2) {
            int tmp = nums[i - 1];
            nums[i - 1] = nums[i];
            nums[i] = tmp;
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> numberGame(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        for (int i = 1; i < nums.size(); i += 2) {
            swap(nums[i - 1], nums[i]);
        }
        return nums;
    }
};
```

```go [sol-Go]
func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
