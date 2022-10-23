[视频讲解](https://www.bilibili.com/video/BV1ne4y1e7nu) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

#### 提示 1

如果把问题中的 $+2$ 和 $-2$ 改成 $+1$ 和 $-1$，要怎么做？

例如 $\textit{nums}=[2,8]$，$\textit{target}=[4,6]$，那么应该让 $2$ 和 $4$ 一对，$8$ 和 $6$ 一对。如果让 $2$ 和 $6$ 一对，$8$ 和 $4$ 一对，是会让变化量的和变得更大的。

通过这种**邻项交换法**，我们可以证明，让最小的一对，次小的一对，第三小的一对，……，累加每对元素的差的绝对值，就得到了每个数的变化量的和的最小值。

#### 提示 2

回到原问题，$+2$ 和 $-2$ 会导致无法直接排序然后一一匹配，但注意到 $+2$ 和 $-2$ 并不会改变元素的**奇偶性**，因此我们可以把偶数分为一组，奇数分为一组，每组分别计算，这样就像提示 1 那样一一匹配了。

最后把变化量的和除以 $4$，即为答案。

#### 提示 3

代码实现时有两种做法，可以按照先偶数再奇数，然后奇数偶数内部再排序；也可以直接排序，然后用两个指针在 $\textit{target}$ 上遍历，分别表示偶数元素的下标和奇数元素的下标。

由于排序是瓶颈，尽量减少排序的时间开销比较好，所以下面用的是直接排序的做法。

```py [sol1-Python3]
class Solution:
    def makeSimilar(self, nums: List[int], target: List[int]) -> int:
        nums.sort()
        target.sort()
        ans, j = 0, [0, 0]  # 用数组表示两个下标，这样不用讨论奇偶性
        for x in nums:
            p = x % 2
            while target[j[p]] % 2 != p:  # 找 target 中奇偶性相同的元素
                j[p] += 1
            ans += abs(x - target[j[p]])
            j[p] += 1
        return ans // 4
```

```java [sol1-Java]
class Solution {
    public long makeSimilar(int[] nums, int[] target) {
        Arrays.sort(nums);
        Arrays.sort(target);
        var ans = 0L;
        var j = new int[2]; // 用数组表示两个下标，这样不用讨论奇偶性
        for (var x : nums) {
            var p = x % 2;
            while (target[j[p]] % 2 != p) ++j[p]; // 找 target 中奇偶性相同的元素
            ans += Math.abs(x - target[j[p]++]);
        }
        return ans / 4;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long makeSimilar(vector<int> &nums, vector<int> &target) {
        sort(nums.begin(), nums.end());
        sort(target.begin(), target.end());
        long long ans = 0L;
        int js[2]{}; // 用数组表示两个下标，这样不用讨论奇偶性
        for (int x : nums) {
            int p = x % 2, &j = js[p];
            while (target[j] % 2 != p) ++j; // 找 target 中奇偶性相同的元素
            ans += abs(x - target[j++]);
        }
        return ans / 4;
    }
};
```

```go [sol1-Go]
func makeSimilar(nums, target []int) (ans int64) {
	sort.Ints(nums)
	sort.Ints(target)
	j := [2]int{} // 用数组表示两个下标，这样不用讨论奇偶性
	for _, x := range nums {
		p := x % 2
		for target[j[p]]%2 != p { // 找 target 中奇偶性相同的元素
			j[p]++
		}
		ans += int64(abs(x - target[j[p]]))
		j[p]++
	}
	return ans / 4
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略快排的栈开销。
