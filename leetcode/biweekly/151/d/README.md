## 什么时候返回空列表

由于相邻元素奇偶性不同，确定第一个元素填什么，后面元素的奇偶性就确定了。比如第一个数填的是偶数，那么后面元素一定是按照奇偶奇偶的顺序填。

$[1,n]$ 中有 $\left\lfloor n/2\right\rfloor$ 个偶数，$\left\lceil n/2\right\rceil$ 个奇数。

- 这些偶数有 $\left\lfloor n/2\right\rfloor!$ 个不同的排列。
- 这些奇数有 $\left\lceil n/2\right\rceil!$ 个不同的排列。

如果 $n$ 是奇数，那么只能按照奇偶奇偶的顺序填，由于奇偶位置互相独立，根据乘法原理可得方案数为 $\left\lfloor n/2\right\rfloor!\left\lceil n/2\right\rceil!$。

如果 $n$ 是偶数，那么可以按照奇偶奇偶的顺序填，也可以按照偶奇偶奇的顺序填，方案数为 $2\left\lfloor n/2\right\rfloor!\left\lceil n/2\right\rceil!$。

如果 $k$ 比上述方案数还大，返回空列表。

代码实现时，可以预处理 $f$ 数组，其中 $f[n] = \left\lfloor n/2\right\rfloor!\left\lceil n/2\right\rceil!$。这可以通过计算 $1,1,2,2,3,3,4,4,\cdots$ 的**前缀积**得到。

## 如何填数字

为方便计算，先把 $k$ 减一，也就是改成从 $0$ 开始。

看示例 1，**按照第一个数分组**，每一组的大小都是 $2$，也就是 $f[n-1]=f[3]=2$。

- 当 $k\in [0,1]$ 时，第一个数在第一组中，一定是 $1$。
- 当 $k\in [2,3]$ 时，第一个数在第二组中，一定是 $2$。
- 当 $k\in [4,5]$ 时，第一个数在第三组中，一定是 $3$。
- 当 $k\in [6,7]$ 时，第一个数在第四组中，一定是 $4$。

所以根据 $\left\lfloor\dfrac{k}{f[n-1]}\right\rfloor$ 的值，我们可以知道第一个数在第几组，从而确定第一个数填什么。再次强调，$k$ 是从 $0$ 开始的。

设 $k'=k\bmod f[n-1]$，问题变成计算 $n-1$ 个数的字典序第 $k'$ 小的交替排列。这是一个规模更小的子问题，可以用递归/迭代解决。实现细节见代码注释。

注意 $n$ 是偶数的情况，有奇偶奇偶、偶奇偶奇两种顺序，需要特殊处理第一个数怎么填。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m39bYiEVV/?t=10m23s)，欢迎点赞关注~

```py [sol-Python3]
# 预处理交替排列的方案数
f = [1]
i = 1
while f[-1] < 10 ** 15:
    f.append(f[-1] * i)
    f.append(f[-1] * i)
    i += 1

class Solution:
    def permute(self, n: int, k: int) -> List[int]:
        # k 改成从 0 开始，方便计算
        k -= 1
        if n < len(f) and k >= f[n] * (2 - n % 2):  # n 是偶数的时候，方案数乘以 2
            return []

        # cand 表示剩余未填入 ans 的数字
        # cand[0] 保存偶数，cand[1] 保存奇数
        cand = [list(range(2, n + 1, 2)), list(range(1, n + 1, 2))]

        ans = [0] * n
        parity = 1  # 当前要填入 ans 的数字的奇偶性
        for i in range(n):
            if n - 1 - i < len(f):
                # 比如示例 1，按照第一个数分组，每一组的大小都是 size=f[n-1-i]=f[3]=2
                # 知道 k 和 size 就知道我们要去哪一组（第 j 组）
                j, k = divmod(k, f[n - 1 - i])
                # n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
                if n % 2 == 0 and i == 0:
                    parity = 1 - j % 2
                    j //= 2
            else:
                j = 0  # n 很大的情况下，只能按照 1,2,3,... 的顺序填
            ans[i] = cand[parity].pop(j)
            parity ^= 1  # 下一个数的奇偶性
        return ans
```

```java [sol-Java]
class Solution {
    // 预处理交替排列的方案数
    private static final List<Long> f = new ArrayList<>();

    static {
        f.add(1L);
        for (int i = 1; f.getLast() < 1e15; i++) {
            f.add(f.getLast() * i);
            f.add(f.getLast() * i);
        }
    }

    public int[] permute(int n, long k) {
        // k 改成从 0 开始，方便计算
        k--;
        if (n < f.size() && k >= f.get(n) * (2 - n % 2)) { // n 是偶数的时候，方案数乘以 2
            return new int[0];
        }

        // cand 表示剩余未填入 ans 的数字
        // cand[0] 保存偶数，cand[1] 保存奇数
        List<Integer>[] cand = new ArrayList[2];
        cand[0] = new ArrayList<>();
        for (int i = 2; i <= n; i += 2) {
            cand[0].add(i);
        }
        cand[1] = new ArrayList<>();
        for (int i = 1; i <= n; i += 2) {
            cand[1].add(i);
        }

        int[] ans = new int[n];
        int parity = 1; // 当前要填入 ans[i] 的数的奇偶性
        for (int i = 0; i < n; i++) {
            int j = 0;
            if (n - 1 - i < f.size()) {
                // 比如示例 1，按照第一个数分组，每一组的大小都是 size=2
                // 知道 k 和 size 就知道我们要去哪一组
                long size = f.get(n - 1 - i);
                j = (int) (k / size); // 去第 j 组
                k %= size;
                // n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
                if (n % 2 == 0 && i == 0) {
                    parity = 1 - j % 2;
                    j /= 2;
                }
            } // else j=0，在 n 很大的情况下，只能按照 1,2,3,... 的顺序填
            ans[i] = cand[parity].remove(j);
            parity ^= 1; // 下一个数的奇偶性
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 预处理交替排列的方案数
vector<long long> f = {1};
int init = []() {
    for (int i = 1; f.back() < 1e15; i++) {
        f.push_back(f.back() * i);
        f.push_back(f.back() * i);
    }
    return 0;
}();

class Solution {
public:
    vector<int> permute(int n, long long k) {
        // k 改成从 0 开始，方便计算
        k--;
        if (n < f.size() && k >= f[n] * (2 - n % 2)) { // n 是偶数的时候，方案数乘以 2
            return {};
        }

        // cand 表示剩余未填入 ans 的数字
        // cand[0] 保存偶数，cand[1] 保存奇数
        vector<int> cand[2];
        for (int i = 2; i <= n; i += 2) {
            cand[0].push_back(i);
        }
        for (int i = 1; i <= n; i += 2) {
            cand[1].push_back(i);
        }

        vector<int> ans(n);
        int parity = 1; // 当前要填入 ans 的数字的奇偶性
        for (int i = 0; i < n; i++) {
            int j = 0;
            if (n - 1 - i < f.size()) {
                // 比如示例 1，按照第一个数分组，每一组的大小都是 size=2
                // 知道 k 和 size 就知道我们要去哪一组
                long long size = f[n - 1 - i];
                j = k / size; // 去第 j 组
                k %= size;
                // n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
                if (n % 2 == 0 && i == 0) {
                    parity = 1 - j % 2;
                    j /= 2;
                }
            } // else j=0，在 n 很大的情况下，只能按照 1,2,3,... 的顺序填
            ans[i] = cand[parity][j];
            cand[parity].erase(cand[parity].begin() + j);
            parity ^= 1; // 下一个数的奇偶性
        }
        return ans;
    }
};
```

```go [sol-Go]
// 预处理交替排列的方案数
var f = []int{1}

func init() {
	for i := 1; f[len(f)-1] < 1e15; i++ {
		f = append(f, f[len(f)-1]*i)
		f = append(f, f[len(f)-1]*i)
	}
}

func permute(n int, K int64) []int {
	// k 改成从 0 开始，方便计算
	k := int(K - 1)
	if n < len(f) && k >= f[n]*(2-n%2) { // n 是偶数的时候，方案数乘以 2
		return nil
	}

	// cand 表示剩余未填入 ans 的数字
	// cand[0] 保存偶数，cand[1] 保存奇数
	cand := [2][]int{}
	for i := 2; i <= n; i += 2 {
		cand[0] = append(cand[0], i)
	}
	for i := 1; i <= n; i += 2 {
		cand[1] = append(cand[1], i)
	}

	ans := make([]int, n)
	parity := 1 // 当前要填入 ans[i] 的数的奇偶性
	for i := range n {
		j := 0
		if n-1-i < len(f) {
			// 比如示例 1，按照第一个数分组，每一组的大小都是 size=2
			// 知道 k 和 size 就知道我们要去哪一组
			size := f[n-1-i]
			j = k / size // 去第 j 组
			k %= size
			// n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
			if n%2 == 0 && i == 0 {
				parity = 1 - j%2
				j /= 2
			}
		} // else j=0，在 n 很大的情况下，只能按照 1,2,3,... 的顺序填
		ans[i] = cand[parity][j]
		cand[parity] = slices.Delete(cand[parity], j, j+1)
		parity ^= 1 // 下一个数的奇偶性
	}
	return ans
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n^2)$。**注**：如果用有序集合或者树状数组维护剩余元素，可以做到 $\mathcal{O}(n\log n)$。考虑到本题 $n$ 很小，直接删除元素是最快的（常数小）。
- 空间复杂度：$\mathcal{O}(n)$。

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
