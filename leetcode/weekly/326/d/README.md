分两步：

1. 筛质数，做法见 [204. 计数质数](https://leetcode.cn/problems/count-primes/)。
2. 找 $[\textit{left},\textit{right}]$ 范围内的最小质数间隙（prime gap），暴力枚举范围内的所有相邻质数即可。

附：[视频讲解](https://www.bilibili.com/video/BV1H8411E7hn)

#### 优化

1. 找范围内的第一个质数可以用二分查找。
2. 可以往质数表末尾额外插入 $2$ 个 $10^6+1$，这样无需判断下标是否越界。

```py [sol1-Python3]
MX = 10 ** 6 + 1
primes = []
is_prime = [True] * MX
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False
primes.extend((MX, MX))  # 保证下面下标不会越界

class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        p = q = -1
        i = bisect_left(primes, left)
        while primes[i + 1] <= right:
            if p < 0 or primes[i + 1] - primes[i] < q - p:
                p, q = primes[i], primes[i + 1]
            i += 1
        return [p, q]
```

```java [sol1-Java]
class Solution {
    private final static int MX = (int) 1e6;
    private final static int[] primes = new int[78500];

    static {
        var np = new boolean[MX + 1];
        var pi = 0;
        for (var i = 2; i <= MX; ++i)
            if (!np[i]) {
                primes[pi++] = i;
                for (var j = i; j <= MX / i; ++j) // 避免溢出的写法
                    np[i * j] = true;
            }
        primes[pi++] = MX + 1;
        primes[pi++] = MX + 1; // 保证下面下标不会越界
    }

    public int[] closestPrimes(int left, int right) {
        int p = -1, q = -1;
        for (var i = lowerBound(primes, left); primes[i + 1] <= right; ++i)
            if (p < 0 || primes[i + 1] - primes[i] < q - p) {
                p = primes[i];
                q = primes[i + 1];
            }
        return new int[]{p, q};
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
const int MX = 1e6;
vector<int> primes;

int init = []() {
    bool np[MX + 1]{};
    for (int i = 2; i <= MX; ++i)
        if (!np[i]) {
            primes.push_back(i);
            for (int j = i; j <= MX / i; ++j) // 避免溢出的写法
                np[i * j] = true;
        }
    primes.push_back(MX + 1);
    primes.push_back(MX + 1); // 保证下面下标不会越界
    return 0;
}();

class Solution {
public:
    vector<int> closestPrimes(int left, int right) {
        int p = -1, q = -1;
        int i = lower_bound(primes.begin(), primes.end(), left) - primes.begin();
        for (; primes[i + 1] <= right; ++i)
            if (p < 0 || primes[i + 1] - primes[i] < q - p) {
                p = primes[i];
                q = primes[i + 1];
            }
        return {p, q};
    }
};
```

```go [sol1-Go]
const mx = 1e6 + 1
var primes = make([]int, 0, 78500)

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	primes = append(primes, mx, mx) // 保证下面下标不会越界
}

func closestPrimes(left, right int) []int {
	p, q := -1, -1
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}
```

也可以用线性筛（欧拉筛）做，具体原理见 [视频讲解](https://www.bilibili.com/video/BV1H8411E7hn)。

```py [sol2-Python3]
MX = 10 ** 6 + 1
primes = []
is_prime = [True] * MX
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
    for p in primes:
        if i * p >= MX: break
        is_prime[i * p] = False
        if i % p == 0: break
primes.extend((MX, MX))  # 保证下面下标不会越界

class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        p = q = -1
        i = bisect_left(primes, left)
        while primes[i + 1] <= right:
            if p < 0 or primes[i + 1] - primes[i] < q - p:
                p, q = primes[i], primes[i + 1]
            i += 1
        return [p, q]
```

```java [sol2-Java]
class Solution {
    private final static int MX = (int) 1e6;
    private final static int[] primes = new int[78500];

    static {
        var np = new boolean[MX + 1];
        var pi = 0;
        for (var i = 2; i <= MX; ++i) {
            if (!np[i]) primes[pi++] = i;
            for (var j = 0; j < pi; ++j) {
                var p = primes[j];
                if (i * p > MX) break;
                np[i * p] = true;
                if (i % p == 0) break;
            }
        }
        primes[pi++] = MX + 1;
        primes[pi++] = MX + 1; // 保证下面下标不会越界
    }

    public int[] closestPrimes(int left, int right) {
        int p = -1, q = -1;
        for (var i = lowerBound(primes, left); primes[i + 1] <= right; ++i)
            if (p < 0 || primes[i + 1] - primes[i] < q - p) {
                p = primes[i];
                q = primes[i + 1];
            }
        return new int[]{p, q};
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = left + (right - left) / 2;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol2-C++]
const int MX = 1e6;
vector<int> primes;

int init = []() {
    bool np[MX + 1]{};
    for (int i = 2; i <= MX; ++i) {
        if (!np[i]) primes.push_back(i);
        for (int p: primes) {
            if (i * p > MX) break;
            np[i * p] = true;
            if (i % p == 0) break;
        }
    }
    primes.push_back(MX + 1);
    primes.push_back(MX + 1); // 保证下面下标不会越界
    return 0;
}();

class Solution {
public:
    vector<int> closestPrimes(int left, int right) {
        int p = -1, q = -1;
        int i = lower_bound(primes.begin(), primes.end(), left) - primes.begin();
        for (; primes[i + 1] <= right; ++i)
            if (p < 0 || primes[i + 1] - primes[i] < q - p) {
                p = primes[i];
                q = primes[i + 1];
            }
        return {p, q};
    }
};
```

```go [sol2-Go]
const mx = 1e6 + 1
var primes = make([]int, 0, 78500)

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= mx {
				break
			}
			np[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
	primes = append(primes, mx, mx) // 保证下面下标不会越界
}

func closestPrimes(left, right int) []int {
	p, q := -1, -1
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}
```


#### 复杂度分析

- 时间复杂度：$O(\textit{right})$，忽略预处理的时间复杂度。严谨地说，由于范围内有 $O\left(\dfrac{\textit{right}}{\log\textit{right}}-\dfrac{\textit{left}}{\log\textit{left}}\right)$ 个质数（根据质数密度），所以遍历的时间复杂度为 $O\left(\dfrac{\textit{right}}{\log\textit{right}}-\dfrac{\textit{left}}{\log\textit{left}}\right)$，再算上二分质数的时间 $O(\log\pi(U))$（$\pi(U)$ 表示 $U=10^6$ 内的质数个数，这有 $78498$ 个），总的时间复杂度为 $O\left(\log\pi(U) + \dfrac{\textit{right}}{\log\textit{right}}-\dfrac{\textit{left}}{\log\textit{left}}\right)$。
- 空间复杂度：$O(1)$，忽略预处理的空间复杂度。仅用到若干变量。
