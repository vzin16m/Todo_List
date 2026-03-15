package views

import (
	"bufio"
	"strings"

	"github.com/google/uuid"
)

func ReadInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ReadUUID(scanner *bufio.Scanner) (uuid.UUID, error) {
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())
	return uuid.Parse(text)
}
