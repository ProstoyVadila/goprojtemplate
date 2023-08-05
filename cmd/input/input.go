package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ProstoyVadila/goproj/internal/models"
)

func readInput(scanner *bufio.Scanner, previousMessage string) (string, error) {
	fmt.Print(previousMessage)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

// GetSetup tries to get information about project from CLI or from Input
func GetSetup() (*models.Setup, error) {

	scanner := bufio.NewScanner(os.Stdin)

	author, err := readInput(scanner, "Please, enter your name: ")
	if err != nil {
		return &models.Setup{}, err
	}
	packageName, err := readInput(scanner, "Please, enter your new project (package) name: ")
	if err != nil {
		return &models.Setup{}, err
	}
	description, err := readInput(scanner, "Please, add a description to your project: ")
	if err != nil {
		return &models.Setup{}, err
	}

	return models.NewInputData(packageName, author, description, make([]string, 0)), nil
}