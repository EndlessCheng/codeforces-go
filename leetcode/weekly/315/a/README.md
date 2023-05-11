### 视频讲解

见[【周赛 315】](https://www.bilibili.com/video/BV1Ae4y1i7PM/)。

### 思路

用一个哈希表记录出现过的数字。一边遍历，一边看 $-\textit{nums}[i]$ 是否在哈希表中，如果在，就更新答案的最大值为 $|\textit{nums}[i]|$。

```py [sol1-Python3]
class Solution:
    def findMaxK(self, nums: List[int]) -> int:
        ans = -1
        s = set()
        for x in nums:
            if -x in s:
                ans = max(ans, abs(x))
            s.add(x)
        return ans
```

```java [sol1-Java]
class Solution {
    public int findMaxK(int[] nums) {
        int ans = -1;
        var s = new HashSet<Integer>();
        for (int x : nums) {
            if (s.contains(-x))
                ans = Math.max(ans, Math.abs(x));
            s.add(x);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int findMaxK(vector<int> &nums) {
        int ans = -1;
        unordered_set<int> s;
        for (int x: nums) {
            if (s.count(-x))
                ans = max(ans, abs(x));
            s.insert(x);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func findMaxK(nums []int) int {
	ans := -1
	has := map[int]bool{}
	for _, x := range nums {
		if abs(x) > ans && has[-x] {
			ans = abs(x)
		}
		has[x] = true
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

---

[往期每日一题题解](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
