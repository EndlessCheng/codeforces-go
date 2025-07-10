用两个变量 $f_0$ 和 $f_1$ 记录字符串中是否有辅音或元音，必须都为 $\texttt{true}$ 才返回 $\texttt{true}$。

如果字符串长度不足 $3$ 或者包含除了数字或字母以外的字符，返回 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def isValid(self, word: str) -> bool:
        if len(word) < 3:
            return False
        f = [False] * 2
        for c in word:
            if c.isalpha():
                f[c.lower() in "aeiou"] = True
            elif not c.isdigit():
                return False
        return all(f)
```

```java [sol-Java]
class Solution {
    public boolean isValid(String word) {
        if (word.length() < 3) {
            return false;
        }
        boolean[] f = new boolean[2];
        for (char c : word.toCharArray()) {
            if (Character.isAlphabetic(c)) {
                c = Character.toLowerCase(c);
                f[c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ? 1 : 0] = true;
            } else if (!Character.isDigit(c)) {
                return false;
            }
        }
        return f[0] && f[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isValid(string word) {
        if (word.size() < 3) {
            return false;
        }
        bool f[2]{};
        for (char c : word) {
            if (isalpha(c)) {
                c = tolower(c);
                f[c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'] = true;
            } else if (!isdigit(c)) {
                return false;
            }
        }
        return f[0] && f[1];
    }
};
```

```go [sol-Go]
func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var f0, f1 bool
	for _, c := range word {
		if unicode.IsLetter(c) {
			if strings.ContainsRune("aeiou", unicode.ToLower(c)) {
				f1 = true
			} else {
				f0 = true
			}
		} else if !unicode.IsDigit(c) {
			return false
		}
	}
	return f0 && f1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
