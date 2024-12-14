## 写法一

**核心观察**：对于两个数 $x$ 和 $y$，如果 $x$ 在 $y$ 左边，且 $x\le y$ 以及 $x\cdot \textit{multiplier} > y$，那么我们会先操作 $x$，然后操作 $y$。由于 $x\cdot \textit{multiplier} \le y\cdot \textit{multiplier}$，这意味着下一次一定会操作 $x$。继续推导下去，后面的操作顺序一定是 $y,x,y,x,\cdots$

这意味着当两个数接近时，我们会**交替操作**这两个数，而**不会连续操作同一个数**。

对于更多的数的情况也同理，当这些数接近时，我们会按照从小到大的顺序依次操作这些数。

那么，首先用最小堆手动模拟操作，直到原数组的最大值 $\textit{mx}$ 成为这 $n$ 个数的最小值。根据上面的结论，后面的操作就不需要手动模拟了。

设此时还剩下 $k$ 次操作，那么：

- 对于前 $k\bmod n$ 小的数，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor+1$ 次。
- 其余元素，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor$ 次。

用**快速幂**计算操作这么多次后的结果，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cMW6ePEwC/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mx = max(nums)
        h = [(x, i) for i, x in enumerate(nums)]
        heapify(h)

        # 模拟，直到堆顶是 mx
        while k and h[0][0] < mx:
            x, i = h[0]
            heapreplace(h, (x * multiplier, i))
            k -= 1

        # 剩余的操作可以直接用公式计算
        h.sort()
        for i, (x, j) in enumerate(h):
            nums[j] = x * pow(multiplier, k // n + (i < k % n), MOD) % MOD
        return nums
```

```py [sol-Python3 写法二]
# 也可以模拟到 k 刚好是 n 的倍数时才停止，这样最后无需排序
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mx = max(nums)
        h = [(x, i) for i, x in enumerate(nums)]
        heapify(h)

        # 模拟，直到堆顶是 mx
        while k and (h[0][0] < mx or k % n):
            x, i = h[0]
            heapreplace(h, (x * multiplier, i))
            k -= 1

        # 剩余的操作可以直接用公式计算
        for x, j in h:
            nums[j] = x * pow(multiplier, k // n, MOD) % MOD
        return nums
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int[] getFinalState(int[] nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return nums;
        }

        int n = nums.length;
        int mx = 0;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> a[0] != b[0] ? Long.compare(a[0], b[0]) : Long.compare(a[1], b[1]));
        for (int i = 0; i < n; i++) {
            mx = Math.max(mx, nums[i]);
            pq.offer(new long[]{nums[i], i});
        }

        // 模拟，直到堆顶是 mx
        for (; k > 0 && pq.peek()[0] < mx; k--) {
            long[] p = pq.poll();
            p[0] *= multiplier;
            pq.offer(p);
        }

        // 剩余的操作可以直接用公式计算
        for (int i = 0; i < n; i++) {
            long[] p = pq.poll();
            nums[(int) p[1]] = (int) (p[0] % MOD * pow(multiplier, k / n + (i < k % n ? 1 : 0)) % MOD);
        }
        return nums;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    vector<int> getFinalState(vector<int>& nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return move(nums);
        }

        int n = nums.size();
        int mx = ranges::max(nums);
        vector<pair<long long, int>> h(n);
        for (int i = 0; i < n; i++) {
            h[i] = {nums[i], i};
        }
        ranges::make_heap(h, greater<>()); // 最小堆，O(n) 堆化

        // 模拟，直到堆顶是 mx
        for (; k && h[0].first < mx; k--) {
            ranges::pop_heap(h, greater<>());
            h.back().first *= multiplier;
            ranges::push_heap(h, greater<>());
        }

        // 剩余的操作可以直接用公式计算
        ranges::sort(h);
        for (int i = 0; i < n; i++) {
            auto& [x, j] = h[i];
            nums[j] = x % MOD * pow(multiplier, k / n + (i < k % n)) % MOD;
        }
        return move(nums);
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n\log_{m} U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$m=\textit{multiplier}$。瓶颈在模拟那，每个数至多操作 $\mathcal{O}(\log_{m} U)$ 次。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：优化

设把每个 $\textit{nums}[i]$ 都操作到至少为 $\textit{nums}$ 的最大值，总共需要操作 $\textit{total}$ 次。

分类讨论：

- 如果 $k < \textit{total}$，那么直接用最小堆暴力模拟 $k$ 次。
- 如果 $k\ge \textit{total}$，我们可以先把每个 $\textit{nums}[i]$ 暴力操作到 $\textit{nums}$ 的最大值，然后剩余的操作直接用公式计算。

此外，可以先把 `pow` 算出来，而不是在循环中计算。

```py [sol-Python3]
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mn = min(nums)
        mx = max(nums)
        t = 0
        while mn < mx:
            mn *= multiplier
            t += 1

        if k < t * n:
            # 暴力模拟
            h = [(x, i) for i, x in enumerate(nums)]
            heapify(h)
            for _ in range(k):
                x, i = h[0]
                heapreplace(h, (x * multiplier, i))
            for x, j in h:
                nums[j] = x % MOD
            return nums

        # 每个数直接暴力操作到 >= mx
        for i, x in enumerate(nums):
            while x < mx:
                x *= multiplier
                k -= 1
            nums[i] = x

        # 剩余的操作可以直接用公式计算
        pow1 = pow(multiplier, k // n, MOD)
        pow2 = pow1 * multiplier % MOD
        for i, (x, j) in enumerate(sorted((x, i) for i, x in enumerate(nums))):
            nums[j] = x * (pow2 if i < k % n else pow1) % MOD
        return nums
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int[] getFinalState(int[] nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return nums;
        }

        int n = nums.length;
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        // 每个数直接暴力操作到 >= mx
        long[] a = new long[n];
        int left = k;
        outer:
        for (int i = 0; i < n; i++) {
            long x = nums[i];
            while (x < mx) {
                x *= multiplier;
                if (--left < 0) {
                    break outer;
                }
            }
            a[i] = x;
        }

        if (left < 0) {
            // 暴力模拟
            PriorityQueue<long[]> pq = new PriorityQueue<>((p, q) -> p[0] != q[0] ? Long.compare(p[0], q[0]) : Long.compare(p[1], q[1]));
            for (int i = 0; i < n; i++) {
                pq.offer(new long[]{nums[i], i});
            }
            while (k-- > 0) {
                long[] p = pq.poll();
                p[0] *= multiplier;
                pq.offer(p);
            }
            while (!pq.isEmpty()) {
                long[] p = pq.poll();
                nums[(int) p[1]] = (int) (p[0] % MOD);
            }
            return nums;
        }

        Integer[] ids = new Integer[n];
        Arrays.setAll(ids, i -> i);
        Arrays.sort(ids, (i, j) -> Long.compare(a[i], a[j]));

        // 剩余的操作可以直接用公式计算
        k = left;
        long pow1 = pow(multiplier, k / n);
        long pow2 = pow1 * multiplier % MOD;
        for (int i = 0; i < n; i++) {
            int j = ids[i];
            nums[j] = (int) (a[j] % MOD * (i < k % n ? pow2 : pow1) % MOD);
        }
        return nums;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    vector<int> getFinalState(vector<int>& nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return move(nums);
        }

        int n = nums.size();
        long long mx = ranges::max(nums);
        vector<pair<long long, int>> h(n);
        for (int i = 0; i < n; i++) {
            h[i] = {nums[i], i};
        }
        auto clone = h;

        // 每个数直接暴力操作到 >= mx
        int left = k;
        for (auto& [x, _] : h) {
            while (x < mx) {
                x *= multiplier;
                if (--left < 0) {
                    goto outer;
                }
            }
        }
        outer:;

        if (left < 0) {
            // 暴力模拟
            h = move(clone);
            ranges::make_heap(h, greater<>()); // 最小堆，O(n) 堆化
            while (k--) {
                ranges::pop_heap(h, greater<>());
                h.back().first *= multiplier;
                ranges::push_heap(h, greater<>());
            }
            for (auto& [x, j] : h) {
                nums[j] = x % MOD;
            }
            return move(nums);
        }

        // 剩余的操作可以直接用公式计算
        k = left;
        long long pow1 = pow(multiplier, k / n);
        long long pow2 = pow1 * multiplier % MOD;
        // ranges::sort(h) 换成快速选择
        ranges::nth_element(h, h.begin() + k % n);
        for (int i = 0; i < n; i++) {
            auto& [x, j] = h[i];
            nums[j] = x % MOD * (i < k % n ? pow2 : pow1) % MOD;
        }
        return move(nums);
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	clone := slices.Clone(h)

	// 每个数直接暴力操作到 >= mx
	left := k
outer:
	for i := range h {
		for h[i].x < mx {
			h[i].x *= multiplier
			left--
			if left < 0 {
				break outer
			}
		}
	}

	if left < 0 {
		// 暴力模拟
		h = clone
		heap.Init(&h)
		for ; k > 0; k-- {
			h[0].x *= multiplier
			heap.Fix(&h, 0)
		}
		for _, p := range h {
			nums[p.i] = p.x % mod
		}
		return nums
	}

	// 剩余的操作可以直接用公式计算
	k = left
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

设 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$m=\textit{multiplier}$。

- 时间复杂度：
    - $k$ 较小时为 $\mathcal{O}(n+k\log n)$。Java 是 $\mathcal{O}(n\log n+k\log n)$。
    - $k$ 较大时为 $\mathcal{O}(n\log_{m} U+ n\log n)$ 或 $\mathcal{O}(n\log_{m} U)$。如果像 C++ 那样使用**快速选择**，时间复杂度为 $\mathcal{O}(n\log_{m} U)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法三：进一步优化

方法二瓶颈在暴力把每个 $\textit{nums}[i]$ 乘到至少为 $\textit{mx}$ 上。

通过预处理，把 $\textit{nums}[i]$ 乘到至少为 $\textit{mx}$，这一步可以做到 $\mathcal{O}(1)$。

根据 [@hqztrue](/u/hqztrue) 的 [题解](https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-ii/solutions/2892649/xian-xing-suan-fa-by-hqztrue-30fk/)，代码实现如下。

> 注：由于常数比较大，该优化并不明显。仅用来说明当 $k$ 比较大时，存在 $\mathcal{O}(n)$ 的做法。

```py [sol-Python3]
class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        if multiplier == 1:  # 数组不变
            return nums

        MOD = 1_000_000_007
        n = len(nums)
        mx = max(nums)

        # 打表，计算出最小的 e 满足 multiplier^e >= 2^i
        e_pow_m = []
        pow_m = pow2 = 1
        e = 0
        while pow2 <= mx:
            if pow_m < pow2:  # 由于 multiplier >= 2，这里只需写 if 而不是 while
                pow_m *= multiplier
                e += 1
            e_pow_m.append((e, pow_m))
            pow2 <<= 1

        # 把每个数都操作到 >= mx
        left = k
        clone = nums.copy()
        mx_len = mx.bit_length()
        for i, x in enumerate(nums):
            e, pow_m = e_pow_m[mx_len - x.bit_length()]
            if pow_m // multiplier * x >= mx:  # 多操作了一次
                pow_m //= multiplier
                e -= 1
            elif x * pow_m < mx:  # 少操作了一次
                pow_m *= multiplier
                e += 1
            left -= e
            if left < 0:
                break
            nums[i] *= pow_m

        if left < 0:
            # 暴力模拟
            h = [(x, i) for i, x in enumerate(clone)]
            heapify(h)
            for _ in range(k):
                x, i = h[0]
                heapreplace(h, (x * multiplier, i))
            for x, j in h:
                clone[j] = x % MOD
            return clone

        # 剩余的操作可以直接用公式计算
        k = left
        pow1 = pow(multiplier, k // n, MOD)
        pow2 = pow1 * multiplier % MOD
        for i, (x, j) in enumerate(sorted((x, i) for i, x in enumerate(nums))):
            nums[j] = x * (pow2 if i < k % n else pow1) % MOD
        return nums
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int[] getFinalState(int[] nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return nums;
        }

        int n = nums.length;
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        // 打表，计算出最小的 exp 满足 multiplier^exp >= 2^i
        int mxLZ = Integer.numberOfLeadingZeros(mx);
        int mxLen = Integer.SIZE - mxLZ;
        long[] powMs = new long[mxLen];
        int[] exps = new int[mxLen];
        int exp = 0, m = 0;
        for (long pow2 = 1, powM = 1; pow2 <= mx; pow2 <<= 1) {
            if (powM < pow2) { // 由于 multiplier >= 2，这里只需写 if 而不是 while
                powM *= multiplier;
                exp++;
            }
            powMs[m] = powM;
            exps[m++] = exp;
        }

        // 把每个数都操作到 >= mx
        long[] a = new long[n];
        int left = k;
        for (int i = 0; i < n && left >= 0; i++) {
            int j = Integer.numberOfLeadingZeros(nums[i]) - mxLZ;
            long x = nums[i];
            long powM = powMs[j];
            int e = exps[j];
            if (powM / multiplier * x >= mx) { // 多操作了一次
                powM /= multiplier;
                e--;
            } else if (x * powM < mx) { // 少操作了一次
                powM *= multiplier;
                e++;
            }
            left -= e;
            a[i] = x * powM;
        }

        if (left < 0) {
            // 暴力模拟
            PriorityQueue<long[]> pq = new PriorityQueue<>((p, q) -> p[0] != q[0] ? Long.compare(p[0], q[0]) : Long.compare(p[1], q[1]));
            for (int i = 0; i < n; i++) {
                pq.offer(new long[]{nums[i], i});
            }
            while (k-- > 0) {
                long[] p = pq.poll();
                p[0] *= multiplier;
                pq.offer(p);
            }
            while (!pq.isEmpty()) {
                long[] p = pq.poll();
                nums[(int) p[1]] = (int) (p[0] % MOD);
            }
            return nums;
        }

        Integer[] ids = new Integer[n];
        Arrays.setAll(ids, i -> i);
        Arrays.sort(ids, (i, j) -> Long.compare(a[i], a[j]));

        // 剩余的操作可以直接用公式计算
        k = left;
        long pow1 = pow(multiplier, k / n);
        long pow2 = pow1 * multiplier % MOD;
        for (int i = 0; i < n; i++) {
            int j = ids[i];
            nums[j] = (int) (a[j] % MOD * (i < k % n ? pow2 : pow1) % MOD);
        }
        return nums;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    vector<int> getFinalState(vector<int>& nums, int k, int multiplier) {
        if (multiplier == 1) { // 数组不变
            return move(nums);
        }

        int n = nums.size();
        int mx = ranges::max(nums);
        // 打表，计算出最小的 e 满足 multiplier^e >= 2^i
        vector<pair<int, long long>> e_pow_m;
        long long pow_m = 1;
        int e = 0;
        for (int pow2 = 1; pow2 <= mx; pow2 <<= 1) {
            if (pow_m < pow2) { // 由于 multiplier >= 2，这里只需写 if 而不是 while
                pow_m *= multiplier;
                e++;
            }
            e_pow_m.emplace_back(e, pow_m);
        }

        vector<pair<long long, int>> h(n);
        for (int i = 0; i < n; i++) {
            h[i] = {nums[i], i};
        }
        auto clone = h;

        // 把每个数都操作到 >= mx
        int left = k;
        int mx_clz = __builtin_clz(mx);
        for (int i = 0; i < n && left >= 0; i++) {
            auto [e, pow_m] = e_pow_m[__builtin_clz(nums[i]) - mx_clz];
            long long x = nums[i];
            if (pow_m / multiplier * x >= mx) { // 多操作了一次
                pow_m /= multiplier;
                e--;
            } else if (x * pow_m < mx) { // 少操作了一次
                pow_m *= multiplier;
                e++;
            }
            left -= e;
            h[i].first *= pow_m;
        }

        if (left < 0) {
            // 暴力模拟
            h = move(clone);
            ranges::make_heap(h, greater<>()); // 最小堆，O(n) 堆化
            while (k--) {
                ranges::pop_heap(h, greater<>());
                h.back().first *= multiplier;
                ranges::push_heap(h, greater<>());
            }
            for (auto& [x, j] : h) {
                nums[j] = x % MOD;
            }
            return move(nums);
        }

        // 剩余的操作可以直接用公式计算
        k = left;
        long long pow1 = pow(multiplier, k / n);
        long long pow2 = pow1 * multiplier % MOD;
        // ranges::sort(h) 换成快速选择
        ranges::nth_element(h, h.begin() + k % n);
        for (int i = 0; i < n; i++) {
            auto& [x, j] = h[i];
            nums[j] = x % MOD * (i < k % n ? pow2 : pow1) % MOD;
        }
        return move(nums);
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	clone := slices.Clone(h)

	// 打表，计算出最小的 e 满足 multiplier^e >= 2^i
	mxLen := bits.Len(uint(mx))
	type ep struct{ e, powM int }
	ePowM := make([]ep, 0, mxLen)
	for pow2, powM, e := 1, 1, 0; pow2 <= mx; pow2 <<= 1 {
		if powM < pow2 { // 由于 multiplier >= 2，这里只需写 if 而不是 for
			powM *= multiplier
			e++
		}
		ePowM = append(ePowM, ep{e, powM})
	}

	// 把每个数都操作到 >= mx
	left := k
	for i := range h {
		x := h[i].x
		p := ePowM[mxLen-bits.Len(uint(x))]
		e, powM := p.e, p.powM
		if powM/multiplier*x >= mx { // 多操作了一次
			powM /= multiplier
			e--
		} else if x*powM < mx { // 少操作了一次
			powM *= multiplier
			e++
		}
		left -= e
		if left < 0 {
			break
		}
		h[i].x *= powM
	}

	if left < 0 {
		// 暴力模拟
		h = clone
		heap.Init(&h)
		for ; k > 0; k-- {
			h[0].x *= multiplier
			heap.Fix(&h, 0)
		}
		sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
		for _, p := range h {
			nums[p.i] = p.x % mod
		}
		return nums
	}

	// 剩余的操作可以直接用公式计算
	k = left
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

设 $n$ 是 $\textit{nums}$ 的长度。

打表和计算快速幂的时间由于只有 $\mathcal{O}(\log)$，忽略不计。

- 时间复杂度：
  - $k$ 较小时为 $\mathcal{O}(n+k\log n)$。Java 是 $\mathcal{O}(n\log n+k\log n)$。
  - $k$ 较大时为 $\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$。如果像 C++ 那样使用**快速选择**，时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**堆**」。

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
