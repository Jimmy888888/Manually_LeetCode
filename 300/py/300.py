# use dp
# dp[i] = dp[j] + 1 , if nums[i] > nums[j]

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        # check 
        n = len(nums)
        if n == 1:
            return 1
        
        # make dp
        dp = [1 for _ in range(n)]

        # state transation
        for i in range(n):
            for j in range(i):
                if nums[i] > nums[j] :
                    if dp[i] < dp[j] + 1 :
                        dp[i] = dp[j] + 1
        
        # find the longest
        ans = 1
        for n in dp :
            if ans < n :
                ans = n

        return ans
