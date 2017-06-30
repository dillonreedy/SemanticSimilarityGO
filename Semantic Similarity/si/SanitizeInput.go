package si

import (
	"strings"
)

func SanitizeUserInput(givenWord string) (string) {
	givenWord = strings.TrimSpace(givenWord)
	givenWord = strings.ToLower(givenWord)
	
	return givenWord
}

func SanitizeLexicon(lexicon string) (string) {

	lexicon = removeInnerPunctuation(lexicon)
	lexicon = strings.ToLower(lexicon)

	return lexicon
}

func removeInnerPunctuation(book string) (string) {
	book = strings.Replace(book, ",", " ", -1)
	book = strings.Replace(book, "-", " ", -1)
	book = strings.Replace(book, "--", " ", -1)
	book = strings.Replace(book, ":", " ", -1)
	book = strings.Replace(book, ";", " ", -1)
	book = strings.Replace(book, "\"", " ", -1)
	book = strings.Replace(book, "'", "", -1)
	book = strings.Replace(book, "/", " ", -1)
	book = strings.Replace(book, ")", " ", -1)
	book = strings.Replace(book, "(", " ", -1)
	book = strings.Replace(book, "\n", " ", -1)
	book = strings.Replace(book, "\t", " ", -1)
	return book
}