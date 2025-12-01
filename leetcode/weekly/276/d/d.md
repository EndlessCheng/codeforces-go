## 方法一：二分答案

如果可以让 $n$ 台电脑同时运行 $x$ 分钟，那么必然可以同时运行 $x-1,x-2,\ldots$ 分钟（要求更宽松）；如果无法让 $n$ 台电脑同时运行 $x$ 分钟，那么必然无法同时运行 $x+1,x+2,\ldots$ 分钟（要求更苛刻）。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

假设可以让 $n$ 台电脑同时运行 $x$ 分钟，那么对于电量大于 $x$ 的电池，其只能被使用 $x$ 分钟，因此每个电池的使用时间至多为 $\min(\textit{batteries}[i], x)$。累加所有电池的使用时间，记作 $\textit{sum}$。那么要让 $n$ 台电脑同时运行 $x$ 分钟，**必要条件**是 $n\cdot x\le \textit{sum}$。

下面证明该条件也是**充分**的，即如果 $n\cdot x\le \textit{sum}$ 成立，那么一定存在一种安排电池的方式，可以让 $n$ 台电脑同时运行 $x$ 分钟。

构造方法如下：

对于电量 $\ge x$ 的电池，我们可以让其给一台电脑供电 $x$ 分钟。由于一个电池不能同时给多台电脑供电，因此该电池若给一台电脑供电 $x$ 分钟，那它就不能用于其他电脑了（我们只有 $x$ 分钟，无论电池有多少电，电池只能用 $x$ 分钟，剩余的电再多也没机会用）。我们可以将所有电量 $\ge x$ 的电池各给一台电脑供电。

对于其余电池，设其电量和为 $\textit{sum}'$，剩余 $n'$ 台电脑未被供电。我们可以随意选择剩下的电池，供给剩余的第一台电脑（用完一个电池就换下一个电池），多余的电池电量与剩下的电池一起供给剩余的第二台电脑，依此类推。注意由于这些电池的电量均小于 $x$，按照这种做法是不会出现同一个电池在同一时间供给多台电脑的（如果某个电池供给了两台电脑，可以将这个电池的供电时间**划分到第一台电脑的末尾和第二台电脑的开头**，因为电池电量小于 $x$，所以时间不会重叠）。注：想象我们把其余电池连成一条线，每隔 $x$ 切一刀，每段分给一台电脑。

由于 $\textit{sum}'=\textit{sum}-(n-n')\cdot x$，结合 $n\cdot x\le \textit{sum}$ 可以得到 $n'\cdot x\le \textit{sum}'$，按照上述供电方案（用完一个电池就换下一个电池），这 $n'$ 台电脑可以运行至少 $x$ 分钟。充分性得证。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。不运行任何电脑，一定满足要求。
- 开区间右端点初始值：平均值加一，即 $\left\lfloor\dfrac{\sum \textit{batteries}[i]}{n}\right\rfloor + 1$。一定无法满足要求。

```py [sol-Python3]
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        l, r = 0, sum(batteries) // n + 1
        while l + 1 < r:
            x = (l + r) // 2
            if n * x <= sum(min(b, x) for b in batteries):
                l = x
            else:
                r = x
        return l
```

```py [sol-Python3 库函数]
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        r = sum(batteries) // n
        # 二分找最小的不满足要求的 x+1，那么最大的满足要求的就是 x
        check = lambda x: n * (x + 1) > sum(min(b, x + 1) for b in batteries)
        return bisect_left(range(r), True, key=check)
```

```java [sol-Java]
class Solution {
    public long maxRunTime(int n, int[] batteries) {
        long tot = 0;
        for (int b : batteries) {
            tot += b;
        }

        long l = 0;
        long r = tot / n + 1;
        while (l + 1 < r) {
            long x = l + (r - l) / 2;
            long sum = 0;
            for (int b : batteries) {
                sum += Math.min(b, x);
            }
            if (n * x <= sum) {
                l = x;
            } else {
                r = x;
            }
        }
        return l;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxRunTime(int n, vector<int>& batteries) {
        long long tot = reduce(batteries.begin(), batteries.end(), 0LL);
        long long l = 0, r = tot / n + 1;
        while (l + 1 < r) {
            long long x = l + (r - l) / 2;
            long long sum = 0;
            for (long long b : batteries) {
                sum += min(b, x);
            }
            (n * x <= sum ? l : r) = x;
        }
        return l;
    }
};
```

```go [sol-Go]
func maxRunTime(n int, batteries []int) int64 {
	tot := 0
	for _, b := range batteries {
		tot += b
	}

	return int64(sort.Search(tot/n, func(x int) bool {
		x++
		sum := 0
		for _, b := range batteries {
			sum += min(b, x)
		}
		return n*x > sum
	}))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log (S/n))$，其中 $m$ 是 $\textit{batteries}$ 的长度，$S$ 是 $\textit{batteries}$ 的元素和。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：排序 + 贪心

受解法一的启发，我们可以得出如下贪心策略：

记电池电量和为 $\textit{sum}$，则理论上至多可以供电 $x=\Big\lfloor\dfrac{\textit{sum}}{n}\Big\rfloor$ 分钟。我们对电池电量从大到小排序，然后从电量最大的电池开始遍历：

- 若该电池电量超过 $x$，则将其供给一台电脑，问题缩减为 $n-1$ 台电脑的子问题。
- 若该电池电量不超过 $x$，则其余电池的电量均不超过 $x$，此时有

  $$
  n\cdot x=n\cdot\Big\lfloor\dfrac{\textit{sum}}{n}\Big\rfloor \le \textit{sum}
  $$

  根据解法一的结论，这些电池可以给 $n$ 台电脑供电 $x$ 分钟。

由于随着问题规模减小，$x$ 不会增加，因此若遍历到一个电量不超过 $x$ 的电池时，可以直接返回 $x$ 作为答案。

```py [sol-Python3]
class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        batteries.sort(reverse=True)
        s = sum(batteries)
        for b in batteries:
            if b <= s // n:
                return s // n
            s -= b
            n -= 1
```

```java [sol-Java]
class Solution {
    public long maxRunTime(int n, int[] batteries) {
        Arrays.sort(batteries);

        long sum = 0;
        for (int b : batteries) {
            sum += b;
        }

        for (int i = batteries.length - 1; ; i--) {
            if (batteries[i] <= sum / n) {
                return sum / n;
            }
            sum -= batteries[i];
            n--;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxRunTime(int n, vector<int>& batteries) {
        ranges::sort(batteries, greater());
        long long sum = reduce(batteries.begin(), batteries.end(), 0LL);
        for (int i = 0; ; i++) {
            if (batteries[i] <= sum / n) {
                return sum / n;
            }
            sum -= batteries[i];
            n--;
        }
    }
};
```

```go [sol-Go]
func maxRunTime(n int, batteries []int) int64 {
	slices.Sort(batteries)
	sum := 0
	for _, b := range batteries {
		sum += b
	}
	for i := len(batteries) - 1; ; i-- {
		if batteries[i] <= sum/n {
			return int64(sum / n)
		}
		sum -= batteries[i]
		n--
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 是 $\textit{batteries}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

1. 二分题单的「**§2.2 求最大**」。
2. 贪心题单的「**§1.1 从最小/最大开始贪心**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
