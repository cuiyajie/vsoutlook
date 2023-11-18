package db

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type FilteredWriter struct {
	writer    io.Writer
	filterStr string
}

func NewFilteredWriter(w io.Writer, filterStr string) *FilteredWriter {
	return &FilteredWriter{writer: w, filterStr: filterStr}
}

func (fw *FilteredWriter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, fw.filterStr) {
			fw.writer.Write([]byte(line + "\n"))
		}
	}

	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return len(p), nil
}

type StrippedLogger struct {
	Logger *log.Logger
}

func (s *StrippedLogger) Printf(format string, v ...any) {
	if len(v) >= 4 {
		s.Logger.Println(v[3])
	} else {
		s.Logger.Printf(format, v...)
	}
}
