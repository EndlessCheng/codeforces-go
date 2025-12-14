按照题意模拟即可。

[本题视频讲解](https://www.bilibili.com/video/BV1a1meBiETs/?t=1m46s)，欢迎点赞关注~

```py [sol-Python3]
def count_vowel(s: str) -> int:
    return sum(c in "aeiou" for c in s)

class Solution:
    def reverseWords(self, s: str) -> str:
        a = s.split()
        cnt0 = count_vowel(a[0])
        for i in range(1, len(a)):
            if count_vowel(a[i]) == cnt0:
                a[i] = a[i][::-1]
        return ' '.join(a)
```

```java [sol-Java]
class Solution {
    public String reverseWords(String s) {
        String[] a = s.split(" ");
        int cnt0 = countVowel(a[0]);
        for (int i = 1; i < a.length; i++) {
            if (countVowel(a[i]) == cnt0) {
                a[i] = new StringBuilder(a[i]).reverse().toString();
            }
        }
        return String.join(" ", a);
    }

    private int countVowel(String s) {
        int cnt = 0;
        for (char c : s.toCharArray()) {
            if ("aeiou".indexOf(c) >= 0) {
                cnt++;
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
#include<ranges>
class Solution {
    constexpr static string VOWELS = "aeiou";

    template<ranges::input_range R>
    int count_vowel(const R& s) {
        int vowel = 0;
        for (char c : s) {
            if (VOWELS.find(c) != string::npos) {
                vowel++;
            }
        }
        return vowel;
    }

public:
    string reverseWords(string s) {
        int cnt0 = -1;
        for (auto t : s | views::split(' ')) {
            int cnt = count_vowel(t);
            if (cnt0 < 0) {
                cnt0 = cnt;
            } else if (cnt == cnt0) {
                ranges::reverse(t);
            }
        }
        return s;
    }
};
```

```go [sol-Go]
func countVowel(s string) (vowel int) {
	for _, c := range s {
		if strings.IndexRune("aeiou", c) >= 0 {
			vowel++
		}
	}
	return
}

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	cnt0 := countVowel(a[0])
	for i := 1; i < len(a); i++ {
		if countVowel(a[i]) == cnt0 {
			t := []byte(a[i])
			slices.Reverse(t)
			a[i] = string(t)
		}
	}
	return strings.Join(a, " ")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，取决于能否原地修改字符串。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
