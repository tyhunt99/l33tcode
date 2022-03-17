from typing import List, Optional
import time


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def generateLinkedList(values: List) -> ListNode:
    """
    Helper function to create a linked list from a list, returns head
    """
    if not values:
        return None
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


def mergeKLists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    if not lists:
        return None
    newListHead = newList = None
    for l in lists:
        lastNode = node = None
        # no list so nothing to do
        if not l:
            continue

        if not newListHead:
            newListHead = l

        # determine if head needs updated and swap if it does
        if newListHead and newListHead.val > l.val:
            newListHead = l
            l = newList
            newList = newListHead

        while True:
            if newList and (not l or newList.val <= l.val):
                node = newList
                newList = newList.next
            else:
                node = l
                l = l.next
            if lastNode:
                lastNode.next = node
            lastNode = node
            if newList is None and l is None:
                break

        # reset newList
        newList = newListHead

    return newListHead


if __name__ == "__main__":
    tests = [
        [
            generateLinkedList([1,4,5]),
            generateLinkedList([1,3,4]),
            generateLinkedList([2,6]),
        ],
        [
            generateLinkedList([2,4,5]),
            generateLinkedList([1,3,4]),
        ],
        [],
        [generateLinkedList([])],
        [
            generateLinkedList([]),
            generateLinkedList([1])
        ],
    ]

    for t in tests:
        result = mergeKLists(t)
        print(printListNode(result))
        print("==============================")
