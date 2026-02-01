**前置题目**：[344. 反转字符串](https://leetcode.cn/problems/reverse-string/)，[我的题解](https://leetcode.cn/problems/reverse-string/solutions/2376290/ji-chong-bu-tong-de-xie-fa-pythonjavacgo-9trb/)。

```py [sol-Python3]
class Solution:
    def reverse(self, t: List[str], f: Callable[[str], bool]) -> None:
        i, j = 0, len(t) - 1
        while i < j:
            while i < j and f(t[i]):
                i += 1
            while i < j and f(t[j]):
                j -= 1
            t[i], t[j] = t[j], t[i]
            i += 1
            j -= 1

    def reverseByType(self, s: str) -> str:
        t = list(s)
        self.reverse(t, str.islower)
        self.reverse(t, lambda ch: not ch.islower())
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String reverseByType(String s) {
        byte[] t = s.getBytes();
        reverse(t, ch -> 'a' <= ch && ch <= 'z');
        reverse(t, ch -> !('a' <= ch && ch <= 'z'));
        return new String(t);
    }

    private void reverse(byte[] t, Predicate<Byte> f) {
        int i = 0;
        int j = t.length - 1;
        while (i < j) {
            while (i < j && f.test(t[i])) {
                i++;
            }
            while (i < j && f.test(t[j])) {
                j--;
            }
            byte tmp = t[i];
            t[i] = t[j];
            t[j] = tmp;
            i++;
            j--;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    void reverse(string& s, auto&& f) {
        int i = 0;
        int j = s.size() - 1;
        while (i < j) {
            while (i < j && f(s[i])) {
                i++;
            }
            while (i < j && f(s[j])) {
                j--;
            }
            swap(s[i], s[j]);
            i++;
            j--;
        }
    }

public:
    string reverseByType(string s) {
        reverse(s, ::islower);
        reverse(s, [](char ch) { return !islower(ch); });
        return s;
    }
};
```

```go [sol-Go]
func reverse(t []byte, f func(byte) bool) {
	i, j := 0, len(t)-1
	for i < j {
		for i < j && f(t[i]) {
			i++
		}
		for i < j && f(t[j]) {
			j--
		}
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func reverseByType(s string) string {
	t := []byte(s)
	reverse(t, func(ch byte) bool { return 'a' <= ch && ch <= 'z' })
	reverse(t, func(ch byte) bool { return !('a' <= ch && ch <= 'z') })
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，取决于能否原地修改字符串。

## 专题训练

见下面双指针题单的「**§3.1 反转字符串**」。

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
