class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # make lists for checking
        rowCheck = [[0 for _ in range(9)] for _ in range(9)]
        colCheck = [[0 for _ in range(9)] for _ in range(9)]
        boxCheck = [[0 for _ in range(9)] for _ in range(9)]

        n = len(board)

        for i in range(n) :
            for j in range(n) :
                chr = board[i][j]
                num = 0
                if chr == '.' :
                    continue
                else :
                    num = int(chr) - 1
                
                # defind the id
                r = i
                c = j
                b = (i//3)*3 + (j//3)

                if rowCheck[r][num] > 0 or colCheck[c][num] > 0 or boxCheck[b][num] > 0 :
                    return False

                rowCheck[r][num] += 1
                colCheck[c][num] += 1
                boxCheck[b][num] += 1

        return True
