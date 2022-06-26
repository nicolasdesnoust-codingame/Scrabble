package domain

var LetterPoints = map[string]int{
	"e": 1, "a": 1, "i": 1, "o": 1, "n": 1, "r": 1, "t": 1, "l": 1, "s": 1, "u": 1,
	"d": 2, "g": 2,
	"b": 3, "c": 3, "m": 3, "p": 3,
	"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
	"k": 5,
	"j": 8, "x": 8,
	"q": 10, "z": 10,
}

type Word struct {
	value  string
	points int
}

func newWord(rawWord string) *Word {
	return &Word{value: rawWord, points: calculatePoints(rawWord)}
}

func calculatePoints(rawWord string) int {
	points := 0

	for _, letter := range rawWord {
		points += LetterPoints[string(letter)]
	}

	return points
}

func (word *Word) isFeasibleWith(lettersAvailable string) bool {
	letterOccurences := countLetterOccurences(lettersAvailable)

	for _, letter := range word.value {
		if letterOccurences[string(letter)] == 0 {
			return false
		}

		letterOccurences[string(letter)] -= 1
	}

	return true
}

func countLetterOccurences(letters string) map[string]int {
	letterOccurences := make(map[string]int)

	for _, letter := range letters {
		letterOccurences[string(letter)] += 1
	}

	return letterOccurences
}

func (word *Word) GetValue() string {
	return word.value
}
