用栈维护，遇到数字则弹出栈顶，否则把字符入栈。最后从栈底到栈顶就是答案。

> 注：题目保证所有数字都可以按以上操作被删除。

如果你没有想到栈，推荐做做 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/) 中的 3.1 节和 3.3 节。

```py [sol-Python3]
class Solution:
    def clearDigits(self, s: str) -> str:
        st = []
        for c in s:
            if c.isdigit():
                st.pop()
            else:
                st.append(c)
        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String clearDigits(String s) {
        StringBuilder st = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (Character.isDigit(c)) {
                st.deleteCharAt(st.length() - 1);
            } else {
                st.append(c);
            }
        }
        return st.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearDigits(string s) {
        string st;
        for (char c : s) {
            if (isdigit(c)) {
                st.pop_back();
            } else {
                st += c;
            }
        }
        return st;
    }
};
```

```c [sol-C]
char* clearDigits(char* s) {
    int top = 0; // 栈顶
    for (int i = 0; s[i]; i++) {
        if (isdigit(s[i])) {
            top--; // 出栈
        } else {
            s[top++] = s[i]; // 入栈（把 s 当作栈）
        }
    }
    s[top] = '\0';
    return s;
}
```

```go [sol-Go]
func clearDigits(s string) string {
	st := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
```

```js [sol-JavaScript]
var clearDigits = function(s) {
    const st = [];
    for (const c of s) {
        if ('0' <= c && c <= '9') {
            st.pop();
        } else {
            st.push(c);
        }
    }
    return st.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn clear_digits(s: String) -> String {
        let mut st = vec![];
        for c in s.bytes() {
            if c.is_ascii_digit() {
                st.pop();
            } else {
                st.push(c);
            }
        }
        unsafe { String::from_utf8_unchecked(st) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。如果把 $s$ 当作栈，则空间复杂度为 $\mathcal{O}(1)$，见 C 语言。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
