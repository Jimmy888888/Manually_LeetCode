class Solution:

    def canConstruct(self, ransomNote: str, magazine: str) -> bool:

        for chr in ransomNote :
            if chr in magazine :
                magazine = magazine.replace(chr, '', 1)
            else :
                return False

        return True

