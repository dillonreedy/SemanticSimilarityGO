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

func removeDuplicates(elements []string) ([]string) {
    // Use map to record duplicates as we find them.
    encountered := map[string]bool{}
    result := []string{}

    for v := range elements {
        if encountered[elements[v]] == true {
            // Do not add duplicate.
        } else {
            // Record this element as an encountered element.
            encountered[elements[v]] = true
            // Append to result slice.
            result = append(result, elements[v])
        }
    }
    // Return the new slice.
    return result
}

func GetUniqueKeys(sanitizedLexicon string) ([]string) {
	sanitizedLexicon = strings.Replace(sanitizedLexicon, ".", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "!", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "?", " ", -1)

	var lexiconsWords []string = strings.Split(sanitizedLexicon, " ")

	return removeDuplicates(lexiconsWords)
}