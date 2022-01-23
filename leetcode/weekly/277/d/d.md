二进制枚举

由于至多有 $n=15$ 个人，我们可以枚举这 $n$ 个人谁是好人，谁是坏人，这一共有 $2^n$ 种不同的情况。

我们可以用二进制数表示这 $n$ 个人中谁好谁坏，其中 $1$ 表示好人，$0$ 表示坏人。这样就可以枚举 $i\in [0, 2^n-1]$ 中的所有数字，然后判断 $i$ 中好人的陈述是否与实际情况矛盾，若不矛盾则 $i$ 为一种合法的情况。所有合法情况中的好人个数的最大值即为答案。

代码实现时，可以从 $i=1$ 开始枚举。

```go [sol1-Go]
func maximumGood(statements [][]int) (ans int) {
next:
	for i := 1; i < 1<<len(statements); i++ {
		cnt := 0 // i 中好人个数
		for j, row := range statements {
			if i>>j&1 == 1 { // 枚举 i 中的好人 j
				for k, st := range row { // 枚举 j 的所有陈述 st
					if st < 2 && st != i>>k&1 { // 该陈述与实际情况矛盾
						continue next
					}
				}
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
```

```C++ [sol1-C++]
class Solution {
public:
    int maximumGood(vector<vector<int>> &statements) {
        int ans = 0, n = statements.size();
        for (int i = 1; i < 1 << n; ++i) {
            int cnt = 0; // i 中好人个数
            for (int j = 0; j < n; ++j) {
                if ((i >> j) & 1) { // 枚举 i 中的好人 j
                    for (int k = 0; k < n; ++k) { // 枚举 j 的所有陈述
                        if (statements[j][k] < 2 && statements[j][k] != ((i >> k) & 1)) { // 该陈述与实际情况矛盾
                            goto next;
                        }
                    }
                    ++cnt;
                }
            }
            ans = max(ans, cnt);
            next:;
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def maximumGood(self, statements: List[List[int]]) -> int:
        def check(i: int) -> int:
            cnt = 0  # i 中好人个数
            for j, row in enumerate(statements):  # 枚举 i 中的好人 j
                if (i >> j) & 1:
                    if any(st < 2 and st != (i >> k) & 1 for k, st in enumerate(row)):
                        return 0  # 好人 j 的某个陈述 st 与实际情况矛盾
                    cnt += 1
            return cnt

        return max(check(i) for i in range(1, 1 << len(statements)))
```

```java [sol1-Java]
class Solution {
    public int maximumGood(int[][] statements) {
        var ans = 0;
        var n = statements.length;
        next:
        for (var i = 1; i < 1 << n; i++) {
            var cnt = 0; // i 中好人个数
            for (var j = 0; j < n; j++) {
                if (((i >> j) & 1) == 1) { // 枚举 i 中的好人 j
                    for (var k = 0; k < n; k++) { // 枚举 j 的所有陈述
                        if (statements[j][k] < 2 && statements[j][k] != ((i >> k) & 1)) { // 该陈述与实际情况矛盾
                            continue next;
                        }
                    }
                    cnt++;
                }
            }
            ans = Math.max(ans, cnt);
        }
        return ans;
    }
}
```
