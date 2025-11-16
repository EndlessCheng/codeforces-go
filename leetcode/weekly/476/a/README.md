要让 $a+b-c$ 尽量大，$a$ 和 $b$ 要尽量大，$c$ 要尽量小。

把 $\textit{nums}$ 从小到大排序，取 $a= \textit{nums}[n-1],b=\textit{nums}[n-2],c=\textit{nums}[0]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maximizeExpressionOfThree(self, nums: List[int]) -> int:
        nums.sort()
        return nums[-1] + nums[-2] - nums[0]
```

```java [sol-Java]
class Solution {
    public int maximizeExpressionOfThree(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        return nums[n - 1] + nums[n - 2] - nums[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeExpressionOfThree(vector<int>& nums) {
        ranges::sort(nums);
        int n = nums.size();
        return nums[n - 1] + nums[n - 2] - nums[0];
    }
};
```

```go [sol-Go]
func maximizeExpressionOfThree(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	return nums[n-1] + nums[n-2] - nums[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 优化

手动维护前二大（或者用堆），以及最小值。

```py [sol-Python3]
class Solution:
    def maximizeExpressionOfThree(self, nums: List[int]) -> int:
        return sum(nlargest(2, nums)) - min(nums)
```

```java [sol-Java]
class Solution {
    public int maximizeExpressionOfThree(int[] nums) {
        int mx = Integer.MIN_VALUE;
        int mx2 = Integer.MIN_VALUE;
        int mn = Integer.MAX_VALUE;
        for (int x : nums) {
            if (x > mx) {
                mx2 = mx;
                mx = x;
            } else if (x > mx2) {
                mx2 = x;
            }
            mn = Math.min(mn, x);
        }
        return mx + mx2 - mn;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeExpressionOfThree(vector<int>& nums) {
        int mx = INT_MIN, mx2 = INT_MIN, mn = INT_MAX;
        for (int x : nums) {
            if (x > mx) {
                mx2 = mx;
                mx = x;
            } else if (x > mx2) {
                mx2 = x;
            }
            mn = min(mn, x);
        }
        return mx + mx2 - mn;
    }
};
```

```go [sol-Go]
func maximizeExpressionOfThree(nums []int) int {
	mx, mx2, mn := math.MinInt, math.MinInt, math.MaxInt
	for _, x := range nums {
		if x > mx {
			mx2 = mx
			mx = x
		} else if x > mx2 {
			mx2 = x
		}
		mn = min(mn, x)
	}
	return mx + mx2 - mn
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
