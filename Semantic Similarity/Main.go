package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var text string = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}
	return text, scanner.Err()
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

func splitOutSentences(book string) ([]string) {
	var pSentences []string = strings.Split(book, ".")
	var pqSentences []string
	for v := range pSentences {
		var pqChunkSentences []string = strings.Split(pSentences[v], "?")
		
		for q := range pqChunkSentences {
			pqSentences = append(pqSentences, pqChunkSentences[q])
		}
	}

	var pqeSentences []string
	for v := range pqSentences {
		var pqeChunkSentences []string = strings.Split(pqSentences[v], "!")
		
		for q := range pqeChunkSentences {
			pqeSentences = append(pqeSentences, pqeChunkSentences[q])
		}
	}

	return pqeSentences
}

func sanitizeLexicon(lexicon string) (string) {

	lexicon = removeInnerPunctuation(lexicon)
	lexicon = strings.ToLower(lexicon)

	return lexicon
}

func createMap(booksSentences []string, keys []string) (map[string]map[string]int) {
	
	var result map[string]map[string]int = make(map[string]map[string]int)

	println(len(keys))
	for k := range keys {
		//if (keys[k] != nil) {
		result[keys[k]] = make(map[string]int)
		//}
	}
	println("All initialized")

	for bs := range booksSentences {
		println(booksSentences[bs])
		var sentenceWords []string = strings.Split(booksSentences[bs], " ")

		var count int = len(sentenceWords)
		var halfCount int = int(math.Ceil(float64(count) / float64(2)))

		for i := 0; i < halfCount; i++ {
			var aWord string = sentenceWords[i]

			for j := i+1; j < count; j++ {
				var bWord string = sentenceWords[j]

				result[aWord][bWord] += 1
				result[bWord][aWord] += 1
			}
		}  	 
	}

	return result
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

func getUniqueKeys(sanitizedLexicon string) ([]string) {
	sanitizedLexicon = strings.Replace(sanitizedLexicon, ".", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "!", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "?", " ", -1)

	var lexiconsWords []string = strings.Split(sanitizedLexicon, " ")

	return removeDuplicates(lexiconsWords)
}


func main() {
	var lexicon string

	for i := 0; i < 7; i++ {

		var bookName string = "HarryPotter" + strconv.Itoa(i+1) + ".txt"
		bookText, someErr := readLines(bookName)
		if someErr != nil {
			os.Exit(3)
		}
		lexicon += bookText
	}

	var sanitizedLexicon string = sanitizeLexicon(lexicon)
	var booksSentences []string = splitOutSentences(sanitizedLexicon)
	var booksKeys []string = getUniqueKeys(sanitizedLexicon)

	var wordMap map[string]map[string]int = createMap(booksSentences, booksKeys)
	
	for k := range booksKeys {
		println(booksKeys[k])
		wordMap[booksKeys[k]]
	}
}