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

	file, err := os.OpenFile(desktopPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Ошибка при открытие/создание файла")
		return
	}
	defer file.Close()

	fmt.Print("Программа готова к работе. Для выхода введите exit. \n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите текст: ")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			log.Println("Пользователь решил выйти.")
			break
		}

		if input == "" {
			continue
		}

		log.Println("Пользователь ввел: ", input)

		_, err = file.WriteString(input + "\n")
		if err != nil {
			log.Println("Ошибка при записи текста: ", err)
			return
		}

		fmt.Print("Введите текст: ")
	}
	if err := scanner.Err(); err != nil {
		log.Println("Ошика при чтении ввода: ", err)
		return
	}

	log.Printf("Файл успешно создан на рабочем столе %s\n", desktopPath)

	log.Println("Закончил работу...\n")

}
