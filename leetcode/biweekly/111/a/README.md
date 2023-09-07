### 前置知识：相向双指针

请看[【基础算法精讲】](https://www.bilibili.com/video/BV1bP411c7oJ/)。

### 思路

为什么可以排序呢？题目相当于从数组中选两个数，**我们只关心这两个数的和是否小于** $\textit{target}$，由于 $a+b=b+a$，无论如何排列数组元素，都不会影响加法的结果，所以排序不影响答案。

排序后：

- 初始化左右指针 $\textit{left}=0,\textit{right}=n-1$。
- 如果 $\textit{nums}[\textit{left}]+\textit{nums}[\textit{right}] < \textit{target}$，由于数组是有序的，$\textit{nums}[\textit{left}]$ 与下标 $i$ 在 $[\textit{left}+1,\textit{right}]$ 中的任何 $\textit{nums}[i]$ 相加，都是 $<\textit{target}$ 的，因此直接找到了 $\textit{right}-\textit{left}$ 个合法数对，加到答案中，然后将 $\textit{left}$ 加一。
- 如果 $\textit{nums}[\textit{left}]+\textit{nums}[\textit{right}] \ge \textit{target}$，由于数组是有序的，$\textit{nums}[\textit{right}]$ 与下标 $i$ 在 $[\textit{left},\textit{right}-1]$ 中的任何 $\textit{nums}[i]$ 相加，都是 $\ge\textit{target}$ 的，因此后面无需考虑 $\textit{nums}[\textit{right}]$，将 $\textit{right}$ 减一。
- 重复上述过程直到 $\textit{left}\ge \textit{right}$ 为止。

```py [sol-Python3]
class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans = left = 0
        right = len(nums) - 1
        while left < right:
            if nums[left] + nums[right] < target:
                ans += right - left
                left += 1
            else:
                right -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countPairs(List<Integer> nums, int target) {
        Collections.sort(nums);
        int ans = 0, left = 0, right = nums.size() - 1;
        while (left < right) {
            if (nums.get(left) + nums.get(right) < target) {
                ans += right - left;
                left++;
            } else {
                right--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPairs(vector<int> &nums, int target) {
        sort(nums.begin(), nums.end());
        int ans = 0, left = 0, right = nums.size() - 1;
        while (left < right) {
            if (nums[left] + nums[right] < target) {
                ans += right - left;
                left++;
            } else {
                right--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}
```

```js [sol-JavaScript]
var countPairs = function (nums, target) {
    nums.sort((a, b) => a - b);
    let ans = 0, left = 0, right = nums.length - 1;
    while (left < right) {
        if (nums[left] + nums[right] < target) {
            ans += right - left;
            left++;
        } else {
            right--;
        }
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。不计入排序的栈开销，仅用到若干额外变量。
