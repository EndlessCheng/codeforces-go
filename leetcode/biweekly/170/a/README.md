## 方法一：转成字符串处理

设 $n$ 的二进制字符串为 $s$，长度为 $m$。

根据题意，遍历 $i=0,1,2,\ldots,n-1$，如果 $s[i]\ne s[n-1-i]$，那么需要反转 $s[i]$，答案加一。

也可以遍历 $i=0,1,2,\ldots,\left\lfloor\dfrac{n}{2}\right\rfloor$，如果 $s[i]\ne s[n-1-i]$，那么需要反转 $s[i]$ 和 $s[n-1-i]$，答案加二。

示例 2 在 $i=0$ 和 $i=1$ 处都满足 $s[i]\ne s[n-1-i]$，答案增加两次 $2$，得到答案 $4$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minimumFlips(self, n: int) -> int:
        s = bin(n)[2:]
        ans = 0
        for i in range(len(s) // 2):
            if s[i] != s[-1 - i]:
                ans += 2
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumFlips(int n) {
        char[] s = Long.toBinaryString(n).toCharArray();
        int m = s.length;
        int ans = 0;
        for (int i = 0; i < m / 2; i++) {
            if (s[i] != s[m - 1 - i]) {
                ans += 2;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumFlips(int n) {
        string s = format("{:b}", n); // #include<format>
        int m = s.size();
        int ans = 0;
        for (int i = 0; i < m / 2; i++) {
            if (s[i] != s[m - 1 - i]) {
                ans += 2;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumFlips(n int) (ans int) {
	s := strconv.FormatUint(uint64(n), 2)
	m := len(s)
	for i := range m / 2 {
		if s[i] != s[m-1-i] {
			ans += 2
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。其中 $n$ 的二进制长度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。

## 方法二：位运算

根据 [190. 颠倒二进制位](https://leetcode.cn/problems/reverse-bits/)，我们可以 $\mathcal{O}(1)$ 得到 $n$ 反转后的值 $\textit{rev}$。

计算 $n$ 和 $\textit{rev}$ 有多少个位置不同，等价于计算 $n\oplus \textit{rev}$ 的二进制中有多少个 $1$。其中 $\oplus$ 是异或运算，对于每个比特位，两个数不同时结果才是 $1$。

```py [sol-Python3]
class Solution:
    def minimumFlips(self, n: int) -> int:
        rev = self.reverseBits(n) >> (32 - n.bit_length())
        return (n ^ rev).bit_count()

    # 190. 颠倒二进制位
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
    public int minimumFlips(int n) {
        int rev = Integer.reverse(n) >>> Integer.numberOfLeadingZeros(n);
        return Integer.bitCount(n ^ rev);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumFlips(int num) {
        uint32_t n = num;
        uint32_t rev = __builtin_bitreverse32(n) >> countl_zero(n);
        return popcount(n ^ rev);
    }
};
```

```go [sol-Go]
func minimumFlips(num int) int {
	n := uint32(num)
	rev := bits.Reverse32(n) >> bits.LeadingZeros32(n)
	return bits.OnesCount32(n ^ rev)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面位运算题单的「**一、基础题**」。

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
