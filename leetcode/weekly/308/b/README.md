**套路**：需要考虑相邻元素 + 有消除操作 = 栈。

请看 [视频讲解](https://www.bilibili.com/video/BV1mG411V7fj) 第二题。

```py [sol-Python3]
class Solution:
    def removeStars(self, s: str) -> str:
        st = []
        for c in s:
            if c == '*':
                st.pop()
            else:
                st.append(c)
        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String removeStars(String s) {
        StringBuilder st = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (c == '*') {
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
    string removeStars(string s) {
        string st;
        for (char c : s) {
            if (c == '*') {
                st.pop_back();
            } else {
                st += c;
            }
        }
        return st;
    }
};
```

```go [sol-Go]
func removeStars(s string) string {
	st := []rune{}
	for _, c := range s {
		if c == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
```

```js [sol-JavaScript]
var removeStars = function(s) {
    const st = [];
    for (const c of s) {
        if (c === '*') {
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
    pub fn remove_stars(s: String) -> String {
        let mut st = vec![];
        for &c in s.as_bytes() {
            if c == b'*' {
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

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。其中 C++ 不计入返回值的空间。

## 相似题目

- [1047. 删除字符串中的所有相邻重复项](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/)
- [2197. 替换数组中的非互质数](https://leetcode.cn/problems/replace-non-coprime-numbers-in-array/)
- [2216. 美化数组的最少删除数](https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/)
- [2273. 移除字母异位词后的结果数组](https://leetcode.cn/problems/find-resultant-array-after-removing-anagrams/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
