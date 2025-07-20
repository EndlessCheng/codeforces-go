用埃氏筛预处理 $[0,n-1]$ 中的每个数是不是质数。

累加质数下标的元素，减去非质数下标的元素，取绝对值，即为答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
MX = 100_000
is_prime = [False] * 2 + [True] * (MX - 2)  # 0 和 1 不是质数
for i in range(2, isqrt(MX) + 1):
    if is_prime[i]:
        for j in range(i * i, MX, i):
            is_prime[j] = False  # j 是质数 i 的倍数

class Solution:
    def splitArray(self, nums: List[int]) -> int:
        ans = sum(x if p else -x for x, p in zip(nums, is_prime))
        return abs(ans)
```

```java [sol-Java]
class Solution {
    private static final int MX = 100_000;
    private static final boolean[] isPrime = new boolean[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.fill(isPrime, true);
        isPrime[0] = isPrime[1] = false; // 0 和 1 不是质数
        for (int i = 2; i * i < MX; i++) {
            if (isPrime[i]) {
                for (int j = i * i; j < MX; j += i) {
                    isPrime[j] = false; // j 是质数 i 的倍数
                }
            }
        }
    }

    public long splitArray(int[] nums) {
        init();
        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            ans += isPrime[i] ? nums[i] : -nums[i];
        }
        return Math.abs(ans);
    }
}
```

```cpp [sol-C++]
const int MX = 100'000;
bool is_prime[MX];

auto init = [] {
    ranges::fill(is_prime, true);
    is_prime[0] = is_prime[1] = false; // 0 和 1 不是质数
    for (int i = 2; i * i < MX; i++) {
        if (is_prime[i]) {
            for (int j = i * i; j < MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();

class Solution {
public:
    long long splitArray(vector<int>& nums) {
        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            ans += is_prime[i] ? nums[i] : -nums[i];
        }
        return abs(ans);
    }
};
```

```go [sol-Go]
const mx = 100_000

var np = [mx]bool{true, true} // 0 和 1 不是质数

func init() {
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func splitArray(nums []int) (ans int64) {
	for i, x := range nums {
		if np[i] {
			ans -= int64(x)
		} else {
			ans += int64(x)
		}
	}
	if ans < 0 {
		ans = -ans
	}
	return
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§1.2 预处理质数（筛质数）**」。

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
