题目要求子串内的字符都相同，那么子串越短，越容易满足要求，修改次数越少。例如子串 $\texttt{0011}$ 需要修改两个字符，分割成 $\texttt{00} + \texttt{11}$ 则无需修改字符。

一般地，对于任意最优分割方案，如果有长度大于 $2$ 的子串，那么把该子串继续分割成若干长为 $2$ 的子串，仍然满足要求。所以只需考虑分割成长为 $2$ 的子串。这个分割方案只有一种。

对于下标 $i=0,2,4,\ldots,n-2$，如果 $s[i]\ne s[i+1]$，那么必须修改其中一个字符。

> 注：题目保证 $s$ 的长度是偶数。

```py [sol-Python3]
class Solution:
    def minChanges(self, s: str) -> int:
        return sum(s[i] != s[i + 1] for i in range(0, len(s), 2))
```

```java [sol-Java]
class Solution {
    public int minChanges(String S) {
        char[] s = S.toCharArray();
        int ans = 0;
        for (int i = 0; i < s.length; i += 2) {
            if (s[i] != s[i + 1]) {
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
    int minChanges(string s) {
        int ans = 0;
        for (int i = 0; i < s.size(); i += 2) {
            if (s[i] != s[i + 1]) {
                ans++;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int minChanges(char* s) {
    int ans = 0;
    for (int i = 0; s[i]; i += 2) {
        if (s[i] != s[i + 1]) {
            ans++;
        }
    }
    return ans;
}
```

```go [sol-Go]
func minChanges(s string) (ans int) {
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var minChanges = function(s) {
    let ans = 0;
    for (let i = 0; i < s.length; i += 2) {
        if (s[i] !== s[i + 1]) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_changes(s: String) -> i32 {
        let s = s.as_bytes();
        let mut ans = 0;
        for i in (0..s.len()).step_by(2) {
            if s[i] != s[i + 1] {
                ans += 1;
            }
        }
        ans
    }
}
```

```rust [sol-Rust 函数式]
impl Solution {
    pub fn min_changes(s: String) -> i32 {
        s.as_bytes()
            .chunks_exact(2)
            .filter(|pair| pair[0] != pair[1])
            .count() as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

在修改次数最少的前提下，最少要分割成多少个子串（子串内的字符都一样）？

欢迎在评论区分享你的思路/代码。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
