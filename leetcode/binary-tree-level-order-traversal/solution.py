
from collections import deque
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        level = 0
        levels = {}
        cur_level = deque([root])
        next_level = deque([])
        while cur_level:
            node = cur_level.popleft()
            if levels.get(level, False) is False:
                levels[level] = []
            if node:
                levels[level].append(node.val)
                next_level.append(node.left)
                next_level.append(node.right)

            # last node in current level
            # rotate to next level
            if not cur_level:
                cur_level = next_level
                next_level = deque([])
                level += 1

        return [l for l in levels.values() if l]
