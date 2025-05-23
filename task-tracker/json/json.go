package json

import (
	"fmt"
	"os"
	"strings"
)

func CreateFile(filename string) error {
	filename = verifyJSONExtension(filename)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	return nil
}

func UpdateFile(filename string, data []byte) error {
	filename = verifyJSONExtension(filename)

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func LoadFile(filename string) ([]byte, error) {
	filename = verifyJSONExtension(filename)

	return os.ReadFile(filename)
}

func verifyJSONExtension(filename string) string {
	if !strings.HasSuffix(filename, ".json") {
		return filename + ".json"
	}

	return filename
}
