package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Vowels Characters : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println(text)
	vow, con := sortChar(text)
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
	var vowels string
	var consonant string
	var sortedConsonant []data
	var sortedVowels []data

	order := 1

	var vowelChars = map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	for _, val := range text {
		v := fmt.Sprint(string(val))
		ok := vowelChars[v]

		if ok {
			if val, ok := vowelsMap[v]; !ok {
				vowelsMap[v] = data{
					Sort:  int(order),
					Value: 1,
				}
				order++
			} else {
				vowelsMap[v] = data{
					Sort:  val.Sort,
					Value: val.Value + 1,
				}
			}
		} else {
			if val, ok := consonantMap[v]; !ok {
				consonantMap[v] = data{
					Sort:  int(order),
					Value: 1,
				}
				order++
			} else {
				consonantMap[v] = data{
					Sort:  val.Sort,
					Value: val.Value + 1,
				}
			}
		}
	}

	for k, v := range vowelsMap {
		sortedVowels = append(sortedVowels, data{
			Key:   k,
			Value: v.Value,
			Sort:  v.Sort,
		})
	}
	sort.Slice(sortedVowels, func(i, j int) bool {
		return sortedVowels[i].Sort < sortedVowels[j].Sort
	})

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

	for _, val := range sortedVowels {
		for i := 0; i < val.Value; i++ {
			vowels = fmt.Sprintf(vowels + val.Key)
		}
	}

	for _, val := range sortedConsonant {
		for i := 0; i < val.Value; i++ {
			consonant = fmt.Sprintf(consonant + val.Key)
		}
	}

	return vowels, consonant
}
