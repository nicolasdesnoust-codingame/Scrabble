package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBestWord_ShouldReturnTheWordThatGivesMorePoints(t *testing.T) {
	dictionnary := NewDictionnary([]string{"dialog", "valgoid"})

	actualBestWord := dictionnary.FindBestWordUsing("aldigov").GetValue()

	expectedBestWord := "valgoid"
	assert.Equal(t, expectedBestWord, actualBestWord)
}

func TestFindBestWord_ShouldReturnTheWordThatComesFirstInDictionnary(t *testing.T) {
	dictionnary := NewDictionnary([]string{"algoid", "dialog"})

	actualBestWord := dictionnary.FindBestWordUsing("aldigov").GetValue()

	expectedBestWord := "algoid"
	assert.Equal(t, expectedBestWord, actualBestWord)
}
