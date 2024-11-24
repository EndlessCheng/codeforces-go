## 方法一：模拟

具体请看 [视频讲解](https://www.bilibili.com/video/BV1uzBxYoEJC/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def canAliceWin(self, n: int) -> bool:
        pick = 10
        while n >= pick:
            n -= pick
            pick -= 1
        return (10 - pick) % 2 > 0
```

```java [sol-Java]
class Solution {
    public boolean canAliceWin(int n) {
        int pick = 10;
        while (n >= pick) {
            n -= pick;
            pick--;
        }
        return (10 - pick) % 2 > 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canAliceWin(int n) {
        int pick = 10;
        while (n >= pick) {
            n -= pick;
            pick--;
        }
        return (10 - pick) % 2;
    }
};
```

```go [sol-Go]
func canAliceWin(n int) bool {
	pick := 10
	for n >= pick {
		n -= pick
		pick--
	}
	return (10-pick)%2 > 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(C)$。其中 $C=10$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学公式

两人交替操作，设两人一共操作了 $x$ 次。

那么有

$$
10+9+8+\cdots+(11-x) = \dfrac{(21-x)x}{2} \le n
$$

解一元二次不等式，得

$$
x \le \dfrac{21 - \sqrt {21^2-8n}}{2}
$$

由于 $x$ 是整数，所以 $x$ 的最大值为

$$
\left\lfloor\dfrac{21 - \sqrt {21^2-8n}}{2}\right\rfloor = \left\lfloor\dfrac{21 - \lceil\sqrt {21^2-8n}\rceil}{2}\right\rfloor
$$

如果 $x$ 是奇数，Alice 胜，否则 Bob 胜。

```py [sol-Python3]
class Solution:
    def canAliceWin(self, n: int) -> bool:
        x = (21 - ceil(sqrt(441 - n * 8))) // 2
        return x % 2 > 0
```

```java [sol-Java]
class Solution {
    public boolean canAliceWin(int n) {
        int x = (21 - (int) Math.ceil(Math.sqrt(441 - n * 8))) / 2;
        return x % 2 > 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canAliceWin(int n) {
        int x = (21 - (int) ceil(sqrt(441 - n * 8))) / 2;
        return x % 2;
    }
};
```

```go [sol-Go]
func canAliceWin(n int) bool {
	x := (21 - int(math.Ceil(math.Sqrt(float64(441-n*8))))) / 2
	return x%2 > 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
