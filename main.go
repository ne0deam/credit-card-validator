package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	for i := 0; i < len(banks); i++ {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func LuhnCheck(cardNumber string) bool {
	sum := 0
	var even bool

	if cardNumber == "" {
		return false
	}

	for i := len(cardNumber) - 1; i >= 0; i-- {
		if cardNumber[i] < '0' || cardNumber[i] > '9' {
			return false
		}

		digit := int(cardNumber[i] - '0')

		if even {
			digit *= 2
		}

		if digit > 9 {
			digit -= 9
		}

		sum += digit
		even = !even
	}

	return sum%10 == 0
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %w", err)
	}
	defer file.Close()

	banks := []Bank{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		line = strings.TrimSpace(line)
		items := strings.Split(line, ",")
		if len(items) != 2 {
			return nil, fmt.Errorf("неверный формат строки: %q", line)
		}
		bankName := strings.TrimSpace(items[0])
		bankPrefix := strings.TrimSpace(items[1])

		banks = append(banks, Bank{bankName, bankPrefix})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Загружено банков:", len(banks))

	return banks, nil
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите номер карты (или Enter для выхода):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "-", "")

	return input
}

func main() {
	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		userInput := getUserInput()
		if userInput != "" {
			isValid := LuhnCheck(userInput)
			fmt.Println("Валиден по Луне:", isValid)

			if !isValid {
				fmt.Println("Банк: не определен")
			} else {
				fmt.Println("Банк:", DetectBank(userInput, banks).Name)
			}
		} else {
			fmt.Println("До встречи!")
			break
		}
	}
}
