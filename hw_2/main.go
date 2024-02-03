package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	/*
		На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет:
			h - 1 0.1
			e - 1 0.1
			l - 3 0.33
			o - 2 0.2
			w - 1 0.1
			r - 1 0.1
			d - 1 0.1

	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите предложение: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)

	res := make(map[rune]int)
	count := 0

	for _, v := range input {
		if unicode.IsLetter(v) {
			count++
			res[v]++
		}
	}

	for k, v := range res {
		fmt.Printf("%c - %d %.1f\n", k, v, float32(v)/float32(count))
	}
}
