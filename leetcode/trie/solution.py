class TrieNode():

    def __init__(self, is_end=False) -> None:
        self.letter_nodes = {}
        self.is_end = False

    def get_letter(self, letter):
        return self.letter_nodes.get(letter)

    def add_letter(self, letter):
        self.letter_nodes[letter] = TrieNode(letter)
        return self.get_letter(letter)


class Trie:

    root = None

    def __init__(self) -> None:
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for c in word:
            letter_node = node.get_letter(c)
            if not letter_node:
                letter_node = node.add_letter(c)
            # point to next node for next iteration
            node = letter_node

        # last letter so flag as word end
        node.is_end =  True

    def search(self, word: str) -> bool:
        node = self.root
        for c in word:
            letter_node = node.get_letter(c)
            if not letter_node:
                return False
            node = letter_node

        # if consdered word end
        return node.is_end


    def startsWith(self, prefix: str) -> bool:
        node = self.root
        for c in prefix:
            letter_node = node.get_letter(c)
            if not letter_node:
                return False
            node = letter_node

        # if consdered word end
        return True
