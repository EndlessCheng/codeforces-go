下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

树状数组/线段树逐渐成为周赛必备技能了。

本题用到的技巧是，合并下标相同的元素。

式子变形得

$$
\textit{nums}_1[i]-\textit{nums}_2[i]\le\textit{nums}_1[j]-\textit{nums}_2[j]+\textit{diff}
$$

记 $a[i]=\textit{nums}_1[i]-\textit{nums}_2[i]$，上式为

$$
a[i]\le a[j]+\textit{diff}
$$

因此本题和 [剑指 Offer 51. 数组中的逆序对](https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)、[315. 计算右侧小于当前元素的个数](https://leetcode.cn/problems/count-of-smaller-numbers-after-self/) 等题目实质上是同一类题，用**归并排序**或者**树状数组**等均可以通过。下午的直播会讲这两种做法。

**下面代码用的离散化树状数组，即使元素范围达到 $10^9$ 也适用。**

```py [sol1-Python3]
class Solution:
    def numberOfPairs(self, a: List[int], nums2: List[int], diff: int) -> int:
        for i, x in enumerate(nums2):
            a[i] -= x
        b = a.copy()
        b.sort()  # 配合下面的二分，离散化

        ans = 0
        t = BIT(len(a) + 1)
        for x in a:
            ans += t.query(bisect_right(b, x + diff))
            t.add(bisect_left(b, x) + 1)
        return ans

class BIT:
    def __init__(self, n):
        self.tree = [0] * n

    def add(self, x):
        while x < len(self.tree):
            self.tree[x] += 1
            x += x & -x

    def query(self, x):
        res = 0
        while x > 0:
            res += self.tree[x]
            x &= x - 1
        return res
```

```java [sol1-Java]
class Solution {
    public long numberOfPairs(int[] a, int[] nums2, int diff) {
        var n = a.length;
        for (var i = 0; i < n; ++i)
            a[i] -= nums2[i];
        var b = a.clone();
        Arrays.sort(b); // 配合下面的二分，离散化

        var ans = 0L;
        var t = new BIT(n + 1);
        for (var x : a) {
            ans += t.query(lowerBound(b, x + diff + 1));
            t.add(lowerBound(b, x) + 1);
        }
        return ans;
    }

    private int lowerBound(int[] a, int x) {
        int left = 0, right = a.length;
        while (left < right) {
            var mid = left + (right - left) / 2;
            if (a[mid] < x) left = mid + 1;
            else right = mid;
        }
        return left;
    }
}

class BIT {
    private final int[] tree;

    public BIT(int n) {
        tree = new int[n];
    }

    public void add(int x) {
        while (x < tree.length) {
            ++tree[x];
            x += x & -x;
        }
    }

    public int query(int x) {
        var res = 0;
        while (x > 0) {
            res += tree[x];
            x &= x - 1;
        }
        return res;
    }
}
```

```cpp [sol1-C++]
class BIT {
private:
    vector<int> tree;

public:
    BIT(int n) : tree(n) {}

    void add(int x) {
        while (x < tree.size()) {
            ++tree[x];
            x += x & -x;
        }
    }

    int query(int x) {
        int res = 0;
        while (x > 0) {
            res += tree[x];
            x &= x - 1;
        }
        return res;
    }
};

class Solution {
public:
    long long numberOfPairs(vector<int> &a, vector<int> &nums2, int diff) {
        int n = a.size();
        for (int i = 0; i < n; ++i)
            a[i] -= nums2[i];
        auto b = a;
        sort(b.begin(), b.end()); // 配合下面的二分，离散化

        long ans = 0L;
        auto t = new BIT(n + 1);
        for (int x : a) {
            ans += t->query(upper_bound(b.begin(), b.end(), x + diff) - b.begin());
            t->add(lower_bound(b.begin(), b.end(), x) - b.begin() + 1);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func numberOfPairs(a, nums2 []int, diff int) (ans int64) {
	for i, x := range nums2 {
		a[i] -= x
	}
	b := append(sort.IntSlice{}, a...)
	b.Sort() // 配合下面的二分，离散化

	t := make(BIT, len(a)+1)
	for _, x := range a {
		ans += int64(t.query(b.Search(x + diff + 1)))
		t.add(b.Search(x) + 1)
	}
	return
}

type BIT []int

func (t BIT) add(x int) {
	for x < len(t) {
		t[x]++
		x += x & -x
	}
}

func (t BIT) query(x int) (res int) {
	for x > 0 {
		res += t[x]
		x &= x - 1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$O(n)$。
