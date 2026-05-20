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

func validateInput(cardNumber string) error {
	minLen := 13
	maxLen := 19
	lenInput := len(cardNumber)
	for _, char := range cardNumber {
		if char < '0' || char > '9' {
			return fmt.Errorf("Ошибка: номер должен содержать только цифры")
		}
	}
	if lenInput < minLen || lenInput > maxLen {
		return fmt.Errorf("Ошибка: номер должен содержать от %d до %d цифр", minLen, maxLen)
	}
	return nil
}

func main() {
	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		userInput := getUserInput()
		if userInput == "" {
			fmt.Println("До свидания!")
			break
		} else {
			err := validateInput(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				isValid := LuhnCheck(userInput)
				if !isValid {
					fmt.Println("Ошибка: номер не прошёл проверку Луна")
					continue
				} else {
					bank := DetectBank(userInput, banks)
					if bank != nil {
						fmt.Println("Банк:", bank.Name)
					} else {
						fmt.Println("Неизвестный банк")
					}
				}
			}
		}
	}
}
