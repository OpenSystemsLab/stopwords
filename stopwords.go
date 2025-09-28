package stopwords

import "github.com/OpenSystemsLab/stopwords/data"

func IsStopWord(lang, word string) bool {
	_, ok := data.Languages[lang][word]

	return ok
}
