from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        reachable = 0
        for i, n in enumerate(nums):
            if i > reachable:
                return False
            if i + n > reachable:
                reachable = i + n
        return True
