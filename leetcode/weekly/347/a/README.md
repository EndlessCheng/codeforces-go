由于输入保证是**正整数**，所以去掉所有的尾零 $0$ 即可。

```py [sol-Python3]
class Solution:
    def removeTrailingZeros(self, num: str) -> str:
        return num.rstrip('0')
```

```java [sol-Java]
class Solution {
    public String removeTrailingZeros(String num) {
        return num.replaceAll("0+$", "");
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string removeTrailingZeros(string s) {
        s.erase(s.begin() + 1 + s.find_last_not_of('0'), s.end()); // 原地操作
        return s;
    }
};
```

```go [sol-Go]
func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
```

```js [sol-JavaScript]
var removeTrailingZeros = function(num) {
    return num.replace(/0*$/, '');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn remove_trailing_zeros(num: String) -> String {
        num.trim_end_matches('0').to_string()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
