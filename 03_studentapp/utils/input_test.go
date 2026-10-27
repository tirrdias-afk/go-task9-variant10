package utils

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestInputRetriesAndEOF(t *testing.T) {
	oldScanner, oldErr := scanner, inputErr
	t.Cleanup(func() { scanner, inputErr = oldScanner, oldErr })
	scanner = bufio.NewScanner(strings.NewReader("abc\n4\n  Қазақ аты  \n"))
	inputErr = nil
	if ReadInt("") != 4 {
		t.Fatal("integer retry failed")
	}
	if ReadString("") != "Қазақ аты" {
		t.Fatal("unicode input failed")
	}
	if ReadInt("") != 0 || Err() != io.EOF {
		t.Fatal("EOF not handled")
	}
}
