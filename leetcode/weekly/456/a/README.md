按题意模拟即可。

为了快速判断当前字符串 $t$ 是否在答案中，用一个哈希集合保存在答案中的字符串。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def partitionString(self, s: str) -> List[str]:
        ans = []
        vis = set()
        t = ''
        for c in s:
            t += c
            if t not in vis:
                vis.add(t)
                ans.append(t)
                t = ''
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> partitionString(String s) {
        List<String> ans = new ArrayList<>();
        Set<String> vis = new HashSet<>();
        String t = "";
        for (char c : s.toCharArray()) {
            t += c;
            if (vis.add(t)) { // t 不在 vis 中
                ans.add(t);
                t = "";
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> partitionString(string s) {
        vector<string> ans;
        unordered_set<string> vis;
        string t;
        for (char c : s) {
            t += c;
            if (vis.insert(t).second) { // t 不在 vis 中
                ans.push_back(t);
                t.clear();
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func partitionString(s string) (ans []string) {
	vis := map[string]bool{}
	t := ""
	for _, c := range s {
		t += string(c)
		if !vis[t] {
			vis[t] = true
			ans = append(ans, t)
			t = ""
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n)$，其中 $n$ 是 $s$ 的长度。最坏情况下 $n$ 个一样的字母，加到答案中的字符串长度为 $1,2,3,\ldots,k$，解不等式 $1+2+3+\cdots+k = \dfrac{k(k+1)}{2}\le n$，得 $k = \mathcal{O}(\sqrt n)$。每次判断一个长为 $\mathcal{O}(\sqrt n)$ 的字符串是否在哈希集合中，需要 $\mathcal{O}(\sqrt n)$ 的时间，一共判断 $n$ 次。所以时间复杂度为 $\mathcal{O}(n\sqrt n)$。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：用字典树（或者字符串哈希），可以做到 $\mathcal{O}(n)$。直播结束后补充。

## 专题训练

见下面数据结构题单的「**六、字典树（trie）**」。

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
