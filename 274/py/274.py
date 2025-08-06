class Solution:
    def hIndex(self, citations: List[int]) -> int:
        citations.sort(reverse=True)

        hId = 0
        i = 0
        for cited in citations :
            if i+1 > cited :
                hId = i
                return hId
            i = i+1

        hId = len(citations)

        return hId
