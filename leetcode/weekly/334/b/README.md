[视频讲解](https://www.bilibili.com/video/BV1wj411G7sH/)

设 $a=k_1m+r_1,\ b=k_2m+r_2$。

那么 $(a+b)\bmod m = (r_1+r_2)\bmod m = (a\bmod m + b\bmod m)\bmod m$。

这意味着我们可以在计算中取模，而不是到最后才取模。 

从左到右遍历 $\textit{word}$。初始化 $x=0$，每遇到一个数字 $d$，就把 $x$ 更新为 $(10x+d)\bmod m$。

```py [sol-Python3]
class Solution:
    def divisibilityArray(self, word: str, m: int) -> List[int]:
        ans = []
        x = 0
        for d in map(int, word):
            x = (x * 10 + d) % m
            ans.append(0 if x else 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] divisibilityArray(String word, int m) {
        char[] s = word.toCharArray();
        int[] ans = new int[s.length];
        long x = 0;
        for (int i = 0; i < s.length; i++) {
            x = (x * 10 + (s[i] - '0')) % m;
            if (x == 0) {
                ans[i] = 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> divisibilityArray(string word, int m) {
        vector<int> ans(word.length());
        long long x = 0;
        for (int i = 0; i < word.length(); i++) {
            x = (x * 10 + (word[i] - '0')) % m;
            ans[i] = x == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	for i, c := range word {
		x = (x*10 + int(c-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var divisibilityArray = function(word, m) {
    const ans = Array(word.length);
    let x = 0;
    for (let i = 0; i < word.length; i++) {
        x = (x * 10 + (word.charCodeAt(i) - '0'.charCodeAt(0))) % m;
        ans[i] = x === 0 ? 1 : 0;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn divisibility_array(word: String, m: i32) -> Vec<i32> {
        let mut ans = vec![0; word.len()];
        let mut x = 0i64;
        for (i, &c) in word.as_bytes().iter().enumerate() {
            x = (x * 10 + (c - b'0') as i64) % m as i64;
            if x == 0 {
                ans[i] = 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值的空间不计入。Java 忽略 $s$ 的空间。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
