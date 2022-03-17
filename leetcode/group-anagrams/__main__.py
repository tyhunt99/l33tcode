from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for w in strs:
            idx = ''.join(sorted(w))
            if idx in groups:
                groups[idx].append(w)
            else:
                groups[idx] = [w]

        return list(groups.values())


if __name__ == "__main__":
    tests = [
        {
            "input": ["eat","tea","tan","ate","nat","bat"],
        },
        {
            "input": [""],
        },
        {
            "input": ["a"]
        },
    ]


    solu = Solution()
    for t in tests:
        print(solu.groupAnagrams(t["input"]))
        print("==============================")
