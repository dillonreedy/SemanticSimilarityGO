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

	for k := range keys {
		result[keys[k]] = make(map[string]int)
	}

	for bs := range booksSentences {
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

func squared_sum(givenMap map[string]int) (int) {
	var sum float64 = 0.0
	for _, value := range givenMap {
		sum += math.Pow(float64(value), float64(2))
	}
	return int(sum)
}

func calculate_operand(wordMap map[string]int, synonymMap map[string]int) (float64) {
	var sum float64 = 0.0
	for key, value := range wordMap {
		_, ok := synonymMap[key]
		if (ok) {
			sum += (float64(value) * float64(synonymMap[key]))
		}
	}
	return sum
}


func cosine_similarity(word string, synonym string, lexicon map[string]map[string]int) (float64) {
	_, ok := lexicon[synonym]
	if (!ok) {
		return 0
	}

	var wordMap map[string]int = lexicon[word]
	var synonymMap map[string]int = lexicon[synonym]

	var operand float64 = calculate_operand(wordMap, synonymMap)
	var divisor float64 = math.Sqrt(float64(squared_sum(wordMap) * squared_sum(synonymMap))) 

	return (operand / divisor)
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

	var lexiconMap map[string]map[string]int = createMap(booksSentences, booksKeys)

	for {

		println("Enter main word (input \"0\" to end the program): ")
		reader := bufio.NewReader(os.Stdin)
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		//println(word, len(word))
		if word == "0" {
			os.Exit(0)
		}

		_, ok := lexiconMap[word]

		if (ok) {
			println("Please enter the 4 synonyms: ")
			var synonym1 string = ""
			var synonym2 string = ""
			var synonym3 string = ""
			var	synonym4 string = ""

			synonym1, _ = reader.ReadString('\n')
			synonym1 = strings.TrimSpace(synonym1)
			synonym1 = strings.ToLower(synonym1)
			
			synonym2, _ = reader.ReadString('\n')
			synonym2 = strings.TrimSpace(synonym2)
			synonym2 = strings.ToLower(synonym2)
			
			synonym3, _ = reader.ReadString('\n')
			synonym3 = strings.TrimSpace(synonym3)
			synonym3 = strings.ToLower(synonym3)
			
			synonym4, _ = reader.ReadString('\n')
			synonym4 = strings.TrimSpace(synonym4)
			synonym4 = strings.ToLower(synonym4)

			println("The cosine similarity of ", word, " to the following synonyms are: ")
			println(synonym1, " = ", strconv.FormatFloat(cosine_similarity(word, synonym1, lexiconMap), 'f', 2, 64))
			println(synonym2, " = ", strconv.FormatFloat(cosine_similarity(word, synonym2, lexiconMap), 'f', 2, 64))
			println(synonym3, " = ", strconv.FormatFloat(cosine_similarity(word, synonym3, lexiconMap), 'f', 2, 64))
			println(synonym4, " = ", strconv.FormatFloat(cosine_similarity(word, synonym4, lexiconMap), 'f', 2, 64))
		} else {
			println("The word was not in the lexicon.")
		}
	}

}