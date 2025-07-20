逐个遍历 $n$ 的每个数位，统计能整除 $n$ 的数位个数。

代码实现时，可以不用转成字符串处理，而是不断取最低位（模 $10$），去掉最低位（除以 $10$），直到数字为 $0$。

例如 $n=123$：

1. 初始化 $x=n$。
2. 通过 $x\bmod 10$ 取到个位数 $3$，然后把 $x$ 除以 $10$（下取整），得到 $x=12$。
3. 再次 $x\bmod 10$ 取到十位数 $2$，然后把 $x$ 除以 $10$（下取整），得到 $x=1$。
4. 最后 $x\bmod 10$ 取到百位数 $1$，然后把 $x$ 除以 $10$（下取整），得到 $x=0$。此时完成了遍历 $n$ 的每个数位，退出循环。

在这个过程中，计算数位和 $s$ 和数位乘积 $m$。最后判断 $n$ 是否为 $s+m$ 的倍数。注意 $n>0$，所以 $s>0$，所以不会发生除 $0$ 异常。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1tbg8z3EaP/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def checkDivisibility(self, n: int) -> bool:
        s, m = 0, 1
        x = n
        while x:
            x, d = divmod(x, 10)
            s += d
            m *= d
        return n % (s + m) == 0
```

```java [sol-Java]
class Solution {
    public boolean checkDivisibility(int n) {
        int s = 0, m = 1;
        for (int x = n; x > 0; x /= 10) {
            int d = x % 10;
            s += d;
            m *= d;
        }
        return n % (s + m) == 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkDivisibility(int n) {
        int s = 0, m = 1;
        for (int x = n; x > 0; x /= 10) {
            int d = x % 10;
            s += d;
            m *= d;
        }
        return n % (s + m) == 0;
    }
};
```

```go [sol-Go]
func checkDivisibility(n int) bool {
	s, m := 0, 1
	for x := n; x > 0; x /= 10 {
		d := x % 10
		s += d
		m *= d
	}
	return n%(s+m) == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [2520. 统计能整除数字的位数](https://leetcode.cn/problems/count-the-digits-that-divide-a-number/)

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
