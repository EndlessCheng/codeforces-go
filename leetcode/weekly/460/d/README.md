枚举 $\textit{nums}$ 的子集作为 $B$，这样我们只需专注 $\text{XOR}$ 的问题。

设 $\textit{nums}$ 除去 $B$ 的剩余元素为 $S$，问题变成：

- 把 $S$ 划分成两个子集 $A$ 和 $C$，计算 $\text{XOR}(A) + \text{XOR}(C)$ 的最大值。

由于异或运算每个比特位互相独立，讨论其中一个比特位：

- 如果 $S$ 在这一位上一共有奇数个 $1$，由于奇数 = 奇数 + 偶数，无论怎么划分，一定是奇数个 $1$ 的异或和 + 偶数个 $1$ 的异或和 = $1$。即 $\text{XOR}(A) + \text{XOR}(C)$ 的这一位恒 $1$。
  - 由于奇数个 $1$ 的异或和是 $1$，偶数个 $1$ 的异或和是 $0$，所以 $S$ 中的包含奇数个 $1$ 的比特位所组成的二进制数，就是 $\text{XOR}(S)$。
- 如果 $S$ 在这一位上一共有偶数个 $1$，那么可以是：
  - 偶数个 $1$ 的异或和 + 偶数个 $1$ 的异或和 = $0$。
  - 奇数个 $1$ 的异或和 + 奇数个 $1$ 的异或和 = $2$。
  - 虽然有多种结果，但有一样东西是不变的：$\text{XOR}(A)$ 在这一位上的异或和 = $\text{XOR}(C)$ 在这一位上的异或和。
  - 只考虑有偶数个 $1$ 的比特位（称其为特殊比特位），把 $S$ 中每个数的非特殊比特位改成 $0$，得到一个新的集合 $S'$。然后把 $S'$ 划分为 $A'$ 和 $C'$，由于 $\text{XOR}(A')=\text{XOR}(C')$，得 $\text{XOR}(A') + \text{XOR}(C') = 2\cdot \text{XOR}(A')$。

所以 $\text{XOR}(A) + \text{XOR}(C)$ 的最大值就是

$$
\text{XOR}(S) + 2\cdot \text{XOR}(A')
$$

的最大值。

由于 $\text{XOR}(S)$ 是个定值，问题变成计算 $\text{XOR}(A')$ 的最大值，即

- 从 $S'$ 中选择一些数 $A'$，计算这些数的最大异或和。

这是 [线性基](https://oi-wiki.org/math/basis/) 的标准应用，具体请看 [视频讲解](https://www.bilibili.com/video/BV1pm8vzAEXx/)，欢迎点赞关注~

**注**：异或运算本质是 $w$ 维线性空间中的模 $2$ 加法，一个二进制数可以视作一个 $w$ 维的向量，本题 $w\le 30$。线性基（线性异或基）计算的是这个线性空间中的由 $S'$ 张成的一组基，$S'$ 中的每个二进制数（视作向量）都可以被这组基表出。

## 优化前

```py [sol-Python3]
# 线性基模板
class XorBasis:
    def __init__(self, n: int):
        self.b = [0] * n

    def insert(self, x: int) -> None:
        b = self.b
        while x:
            i = x.bit_length() - 1  # x 的最高位
            if b[i] == 0:  # x 和之前的基是线性无关的
                b[i] = x  # 新增一个基，最高位为 i
                return
            x ^= b[i]  # 保证参与 max_xor 的基的最高位是互不相同的，方便我们贪心
        # 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基

    def max_xor(self) -> int:
        b = self.b
        res = 0
        # 从高到低贪心：越高的位，越必须是 1
        # 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for i in range(len(b) - 1, -1, -1):
            if res ^ b[i] > res:
                res ^= b[i]
        return res


class Solution:
    def maximizeXorAndXor(self, nums: List[int]) -> int:
        n = len(nums)
        sz = max(nums).bit_length()

        # 预处理所有子集的 AND 和 XOR（刷表法）
        u = 1 << n
        sub_and = [0] * u
        sub_xor = [0] * u
        sub_and[0] = -1
        for i, x in enumerate(nums):
            high_bit = 1 << i
            for mask in range(high_bit):
                sub_and[high_bit | mask] = sub_and[mask] & x
                sub_xor[high_bit | mask] = sub_xor[mask] ^ x
        sub_and[0] = 0

        def max_xor2(sub: int) -> int:
            b = XorBasis(sz)
            xor = sub_xor[sub]
            for i, x in enumerate(nums):
                if sub >> i & 1:
                    # 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
                    b.insert(x & ~xor)
            return xor + b.max_xor() * 2

        return max(sub_and[i] + max_xor2((u - 1) ^ i) for i in range(u))
```

```java [sol-Java]
// 线性基模板
class XorBasis {
    private final int[] b;

    public XorBasis(int n) {
        b = new int[n];
    }

    public void insert(int x) {
        while (x > 0) {
            int i = 31 - Integer.numberOfLeadingZeros(x); // x 的最高位
            if (b[i] == 0) { // x 和之前的基是线性无关的
                b[i] = x; // 新增一个基，最高位为 i
                return;
            }
            x ^= b[i]; // 保证参与 maxXor 的基的最高位是互不相同的，方便我们贪心
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    public int maxXor() {
        int res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.length - 1; i >= 0; i--) {
            res = Math.max(res, res ^ b[i]);
        }
        return res;
    }
}

class Solution {
    public long maximizeXorAndXor(int[] nums) {
        int n = nums.length;
        int maxVal = Arrays.stream(nums).max().getAsInt();
        int sz = 32 - Integer.numberOfLeadingZeros(maxVal);

        // 预处理所有子集的 AND 和 XOR（刷表法）
        int u = 1 << n;
        int[] subAnd = new int[u];
        int[] subXor = new int[u];
        subAnd[0] = -1;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int highBit = 1 << i;
            for (int mask = 0; mask < highBit; mask++) {
                subAnd[highBit | mask] = subAnd[mask] & x;
                subXor[highBit | mask] = subXor[mask] ^ x;
            }
        }
        subAnd[0] = 0;

        long ans = 0;
        for (int i = 0; i < u; i++) {
            int j = (u - 1) ^ i;
            ans = Math.max(ans, subAnd[i] + maxXor2(j, subXor[j], nums, sz));
        }
        return ans;
    }

    private long maxXor2(int sub, int xor, int[] nums, int sz) {
        XorBasis b = new XorBasis(sz);
        for (int i = 0; i < nums.length; i++) {
            if ((sub >> i & 1) > 0) {
                // 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
                b.insert(nums[i] & ~xor);
            }
        }
        return xor + b.maxXor() * 2L;
    }
}
```

```cpp [sol-C++]
// 线性基模板
class XorBasis {
    vector<uint32_t> b;

public:
    XorBasis(int n) : b(n) {}

    void insert(uint32_t x) {
        while (x) {
            int i = bit_width(x) - 1; // x 的最高位
            if (b[i] == 0) { // x 和之前的基是线性无关的
                b[i] = x; // 新增一个基，最高位为 i
                return;
            }
            x ^= b[i]; // 保证参与 max_xor 的基的最高位是互不相同的，方便我们贪心
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    uint32_t max_xor() {
        uint32_t res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.size() - 1; i >= 0; i--) {
            res = max(res, res ^ b[i]);
        }
        return res;
    }
};

class Solution {
public:
    long long maximizeXorAndXor(vector<int>& nums) {
        int n = nums.size();
        int sz = bit_width((uint32_t) ranges::max(nums));

        // 预处理所有子集的 AND 和 XOR（刷表法）
        int u = 1 << n;
        vector<int> sub_and(u), sub_xor(u);
        sub_and[0] = -1;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int high_bit = 1 << i;
            for (int mask = 0; mask < high_bit; mask++) {
                sub_and[high_bit | mask] = sub_and[mask] & x;
                sub_xor[high_bit | mask] = sub_xor[mask] ^ x;
            }
        }
        sub_and[0] = 0;

        auto max_xor2 = [&](int sub) -> long long {
            XorBasis b(sz);
            int xor_ = sub_xor[sub];
            for (int i = 0; i < n; i++) {
                if (sub >> i & 1) {
                    // 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
                    b.insert(nums[i] & ~xor_);
                }
            }
            return xor_ + b.max_xor() * 2LL;
        };

        long long ans = 0;
        for (int i = 0; i < u; i++) {
            ans = max(ans, sub_and[i] + max_xor2((u - 1) ^ i));
        }
        return ans;
    }
};
```

```go [sol-Go]
// 线性基模板
type xorBasis []int

func (b xorBasis) insert(x int) {
	for x > 0 {
		i := bits.Len(uint(x)) - 1 // x 的最高位
		if b[i] == 0 { // x 和之前的基是线性无关的
			b[i] = x // 新增一个基，最高位为 i
			return
		}
		x ^= b[i] // 保证参与 maxXor 的基的最高位是互不相同的，方便我们贪心
	}
	// 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
}

func (b xorBasis) maxXor() (res int) {
	// 从高到低贪心：越高的位，越必须是 1
	// 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
	for i := len(b) - 1; i >= 0; i-- {
		res = max(res, res^b[i])
	}
	return
}

func maximizeXorAndXor(nums []int) int64 {
	n := len(nums)
	// 预处理所有子集的 AND 和 XOR（刷表法）
	type pair struct{ and, xor int }
	subSum := make([]pair, 1<<n)
	subSum[0].and = -1
	for i, x := range nums {
		highBit := 1 << i
		for mask, p := range subSum[:highBit] {
			subSum[highBit|mask] = pair{p.and & x, p.xor ^ x}
		}
	}
	subSum[0].and = 0

	sz := bits.Len(uint(slices.Max(nums)))
	b := make(xorBasis, sz)
	maxXor2 := func(sub uint) (res int) {
		clear(b)
		xor := subSum[sub].xor
		for ; sub > 0; sub &= sub - 1 {
			x := nums[bits.TrailingZeros(sub)]
			b.insert(x &^ xor) // 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
		}
		return xor + b.maxXor()*2
	}

	ans := 0
	u := 1<<n - 1
	for i, p := range subSum {
		ans = max(ans, p.and+maxXor2(uint(u^i)))
	}
	return int64(ans)
}
```

## 优化

最优性剪枝：$\text{XOR}(A)$ 和 $\text{XOR}(C)$ 的理论最大值是 $\text{OR}(\complement_UB)$，如果 

$$
\text{AND}(B) + 2\cdot \text{OR}(\complement_UB) \le \textit{ans}
$$

那么答案不可能变大，直接计算下一个子集。

进一步地，前文分析过，对于 $\text{XOR}(\complement_UB)$ 中的等于 $1$ 的比特位，$\text{XOR}(A) + \text{XOR}(C)$ 在这些比特位上恒为 $1$，不可能是 $2$。所以更精细的剪枝条件是

$$
\text{AND}(B) + 2\cdot \text{OR}(\complement_UB) - \text{XOR}(\complement_UB) \le \textit{ans}
$$

若不满足则跳过线性基的计算。

```py [sol-Python3]
# 线性基模板
class XorBasis:
    def __init__(self, n: int):
        self.b = [0] * n

    def insert(self, x: int) -> None:
        b = self.b
        while x:
            i = x.bit_length() - 1  # x 的最高位
            if b[i] == 0:  # x 和之前的基是线性无关的
                b[i] = x  # 新增一个基，最高位为 i
                return
            x ^= b[i]  # 保证参与 max_xor 的基的最高位是互不相同的，方便我们贪心
        # 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基

    def max_xor(self) -> int:
        b = self.b
        res = 0
        # 从高到低贪心：越高的位，越必须是 1
        # 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for i in range(len(b) - 1, -1, -1):
            if res ^ b[i] > res:
                res ^= b[i]
        return res


class Solution:
    def maximizeXorAndXor(self, nums: List[int]) -> int:
        n = len(nums)
        sz = max(nums).bit_length()

        # 多算一个子集 OR，用于剪枝
        u = 1 << n
        sub_and = [0] * u
        sub_xor = [0] * u
        sub_or = [0] * u
        sub_and[0] = -1
        for i, x in enumerate(nums):
            high_bit = 1 << i
            for mask in range(high_bit):
                sub_and[high_bit | mask] = sub_and[mask] & x
                sub_xor[high_bit | mask] = sub_xor[mask] ^ x
                sub_or[high_bit | mask] = sub_or[mask] | x
        sub_and[0] = 0

        def max_xor2(sub: int) -> int:
            b = XorBasis(sz)
            xor = sub_xor[sub]
            for i, x in enumerate(nums):
                if sub >> i & 1:
                    b.insert(x & ~xor)
            return xor + b.max_xor() * 2

        ans = 0
        for i in range(u):
            j = (u - 1) ^ i
            if sub_and[i] + sub_or[j] * 2 - sub_xor[j] > ans:  # 有机会让 ans 变得更大
                ans = max(ans, sub_and[i] + max_xor2(j))
        return ans
```

```java [sol-Java]
// 线性基模板
class XorBasis {
    private final int[] b;

    public XorBasis(int n) {
        b = new int[n];
    }

    public void insert(int x) {
        while (x > 0) {
            int i = 31 - Integer.numberOfLeadingZeros(x); // x 的最高位
            if (b[i] == 0) { // x 和之前的基是线性无关的
                b[i] = x; // 新增一个基，最高位为 i
                return;
            }
            x ^= b[i]; // 保证参与 maxXor 的基的最高位是互不相同的，方便我们贪心
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    public int maxXor() {
        int res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.length - 1; i >= 0; i--) {
            res = Math.max(res, res ^ b[i]);
        }
        return res;
    }
}

class Solution {
    public long maximizeXorAndXor(int[] nums) {
        int n = nums.length;
        int maxVal = Arrays.stream(nums).max().getAsInt();
        int sz = 32 - Integer.numberOfLeadingZeros(maxVal);

        // 多算一个子集 OR，用于剪枝
        int u = 1 << n;
        int[] subAnd = new int[u];
        int[] subXor = new int[u];
        int[] subOr = new int[u];
        subAnd[0] = -1;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int highBit = 1 << i;
            for (int mask = 0; mask < highBit; mask++) {
                subAnd[highBit | mask] = subAnd[mask] & x;
                subXor[highBit | mask] = subXor[mask] ^ x;
                subOr[highBit | mask] = subOr[mask] | x;
            }
        }
        subAnd[0] = 0;

        long ans = 0;
        for (int i = 0; i < u; i++) {
            int j = (u - 1) ^ i;
            if (subAnd[i] + subOr[j] * 2L - subXor[j] > ans) { // 有机会让 ans 变得更大
                ans = Math.max(ans, subAnd[i] + maxXor2(j, subXor[j], nums, sz));
            }
        }
        return ans;
    }

    private long maxXor2(int sub, int xor, int[] nums, int sz) {
        XorBasis b = new XorBasis(sz);
        for (int i = 0; i < nums.length; i++) {
            if ((sub >> i & 1) > 0) {
                b.insert(nums[i] & ~xor);
            }
        }
        return xor + b.maxXor() * 2L;
    }
}
```

```cpp [sol-C++]
// 线性基模板
class XorBasis {
    vector<uint32_t> b;

public:
    XorBasis(int n) : b(n) {}

    void insert(uint32_t x) {
        while (x) {
            int i = bit_width(x) - 1; // x 的最高位
            if (b[i] == 0) { // x 和之前的基是线性无关的
                b[i] = x; // 新增一个基，最高位为 i
                return;
            }
            x ^= b[i]; // 保证参与 max_xor 的基的最高位是互不相同的，方便我们贪心
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    uint32_t max_xor() {
        uint32_t res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.size() - 1; i >= 0; i--) {
            res = max(res, res ^ b[i]);
        }
        return res;
    }
};

class Solution {
public:
    long long maximizeXorAndXor(vector<int>& nums) {
        int n = nums.size();
        int sz = bit_width((uint32_t) ranges::max(nums));

        // 多算一个子集 OR，用于剪枝
        int u = 1 << n;
        vector<int> sub_and(u), sub_xor(u), sub_or(u);
        sub_and[0] = -1;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int high_bit = 1 << i;
            for (int mask = 0; mask < high_bit; mask++) {
                sub_and[high_bit | mask] = sub_and[mask] & x;
                sub_xor[high_bit | mask] = sub_xor[mask] ^ x;
                sub_or[high_bit | mask] = sub_or[mask] | x;
            }
        }
        sub_and[0] = 0;

        auto max_xor2 = [&](int sub) -> long long {
            XorBasis b(sz);
            int xor_ = sub_xor[sub];
            for (int i = 0; i < n; i++) {
                if (sub >> i & 1) {
                    b.insert(nums[i] & ~xor_);
                }
            }
            return xor_ + b.max_xor() * 2LL;
        };

        long long ans = 0;
        for (int i = 0; i < u; i++) {
            int j = (u - 1) ^ i;
            if (sub_and[i] + sub_or[j] * 2LL - sub_xor[j] > ans) { // 有机会让 ans 变得更大
                ans = max(ans, sub_and[i] + max_xor2(j));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 线性基模板
type xorBasis []int

func (b xorBasis) insert(x int) {
	for x > 0 {
		i := bits.Len(uint(x)) - 1 // x 的最高位
		if b[i] == 0 { // x 和之前的基是线性无关的
			b[i] = x // 新增一个基，最高位为 i
			return
		}
		x ^= b[i] // 保证参与 maxXor 的基的最高位是互不相同的，方便我们贪心
	}
	// 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
}

func (b xorBasis) maxXor() (res int) {
	// 从高到低贪心：越高的位，越必须是 1
	// 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
	for i := len(b) - 1; i >= 0; i-- {
		res = max(res, res^b[i])
	}
	return
}

func maximizeXorAndXor(nums []int) int64 {
	n := len(nums)
	type pair struct{ and, xor, or int } // 多算一个子集 OR，用于剪枝
	subSum := make([]pair, 1<<n)
	subSum[0].and = -1
	for i, x := range nums {
		highBit := 1 << i
		for mask, p := range subSum[:highBit] {
			subSum[highBit|mask] = pair{p.and & x, p.xor ^ x, p.or | x}
		}
	}
	subSum[0].and = 0

	sz := bits.Len(uint(slices.Max(nums)))
	b := make(xorBasis, sz)
	maxXor2 := func(sub uint) (res int) {
		clear(b)
		xor := subSum[sub].xor
		for ; sub > 0; sub &= sub - 1 {
			x := nums[bits.TrailingZeros(sub)]
			b.insert(x &^ xor)
		}
		return xor + b.maxXor()*2
	}

	ans := 0
	u := 1<<n - 1
	for i, p := range subSum {
		j := u ^ i
		if p.and+subSum[j].or*2-subSum[j].xor > ans { // 有机会让 ans 变得更大
			ans = max(ans, p.and+maxXor2(uint(j)))
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(2^n + \log U)$。

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
