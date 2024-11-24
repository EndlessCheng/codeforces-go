## 方法一：暴力枚举

暴力枚举子数组的左端点 $i$，然后枚举右端点 $j$。

可以在枚举的过程中累加元素和。

```py [sol-Python3]
class Solution:
    def minimumSumSubarray(self, nums: list[int], l: int, r: int) -> int:
        ans = inf
        n = len(nums)
        for i in range(n - l + 1):
            s = 0
            for j in range(i, min(i + r, n)):
                s += nums[j]
                if s > 0 and j - i + 1 >= l:
                    ans = min(ans, s)
        return -1 if ans == inf else ans
```

```java [sol-Java]
class Solution {
    public int minimumSumSubarray(List<Integer> nums, int l, int r) {
        Integer[] a = nums.toArray(Integer[]::new);
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i <= a.length - l; i++) {
            int s = 0;
            for (int j = i; j < a.length && j - i + 1 <= r; j++) {
                s += a[j];
                if (s > 0 && j - i + 1 >= l) {
                    ans = Math.min(ans, s);
                }
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSumSubarray(vector<int>& nums, int l, int r) {
        int ans = INT_MAX;
        for (int i = 0; i <= nums.size() - l; i++) {
            int s = 0;
            for (int j = i; j < nums.size() && j - i + 1 <= r; j++) {
                s += nums[j];
                if (s > 0 && j - i + 1 >= l) {
                    ans = min(ans, s);
                }
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minimumSumSubarray(nums []int, l, r int) int {
	ans := math.MaxInt
	n := len(nums)
	for i := range n - l + 1 {
		s := 0
		for j := i; j < n && j-i+1 <= r; j++ {
			s += nums[j]
			if s > 0 && j-i+1 >= l {
				ans = min(ans, s)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-l)r)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：前缀和+定长滑窗+有序集合

利用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)，问题变成：

- 找到一个小于 $s[j]$ 且离 $s[j]$ 最近的前缀和 $s[i]$，满足 $l\le j-i\le r$。

枚举 $j$，那么 $i$ 需要满足 $j-r\le i\le j-l$。

用**有序集合**维护这个范围内的 $s[i]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1fFB4YGEZY/?t=3m47s)，欢迎点赞关注~

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        ans = inf
        s = list(accumulate(nums, initial=0))
        sl = SortedList()
        for j in range(l, len(nums) + 1):
            sl.add(s[j - l])
            k = sl.bisect_left(s[j])
            if k:
                ans = min(ans, s[j] - sl[k - 1])
            if j >= r:
                sl.remove(s[j - r])
        return -1 if ans == inf else ans
```

```java [sol-Java]
class Solution {
    public int minimumSumSubarray(List<Integer> nums, int l, int r) {
        Integer[] a = nums.toArray(Integer[]::new);
        int ans = Integer.MAX_VALUE;
        int n = a.length;
        int[] s = new int[n + 1];
        TreeMap<Integer, Integer> cnt = new TreeMap<>();
        for (int j = 1; j <= n; j++) {
            s[j] = s[j - 1] + a[j - 1];
            if (j < l) {
                continue;
            }
            cnt.merge(s[j - l], 1, Integer::sum); // cnt[s[j-l]]++
            Integer lower = cnt.lowerKey(s[j]);
            if (lower != null) {
                ans = Math.min(ans, s[j] - lower);
            }
            if (j >= r) {
                int v = s[j - r];
                int c = cnt.get(v);
                if (c == 1) {
                    cnt.remove(v);
                } else {
                    cnt.put(v, c - 1);
                }
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSumSubarray(vector<int>& nums, int l, int r) {
        int ans = INT_MAX;
        int n = nums.size();
        vector<int> s(n + 1);
        multiset<int> s_set;
        for (int j = 1; j <= n; j++) {
            s[j] = s[j - 1] + nums[j - 1];
            if (j < l) {
                continue;
            }
            s_set.insert(s[j - l]);
            auto it = s_set.lower_bound(s[j]);
            if (it != s_set.begin()) {
                ans = min(ans, s[j] - *--it);
            }
            if (j >= r) {
                s_set.erase(s_set.find(s[j - r]));
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minimumSumSubarray(nums []int, l, r int) int {
	ans := math.MaxInt
	n := len(nums)
	s := make([]int, n+1)
	cnt := redblacktree.New[int, int]() // "github.com/emirpasic/gods/v2/trees/redblacktree"
	for j := 1; j <= n; j++ {
		s[j] = s[j-1] + nums[j-1]
		if j < l {
			continue
		}
		c, _ := cnt.Get(s[j-l])
		cnt.Put(s[j-l], c+1)
		if lower, ok := cnt.Floor(s[j] - 1); ok {
			ans = min(ans, s[j]-lower.Key)
		}
		if j >= r {
			v := s[j-r]
			c, _ := cnt.Get(v)
			if c == 1 {
				cnt.Remove(v)
			} else {
				cnt.Put(v, c-1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + (n-l)\log (r-l))$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
