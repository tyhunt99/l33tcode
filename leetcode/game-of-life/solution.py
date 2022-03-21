from ctypes.wintypes import PINT
from typing import List

def printBoard(b):
    for row in b:
        print(row)

class Solution:
    def getLiveNeighbors(self, board, m, n):
        possible_neighbors = [
            (1,0),
            (1,-1),
            (0,-1),
            (-1,-1),
            (-1,0),
            (1,1),
            (0,1),
            (-1,1),
        ]
        live = 0
        for pn in possible_neighbors:
            x = m+pn[0]
            y = n+pn[1]
            if x >= 0 and y >= 0 and x < len(board) and y < len(board[0]):
                if board[x][y]:
                    live += 1

        return live

    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        # grab snapshot of current state for calcs
        prev = [[v for v in r] for r in board]
        #now loop through
        for m, row in enumerate(board):
            for n, _ in enumerate(row):
                live_neighbors = self.getLiveNeighbors(prev,m,n)
                # dead_neighbors = sum([l for l in neighbors if not l])

                # rebirth
                if not board[m][n] and live_neighbors == 3:
                    board[m][n] = 1
                else:
                    # under pop
                    if live_neighbors < 2:
                        board[m][n] = 0
                    # live on
                    elif live_neighbors >= 2 and live_neighbors <=3:
                        pass
                    # over population
                    elif live_neighbors > 3:
                        board[m][n] = 0
