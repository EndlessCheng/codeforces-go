**核心思路**：只有 $26$ 种字母，写一个 $26^2$ 的暴力枚举。

哈希表统计 $\textit{word}_1$ 字母出现次数 $c_1$，以及 $\textit{word}_2$ 字母出现次数 $c_2$：

- 枚举交换 $\textit{word}_1$ 中的字母 $x$ 和 $\textit{word}_2$ 中的字母 $y$。
- 如果 $x=y$，那么交换后不同字母数不变，并且哈希表 $c_1$ 和 $c_2$ 的大小相同，那么返回 $\texttt{true}$。
- 如果 $x\ne y$，那么就看 $x$ 的个数是否为 $1$，$y$ 的个数是否为 $1$，$x$ 是否出现在 $\textit{word}_2$ 中，$y$ 是否出现在 $\textit{word}_1$ 中，从而计算不同字母的变化量：
    - 如果 $x$ 个数为 $1$，那么交换后 $c_1$ 的大小减少了 $1$；
    - 如果 $y$ 个数为 $1$，那么交换后 $c_2$ 的大小减少了 $1$；
    - 如果 $y$ 不在 $\textit{word}_1$ 中，那么交换后 $c_1$ 的大小增加了 $1$；
    - 如果 $x$ 不在 $\textit{word}_2$ 中，那么交换后 $c_2$ 的大小增加了 $1$。

[视频讲解](https://www.bilibili.com/video/BV1KG4y1j73o/?t=8m51s)

```py [sol-Python3]
class Solution:
    def isItPossible(self, word1: str, word2: str) -> bool:
        c1 = Counter(word1)
        c2 = Counter(word2)
        for x, c in c1.items():
            for y, d in c2.items():
                if y == x:  # 无变化
                    if len(c1) == len(c2):
                        return True
                elif len(c1) - (c == 1) + (y not in c1) == \
                     len(c2) - (d == 1) + (x not in c2):  # 基于长度计算变化量
                    return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean isItPossible(String word1, String word2) {
        Map<Character, Integer> c1 = new HashMap<>();
        Map<Character, Integer> c2 = new HashMap<>();
        for (char c : word1.toCharArray()) {
            c1.merge(c, 1, Integer::sum);
        }
        for (char c : word2.toCharArray()) {
            c2.merge(c, 1, Integer::sum);
        }
        for (var e : c1.entrySet()) {
            for (var f : c2.entrySet()) {
                char x = e.getKey(), y = f.getKey();
                if (x == y) {
                    if (c1.size() == c2.size()) {
                        return true;
                    }
                } else if (c1.size() - (e.getValue() == 1 ? 1 : 0) + (c1.containsKey(y) ? 0 : 1) ==
                           c2.size() - (f.getValue() == 1 ? 1 : 0) + (c2.containsKey(x) ? 0 : 1)) { // 基于长度计算变化量
                    return true;
                }
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isItPossible(string word1, string word2) {
        unordered_map<char, int> c1, c2;
        for (char c : word1) c1[c]++;
        for (char c : word2) c2[c]++;
        for (auto& [x, c] : c1) {
            for (auto& [y, d] : c2) {
                if (y == x) {
                    if (c1.size() == c2.size()) {
                        return true;
                    }
                } else if (c1.size() - (c == 1) + !c1.contains(y) ==
                           c2.size() - (d == 1) + !c2.contains(x)) { // 基于长度计算变化量
                    return true;
                }
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func isItPossible(word1, word2 string) bool {
	c1 := map[rune]int{}
	for _, c := range word1 {
		c1[c]++
	}
	c2 := map[rune]int{}
	for _, c := range word2 {
		c2[c]++
	}
	for x, c := range c1 {
		for y, d := range c2 {
			if y == x { // 无变化
				if len(c1) == len(c2) {
					return true
				}
			} else if len(c1)-b2i(c == 1)+b2i(c1[y] == 0) ==
				      len(c2)-b2i(d == 1)+b2i(c2[x] == 0) { // 基于长度计算变化量
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m+|\Sigma|^2)$，其中 $n$ 为 $\textit{word}_1$ 的长度，$m$ 为 $\textit{word}_2$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
