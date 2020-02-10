package unit_4

/*
Write a function to count the frequency of words in a string of text and return a map
of words with their counts. The function should convert the text to lowercase, and punctuation should be trimmed from words.
The strings package contains several help- ful functions for this task, including Fields, ToLower, and Trim.
Use your function to count the frequency of words in the following passage and then display the count for any word that occurs more than once.

*/

import (
	"fmt"
	"strings"
)

func Freq(text string) {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,"-`)
		frequency[word]++
	}
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}

func FrequencyOfWords() {
	text := `As far as eye could reach he saw nothing but the stems of the
	great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of 
	twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first 
	thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; 
	nothing to fear -- except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles
	beyond the reach or knowledge of man.`
	Freq(text)
}
