根据题意，$t$ 中字符是先进后出的，是一个栈。

问题相当于从左到右遍历 $s$，在允许用一个辅助栈的前提下，计算能得到的字典序最小的字符串。

看示例 2，$s=\texttt{bac}$：

- $s[0]=\texttt{b}$ 入栈，能立刻出栈吗？
   - 如果立刻出栈，那么答案的第一个字母是 $\texttt{b}$。
   - 如果不出栈，后面遍历到 $\texttt{a}$ 的时候把 $\texttt{a}$ 出栈，那么答案的第一个字母就会是 $\texttt{a}$，字典序更小。
   - 因此，如果我们发现后面有更小的字母，那就不出栈。
- $s[1]=\texttt{a}$ 入栈，能立刻出栈吗？
   - 如果立刻出栈，那么答案的第一个字母是 $\texttt{a}$。
   - 如果不出栈，后面也没有更小的字母了。把 $\texttt{c}$ 入栈再出栈，答案的第一个字母是 $\texttt{c}$，不是最优的。
   - 因此，如果我们发现当前字母比剩余字母（后缀 $s[i+1:]$）中的最小值还小，那么就立刻出栈。
   - 此外，$\texttt{a}$ 出栈后，应该继续把栈顶的 $\texttt{b}$ 也出栈，不然后面 $\texttt{c}$ 入栈出栈，会把 $\texttt{c}$ 填到答案的第二个位置上，不是最优的。
- $s[2]=\texttt{c}$ 入栈，再出栈。
- 最终答案为 $\texttt{abc}$。

如果 $s$ 中有相同字母，要如何处理呢？

以 $s=\texttt{caba}$ 为例。对于 $s[1]=\texttt{a}$，应该立刻出栈。如果不出栈，遍历到 $s[3]=\texttt{a}$ 才出栈，那么结果是 $\texttt{abac}$，但正确答案是 $\texttt{aabc}$。所以，如果栈顶等于剩余字母（后缀 $s[i+1:]$）中的最小值，也应该立刻出栈。

**总结**：如果栈顶 $\le$ 剩余字母（后缀 $s[i+1:]$）中的最小值，就立刻出栈。

```py [sol-Python3]
class Solution:
    def robotWithString(self, s: str) -> str:
      n = len(s)
      # 计算后缀最小值
      suf_min = ['z'] * (n + 1)
      for i in range(n - 1, -1, -1):
          suf_min[i] = min(suf_min[i + 1], s[i])

      ans = []
      st = []
      for i, ch in enumerate(s):
          st.append(ch)
          while st and st[-1] <= suf_min[i + 1]:
              ans.append(st.pop())
      return ''.join(ans)
```

```java [sol-Java]
// 更快的写法见【Java 数组】
class Solution {
    public String robotWithString(String s) {
        int n = s.length();
        // 计算后缀最小值
        char[] sufMin = new char[n + 1];
        sufMin[n] = Character.MAX_VALUE;
        for (int i = n - 1; i >= 0; i--) {
            sufMin[i] = (char) Math.min(sufMin[i + 1], s.charAt(i));
        }

        StringBuilder ans = new StringBuilder(n);
        Deque<Character> st = new ArrayDeque<>();
        for (int i = 0; i < n; i++) {
            st.push(s.charAt(i));
            while (!st.isEmpty() && st.peek() <= sufMin[i + 1]) {
                ans.append(st.pop());
            }
        }
        return ans.toString();
    }
}
```

```java [sol-Java 数组]
class Solution {
    public String robotWithString(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        // 计算后缀最小值
        char[] sufMin = new char[n + 1];
        sufMin[n] = Character.MAX_VALUE;
        for (int i = n - 1; i >= 0; i--) {
            sufMin[i] = (char) Math.min(sufMin[i + 1], s[i]);
        }

        char[] ans = s;
        char[] st = sufMin;
        int idx = 0;
        int top = -1;
        for (int i = 0; i < n; i++) {
            st[++top] = s[i];
            while (top >= 0 && st[top] <= sufMin[i + 1]) {
                ans[idx++] = st[top--];
            }
        }
        return new String(ans);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string robotWithString(string s) {
        int n = s.size();
        // 计算后缀最小值
        vector<char> suf_min(n + 1);
        suf_min[n] = 'z';
        for (int i = n - 1; i >= 0; i--) {
            suf_min[i] = min(suf_min[i + 1], s[i]);
        }

        string ans;
        stack<char> st;
        for (int i = 0; i < n; i++) {
            st.push(s[i]);
            while (!st.empty() && st.top() <= suf_min[i + 1]) {
                ans += st.top();
                st.pop();
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

char* robotWithString(char* s) {
    int n = strlen(s);
    // 计算后缀最小值
    char* suf_min = malloc((n + 1) * sizeof(char));
    suf_min[n] = 'z';
    for (int i = n - 1; i >= 0; i--) {
        suf_min[i] = MIN(suf_min[i + 1], s[i]);
    }

    char* ans = s;
    char* st = suf_min;
    int idx = 0, top = -1;
    for (int i = 0; i < n; i++) {
        st[++top] = s[i];
        while (top >= 0 && st[top] <= suf_min[i + 1]) {
            ans[idx++] = st[top--];
        }
    }
    ans[idx] = '\0';

    free(suf_min);
    return ans;
}
```

```go [sol-Go]
func robotWithString(s string) string {
	n := len(s)
	// 计算后缀最小值
	sufMin := make([]byte, n+1)
	sufMin[n] = math.MaxUint8
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}

	ans := make([]byte, 0, n)
	st := sufMin[:0]
	for i, ch := range s {
		st = append(st, byte(ch))
		for len(st) > 0 && st[len(st)-1] <= sufMin[i+1] {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return string(ans)
}
```

```js [sol-JavaScript]
var robotWithString = function(s) {
    const n = s.length;
    // 计算后缀最小值
    const sufMin = Array(n + 1);
    sufMin[n] = 'z';
    for (let i = n - 1; i >= 0; i--) {
        sufMin[i] = s[i] < sufMin[i + 1] ? s[i] : sufMin[i + 1];
    }

    const ans = [];
    const st = [];
    for (let i = 0; i < n; i++) {
        st.push(s[i]);
        while (st.length > 0 && st[st.length - 1] <= sufMin[i + 1]) {
            ans.push(st.pop());
        }
    }
    return ans.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn robot_with_string(s: String) -> String {
        let n = s.len();
        // 计算后缀最小值
        let mut suf_min = vec![u8::MAX; n + 1];
        for (i, ch) in s.bytes().enumerate().rev() {
            suf_min[i] = suf_min[i + 1].min(ch);
        }

        let mut ans = Vec::with_capacity(n);
        let mut st = vec![];
        for (i, ch) in s.bytes().enumerate() {
            st.push(ch);
            while let Some(&top) = st.last() {
                if top > suf_min[i + 1] {
                    break;
                }
                ans.push(st.pop().unwrap());
            }
        }
        unsafe { String::from_utf8_unchecked(ans) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。每个字母入栈出栈各恰好一次，所以二重循环的时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

