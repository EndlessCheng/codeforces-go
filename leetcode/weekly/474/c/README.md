这题和 [2513. 最小化两个数组中的最大值](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/) 是一样的，可以看 [我的题解](https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/solutions/2031827/er-fen-da-an-by-endlesscheng-y8fp/)。

下面介绍另外一种思考方式，目标是推导出方法二的结论。

## 方法一：二分答案

### 转化

如果能用 $5$ 小时完成送货，那么用 $6$ 小时或者更长时间，也能完成送货。

如果无法用 $5$ 小时完成送货，那么用 $4$ 小时或者更短时间，也无法完成送货。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定时限 $t$，能否在 $t$ 小时内完成送货？

如果可以完成，说明答案 $\le t$，否则答案 $> t$。

### 思路（具体构造方案）

两台无人机分别称作 $A$ 和 $B$。

$[1,t]$ 一共有 $t$ 个小时。按照无人机是否在充电，这 $t$ 个小时可以分为如下四类：

- $A$ 和 $B$ 都在充电：如果时刻是 $r_1$ 和 $r_2$ 的倍数，即 $L = \text{lcm}(r_1,r_2)$ 的倍数，那么 $A$ 和 $B$ 都在充电。这有 $\left\lfloor\dfrac{t}{L}\right\rfloor$ 个小时。
- 只能用 $A$：即 $A$ 可用，但 $B$ 在充电的时刻。$B$ 在充电的时刻有 $\left\lfloor\dfrac{t}{r_2}\right\rfloor$ 个，减去二者都在充电的时刻 $\left\lfloor\dfrac{t}{L}\right\rfloor$，得到只能用 $A$ 的时长 $t_A = \left\lfloor\dfrac{t}{r_2}\right\rfloor - \left\lfloor\dfrac{t}{L}\right\rfloor$。
- 只能用 $B$：即 $B$ 可用，但 $A$ 在充电的时刻。$A$ 在充电的时刻有 $\left\lfloor\dfrac{t}{r_1}\right\rfloor$ 个，减去二者都在充电的时刻 $\left\lfloor\dfrac{t}{L}\right\rfloor$，得到只能用 $B$ 的时长 $t_B = \left\lfloor\dfrac{t}{r_1}\right\rfloor - \left\lfloor\dfrac{t}{L}\right\rfloor$。
- $A$ 和 $B$ 都可用：用总时长减去上面三种情况，得到 $A$ 和 $B$ 都可用的时长 $t_C = t - \left\lfloor\dfrac{t}{r_1}\right\rfloor - \left\lfloor\dfrac{t}{r_2}\right\rfloor + \left\lfloor\dfrac{t}{L}\right\rfloor$。注：也可以用容斥原理得到这个式子。

$A$ 需要消耗 $d_1$ 小时，我们可以优先消耗「只能用 $A$」的时间，然后再消耗「$A$ 和 $B$ 都可用」的时间。所以必须满足

$$
d_1 \le t_A + t_C = t - \left\lfloor\dfrac{t}{r_1}\right\rfloor
$$

$B$ 需要消耗 $d_2$ 小时，我们可以优先消耗「只能用 $B$」的时间，然后再消耗「$A$ 和 $B$ 都可用」的时间。所以必须满足

$$
d_2 \le t_B + t_C = t - \left\lfloor\dfrac{t}{r_2}\right\rfloor
$$

$A$ 和 $B$ 一共需要消耗 $d_1+d_2$ 小时，这不能超过三类可用时长之和，或者说 $t$ 减去 $A$ 和 $B$ 都在充电的时长，即

$$
d_1+d_2 \le t_A + t_B + t_C =  t - \left\lfloor\dfrac{t}{L}\right\rfloor
$$

上面三个不等式是**必要条件**。下面证明其**充分性**，也就是说，如果上面三个不等式成立，那么上述方案一定能完成送货。

$d_1$ 消耗掉 $t_A$ 后，剩余值为 $d'_1 = \max(d_1-t_A, 0)$。

$d_2$ 消耗掉 $t_B$ 后，剩余值为 $d'_2 = \max(d_2-t_B, 0)$。

我们需要证明 $d'_1 + d'_2 \le t_C$。

这可以按照 $d_1\le t_A$ 是否成立，$d_2\le t_B$ 是否成立，分四种情况讨论。

比如 $d_1 > t_A$ 且 $d_2 > t_B$ 的情况，那么 $d'_1 + d'_2 \le t_C$ 相当于 $d_1-t_A + d_2-t_B\le t_C$，根据第三个不等式，这是成立的。其他情况类似。故充分性得证。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$d_1+d_2-1$。一定无法完成送货。
- 开区间右端点初始值：$2(d_1+d_2)-1$。即使是 $r_1=r_2=2$ 的最坏情况（有一半时间无法送货），也一定能完成送货。

> 对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

[本题视频讲解](https://www.bilibili.com/video/BV1MgyfBoEuX/?t=5m53s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumTime(self, d: List[int], r: List[int]) -> int:
        d1, d2 = d
        r1, r2 = r
        l = lcm(r1, r2)

        def check(t: int) -> bool:
            return d1 <= t - t // r1 and d2 <= t - t // r2 and d1 + d2 <= t - t // l

        left = d1 + d2 - 1
        right = (d1 + d2) * 2 - 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                right = mid
            else:
                left = mid
        return right
```

```py [sol-Python3 库函数]
class Solution:
    def minimumTime(self, d: List[int], r: List[int]) -> int:
        d1, d2 = d
        r1, r2 = r
        l = lcm(r1, r2)

        def check(t: int) -> bool:
            return d1 <= t - t // r1 and d2 <= t - t // r2 and d1 + d2 <= t - t // l

        # 库函数是左闭右开区间
        left = d1 + d2
        right = (d1 + d2) * 2 - 1
        return bisect_left(range(right), True, lo=left, key=check)
```

```java [sol-Java]
class Solution {
    public long minimumTime(int[] d, int[] r) {
        int d1 = d[0], d2 = d[1];
        int r1 = r[0], r2 = r[1];
        int l = lcm(r1, r2);

        long left = d1 + d2 - 1;
        long right = (d1 + d2) * 2L - 1;
        while (left + 1 < right) {
            long mid = left + (right - left) / 2;
            if (check(mid, d1, d2, r1, r2, l)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(long t, int d1, int d2, int r1, int r2, int l) {
        return d1 <= t - t / r1 && d2 <= t - t / r2 && d1 + d2 <= t - t / l;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumTime(vector<int>& d, vector<int>& r) {
        int d1 = d[0], d2 = d[1];
        int r1 = r[0], r2 = r[1];
        int l = lcm(r1, r2); // 注：如果计算结果是 long long，可以把 r1 r2 的类型改成 long long 避免溢出

        auto check = [&](long long t) -> bool {
            return d1 <= t - t / r1 && d2 <= t - t / r2 && d1 + d2 <= t - t / l;
        };

        long long left = d1 + d2 - 1;
        long long right = (d1 + d2) * 2LL - 1;
        while (left + 1 < right) {
            long long mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minimumTime(d, r []int) int64 {
	d1, d2 := d[0], d[1]
	r1, r2 := r[0], r[1]
	l := lcm(r1, r2)

	// 库函数是左闭右开区间
	left := d1 + d2
	right := (d1+d2)*2 - 1
	ans := left + sort.Search(right-left, func(t int) bool {
		t += left
		return d1 <= t-t/r1 && d2 <= t-t/r2 && d1+d2 <= t-t/l
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \max(r_1,r_2) + \log (d_1+d_2))$。计算 $\text{lcm}$ 需要 $\mathcal{O}(\log \max(r_1,r_2))$ 时间，二分需要 $\mathcal{O}(\log (d_1+d_2))$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学公式

对于给定的一对 $(d,r)$，问题相当于计算最小的 $t$，满足

$$
t - \left\lfloor\dfrac{t}{r}\right\rfloor \ge d
$$

直接解这个不等式不好做，考虑其实际含义。对于给定的 $(d,r)$，最小的 $t$ 等于 $d$ 加上充电时长。所以我们要关注的其实是：需要额外花费多少小时用来充电？

每经过 $r$ 小时，就需要花 $1$ 小时充电（最后一段不需要充电）。或者说，每 $r$ 小时只能工作 $r-1$ 小时。

- 比如 $d=5$，$r=3$，那么分成三段（$d=2+2+1$），除了最后一段不需要充电，前两段都需要充电。
- 比如 $d=6$，$r=3$，那么分成三段（$d=2+2+2$），除了最后一段不需要充电，前两段都需要充电。
- 比如 $d=7$，$r=3$，那么分成四段（$d=2+2+2+1$），除了最后一段不需要充电，前三段都需要充电。
- 比如 $d=8$，$r=3$，那么分成四段（$d=2+2+2+2$），除了最后一段不需要充电，前三段都需要充电。

一般地，我们分成

$$
\left\lceil\dfrac{d}{r-1}\right\rceil = \left\lfloor\dfrac{d-1}{r-1}\right\rfloor + 1
$$

段。见 [上取整下取整转换公式的证明](https://zhuanlan.zhihu.com/p/1890356682149838951)。

其中最后一段不充电，所以充电时长为

$$
\left\lfloor\dfrac{d-1}{r-1}\right\rfloor
$$

所以最小的 $t$ 为

$$
d + \left\lfloor\dfrac{d-1}{r-1}\right\rfloor
$$

方法一的三个不等式，算出三个 $t$，其中最大的 $t$ 可以同时满足三个不等式。

```py [sol-Python3]
class Solution:
    def minimumTime(self, d: List[int], r: List[int]) -> int:
        d1, d2 = d
        r1, r2 = r
        l = lcm(r1, r2)

        def f(d: int, r: int) -> int:
            return d + (d - 1) // (r - 1)

        return max(f(d1, r1), f(d2, r2), f(d1 + d2, l))
```

```java [sol-Java]
class Solution {
    public long minimumTime(int[] d, int[] r) {
        int d1 = d[0], d2 = d[1];
        int r1 = r[0], r2 = r[1];
        int l = lcm(r1, r2);
        return Math.max(Math.max(f(d1, r1), f(d2, r2)), f(d1 + d2, l));
    }

    private long f(int d, int r) {
        return (long) d + (d - 1) / (r - 1);
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }
}
```

```cpp [sol-C++]
class Solution {
    long long f(int d, int r) {
        return 1LL * d + (d - 1) / (r - 1);
    }

public:
    long long minimumTime(vector<int>& d, vector<int>& r) {
        int d1 = d[0], d2 = d[1];
        int r1 = r[0], r2 = r[1];
        int l = lcm(r1, r2); // 注：如果计算结果是 long long，可以把 r1 r2 的类型改成 long long 避免溢出
        return max({f(d1, r1), f(d2, r2), f(d1 + d2, l)});
    }
};
```

```go [sol-Go]
func f(d, r int) int {
	return d + (d-1)/(r-1)
}

func minimumTime(d, r []int) int64 {
	d1, d2 := d[0], d[1]
	r1, r2 := r[0], r[1]
	l := lcm(r1, r2)
	return int64(max(f(d1, r1), f(d2, r2), f(d1+d2, l)))
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \max(r_1,r_2))$。计算 $\text{lcm}$ 需要 $\mathcal{O}(\log \max(r_1,r_2))$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面二分题单的「**§2.1 求最小**」。

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
