## 方法一：二分查找

没有逆序对意味着子数组是**递增**的（允许元素相等）。

把 $\textit{nums}$ 分成若干递增段。

比如 $\textit{nums}=[2,3,3,1,3,2,1]$，分成四个递增段 $[2,3,3] + [1,3] + [2] + [1]$。

分类讨论：

- 如果询问的区间 $[l,r]$ 落在某个递增段的内部，那么 $[l,r]$ 内的任意子数组都是递增的。设 $m=r-l+1$，一共有 $1+2+\cdots+m=\dfrac{m(m+1)}{2}$ 个递增子数组。**注**：计算方式为，有 $1$ 个长为 $m$ 的，$2$ 个长为 $m-1$ 的，$m$ 个长为 $1$ 的。
- 如果 $l$ 和 $r$ 在不同的递增段，那么可以把 $[l,r]$ 分成三部分：
  - $[l,\textit{left}_1)$，其中 $\textit{left}_1$ 是在 $l$ 右边的下一个递增段的左端点。设 $m=\textit{left}_1-l$，一共有 $\dfrac{m(m+1)}{2}$ 个递增子数组。
  - $[\textit{left}_1, \textit{left}_2)$，其中 $\textit{left}_2$ 是包含 $r$ 的递增段的左端点。我们可以计算 $[\textit{left}_1, \textit{left}_2)$ 中的每个递增段的子数组个数，预处理子数组个数的**前缀和**数组 $s$，就可以 $\mathcal{O}(1)$ 算出 $[\textit{left}_1, \textit{left}_2)$ 中的递增子数组个数。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。
  - $[\textit{left}_2,r]$。设 $m_2=r-\textit{left}_2+1$，一共有 $\dfrac{m_2(m_2+1)}{2}$ 个递增子数组。

[本题视频讲解](https://www.bilibili.com/video/BV1ZuCQBJEjD/?t=17m55s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countStableSubarrays(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        # 找递增段
        left = []  # 递增段的左端点
        s = [0]  # 递增子数组个数的前缀和
        start = 0
        for i, x in enumerate(nums):
            if i == n - 1 or x > nums[i + 1]:
                # 找到了一个递增段 [start, i]
                left.append(start)
                m = i - start + 1
                # 长为 m 的子数组中有 m*(m+1)/2 个递增子数组
                # 计算 m*(m+1)/2 的前缀和
                s.append(s[-1] + m * (m + 1) // 2)
                start = i + 1  # 下一个递增段的左端点

        ans = []
        for l, r in queries:
            i = bisect_right(left, l)  # 左端点严格大于 l 的第一个区间
            j = bisect_right(left, r) - 1  # 包含 r 的最后一个区间

            # l 和 r 在同一个区间
            if i > j:
                m = r - l + 1
                ans.append(m * (m + 1) // 2)
                continue

            # l 和 r 在不同区间
            # 分成三段 [l, left[i]) + [left[i], left[j]) + [left[j], r]
            # 中间那段的子数组个数用前缀和计算
            m = left[i] - l
            m2 = r - left[j] + 1
            ans.append(m * (m + 1) // 2 + (s[j] - s[i]) + m2 * (m2 + 1) // 2)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] countStableSubarrays(int[] nums, int[][] queries) {
        int n = nums.length;
        // 找递增段
        List<Integer> left = new ArrayList<>(); // 递增段的左端点
        List<Long> s = new ArrayList<>(); // 递增子数组个数的前缀和
        s.add(0L);
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i == n - 1 || x > nums[i + 1]) {
                // 找到了一个递增段 [start, i]
                left.add(start);
                long m = i - start + 1;
                // 长为 m 的子数组中有 m*(m+1)/2 个递增子数组
                // 计算 m*(m+1)/2 的前缀和
                s.add(s.getLast() + m * (m + 1) / 2);
                start = i + 1; // 下一个递增段的左端点
            }
        }

        int qn = queries.length;
        long[] ans = new long[qn];
        for (int k = 0; k < qn; k++) {
            int l = queries[k][0];
            int r = queries[k][1];

            int i = upperBound(left, l); // 左端点严格大于 l 的第一个区间
            int j = upperBound(left, r) - 1; // 包含 r 的最后一个区间

            // l 和 r 在同一个区间
            if (i > j) {
                long m = r - l + 1;
                ans[k] = m * (m + 1) / 2;
                continue;
            }

            // l 和 r 在不同区间
            // 分成三段 [l, left[i]) + [left[i], left[j]) + [left[j], r]
            // 中间那段的子数组个数用前缀和计算
            long m = left.get(i) - l;
            long m2 = r - left.get(j) + 1;
            ans[k] = m * (m + 1) / 2 + (s.get(j) - s.get(i)) + m2 * (m2 + 1) / 2;
        }
        return ans;
    }

    private int upperBound(List<Integer> a, int x) {
        int l = -1;
        int r = a.size();
        while (l + 1 < r) {
            int m = l + (r - l) / 2;
            if (a.get(m) > x) {
                r = m;
            } else {
                l = m;
            }
        }
        return r;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> countStableSubarrays(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        // 找递增段
        vector<int> left; // 递增段的左端点
        vector<long long> s = {0}; // 递增子数组个数的前缀和
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i == n - 1 || x > nums[i + 1]) {
                // 找到了一个递增段 [start, i]
                left.push_back(start);
                long long m = i - start + 1;
                // 长为 m 的子数组中有 m*(m+1)/2 个递增子数组
                // 计算 m*(m+1)/2 的前缀和
                s.push_back(s.back() + m * (m + 1) / 2);
                start = i + 1; // 下一个递增段的左端点
            }
        }

        vector<long long> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            int l = q[0], r = q[1];
            int i = ranges::upper_bound(left, l) - left.begin(); // 左端点严格大于 l 的第一个区间
            int j = ranges::upper_bound(left, r) - left.begin() - 1; // 包含 r 的最后一个区间

            // l 和 r 在同一个区间
            if (i > j) {
                long long m = r - l + 1;
                ans.push_back(m * (m + 1) / 2);
                continue;
            }

            // l 和 r 在不同区间
            // 分成三段 [l, left[i]) + [left[i], left[j]) + [left[j], r]
            // 中间那段的子数组个数用前缀和计算
            long long m = left[i] - l;
            long long m2 = r - left[j] + 1;
            ans.push_back(m * (m + 1) / 2 + (s[j] - s[i]) + m2 * (m2 + 1) / 2);
        }
        return ans;
    }
};
```

```go [sol-Go]
func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)
	// 找递增段
	left := []int{}   // 递增段的左端点
	sum := []int64{0} // 递增子数组个数的前缀和
	start := 0
	for i, x := range nums {
		if i == n-1 || x > nums[i+1] {
			// 找到了一个递增段 [start, i]
			left = append(left, start)
			m := int64(i - start + 1)
			// 长为 m 的子数组中有 m*(m+1)/2 个递增子数组
			// 计算 m*(m+1)/2 的前缀和
			sum = append(sum, sum[len(sum)-1]+m*(m+1)/2)
			start = i + 1 // 下一个递增段的左端点
		}
	}

	ans := make([]int64, len(queries))
	for k, q := range queries {
		l, r := q[0], q[1]
		i := sort.SearchInts(left, l+1)     // 左端点严格大于 l 的第一个区间
		j := sort.SearchInts(left, r+1) - 1 // 包含 r 的最后一个区间

		// l 和 r 在同一个区间
		if i > j {
			m := int64(r - l + 1)
			ans[k] = m * (m + 1) / 2
			continue
		}

		// l 和 r 在不同区间
		// 分成三段 [l, left[i]) + [left[i], left[j]) + [left[j], r]
		// 中间那段的子数组个数用前缀和计算
		m := int64(left[i] - l)
		m2 := int64(r - left[j] + 1)
		ans[k] = m*(m+1)/2 + (sum[j] - sum[i]) + m2*(m2+1)/2
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：记录下一个递增段的左端点

定义 $\textit{nxt}[i]$ 表示在 $i$ 右边的下一个递增段的左端点（不存在则为 $n$）。代替方法一的二分查找。

分类讨论：

- 如果 $\textit{nums}[i] > \textit{nums}[i+1]$，根据定义，$\textit{nxt}[i] = i+1$。
- 如果 $\textit{nums}[i] \le \textit{nums}[i+1]$，那么在 $i$ 右边的下一个递增段的左端点也是在 $i+1$ 右边的下一个递增段的左端点，所以 $\textit{nxt}[i] = \textit{nxt}[i+1]$。这也意味着我们需要倒着计算。

```py [sol-Python3]
class Solution:
    def countStableSubarrays(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        # 计算递增子数组个数的前缀和
        s = [0] * (n + 1)
        cnt = 0
        for i, x in enumerate(nums):
            if i > 0 and x < nums[i - 1]:
                cnt = 0
            cnt += 1
            # 现在 cnt 表示以 i 为右端点的递增子数组个数
            s[i + 1] = s[i] + cnt

        # nxt[i] 表示 i 右边下一个递增段的左端点，若不存在则为 n
        nxt = [0] * n
        nxt[-1] = n
        for i in range(n - 2, -1, -1):
            nxt[i] = nxt[i + 1] if nums[i] <= nums[i + 1] else i + 1

        ans = []
        for l, r in queries:
            l2 = nxt[l]
            if l2 > r:  # l 和 r 在同一个区间
                m = r - l + 1
                ans.append(m * (m + 1) // 2)
            else:  # l 和 r 在不同区间
                # 分成 [l, l2) + [l2, r]
                # 由于 [l2, r] 中的每个右端点所在递增段的左端点都在 [l2, r] 内，所以可以用前缀和计算
                m = l2 - l
                ans.append(m * (m + 1) // 2 + s[r + 1] - s[l2])
        return ans
```

```java [sol-Java]
class Solution {
    public long[] countStableSubarrays(int[] nums, int[][] queries) {
        int n = nums.length;
        // 计算递增子数组个数的前缀和
        long[] sum = new long[n + 1];
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i > 0 && x < nums[i - 1]) {
                cnt = 0;
            }
            cnt++;
            // 现在 cnt 表示以 i 为右端点的递增子数组个数
            sum[i + 1] = sum[i] + cnt;
        }

        // nxt[i] 表示 i 右边下一个递增段的左端点，若不存在则为 n
        int[] nxt = new int[n];
        nxt[n - 1] = n;
        for (int i = n - 2; i >= 0; i--) {
            nxt[i] = nums[i] <= nums[i + 1] ? nxt[i + 1] : i + 1;
        }

        int qn = queries.length;
        long[] ans = new long[qn];
        for (int k = 0; k < qn; k++) {
            int l = queries[k][0];
            int r = queries[k][1];

            int l2 = nxt[l];
            if (l2 > r) { // l 和 r 在同一个区间
                long m = r - l + 1;
                ans[k] = m * (m + 1) / 2;
            } else { // l 和 r 在不同区间
                // 分成 [l, l2) + [l2, r]
                // 由于 [l2, r] 中的每个右端点所在递增段的左端点都在 [l2, r] 内，所以可以用前缀和计算
                long m = l2 - l;
                ans[k] = m * (m + 1) / 2 + sum[r + 1] - sum[l2];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> countStableSubarrays(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        // 计算递增子数组个数的前缀和
        vector<long long> sum(n + 1);
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i > 0 && x < nums[i - 1]) {
                cnt = 0;
            }
            cnt++;
            // 现在 cnt 表示以 i 为右端点的递增子数组个数
            sum[i + 1] = sum[i] + cnt;
        }

        // nxt[i] 表示 i 右边下一个递增段的左端点，若不存在则为 n
        vector<int> nxt(n);
        nxt[n - 1] = n;
        for (int i = n - 2; i >= 0; i--) {
            nxt[i] = nums[i] <= nums[i + 1] ? nxt[i + 1] : i + 1;
        }

        vector<long long> ans;
        ans.reserve(queries.size()); // 预分配空间
        for (auto& q : queries) {
            int l = q[0], r = q[1];
            int l2 = nxt[l];
            if (l2 > r) { // l 和 r 在同一个区间
                long long m = r - l + 1;
                ans.push_back(m * (m + 1) / 2);
            } else { // l 和 r 在不同区间
                // 分成 [l, l2) + [l2, r]
                // 由于 [l2, r] 中的每个右端点所在递增段的左端点都在 [l2, r] 内，所以可以用前缀和计算
                long long m = l2 - l;
                ans.push_back(m * (m + 1) / 2 + sum[r + 1] - sum[l2]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)
	// 计算递增子数组个数的前缀和
	sum := make([]int64, n+1)
	cnt := 0
	for i, x := range nums {
		if i > 0 && x < nums[i-1] {
			cnt = 0
		}
		cnt++
		// 现在 cnt 表示以 i 为右端点的递增子数组个数
		sum[i+1] = sum[i] + int64(cnt)
	}

	// nxt[i] 表示 i 右边下一个递增段的左端点，若不存在则为 n
	nxt := make([]int, n)
	nxt[n-1] = n
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			nxt[i] = nxt[i+1]
		} else {
			nxt[i] = i + 1
		}
	}

	ans := make([]int64, len(queries))
	for k, q := range queries {
		l, r := q[0], q[1]
		l2 := nxt[l]
		if l2 > r { // l 和 r 在同一个区间
			m := int64(r - l + 1)
			ans[k] = m * (m + 1) / 2
		} else { // l 和 r 在不同区间
			// 分成 [l, l2) + [l2, r]
			// 由于 [l2, r] 中的每个右端点所在递增段的左端点都在 [l2, r] 内，所以可以用前缀和计算
			m := int64(l2 - l)
			ans[k] = m*(m+1)/2 + sum[r+1] - sum[l2]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 进阶问题

额外输入一个 $\textit{queries}$ 数组，每个询问包含三个数：

- $[1,i,x]$ 表示修改操作，把 $\textit{nums}[i]$ 改成 $x$。
- $[2,l,r]$ 表示查询操作，询问 $[l,r]$ 中的递增子数组个数。

欢迎在评论区分享你的思路/代码。

## 相似题目

[3501. 操作后最大活跃区段数 II](https://leetcode.cn/problems/maximize-active-section-with-trade-ii/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
