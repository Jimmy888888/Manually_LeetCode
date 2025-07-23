package anagrams_test

import (
	"reflect"
	"sort"
	"testing"

	"example.com/anagrams"
)

// Helper function to sort slices of strings for consistent comparison
func sortSlices(slices [][]string) {
	for _, s := range slices {
		sort.Strings(s)
	}
	sort.Slice(slices, func(i, j int) bool {
		if len(slices[i]) == 0 || len(slices[j]) == 0 {
			return len(slices[i]) < len(slices[j])
		}
		return slices[i][0] < slices[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "Example 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name:  "Empty string",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "Single character strings",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
        {
			name:  "Empty input",
			input: []string{},
			want:  [][]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := anagrams.GroupAnagrams(tc.input)

			sortSlices(got)
			sortSlices(tc.want)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tc.want)
			}
		})
	}
}
