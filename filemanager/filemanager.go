package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func NewFileManger(inputFile, outputFile string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFile,
		OutputFilePath: outputFile,
	}
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteToJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode file")
	}

	return nil
}
