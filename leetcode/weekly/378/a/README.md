[本题视频讲解](https://www.bilibili.com/video/BV1XG411B7bX/)

$\textit{nums}$ 中的奇数肯定不能参与或运算，这会导致或运算结果的最低位必然是 $1$。

所以只有偶数能参与或运算，这样最低位必然是 $0$。

所以判断 $\textit{nums}$ 中是否至少有两个偶数即可。

```py [sol-Python3]
class Solution:
    def hasTrailingZeros(self, nums: List[int]) -> bool:
        return len(nums) - sum(x % 2 for x in nums) >= 2
```

```java [sol-Java]
class Solution {
    public boolean hasTrailingZeros(int[] nums) {
        int even = nums.length;
        for (int x : nums) {
            even -= x % 2;
        }
        return even >= 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool hasTrailingZeros(vector<int> &nums) {
        int even = nums.size();
        for (int x: nums) {
            even -= x % 2;
        }
        return even >= 2;
    }
};
```

```go [sol-Go]
func hasTrailingZeros(nums []int) bool {
	even := len(nums)
	for _, x := range nums {
		even -= x % 2
	}
	return even >= 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
