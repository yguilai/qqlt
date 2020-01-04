package qqlt

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewDefaultServer()

	go s.DefaultRun()

	for update := range s.Updates {
		fmt.Println(update)
	}
}
