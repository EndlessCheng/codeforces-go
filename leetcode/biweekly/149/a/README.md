统计 $s$ 中每个数字的出现次数。

遍历 $s$，如果相邻数字满足要求，返回这两个数字组成的字符串。

如果没有满足要求的相邻数字，返回空字符串。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1eUF6eaERQ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def findValidPair(self, s: str) -> str:
        cnt = Counter(s)
        for x, y in pairwise(s):
            if x != y and cnt[x] == int(x) and cnt[y] == int(y):
                return x + y
        return ""
```

```java [sol-Java]
class Solution {
    public String findValidPair(String S) {
        char[] s = S.toCharArray();
        int[] cnt = new int[10];
        for (char b : s) {
            cnt[b - '0']++;
        }
        for (int i = 1; i < s.length; i++) {
            int x = s[i - 1] - '0';
            int y = s[i] - '0';
            if (x != y && cnt[x] == x && cnt[y] == y) {
                return S.substring(i - 1, i + 1);
            }
        }
        return "";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string findValidPair(string s) {
        int cnt[10]{};
        for (char b : s) {
            cnt[b - '0']++;
        }
        for (int i = 1; i < s.size(); i++) {
            int x = s[i - 1] - '0', y = s[i] - '0';
            if (x != y && cnt[x] == x && cnt[y] == y) {
                return s.substr(i - 1, 2);
            }
        }
        return "";
    }
};
```

```go [sol-Go]
func findValidPair(s string) (ans string) {
	cnt := [10]int{}
	for _, b := range s {
		cnt[b-'0']++
	}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'0', s[i]-'0'
		if x != y && cnt[x] == int(x) && cnt[y] == int(y) {
			return s[i-1 : i+1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=10$ 是字符集合的大小。
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
