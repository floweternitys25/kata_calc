package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var err1 = "Ошибка, формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
var err2 = "Ошибка, используются одновременно разные системы счисления или числа вне диапазона от 1 до 10."
var err3 = "Ошибка, доступны следующие операции: a + b, a - b, a * b, a / b"
var err4 = "Ошибка, строка не является математической операцией."
var err5 = "Ошибка, в римской системе нет ноля и отрицательных чисел."
var romanNums = [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV",
	"XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX",
	"XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII",
	"XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX",
	"LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII",
	"LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV",
	"LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XVII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func main() {
	var num1, operator, num2, operator2, num3 string
	fmt.Scanln(&num1, &operator, &num2, &operator2, &num3)
	if operator2 != "" || num3 != "" {
		fmt.Println(err1)
	} else {
		if operator != "" {
			operators := [4]string{"+", "-", "/", "*"}
			for _, v := range operators {
				if operator == v {
					if isGreek(num1) && isGreek(num2) {
						num1, _ := strconv.Atoi(num1)
						num2, _ := strconv.Atoi(num2)
						calc(num1, num2, operator, false)
					} else if isRoman(num1) && isRoman(num2) {
						calc(roman2greek(num1), roman2greek(num2), operator, true)
					} else {
						fmt.Println(err2)
					}
					os.Exit(0)
				}
			}
			fmt.Println(err3)
		} else {
			fmt.Println(err4)
		}
	}
}

func calc(num1 int, num2 int, operator string, roman bool) {
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	}
	result = int(math.Floor(float64(result)))
	if roman {
		if result > 0 {
			fmt.Print(romanNums[result-1])
		} else {
			fmt.Println(err5)
		}
	} else {
		fmt.Print(result)
	}
}

func isGreek(num string) bool {
	var n, _ = strconv.Atoi(num)
	if n > 0 && n < 11 {
		return true
	} else {
		return false
	}
}

func isRoman(num string) bool {
	for i, v := range romanNums {
		if num == v && i < 11 {
			return true
		}
	}
	return false
}

func roman2greek(num string) int {
	for i, v := range romanNums {
		if num == v {
			return i + 1
		}
	}
	return 0
}
