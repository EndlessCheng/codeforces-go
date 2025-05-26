## 思路

想一想，第一个（最左边的）子串选哪个最好？

贪心地想，这个子串的右端点越小越好，这样后面能选的子串就越多。

选了右端点为 $i$ 的子串后，问题变成从 $[i+1,n-1]$ 中能选多少个子串，这是一个和原问题相似，规模更小的子问题。可以递归/迭代解决。

## 算法

遍历 $\textit{word}$，同时用 $\textit{pos}[\textit{ch}]$ 记录字母 $\textit{ch}=\textit{word}[i]$ 首次出现的位置。

- 如果 $\textit{ch}$ 首次出现，那么记录下标，即 $\textit{pos}[\textit{ch}] = i$。
- 否则，如果子串长度 $\ge 4$，即 $i-\textit{pos}[\textit{ch}]+1\ge 4$，即 $i-\textit{pos}[\textit{ch}] > 2$，那么找到了一个右端点尽量小的子串，答案加一。然后把 $\textit{pos}$ 重置，继续寻找下一个子串。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cqjgzdEPP/?t=9m12s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxSubstrings(self, word: str) -> int:
        ans = 0
        pos = {}
        for i, ch in enumerate(word):
            if ch not in pos:  # 之前没有遇到
                pos[ch] = i
            elif i - pos[ch] > 2:  # 再次遇到，且子串长度 >= 4
                ans += 1
                # 找下一个子串
                pos.clear()
        return ans
```

```java [sol-Java]
class Solution {
    public int maxSubstrings(String word) {
        int ans = 0;
        int[] pos = new int[26];
        Arrays.fill(pos, -1);
        for (int i = 0; i < word.length(); i++) {
            int ch = word.charAt(i) - 'a';
            if (pos[ch] < 0) { // 之前没有遇到
                pos[ch] = i;
            } else if (i - pos[ch] > 2) { // 再次遇到，且子串长度 >= 4
                ans++;
                // 找下一个子串
                Arrays.fill(pos, -1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubstrings(string word) {
        int ans = 0;
        int pos[26];
        ranges::fill(pos, -1);
        for (int i = 0; i < word.size(); i++) {
            int ch = word[i] - 'a';
            if (pos[ch] < 0) { // 之前没有遇到
                pos[ch] = i;
            } else if (i - pos[ch] > 2) { // 再次遇到，且子串长度 >= 4
                ans++;
                // 找下一个子串
                ranges::fill(pos, -1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go 写法一]
func maxSubstrings(word string) (ans int) {
	pos := [26]int{}
	for i := range pos {
		pos[i] = -1
	}
	for i, b := range word {
		b -= 'a'
		if pos[b] < 0 { // 之前没有遇到
			pos[b] = i
		} else if i-pos[b] > 2 { // 再次遇到，且子串长度 >= 4
			ans++
			// 找下一个子串
			for j := range pos {
				pos[j] = -1
			}
		}
	}
	return
}
```

```go [sol-Go 写法二]
func maxSubstrings(word string) (ans int) {
	pos := [26]int{}
	for i, b := range word {
		b -= 'a'
		if pos[b] == 0 {
			pos[b] = i + 1
		} else if i-pos[b] > 1 {
			ans++
			clear(pos[:])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 优化：位运算

其实，我们只需要知道在下标 $\le i-3$ 的字符中，是否有 $\textit{word}[i]$ 就行。所以可以用一个布尔数组记录。

进一步地，用 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/) 中的技巧，布尔数组压缩成一个二进制数。

```py [sol-Python3]
class Solution:
    def maxSubstrings(self, word: str) -> int:
        ans = seen = 0
        i = 3
        while i < len(word):
            seen |= 1 << (ord(word[i - 3]) - ord('a'))
            if seen >> (ord(word[i]) - ord('a')) & 1:  # 再次遇到 word[i]
                ans += 1
                seen = 0
                i += 3
            i += 1
        return ans  
```

```java [sol-Java]
class Solution {
    public int maxSubstrings(String word) {
        int ans = 0;
        int seen = 0;
        for (int i = 3; i < word.length(); i++) {
            seen |= 1 << (word.charAt(i - 3) - 'a');
            if ((seen >> (word.charAt(i) - 'a') & 1) > 0) { // 再次遇到 word[i]
                ans++;
                seen = 0;
                i += 3;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubstrings(string word) {
        int ans = 0, seen = 0;
        for (int i = 3; i < word.size(); i++) {
            seen |= 1 << (word[i - 3] - 'a');
            if (seen >> (word[i] - 'a') & 1) { // 再次遇到 word[i]
                ans++;
                seen = 0;
                i += 3;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubstrings(word string) (ans int) {
	seen := 0
	for i := 3; i < len(word); i++ {
		seen |= 1 << (word[i-3] - 'a')
		if seen>>(word[i]-'a')&1 > 0 { // 再次遇到 word[i]
			ans++
			seen = 0
			i += 3
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面贪心题单的「**§1.5 划分型贪心**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
