package gi


import (
	"bufio"
	"os"
	"strconv"
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

func GetBookText() (string) {
	var booksText string

	for i := 0; i < 7; i++ {

		var bookName string = "HarryPotter" + strconv.Itoa(i+1) + ".txt"
		bookText, someErr := readLines(bookName)
		if someErr != nil {
			os.Exit(3)
		}
		booksText += bookText
	}

	return booksText
}