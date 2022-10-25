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

代码实现时可以先奇数再偶数，然后奇数偶数内部再排序。

由于数组元素都是正数，可以先**把所有奇数变成相反数**，然后排序，奇偶就自动分开了。

```py [sol1-Python3]
def f(a: List[int]) -> None:
    for i, x in enumerate(a):
        if x % 2: a[i] = -x  # 由于元素都是正数，把奇数变成相反数，这样排序后奇偶就自动分开了
    a.sort()

class Solution:
    def makeSimilar(self, nums: List[int], target: List[int]) -> int:
        f(nums)
        f(target)
        return sum(abs(x - y) for x, y in zip(nums, target)) // 4
```

```java [sol1-Java]
class Solution {
    public long makeSimilar(int[] nums, int[] target) {
        f(nums);
        f(target);
        var ans = 0L;
        for (var i = 0; i < nums.length; ++i)
            ans += Math.abs(nums[i] - target[i]);
        return ans / 4;
    }

    private void f(int[] a) {
        // 由于元素都是正数，把奇数变成相反数，这样排序后奇偶就自动分开了
        for (var i = 0; i < a.length; ++i)
            if (a[i] % 2 != 0) a[i] = -a[i];
        Arrays.sort(a);
    }
}
```

```cpp [sol1-C++]
class Solution {
    void f(vector<int> &a) {
        for (int &x : a)
            if (x % 2) x = -x; // 由于元素都是正数，把奇数变成相反数，这样排序后奇偶就自动分开了
        sort(a.begin(), a.end());
    }

public:
    long long makeSimilar(vector<int> &nums, vector<int> &target) {
        f(nums);
        f(target);
        long long ans = 0L;
        for (int i = 0; i < nums.size(); ++i)
            ans += abs(nums[i] - target[i]);
        return ans / 4;
    }
};
```

```go [sol1-Go]
func f(a []int) {
	for i, x := range a {
		if x%2 > 0 {
			a[i] = -x // 由于元素都是正数，把奇数变成相反数，这样排序后奇偶就自动分开了
		}
	}
	sort.Ints(a)
}

func makeSimilar(nums, target []int) (ans int64) {
	f(nums)
	f(target)
	for i, x := range nums {
		ans += int64(abs(x - target[i]))
	}
	return ans / 4
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略快排的栈开销。
