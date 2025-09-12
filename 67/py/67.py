class Solution:
    def addBinary(self, a: str, b: str) -> str:
        anum = int(a, 2)
        bnum = int(b, 2)

        result = anum + bnum

        return bin(result)[2:]#for cuting "0b"