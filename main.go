package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var romanMap = []struct {
	decVal int
	symbol string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	var formula string
	var roman_numerals = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4,
		"V": 5, "VI": 6, "VII": 7, "VIII": 8,
		"IX": 9, "X": 10,
	}

	for true {
		fmt.Println("Для выхода из калькулятора введите exit")
		fmt.Print("Введите уравнение -> ")
		formula = Input()

		if formula == "exit" {
			fmt.Println("Введена команда выхода")
			break
		}

		lst := strings.Split(formula, " ")
		keys := make([]string, 0, len(roman_numerals))

		for k := range roman_numerals {
			keys = append(keys, k)
		}

		if len(lst) == 3 {
			ind0 := slices.IndexFunc(keys, func(elt string) bool { return elt == strings.Split(lst[0], "")[0] })
			ind1 := slices.IndexFunc(keys, func(elt string) bool { return elt == strings.Split(lst[2], "")[0] })

			if ind0 == -1 && ind1 == -1 {
				a, err := strconv.Atoi(lst[0])
				symbol := lst[1]
				b, err := strconv.Atoi(lst[2])

				if err != nil {
					panic("Выдача паники, так как не подходящие число")
				}
				o
				result := calculator(a, b, symbol)
				fmt.Println(result)

			} else if ind0 > -1 && ind1 > -1 {
				a := roman_numerals[lst[0]]
				symbol := lst[1]
				b := roman_numerals[lst[2]]
				if a == 0 || b == 0 {
					if a == 0 {
						panic("Выдача паники, потому что неправильное значение: " + lst[0])
					} else {
						panic("Выдача паники, потому что неправильное значение: " + lst[2])
					}
				}
				answer := calculator(a, b, symbol)
				if answer <= 0 {
					panic("Выдача паники, так как в римской системе нет нуля и отрицательных чисел")
				}

				result := decimalToRoman(answer)
				fmt.Println(result)

			} else {
				panic("Выдача паники, так как используются одновременно разные системы счисления.")
			}

		} else {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию")
		}
	}
}

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func calculator(a int, b int, symbol string) int {
	var result int
	if symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/" {
		if a > 0 && a <= 10 && b > 0 && b <= 10 {
			switch symbol {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			return result
		} else {
			panic("Выдача паники, так как не подходящие числа")
		}
	} else {
		panic("Неизвестная математичаская операция")
	}

}

func decimalToRoman(n int) string {
	for _, v := range romanMap {
		if n >= v.decVal {
			return v.symbol + decimalToRoman(n-v.decVal)
		}
	}
	return ""
}
