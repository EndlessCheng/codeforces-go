## 视频讲解

见[【双周赛 91】](https://www.bilibili.com/video/BV1gd4y1b7qj) 第一题。

## 思路

排序后，每次取出的最小和最大的数就是 $\textit{nums}[i]$ 和 $\textit{nums}[n-1-i]$。

把这两个数的和放入哈希表中（不需要除以 $2$，因为只计算不同平均值的个数，两个平均值不同，等价于两数之和不同）。

```py [sol-Python3]
class Solution:
    def distinctAverages(self, nums: List[int]) -> int:
        nums.sort()
        return len(set(nums[i] + nums[-i - 1] for i in range(len(nums) // 2)))
```

```java [sol-Java]
class Solution {
    public int distinctAverages(int[] nums) {
        Arrays.sort(nums);
        var set = new HashSet<Integer>();
        for (int i = 0, n = nums.length; i < n / 2; i++)
            set.add(nums[i] + nums[n - 1 - i]);
        return set.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int distinctAverages(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        unordered_set<int> s;
        for (int i = 0, n = nums.size(); i < n / 2; i++)
            s.insert(nums[i] + nums[n - 1 - i]);
        return s.size();
    }
};
```

```go [sol-Go]
func distinctAverages(nums []int) int {
	sort.Ints(nums)
	set := map[int]struct{}{}
	for i, n := 0, len(nums); i < n/2; i++ {
		set[nums[i]+nums[n-1-i]] = struct{}{}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
