package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при создание файла для логов", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println("Начинаю работу...")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Не удалось найти директорию")
		return
	}

	desktopPath := filepath.Join(homeDir, "Desktop", "hello.txt")

	file, err := os.Create(desktopPath)
	if err != nil {
		log.Println("Ошибка при создание файла")
		return
	}
	err = file.Close()
	if err != nil {
		log.Println("Ошибка при закрытие файла: ", err)
	}

	fmt.Print("Введите текст: ")
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		log.Println("Ошика при чтении ввода: ", err)
		return
	}

	var input string

	if scanner.Scan() {
		input = scanner.Text()
		log.Println("Пользователь ввел: ", input)
	}

	_, err = file.WriteString(input)
	if err != nil {
		log.Println("Ошибка при написание текста")
		return
	}

	log.Printf("Файл успешно создан на рабочем столе %s\n", desktopPath)

	log.Println("Закончил работу...\n")

}
