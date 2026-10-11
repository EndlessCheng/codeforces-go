设斐波那契数列为 $F$。

连续三项斐波那契数之和为

$$
F_{i-2} + F_{i-1} + F_i = F_i + F_i = 2F_i
$$

所以 $n$ 必须是偶数。

问题变成判断 $\dfrac{n}{2}$ 是否为斐波那契数。

预处理 $[1,10^9]$ 内的斐波那契数，添加到一个哈希集合后，这样就可以 $\mathcal{O}(1)$ 判断了。

下面预处理斐波那契数的算法，见 [509. 斐波那契数](https://leetcode.cn/problems/fibonacci-number/)，[我的题解](https://leetcode.cn/problems/fibonacci-number/solutions/3811875/san-chong-xie-fa-shu-zu-di-tui-kong-jian-3lwj/)。

[本题视频讲解](https://www.bilibili.com/video/BV1wnpj6VEtP/)，欢迎点赞关注~

```py [sol-Python3]
MX = 1_000_000_000
fib_set = set()
a, b = 0, 1
while b < MX:
    a, b = b, a + b
    fib_set.add(b)

class Solution:
    def threeFibonacciSum(self, n: int) -> bool:
        return n % 2 == 0 and n // 2 in fib_set
```

```java [sol-Java]
class Solution {
    private static final Set<Integer> fibSet = new HashSet<>();

    static {
        int a = 0;
        int b = 1;
        while (b < 1e9) {
            int x = a + b;
            a = b;
            b = x;
            fibSet.add(b);
        }
    }

    public boolean threeFibonacciSum(int n) {
        return n % 2 == 0 && fibSet.contains(n / 2);
    }
}
```

```cpp [sol-C++]
unordered_set<int> fib_set;

int init = [] {
    int a = 0, b = 1;
    while (b < 1e9) {
        tie(a, b) = pair(b, a + b);
        fib_set.insert(b);
    }
    return 0;
}();

class Solution {
public:
    bool threeFibonacciSum(int n) {
        return n % 2 == 0 && fib_set.contains(n / 2);
    }
};
```

```go [sol-Go]
var fibSet = map[int]bool{}

func init() {
	a, b := 0, 1
	for b < 1e9 {
		a, b = b, a+b
		fibSet[b] = true
	}
}

func threeFibonacciSum(n int) bool {
	return n%2 == 0 && fibSet[n/2]
}
```

#### 复杂度分析

由于斐波那契数列是指数增长的，所以预处理的时间和空间为 $\mathcal{O}(\log U)$，其中 $U=10^9$。

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
