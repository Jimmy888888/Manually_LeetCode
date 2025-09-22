class Solution:
    def fib(self, n: int) -> int:
        # check
        if n == 0 :
            return 0
        if n == 1 or n == 2:
            return 1

        pre1 = 1
        pre2 = 0
        res = 0

        for i in range(2, n+1):
            res = pre1 + pre2
            pre2 = pre1
            pre1 = res
        
        return res