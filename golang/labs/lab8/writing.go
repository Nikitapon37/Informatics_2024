package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AppendDataToFile(fileName string) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите строку для записи (или 'Нет' для завершения):")
		newLine, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("ошибка чтения ввода: %w", err)
		}

		newLine = strings.TrimSpace(newLine)
		if strings.EqualFold(newLine, "Нет") {
			fmt.Println("Запись завершена.")
			break
		}

		if err := writeToFile(fileName, newLine); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}

	return nil
}

func writeToFile(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(text + "\n"); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return nil
}
