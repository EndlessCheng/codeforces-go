#### 提示 1

如果有三个相邻且可以合并的数 $x,y,z$，那么先合并 $x$ 和 $y$ 再合并 $z$ 的结果，与先合并 $y$ 和 $z$ 再合并 $x$ 的结果是一样的。

#### 提示 2

不妨总是优先与左侧元素合并，即：如果合并后的数能与左侧元素非互质，那么就合并左侧的元素，否则尝试合并右侧的元素。

#### 提示 3

用栈来模拟上述过程。

---

创建一个栈，初始元素为 $\textit{nums}[0]$。

从 $\textit{nums}[1]$ 开始遍历，将其入栈。循环，每次从栈顶取出两个元素，若其互质则退出循环，否则将这两个元素的最小公倍数入栈，循环直至栈内只有一个元素为止。

遍历 $\textit{nums}$ 结束后的栈就是答案。

```Python [sol1-Python3]
class Solution:
    def replaceNonCoprimes(self, nums: List[int]) -> List[int]:
        s = [nums[0]]
        for num in nums[1:]:
            s.append(num)
            while len(s) > 1:
                x, y = s[-1], s[-2]
                g = gcd(x, y)
                if g == 1:
                    break
                s.pop()
                s[-1] *= x // g
        return s
```

```go [sol1-Go]
func replaceNonCoprimes(nums []int) []int {
	s := []int{nums[0]}
	for _, num := range nums[1:] {
		s = append(s, num)
		for len(s) > 1 {
			x, y := s[len(s)-1], s[len(s)-2]
			g := gcd(x, y)
			if g == 1 {
				break
			}
			s = s[:len(s)-1]
			s[len(s)-1] *= x / g
		}
	}
	return s
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

```C++ [sol1-C++]
class Solution {
public:
    vector<int> replaceNonCoprimes(vector<int> &nums) {
        vector<int> s = {nums[0]};
        for (int i = 1; i < nums.size(); ++i) {
            s.push_back(nums[i]);
            while (s.size() > 1) {
                int x = s.back(), y = s[s.size() - 2];
                int g = gcd(x, y);
                if (g == 1) break;
                s.pop_back();
                s.back() *= x / g;
            }
        }
        return s;
    }
};
```

```java [sol1-Java]
class Solution {
    public List<Integer> replaceNonCoprimes(int[] nums) {
        var s = new ArrayList<Integer>();
        s.add(nums[0]);
        for (var i = 1; i < nums.length; ++i) {
            s.add(nums[i]);
            while (s.size() > 1) {
                var x = s.get(s.size() - 1);
                var y = s.get(s.size() - 2);
                var g = gcd(x, y);
                if (g == 1) break;
                s.remove(s.size() - 1);
                s.set(s.size() - 1, x / g * y);
            }
        }
        return s;
    }

    int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }
}
```
