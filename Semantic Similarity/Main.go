package main

import (
	"bufio"
	"dc"
	"gi"
	"os"
	"sc"
	"si"
	"strconv"
)

func main() {
	var lexicon string = gi.GetBookText()
	var sanitizedLexicon string = si.SanitizeLexicon(lexicon)
	var booksSentences []string = dc.CreateSentences(sanitizedLexicon)
	var booksKeys []string = dc.CreateUniqueKeys(sanitizedLexicon)
	var lexiconMap map[string]map[string]int = dc.CreateMap(booksSentences, booksKeys)
	reader := bufio.NewReader(os.Stdin)

	for {

		println("Enter main word (input \"0\" to end the program): ")
		word, _ := reader.ReadString('\n')
		word = si.SanitizeUserInput(word)
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
			var synonym4 string = ""

			synonym1, _ = reader.ReadString('\n')
			synonym1 = si.SanitizeUserInput(synonym1)

			synonym2, _ = reader.ReadString('\n')
			synonym2 = si.SanitizeUserInput(synonym2)
			
			synonym3, _ = reader.ReadString('\n')
			synonym3 = si.SanitizeUserInput(synonym3)
			
			synonym4, _ = reader.ReadString('\n')
			synonym4 = si.SanitizeUserInput(synonym4)

			println("The cosine similarity of ", word, " to the following synonyms are: ")
			println(synonym1, " = ", strconv.FormatFloat(sc.CosineSimilarity(word, synonym1, lexiconMap), 'f', 2, 64))
			println(synonym2, " = ", strconv.FormatFloat(sc.CosineSimilarity(word, synonym2, lexiconMap), 'f', 2, 64))
			println(synonym3, " = ", strconv.FormatFloat(sc.CosineSimilarity(word, synonym3, lexiconMap), 'f', 2, 64))
			println(synonym4, " = ", strconv.FormatFloat(sc.CosineSimilarity(word, synonym4, lexiconMap), 'f', 2, 64))
		} else {
			println("The word was not in the lexicon.")
		}
		println()
	}

}