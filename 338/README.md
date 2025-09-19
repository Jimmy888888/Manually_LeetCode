# 338. Counting Bits

## Brute Force

convert to binary and count the number of 1s

```python
class Solution:
    def countBits(self, n: int) -> List[int]:
        res = [0]

        if n == 0 :
            return res

        for i in range(1, n+1):
            num = 0
            strBin = bin(i)
            for ch in strBin:
                if ch == "1":
                    num += 1
            
            res.append(num)
        
        return res


# 6 -> 110
# 7 -> 111
# 8 -> 1000

# 2^x + 2^y + 2^z + ... + 2^0
```
