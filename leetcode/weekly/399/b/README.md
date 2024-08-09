统计连续相同字母的长度，按题意模拟即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def compressedString(self, word: str) -> str:
        t = []
        i0 = -1
        for i, c in enumerate(word):
            if i + 1 == len(word) or c != word[i + 1]:
                k, rem = divmod(i - i0, 9)
                t.append(("9" + c) * k)
                if rem:
                    t.append(str(rem))
                    t.append(c)
                i0 = i
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String compressedString(String word) {
        StringBuilder t = new StringBuilder();
        char[] s = word.toCharArray();
        int i0 = -1;
        for (int i = 0; i < s.length; i++) {
            char c = s[i];
            if (i + 1 == s.length || c != s[i + 1]) {
                int k = i - i0;
                for (int j = 0; j < k / 9; j++) {
                    t.append('9').append(c);
                }
                if (k % 9 > 0) {
                    t.append((char) ('0' + (k % 9))).append(c);
                }
                i0 = i;
            }
        }
        return t.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string compressedString(string word) {
        string t;
        int i0 = -1;
        for (int i = 0; i < word.length(); i++) {
            char c = word[i];
            if (i + 1 == word.length() || c != word[i + 1]) {
                int k = i - i0;
                for (int j = 0; j < k / 9; j++) {
                    t += '9';
                    t += c;
                }
                if (k % 9) {
                    t += '0' + (k % 9);
                    t += c;
                }
                i0 = i;
            }
        }
        return t;
    }
};
```

```go [sol-Go]
func compressedString(word string) string {
	t := []byte{}
	i0 := -1
	for i := range word {
		c := word[i]
		if i+1 == len(word) || c != word[i+1] {
			k := i - i0
			t = append(t, bytes.Repeat([]byte{'9', c}, k/9)...)
			if k%9 > 0 {
				t = append(t, '0'+byte(k%9), c)
			}
			i0 = i
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
