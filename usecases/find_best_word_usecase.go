package usecases

import (
	"scrabble/domain"
)

func FindBestWordUsecase(rawDictionnary []string, lettersAvailable string) string {
	dictionnary := domain.NewDictionnary(rawDictionnary)

	bestWord := dictionnary.FindBestWordUsing(lettersAvailable)

	return bestWord.GetValue()
}
