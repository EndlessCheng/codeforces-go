如果 $\textit{nums}_1$ 全是奇数或者全是偶数，那么只用第一种操作 `nums2[i] = nums1[i]` 即可满足要求。

否则，$\textit{nums}_1$ 奇数偶数都有。

分类讨论：

- 可以把奇数都变成偶数吗？注意奇数只能减去奇数才能变成偶数，那么不断操作奇数，直到只剩下一个奇数时，这个奇数只能减去自己。但题目要求 `j != i`，所以无法操作。所以无法把奇数都变成偶数。
- 可以把偶数都变成奇数吗？注意偶数只能减去奇数才能变成奇数。为了尽可能地满足 `nums1[i] - nums1[j] >= 1` 的要求，我们可以选一个**最小的奇数** $x$，把每个偶数都减去 $x$。这要求所有的偶数都要比 $x$ 大。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3 普通写法]
class Solution:
    def uniformArray(self, nums1: List[int]) -> bool:
        min_odd = inf
        for x in nums1:
            if x % 2:
                min_odd = min(min_odd, x)

        # 没有奇数，都是偶数
        if min_odd == inf:
            return True

        for x in nums1:
            # 把偶数减去奇数，变成奇数，前提是偶数 > 奇数
            if x % 2 == 0 and x < min_odd:
                return False

        return True
```

```py [sol-Python3 写法二]
class Solution:
    def uniformArray(self, nums1: List[int]) -> bool:
        min_odd = min((x for x in nums1 if x % 2), default=inf)
        return min_odd == inf or all(x % 2 or x > min_odd for x in nums1)
```

```java [sol-Java]
class Solution {
    public boolean uniformArray(int[] nums1) {
        int minOdd = Integer.MAX_VALUE;
        for (int x : nums1) {
            if (x % 2 != 0) {
                minOdd = Math.min(minOdd, x);
            }
        }

        // 没有奇数，都是偶数
        if (minOdd == Integer.MAX_VALUE) {
            return true;
        }

        for (int x : nums1) {
            // 把偶数减去奇数，变成奇数，前提是偶数 > 奇数
            if (x % 2 == 0 && x < minOdd) {
                return false;
            }
        }

        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool uniformArray(vector<int>& nums1) {
        int min_odd = INT_MAX;
        for (int x : nums1) {
            if (x % 2) {
                min_odd = min(min_odd, x);
            }
        }

        // 没有奇数，都是偶数
        if (min_odd == INT_MAX) {
            return true;
        }

        for (int x : nums1) {
            // 把偶数减去奇数，变成奇数，前提是偶数 > 奇数
            if (x % 2 == 0 && x < min_odd) {
                return false;
            }
        }

        return true;
    }
};
```

```go [sol-Go]
func uniformArray(nums1 []int) bool {
	minOdd := math.MaxInt
	for _, x := range nums1 {
		if x%2 != 0 {
			minOdd = min(minOdd, x)
		}
	}

	// 没有奇数，都是偶数
	if minOdd == math.MaxInt {
		return true
	}

	for _, x := range nums1 {
		// 把偶数减去奇数，变成奇数，前提是偶数 > 奇数
		if x%2 == 0 && x < minOdd {
			return false
		}
	}

	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」和「**§5.7 分类讨论**」。

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
