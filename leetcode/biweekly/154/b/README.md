虽然题目要求 $i\le j\le k$，但因为异或运算满足交换律 $a\oplus b = b\oplus a$，实际上我们可以随意选。所以本质上，这题就是从 $\textit{nums}$ 中（可重复地）选三个数。

设 $n$ 为 $\textit{nums}$ 的长度。由于 $\textit{nums}$ 是一个排列，所以 $\textit{nums}$ 包含 $[1,n]$ 中的所有整数。

如果 $n=1$，只能选三个 $1$，异或值为 $1$，答案为 $1$。

如果 $n=2$，三个数中必然有两个数相等（鸽巢原理），这两个数的异或为 $0$，另一个数可以是 $1$ 也可以是 $2$，所以三数异或的结果是 $1$ 或者 $2$，答案为 $2$。

如果 $n\ge 3$，我们可以得到哪些异或值？

- 可以得到 $0$，选择 $1,2,3$ 这三个数。
- 可以得到 $[1,n]$ 中的任意整数 $a$，选择三个 $a$ 即可。
- 可以得到 $[n+1, 2^L-1]$ 中的任意整数 $a$，其中 $L=\left\lfloor \log_2 n\right\rfloor + 1$，也就是 $n$ 的二进制长度。
   - 比如二进制数 $a=1100$，我们可以先选一个最高位 $2^{L-1}=1000$，剩余的二进制数为 $a'=100$，可以通过 $a'\oplus 1 = 101$ 和 $1$ 异或得到。
   - 一般地，选择 $2^{L-1},a\oplus 2^{L-1}\oplus 1, 1$ 这三个数，可以异或得到 $[n+1, 2^L-1]$ 中的任意整数 $a$。
   - 特殊情况，如果 $a=2^{L-1}+1$，那么 $a\oplus 2^{L-1}\oplus 1=0$，可以改为选择 $2^{L-1},3,2$，异或得到 $a$。

所以我们可以得到 $[0,2^L-1]$ 中的任意数，这有 $2^L$ 个。

综上所述，如果 $n\le 2$，返回 $n$，否则返回 $2^L$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        n = len(nums)
        return n if n <= 2 else 1 << n.bit_length()
```

```java [sol-Java]
class Solution {
    public int uniqueXorTriplets(int[] nums) {
        int n = nums.length;
        return n <= 2 ? n : 1 << (32 - Integer.numberOfLeadingZeros(n));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int uniqueXorTriplets(vector<int>& nums) {
        size_t n = nums.size();
        return n <= 2 ? n : 1 << bit_width(n);
    }
};
```

```go [sol-Go]
func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	return 1 << bits.Len(uint(n))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
