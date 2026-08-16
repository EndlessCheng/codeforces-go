先把 $k$ 减少 $1$，改成从 $0$ 开始，方便计算。

设答案在正整数 $x$ 中。我们先求出 $x$ 的十进制长度。

十进制长为 $\textit{length}$ 的正整数，最高位不能填 $0$，有 $9$ 种填法，其余 $\textit{length}-1$ 位有 $10$ 种填法。根据乘法原理，十进制长为 $\textit{length}$ 的正整数个数为

$$
\textit{cnt} = 9 \cdot 10^{\textit{length}-1}
$$

这 $\textit{cnt}$ 个数的十进制长度之和为

$$
\textit{cnt}\cdot \textit{length}
$$

枚举 $\textit{length}=1,2,3,\ldots$ 把 $k$ 减少 $\textit{cnt}\cdot \textit{length}$，直到 $\textit{cnt}\cdot \textit{length} > k$ 为止。

$x$ 是第 $\left\lfloor\dfrac{k}{\textit{length}}\right\rfloor$ 个（从 $0$ 开始）十进制长为 $\textit{length}$ 的正整数，即

$$
x = 10^{\textit{length}-1} + \left\lfloor\dfrac{k}{\textit{length}}\right\rfloor
$$

如果 $\left\lfloor\dfrac{x}{10}\right\rfloor$ 是奇数，$x$ 所在块是递减顺序，需要把 $x$ 的个位数 $d = x\bmod 10$ 替换成 $9 - d$，即 $x$ 改成

$$
(x - d) + (9 - d) = x+9-2d
$$

最后，答案在 $x$ 的从高到低第 $k\bmod \textit{length}$（从 $0$ 开始）个数字中。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def kthDigit(self, k: int) -> int:
        k -= 1  # k 改成从 0 开始，方便计算

        # 十进制长为 length 的正整数有 cnt = 9 * 10**(length-1) 个
        cnt, length = 9, 1
        while cnt * length <= k:
            k -= cnt * length  # 这里减小了 k
            cnt *= 10
            length += 1

        # 答案在正整数 x 中
        x = cnt // 9 + k // length
        if x // 10 % 2:
            # 改成递减顺序，例如 10 变成 19，11 变成 18 ……
            x += 9 - x % 10 * 2

        # 计算 x 从高到低第 k%length（从 0 开始）个数字
        tail = 10 ** (length - k % length - 1)
        return x // tail % 10
```

```java [sol-Java]
class Solution {
    public int kthDigit(long k) {
        k--; // k 改成从 0 开始，方便计算

        // 十进制长为 length 的正整数有 cnt = 9 * 10^(length-1) 个
        long cnt = 9;
        int length = 1;
        while (cnt * length <= k) {
            k -= cnt * length; // 这里减小了 k
            cnt *= 10;
            length++;
        }

        // 答案在正整数 x 中
        long x = cnt / 9 + k / length;
        if (x / 10 % 2 > 0) {
            // 改成递减顺序，例如 10 变成 19，11 变成 18 ……
            x += 9 - x % 10 * 2;
        }

        // 计算 x 从高到低第 k%length（从 0 开始）个数字
        for (int i = 0; i < length - k % length - 1; i++) {
            x /= 10;
        }
        return (int) (x % 10);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int kthDigit(long long k) {
        k--; // k 改成从 0 开始，方便计算

        // 十进制长为 length 的正整数有 cnt = 9 * 10^(length-1) 个
        long long cnt = 9;
        int length = 1;
        while (cnt * length <= k) {
            k -= cnt * length; // 这里减小了 k
            cnt *= 10;
            length++;
        }

        // 答案在正整数 x 中
        long long x = cnt / 9 + k / length;
        if (x / 10 % 2) {
            // 改成递减顺序，例如 10 变成 19，11 变成 18 ……
            x += 9 - x % 10 * 2;
        }

        // 计算 x 从高到低第 k%length（从 0 开始）个数字
        for (int i = 0; i < length - k % length - 1; i++) {
            x /= 10;
        }
        return x % 10;
    }
};
```

```go [sol-Go]
func kthDigit(K int64) int {
	k := int(K - 1) // k 改成从 0 开始，方便计算

	// 十进制长为 length 的正整数有 cnt = 9 * 10^(length-1) 个
	cnt, length := 9, 1
	for cnt*length <= k {
		k -= cnt * length // 这里减小了 k
		cnt *= 10
		length++
	}

	// 答案在正整数 x 中
	x := cnt/9 + k/length
	if x/10%2 > 0 {
		// 改成递减顺序，例如 10 变成 19，11 变成 18 ……
		x += 9 - x%10*2
	}

	// 计算 x 从高到低第 k%length（从 0 开始）个数字
	for range length - k%length - 1 {
		x /= 10
	}
	return x % 10
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log k)$。注：更精细的复杂度分析需要用到 Lambert W 函数，但结论是一样的，时间复杂度为 $\Theta(\log k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [400. 第 N 位数字](https://leetcode.cn/problems/nth-digit/)

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
