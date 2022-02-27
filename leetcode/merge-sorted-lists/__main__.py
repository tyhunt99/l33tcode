from typing import List, Optional


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


def mergeTwoLists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    sortedHead = lastNode = node = None
    # handle empty cases
    if not list1 and not list2:
        return None
    elif list1 and (not list2 or list1.val < list2.val):
        sortedHead = list1
    else:
        sortedHead = list2
    while True:
        if list1 and (not list2 or list1.val < list2.val):
            node = list1
            list1 = list1.next
        else:
            node = list2
            list2 = list2.next
        if lastNode:
            lastNode.next = node
        lastNode = node
        if list1 is None and list2 is None:
            break
    return sortedHead


if __name__ == "__main__":
    tests = [
        {
            "head1": generateLinkedList([1,2,4]),
            "head2": generateLinkedList([1,3,4]),
        },
        {
            "head1": generateLinkedList([]),
            "head2": generateLinkedList([]),
        },
        {
            "head1": generateLinkedList([]),
            "head2": generateLinkedList([1,2]),
        },
        {
            "head1": generateLinkedList([2,3]),
            "head2": generateLinkedList([]),
        }
    ]

    for t in tests:
        result = mergeTwoLists(t["head1"], t["head2"])
        print(printListNode(result))
        print("==============================")
