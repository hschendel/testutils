package testutils

import (
	"io/ioutil"
	"testing"
)

func TestOpenTestDataFile(t *testing.T) {
	file := OpenTestDataFile(t, "test.txt")
	if file == nil {
		return
	}
	content, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		t.Fatal(readErr)
		return
	}
	if len(content) < 4 || string(content[0:4]) != "Blah" {
		t.Error("Expected \"Blah\", but found: ", string(content))
	}
}