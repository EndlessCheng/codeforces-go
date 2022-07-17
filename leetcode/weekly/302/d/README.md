本题 [视频讲解](https://www.bilibili.com/video/BV1GV4y1J7kc) 已出炉，欢迎点赞三连~

---

#### 提示 1

元素 $x$ 若能整除 $\textit{numsDivide}$ 的所有元素，等价于 $x$ 是所有 $\textit{numsDivide}[i]$ 的因子，这也等价于 $x$ 是 $\textit{numsDivide}$ 所有元素的最大公因数 $g$ 的因子。

#### 提示 2

由于要求用 $\textit{nums}$ 的最小元素去整除 $g$，不妨将 $\textit{nums}$ 排序后，从小到大找到第一个能整除 $g$ 的元素 $x$，所有小于 $x$ 的元素都需要删除。

```py [sol1-Python3]
class Solution:
    def minOperations(self, nums: List[int], numsDivide: List[int]) -> int:
        g = reduce(gcd, numsDivide)
        nums.sort()
        return next((i for i, x in enumerate(nums) if g % x == 0), -1)
```

```java [sol1-Java]
class Solution {
    public int minOperations(int[] nums, int[] numsDivide) {
        var g = 0;
        for (var x : numsDivide) g = gcd(g, x);
        Arrays.sort(nums);
        for (var i = 0; i < nums.length; i++) if (g % nums[i] == 0) return i;
        return -1;
    }

    int gcd(int a, int b) {
        return a == 0 ? b : gcd(b % a, a);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, vector<int> &numsDivide) {
        int g = 0;
        for (int x : numsDivide) g = gcd(g, x);
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size(); i++) if (g % nums[i] == 0) return i;
        return -1;
    }
};
```

```go [sol1-Go]
func minOperations(nums, numsDivide []int) int {
	g := 0
	for _, x := range numsDivide {
		g = gcd(g, x)
	}
	sort.Ints(nums)
	for i, x := range nums {
		if g%x == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

也可以不用排序，通过两次遍历得到答案。

```py [sol2-Python3]
class Solution:
    def minOperations(self, nums: List[int], numsDivide: List[int]) -> int:
        g = reduce(gcd, numsDivide)
        mn = min((x for x in nums if g % x == 0), default=0)
        return sum(x < mn for x in nums) if mn else -1
```

```java [sol2-Java]
class Solution {
    public int minOperations(int[] nums, int[] numsDivide) {
        var g = 0;
        for (var x : numsDivide) g = gcd(g, x);
        var min = Integer.MAX_VALUE;
        for (int num : nums) if (g % num == 0) min = Math.min(min, num);
        if (min == Integer.MAX_VALUE) return -1;
        var ans = 0;
        for (var x : nums) if (x < min) ++ans;
        return ans;
    }

    int gcd(int a, int b) {
        return a == 0 ? b : gcd(b % a, a);
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, vector<int> &numsDivide) {
        int g = 0;
        for (int x : numsDivide) g = gcd(g, x);
        int mn = INT_MAX;
        for (int x : nums) if (g % x == 0) mn = min(mn, x);
        if (mn == INT_MAX) return -1;
        int ans = 0;
        for (int x : nums) if (x < mn) ++ans;
        return ans;
    }
};
```

```go [sol2-Go]
func minOperations(nums, numsDivide []int) (ans int) {
	g := 0
	for _, x := range numsDivide {
		g = gcd(g, x)
	}
	min := math.MaxInt32
	for _, x := range nums {
		if g%x == 0 && x < min {
			min = x
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	for _, x := range nums {
		if x < min {
			ans++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(m+\log U + n)$，其中 $m$ 为数组 $\textit{numsDivide}$ 的长度，$U=\max(\textit{numsDivide})$，$n$ 为数组 $\textit{nums}$ 的长度。注意到求最大公因数 $g$ 的过程（设初始 $g=U$），要么使 $g$ 不变，要么使 $g$ 至少减半，而 $g$ 至多减半 $O(\log U)$ 次，因此求最大公因数的迭代次数为 $O(m+\log U)$ 次。总的时间复杂度为 $O(m+\log U + n)$。
- 空间复杂度：$O(1)$。仅用到若干变量。
