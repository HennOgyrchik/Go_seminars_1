package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPrime(n int) bool {
	// Напишите код, проверяющий является ли число простым здесь
	//
	//
	if n < 0 {
		n *= -1
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	/*//task 1
	// В данном примере мы сознатель опускаем проверку типов.
	// Если передать в качестве ввода значение типа строка или булево значение, то программа аварийно завершится
	// Пакет log предоставляет нам базовые функции логирования.
	// log.Fatal печатает на экране переданный текст ошибки и звершается
	var value int
	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal(err)
	}

	if isPrime(value) {
		fmt.Println("Число", value, "простое")
	} else {
		fmt.Println("Число", value, "не простое")
	}*/

	//task 2
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordLine, symbolNumber, bytesNumber int
	// Напишите код здесь

	arr := strings.Fields(text[:len(text)-2])

	wordLine = len(arr)

	for _, s := range arr {
		symbolNumber += len([]rune(s))
		bytesNumber += len(s)
	}

	// Подсказка. Чтобы обойти все символы в строке используйте цикл for _, v := range value
	// Помните вам нужно вывести количество букв, а не символов.
	// Используйте unicode.IsSpace() для определения пробела
	// Используйте unicode.IsLetter() для определения буквы
	fmt.Printf("Количество слов: %d\n", wordLine)
	fmt.Printf("Количество букв: %d\n", symbolNumber)
	fmt.Printf("Количество байт: %d\n", bytesNumber)

}
