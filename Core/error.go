package Core

import (
	"fmt"
	"os"
)

func ReportError(errType string, msg string) string {
	fmt.Printf("%sError: %s\n", errType, msg)
	os.Exit(1)
	return "Exit code 1"
}
