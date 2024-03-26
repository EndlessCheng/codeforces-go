这种三元组的题目，通常是枚举中间的数。

枚举 $\textit{nums}[j]$，我们需要求出 $j$ **左边所有元素的最小值**和**右边所有元素的最小值**。

这可以**递推**计算。定义 $\textit{suf}[i]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 的最小值（后缀最小值），则有

$$
\textit{suf}[i] = \min(\textit{suf}[i+1], \textit{nums}[i])
$$

前缀最小值 $\textit{pre}$ 的计算方式同理，可以和答案一起算，所以只需要一个变量。

那么答案就是

$$
\textit{pre} + \textit{nums}[j] + \textit{suf}[j+1]
$$

的最小值。

附：[视频讲解](https://www.bilibili.com/video/BV12w411B7ia/)。

```py [sol-Python3]
class Solution:
    def minimumSum(self, nums: List[int]) -> int:
        n = len(nums)
        suf = [0] * n
        suf[-1] = nums[-1]  # 后缀最小值
        for i in range(n - 2, 1, -1):
            suf[i] = min(suf[i + 1], nums[i])

        ans = inf
        pre = nums[0]  # 前缀最小值
        for j in range(1, n - 1):
            if pre < nums[j] > suf[j + 1]:  # 山形
                ans = min(ans, pre + nums[j] + suf[j + 1])  # 更新答案
            pre = min(pre, nums[j])
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSum(int[] nums) {
        int n = nums.length;
        int[] suf = new int[n]; // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = Math.min(suf[i + 1], nums[i]);
        }

        int ans = Integer.MAX_VALUE;
        int pre = nums[0]; // 前缀最小值
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
                ans = Math.min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
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
    int minimumSum(vector<int> &nums) {
        int n = nums.size();
        vector<int> suf(n); // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for (int i = n - 2; i > 1; i--) {
            suf[i] = min(suf[i + 1], nums[i]);
        }

        int ans = INT_MAX;
        int pre = nums[0]; // 前缀最小值
        for (int j = 1; j < n - 1; j++) {
            if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
                ans = min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
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
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0] // 前缀最小值
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] { // 山形
			ans = min(ans, pre+nums[j]+suf[j+1]) // 更新答案
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

```js [sol-JavaScript]
var minimumSum = function(nums) {
    const n = nums.length;
    const suf = Array(n); // 后缀最小值
    suf[n - 1] = nums[n - 1];
    for (let i = n - 2; i > 1; i--) {
        suf[i] = Math.min(suf[i + 1], nums[i]);
    }

    let ans = Infinity;
    let pre = nums[0]; // 前缀最小值
    for (let j = 1; j < n - 1; j++) {
        if (pre < nums[j] && nums[j] > suf[j + 1]) { // 山形
            ans = Math.min(ans, pre + nums[j] + suf[j + 1]); // 更新答案
        }
        pre = Math.min(pre, nums[j]);
    }
    return ans === Infinity ? -1 : ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut suf = vec![0; n]; // 后缀最小值
        suf[n - 1] = nums[n - 1];
        for i in (2..n - 1).rev() {
            suf[i] = suf[i + 1].min(nums[i]);
        }

        let mut ans = i32::MAX;
        let mut pre = nums[0]; // 前缀最小值
        for j in 1..n - 1 {
            if pre < nums[j] && nums[j] > suf[j + 1] { // 山形
                ans = ans.min(pre + nums[j] + suf[j + 1]); // 更新答案
            }
            pre = pre.min(nums[j]);
        }
        if ans == i32::MAX { -1 } else { ans }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 题单：前后缀分解（右边数字为题目难度）

> 部分题目也可以用状态机 DP 解决。

- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)（[视频讲解](https://www.bilibili.com/video/BV1Qg411q7ia/?t=3m05s)）
- [123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) 注：拆分成两个 121 题
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395
- [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/) 1423
- [845. 数组中的最长山脉](https://leetcode.cn/problems/longest-mountain-in-array/) 1437
- [2909. 元素和最小的山形三元组 II](https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/) 1479
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495
- [1525. 字符串的好分割数目](https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/) 1500
- [2874. 有序三元组中的最大值 II](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/) 1583
- [1031. 两个非重叠子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/) 1680
- [689. 三个无重叠子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/)
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695
- [2100. 适合野炊的日子](https://leetcode.cn/problems/find-good-days-to-rob-the-bank/) 1702
- [1653. 使字符串平衡的最少删除次数](https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/) 1794
- [926. 将字符串翻转到单调递增](https://leetcode.cn/problems/flip-string-to-monotone-increasing/)
- [1477. 找两个和为目标值且不重叠的子数组](https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/) 1851
- [1671. 得到山形数组的最少删除次数](https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/) 1913
- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) ~2000
- [1888. 使二进制字符串字符交替的最少反转次数](https://leetcode.cn/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/) 2006
- [2906. 构造乘积矩阵](https://leetcode.cn/problems/construct-product-matrix/) 2075
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223
- [2163. 删除元素后和的最小差值](https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements/) 2225
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433
- [3003. 执行操作后的最大分割数量](https://leetcode.cn/problems/maximize-the-number-of-partitions-after-operations/) 3039
- [487. 最大连续 1 的个数 II](https://leetcode.cn/problems/max-consecutive-ones-ii/)（会员题）
- [1746. 经过一次操作后的最大子数组和](https://leetcode.cn/problems/maximum-subarray-sum-after-one-operation/)（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
