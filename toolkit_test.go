package toolkit

import (
	"runtime"
	"testing"
)

func TestOpenFolder(t *testing.T) {
	path := "/Users/zwj/JzTranslator/plugins"
	if runtime.GOOS == "windows" {
		path = "/xxxxx"
	}
	err := OpenFolder(path)
	if err != nil {
		t.Error(err)
	}
}
