package task4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CheckSpam(inputFile, outputFile string) {
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка открытия входного файла:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка создания выходного файла:", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)

	re := regexp.MustCompile(`[^\d]`)

	for scanner.Scan() {
		originalNumber := scanner.Text()
		cleanedNumber := re.ReplaceAllString(originalNumber, "")

		status := "OK"
		if len(cleanedNumber) != 11 {
			status = "ERROR"
		} else if numberIsSpam(cleanedNumber) {
			status = "SPAM"
		}

		result := fmt.Sprintf("%s %s\n", originalNumber, status)
		_, err := outFile.WriteString(result)
		if err != nil {
			fmt.Println("Ошибка записи в выходной файл:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}
}

func numberIsSpam(number string) bool {
	cityCodeFirstDigit := number[1:4][0]
	if cityCodeFirstDigit != '8' {
		return true
	}

	phoneNumberFirstDigit := number[4]
	for i := 4; i < len(number); i++ {
		if phoneNumberFirstDigit != number[i] {
			return false
		}
	}

	return true
}
