package main

import (
	"sort"
	"strings"
)

// func main() {
// 	dict := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

// 	anagrams := FindAnagrams(&dict)
// 	fmt.Println(anagrams)
// }

// FindAnagrams ...
func FindAnagrams(dict *[]string) *map[string]*[]string {
	result := make(map[string]*[]string)
	checkmap := make(map[string]*[]string)
	tmpSlice := []rune{}
	tmpString := ""

	for _, v := range *dict {
		if len(v) < 2 {
			continue
		}

		tmpString = strings.ToLower(v)
		// fmt.Println(tmpString)
		tmpSlice = []rune(tmpString)
		sort.Slice(tmpSlice, func(i, j int) bool {
			return tmpSlice[i] <= tmpSlice[j]
		})
		if mapValue, ok := checkmap[string(tmpSlice)]; ok {
			*mapValue = append(*mapValue, strings.ToLower(v))
			sort.Strings(*mapValue)
			checkmap[string(tmpSlice)] =  mapValue
		} else {
			newMapValue := &[]string{strings.ToLower(v)}
			checkmap[string(tmpSlice)] = newMapValue
		}
	}

	for _, v := range checkmap {
		if *v != nil && len(*v) > 1 {
			cutSlice := (*v)
			cutSlice = cutSlice[1:]
			result[(*v)[0]] = &cutSlice
		}
	 }

	return &result
}