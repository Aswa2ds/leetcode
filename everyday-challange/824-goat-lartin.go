package everydaychallange

import "strings"

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	result := make([]string, 0)

	var startWithVowel func(word string) bool

	startWithVowel = func(word string) bool {
		vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
		word = strings.ToLower(word)
		_, ok := vowels[word[0]]
		return ok
	}

	for i, word := range words {
		if !startWithVowel(word) {
			startle := word[:1]
			word = word[1:] + startle
		}
		word = word + "ma"
		for j := 0; j <= i; j++ {
			word = word + "a"
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}
