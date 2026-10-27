package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var inputErr error

// Err енгізу ағыны аяқталғанын не оқу қатесін анықтауға мүмкіндік береді.
func Err() error { return inputErr }

func ReadString(prompt string) string {
	if inputErr != nil {
		return ""
	}
	fmt.Print(prompt)
	if !scanner.Scan() {
		inputErr = scanner.Err()
		if inputErr == nil {
			inputErr = io.EOF
		}
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}

func ReadInt(prompt string) int {
	for {
		text := ReadString(prompt)
		if inputErr != nil {
			return 0
		}
		value, err := strconv.Atoi(text)
		if err == nil {
			return value
		}
		fmt.Println("Қате: бүтін сан енгізіңіз.")
	}
}
