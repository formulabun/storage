package storage

import (
	"path"
	"testing"
)

func TestUrlIsOnlyFilename(t *testing.T) {
	testData := []struct {
		path           string
		isOnlyFilename bool
	}{
		{"", false},
		{"/", false},
		{"/filename", true},
		{"/filename/", false},
		{"/filename/checksum", false},
		{"/filename/checksum/", false},
		{"/filename/checksum/filename", false},
		{"/filename/checksum/dir/", false},
	}

	for _, td := range testData {
		res := urlIsOnlyFilename(td.path)
		if res != td.isOnlyFilename {
			clean := path.Clean(td.path)
			left, right := path.Split(clean)
			t.Logf("path: %#v, cleaned: %#v, split: %#v, %#v", td.path, clean, left, right)
			t.Errorf("urlIsOnlyFilename(%#v) == %#v but expected %#v", td.path, res, td.isOnlyFilename)
		}
	}
}
