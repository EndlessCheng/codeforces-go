请看 [视频讲解](https://www.bilibili.com/video/BV1wh4y1Q7XW/)。

用一个长为 $10$ 的数组 $\textit{maxVal}[i]$ 维护最大数位为 $i$ 的元素的最大值。

当我们遍历到 $\textit{nums}[i]$ 时，设其最大数位为 $\textit{maxD}$，那么用

$$
\textit{nums}[i] + \textit{maxVal}[\textit{maxD}]
$$

更新答案。

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int]) -> int:
        ans = -1
        max_val = [-inf] * 10
        for v in nums:
            max_d = max(map(int, str(v)))
            ans = max(ans, v + max_val[max_d])
            max_val[max_d] = max(max_val[max_d], v)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxSum(int[] nums) {
        int ans = -1;
        var maxVal = new int[10];
        Arrays.fill(maxVal, Integer.MIN_VALUE);
        for (int v : nums) {
            int maxD = 0;
            for (int x = v; x > 0; x /= 10)
                maxD = Math.max(maxD, x % 10);
            ans = Math.max(ans, v + maxVal[maxD]);
            maxVal[maxD] = Math.max(maxVal[maxD], v);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int> &nums) {
        int ans = -1;
        vector<int> max_val(10, INT_MIN);
        for (int v: nums) {
            int max_d = 0;
            for (int x = v; x; x /= 10)
                max_d = max(max_d, x % 10);
            ans = max(ans, v + max_val[max_d]);
            max_val[max_d] = max(max_val[max_d], v);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int) int {
	ans := -1
	maxVal := [10]int{}
	for i := range maxVal {
		maxVal[i] = math.MinInt // 表示不存在最大值
	}
	for _, v := range nums {
		maxD := 0
		for x := v; x > 0; x /= 10 {
			maxD = max(maxD, x%10)
		}
		ans = max(ans, v+maxVal[maxD])
		maxVal[maxD] = max(maxVal[maxD], v)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var maxSum = function (nums) {
    let ans = -1;
    let maxVal = Array(10).fill(Number.MIN_SAFE_INTEGER);
    for (const v of nums) {
        let maxD = 0;
        for (let x = v; x; x = Math.floor(x / 10))
            maxD = Math.max(maxD, x % 10);
        ans = Math.max(ans, v + maxVal[maxD]);
        maxVal[maxD] = Math.max(maxVal[maxD], v);
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
