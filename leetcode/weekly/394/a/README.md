**一句话思路**：用两个集合分别记录小写字母和大写字母（转成小写字母），两个集合的**交集**大小就是答案。

集合可以用二进制数表示，两个二进制数的 AND 即为交集，AND 结果中的 $1$ 的个数即为答案。具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

附 ASCII 字符的性质：

- 对于大写英文字母：二进制从右往左第 $6$ 个比特一定是 $0$。
- 对于小写英文字母：二进制从右往左第 $6$ 个比特一定是 $1$。
- 对于任何英文字母：小写字母的二进制低 $5$ 位，等于其大写字母的二进制低 $5$ 位。

[本题视频讲解](https://www.bilibili.com/video/BV1gu4m1F7B8/)，欢迎点赞关注~

```py [sol-Python3]
# 不用位运算的写法见【Python3 set】
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        mask = [0] * 2  # 大写字母集合、小写字母集合
        for c in map(ord, word):
            # 用 c>>5&1 区分大小写，c&31 获取 c 是第几个字母
            mask[c >> 5 & 1] |= 1 << (c & 31)
        return (mask[0] & mask[1]).bit_count()  # 计算交集大小
```

```py [sol-Python3 set]
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        lower = set()
        upper = set()
        for c in word:
            if c.islower():
                lower.add(c)
            else:
                upper.add(c.lower())
        return len(lower & upper)
```

```java [sol-Java]
class Solution {
    public int numberOfSpecialChars(String word) {
        int[] mask = new int[2]; // 大写字母集合、小写字母集合
        for (char c : word.toCharArray()) {
            // 用 c>>5&1 区分大小写，c&31 获取 c 是第几个字母
            mask[c >> 5 & 1] |= 1 << (c & 31);
        }
        return Integer.bitCount(mask[0] & mask[1]); // 计算交集大小
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSpecialChars(string word) {
        int mask[2]{}; // 大写字母集合、小写字母集合
        for (char c : word) {
            // 用 c>>5&1 区分大小写，c&31 获取 c 是第几个字母
            mask[c >> 5 & 1] |= 1 << (c & 31);
        }
        return popcount(1u * (mask[0] & mask[1])); // 计算交集大小
    }
};
```

```go [sol-Go]
func numberOfSpecialChars(word string) int {
	mask := [2]int{} // 大写字母集合、小写字母集合
	for _, c := range word {
		// 用 c>>5&1 区分大小写，c&31 获取 c 是第几个字母
		mask[c>>5&1] |= 1 << (c & 31)
	}
	return bits.OnesCount(uint(mask[0] & mask[1])) // 计算交集大小
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
