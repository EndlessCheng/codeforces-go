题目要求任意两个子字符串的首字符不能相同，那么答案的理论最大值为 $s$ 中的**不同字母个数**。

可以达到理论最大值吗？可以，在每种字母首次出现的位置前切一刀即可。

[本题视频讲解](https://www.bilibili.com/video/BV1D4SiB5Ee3/?t=54m08s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxDistinct(self, s: str) -> int:
        return len(set(s))
```

```java [sol-Java]
class Solution {
    public int maxDistinct(String s) {
        boolean[] vis = new boolean[26];
        int ans = 0;
        for (char c : s.toCharArray()) {
            c -= 'a';
            if (!vis[c]) {
                vis[c] = true;
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistinct(string s) {
        unordered_set<char> st(s.begin(), s.end());
        return st.size();
    }
};
```

```go [sol-Go]
func maxDistinct(s string) (ans int) {
	vis := [26]bool{}
	for _, c := range s {
		c -= 'a'
		if !vis[c] {
			vis[c] = true
			ans++
		}
	}
	return
}
```

## 优化

用二进制数表示集合/布尔数组，用位运算实现元素的添加。

请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def maxDistinct(self, s: str) -> int:
        st = 0
        for c in s:
            st |= 1 << (ord(c) - ord('a'))
        return st.bit_count()
```

```java [sol-Java]
class Solution {
    public int maxDistinct(String s) {
        int set = 0;
        for (char c : s.toCharArray()) {
            set |= 1 << (c - 'a');
        }
        return Integer.bitCount(set);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDistinct(string s) {
        uint32_t st = 0;
        for (char c : s) {
            st |= 1 << (c - 'a');
        }
        return popcount(st);
    }
};
```

```go [sol-Go]
func maxDistinct(s string) int {
	set := 0
	for _, c := range s {
		set |= 1 << (c - 'a')
	}
	return bits.OnesCount(uint(set))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面思维题单的「**§5.2 脑筋急转弯**」。

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
