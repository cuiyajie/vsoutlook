package utils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

func StringFallback(s string, fallback string) string {
	if len(s) > 0 {
		return s
	}
	return fallback
}

func Md5(strings ...string) string {
	hasher := md5.New()
	for _, s := range strings {
		hasher.Write([]byte(s))
	}
	h := hasher.Sum(nil)
	return hex.EncodeToString(h)
}

func ValidOneLineString(s string) string {
	return RemoveNonPrintable(strings.TrimSpace(s))
}

func ValidMultiLineString(s string) string {
	return RemoveNonPrintableExceptNewLine(strings.TrimSpace(s))
}

func RemoveNonPrintable(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
}

func isPrintOrNewline(r rune) bool {
	return unicode.IsPrint(r) || r == '\n'
}

func RemoveNonPrintableExceptNewLine(s string) string {
	return strings.Map(func(r rune) rune {
		if isPrintOrNewline(r) {
			return r
		}
		return -1
	}, s)
}

func FormatEmail(email string) string {
	return strings.ToLower(RemoveNonPrintable(strings.TrimSpace(email)))
}

var word = regexp.MustCompile(`^\w+$`)

func IsWord(s string) bool {
	return word.MatchString(s)
}

func String(s any) string {
	if s1, ok := s.(string); ok {
		return s1
	}
	return ""
}

func Collect[T any](rows []T, call func(T) string) []string {
	var result = make([]string, 0, len(rows))
	for _, row := range rows {
		val := call(row)
		if !slices.Contains(result, val) {
			result = append(result, val)
		}
	}
	return result
}

func ToValidUTF8(s string) string {
	s = strings.ToValidUTF8(s, "")
	s = strings.ReplaceAll(s, "\x00", "")
	return s
}
