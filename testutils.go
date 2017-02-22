package testutils

import (
	"path/filepath"
	"os"
	"testing"
)

// Opens a file from the package's testdata directory.
// If that does not work, does a t.Fatalln() with the error and returns nil.
func OpenTestDataFile(t *testing.T, fileName string) *os.File {
	workingDir, getwdErr := os.Getwd()
	if getwdErr != nil {
		t.Fatal(getwdErr)
		return nil
	}
	testFilePath := filepath.Join(workingDir, "testdata", fileName)
	file, openErr := os.Open(testFilePath)
	if openErr != nil {
		t.Fatal(openErr)
		return nil
	} else {
		return file
	}
}

func Equal(t *testing.T, expected interface{}, found interface{}, context string) {
	if expected != found {
		t.Errorf("%s: expected %v , found %v", context, expected, found)
	}
}
