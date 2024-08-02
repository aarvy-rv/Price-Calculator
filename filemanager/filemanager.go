package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error in opening file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("error in reading line")
	}

	file.Close()
	return lines, nil
}
