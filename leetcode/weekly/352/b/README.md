首先可以预处理出 $10^6$ 内的所有质数，这个之前周赛出过，我讲了**埃氏筛**和**线性筛**两种做法，可以看 [【周赛 326】](https://www.bilibili.com/video/BV1H8411E7hn/)第四题。

[两种筛法的代码模板（Python/Java/C++/Go）](https://leetcode.cn/problems/closest-prime-numbers-in-range/solution/yu-chu-li-zhi-shu-mei-ju-by-endlesscheng-uw2b/)

然后就是暴力枚举质数 $x$ 和 $y=n-x$ 了，如果 $x\le y$ 且 $y$ 是质数，那么就把 $[x,y]$ 加入答案。

代码实现时，有一些优化之处：如果 $n$ 是奇数，由于只有奇数+偶数=奇数，而偶数中只有 $2$ 是质数，所以如果 $n$ 是奇数时，至多只有一个质数对 $(2,n-2)$。

注：把 $n$ 分解成两个质数的方案有多少呢？请看 [OEIS A061358](https://oeis.org/A061358/graph)。

```py [sol-Python3]
MX = 10 ** 6 + 1
primes = []
is_prime = [True] * MX
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False

class Solution:
    def findPrimePairs(self, n: int) -> List[List[int]]:
        if n % 2:
            return [[2, n - 2]] if n > 4 and is_prime[n - 2] else []
        ans = []
        for x in primes:
            y = n - x
            if y < x:
                break
            if is_prime[y]:
                ans.append([x, y])
        return ans
```

```java [sol-Java]
class Solution {
    private final static int MX = (int) 1e6;
    private final static int[] primes = new int[78498];
    private final static boolean[] np = new boolean[MX + 1];

    static {
        var pi = 0;
        for (var i = 2; i <= MX; ++i) {
            if (!np[i]) {
                primes[pi++] = i;
                for (var j = i; j <= MX / i; ++j) // 避免溢出的写法
                    np[i * j] = true;
            }
        }
    }

    public List<List<Integer>> findPrimePairs(int n) {
        if (n % 2 > 0)
            return n > 4 && !np[n - 2] ? List.of(List.of(2, n - 2)) : List.of();
        var ans = new ArrayList<List<Integer>>();
        for (int x : primes) {
            int y = n - x;
            if (y < x) break;
            if (!np[y]) ans.add(List.of(x, y));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
const int MX = 1e6;
vector<int> primes;
bool np[MX + 1];

int init = []() {
    for (int i = 2; i <= MX; i++) {
        if (!np[i]) {
            primes.push_back(i);
            for (int j = i; j <= MX / i; j++) // 避免溢出的写法
                np[i * j] = true;
        }
    }
    return 0;
}();

class Solution {
public:
    vector<vector<int>> findPrimePairs(int n) {
        vector<vector<int>> ans;
        if (n % 2) {
            if (n > 4 && !np[n - 2])
                ans.push_back({2, n - 2});
            return ans;
        }
        for (int x: primes) {
            int y = n - x;
            if (y < x) break;
            if (!np[y]) ans.push_back({x, y});
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 1e6
var primes []int
var isP = [mx + 1]bool{}

func init() {
	for i := 2; i <= mx; i++ {
		isP[i] = true
	}
	for i := 2; i <= mx; i++ {
		if isP[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				isP[j] = false
			}
		}
	}
}

func findPrimePairs(n int) (ans [][]int) {
	if n%2 > 0 {
		if n > 4 && isP[n-2] {
			return [][]int{{2, n - 2}}
		}
		return
	}
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if isP[y] {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
```

#### 复杂度分析

这里忽略预处理质数的时间和空间。

- 时间复杂度：$\mathcal{O}\left(\dfrac{n}{\log n}\right)$。$n$ 以内有 $\mathcal{O}\left(\dfrac{n}{\log n}\right)$ 个质数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
