从左到右遍历 $s$，用一个哈希集合保存遍历过的字母。

如果哈希集合的大小等于遍历过的字母个数模 $3$，即 $(i+1)\bmod 3$，那么把答案加一。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 方法一

```py [sol-Python3]
class Solution:
    def residuePrefixes(self, s: str) -> int:
        st = set()
        ans = 0
        for i, ch in enumerate(s):
            st.add(ch)
            if len(st) == (i + 1) % 3:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int residuePrefixes(String s) {
        Set<Character> set = new HashSet<>();
        int ans = 0;
        for (int i = 0; i < s.length(); i++) {
            set.add(s.charAt(i));
            if (set.size() == (i + 1) % 3) {
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
    int residuePrefixes(string s) {
        unordered_set<char> st;
        int ans = 0;
        for (int i = 0; i < s.size(); i++) {
            st.insert(s[i]);
            if (st.size() == (i + 1) % 3) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func residuePrefixes(s string) (ans int) {
	set := map[rune]struct{}{}
	for i, ch := range s {
		set[ch] = struct{}{}
		if len(set) == (i+1)%3 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(\min(n,|\Sigma|))$，其中 $|\Sigma|=26$ 是字符集合的大小。

## 方法二

用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

```py [sol-Python3]
class Solution:
    def residuePrefixes(self, s: str) -> int:
        ans = st = 0
        for i, ch in enumerate(s):
            st |= 1 << (ord(ch) - ord('a'))  # 把 ch 添加到 st 中
            if st.bit_count() == (i + 1) % 3:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int residuePrefixes(String s) {
        int set = 0;
        int ans = 0;
        for (int i = 0; i < s.length(); i++) {
            set |= 1 << (s.charAt(i) - 'a'); // 把 s[i] 添加到 set 中
            if (Integer.bitCount(set) == (i + 1) % 3) {
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
    int residuePrefixes(string s) {
        int st = 0;
        int ans = 0;
        for (int i = 0; i < s.size(); i++) {
            st |= 1 << (s[i] - 'a'); // 把 s[i] 添加到 st 中
            if (popcount((uint32_t) st) == (i + 1) % 3) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func residuePrefixes(s string) (ans int) {
	set := 0
	for i, ch := range s {
		set |= 1 << (ch - 'a') // 把 ch 添加到 set 中
		if bits.OnesCount(uint(set)) == (i+1)%3 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
