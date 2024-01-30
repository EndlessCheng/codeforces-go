package copypasta

/*
朗伯 W 函数   Lambert W function
https://en.wikipedia.org/wiki/Lambert_W_function
Bounds and inequalities https://en.wikipedia.org/wiki/Lambert_W_function#Bounds_and_inequalities
W(x) = Θ(log x)

https://math.stackexchange.com/questions/433717/how-to-solve-equations-with-logarithms-like-this-ax-b-logx-c-0
a*x + b*ln(x) + c = 0 的解：
x = (b/a)*W((a/b)*e^(-c/b))
注意 W(x) 当 x < 0 的时候有两个值 W_0(x) 和 W_{-1}(x) 

用于一些题目的时间复杂度的说明上
LC2749 https://leetcode.cn/problems/minimum-operations-to-make-the-integer-zero/

一元五次方程的根 布灵根式 Bring radical https://en.wikipedia.org/wiki/Bring_radical
*/
