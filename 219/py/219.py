class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        check = dict()

        for i in range(0, n):
            num = nums[i]
            
            if num in check:
                preId = check[num]
                if i - preId <= k :
                    return True
            
            check[num] = i

        return False