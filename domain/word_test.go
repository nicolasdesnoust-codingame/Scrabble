package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoints_ShouldSetCorrectAmountOfPoints(t *testing.T) {
	word := newWord("xylophone")

	actualPoints := word.points

	expectedPoints := 24
	assert.Equal(t, expectedPoints, actualPoints)
}

func TestIsFeasibleWith_ShouldReturnFalseIfSomeLettersAreMissing(t *testing.T) {
	word := newWord("xylophone")

	isFeasible := word.isFeasibleWith("lenoyop")

	assert.False(t, isFeasible)
}

func TestIsFeasibleWith_ShouldReturnTrueIfAllLettersAreAvailable(t *testing.T) {
	word := newWord("diamond")

	isFeasible := word.isFeasibleWith("dniadmo")

	assert.True(t, isFeasible)
}

func TestIsFeasibleWith_ShouldNotUseSameLetterTwice(t *testing.T) {
	word := newWord("diamond")

	isFeasible := word.isFeasibleWith("dniazmo")

	assert.False(t, isFeasible)
}

func TestGetValue_ShouldReturnRawValue(t *testing.T) {
	word := newWord("xylophone")

	actualValue := word.GetValue()

	expectedValue := "xylophone"
	assert.Equal(t, expectedValue, actualValue)
}
