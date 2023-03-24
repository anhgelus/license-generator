package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Scan(s *string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sc := scanner.Text()
		*s = sc
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}
	return nil
}
