package dc

import (
	"math"
	"strings"
)

func CreateSentences(book string) ([]string) {
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

func CreateMap(booksSentences []string, keys []string) (map[string]map[string]int) {
	
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

func CreateUniqueKeys(sanitizedLexicon string) ([]string) {
	sanitizedLexicon = strings.Replace(sanitizedLexicon, ".", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "!", " ", -1)
	sanitizedLexicon = strings.Replace(sanitizedLexicon, "?", " ", -1)

	var lexiconsWords []string = strings.Split(sanitizedLexicon, " ")

	return removeDuplicates(lexiconsWords)
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

