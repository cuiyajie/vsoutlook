package site

import (
	"log"
	"os"
	"strings"
)

func short(s string, n int) string {
	if len(s) > n {
		return strings.TrimSpace(s[:n])
	}
	return strings.TrimSpace(s)
}
func appendLog(file, message string) {
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(message + "\n"); err != nil {
		log.Println(err)
	}
}
