class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1 :
            return nums[0]

        curS = nums[0]
        golS = nums[0]

        for i in range(1, n):
            num = nums[i]
            if num > num + curS:
                curS = num
            else:
                curS += num
            
            if golS < curS :
                golS = curS
        
        return golS