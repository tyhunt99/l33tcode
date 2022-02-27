from platform import node
from typing import List, Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def generateHead(values: List) -> ListNode:
    lastNode = head = ListNode(val=values[0], next=None)
    for v in values[1:]:
        node = ListNode(val=v, next=None)
        lastNode.next = node
        lastNode = node
    return head


def printListNode(head):
    nodes = []
    node = head
    while node:
        nodes.append(node.val)
        node = node.next
    return nodes


def removeNthFromEnd(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    node = head
    nodes = []
    while node:
        nodes.append(node)
        node = node.next
    length = len(nodes)
    nodeToRemoveIdx = length-n
    if nodeToRemoveIdx < 0:
        # index out of bounds, return original head
        return head
    elif length == 1:
        # if removing the only element from the list
        head = None
    elif n == 1:
        # if removing the last node
        # avoids infinite loop
        nodes[nodeToRemoveIdx-1].next = None
    elif length-n == 0:
        # if removing the first node
        head = nodes[1]
    else:
        # default case of removing node
        nodes[nodeToRemoveIdx-1].next = nodes[nodeToRemoveIdx+1]
    return head


if __name__ == "__main__":
    tests = [
        {
            "head": generateHead([1,2,3,4,5]),
            "n": 2,
        },
        {
            "head": generateHead([1]),
            "n": 1,
        },
        {
            "head": generateHead([1,2]),
            "n": 1,
        },
        {
            "head": generateHead([1,2]),
            "n": 2,
        },
        {
            "head": generateHead([1,2,3]),
            "n": 1,
        },
        {
            "head": generateHead([1,2,3,4,5,6,7,8,9,10]),
            "n": 7,
        }
    ]

    for t in tests:
        print("++++++++++++++++++")
        result = removeNthFromEnd(t["head"], t["n"])
        print(printListNode(result))
        print("++++++++++++++++++")
