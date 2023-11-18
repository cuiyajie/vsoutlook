package utils

import (
	"log"
	"testing"
)

func TestRemove(t *testing.T) {
	var a = []string{"a", "b"}
	var b = "b"
	Remove(a, b)
	log.Printf("a %#v\n", a)
	log.Fatal("show log")
}
