package anagrams

import "slices"

func GroupAnagrams(strs []string) [][]string {

	sortMaps := make(map[string]int)
	outPut := make([][]string, 0)

	for _, str := range strs {

		runesStr := []rune(str)
		slices.Sort(runesStr)
		sortStr := string(runesStr)

		GroupId, isInMap := sortMaps[sortStr]

		if isInMap == true {
			outPut[GroupId] = append(outPut[GroupId], str)
		}

		if isInMap == false {
			newGroupId := len(sortMaps)
			sortMaps[sortStr] = newGroupId
			newGroup := make([]string, 0)
			newGroup = append(newGroup, str)
			outPut = append(outPut, newGroup)
		}

	}

	return outPut
}
