请看 [视频讲解](https://www.bilibili.com/video/BV12w411B7ia/)。

遇到这种三元组的题目，一个通常的做法是枚举中间的数。

知道了 $\textit{nums}[j]$，只需要知道 $j$ 左边的最小值和右边的最小值，就知道了三元组的和的最小值。

左右的最小值可以递推算出来。定义 $\textit{suf}[i]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 的最小值（后缀最小值），则有

$$
\textit{suf}[i] = \min(\textit{suf}[i+1], \textit{nums}[i])
$$

前缀最小值 $\textit{pre}$ 的计算方式同理，可以和答案一起算，所以只需要一个变量。

那么答案就是

$$
\textit{pre} + \textit{nums}[j] + \textit{suf}[j+1]
$$

的最小值。

```py [sol-Python3]
class Solution:
    def minimumSum(self, nums: List[int]) -> int:
        n = len(nums)
        suf = [0] * n
        suf[-1] = nums[-1]
        for i in range(n - 2, 1, -1):
            suf[i] = min(suf[i + 1], nums[i])

        ans = inf
        pre = nums[0]
        for j in range(1, n - 1):
            if pre < nums[j] > suf[j + 1]:
                ans = min(ans, pre + nums[j] + suf[j + 1])
            pre = min(pre, nums[j])
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSum(int[] nums) {
        int n = nums.length;
        int[] suf = new int[n];
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = Math.min(suf[i + 1], nums[i]);
        }

        int ans = Integer.MAX_VALUE;
        int pre = nums[0];
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) {
                ans = Math.min(ans, pre + nums[j] + suf[j + 1]);
            }
            pre = Math.min(pre, nums[j]);
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSum(vector<int>& nums) {
        int n = nums.size();
        vector<int> suf(n);
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = min(suf[i + 1], nums[i]);
        }

        int ans = INT_MAX;
        int pre = nums[0];
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) {
                ans = min(ans, pre + nums[j] + suf[j + 1]);
            }
            pre = min(pre, nums[j]);
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0]
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] {
			ans = min(ans, pre+nums[j]+suf[j+1])
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 练习：前后缀分解（右边数字为题目难度）

- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)
- [2906. 构造乘积矩阵](https://leetcode.cn/problems/construct-product-matrix/)
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)
