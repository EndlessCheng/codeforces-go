[本题视频讲解](https://www.bilibili.com/video/BV1ae411e7fn/)

## 方法一：从「移除」的角度思考

#### 提示 1

设 $\textit{nums}_1$ 中有 $n_1$ 个不同元素，$\textit{nums}_2$ 中有 $n_2$ 个不同元素，它们的交集有 $\textit{common}$ 个元素。

如果不移除任何元素，根据容斥原理，$\textit{nums}_1$ 和 $\textit{nums}_2$ 的并集一共有

$$
\textit{ans} = n_1+n_2-\textit{common}
$$

个不同元素。

#### 提示 2

我们可以先移除每个数组中的重复元素，再考虑从剩下的数中移除元素。

设 $m = n/2$。对于 $\textit{nums}_1$ 来说，如果 $n_1 > m$，**先从交集中移除元素**：

- 如果交集元素少，那么全部移除，即移除 $\textit{common}$ 个元素。
- 如果交集元素多，那么移除交集中的 $n_1-m$ 个元素，就可以让 $n_1=m$。

所以要从交集中移除

$$
\min(n_1-m, \textit{common})
$$

个元素。

移除后，如果 $n_1$ 仍然大于 $m$，那么必须把 $\textit{ans}$ 减少 $n_1-m$。

对于 $\textit{nums}_2$ 也同理。

```py [sol-Python3]
class Solution:
    def maximumSetSize(self, nums1: List[int], nums2: List[int]) -> int:
        set1 = set(nums1)
        set2 = set(nums2)
        common = len(set1 & set2)

        n1 = len(set1)
        n2 = len(set2)
        ans = n1 + n2 - common

        m = len(nums1) // 2
        if n1 > m:
            mn = min(n1 - m, common)
            ans -= n1 - mn - m
            common -= mn

        if n2 > m:
            n2 -= min(n2 - m, common)
            ans -= n2 - m

        return ans
```

```java [sol-Java]
class Solution {
    public int maximumSetSize(int[] nums1, int[] nums2) {
        Set<Integer> set1 = new HashSet<>();
        for (int x : nums1) {
            set1.add(x);
        }
        Set<Integer> set2 = new HashSet<>();
        for (int x : nums2) {
            set2.add(x);
        }
        int common = 0;
        for (int x : set1) {
            if (set2.contains(x)) {
                common++;
            }
        }

        int n1 = set1.size();
        int n2 = set2.size();
        int ans = n1 + n2 - common;

        int m = nums1.length / 2;
        if (n1 > m) {
            int mn = Math.min(n1 - m, common);
            ans -= n1 - mn - m;
            common -= mn;
        }

        if (n2 > m) {
            n2 -= Math.min(n2 - m, common);
            ans -= n2 - m;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSetSize(vector<int> &nums1, vector<int> &nums2) {
        unordered_set<int> set1(nums1.begin(), nums1.end());
        unordered_set<int> set2(nums2.begin(), nums2.end());
        int common = 0;
        for (int x : set1) {
            common += set2.count(x);
        }

        int n1 = set1.size();
        int n2 = set2.size();
        int ans = n1 + n2 - common;

        int m = nums1.size() / 2;
        if (n1 > m) {
            int mn = min(n1 - m, common);
            ans -= n1 - mn - m;
            common -= mn;
        }

        if (n2 > m) {
            n2 -= min(n2 - m, common);
            ans -= n2 - m;
        }

        return ans;
    }
};
```

```go [sol-Go]
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	set2 := map[int]bool{}
	for _, x := range nums2 {
		set2[x] = true
	}
	common := 0
	for x := range set1 {
		if set2[x] {
			common++
		}
	}

	n1 := len(set1)
	n2 := len(set2)
	ans := n1 + n2 - common

	m := len(nums1) / 2
	if n1 > m {
		mn := min(n1-m, common)
		ans -= n1 - mn - m
		common -= mn
	}

	if n2 > m {
		n2 -= min(n2-m, common)
		ans -= n2 - m
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 方法二：从「添加」的角度思考

设 $\textit{nums}_1$ 中有 $n_1$ 个不同元素，$\textit{nums}_2$ 中有 $n_2$ 个不同元素，它们的交集有 $\textit{common}$ 个元素。

考虑怎么从两个数组中选择不同元素，添加到集合 $s$ 中，使 $s$ 的大小最大：

- 对于 $\textit{nums}_1$，优先选**不在**交集中的元素，这可以选 $n_1-\textit{common}$ 个，但不能超过题目规定的 $n/2$，所以至多选 $c_1 = \min(n_1-\textit{common}, n/2)$ 个不在交集中的元素。
- 对于 $\textit{nums}_2$，优先选**不在**交集中的元素，同理，至多选 $c_2 = \min(n_2-\textit{common}, n/2)$ 个。
- 由于都和 $n/2$ 取最小值，所以 $c_1 + c_2 \le n/2 + n/2 = n$。
- 如果 $c_1 + c_2 < n$，那么还可以再选 $n-c_1-c_2$ 个数，且这些数只能从交集中选，所以不能超过 $\textit{common}$ 个，所以还可以再选 $\min(n-c_1-c_2, \textit{common})$ 个数。

最终答案为

$$
\begin{aligned}
&c_1 + c_2 + \min(n-c_1-c_2, \textit{common})\\
=\ &\min(n, c_1 + c_2 + \textit{common})
\end{aligned}
$$

```py [sol-Python3]
class Solution:
    def maximumSetSize(self, nums1: List[int], nums2: List[int]) -> int:
        set1 = set(nums1)
        set2 = set(nums2)
        common = len(set1 & set2)
        n = len(nums1)
        c1 = min(len(set1) - common, n // 2)
        c2 = min(len(set2) - common, n // 2)
        return min(n, c1 + c2 + common)
```

```java [sol-Java]
class Solution {
    public int maximumSetSize(int[] nums1, int[] nums2) {
        Set<Integer> set1 = new HashSet<>();
        for (int x : nums1) {
            set1.add(x);
        }
        Set<Integer> set2 = new HashSet<>();
        int common = 0;
        for (int x : nums2) {
            if (set2.contains(x)) {
                continue;
            }
            set2.add(x);
            // 相比方法一，这样写会略快一点
            if (set1.contains(x)) {
                common++;
            }
        }

        int n = nums1.length;
        int c1 = Math.min(set1.size() - common, n / 2);
        int c2 = Math.min(set2.size() - common, n / 2);
        return Math.min(n, c1 + c2 + common);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSetSize(vector<int> &nums1, vector<int> &nums2) {
        unordered_set<int> set1(nums1.begin(), nums1.end());
        unordered_set<int> set2(nums2.begin(), nums2.end());
        int common = 0;
        for (int x : set1) {
            common += set2.count(x);
        }

        int n = nums1.size();
        int c1 = min((int) set1.size() - common, n / 2);
        int c2 = min((int) set2.size() - common, n / 2);
        return min(n, c1 + c2 + common);
    }
};
```

```go [sol-Go]
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	set2 := map[int]bool{}
	common := 0
	for _, x := range nums2 {
		if set2[x] {
			continue
		}
		set2[x] = true
		// 另一种求 common 的写法
		if set1[x] {
			common++
		}
	}

	n := len(nums1)
	c1 := min(len(set1)-common, n/2)
	c2 := min(len(set2)-common, n/2)
	return min(n, c1+c2+common)
}
```

还有一种理解方式：

- 从 $\textit{nums}_1$ 中选择不超过 $n/2$ 个元素，并让这些元素尽量与 $\textit{nums}_2$ 没有交集。这可以选 $c_1=\min(n_1, n/2)$ 个。
- 从 $\textit{nums}_2$ 中选择不超过 $n/2$ 个元素，并让这些元素尽量与 $\textit{nums}_1$ 没有交集。这可以选 $c_2=\min(n_2, n/2)$ 个。
- 如果这两部分没有交集，那么答案就是 $c_1+c_2$。
- 如果这两部分有交集，那么答案就是 $\textit{nums}_1$ 和 $\textit{nums}_2$ 的**并集**的大小。
- 这两种情况取最小值。

```py [sol-Python3]
class Solution:
    def maximumSetSize(self, nums1: List[int], nums2: List[int]) -> int:
        set1 = set(nums1)
        set2 = set(nums2)
        n = len(nums1)
        c1 = min(len(set1), n // 2)
        c2 = min(len(set2), n // 2)
        return min(len(set1 | set2), c1 + c2)
```

```java [sol-Java]
class Solution {
    public int maximumSetSize(int[] nums1, int[] nums2) {
        Set<Integer> set1 = new HashSet<>();
        for (int x : nums1) {
            set1.add(x);
        }
        int all = set1.size();
        Set<Integer> set2 = new HashSet<>();
        for (int x : nums2) {
            if (set2.contains(x)) {
                continue;
            }
            set2.add(x);
            if (!set1.contains(x)) {
                all++;
            }
        }

        int n = nums1.length;
        int c1 = Math.min(set1.size(), n / 2);
        int c2 = Math.min(set2.size(), n / 2);
        return Math.min(all, c1 + c2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSetSize(vector<int> &nums1, vector<int> &nums2) {
        unordered_set<int> set1(nums1.begin(), nums1.end());
        unordered_set<int> set2(nums2.begin(), nums2.end());
        int all = set1.size() + set2.size();
        for (int x : set1) {
            all -= set2.count(x); // 去掉重复的
        }

        int n = nums1.size();
        int c1 = min((int) set1.size(), n / 2);
        int c2 = min((int) set2.size(), n / 2);
        return min(all, c1 + c2);
    }
};
```

```go [sol-Go]
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	all := len(set1)
	set2 := map[int]bool{}
	for _, x := range nums2 {
		if set2[x] {
			continue
		}
		set2[x] = true
		if !set1[x] {
			all++
		}
	}

	n := len(nums1)
	c1 := min(len(set1), n/2)
	c2 := min(len(set2), n/2)
	return min(all, c1+c2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
