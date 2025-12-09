用库函数 $\mathcal{O}(1)$ 得到二进制反转后的结果。Python 需要手写。

[本题视频讲解](https://www.bilibili.com/video/BV1sv2fB4Evi/?t=18m26s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sortByReflection(self, nums: List[int]) -> List[int]:
        # O(1) 反转见另一份代码【Python3 位运算】
        nums.sort(key=lambda x: (int(bin(x)[2:][::-1], 2), x))
        return nums
```

```py [sol-Python3 位运算]
class Solution:
    def sortByReflection(self, nums: List[int]) -> List[int]:
        nums.sort(key=lambda x: (self.reverseBits(x) >> (32 - x.bit_length()), x))
        return nums

    # 190. 颠倒二进制位
    # https://leetcode.cn/problems/reverse-bits/
    def reverseBits(self, n: int) -> int:
        # 交换 16 位
        n = ((n >> 16) | (n << 16)) & 0xFFFFFFFF
        # 交换每个 8 位块
        n = (((n & 0xFF00FF00) >> 8) | ((n & 0x00FF00FF) << 8)) & 0xFFFFFFFF
        # 交换每个 4 位块
        n = (((n & 0xF0F0F0F0) >> 4) | ((n & 0x0F0F0F0F) << 4)) & 0xFFFFFFFF
        # 交换每个 2 位块
        n = (((n & 0xCCCCCCCC) >> 2) | ((n & 0x33333333) << 2)) & 0xFFFFFFFF
        # 交换相邻位
        n = (((n & 0xAAAAAAAA) >> 1) | ((n & 0x55555555) << 1)) & 0xFFFFFFFF
        return n
```

```java [sol-Java]
class Solution {
    public int[] sortByReflection(int[] nums) {
        Integer[] arr = Arrays.stream(nums).boxed().toArray(Integer[]::new);

        Arrays.sort(arr, (a, b) -> {
            int revA = Integer.reverse(a) >>> Integer.numberOfLeadingZeros(a);
            int revB = Integer.reverse(b) >>> Integer.numberOfLeadingZeros(b);
            return revA != revB ? revA - revB : a - b;
        });

        for (int i = 0; i < nums.length; i++) {
            nums[i] = arr[i];
        }
        return nums;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public int[] sortByReflection(int[] nums) {
        int n = nums.length;
        long[] arr = new long[n];
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            long rev = Integer.reverse(x) >>> Integer.numberOfLeadingZeros(x);
            arr[i] = rev << 32 | x;
        }

        Arrays.sort(arr); // 比较 long 的高 32 位，相同的话比低 32 位

        for (int i = 0; i < n; i++) {
            nums[i] = (int) arr[i]; // 去掉 long 的高位
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> sortByReflection(vector<int>& nums) {
        ranges::sort(nums, {}, [&](int x) {
            int rev = __builtin_bitreverse32(x) >> countl_zero((uint32_t) x);
            return pair(rev, x);
        });
        return nums;
    }
};
```

```go [sol-Go]
func sortByReflection(nums []int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		revA := int(bits.Reverse(uint(a)) >> bits.LeadingZeros(uint(a)))
		revB := int(bits.Reverse(uint(b)) >> bits.LeadingZeros(uint(b)))
		return cmp.Or(revA-revB, a-b)
	})
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
