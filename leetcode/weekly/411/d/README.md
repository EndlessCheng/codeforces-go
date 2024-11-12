## 方法一：滑动窗口+前缀和+二分查找

**核心思路**：对于每个询问，计算以 $l$ 为右端点的合法子串个数，以 $l+1$ 为右端点的合法子串个数，……，以 $r$ 为右端点的合法子串个数。

我们需要知道以 $i$ 为右端点的合法子串，其左端点最小是多少。

由于随着 $i$ 的变大，窗口内的字符数量变多，越不能满足题目要求，所以最小左端点会随着 $i$ 的增大而增大，有**单调性**，因此可以用 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 计算。

设以 $i$ 为右端点的合法子串，其左端点**最小**是 $\textit{left}[i]$。

那么以 $i$ 为右端点的合法子串，其左端点可以是 $\textit{left}[i],\textit{left}[i]+1, \ldots, i$，一共 

$$
i-\textit{left}[i]+1
$$ 

个。

回答询问时，分类讨论：

- 如果 $\textit{left}[r] \le l$，说明 $[l,r]$ 内的所有子串都是合法的，这一共有 $1+2+\cdots + (r-l+1) = \dfrac{(r-l+2)(r-l+1)}{2}$ 个。
- 否则，由于 $\textit{left}$ 是**有序数组**，我们可以在 $[l,r]$ 中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) $\textit{left}$ 中的第一个满足 $\textit{left}[j]\ge l$ 的下标 $j$，那么：
   - 由于 $\textit{left}[j-1] < l$，所以 $[l,j-1]$ 内的所有子串都是合法的，这一共有 $1+2+\cdots + (j-l) = \dfrac{(j-l+1)(j-l)}{2}$ 个。
   - 右端点在 $[j,r]$ 内的子串，可以累加下标在 $[j,r]$ 内的所有 $i-\textit{left}[i]+1$ 的和。这可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 预处理。
   - 上述两项累加，即为答案。

代码实现时，两种情况可以合并为一种。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hH4y1c7T5/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def countKConstraintSubstrings(self, s: str, k: int, queries: List[List[int]]) -> List[int]:
        n = len(s)
        left = [0] * n
        pre = [0] * (n + 1)
        cnt = [0, 0]
        l = 0
        for i, c in enumerate(s):
            cnt[ord(c) & 1] += 1
            while cnt[0] > k and cnt[1] > k:
                cnt[ord(s[l]) & 1] -= 1
                l += 1
            left[i] = l  # 记录合法子串右端点 i 对应的最小左端点 l
            # 计算 i-left[i]+1 的前缀和
            pre[i + 1] = pre[i] + i - l + 1

        ans = []
        for l, r in queries:
            j = bisect_left(left, l, l, r + 1)  # 如果区间内所有数都小于 l，结果是 j=r+1
            ans.append(pre[r + 1] - pre[j] + (j - l + 1) * (j - l) // 2)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] countKConstraintSubstrings(String S, int k, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] left = new int[n];
        long[] sum = new long[n + 1];
        int[] cnt = new int[2];
        int l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l++] & 1]--;
            }
            left[i] = l; // 记录合法子串右端点 i 对应的最小左端点 l
            // 计算 i-left[i]+1 的前缀和
            sum[i + 1] = sum[i] + i - l + 1;
        }

        long[] ans = new long[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int ql = queries[i][0];
            int qr = queries[i][1];
            // 如果区间内所有数都小于 ql，结果是 j=qr+1
            int j = lowerBound(left, ql - 1, qr + 1, ql);
            ans[i] = sum[qr + 1] - sum[j] + (long) (j - ql + 1) * (j - ql) / 2;
        }
        return ans;
    }

    // 返回在开区间 (left, right) 中的最小的 j，满足 nums[j] >= target
    // 如果没有这样的数，返回 right
    // 原理见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int left, int right, int target) {
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
    vector<long long> countKConstraintSubstrings(string s, int k, vector<vector<int>>& queries) {
        int n = s.length();
        vector<int> left(n);
        vector<long long> sum(n + 1);
        int cnt[2]{}, l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l++] & 1]--;
            }
            left[i] = l; // 记录合法子串右端点 i 对应的最小左端点 l
            // 计算 i-left[i]+1 的前缀和
            sum[i + 1] = sum[i] + i - l + 1;
        }

        vector<long long> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            int l = queries[i][0], r = queries[i][1];
            // 如果区间内所有数都小于 l，结果是 j=r+1
            int j = lower_bound(left.begin() + l, left.begin() + r + 1, l) - left.begin();
            ans[i] = sum[r + 1] - sum[j] + (long long) (j - l + 1) * (j - l) / 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l // 记录合法子串右端点 i 对应的最小左端点 l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := l + sort.SearchInts(left[l:r+1], l) // 如果区间内所有数都小于 l，结果是 j=r+1
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。注意 $\textit{l}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：预处理

上面的做法，每次都要二分找最小 $j$，满足 $\textit{left}[j]\ge l$。能否不用二分呢？

也可以直接预处理，对每个左端点 $l=0,1,2,\ldots,n-1$，计算出最小的 $j$，满足 $\textit{left}[j]\ge l$。

由于 $\textit{left}$ 数组是有序的，这个过程可以用**双指针**实现。

将计算出的 $j$ 保存到 $\textit{right}[l]$ 中。如果不存在满足 $\textit{left}[j]\ge l$ 的 $j$，则 $\textit{right}[l] = n$。

现在 $\textit{left}[\textit{right}[l]]\ge l$ 且 $\textit{left}[\textit{right}[l]-1] < l$。

预处理后，回答询问时，$j$ 可以直接通过 $\textit{right}[l]$ 获取到。注意这个数不能超过 $r+1$，所以有

$$
j = \min(\textit{right}[l], r+1)
$$

### 写法一

更快的写法见下面的写法二。

```py [sol-Python3]
class Solution:
    def countKConstraintSubstrings(self, s: str, k: int, queries: List[List[int]]) -> List[int]:
        n = len(s)
        left = [0] * n
        pre = [0] * (n + 1)
        cnt = [0, 0]
        l = 0
        for i, c in enumerate(s):
            cnt[ord(c) & 1] += 1
            while cnt[0] > k and cnt[1] > k:
                cnt[ord(s[l]) & 1] -= 1
                l += 1
            left[i] = l
            pre[i + 1] = pre[i] + i - l + 1

        right = [0] * n
        l = 0
        for i in range(n):
            while l < n and left[l] < i:
                l += 1
            right[i] = l

        ans = []
        for l, r in queries:
            j = min(right[l], r + 1)
            ans.append(pre[r + 1] - pre[j] + (j - l + 1) * (j - l) // 2)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] countKConstraintSubstrings(String S, int k, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] left = new int[n];
        long[] sum = new long[n + 1];
        int[] cnt = new int[2];
        int l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l++] & 1]--;
            }
            left[i] = l;
            sum[i + 1] = sum[i] + i - l + 1;
        }

        int[] right = new int[n];
        l = 0;
        for (int i = 0; i < n; i++) {
            while (l < n && left[l] < i) {
                l++;
            }
            right[i] = l;
        }

        long[] ans = new long[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int ql = queries[i][0];
            int qr = queries[i][1];
            int j = Math.min(right[ql], qr + 1);
            ans[i] = sum[qr + 1] - sum[j] + (long) (j - ql + 1) * (j - ql) / 2;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> countKConstraintSubstrings(string s, int k, vector<vector<int>>& queries) {
        int n = s.length();
        vector<int> left(n);
        vector<long long> sum(n + 1);
        int cnt[2]{}, l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l++] & 1]--;
            }
            left[i] = l;
            sum[i + 1] = sum[i] + i - l + 1;
        }

        vector<int> right(n);
        l = 0;
        for (int i = 0; i < n; i++) {
            while (l < n && left[l] < i) {
                l++;
            }
            right[i] = l;
        }

        vector<long long> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            int l = queries[i][0], r = queries[i][1];
            int j = min(right[l], r + 1);
            ans[i] = sum[r + 1] - sum[j] + (long long) (j - l + 1) * (j - l) / 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		sum[i+1] = sum[i] + i - l + 1
	}

	right := make([]int, n)
	l = 0
	for i := range right {
		for l < n && left[l] < i {
			l++
		}
		right[i] = l
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(right[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
```

### 写法二

在上面的方法中，我们找的是最小的 $j$，满足 $\textit{left}[j]\ge l$。

也可以找最小的 $j$，满足 $\textit{left}[j] > l$，这不会影响计算出的合法子串个数。

现在 $\textit{left}[\textit{right}[l]] > l$ 且 $\textit{left}[\textit{right}[l]-1] \le l$。

在滑窗的过程中，如果发现窗口不满足要求，那么在移动左端点 $l$ 之前，可以记录 $\textit{right}[l]=i$。

```py [sol-Python3]
class Solution:
    def countKConstraintSubstrings(self, s: str, k: int, queries: List[List[int]]) -> List[int]:
        n = len(s)
        right = [n] * n
        pre = [0] * (n + 1)
        cnt = [0, 0]
        l = 0
        for i, c in enumerate(s):
            cnt[ord(c) & 1] += 1
            while cnt[0] > k and cnt[1] > k:
                cnt[ord(s[l]) & 1] -= 1
                right[l] = i
                l += 1
            pre[i + 1] = pre[i] + i - l + 1

        ans = []
        for l, r in queries:
            j = min(right[l], r + 1)
            ans.append(pre[r + 1] - pre[j] + (j - l + 1) * (j - l) // 2)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] countKConstraintSubstrings(String S, int k, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] right = new int[n];
        long[] sum = new long[n + 1];
        int[] cnt = new int[2];
        int l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l] & 1]--;
                right[l++] = i;
            }
            sum[i + 1] = sum[i] + i - l + 1;
        }
        // 剩余没填的 right[l] 均为 n
        Arrays.fill(right, l, n, n);

        long[] ans = new long[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int ql = queries[i][0];
            int qr = queries[i][1];
            int j = Math.min(right[ql], qr + 1);
            ans[i] = sum[qr + 1] - sum[j] + (long) (j - ql + 1) * (j - ql) / 2;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> countKConstraintSubstrings(string s, int k, vector<vector<int>>& queries) {
        int n = s.length();
        vector<int> right(n, n);
        vector<long long> sum(n + 1);
        int cnt[2]{}, l = 0;
        for (int i = 0; i < n; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[l] & 1]--;
                right[l++] = i;
            }
            sum[i + 1] = sum[i] + i - l + 1;
        }

        vector<long long> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            int l = queries[i][0], r = queries[i][1];
            int j = min(right[l], r + 1);
            ans[i] = sum[r + 1] - sum[j] + (long long) (j - l + 1) * (j - l) / 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	right := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			right[l] = i
			l++
		}
		sum[i+1] = sum[i] + i - l + 1
	}
	// 剩余没填的 right[l] 均为 n
	for ; l < n; l++ {
		right[l] = n
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(right[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。注意 $\textit{l}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
