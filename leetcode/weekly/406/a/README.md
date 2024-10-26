只能交换一次。要使字典序最小，需要满足下面两个要求：

1. 交换的两个数字，**左边必须大于右边**，否则交换不会使字典序变小。
2. 交换的位置**越靠左越好**。比如示例 1 的 $45320$，交换 $5$ 和 $3$ 得到 $43520$，而交换更靠右的 $2$ 和 $0$ 得到 $45302$，这比 $43520$ 更大。

细节：字符对应的数字的奇偶性，等于字符的 ASCII 值的奇偶性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1LZ421u7Ut/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def getSmallestString(self, s: str) -> str:
        t = list(s)
        for i in range(1, len(t)):
            x, y = t[i - 1], t[i]
            if x > y and ord(x) % 2 == ord(y) % 2:
                t[i - 1], t[i] = y, x
                break
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String getSmallestString(String s) {
        char[] t = s.toCharArray();
        for (int i = 1; i < t.length; i++) {
            char x = t[i - 1];
            char y = t[i];
            if (x > y && x % 2 == y % 2) {
                t[i - 1] = y;
                t[i] = x;
                break;
            }
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string getSmallestString(string s) {
        for (int i = 1; i < s.length(); i++) {
            char x = s[i - 1], y = s[i];
            if (x > y && x % 2 == y % 2) {
                swap(s[i - 1], s[i]);
                break;
            }
        }
        return s;
    }
};
```

```c [sol-C]
char* getSmallestString(char* s) {
    for (int i = 1; s[i]; i++) {
        char x = s[i - 1], y = s[i];
        if (x > y && x % 2 == y % 2) {
            s[i - 1] = y;
            s[i] = x;
            break;
        }
    }
    return s;
}
```

```go [sol-Go]
func getSmallestString(s string) string {
    t := []byte(s)
    for i := 1; i < len(t); i++ {
        x, y := t[i-1], t[i]
        if x > y && x%2 == y%2 {
            t[i-1], t[i] = y, x
            break
        }
    }
    return string(t)
}
```

```js [sol-JavaScript]
var getSmallestString = function(s) {
    const t = s.split('');
    for (let i = 1; i < t.length; i++) {
        const x = t[i - 1], y = t[i];
        if (x > y && x.charCodeAt(0) % 2 === y.charCodeAt(0) % 2) {
            t[i - 1] = y;
            t[i] = x;
            break;
        }
    }
    return t.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn get_smallest_string(mut s: String) -> String {
        unsafe {
            let t = s.as_bytes_mut();
            for i in 1..t.len() {
                let x = t[i - 1];
                let y = t[i];
                if x > y && x % 2 == y % 2 {
                    t[i - 1] = y;
                    t[i] = x;
                    break;
                }
            }
            s
        }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，取决于能否直接修改 $s$。

## 思考题

如果交换的是**任意两位**呢？

见 [670. 最大交换](https://leetcode.cn/problems/maximum-swap/)。

更多相似题目，见 [贪心题单](https://leetcode.cn/circle/discuss/g6KTKL/) 中的「**§3.1 字典序最小/最大**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
