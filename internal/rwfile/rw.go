package rwfile

import (
	"fmt"
	"os"
)

func Read(filename string) ([]byte, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return data, nil
}

func OverWrite(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to over write file %s: %w", filename, err)
	}
	return nil
}
