package domain

import (
	"sort"
)

type Dictionnary struct {
	words []Word
}

func NewDictionnary(rawDictionnary []string) *Dictionnary {
	words := make([]Word, len(rawDictionnary))
	for i, rawWord := range rawDictionnary {
		words[i] = *newWord(rawWord)
	}

	sort.SliceStable(words, func(i, j int) bool {
		return words[i].points > words[j].points
	})

	return &Dictionnary{words: words}
}

func (dictionnary *Dictionnary) FindBestWordUsing(lettersAvailable string) *Word {
	for _, word := range dictionnary.words {
		if word.isFeasibleWith(lettersAvailable) {
			return &word
		}
	}

	return nil
}
