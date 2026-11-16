## 方法一：模拟

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

- 时间复杂度：$\mathcal{O}(C)$，其中 $C=10$。
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
