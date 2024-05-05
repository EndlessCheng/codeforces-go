用两个变量 $f_0$ 和 $f_1$ 记录字符串中是否有辅音或元音，必须都为 $\texttt{true}$ 才返回 $\texttt{true}$。

如果字符串长度不足 $3$ 或者包含除了数字或字母以外的字符，返回 $\texttt{false}$。

请看 [视频讲解](https://www.bilibili.com/video/BV1Nf421U7em/)，欢迎点赞关注！

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
        Arrays.fill(f, false);
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
        if (word.length() < 3) {
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

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
