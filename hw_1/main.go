package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var filePth string
	fmt.Print("Укажите полный путь до файла: ")
	if _, err := fmt.Fscan(os.Stdin, &filePth); err != nil {
		log.Fatal(err)
	}

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex

	if strings.Contains(filePth, ".") {
		fileName = filePth[strings.LastIndex(filePth, "/")+1 : strings.LastIndex(filePth, ".")]
		fileExt = filePth[strings.LastIndex(filePth, ".")+1:]
	} else {
		fileName = filePth[strings.LastIndex(filePth, "/")+1:]
		fileExt = ""
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
