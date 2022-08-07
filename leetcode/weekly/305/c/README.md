下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

定义 $f[i+1]$ 表示从 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的这些元素能否有效划分。那么 $f[0] = \texttt{true}$，答案为 $f[n]$。

根据题意，有

$$
f[i+1] = \text{OR}
\begin{cases} 
f[i-1]\ \text{AND}\ \textit{nums}[i] = \textit{nums}[i-1],&i>0\\
f[i-2]\ \text{AND}\ \textit{nums}[i] = \textit{nums}[i-1] = \textit{nums}[i-2],&i>1\\
f[i-2]\ \text{AND}\ \textit{nums}[i] = \textit{nums}[i-1]+1 = \textit{nums}[i-2]+2,&i>1
\end{cases}
$$

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def validPartition(self, nums: List[int]) -> bool:
        n = len(nums)
        f = [True] + [False] * n
        for i, x in enumerate(nums):
            if i > 0 and f[i - 1] and x == nums[i - 1] or \
               i > 1 and f[i - 2] and (x == nums[i - 1] == nums[i - 2] or
                                       x == nums[i - 1] + 1 == nums[i - 2] + 2):
               f[i + 1] = True
        return f[n]
```

```java [sol1-Java]
class Solution {
    public boolean validPartition(int[] nums) {
        var n = nums.length;
        var f = new boolean[n + 1];
        f[0] = true;
        for (var i = 1; i < n; ++i)
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2))
                f[i + 1] = true;
        return f[n];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool validPartition(vector<int> &nums) {
        int n = nums.size();
        bool f[n + 1]; memset(f, 0, sizeof(f));
        f[0] = true;
        for (int i = 1; i < n; ++i)
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2))
                f[i + 1] = true;
        return f[n];
    }
};
```

```go [sol1-Go]
func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				                x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
```
