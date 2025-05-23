首先遍历 $s$，统计每个字母的出现次数，记录到一个 $\textit{cnt}$ 数组（哈希表）中。

然后遍历所有元音 $\texttt{aeiou}$，在 $\textit{cnt}$ 中查找元音的出现次数，统计最大值，然后把元音的出现次数置为 $0$。

去掉元音后，求整个 $\textit{cnt}$ 的最大值，就是辅音的出现次数的最大值了。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m7EuzqEqr/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxFreqSum(self, s: str) -> int:
        cnt = Counter(s)

        max_vowel = 0
        for b in "aeiou":
            max_vowel = max(max_vowel, cnt[b])
            del cnt[b]  # 这样下面计算的一定是辅音出现次数的最大值

        max_consonant = max(cnt.values()) if cnt else 0
        return max_vowel + max_consonant
```

```java [sol-Java]
class Solution {
    public int maxFreqSum(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }

        int maxVowel = 0;
        for (char b : "aeiou".toCharArray()) {
            maxVowel = Math.max(maxVowel, cnt[b - 'a']);
            cnt[b - 'a'] = 0; // 这样下面计算的一定是辅音出现次数的最大值
        }

        int maxConsonant = 0;
        for (int c : cnt) {
            maxConsonant = Math.max(maxConsonant, c);
        }

        return maxVowel + maxConsonant;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreqSum(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }

        int max_vowel = 0;
        for (char b : string("aeiou")) {
            max_vowel = max(max_vowel, cnt[b - 'a']);
            cnt[b - 'a'] = 0; // 这样下面计算的一定是辅音出现次数的最大值
        }

        int max_consonant = ranges::max(cnt);
        return max_vowel + max_consonant;
    }
};
```

```go [sol-Go]
func maxFreqSum(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	maxVowel := 0
	for _, b := range "aeiou" {
		maxVowel = max(maxVowel, cnt[b-'a'])
		cnt[b-'a'] = 0 // 这样下面计算的一定是辅音出现次数的最大值
	}

	maxConsonant := slices.Max(cnt[:])
	return maxVowel + maxConsonant
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
