package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFilePath(scanner *bufio.Scanner) string {
	fmt.Println("Введите название файла:")
	for scanner.Scan() {
		filePath := scanner.Text()
		if filePath != "" {
			return filePath
		}
		fmt.Println("Имя файла не может быть пустым. Попробуйте снова:")
	}
	return ""
}

func createFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Println("Файл успешно создан:", filePath)
	return file, nil
}

func writeFile(file *os.File, scanner *bufio.Scanner) error {
	fmt.Println("Введите текст для записи в файл. Для завершения записи введите пустую строку и нажмите Enter:")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		if _, err := file.WriteString(input + "\n"); err != nil {
			return err
		}
	}
	fmt.Println("Данные успешно записаны в файл.")
	return nil
}

func readFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("\nСодержимое файла:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func searchInFile(filePath, searchText string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			fmt.Printf("Найдено совпадение на строке %d: %s\n", lineNumber, line)
			found = true
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if !found {
		fmt.Println("Совпадений не найдено.")
	}
	return nil
}

func RunLab8() {
	scanner := bufio.NewScanner(os.Stdin)

	filePath := getFilePath(scanner)
	if filePath == "" {
		fmt.Println("Ошибка: имя файла не задано.")
		return
	}

	file, err := createFile(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	if err := writeFile(file, scanner); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	if err := readFile(filePath); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println("\nВведите текст для поиска в файле:")
	var searchText string
	if scanner.Scan() {
		searchText = scanner.Text()
	}

	if searchText != "" {
		if err := searchInFile(filePath, searchText); err != nil {
			fmt.Println("Ошибка при поиске в файле:", err)
		}
	} else {
		fmt.Println("Поиск не выполнен: текст для поиска не был введен.")
	}
}
