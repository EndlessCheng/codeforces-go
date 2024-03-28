**提示**：把反转看成是往字符串的头部添加字符。

具体来说：

- 如果当前处于「往字符串尾部添加字符」的状态，那么遇到 `i` 后，改成「往字符串头部添加字符」的状态。
- 如果当前处于「往字符串头部添加字符」的状态，那么遇到 `i` 后，改成「往字符串尾部添加字符」的状态。

这可以用**双端队列**实现。

不想用双端队列的话，可以参考下面 Go 和 JS 的实现。

附：[视频讲解](https://www.bilibili.com/video/BV1Yr4y1o7aP/)

```py [sol-Python3]
class Solution:
    def finalString(self, s: str) -> str:
        q = deque()
        tail = True
        for c in s:
            if c == 'i':
                tail = not tail  # 修改添加方向
            elif tail:  # 加尾部
                q.append(c)
            else:  # 加头部
                q.appendleft(c)
        return ''.join(q if tail else reversed(q))
```

```java [sol-Java]
class Solution {
    public String finalString(String s) {
        Deque<Character> q = new ArrayDeque<>();
        boolean tail = true;
        for (char c : s.toCharArray()) {
            if (c == 'i') {
                tail = !tail; // 修改添加方向
            } else if (tail) {
                q.addLast(c); // 加尾部
            } else {
                q.addFirst(c); // 加头部
            }
        }
        StringBuilder ans = new StringBuilder();
        for (char c : q) {
            ans.append(c);
        }
        if (!tail) {
            ans.reverse();
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string finalString(string s) {
        deque<char> q;
        bool tail = true;
        for (char c : s) {
            if (c == 'i') tail = !tail; // 修改添加方向
            else if (tail) q.push_back(c); // 加尾部
            else q.push_front(c); // 加头部
        }
        return tail ? string(q.begin(), q.end()) : string(q.rbegin(), q.rend());
    }
};
```

```go [sol-Go]
func finalString(s string) string {
	q := [2][]rune{} // 两个 slice 背靠背，q[0] 向左，q[1] 向右
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1 // 修改添加方向
		} else {
			q[dir] = append(q[dir], c)
		}
	}
	slices.Reverse(q[dir^1])
	return string(append(q[dir^1], q[dir]...))
}
```

```js [sol-JavaScript]
var finalString = function(s) {
    const q = [[], []] // 两个 list 背靠背，q[0] 向左，q[1] 向右
    let dir = 1;
    for (const c of s) {
        if (c === 'i') {
            dir ^= 1; // 修改添加方向
        } else {
            q[dir].push(c);
        }
    }
    return q[dir ^ 1].reverse().concat(q[dir]).join('');
};
```

```rust [sol-Rust]
use std::collections::VecDeque;

impl Solution {
    pub fn final_string(s: String) -> String {
        let mut q = VecDeque::new();
        let mut tail = true;
        for c in s.chars() {
            if c == 'i' {
                tail = !tail; // 修改添加方向
            } else if tail {
                q.push_back(c); // 加尾部
            } else {
                q.push_front(c); // 加头部
            }
        }
        if tail {
            q.iter().collect()
        } else {
            q.iter().rev().collect()
        }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
