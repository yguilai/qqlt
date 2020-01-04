package qqlt

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	fmt.Println(NewMessage().At("111").Line().Text("asdaasd").ToString())
}
