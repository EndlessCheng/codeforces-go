## 方法一：遍历二进制数

遍历二进制 $n$，如果相邻两个数字都是 $1$，计数器 $\textit{cnt}$ 增加一。

最后判断 $\textit{cnt}$ 是否等于 $1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def consecutiveSetBits(self, n: int) -> bool:
        cnt = 0
        while n > 1:
            if n & 3 == 3:  # 最低两位是 11
                cnt += 1
            n >>= 1
        return cnt == 1
```

```java [sol-Java]
class Solution {
    public boolean consecutiveSetBits(int n) {
        int cnt = 0;
        for (; n > 1; n >>= 1) {
            if ((n & 3) == 3) { // 最低两位是 11
                cnt++;
            }
        }
        return cnt == 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool consecutiveSetBits(int n) {
        int cnt = 0;
        for (; n > 1; n >>= 1) {
            cnt += (n & 3) == 3; // n 的最低两位是 11
        }
        return cnt == 1;
    }
};
```

```go [sol-Go]
func consecutiveSetBits(n int) bool {
	cnt := 0
	for ; n > 1; n >>= 1 {
		if n&3 == 3 { // 最低两位是 11
			cnt++
		}
	}
	return cnt == 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：位运算技巧

由于 `n` 和 `n >> 1` 错开一位，所以计算 `n & (n >> 1)` 等价于计算所有相邻比特位的 `&`。所有相邻的 $11$ 变成 $1$，其余变成 $0$。

例如二进制数 $n = 10011001110$，`n & (n >> 1)` 为 $1000110$。 

如果 $n$ 只有一个 $11$，那么 `n & (n >> 1)` 只有一个 $1$，问题变成判断一个数是不是 [231. 2 的幂](https://leetcode.cn/problems/power-of-two/)，请看 [我的题解](https://leetcode.cn/problems/power-of-two/solutions/2973442/yan-ge-zheng-ming-yi-xing-xie-fa-pythonj-h04o/)。

```py [sol-Python3]
class Solution:
    def consecutiveSetBits(self, n: int) -> bool:
        m = n & (n >> 1)
        return m > 0 and m & (m - 1) == 0
```

```java [sol-Java]
class Solution {
    public boolean consecutiveSetBits(int n) {
        int m = n & (n >> 1);
        return m > 0 && (m & (m - 1)) == 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool consecutiveSetBits(int n) {
        int m = n & (n >> 1);
        return m > 0 && (m & (m - 1)) == 0;
    }
};
```

```go [sol-Go]
func consecutiveSetBits(n int) bool {
	m := n & (n >> 1)
	return m > 0 && m&(m-1) == 0
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
