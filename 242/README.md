# 242

## way 1

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        # check
        if len(s) != len(t) :
            return False

        # make dict
        dictT = {}
        for chT in t :
            val = dictT.get(chT)
            if val :
                dictT[chT] = val + 1
            else :
                dictT[chT] = 1
        
        # check anagram
        for chS in s :
            val = dictT.get(chS)
            if val :
                if val - 1 < 0 :
                    return False
                else :
                    dictT[chS] = val-1
            else :
                return False

        return True
```

## way 2

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        for char in string.ascii_lowercase:
            if s.count(char) != t.count(char):
                return False
        return True
```