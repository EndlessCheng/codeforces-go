下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

### 提示 1

要让答案最大，首先应当最大化答案的二进制的长度。

### 提示 2

把「乘 $2$」分配给多个数，不如只分配给一个数，这样更有可能得到更大的答案。

### 提示 3

枚举把哪个 $\textit{nums}[i]$ 乘 $k$ 次 $2$（左移 $k$ 次）。

代码实现时，可以仿照 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)，预处理每个 $\textit{nums}[i]$ 左侧元素的或值 $\textit{pre}$，以及右侧元素的或值 $\textit{suf}$，从而快速计算出把 $\textit{nums}[i]$ 乘 $k$ 次 $2$ 后的所有元素的或值。进一步地，只需要预处理右侧元素的或值，左侧的或值可以一边枚举一边计算。

```py [sol1-Python3]
class Solution:
    def maximumOr(self, nums: List[int], k: int) -> int:
        n = len(nums)
        suf = [0] * (n + 1)
        for i in range(n - 1, 0, -1):
            suf[i] = suf[i + 1] | nums[i]
        ans = pre = 0
        for i, x in enumerate(nums):
            ans = max(ans, pre | (x << k) | suf[i + 1])
            pre |= x
        return ans
```

```java [sol1-Java]
class Solution {
    public long maximumOr(int[] nums, int k) {
        int n = nums.length;
        var suf = new int[n + 1];
        for (int i = n - 1; i > 0; i--)
            suf[i] = suf[i + 1] | nums[i];
        long ans = 0;
        for (int i = 0, pre = 0; i < n; i++) {
            ans = Math.max(ans, pre | ((long) nums[i] << k) | suf[i + 1]);
            pre |= nums[i];
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long maximumOr(vector<int> &nums, int k) {
        int n = nums.size(), suf[n + 1];
        suf[n] = 0;
        for (int i = n - 1; i; i--)
            suf[i] = suf[i + 1] | nums[i];
        long long ans = 0;
        for (int i = 0, pre = 0; i < n; i++) {
            ans = max(ans, pre | ((long long) nums[i] << k) | suf[i + 1]);
            pre |= nums[i];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i := n - 1; i > 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}
	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i+1])
		pre |= x
	}
	return int64(ans)
}

func max(a, b int) int { if a < b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 思考题

把 $2$ 换成其它数，方法是否一样？

把乘法改成除法，把 OR 改成 AND，要怎么做？
