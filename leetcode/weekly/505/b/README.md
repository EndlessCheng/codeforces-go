对于 `x` 和 `x >> 1` 这两个二进制数，同一个比特位上的数字，对应着 $x$ 中一对相邻比特位上的数字。所以计算 `x & (x >> 1)` 等价于计算 $x$ 的所有相邻比特位的 `&`。所有相邻的 $11$ 变成 $1$，其余变成 $0$。

例如二进制数 $x = 10011001110$，`x & (x >> 1)` 为 $1000110$。

$x$ 不含 $11$，等价于 `n & (n >> 1)` 等于 $0$。

计算 $x$ 的成本，我们可以去掉 $x$ 中的一个比特位（最低位还是最高位都可以），计算剩余二进制数的成本，这是一个规模更小的子问题，所以可以 DP。

[本题视频讲解](https://www.bilibili.com/video/BV1yfEx6YEBx/)，欢迎点赞关注~

```py [sol-Python3]
cost = [0] * (1 << 12)
for x in range(1, len(cost)):
    if x & (x >> 1):  # 有两个连续的 1
        cost[x] = inf  # 不合法
    else:
        # 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
        i = x.bit_length() - 1
        cost[x] = cost[x ^ (1 << i)] + i

class Solution:
    def generateValidStrings(self, n: int, k: int) -> List[str]:
        ans = []
        s = [''] * n
        for x in range(1 << n):
            if cost[x] > k:
                continue
            for j in range(n):  # 注意左边是低位，右边是高位
                s[j] = str(x & 1)
                x >>= 1
            ans.append(''.join(s))
        return ans
```

```py [sol-Python3 库函数]
cost = [0] * (1 << 12)
for x in range(1, len(cost)):
    if x & (x >> 1):  # 有两个连续的 1
        cost[x] = inf  # 不合法
    else:
        # 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
        i = x.bit_length() - 1
        cost[x] = cost[x ^ (1 << i)] + i

class Solution:
    def generateValidStrings(self, n: int, k: int) -> List[str]:
        ans = []
        for x in range(1 << n):
            if cost[x] > k:
                continue
            s = bin(x)[2:].zfill(n)[::-1]  # 或者 format(x, f"0{n}b")[::-1]
            ans.append(s)
        return ans
```

```java [sol-Java]
class Solution {
    private static final int[] cost = new int[1 << 12];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        for (int x = 1; x < cost.length; x++) {
            if ((x & (x >> 1)) > 0) { // 有两个连续的 1
                cost[x] = Integer.MAX_VALUE; // 不满足要求
            } else {
                // 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
                cost[x] = cost[x & (x - 1)] + Integer.numberOfTrailingZeros(x);
            }
        }
    }

    public List<String> generateValidStrings(int n, int k) {
        List<String> ans = new ArrayList<>();
        char[] s = new char[n];
        for (int x = 0; x < (1 << n); x++) {
            if (cost[x] > k) {
                continue;
            }
            int y = x;
            for (int j = 0; j < n; j++) { // 注意左边是低位，右边是高位
                s[j] = (char) ('0' + (y & 1));
                y >>= 1;
            }
            ans.add(new String(s));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 1 << 12;
int cost[MX];

int init = [] {
    for (int x = 1; x < MX; x++) {
        if (x & (x >> 1)) { // 有两个连续的 1
            cost[x] = INT_MAX; // 不满足要求
        } else {
            // 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
            cost[x] = cost[x & (x - 1)] + countr_zero(1u * x);
        }
    }
    return 0;
}();

class Solution {
public:
    vector<string> generateValidStrings(int n, int k) {
        vector<string> ans;
        string s(n, 0);
        for (int x = 0; x < (1 << n); x++) {
            if (cost[x] > k) {
                continue;
            }
            int y = x;
            for (int j = 0; j < n; j++) { // 注意左边是低位，右边是高位
                s[j] = '0' + (y & 1);
                y >>= 1;
            }
            ans.push_back(s);
        }
        return ans;
    }
};
```

```go [sol-Go]
var cost [1 << 12]int

func init() {
	for x := 1; x < len(cost); x++ {
		if x&(x>>1) > 0 { // 有两个连续的 1
			cost[x] = math.MaxInt // 不合法
		} else {
			// 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
			cost[x] = cost[x&(x-1)] + bits.TrailingZeros(uint(x))
		}
	}
}

func generateValidStrings(n, k int) (ans []string) {
	s := make([]byte, n)
	for x, c := range cost[:1<<n] {
		if c > k {
			continue
		}
		for j := range s { // 注意左边是低位，右边是高位
			s[j] = '0' + byte(x&1)
			x >>= 1
		}
		ans = append(ans, string(s))
	}
	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n2^n)$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。返回值不计入。

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
