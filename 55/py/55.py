from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        step = nums[0]
        newStep = 0
        position = 0
        i = 0
        while i < len(nums) :
            # check success
            if i + step >= len(nums) - 1 :
                return True
            elif i + step <= i :
                return False
            
            newStep = 0
            position = i + 1
            for j in range(i+1, i+step+1) :
                if j + nums[j] >= i + step:
                    newStep = nums[j]
                    position = j
                    break
            

            i = position
            step = newStep

        return True
            
