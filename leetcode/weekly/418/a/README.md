本题是 [179. 最大数](https://leetcode.cn/problems/largest-number/) 的二进制版本，做法是一样的。

把 $\textit{nums}$ 排序，对于相比较的两个元素 $a$ 和 $b$，如果其二进制表示（字符串）满足 $a+b>b+a$，那么 $a$ 排在 $b$ 的左边，否则 $b$ 排在 $a$ 的左边。

```py [sol-Python3]
class Solution:
    def maxGoodNumber(self, nums: list[int]) -> int:
        def cmp(a: int, b: int) -> int:
            len_a = a.bit_length()
            len_b = b.bit_length()
            return (b << len_a | a) - (a << len_b | b)
        nums.sort(key=cmp_to_key(cmp))

        ans = 0
        for x in nums:
            ans = ans << x.bit_length() | x
        return ans
```

```java [sol-Java]
class Solution {
    public int maxGoodNumber(int[] nums) {
        // Integer[] arr = Arrays.stream(nums).boxed().toArray(Integer[]::new);
        Integer[] arr = new Integer[]{nums[0], nums[1], nums[2]};
        Arrays.sort(arr, (a, b) -> {
            int lenA = 32 - Integer.numberOfLeadingZeros(a);
            int lenB = 32 - Integer.numberOfLeadingZeros(b);
            return (b << lenA | a) - (a << lenB | b);
        });

        int ans = 0;
        for (int x : arr) {
            int lenX = 32 - Integer.numberOfLeadingZeros(x);
            ans = ans << lenX | x;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxGoodNumber(vector<int>& nums) {
        ranges::sort(nums, [](int a, int b) {
            int len_a = __lg(a) + 1;
            int len_b = __lg(b) + 1;
            return (a << len_b | b) > (b << len_a | a);
        });

        int ans = 0;
        for (int x : nums) {
            ans = ans << (__lg(x) + 1) | x;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxGoodNumber(nums []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int {
		lenA := bits.Len(uint(a))
		lenB := bits.Len(uint(b))
		return (b<<lenA | a) - (a<<lenB | b)
	})

	for _, x := range nums {
		ans = ans<<bits.Len(uint(x)) | x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

更多相似题目，见下面贪心题单中的「**§1.7 交换论证法**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
