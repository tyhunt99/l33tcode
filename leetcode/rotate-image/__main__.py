from ctypes.wintypes import tagMSG
from typing import List


def rotate(matrix: List[List[int]]) -> None:
    # reflect(matrix)
    transpose(matrix)
    reverse(matrix)

def transpose(matrix):
    n = len(matrix)
    for i in range(n):
        for j in range(i, n):
            matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]

def reverse(matrix):
    for i in range(len(matrix)):
        matrix[i].reverse()

def reflect(matrix):
    n = len(matrix)
    for i in range(n):
        for j in range(n // 2):
            matrix[i][j], matrix[i][-j - 1] = matrix[i][-j - 1], matrix[i][j]


# Function to print the matrix
def displayMatrix(mat):
    for i in range(0, len(mat)):
        for j in range(0, len(mat)):
            print(mat[i][j], end=' ')
        print()


if __name__ == "__main__":
    tests = [
        {
            "matrix": [[1,2,3],[4,5,6],[7,8,9]],
            "target": 0
        },
        {
            "matrix": [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]],
            "target": 3,
        },
    ]

    for t in tests:
        displayMatrix(t["matrix"])
        rotate(t["matrix"])
        print("***************************")
        displayMatrix(t["matrix"])
        print("==============================")
