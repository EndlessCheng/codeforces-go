```py [sol-Python3]
import numpy as np

class Solution:
    def concatHex36(self, n: int) -> str:
        return np.base_repr(n ** 2, base=16) + np.base_repr(n ** 3, base=36)
```

```java [sol-Java]
class Solution {
    public String concatHex36(int n) {
        String s = Integer.toHexString(n * n) + Integer.toString(n * n * n, 36);
        return s.toUpperCase();
    }
}
```

```cpp [sol-C++]
// C++ 只能手写了
class Solution {
    string base_repr(int v, int base) {
        string s;
        while (v > 0) {
            int d = v % base;
            s += d < 10 ? '0' + d : 'A' + d - 10;
            v /= base;
        }
        ranges::reverse(s);
        return s;
    }

public:
    string concatHex36(int n) {
        return base_repr(n * n, 16) + base_repr(n * n * n, 36);
    }
};
```

```go [sol-Go]
func concatHex36(n int) string {
	s := strconv.FormatInt(int64(n*n), 16) + strconv.FormatInt(int64(n*n*n), 36)
	return strings.ToUpper(s)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。$\mathcal{O}(\log n^3) = \mathcal{O}(3\log n) = \mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。

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
