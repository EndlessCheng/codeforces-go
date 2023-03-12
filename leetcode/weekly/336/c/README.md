### 前置知识：前缀异或和

下文中 $\oplus$ 表示异或运算

对于数组 $\textit{nums}$，定义它的前缀异或和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \bigoplus\limits_{j=0}^{i}\textit{nums}[j]$。

根据这个定义，有 $s[i+1]=s[i]\oplus\textit{nums}[i]$。

例如 $\textit{nums}=[4,3,1,2,4]$，对应的前缀异或和数组为 $s=[0,4,7,6,4,0]$。

通过前缀异或和，我们可以把**子数组的异或和转换成两个前缀异或和的异或**，即

$$
\bigoplus_{j=\textit{left}}^{\textit{right}}\textit{nums}[j] = \bigoplus\limits_{j=0}^{\textit{right}}\textit{nums}[j] \oplus \bigoplus\limits_{j=0}^{\textit{left}-1}\textit{nums}[j] = \textit{s}[\textit{right}+1]\oplus \textit{s}[\textit{left}]
$$

例如 $\textit{nums}$ 的子数组 $[3,1,2]$ 的异或和就可以用 $s[4]\oplus s[1]=4\oplus 4=0$ 算出来。

> 注：为方便计算，常用左闭右开区间 $[\textit{left},\textit{right})$ 来表示从 $\textit{nums}[\textit{left}]$ 到 $\textit{nums}[\textit{right}-1]$ 的子数组，此时子数组的异或和为 $\textit{s}[\textit{right}] \oplus \textit{s}[\textit{left}]$。
>
> 注 2：$s[0]=0$ 表示一个空数组的异或和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $\textit{nums}[0]$ 开始），你要用 $s[\textit{right}]$ 异或谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀异或和的异或。

### 提示 1

由于每次操作的都是同一个比特位，可以把每一位单独看。

### 提示 2

每次都去掉两个 $1$，要是美丽子数组，需要子数组内这个比特位的 $1$ 的个数是偶数。

### 提示 3

由于 $1\oplus 1=0$，把所有比特位合起来看，美丽子数组这等价于子数组的异或和等于 $0$。

### 提示 4

利用前缀异或和 $s$，问题相当于在求 $s$ 中有多少对 $s[\textit{left}]$ 和 $s[\textit{right}]$ 满足 $s[\textit{right}]\oplus s[\textit{left}] = 0$，即 $s[\textit{right}]= s[\textit{left}]$，因为异或为 $0$ 的两个数字必然相等。

也就是说，我们实际计算的是 $s$ 中有多少对相同数字。

我们可以一边遍历 $s$，一边用一个哈希表 $\textit{cnt}$ 统计 $s[i]$ 的出现次数，累加遍历中的 $\textit{cnt}[s[i]]$，即为答案。

附：[视频讲解](https://www.bilibili.com/video/BV1d54y1M7Qg/)

```py [sol1-Python3]
class Solution:
    def beautifulSubarrays(self, nums: List[int]) -> int:
        s = list(accumulate(nums, xor, initial=0))
        ans, cnt = 0, Counter()
        for x in s:
            # 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
            ans += cnt[x]
            cnt[x] += 1
        return ans
```

```java [sol1-Java]
class Solution {
    public long beautifulSubarrays(int[] nums) {
        long ans = 0;
        int n = nums.length;
        var s = new int[n + 1];
        for (int i = 0; i < n; ++i)
            s[i + 1] = s[i] ^ nums[i];
        var cnt = new HashMap<Integer, Integer>();
        for (int x : s) {
            // 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
            ans += cnt.getOrDefault(x, 0);
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long beautifulSubarrays(vector<int> &nums) {
        long long ans = 0;
        int n = nums.size();
        vector<int> s(n + 1);
        for (int i = 0; i < n; ++i)
            s[i + 1] = s[i] ^ nums[i];
        unordered_map<int, int> cnt;
        for (int x : s)
            // 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
            ans += cnt[x]++;
        return ans;
    }
};
```

```go [sol1-Go]
func beautifulSubarrays(nums []int) (ans int64) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] ^ x
	}
	cnt := map[int]int{}
	for _, x := range s {
		// 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
		ans += int64(cnt[x])
		cnt[x]++
	}
	return
}
```

### 优化

也可以一边计算 $s$，一边统计答案。

```py [sol2-Python3]
class Solution:
    def beautifulSubarrays(self, nums: List[int]) -> int:
        ans = s = 0
        cnt = Counter({s: 1})  # s[0]
        for x in nums:
            s ^= x
            ans += cnt[s]
            cnt[s] += 1
        return ans
```

```java [sol2-Java]
class Solution {
    public long beautifulSubarrays(int[] nums) {
        long ans = 0;
        int s = 0;
        var cnt = new HashMap<Integer, Integer>();
        cnt.put(s, 1); // s[0]
        for (int x : nums) {
            s ^= x;
            ans += cnt.getOrDefault(s, 0);
            cnt.merge(s, 1, Integer::sum);
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long beautifulSubarrays(vector<int> &nums) {
        long long ans = 0;
        int s = 0;
        unordered_map<int, int> cnt{{s, 1}}; // s[0]
        for (int x : nums) {
            s ^= x;
            ans += cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol2-Go]
func beautifulSubarrays(nums []int) (ans int64) {
	s := 0
	cnt := map[int]int{s: 1} // s[0]
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

### 相似题目

推荐按顺序做。

- [1. 两数之和](https://leetcode.cn/problems/two-sum/)
- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
- [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)
- [1590. 使数组和能被 P 整除](https://leetcode.cn/problems/make-sum-divisible-by-p/)
- [523. 连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
- [面试题 17.05. 字母与数字](https://leetcode.cn/problems/find-longest-subarray-lcci/)
- [1915. 最美子字符串的数目](https://leetcode.cn/problems/number-of-wonderful-substrings/)
- [930. 和相同的二元子数组](https://leetcode-cn.com/problems/binary-subarrays-with-sum/)
- [1371. 每个元音包含偶数次的最长子字符串](https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)
- [1542. 找出最长的超赞子字符串](https://leetcode-cn.com/problems/find-longest-awesome-substring/)
