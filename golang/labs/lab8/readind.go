package lab8

import (
	"fmt"
	"os"
)

// ReadFile читает содержимое файла.
func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return string(data), nil
}