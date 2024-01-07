[本题视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)

1. 求出最长前缀，同时维护最长前缀的元素和 $\textit{sum}$。
2. 不断增加 $\textit{sum}$，直到 $\textit{sum}$ 不在 $\textit{nums}$ 中。
3. 返回 $\textit{sum}$。

```py [sol-Python3]
class Solution:
    def missingInteger(self, nums: List[int]) -> int:
        s = nums[0]
        for x, y in pairwise(nums):
            if x + 1 != y: break
            s += y

        st = set(nums)
        while s in st:  # 至多循环 n 次，例如 1324567
            s += 1
        return s
```

```java [sol-Java]
class Solution {
    public int missingInteger(int[] nums) {
        int sum = nums[0];
        for (int i = 1; i < nums.length && nums[i] == nums[i - 1] + 1; i++) {
            sum += nums[i];
        }

        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }
        while (set.contains(sum)) { // 至多循环 n 次，例如 1324567
            sum++;
        }
        return sum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int missingInteger(vector<int> &nums) {
        int sum = nums[0];
        for (int i = 1; i < nums.size() && nums[i] == nums[i - 1] + 1; i++) {
            sum += nums[i];
        }

        unordered_set<int> s(nums.begin(), nums.end());
        while (s.contains(sum)) { // 至多循环 n 次，例如 1324567
            sum++;
        }
        return sum;
    }
};
```

```go [sol-Go]
func missingInteger(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		sum += nums[i]
	}

	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}
	for has[sum] { // 至多循环 n 次，例如 1324567
		sum++
	}
	return sum
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
