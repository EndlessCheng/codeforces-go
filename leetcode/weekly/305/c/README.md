**前置知识**：[动态规划入门](https://b23.tv/72onpYq)。

如何想出状态定义？

- 如果 $\textit{nums}$ 的最后两个数相等，那么去掉这两个数，问题变成剩下 $n-2$ 个数能否有效划分。
- 如果 $\textit{nums}$ 的最后三个数相等，那么去掉这三个数，问题变成剩下 $n-3$ 个数能否有效划分。
- 如果 $\textit{nums}$ 的最后三个数是连续递增的，那么去掉这三个数，问题变成剩下 $n-3$ 个数能否有效划分。

我们要解决的问题都形如「$\textit{nums}$ 的前 $i$ 个数能否有效划分」。

于是定义 $f[0] = \texttt{true}$，$f[i+1]$ 表示能否有效划分 $\textit{nums}[0]$ 到 $\textit{nums}[i]$。

根据有效划分的定义，有

$$
f[i+1] = \vee
\begin{cases} 
f[i-1]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1],&i>0\\
f[i-2]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1] = \textit{nums}[i-2],&i>1\\
f[i-2]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1]+1 = \textit{nums}[i-2]+2,&i>1
\end{cases}
$$

答案为 $f[n]$。

```py [sol-Python3]
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

```java [sol-Java]
class Solution {
    public boolean validPartition(int[] nums) {
        int n = nums.length;
        boolean[] f = new boolean[n + 1];
        f[0] = true;
        for (int i = 1; i < n; i++) {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool validPartition(vector<int> &nums) {
        int n = nums.size();
        vector<int> f(n + 1);
        f[0] = true;
        for (int i = 1; i < n; i++) {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
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

```js [sol-JavaScript]
var validPartition = function(nums) {
    const n = nums.length;
    const f = Array(n + 1).fill(false);
    f[0] = true;
    for (let i = 1; i < n; i++) {
        if (f[i - 1] && nums[i] === nums[i - 1] ||
            i > 1 && f[i - 2] && (nums[i] === nums[i - 1] && nums[i] === nums[i - 2] ||
                                  nums[i] === nums[i - 1] + 1 && nums[i] === nums[i - 2] + 2)) {
            f[i + 1] = true;
        }
    }
    return f[n];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn valid_partition(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut f = vec![false; n + 1];
        f[0] = true;
        for i in 1..n {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        f[n]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

有多少种有效划分的**方案数**？模 $10^9+7$。

欢迎在评论区发表你的思路/代码。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
