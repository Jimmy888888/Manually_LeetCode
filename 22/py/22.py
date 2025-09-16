class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        result = []

        def backTracking(cur_str :str, left_num :int, right_num :int) :
            if len(cur_str) == 2*n :
                result.append(cur_str)
                return

            if left_num < n :
                backTracking(cur_str + "(", left_num + 1, right_num)

            if right_num < left_num :
                backTracking(cur_str + ")", left_num, right_num + 1)

        backTracking("", 0, 0)

        return result