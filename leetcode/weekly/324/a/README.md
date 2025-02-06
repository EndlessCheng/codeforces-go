推荐先把 [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 做了。

为方便统计，把字符串 $s$ 中出现过的字母视作一个集合，把这个集合压缩成一个二进制数 $\textit{mask}$。其中 $\textit{mask}$ 第 $i$ 位为 $1$ 表示第 $i$ 个小写字母在 $s$ 中，为 $0$ 表示不在。这个技巧的详细解释见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

遍历 $\textit{words}$ 的同时，用一个哈希表 $\textit{cnt}$ 维护 $\textit{words}[i]$ 对应的 $\textit{mask}$ 的出现次数。和 1512 题一样，先把 $\textit{cnt}[\textit{mask}]$ 加到答案中，然后把 $\textit{cnt}[\textit{mask}]$ 加一。这个顺序可以保证我们只会统计 $i<j$ 的下标对，不会把 $i=j$ 的情况也统计进去。

```py [sol-Python3]
class Solution:
    def similarPairs(self, words: List[str]) -> int:
        ans = 0
        cnt = defaultdict(int)
        for s in words:
            mask = 0
            for c in s:
                mask |= 1 << (ord(c) - ord('a'))
            ans += cnt[mask]
            cnt[mask] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int similarPairs(String[] words) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int ans = 0;
        for (String s : words) {
            int mask = 0;
            for (char c : s.toCharArray()) {
                mask |= 1 << (c - 'a');
            }
            int c = cnt.getOrDefault(mask, 0);
            ans += c;
            cnt.put(mask, c + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int similarPairs(vector<string>& words) {
        unordered_map<int, int> cnt;
        int ans = 0;
        for (auto& s : words) {
            int mask = 0;
            for (char c : s) {
                mask |= 1 << (c - 'a');
            }
            ans += cnt[mask]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func similarPairs(words []string) (ans int) {
    cnt := map[int]int{}
    for _, s := range words {
        mask := 0
        for _, c := range s {
            mask |= 1 << (c - 'a')
        }
        ans += cnt[mask]
        cnt[mask]++
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{words}$ 的长度。哈希表需要 $\mathcal{O}(n)$ 的空间。

更多相似题目，见下面数据结构题单中的「**§0.1 枚举右，维护左**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
