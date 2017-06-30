package sc

import (
	"math"
)

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

func calculate_divisor(wordMap map[string]int, synonymMap map[string]int) (float64) {
	return math.Sqrt(float64(squared_sum(wordMap) * squared_sum(synonymMap)))
}


func CosineSimilarity(word string, synonym string, lexicon map[string]map[string]int) (float64) {
	_, ok := lexicon[synonym]
	if (!ok) {
		return 0
	}

	var wordMap map[string]int = lexicon[word]
	var synonymMap map[string]int = lexicon[synonym]

	var operand float64 = calculate_operand(wordMap, synonymMap)
	var divisor float64 = calculate_divisor(wordMap, synonymMap)

	return (operand / divisor)
}
