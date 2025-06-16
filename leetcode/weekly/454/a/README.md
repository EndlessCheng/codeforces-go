模拟题：

1. 把 $\textit{caption}$ 按照空格，分割成若干个单词。
2. 把每个单词的字母全部小写。
3. 把除了第一个单词以外的单词，首字母大写。
4. 在单词列表前面加个井号。
5. 把单词列表 join 起来，得到答案。
6. 如果答案长度超过 $100$，截断，只保留前 $100$ 个字符。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1qsMxz6EEd/)，欢迎点赞关注~

**注**：不用库函数的写法见 Java 写法二。

```py [sol-Py3]
class Solution:
    def generateTag(self, caption: str) -> str:
        s = caption.title().replace(' ', '')
        if not s:
            return '#'
        return '#' + s[0].lower() + s[1:99]
```

```py [sol-Py3 写法二]
class Solution:
    def generateTag(self, caption: str) -> str:
        a = ['#']
        for i, s in enumerate(caption.split()):
            s = s.capitalize() if i else s.lower()
            a.append(s)
        return ''.join(a)[:100]  # join 后，取前 100 个字符
```

```java [sol-Java]
// 更快写法见【Java 写法二】
class Solution {
    public String generateTag(String caption) {
        StringBuilder ans = new StringBuilder("#");
        for (String s : caption.trim().split("\\s+")) { // 用一个或多个空格分隔 caption
            if (ans.length() == 1) { // s 是第一个单词
                s = s.toLowerCase();
            } else {
                s = s.substring(0, 1).toUpperCase() + s.substring(1).toLowerCase();
            }
            ans.append(s);
            if (ans.length() >= 100) {
                ans.setLength(100);
                break;
            }
        }
        return ans.toString();
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public String generateTag(String caption) {
        StringBuilder ans = new StringBuilder("#");
        char[] s = caption.toCharArray();
        for (int i = 0; i < s.length && ans.length() < 100; i++) {
            char c = s[i];
            if (c == ' ') {
                continue;
            }
            // 如果前一个字符是空格，那么当前字符是首字母
            if (ans.length() > 1 && s[i - 1] == ' ') { // 不是第一个单词
                ans.append((char) (c & ~32)); // 变成大写
            } else {
                ans.append((char) (c | 32)); // 变成小写
            }
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string generateTag(string caption) {
        stringstream ss(caption);
        string ans = "#";
        string s;
        while (ss >> s) {
            for (auto& ch : s) {
                ch = tolower(ch);
            }
            if (ans.size() > 1) { // 不是第一个单词，首字母大写
                s[0] = toupper(s[0]);
            }
            ans += s;
            if (ans.size() >= 100) {
                ans.resize(100);
                break;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func generateTag(caption string) string {
	ans := []byte{'#'}
	for i, s := range strings.Fields(caption) {
		s = strings.ToLower(s)
		if i > 0 { // 不是第一个单词，首字母大写
			s = strings.Title(s)
		}
		ans = append(ans, s...)
		if len(ans) >= 100 {
			ans = ans[:100]
			break
		}
	}
	return string(ans)
}
```

```go [sol-Go 写法二]
func generateTag(caption string) string {
	s := strings.ToLower(caption)
	s = strings.Title(s) // 所有单词首字母大写
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return "#"
	}
	s = "#" + string(unicode.ToLower(rune(s[0]))) + s[1:]
	if len(s) >= 100 {
		s = s[:100]
	}
	return s
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{caption}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，取决于实现。

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
