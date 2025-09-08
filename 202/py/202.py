class Solution:
    def isHappy(self, n: int) -> bool:
        seen = set()
        s = 0
        while n != 1 and n not in seen:
            seen.add(n)
            while n > 0 :
                num = n % 10
                s += num * num
                n = n // 10
            
            n = s
            s = 0
                    

        return n == 1
