由于奇数的二进制末尾一定是 $1$，我们可以把一个 $1$ 放在末尾，其余的 $1$ 全部放在开头，这样构造出的奇数尽量大。

```py [sol-Python3]
class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:
        cnt1 = s.count('1')
        return '1' * (cnt1 - 1) + '0' * (len(s) - cnt1) + '1'
```

```java [sol-Java]
class Solution {
    public String maximumOddBinaryNumber(String s) {
        int cnt1 = (int) s.chars().filter(c -> c == '1').count();
        return "1".repeat(cnt1 - 1) + "0".repeat(s.length() - cnt1) + "1";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string maximumOddBinaryNumber(string s) {
        int cnt1 = ranges::count(s, '1');
        return string(cnt1 - 1, '1') + string(s.length() - cnt1, '0') + '1';
    }
};
```

```go [sol-Go]
func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
```

```js [sol-JavaScript]
var maximumOddBinaryNumber = function (s) {
    const cnt1 = _.sumBy(s, c => c === '1')
    return '1'.repeat(cnt1 - 1) + '0'.repeat(s.length - cnt1) + '1';
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_odd_binary_number(s: String) -> String {
        let cnt1 = s.bytes().filter(|&c| c == b'1').count();
        "1".repeat(cnt1 - 1) + &"0".repeat(s.len() - cnt1) + "1"
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
