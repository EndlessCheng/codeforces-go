### 本题视频讲解

见[【周赛 344】](https://www.bilibili.com/video/BV1YL41187Rx/)，欢迎点赞投币！

### 思路

先倒序枚举 $\textit{nums}$，并记录每个后缀的不同元素个数到 $\textit{suf}$ 数组中，这可以用哈希表做到。

然后正序枚举 $\textit{nums}[i]$，同样地，记录每个前缀的不同元素个数，减去 $\textit{suf}[i+1]$，即为答案。

```py [sol1-Python3]
class Solution:
    def distinctDifferenceArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        suf = [0] * (n + 1)
        s = set()
        for i in range(n - 1, 0, -1):
            s.add(nums[i])
            suf[i] = len(s)

        s.clear()
        ans = [0] * n
        for i, x in enumerate(nums):
            s.add(x)
            ans[i] = len(s) - suf[i + 1]
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] distinctDifferenceArray(int[] nums) {
        int n = nums.length;
        var suf = new int[n + 1];
        var s = new HashSet<Integer>();
        for (int i = n - 1; i > 0; i--) {
            s.add(nums[i]);
            suf[i] = s.size();
        }

        s.clear();
        var ans = new int[n];
        for (int i = 0; i < n; i++) {
            s.add(nums[i]);
            ans[i] = s.size() - suf[i + 1];
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> distinctDifferenceArray(vector<int> &nums) {
        int n = nums.size(), suf[n + 1]; suf[n] = 0;
        unordered_set<int> s;
        for (int i = n - 1; i; i--) {
            s.insert(nums[i]);
            suf[i] = s.size();
        }

        s.clear();
        vector<int> ans(n);
        for (int i = 0; i < n; i++) {
            s.insert(nums[i]);
            ans[i] = s.size() - suf[i + 1];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n+1)
	set := map[int]struct{}{}
	for i := n - 1; i >= 0; i-- {
		set[nums[i]] = struct{}{}
		suf[i] = len(set)
	}

	set = make(map[int]struct{}, len(set))
	ans := make([]int, n)
	for i, x := range nums {
		set[x] = struct{}{}
		ans[i] = len(set) - suf[i+1]
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
