统计相邻数字的和，加入哈希表中。一共有 $n-1$ 个相邻数字的和。如果去重后，个数不足 $n-1$ 个，则说明和相等的子数组存在。

```py [sol1-Python3]
class Solution:
    def findSubarrays(self, nums: List[int]) -> bool:
        return len(set(map(sum, pairwise(nums)))) < len(nums) - 1
```

也可以在遍历 $\textit{nums}$ 的过程中去判断是否有相同的和。

```java [sol1-Java]
class Solution {
    public boolean findSubarrays(int[] nums) {
        var set = new HashSet<Integer>();
        for (int i = 1; i < nums.length; ++i)
            if (!set.add(nums[i - 1] + nums[i]))
                return true;
        return false;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool findSubarrays(vector<int> &nums) {
        unordered_set<int> s;
        for (int i = 1; i < nums.size(); ++i)
            if (!s.insert(nums[i - 1] + nums[i]).second)
                return true;
        return false;
    }
};
```

```go [sol1-Go]
func findSubarrays(nums []int) bool {
	vis := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		s := nums[i-1] + nums[i]
		if vis[s] {
			return true
		}
		vis[s] = true
	}
	return false
}
```

### 思考题

1. 如果把子数组的长度改为一个比较大的数字 $k$ 要怎么做？

2. 如果把子数组改成子序列要怎么做？

---

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~

附：[每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)。
