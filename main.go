package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Attention()
	for {
		text := ReadNFormatLine()
		if text[0] == "exit" || text[0] == "выход" {
			break
		}
		CalcNPrint(text)
	}
}

func ReadNFormatLine() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	mtext := strings.Fields(text)
	return mtext
}

var RomeNumbers = [20]string{
	"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX"}

func CalcNPrint(mtext []string) {
	if len(mtext) == 3 {
		a := false
		b := false
		num := [2]int{0, 0}
		for i, i2 := range RomeNumbers {
			if strings.Contains(mtext[0], i2) && len(i2) == len(mtext[0]) {
				a = true
				num[0] = i + 1
			}
			if strings.Contains(mtext[2], i2) && len(i2) == len(mtext[2]) {
				b = true
				num[1] = i + 1
			}
		}
		if a == b {
			if a == false {
				var check error
				num[0], check = strconv.Atoi(mtext[0])
				num[1], check = strconv.Atoi(mtext[2])
				if check != nil {
					fmt.Println("Ошибка!")
				}
			}
			answer := 0
			switch mtext[1] {
			case "+":
				answer = Add(num[0], num[1])
			case "-":
				answer = Sub(num[0], num[1])
			case "/":
				answer = Div(num[0], num[1])
			case "*":
				answer = Mult(num[0], num[1])
			}
			if a == true {
				if answer < 1 {
					fmt.Println("Ошибка!")
				} else {
					number := []struct {
						arabic int
						rome   string
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
					var AnswerString strings.Builder
					for _, num := range number {
						for answer >= num.arabic {
							AnswerString.WriteString(num.rome)
							answer -= num.arabic
						}
					}
					fmt.Println("Ответ:" + AnswerString.String())
				}
			} else {
				fmt.Println("Ответ:" + strconv.Itoa(answer))
			}
		} else {

		}
	} else {
		fmt.Println("Ошибка!")
	}
}

// Базовые операций

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

// Правила-предупреждения

func Attention() {
	fmt.Print(
		"Правила использования калькулятора:\n" +
			"1. Калькулятор работает с простыми операциями(сложение/вычетание, умножение/деление;\n" +
			"2. Калькулятор одновременно работает лишь с двумя числа и одним оператором. Пример: 1+2 - верно, 1+2+3 - неверно;\n" +
			"3. Калькулятор работает и с арабскими, и с римскими цифрами, но не в одной операций. Пример: 1 + 2 - верно, 1 + I - неверно;\n" +
			"4. Числа должны быть в диапазоне от 1 до 10 включительно;\n" +
			"5. Числа писать слитно, а между числами и операндом ставить пробел. Пример: 1 + 2 - верно, 1+2 - неверно" +
			"6. Калькулятор оперирует лишь с целыми числами." +
			"Как при попытке введения чисел с десятичной частью, так и при получение результа с десятичной частью, десятичная часть будет отброшена;\n" +
			"7. При получение отрицательного результата в операций с римскими числами произойдет ошибка;\n\n" +
			"Приятного пользования!\n\n")
}
