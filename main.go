package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "Sample Case"
	vow, con := sortChar(text)
	fmt.Println("Vowels Characters : ", vow)
	fmt.Println("Consonant Characters : ", con)

	text = "Next Case"
	vow, con = sortChar(text)
	fmt.Println("Vowels Characters : ", vow)
	fmt.Println("Consonant Characters : ", con)

}

func sortChar(text string) (string, string) {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")

	type data struct {
		Key   string
		Sort  int
		Value int
	}

	vowelsMap := make(map[string]data)
	consonantMap := make(map[string]data)

	sortVow := 1
	sortCon := 1
	for _, val := range text {
		switch val {
		case 'a':
			if val, ok := vowelsMap["a"]; !ok {
				vowelsMap["a"] = data{
					Sort:  sortVow,
					Value: 1,
				}
			} else {
				vowelsMap["a"] = data{
					Sort:  sortVow,
					Value: val.Value + 1,
				}
			}

			sortVow++
		case 'e':
			if val, ok := vowelsMap["e"]; !ok {
				vowelsMap["e"] = data{
					Sort:  sortVow,
					Value: 1,
				}
			} else {
				vowelsMap["e"] = data{
					Sort:  sortVow,
					Value: val.Value + 1,
				}
			}
		case 'i':
			if val, ok := vowelsMap["i"]; !ok {
				vowelsMap["i"] = data{
					Sort:  sortVow,
					Value: 1,
				}
			} else {
				vowelsMap["i"] = data{
					Sort:  sortVow,
					Value: val.Value + 1,
				}
			}
		case 'o':
			if val, ok := vowelsMap["o"]; !ok {
				vowelsMap["o"] = data{
					Sort:  sortVow,
					Value: 1,
				}
			} else {
				vowelsMap["o"] = data{
					Sort:  sortVow,
					Value: val.Value + 1,
				}
			}
		case 'u':
			if val, ok := vowelsMap["u"]; !ok {
				vowelsMap["u"] = data{
					Sort:  sortVow,
					Value: 1,
				}
			} else {
				vowelsMap["u"] = data{
					Sort:  sortVow,
					Value: val.Value + 1,
				}
			}

		default:
			v := fmt.Sprint(string(val))
			if val, ok := consonantMap[v]; !ok {
				consonantMap[v] = data{
					Sort:  int(sortCon),
					Value: 1,
				}
				sortCon++
			} else {
				consonantMap[v] = data{
					Sort:  val.Sort,
					Value: val.Value + 1,
				}
			}
		}
	}

	var sortedConsonant []data

	for k, v := range consonantMap {
		sortedConsonant = append(sortedConsonant, data{
			Key:   k,
			Value: v.Value,
			Sort:  v.Sort,
		})
	}

	sort.Slice(sortedConsonant, func(i, j int) bool {
		return sortedConsonant[i].Sort < sortedConsonant[j].Sort
	})

	// fmt.Println(vowelsMap)
	// fmt.Println(sortedConsonant)

	var vowels string
	var consonant string
	for key, val := range vowelsMap {
		for i := 0; i < val.Value; i++ {
			vowels = fmt.Sprintf(vowels + key)
		}
	}

	for _, val := range sortedConsonant {
		for i := 0; i < val.Value; i++ {
			consonant = fmt.Sprintf(consonant + val.Key)
		}
	}

	return vowels, consonant
}
