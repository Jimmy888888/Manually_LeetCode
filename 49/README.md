Group Anagrams

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [ ["bat"],["nat","tan"],["ate","eat","tea"] ]

## thought: 
input：string array

processor : map

output : array of string array

## problem:
how to define the key of map

must to process every character of the strings in the InputStr

use  [26]int to check a str is in a anagram


```go
func groupAnagrams(strs []string) [][]string {

	hashMaps := make([][26]int, 0)
	outPut := make([][]string, 0)
	
	for _, str := range strs{
	
		var curMap [26]int
		
		for _, r := range str{
			curMap[r - 'a']++
		}
	
		findMap := false
		ID := 0
		
		for hashID, theMap := range hashMaps{
			
			findMap = true
			for i := range curMap{
				if curMap[i] != theMap[i]{
					findMap = false
					break
				}
			}
			
			if findMap == true{
				ID = hashID
				break
			}
		}
	
		if findMap == true{
			outPut[ID] = append(outPut[ID], str)
		}
	
		if findMap == false{
			hashMaps = append(hashMaps, curMap)
			newGroup := make([]string, 0)
			newGroup = append(newGroup, str)
			outPut = append(outPut, newGroup)
		}
	
	}
	
	return outPut
}
```

## result: pass 
but too slow
### time complexity: 
N = MAX len(strs) , L = MAX len(strs[i])

--> O(N*(N + L))

## another solution:
should make good use of map

the "key" should be found directly

sort the string first

```go
improt "slices"

func groupAnagrams(strs []string) [][]string {

	sortMaps := make(map[string] int)
	outPut := make([][]string, 0)
	
	for _, str := range strs{

		runesStr := []rune(str)
		slices.Sort(runesStr)
		sortStr := string(runesStr)
		
		GroupId, isInMap := sortMaps[sortStr]
	
		if isInMap == true {
			outPut[GroupId] = append(outPut[GroupId], str)
		}
	
		if isInMap == false {
			newGroupId = len(sortMaps)
			sortMaps[sortStr] = newGroupId
			newGroup := make([]string, 0)
			newGroup = append(newGroup, str)
			outPut = append(outPut, newGroup)
		}
	
	}
	
	return outPut
}
```

### time complexity: 
N = MAX len(strs) , L = MAX len(strs[i])

--> O(N⋅LlogL)




