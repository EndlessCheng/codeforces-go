解决本题您需要掌握如下知识：

1. 前缀和，[原理讲解](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)
2. 哈希表
3. 模运算

请看 [视频讲解](https://www.bilibili.com/video/BV1Nj411178Z/) 第三题。

对于本题，由于需要统计 $\textit{cnt}$，我们可以把满足 $\textit{nums}[i]\bmod \textit{modulo} = k$ 的 $\textit{nums}[i]$ 视作 $1$，不满足则视作 $0$。

如此转换后，算出 $\textit{nums}$ 的前缀和数组 $s$，那么题目中的 $\textit{cnt}\bmod \textit{modulo} = k$ 等价于

$$
(s[r+1]-s[l])\bmod \textit{modulo} = k
$$

上式等价于（推导过程请看视频）

$$
s[l]\bmod \textit{modulo} = (s[r+1]-k)\bmod \textit{modulo}
$$

根据上式，我们可以一边枚举 $r$，一边用一个哈希表统计有多少个 $s[r+1]\bmod \textit{modulo}$。这样可以快速知道有多少个 $(s[r+1]-k)\bmod \textit{modulo}$，也就是 $s[l]\bmod \textit{modulo}$ 的个数，把个数加到答案中。

代码实现时，前缀和数组可以优化成一个变量 $s$。

```py [sol-Python3]
class Solution:
    def countInterestingSubarrays(self, nums: List[int], mod: int, k: int) -> int:
        cnt = Counter([0])  # 把 s[0]=0 算进去
        ans = s = 0
        for x in nums:
            s += x % mod == k
            ans += cnt[(s - k) % mod]  # Python 取模可以把负数自动转成非负数
            cnt[s % mod] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countInterestingSubarrays(List<Integer> nums, int mod, int k) {
        var cnt = new HashMap<Integer, Integer>();
        cnt.put(0, 1);  // 把 s[0]=0 算进去
        long ans = 0;
        int s = 0;
        for (int x : nums) {
            if (x % mod == k)
                s = (s + 1) % mod; // 这里取模，下面 cnt[s]++ 就不需要取模了
            ans += cnt.getOrDefault((s - k + mod) % mod, 0); // +mod 避免减法出现负数
            cnt.merge(s, 1, Integer::sum); // cnt[s]++
        }
        return ans;
    }

    // 数组版本，效率更高！
    // 因为 s 至多为 n
    public long countInterestingSubarrays(List<Integer> nums, int mod, int k) {
        int n = nums.size();
        var cnt = new int[n + 1];
        cnt[0] = 1;
        long ans = 0;
        int s = 0;
        for (int x : nums) {
            if (x % mod == k)
                s = (s + 1) % mod;
            int s2 = (s - k + mod) % mod;
            if (s2 <= n)
                ans += cnt[s2];
            cnt[s]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countInterestingSubarrays(vector<int> &nums, int mod, int k) {
        unordered_map<int, int> cnt;
        cnt[0] = 1; // 把 s[0]=0 算进去
        long long ans = 0;
        int s = 0;
        for (int x: nums) {
            s += x % mod == k;
            ans += cnt[(s - k + mod) % mod]; // +mod 避免减法出现负数
            cnt[s % mod]++;
        }
        return ans;
    }
    
    // 数组版本，效率更高！
    // 因为 s 至多为 n
    long long countInterestingSubarrays(vector<int> &nums, int mod, int k) {
        int n = nums.size();
        vector<int> cnt(n + 1);
        cnt[0] = 1;
        long long ans = 0;
        int s = 0;
        for (int x: nums) {
            if (x % mod == k)
                s = (s + 1) % mod;
            int s2 = (s - k + mod) % mod;
            if (s2 <= n)
                ans += cnt[s2];
            cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countInterestingSubarrays(nums []int, mod, k int) (ans int64) {
	cnt := map[int]int{0: 1} // 把 s[0]=0 算进去
	s := 0
	for _, x := range nums {
		if x%mod == k {
			s = (s + 1) % mod // 这里取模，下面 cnt[s]++ 就不需要取模了
		}
		ans += int64(cnt[(s-k+mod)%mod]) // +mod 避免减法出现负数
		cnt[s]++
	}
	return
}
```

```js [sol-JavaScript]
var countInterestingSubarrays = function (nums, mod, k) {
    const n = nums.length;
    var cnt = new Array(n + 1).fill(0);
    cnt[0] = 1;
    let ans = 0;
    let s = 0;
    for (const x of nums) {
        if (x % mod === k)
            s = (s + 1) % mod; // 这里取模，下面 cnt[s]++ 就不需要取模了
        const s2 = (s - k + mod) % mod; // +mod 避免减法出现负数
        if (s2 <= n)
            ans += cnt[s2];
        cnt[s]++;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目（前缀和+哈希表）

推荐按照顺序完成。

- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
- [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)
- [523. 连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
