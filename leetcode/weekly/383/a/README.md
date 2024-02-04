一边遍历数组，一边累加元素，如果发现累加值等于 $0$，说明返回到边界，把答案加一。

```py [sol-Python3]
class Solution:
    def returnToBoundaryCount(self, nums: List[int]) -> int:
        return sum(s == 0 for s in accumulate(nums))
```

```java [sol-Java]
class Solution {
    public int returnToBoundaryCount(int[] nums) {
        int ans = 0;
        int sum = 0;
        for (int x : nums) {
            sum += x;
            if (sum == 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int returnToBoundaryCount(vector<int> &nums) {
        int ans = 0, sum = 0;
        for (int x : nums) {
            sum += x;
            ans += sum == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func returnToBoundaryCount(nums []int) (ans int) {
	sum := 0
	for _, x := range nums {
		sum += x
		if sum == 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
