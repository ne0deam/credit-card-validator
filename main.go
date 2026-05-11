package main

import (
	"fmt";
	"strings"
)

type Bank struct {
	Name string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	for i:= 0; i < len(banks); i++ {
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

func main() {
    banks := []Bank{
	    {Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"}, 
		{Name: "Venus Express", Prefix: "6000"},
	}

    is_valid := LuhnCheck("4000123456789017")
    
	fmt.Println("Валиден по Луне:", is_valid)

	if !is_valid {
		fmt.Println("Банк: не определен")
	}else{
		fmt.Println("Банк:", DetectBank("4000123456789017", banks).Name)
	}

}