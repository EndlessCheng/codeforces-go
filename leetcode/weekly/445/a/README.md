时间等于距离除以速度，由于两人速度相同，转化成比较两人到第三人的**距离**。

设 $a=|x-z|$，$b=|y-z|$。

根据题意：

- 如果 $a=b$，返回 $0$。
- 如果 $a<b$，返回 $1$。
- 如果 $a>b$，返回 $2$。

## 写法一

```py [sol-Python3]
class Solution:
    def findClosest(self, x: int, y: int, z: int) -> int:
        a = abs(x - z)
        b = abs(y - z)
        if a == b:
            return 0
        return 1 if a < b else 2
```

```java [sol-Java]
class Solution {
    public int findClosest(int x, int y, int z) {
        int a = Math.abs(x - z);
        int b = Math.abs(y - z);
        return a == b ? 0 : a < b ? 1 : 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findClosest(int x, int y, int z) {
        int a = abs(x - z);
        int b = abs(y - z);
        return a == b ? 0 : a < b ? 1 : 2;
    }
};
```

```c [sol-C]
int findClosest(int x, int y, int z) {
    int a = abs(x - z);
    int b = abs(y - z);
    return a == b ? 0 : a < b ? 1 : 2;
}
```

```go [sol-Go]
func findClosest(x, y, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return 2
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```go [sol-Go 写法二]
var state = [3]int{1, 0, 2}

func findClosest(x, y, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	return state[cmp.Compare(a, b)+1]
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JS]
var findClosest = function(x, y, z) {
    const a = Math.abs(x - z);
    const b = Math.abs(y - z);
    return a === b ? 0 : a < b ? 1 : 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_closest(x: i32, y: i32, z: i32) -> i32 {
        let a = (x - z).abs();
        let b = (y - z).abs();
        if a == b { 0 } else if a < b { 1 } else { 2 }
    }
}
```

## 写法二

部分语言可以利用 $\texttt{bool}$ 自动转成 $\texttt{int}$ 的特性简化代码逻辑。

```py [sol-Python3]
class Solution:
    def findClosest(self, x: int, y: int, z: int) -> int:
        a = abs(x - z)
        b = abs(y - z)
        return (a > b) << 1 | (a < b)
```

```cpp [sol-C++]
class Solution {
public:
    int findClosest(int x, int y, int z) {
        int a = abs(x - z);
        int b = abs(y - z);
        return (a > b) << 1 | (a < b);
    }
};
```

```c [sol-C]
int findClosest(int x, int y, int z) {
    int a = abs(x - z);
    int b = abs(y - z);
    return (a > b) << 1 | (a < b);
}
```

```js [sol-JavaScript]
var findClosest = function(x, y, z) {
    const a = Math.abs(x - z);
    const b = Math.abs(y - z);
    return (a > b) << 1 | (a < b);
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
