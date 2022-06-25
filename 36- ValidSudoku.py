# straightforward
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row = {}
        col = {}
        box1 = {}
        box2 = {}
        box3 = {}
        for i in range(0, 9):
            for j in range(0, 9):
                if board[i][j] != '.' and board[i][j] not in row:
                    row[board[i][j]] = True
                elif board[i][j] != '.' and board[i][j] in row:
                    return False

                if board[j][i] != '.' and board[j][i] not in col:
                    col[board[j][i]] = True
                elif board[j][i] != '.' and board[j][i] in col:
                    return False

                if j < 3:
                    if board[i][j] != '.' and board[i][j] not in box1:
                        box1[board[i][j]] = True
                    elif board[i][j] != '.' and board[i][j] in box1:
                        return False
                elif j < 6:
                    if board[i][j] != '.' and board[i][j] not in box2:
                        box2[board[i][j]] = True
                    elif board[i][j] != '.' and board[i][j] in box2:
                        return False
                elif j < 9:
                    if board[i][j] != '.' and board[i][j] not in box3:
                        box3[board[i][j]] = True
                    elif board[i][j] != '.' and board[i][j] in box3:
                        return False

            if i == 2 or i == 5 or i == 8:
                box1 = {}
                box2 = {}
                box3 = {}
            row = {}
            col = {}
        return True
