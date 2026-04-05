枚举 $a=1,2,\ldots$ 和 $b=a,a+1,\ldots$，用哈希表统计 $10^9$ 以内的 $a^3+b^3$ 的出现次数。

把出现次数大于 $1$ 的 $a^3+b^3$ 添加到答案中。

代码实现时，可以先预处理所有好整数 $\textit{goodIntegers}$，然后在 $\textit{goodIntegers}$ 中**二分查找**最后一个 $\le n$ 的数的下标 $i$，那么 $\textit{goodIntegers}$ 的前缀 $[0,i]$ 就是答案。

也可以二分查找第一个 $> n$ 的数的下标 $i$，那么 $\textit{goodIntegers}$ 的前缀 $[0,i-1]$ 就是答案。

关于二分查找的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

由于 $b\ge a$，所以当 $2a^3 > 10^9$ 时，跳出外层循环。

```py [sol-Python3]
MX = 1_000_000_000
cnt = defaultdict(int)
a = 1
while a * a * a * 2 <= MX:
    b = a
    while (s := a * a * a + b * b * b) <= MX:
        cnt[s] += 1
        b += 1
    a += 1

good_integers = sorted(x for x, c in cnt.items() if c > 1)

class Solution:
    def findGoodIntegers(self, n: int) -> List[int]:
        i = bisect_right(good_integers, n)
        return good_integers[:i]
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_000_000;
    private static final List<Integer> goodIntegers = new ArrayList<>();
    private static boolean initialized = false;

    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Map<Integer, Integer> cnt = new HashMap<>();
        for (int a = 1; a * a * a <= MX / 2; a++) {
            for (int b = a; a * a * a + b * b * b <= MX; b++) {
                cnt.merge(a * a * a + b * b * b, 1, Integer::sum);
            }
        }

        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            if (e.getValue() > 1) {
                goodIntegers.add(e.getKey());
            }
        }

        Collections.sort(goodIntegers);
    }

    public List<Integer> findGoodIntegers(int n) {
        int i = Collections.binarySearch(goodIntegers, n + 1);
        if (i < 0) {
            i = ~i; // 见 binarySearch 的源码
        }
        return goodIntegers.subList(0, i);
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1'000'000'000;
vector<int> good_integers;

auto init = [] {
    map<int, int> cnt;
    for (long long a = 1; a * a * a <= MX / 2; a++) {
        for (long long b = a; a * a * a + b * b * b <= MX; b++) {
            cnt[a * a * a + b * b * b]++;
        }
    }

    for (auto& [x, c] : cnt) {
        if (c > 1) {
            good_integers.push_back(x);
        }
    }
    return 0;
}();

class Solution {
public:
    vector<int> findGoodIntegers(int n) {
        auto end = ranges::upper_bound(good_integers, n);
        return vector(good_integers.begin(), end);
    }
};
```

```go [sol-Go]
var goodIntegers []int // 1554 个

func init() {
	const mx = 1_000_000_000
	cnt := map[int]int{}
	for a := 1; a*a*a <= mx/2; a++ {
		for b := a; a*a*a+b*b*b <= mx; b++ {
			cnt[a*a*a+b*b*b]++
		}
	}

	for x, c := range cnt {
		if c > 1 {
			goodIntegers = append(goodIntegers, x)
		}
	}

	slices.Sort(goodIntegers)
}

func findGoodIntegers(n int) []int {
	i := sort.SearchInts(goodIntegers, n+1)
	return goodIntegers[:i]
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(G)$ 或 $\mathcal{O}(\log G)$，其中 $G=1554$ 是好整数的个数。如果获取子数组是 $\mathcal{O}(1)$ 时间的（例如 Go 语言），则瓶颈在二分查找上。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
