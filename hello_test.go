package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Response from getName is unexpected value")
	}
}
