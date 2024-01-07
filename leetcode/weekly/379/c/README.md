[本题视频讲解](https://www.bilibili.com/video/BV1ae411e7fn/)

## 提示 1

设 $\textit{nums}_1$ 中有 $n_1$ 个不同元素，$\textit{nums}_2$ 中有 $n_2$ 个不同元素，它们的交集有 $\textit{common}$ 个元素。

如果不移除任何元素，根据容斥原理，$\textit{nums}_1$ 和 $\textit{nums}_2$ 的并集一共有

$$
\textit{ans} = n_1+n_2-\textit{common}
$$

个不同元素。

## 提示 2

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

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
