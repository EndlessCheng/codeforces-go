算法：

1. 找到 $p$ 中 $\texttt{*}$ 的位置，记作 $\textit{star}$。
2. 把 $p$ 分成两部分，$\texttt{*}$ 左边的记作 $p[:\textit{star}]$，右边的记作 $p[\textit{star}+1:]$。
3. 在 $s$ 中找 $p[:\textit{star}]$ 首次出现的位置 $i$。如果没找到，返回 $\texttt{false}$。
4. 继续匹配 $\texttt{*}$ 右边的内容，也就是判断 $p[\textit{star}+1:]$ 是否在 $s[i+\textit{star}:]$ 中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SzrAYMESJ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def hasMatch(self, s: str, p: str) -> bool:
        star = p.find('*')
        i = s.find(p[:star])
        return i >= 0 and p[star + 1:] in s[i + star:]
```

```java [sol-Java]
class Solution {
    public boolean hasMatch(String s, String p) {
        int star = p.indexOf('*');
        int i = s.indexOf(p.substring(0, star));
        return i >= 0 && s.substring(i + star).contains(p.substring(star + 1));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool hasMatch(const string& s, const string& p) {
        int star = p.find('*');
        int i = s.find(p.substr(0, star));
        return i != string::npos && s.substr(i + star).find(p.substr(star + 1)) != string::npos;
    }
};
```

```go [sol-Go]
func hasMatch(s, p string) bool {
	star := strings.IndexByte(p, '*')
	i := strings.Index(s, p[:star])
	return i >= 0 && strings.Contains(s[i+star:], p[star+1:])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $p$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$ 或 $\mathcal{O}(1)$。Go 的切片没有拷贝，是 $\mathcal{O}(1)$ 的。

用 KMP 算法可以进一步优化至 $\mathcal{O}(n+m)$。

## 思考题

如果 $p$ 中有多个 $\texttt{*}$ 呢？

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
12. 【本题相关】[字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
