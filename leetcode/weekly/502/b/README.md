定义 $f(n,k)$ 表示 $[0,n]$ 中的完全 $k$ 次幂的个数，那么答案为 

$$
f(r,k)-f(\ell-1,k)
$$

$x^k\le n$ 即 $x\le n^{1/k}$，所以 $x$ 的范围为 $[0,\lfloor n^{1/k} \rfloor]$，这有 $\lfloor n^{1/k} \rfloor + 1$ 个。

⚠**注意**：由于浮点运算有误差，可能 $n^{1/k} = 6$，但计算机算出来的是 $5.99999\cdots$，下取整后是 $5$。

设 $u = \lfloor n^{1/k} \rfloor$，我们可以额外验证 $(u+1)^k\le n$ 是否成立，如果成立则把 $u$ 加一。保险起见，计算 $(u+1)^k$ 可以用快速幂，原理见 [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)，[【图解】一张图秒懂快速幂！](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)

[本题视频讲解](https://www.bilibili.com/video/BV18gLE6VETZ/?t=33m30s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def f(self, n: int, k: int) -> int:
        if n < 0:
            return 0
        x = floor(n ** (1 / k))
        # 可能 x 的正确值是 6，但算出来的 x = floor(5.99999...) = 5
        if (x + 1) ** k <= n:  # 为避免浮点误差，这里用整数计算 pow
            x += 1
        return x + 1

    def countKthRoots(self, l: int, r: int, k: int) -> int:
        return self.f(r, k) - self.f(l - 1, k)
```

```java [sol-Java]
class Solution {
    public int countKthRoots(int l, int r, int k) {
        return f(r, k) - f(l - 1, k);
    }

    private int f(int n, int k) {
        if (n < 0) {
            return 0;
        }
        int x = (int) Math.pow(n, 1.0 / k);
        // 可能 x 的正确值是 6，但算出来的 x = int(5.99999...) = 5
        if (pow(x + 1, k) <= n) { // 为避免浮点误差，这里用整数计算 pow
            x++;
        }
        return x + 1;
    }

    // 50. Pow(x, n)
    private long pow(long x, int k) {
        long res = 1;
        for (; k > 0; k /= 2) {
            if (k % 2 > 0) {
                res *= x;
            }
            x *= x;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 50. Pow(x, n)
    long long qpow(long long x, int k) {
        long long res = 1;
        for (; k > 0; k /= 2) {
            if (k % 2 > 0) {
                res *= x;
            }
            x *= x;
        }
        return res;
    }

    int f(int n, int k) {
        if (n < 0) {
            return 0;
        }
        int x = pow(n, 1.0 / k);
        // 可能 x 的正确值是 6，但算出来的 x = int(5.99999...) = 5
        if (qpow(x + 1, k) <= n) { // 为避免浮点误差，这里用整数计算 pow
            x++;
        }
        return x + 1;
    }

public:
    int countKthRoots(int l, int r, int k) {
        return f(r, k) - f(l - 1, k);
    }
};
```

```go [sol-Go]
// 50. Pow(x, n)
func pow(x, k int) int {
	res := 1
	for ; k > 0; k /= 2 {
		if k%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}

func f(n, k int) int {
	if n < 0 {
		return 0
	}
	x := int(math.Pow(float64(n), 1/float64(k)))
	// 可能 x 的正确值是 6，但算出来的 x = int(5.99999...) = 5
	if pow(x+1, k) <= n { // 为避免浮点误差，这里用整数计算 pow
		x++
	}
	return x + 1
}

func countKthRoots(l, r, k int) int {
	return f(r, k) - f(l-1, k)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log k)$。瓶颈在计算快速幂上。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
