把 $\textit{coordinate}_1$ 和 $\textit{coordinate}_2$ 简记为 $s$ 和 $t$。

根据题目中的图片，如果 $s[0]$ 和 $s[1]$ 的 ASCII 值的奇偶性相同，那么格子是黑格，否则是白格。

进一步地，由于奇数+奇数=偶数，偶数+偶数=偶数，所以如果 $(s[0] + s[1])\bmod 2$ 是偶数，那么格子是黑格；否则奇数+偶数=奇数，格子是白格。

如果 

$$
(s[0] + s[1])\bmod 2 = (t[0] + t[1])\bmod 2
$$

那么两个格子颜色相同，否则不同。

也可以取 $(s[0] \oplus s[1])$ 的最低位，其中 $\oplus$ 是异或运算。

具体请看 [视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def checkTwoChessboards(self, s: str, t: str) -> bool:
        return (ord(s[0]) + ord(s[1])) % 2 == (ord(t[0]) + ord(t[1])) % 2
```

```java [sol-Java]
class Solution {
    public boolean checkTwoChessboards(String s, String t) {
        int a = (s.charAt(0) + s.charAt(1)) % 2;
        int b = (t.charAt(0) + t.charAt(1)) % 2;
        return a == b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkTwoChessboards(string s, string t) {
        return ((s[0] ^ s[1]) & 1) == ((t[0] ^ t[1]) & 1);
    }
};
```

```go [sol-Go]
func checkTwoChessboards(s, t string) bool {
    return (s[0]^s[1])&1 == (t[0]^t[1])&1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
