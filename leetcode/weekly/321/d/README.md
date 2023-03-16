### 提示 1

先来计算子数组长为奇数的情况。

由于题目保证 $\textit{nums}$ 中的整数互不相同，「$k$ 是长为奇数的子数组的中位数」等价于「子数组中小于 $k$ 的数的个数 $=$ 大于 $k$ 的数的个数」。

这相当于「左侧小于 $+$ 右侧小于 $=$ 左侧大于 $+$ 右侧大于」。

变形得到「左侧小于 $-$ 左侧大于 $=$ 右侧大于 $-$ 右侧小于」。

为了方便计算，把这四类数字等价转换：

- 左侧小于：在 $k$ 左侧且比 $k$ 小的视作 $1$；
- 左侧大于：在 $k$ 左侧且比 $k$ 大的视作 $-1$；
- 右侧大于：在 $k$ 右侧且比 $k$ 大的视作 $1$；
- 右侧小于：在 $k$ 右侧且比 $k$ 小的视作 $-1$。

此外，把 $k$ 视作 $0$。

### 提示 2

用示例 1 来说明。$[3,2,1,4,5]$ 可以视作 $[1,1,1,0,1]$。由于子数组一定要包含中位数，故从下标 $3$ 开始倒序枚举子数组左端点，计算「左侧小于 - 左侧大于」的值 $x$，依次为 $0,1,2,3$，记到一个哈希表 $\textit{cnt}$ 中。然后从下标 $3$ 开始正序枚举子数组右端点，计算「右侧大于 $-$ 右侧小于」的值 $x$，依次为 $0,1$。**对每个右端点的 $x$，去看有多少个一样的左端点的 $x$**。所以中位数为 $4$ 且长为奇数的子数组个数为 $\textit{cnt}[0]+\textit{cnt}[1] = 1+1=2$，对应的子数组为 $[4]$ 和 $[1,4,5]$。

对于子数组长为偶数的情况，「$k$ 是长为偶数的子数组的中位数」等价于「左侧小于 $+$ 右侧小于 $=$ 左侧大于 $+$ 右侧大于 $-1$」，即「左侧小于 $-$ 左侧大于 $=$ 右侧大于 $-$ 右侧小于 $-1$」。相比奇数的情况，等号右侧多了个 $-1$，那么接着上面的「右侧大于 $-$ 右侧小于」的值 $x$ 来说，$\textit{cnt}[x-1]$ 就是该右端点对应的符合题目要求的长为偶数的子数组个数。累加这些 $\textit{cnt}[x-1]$，就是子数组长为偶数时的答案。

例如 $[3,2,1,4,5]$ 中就有 $cnt[1-1]=1$ 个，对应的子数组为 $[4,5]$。

把奇数的和偶数的个数相加，即为答案。

**附**：更多的例子和讲解，可以看我的 [视频](https://www.bilibili.com/video/BV1sD4y1e7pr/?t=12m40s)（视频中先枚举的右侧），从 12:40 开始。

### 答疑

**问**：答案的大小看上去是 $O(n^2)$，这会不会爆 `int`（超过 $2^{31}-1$）？

**答**：这是个有趣的问题，怎么构造能让答案尽量大呢？

- 既然 $k$ 是中位数，不妨取 $k=\left[\dfrac{n}{2}\right]$ 且位于 $\textit{nums}$ 中间，从而让尽量多的子数组包含 $k$；
- 小于 $k$ 和大于 $k$ 的数**交替排布**，从而产生大量重复的 $x$。 

例如 $n=10$ 的时候，取 $k=5$，并构造如下 $\textit{nums}$：

$$
[1,6,2,7,5,8,3,9,4,10]
$$

转换成 $1$，$-1$ 和 $0$：

$$
[1,-1,1,-1,0,1,-1,1,-1,1]
$$

按照上面的算法，从中间开始向左累加，可以得到大约 $\dfrac{n}{4}$ 个 $0$ 和 $\dfrac{n}{4}$ 个 $-1$；从中间开始向右累加，可以得到大约 $\dfrac{n}{4}$ 个 $0$ 和 $\dfrac{n}{4}$ 个 $1$。所以中位数为 $k$ 的奇数长度子数组约有 $\dfrac{n}{4}\times \dfrac{n}{4}$ 个，中位数为 $k$ 的偶数长度子数组约有 $2\times\dfrac{n}{4}\times \dfrac{n}{4}$ 个。

答案约为 

$$
\dfrac{3n^2}{16}
$$

代入 $n=10^5$ 得 $1.875\times 10^9 < 2^{31}  \approx  2.1\times 10^9$，所以不会爆 `int`。

```py [sol1-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        pos = nums.index(k)
        # i=pos 的时候 x 是 0，直接记到 cnt 中，这样下面不是大于 k 就是小于 k
        cnt, x = Counter({0: 1}), 0
        for i in range(pos - 1, -1, -1):  # 从 pos-1 开始累加 x
            x += 1 if nums[i] < k else -1
            cnt[x] += 1

        # i=pos 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
        ans, x = cnt[0] + cnt[-1], 0
        for i in range(pos + 1, len(nums)):  # 从 pos+1 开始累加 x
            x += 1 if nums[i] > k else -1
            ans += cnt[x] + cnt[x - 1]
        return ans
```

```java [sol1-Java]
class Solution {
    public int countSubarrays(int[] nums, int k) {
        int pos = 0, n = nums.length;
        while (nums[pos] != k) ++pos;

        var cnt = new HashMap<Integer, Integer>();
        // i=pos 的时候 x 是 0，直接记到 cnt 中，这样下面不是大于 k 就是小于 k
        cnt.put(0, 1);
        for (int i = pos - 1, x = 0; i >= 0; --i) { // 从 pos-1 开始累加 x
            x += nums[i] < k ? 1 : -1;
            cnt.merge(x, 1, Integer::sum);
        }

        // i=pos 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
        int ans = cnt.get(0) + cnt.getOrDefault(-1, 0);
        for (int i = pos + 1, x = 0; i < n; ++i) { // 从 pos+1 开始累加 x
            x += nums[i] > k ? 1 : -1;
            ans += cnt.getOrDefault(x, 0) + cnt.getOrDefault(x - 1, 0);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int countSubarrays(vector<int> &nums, int k) {
        int pos = find(nums.begin(), nums.end(), k) - nums.begin();
        // i=pos 的时候 x 是 0，直接记到 cnt 中，这样下面不是大于 k 就是小于 k
        unordered_map<int, int> cnt{{0, 1}};
        for (int i = pos - 1, x = 0; i >= 0; --i) { // 从 pos-1 开始累加 x
            x += nums[i] < k ? 1 : -1;
            ++cnt[x];
        }

        // i=pos 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
        int ans = cnt[0] + cnt[-1];
        for (int i = pos + 1, x = 0; i < nums.size(); ++i) { // 从 pos+1 开始累加 x
            x += nums[i] > k ? 1 : -1;
            ans += cnt[x] + cnt[x - 1];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	// i=pos 的时候 x 是 0，直接记到 cnt 中，这样下面不是大于 k 就是小于 k
	cnt, x := map[int]int{0: 1}, 0
	for i := pos - 1; i >= 0; i-- { // 从 pos-1 开始累加 x
		if nums[i] < k {
			x++
		} else {
			x--
		}
		cnt[x]++
	}

	// i=pos 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
	ans, x := cnt[0]+cnt[-1], 0
	for _, v := range nums[pos+1:] { // 从 pos+1 开始累加 x
		if v > k {
			x++
		} else {
			x--
		}
		ans += cnt[x] + cnt[x-1]
	}
	return ans
}
```

### 优化

$x$ 的范围在 $[-(n-1),n-1]$ 之间，所以可以用数组代替哈希表，加快运行速度。

考虑到还需要访问 $\textit{cnt}[x-1]$，所以实际的范围为 $[-n,n-1]$，因此需要一个长为 $2n$ 的数组，此时 $x$ 应当初始化为 $n$。

```py [sol2-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        pos, n = nums.index(k), len(nums)
        cnt, x = [0] * (n * 2), n
        cnt[x] = 1
        for i in range(pos - 1, -1, -1):  # 从 pos-1 开始累加 x
            x += 1 if nums[i] < k else -1
            cnt[x] += 1

        ans, x = cnt[n] + cnt[n - 1], n
        for i in range(pos + 1, len(nums)):  # 从 pos+1 开始累加 x
            x += 1 if nums[i] > k else -1
            ans += cnt[x] + cnt[x - 1]
        return ans
```

```java [sol2-Java]
class Solution {
    public int countSubarrays(int[] nums, int k) {
        int pos = 0, n = nums.length;
        while (nums[pos] != k) ++pos;

        var cnt = new int[n * 2];
        cnt[n] = 1;
        for (int i = pos - 1, x = n; i >= 0; --i) { // 从 pos-1 开始累加 x
            x += nums[i] < k ? 1 : -1;
            ++cnt[x];
        }

        int ans = cnt[n] + cnt[n - 1];
        for (int i = pos + 1, x = n; i < n; ++i) { // 从 pos+1 开始累加 x
            x += nums[i] > k ? 1 : -1;
            ans += cnt[x] + cnt[x - 1];
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int countSubarrays(vector<int> &nums, int k) {
        int pos = find(nums.begin(), nums.end(), k) - nums.begin();
        int n = nums.size(), cnt[n * 2];
        memset(cnt, 0, sizeof(cnt));
        cnt[n] = 1;
        for (int i = pos - 1, x = n; i >= 0; --i) { // 从 pos-1 开始累加 x
            x += nums[i] < k ? 1 : -1;
            ++cnt[x];
        }

        int ans = cnt[n] + cnt[n - 1];
        for (int i = pos + 1, x = n; i < n; ++i) { // 从 pos+1 开始累加 x
            x += nums[i] > k ? 1 : -1;
            ans += cnt[x] + cnt[x - 1];
        }
        return ans;
    }
};
```

```go [sol2-Go]
func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	n := len(nums)
	cnt, x := make([]int, n*2), n
	cnt[x] = 1
	for i := pos - 1; i >= 0; i-- { // 从 pos-1 开始累加 x
		if nums[i] < k {
			x++
		} else {
			x--
		}
		cnt[x]++
	}

	ans, x := cnt[n]+cnt[n-1], n
	for _, v := range nums[pos+1:] { // 从 pos+1 开始累加 x
		if v > k {
			x++
		} else {
			x--
		}
		ans += cnt[x] + cnt[x-1]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

### 相似题目（等价转换）

- [面试题 17.05. 字母与数字](https://leetcode.cn/problems/find-longest-subarray-lcci/)
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
- [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/)

---

附：我的 [每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)，已分类整理好。

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~
