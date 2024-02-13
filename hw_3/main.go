package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type item struct {
	name string
	date time.Time
	tags []string
	link string
}

func main() {

	fmt.Println("Программа для добавления url в список")
	var urls []item

	for {
		showMenu()
		var char string
		fmt.Scanln(&char)

		switch char {
		case "1":
			addUrl(&urls)
		case "2":
			showUrl(urls)
		case "3":
			delUrl(&urls)
		case "4":
			return
		}
	}
}

func showMenu() {
	fmt.Printf("1 - Добавить новый URL\n2 - Показать список всех URL\n3 - Удалить URL\n4 - Выход\nВведите пункт меню: ")
}

func addUrl(list *[]item) {
	fmt.Println("Введите новую запись в формате url описание теги")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	args := strings.Fields(text)
	if len(args) < 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	*list = append(*list, item{link: args[0], name: args[1], tags: args[2:], date: time.Now()})

}

func showUrl(list []item) {
	fmt.Printf("Всего записей %d\n", len(list))
	for i, url := range list {
		fmt.Printf("%d.\nИмя: %s\nURL: %s\nТеги: %s\nДата: %s\n", i+1, url.name, url.link, strings.Join(url.tags, ", "), url.date.Format(time.DateTime))
	}
}

func delUrl(list *[]item) {
	fmt.Println("Введите имя ссылки, которое нужно удалить")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-2]
	index := slices.IndexFunc(*list, func(i item) bool {
		return i.name == text
	})

	if index == -1 {
		fmt.Println("Ссылка с таким именем не найдена")
		return
	}
	*list = slices.Delete(*list, index, index+1)
	fmt.Printf("Ссылка %s успешно удалена\n", text)
}
