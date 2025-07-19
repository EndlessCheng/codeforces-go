题意：把数组分成三段，每一段取第一个数再求和，问和的最小值是多少。

第一段的第一个数是确定的，即 $\textit{nums}[0]$。

如果知道了第二段的第一个数的位置，和第三段的第一个数的位置，那么这个划分方案也就确定了。

这两个下标可以在 $[1,n-1]$ 中随意取。

所以问题变成求下标在 $[1,n-1]$ 中的两个最小的数。

[视频讲解](https://www.bilibili.com/video/BV1oV411D7gB/)

## 方法一：直接排序

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        return nums[0] + sum(sorted(nums[1:])[:2])
```

```java [sol-Java]
class Solution {
    public int minimumCost(int[] nums) {
        Arrays.sort(nums, 1, nums.length);
        return nums[0] + nums[1] + nums[2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(vector<int>& nums) {
        sort(nums.begin() + 1, nums.end());
        return reduce(nums.begin(), nums.begin() + 3, 0);
    }
};
```

```go [sol-Go]
func minimumCost(nums []int) int {
	slices.Sort(nums[1:])
	return nums[0] + nums[1] + nums[2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 方法二：维护最小值和次小值

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        return nums[0] + sum(nsmallest(2, nums[1:]))
```

```java [sol-Java]
class Solution {
    public int minimumCost(int[] nums) {
        int fi = Integer.MAX_VALUE, se = Integer.MAX_VALUE;
        for (int i = 1; i < nums.length; i++) {
            int x = nums[i];
            if (x < fi) {
                se = fi;
                fi = x;
            } else if (x < se) {
                se = x;
            }
        }
        return nums[0] + fi + se;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(vector<int> &nums) {
        int fi = INT_MAX, se = INT_MAX;
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i];
            if (x < fi) {
                se = fi;
                fi = x;
            } else if (x < se) {
                se = x;
            }
        }
        return nums[0] + fi + se;
    }
};
```

```go [sol-Go]
func minimumCost(nums []int) int {
	fi, se := math.MaxInt, math.MaxInt
	for _, x := range nums[1:] {
		if x < fi {
			se = fi
			fi = x
		} else if x < se {
			se = x
		}
	}
	return nums[0] + fi + se
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
