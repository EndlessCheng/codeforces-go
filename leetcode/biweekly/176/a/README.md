「将结果按字母倒序映射到一个小写英文字母」，相当于把 $\texttt{z}$ 减去权重取模后的值：减 $0$ 就是 $\texttt{z}$，减 $1$ 就是 $\texttt{y}$ …… 减 $25$ 就是 $\texttt{a}$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        ans = []
        ord_a = ord('a')
        ord_z = ord('z')
        for w in words:
            s = sum(weights[ord(ch) - ord_a] for ch in w)
            ans.append(chr(ord_z - s % 26))
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String mapWordWeights(String[] words, int[] weights) {
        int n = words.length;
        char[] ans = new char[n];
        for (int i = 0; i < n; i++) {
            int sum = 0;
            for (char ch : words[i].toCharArray()) {
                sum += weights[ch - 'a'];
            }
            ans[i] = (char) ('z' - sum % 26);
        }
        return new String(ans);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string mapWordWeights(vector<string>& words, vector<int>& weights) {
        int n = words.size();
        string ans(n, 0);
        for (int i = 0; i < n; i++) {
            int sum = 0;
            for (char ch : words[i]) {
                sum += weights[ch - 'a'];
            }
            ans[i] = 'z' - sum % 26;
        }
        return ans;
    }
};
```

```go [sol-Go]
func mapWordWeights(words []string, weights []int) string {
	ans := make([]byte, len(words))
	for i, w := range words {
		sum := 0
		for _, ch := range w {
			sum += weights[ch-'a']
		}
		ans[i] = 'z' - byte(sum%26)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 是所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
