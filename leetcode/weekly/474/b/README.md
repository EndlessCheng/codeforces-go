计算绝对值前二大元素。

如果这两个元素取绝对值之前的乘积是负数，那么把第三个数变成 $-10^5$。

否则，把第三个数变成 $10^5$。

所以答案为绝对值前二大元素的乘积，再乘以 $10^5$。

[本题视频讲解](https://www.bilibili.com/video/BV1MgyfBoEuX/?t=1m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        mx = mx2 = 0
        for x in nums:
            x = abs(x)
            if x > mx:
                mx2 = mx  # 原来的最大变成次大
                mx = x  # x 是新的最大
            elif x > mx2:
                mx2 = x  # 最大不变，x 是新的次大
        return mx * mx2 * 10 ** 5
```

```java [sol-Java]
class Solution {
    public long maxProduct(int[] nums) {
        int mx = 0, mx2 = 0;
        for (int x : nums) {
            x = Math.abs(x);
            if (x > mx) {
                mx2 = mx; // 原来的最大变成次大
                mx = x; // x 是新的最大
            } else if (x > mx2) {
                mx2 = x; // 最大不变，x 是新的次大
            }
        }
        return 100000L * mx * mx2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProduct(vector<int>& nums) {
        int mx = 0, mx2 = 0;
        for (int x : nums) {
            x = abs(x);
            if (x > mx) {
                mx2 = mx; // 原来的最大变成次大
                mx = x; // x 是新的最大
            } else if (x > mx2) {
                mx2 = x; // 最大不变，x 是新的次大
            }
        }
        return 100000LL * mx * mx2;
    }
};
```

```go [sol-Go]
func maxProduct(nums []int) int64 {
	mx, mx2 := 0, 0
	for _, x := range nums {
		x = abs(x)
		if x > mx {
			mx2 = mx // 原来的最大变成次大
			mx = x // x 是新的最大
		} else if x > mx2 {
			mx2 = x // 最大不变，x 是新的次大
		}
	}
	return int64(mx * mx2 * 1e5)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心题单的「**§1.1 从最小/最大开始贪心**」。

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
