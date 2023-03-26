### 本题视频讲解

见[【周赛 338】](https://www.bilibili.com/video/BV11o4y1p7Ci/)。

### 前置知识：前缀和

对于数组 $\textit{nums}$，定义它的前缀和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \sum\limits_{j=0}^{i}\textit{nums}[j]$。

根据这个定义，有 $s[i+1]=s[i]+\textit{nums}[i]$。

例如 $\textit{nums}=[1,2,-1,2]$，对应的前缀和数组为 $s=[0,1,3,2,4]$。

通过前缀和，我们可以把**子数组的元素和转换成两个前缀和的差**，即

$$
\sum_{j=\textit{left}}^{\textit{right}}\textit{nums}[j] = \sum\limits_{j=0}^{\textit{right}}\textit{nums}[j] - \sum\limits_{j=0}^{\textit{left}-1}\textit{nums}[j] = \textit{s}[\textit{right}+1] - \textit{s}[\textit{left}]
$$

例如 $\textit{nums}$ 的子数组 $[2,-1,2]$ 的和就可以用 $s[4]-s[1]=4-1=3$ 算出来。

> 注：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示从 $\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}-1]$ 的子数组，此时子数组的和为 $\textit{s}[\textit{right}] - \textit{s}[\textit{left}]$，子数组的长度为 $\textit{right}-\textit{left}$。
>
> 注 2：$s[0]=0$ 表示一个空数组的元素和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $\textit{nums}[0]$ 开始），你要用 $s[\textit{right}]$ 减去谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀和的差。

### 前置知识：二分查找

见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 思路

![t3.png](https://pic.leetcode.cn/1679808210-FVsAou-t3.png)

```py [sol1-Python3]
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

```java [sol1-Java]
class Solution {
    public List<Long> minOperations(int[] nums, int[] queries) {
        Arrays.sort(nums);
        int n = nums.length;
        var sum = new long[n + 1]; // 前缀和
        for (int i = 0; i < n; ++i)
            sum[i + 1] = sum[i] + nums[i];

        var ans = new ArrayList<Long>(queries.length);
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
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<long long> minOperations(vector<int> &nums, vector<int> &queries) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        long long sum[n + 1]; // 前缀和
        sum[0] = 0;
        for (int i = 0; i < n; ++i)
            sum[i + 1] = sum[i] + nums[i];

        int m = queries.size();
        vector<long long> ans(m);
        for (int i = 0; i < m; ++i) {
            int q = queries[i];
            long long j = lower_bound(nums.begin(), nums.end(), q) - nums.begin();
            long long left = q * j - sum[j]; // 蓝色面积
            long long right = sum[n] - sum[j] - q * (n - j); // 绿色面积
            ans[i] = left + right;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func minOperations(nums, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
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

### 复杂度分析

- 时间复杂度：$O((n+q)\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n)$。返回值不计入。
