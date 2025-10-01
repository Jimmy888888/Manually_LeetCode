class Solution:
    def tribonacci(self, n: int) -> int:
        # check n
        match n :
            case 0 :
                return 0
            case 1 :
                return 1
            case 2 :
                return 1

        
        # recorders
        t1 = 0
        t2 = 1
        t3 = 1
        res = 0

        # stat transition equation
        for i in range(3, n+1):
            res = t1 + t2 + t3
            t1 = t2
            t2 = t3
            t3 = res
        
        return res