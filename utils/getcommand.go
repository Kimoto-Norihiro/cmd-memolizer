package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

var path = filepath.Join(".", "data", "commands.txt")

func GetPath() string {
	return path
}

func GetCommands() ([]string, error) {
	var commands []string

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, string(scanner.Text()))
	}

	return commands, nil
}


