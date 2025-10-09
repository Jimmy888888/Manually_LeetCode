class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        arr.sort()
        n = len(arr)
        commDiff = arr[0] - arr[1]

        for i in range(1, n-1) :
            if commDiff != (arr[i] - arr[i+1]) :
                return False

        return True