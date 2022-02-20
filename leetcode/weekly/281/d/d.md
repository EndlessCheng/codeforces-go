#### 提示 1

枚举 $\textit{nums}[j]$，求有多少个符合条件的 $\textit{nums}[i]$。

#### 提示 2

对于一个固定的 $\textit{nums}[j]$，$\textit{nums}[i]$ 必须是某个数 $x$ 的倍数。显然 $\textit{nums}[i]$ 可以是 $x=k$ 的倍数，但为了找到所有的 $\textit{nums}[i]$，$x$ 应当尽可能地小，这个最小的 $x$ 是多少？

---

#### 方法一：统计每个数的因子

从因子的角度考虑。如果 $\textit{nums}[j]$ 和 $k$ 有一些公因子，那么可以从 $k$ 中除去这些公因子，这样 $x$ 会变小，那么除去 $\text{GCD}(\textit{nums}[j],k)$ 是最优的，因此 $\textit{nums}[i]$ 必须是 $x=\dfrac{k}{\text{GCD}(\textit{nums}[j],k)}$ 的倍数。

如何统计是 $x$ 的倍数的 $\textit{nums}[i]$ 呢？我们可以预处理每个数的所有因子，用一个哈希表 $\textit{cnt}$ 统计 $\textit{nums}[j]$ 前面每个数的每个因子的出现次数，这样对于一个固定的 $\textit{nums}[j]$，符合条件的 $\textit{nums}[i]$ 的个数就是 $\textit{cnt}[x]$。累加所有个数即为答案。

在计算一个数的所有因子时，可以枚举因子（见 Java 代码），也可以用 $O(M\log M)$ 的时间预处理出来，这里 $M=10^5$。

注：$[1,10^5]$ 内 $83160$ 的因子个数最多，有 $128$ 个。

```go [sol1-Go]
const mx int = 1e5
var divisors [mx + 1][]int

func init() { // 预处理每个数的所有因子
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func coutPairs(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	for _, v := range nums {
		ans += int64(cnt[k/gcd(v, k)])
		for _, d := range divisors[v] {
			cnt[d]++
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

```C++ [sol1-C++]
const int mx = 100001;
vector<int> divisors[mx + 1];

int init = []() { // 预处理每个数的所有因子
    for (int i = 1; i < mx; ++i)
        for (int j = i; j < mx; j += i)
            divisors[j].push_back(i);
    return 0;
}();

class Solution {
public:
    long long coutPairs(vector<int> &nums, int k) {
        long long ans = 0;
        unordered_map<int, int> cnt;
        for (int v : nums) {
            ans += cnt[k / gcd(v, k)];
            for (int d : divisors[v])
                ++cnt[d];
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
MX = 100001
divisors = [[] for _ in range(MX)]
for i in range(1, MX):  # 预处理每个数的所有因子
    for j in range(i, MX, i):
        divisors[j].append(i)

class Solution:
    def coutPairs(self, nums: List[int], k: int) -> int:
        ans = 0
        cnt = defaultdict(int)
        for v in nums:
            ans += cnt[k / gcd(v, k)]
            for d in divisors[v]:
                cnt[d] += 1
        return ans
```

```java [sol1-Java]
class Solution {
    public long coutPairs(int[] nums, int k) {
        var ans = 0L;
        var cnt = new HashMap<Integer, Integer>();
        for (var v : nums) {
            ans += cnt.getOrDefault(k / gcd(v, k), 0);
            for (var i = 1; i * i <= v; ++i) { // 统计 v 的每个因子
                if (v % i == 0) {
                    cnt.put(i, cnt.getOrDefault(i, 0) + 1);
                    if (i * i <= v) {
                        cnt.put(v / i, cnt.getOrDefault(v / i, 0) + 1);
                    }
                }
            }
        }
        return ans;
    }

    static int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```

#### 方法二：统计 $k$ 的因子

注意到 $x$ 是 $k$ 的因子，因此我们可以将方法一中「统计 $v$ 的因子」改为「统计 $v$ 是 $k$ 的哪些因子的倍数」，这可以通过枚举 $k$ 的所有因子来判断。

```go [sol2-Go]
func coutPairs(nums []int, k int) (ans int64) {
	divisors := []int{} 
	for d := 1; d*d <= k; d++ { // 预处理 k 的所有因子
		if k%d == 0 {
			divisors = append(divisors, d)
			if d*d < k {
				divisors = append(divisors, k/d)
			}
		}
	}
	cnt := map[int]int{}
	for _, v := range nums {
		ans += int64(cnt[k/gcd(v, k)])
		for _, d := range divisors {
			if v%d == 0 {
				cnt[d]++
			}
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

```C++ [sol2-C++]
class Solution {
public:
    long long coutPairs(vector<int> &nums, int k) {
        vector<int> divisors; 
        for (int d = 1; d * d <= k; ++d) { // 预处理 k 的所有因子
            if (k % d == 0) {
                divisors.push_back(d);
                if (d * d < k) divisors.push_back(k / d);
            }
        }
        long long ans = 0;
        unordered_map<int, int> cnt;
        for (int v : nums) {
            ans += cnt[k / gcd(v, k)];
            for (int d : divisors)
                if (v % d == 0) ++cnt[d];
        }
        return ans;
    }
};
```

```Python [sol2-Python3]
class Solution:
    def coutPairs(self, nums: List[int], k: int) -> int:
        divisors = []  
        d = 1
        while d * d <= k:  # 预处理 k 的所有因子
            if k % d == 0:
                divisors.append(d)
                if d * d < k:
                    divisors.append(k / d)
            d += 1
        ans = 0
        cnt = defaultdict(int)
        for v in nums:
            ans += cnt[k / gcd(v, k)]
            for d in divisors:
                if v % d == 0:
                    cnt[d] += 1
        return ans
```

```java [sol2-Java]
class Solution {
    public long coutPairs(int[] nums, int k) {
        var divisors = new ArrayList<Integer>(); // 预处理 k 的所有因子
        for (var d = 1; d * d <= k; d++) {
            if (k % d == 0) {
                divisors.add(d);
                if (d * d < k) divisors.add(k / d);
            }
        }
        var ans = 0L;
        var cnt = new HashMap<Integer, Integer>();
        for (var v : nums) {
            ans += cnt.getOrDefault(k / gcd(v, k), 0);
            for (var d : divisors)
                if (v % d == 0)
                    cnt.put(d, cnt.getOrDefault(d, 0) + 1);
        }
        return ans;
    }

    static int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```
