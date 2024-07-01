## 前置知识

[前缀和及其扩展，附题单！](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)

![t3.png](https://pic.leetcode.cn/1679808210-FVsAou-t3.png)

关于二分算法，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)
        nums.sort()
        s = list(accumulate(nums, initial=0))  # 前缀和
        ans = []
        for q in queries:
            j = bisect_left(nums, q)
            left = q * j - s[j]  # 蓝色面积
            right = s[n] - s[j] - q * (n - j)  # 绿色面积
            ans.append(left + right)
        return ans
```

```java [sol-Java]
class Solution {
    public List<Long> minOperations(int[] nums, int[] queries) {
        Arrays.sort(nums);
        int n = nums.length;
        long[] sum = new long[n + 1]; // 前缀和
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        List<Long> ans = new ArrayList<>(queries.length);
        for (int q : queries) {
            int j = lowerBound(nums, q);
            long left = (long) q * j - sum[j]; // 蓝色面积
            long right = sum[n] - sum[j] - (long) q * (n - j); // 绿色面积
            ans.add(left + right);
        }
        return ans;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> minOperations(vector<int>& nums, vector<int>& queries) {
        ranges::sort(nums);
        int n = nums.size();
        vector<long long> sum(n + 1); // 前缀和
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        int m = queries.size();
        vector<long long> ans(m);
        for (int i = 0; i < m; i++) {
            int q = queries[i];
            long long j = ranges::lower_bound(nums, q) - nums.begin();
            long long left = q * j - sum[j]; // 蓝色面积
            long long right = sum[n] - sum[j] - q * (n - j); // 绿色面积
            ans[i] = left + right;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums, queries []int) []int64 {
	n := len(nums)
	slices.Sort(nums)
	sum := make([]int, n+1) // 前缀和
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - sum[j] // 蓝色面积
		right := sum[n] - sum[j] - q*(n-j) // 绿色面积
		ans[i] = int64(left + right)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
