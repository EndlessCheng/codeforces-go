设 $s = \textit{nums}[0] + \textit{nums}[1]$。答案初始化为 $1$。

从 $i=3$ 开始，如果 $\textit{nums}[i-1] + \textit{nums}[i] = s$ 就继续，把 $i$ 增加 $2$，答案增加 $1$；否则退出循环。

```py [sol-Python3]
class Solution:
    def maxOperations(self, nums: List[int]) -> int:
        s = nums[0] + nums[1]
        for i in range(3, len(nums), 2):
            if nums[i - 1] + nums[i] != s:
                return i // 2
        return len(nums) // 2
```

```java [sol-Java]
class Solution {
    public int maxOperations(int[] nums) {
        int s = nums[0] + nums[1];
        int ans = 1;
        for (int i = 3; i < nums.length && nums[i - 1] + nums[i] == s; i += 2) {
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxOperations(vector<int> &nums) {
        int s = nums[0] + nums[1];
        int ans = 1;
        for (int i = 3; i < nums.size() && nums[i - 1] + nums[i] == s; i += 2) {
            ans++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxOperations(nums []int) int {
	s := nums[0] + nums[1]
	ans := 1
	for i := 3; i < len(nums) && nums[i-1]+nums[i] == s; i += 2 {
		ans++
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
