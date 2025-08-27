class Solution:
    def jump(self, nums: List[int]) -> int:
        if len(nums) <= 1 :
            return 0

        minStep = 0
        curFarthest = 0
        curBound = 0

        for i in range(0, len(nums)-1, 1) :
            if curFarthest < i + nums[i] :
                curFarthest = i + nums[i]

            if i == curBound :
                curBound = curFarthest
                minStep += 1

                if curBound >= len(nums) - 1 :
                    return minStep

        return minStep