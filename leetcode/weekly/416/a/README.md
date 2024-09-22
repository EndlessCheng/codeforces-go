统计 $\textit{message}$ 中有多少个字符串在 $\textit{bannedWords}$ 中，如果超过 $1$ 个，返回 $\texttt{true}$，否则返回 $\texttt{false}$。

把 $\textit{bannedWords}$ 中的字符串全部丢到哈希集合中，从而加快效率。

```py [sol-Python3]
class Solution:
    def reportSpam(self, message: List[str], bannedWords: List[str]) -> bool:
        banned = set(bannedWords)
        return sum(s in banned for s in message) > 1
```

```py [sol-Python3 写法二]
class Solution:
    def reportSpam(self, message: List[str], bannedWords: List[str]) -> bool:
        banned = set(bannedWords)
        seen = False
        for s in message:
            if s in banned:
                if seen:
                    return True
                seen = True
        return False
```

```java [sol-Java]
class Solution {
    public boolean reportSpam(String[] message, String[] bannedWords) {
        Set<String> banned = new HashSet<>(Arrays.asList(bannedWords));
        int cnt = 0;
        for (String s : message) {
            if (banned.contains(s) && ++cnt > 1) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool reportSpam(vector<string>& message, vector<string>& bannedWords) {
        unordered_set<string> banned(bannedWords.begin(), bannedWords.end());
        int cnt = 0;
        for (auto& s : message) {
            if (banned.contains(s) && ++cnt > 1) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func reportSpam(message, bannedWords []string) bool {
	banned := map[string]bool{}
	for _, s := range bannedWords {
		banned[s] = true
	}

	seen := false
	for _, s := range message {
		if banned[s] {
			if seen {
				return true
			}
			seen = true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)l)$，其中 $n$ 是 $\textit{message}$ 的长度，$m$ 是 $\textit{bannedWords}$ 的长度，$l\le 15$ 是字符串的最大长度。
- 空间复杂度：$\mathcal{O}(ml)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
