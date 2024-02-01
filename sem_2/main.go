package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func main() {
	task1()
	fmt.Println()
	task2()
}

func task1() {
	var arr [100]int
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(20)
		arr[i] = value
	}

	// Создайте слайс нужного размера
	slice := make([]int, 20)

	for _, v := range arr {
		slice[v]++
	}
	// Заполните слайс количеством встречаемых чисел

	fmt.Println(arr, "\n", slice)
}

// Напишите функцию фильтрации слайса. Отфильтруйте слайс arr1 так, чтобы он содержал только не четные числа
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] после фильтрации [3, 1, 5]
// Не используйте готовые функции из пакета slices
func filter(arr1 []int) []int {
	res := make([]int, 0)
	for _, v := range arr1 {
		if v%2 == 1 {
			res = append(res, v)
		}
	}
	return res

	///// вариант 2
	/*var n int
	for _, v := range arr1 {
		if v %2 != 0{
			arr1[n]=v
			n++
		}
	}
	return arr1[:n]*/

}

// Напишите функцию вставки элемента в слайс на позицию
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] pos = 1, value = 4. Результат [0, 4, 2, 3, 1, 5, 4]
// Не используйте готовые функции из пакета slices
func insert(pos, value int, arr1 []int) []int {
	res := make([]int, 0, cap(arr1)+1)

	res = append(res, arr1[:pos]...)
	res = append(res, value)
	res = append(res, arr1[pos:]...)

	return res

	//return append(arr1[:pos], append([]int{value}, arr1[pos:]...)...)
}

func task2() {
	// Этот код нужен для ввода массива из стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите элементы массива через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем строку на массив строк
	strValues := strings.Split(input, " ")

	// Преобразуем строки в числа и заполняем массив
	var arr []int
	for _, str := range strValues {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		arr = append(arr, val)
	}

	var pos int
	fmt.Println("Введите позицию для вставки")
	if _, err := fmt.Fscan(os.Stdin, &pos); err != nil {
		log.Fatal("Неправильное значение")
	}
	if pos < 0 || pos > len(arr)-1 {
		log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	}

	var value int
	fmt.Println("Введите значение для вставки")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal("Неправильное значение")
	}
	if pos < 0 || pos > len(arr)-1 {
		log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	}

	// Скопируйте слайс arr в слайс arr1.
	// Это можно сделать через функцию Append или через функцию copy

	arr1 := make([]int, len(arr))
	copy(arr1, arr)

	arr1 = insert(pos, value, arr1)

	arr1 = filter(arr1)

	fmt.Println(arr1)
}

type mySlice struct {
	len int
	cap int
	p   unsafe.Pointer
}

func task3() {
	s := []int64{1, 2, 3} //len 3, cap 3
	slice(s)
	slice1(&s)

	m := make(map[string]struct{})
	mutMap(m)

	mm := make(map[string]map[string]int, 0)
	mm["s"] = make(map[string]int)
}

func slice(arr []int64) {
	arr[1] = 10
	arr = append(arr, 20) // len 3
}

func slice1(arr *[]int64) {
	*arr = append(*arr, 20)
}

func mutMap(s map[string]struct{}) {
	s["s"] = struct{}{}
	for k := range s {
		fmt.Println(k)
	}
}
