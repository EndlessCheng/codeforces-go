### 题意解读

每次操作可以把 $s[i]$ 加一或减一，求在操作次数不超过 $k$ 的前提下，$s$ 的最小字典序。注意 $\texttt{z}$ 加一后变成 $\texttt{a}$，$\texttt{a}$ 减一后变成 $\texttt{z}$。

### 分析

贪心，优先把左边的字母变成 $\texttt{a}$。

把 $s[i]$ 变成 $\texttt{a}$，可以不断减一到 $\texttt{a}$，也可以不断加一到 $\texttt{a}$，二者取最小值，得操作次数 

$$
\textit{dis} = \min(s[i] - \texttt{a}, \texttt{z} - s[i] + 1);
$$

### 算法

1. 从左到右遍历 $s$。
2. 如果把 $s[i]$ 变成 $\texttt{a}$ 的操作次数 $\textit{dis} \le k$，那么就把 $s[i]$ 变成 $\texttt{a}$，同时 $k$ 减少 $\textit{dis}$。
3. 否则无法变成 $\texttt{a}$，直接把 $s[i]$ 减少 $k$，退出循环。

请看 [视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def getSmallestString(self, s: str, k: int) -> str:
        s = list(s)
        for i, c in enumerate(map(ord, s)):
            dis = min(c - ord('a'), ord('z') - c + 1)
            if dis > k:
                s[i] = chr(c - k)
                break
            s[i] = 'a'
            k -= dis
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String getSmallestString(String s, int k) {
        char[] t = s.toCharArray();
        for (int i = 0; i < t.length; i++) {
            int dis = Math.min(t[i] - 'a', 'z' - t[i] + 1);
            if (dis > k) {
                t[i] -= k;
                break;
            }
            t[i] = 'a';
            k -= dis;
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string getSmallestString(string s, int k) {
        for (char& c : s) {
            int dis = min(c - 'a', 'z' - c + 1);
            if (dis > k) {
                c -= k;
                break;
            }
            c = 'a';
            k -= dis;
        }
        return s;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

char* getSmallestString(char* s, int k) {
    for (int i = 0; s[i]; i++) {
        int dis = MIN(s[i] - 'a', 'z' - s[i] + 1);
        if (dis > k) {
            s[i] -= k;
            break;
        }
        s[i] = 'a';
        k -= dis;
    }
    return s;
}
```

```go [sol-Go]
func getSmallestString(s string, k int) string {
	t := []byte(s)
	for i, c := range t {
		dis := int(min(c-'a', 'z'-c+1))
		if dis > k {
			t[i] -= byte(k)
			break
		}
		t[i] = 'a'
		k -= dis
	}
	return string(t)
}
```

```js [sol-JavaScript]
var getSmallestString = function(s, k) {
    let t = s.split('');
    for (let i = 0; i < t.length; i++) {
        const dis = Math.min(t[i].charCodeAt(0) - 'a'.charCodeAt(0), 'z'.charCodeAt(0) - t[i].charCodeAt(0) + 1);
        if (dis > k) {
            t[i] = String.fromCharCode(t[i].charCodeAt(0) - k);
            break;
        }
        t[i] = 'a';
        k -= dis;
    }
    return t.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn get_smallest_string(s: String, mut k: i32) -> String {
        let mut s = s.into_bytes();
        for c in s.iter_mut() {
            let dis = (*c - b'a').min(b'z' - *c + 1) as i32;
            if dis > k {
                *c -= k as u8;
                break;
            }
            *c = b'a';
            k -= dis;
        }
        unsafe { String::from_utf8_unchecked(s) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，其中 C/C++ 可以原地修改字符串。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
