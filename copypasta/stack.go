package copypasta

/*
栈消除
LC735 https://leetcode.cn/problems/asteroid-collision/
- 变形：碰撞后，大行星的质量会减去的行星的质量
- 变形 2：你可以决定每颗行星的方向，问最后的行星质量之和最小是多少      
-- f[i][j] 表示前 i 颗行星，向右的质量之和为 j 时，向左的质量之和的最小值
LC2751 https://leetcode.cn/problems/robot-collisions/
https://codeforces.com/problemset/problem/926/E 1900 难度虚高
https://atcoder.jp/contests/abc328/tasks/abc328_d

栈的应用
栈+懒删除 https://codeforces.com/problemset/problem/1000/F
https://codeforces.com/problemset/problem/190/C 1500
https://codeforces.com/problemset/problem/1092/D1
https://codeforces.com/problemset/problem/1092/D2

括号匹配/有效括号
- [32. 最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/)
- [1614. 括号的最大嵌套深度](https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses/) 1323
- [1963. 使字符串平衡的最小交换次数](https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-string-balanced/) 1689
- [2116. 判断一个括号字符串是否有效](https://leetcode.cn/problems/check-if-a-parentheses-string-can-be-valid/) 2038
下标范围有限制 https://codeforces.com/problemset/problem/543/D
https://leetcode.cn/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/

对顶栈
LC2296 https://leetcode.cn/problems/design-a-text-editor/
http://acm.hdu.edu.cn/showproblem.php?pid=4699
*/
