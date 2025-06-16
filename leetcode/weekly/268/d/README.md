由于数据范围很小，可以直接把 $k$ 在 $[2,9]$，$n$ 在 $[1,30]$ 中的所有答案预处理出来，这一共有 $8\cdot 30=240$ 个。

从小到大枚举回文数，构造方法如下：

- 枚举 $\textit{base}=10^0,10^1,10^2,\cdots$
- **枚举回文数的左半边**。
- 对于 $[\textit{base},10\cdot \textit{base})$ 范围内的每个数，将除了末尾数位的其余部分反转，拼到原数字末尾，这样可以生成奇数长度回文数。例如 $\textit{base}=10$，可以生成在 $[101,999]$ 中的所有回文数。
- 对于 $[\textit{base},10\cdot \textit{base})$ 范围内的每个数，将所有数位反转，拼到原数字末尾，这样可以生成偶数长度回文数。例如 $\textit{base}=10$，可以生成在 $[1001,9999]$ 中的所有回文数。

对每个回文数，判断其是否为 $k=2,3,\ldots,9$ 镜像数字。做法见 [9. 回文数](https://leetcode.cn/problems/palindrome-number/)，[我的题解](https://leetcode.cn/problems/palindrome-number/solutions/3682487/bi-guan-fang-ti-jie-shao-xun-huan-yi-ci-02nkc/)。也可以转成字符串再判断。

当每组 $k$ 镜像数字都计算出 $30$ 个时，结束枚举。

然后对每组 $k$ 镜像数字，计算这 $30$ 个数的前缀和，即为答案。

**注**：经计算，枚举回文数的左半边时，要从 $1$ 枚举到 $644545$，最后一个回文数为 $64454545446$。注意这个回文数超出了 $32$ 位整数范围。

```py [sol-Python3]
MAX_N = 30
ans = [[] for _ in range(10)]

# 力扣 9. 回文数
def is_k_palindrome(x: int, k: int) -> bool:
    if x % k == 0:
        return False
    rev = 0
    while rev < x // k:
        rev = rev * k + x % k
        x //= k
    return rev == x or rev == x // k

def do_palindrome(x: int) -> bool:
    done = True
    for k in range(2, 10):
        if len(ans[k]) < MAX_N and is_k_palindrome(x, k):
            ans[k].append(x)
        if len(ans[k]) < MAX_N:
            done = False
    if not done:
        return False
    for k in range(2, 10):
        ans[k] = list(accumulate(ans[k]))
    return True

def init() -> None:
    base = 1
    while True:
        # 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1][1:])
            if do_palindrome(x):
                return
        # 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1])
            if do_palindrome(x):
                return
        base *= 10
init()

class Solution:
    def kMirror(self, k: int, n: int) -> int:
        return ans[k][n - 1]
```

```java [sol-Java]
class Solution {
    private static final int MAX_N = 30;
    private static final List<Long>[] ans = new ArrayList[10];
    private static boolean initialized = false;

    // 这样写比 static block 快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.setAll(ans, i -> new ArrayList<>());
        for (int base = 1; ; base *= 10) {
            // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
            for (int i = base; i < base * 10; i++) {
                long x = i;
                for (int t = i / 10; t > 0; t /= 10) {
                    x = x * 10 + t % 10;
                }
                if (doPalindrome(x)) {
                    return;
                }
            }
            // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
            for (int i = base; i < base * 10; i++) {
                long x = i;
                for (int t = i; t > 0; t /= 10) {
                    x = x * 10 + t % 10;
                }
                if (doPalindrome(x)) {
                    return;
                }
            }
        }
    }

    private boolean doPalindrome(long x) {
        boolean done = true;
        for (int k = 2; k < 10; k++) {
            if (ans[k].size() < MAX_N && isKPalindrome(x, k)) {
                ans[k].add(x);
            }
            if (ans[k].size() < MAX_N) {
                done = false;
            }
        }
        if (!done) {
            return false;
        }

        for (int k = 2; k < 10; k++) {
            // 原地求前缀和
            List<Long> s = ans[k];
            for (int i = 1; i < MAX_N; i++) {
                s.set(i, s.get(i) + s.get(i - 1));
            }
        }
        return true;
    }

    // 力扣 9. 回文数
    private boolean isKPalindrome(long x, int k) {
        if (x % k == 0) {
            return false;
        }
        int rev = 0;
        while (rev < x / k) {
            rev = rev * k + (int) (x % k);
            x /= k;
        }
        return rev == x || rev == x / k;
    }

    public long kMirror(int k, int n) {
        init();
        return ans[k].get(n - 1);
    }
}
```

```cpp [sol-C++]
const int MAX_N = 30;
vector<long long> ans[10];

// 力扣 9. 回文数
bool is_k_palindrome(long long x, int k) {
    if (x % k == 0) {
        return false;
    }
    int rev = 0;
    while (rev < x / k) {
        rev = rev * k + x % k;
        x /= k;
    }
    return rev == x || rev == x / k;
}

bool do_palindrome(long long x) {
    bool done = true;
    for (int k = 2; k < 10; k++) {
        if (ans[k].size() < MAX_N && is_k_palindrome(x, k)) {
            ans[k].push_back(x);
        }
        if (ans[k].size() < MAX_N) {
            done = false;
        }
    }
    if (!done) {
        return false;
    }

    for (int k = 2; k < 10; k++) {
        // 原地求前缀和
        partial_sum(ans[k].begin(), ans[k].end(), ans[k].begin());
    }
    return true;
}

auto init = []() {
    for (int base = 1; ; base *= 10) {
        // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for (int i = base; i < base * 10; i++) {
            long long x = i;
            for (int t = i / 10; t > 0; t /= 10) {
                x = x * 10 + t % 10;
            }
            if (do_palindrome(x)) {
                return 0;
            }
        }
        // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for (int i = base; i < base * 10; i++) {
            long long x = i;
            for (int t = i; t > 0; t /= 10) {
                x = x * 10 + t % 10;
            }
            if (do_palindrome(x)) {
                return 0;
            }
        }
    }
}();

class Solution {
public:
    long long kMirror(int k, int n) {
        return ans[k][n - 1];
    }
};
```

```go [sol-Go]
const maxN = 30

var ans [10][]int

// 力扣 9. 回文数
func isKPalindrome(x, k int) bool {
	if x%k == 0 {
		return false
	}
	rev := 0
	for rev < x/k {
		rev = rev*k + x%k
		x /= k
	}
	return rev == x || rev == x/k
}

func doPalindrome(x int) bool {
	done := true
	for k := 2; k < 10; k++ {
		if len(ans[k]) < maxN && isKPalindrome(x, k) {
			ans[k] = append(ans[k], x)
		}
		if len(ans[k]) < maxN {
			done = false
		}
	}
	if !done {
		return false
	}

	for k := 2; k < 10; k++ {
		// 计算前缀和 
		for i := 1; i < maxN; i++ {
			ans[k][i] += ans[k][i-1]
		}
	}
	return true
}

func init() {
	for k := 2; k < 10; k++ {
		ans[k] = make([]int, 0, maxN) // 预分配空间
	}
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if doPalindrome(x) {
				return
			}
		}
		// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
		for i := base; i < base*10; i++ {
			x := i
			for t := i; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if doPalindrome(x) {
				return
			}
		}
	}
}

func kMirror(k, n int) int64 {
	return int64(ans[k][n-1])
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

见下面数学题单的「**§7.1 回文数**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
