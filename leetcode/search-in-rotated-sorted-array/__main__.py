from ctypes.wintypes import tagMSG
from typing import List


def searchDistinct(nums: List[int], target: int) -> int:
    prev_val = None
    for i, val in enumerate(nums):
        if val == target:
            return i
        if prev_val and prev_val > val:
            # reached pivot
            return -1
    return -1


def search(nums: List[int], target: int) -> int:
    pass


if __name__ == "__main__":
    tests = [
        {
            "nums": [4,5,6,7,0,1,2],
            "target": 0
        },
        {
            "nums": [4,5,6,7,0,1,2],
            "target": 3,
        },
        {
            "nums": [8,12,15,20,2,3,6],
            "target": 3,
        },
        {
            "nums": [1],
            "target": 3,
        },
    ]

    for t in tests:
        print(searchDistinct(t["nums"], t["target"]))
        print("==============================")
