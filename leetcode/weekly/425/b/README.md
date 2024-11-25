最终只需要关心 $s$ 是否等于 $t$，那么交换 $s$ 中的两个子串，等同于交换 $t$ 中的两个子串。

所以可以视作 $s$ 和 $t$ 中的子串可以随意重排。

思路和 [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) 一样，判断 $s$ 子串和 $t$ 子串排序后的结果是否一样。也可以用哈希表统计每个子串的出现次数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1fFB4YGEZY/)，欢迎点赞关注~

```py [sol-Python3 哈希表]
class Solution:
    def isPossibleToRearrange(self, s: str, t: str, k: int) -> bool:
        n = len(s)
        k = n // k
        cnt_s = Counter(s[i: i + k] for i in range(0, n, k))
        cnt_t = Counter(t[i: i + k] for i in range(0, n, k))
        return cnt_s == cnt_t
```

```py [sol-Python3 排序]
class Solution:
    def isPossibleToRearrange(self, s: str, t: str, k: int) -> bool:
        n = len(s)
        k = n // k
        a = sorted(s[i: i + k] for i in range(0, n, k))
        b = sorted(t[i: i + k] for i in range(0, n, k))
        return a == b
```

```java [sol-Java]
class Solution {
    public boolean isPossibleToRearrange(String s, String t, int k) {
        List<String> a = new ArrayList<>(k); // 预分配空间
        List<String> b = new ArrayList<>(k);
        int n = s.length();
        k = n / k;
        for (int i = k; i <= n; i += k) {
            a.add(s.substring(i - k, i));
            b.add(t.substring(i - k, i));
        }
        Collections.sort(a);
        Collections.sort(b);
        return a.equals(b);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isPossibleToRearrange(string s, string t, int k) {
        vector<string> a, b;
        int n = s.length();
        k = n / k;
        for (int i = k; i <= n; i += k) {
            a.push_back(s.substr(i - k, k));
            b.push_back(t.substr(i - k, k));
        }
        ranges::sort(a);
        ranges::sort(b);
        return a == b;
    }
};
```

```go [sol-Go]
func isPossibleToRearrange(s, t string, k int) bool {
	a := make([]string, 0, k) // 预分配空间
	b := make([]string, 0, k)
	n := len(s)
	k = n / k
	for i := k; i <= n; i += k {
		a = append(a, s[i-k:i])
		b = append(b, t[i-k:i])
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n\log k)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
