根据题意，只要有重复字母，就可以继续操作。

如果每种字母都只剩下一个，则无法操作。

所以答案为 $s$ 中不同字母的个数。

## 写法一：哈希集合

```py [sol-Python3]
class Solution:
    def minimizedStringLength(self, s: str) -> int:
        return len(set(s))
```

```java [sol-Java]
class Solution {
    public int minimizedStringLength(String s) {
        return (int) s.chars().distinct().count();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizedStringLength(string s) {
        return unordered_set(s.begin(), s.end()).size();
    }
};
```

```go [sol-Go]
func minimizedStringLength(s string) int {
    set := map[rune]struct{}{}
    for _, c := range s {
        set[c] = struct{}{}
    }
    return len(set)
}
```

## 写法二：位运算

原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def minimizedStringLength(self, s: str) -> int:
        mask = 0
        for c in s:
            mask |= 1 << (ord(c) - ord('a'))  # 把 c-'a' 加到集合中
        return mask.bit_count()  # 集合的大小
```

```java [sol-Java]
class Solution {
    public int minimizedStringLength(String s) {
        int mask = 0;
        for (char c : s.toCharArray()) {
            mask |= 1 << (c - 'a'); // 把 c-'a' 加到集合中
        }
        return Integer.bitCount(mask); // 集合的大小
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizedStringLength(string s) {
        unsigned mask = 0;
        for (char c : s) {
            mask |= 1 << (c - 'a'); // 把 c-'a' 加到集合中
        }
        return popcount(mask); // 集合的大小
    }
};
```

```go [sol-Go]
func minimizedStringLength(s string) int {
    mask := uint(0)
    for _, c := range s {
        mask |= 1 << (c - 'a') // 把 c-'a' 加到集合中
    }
    return bits.OnesCount(mask) // 集合的大小
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
