package storage

import "testing"

func TestUrlIsOnlyFilename(t *testing.T) {
	testData := []struct {
		path           string
		isOnlyFilename bool
	}{
		{"", false},
		{"/", false},
		{"/filename", true},
		{"/filename/", true},
		{"/filename/checksum", false},
		{"/filename/checksum/", false},
		{"/filename/checksum/filename", false},
		{"/filename/checksum/dir/", false},
	}

	for _, td := range testData {
		res := urlIsOnlyFilename(td.path)
		if res != td.isOnlyFilename {
			t.Errorf("urlIsOnlyFilename(%#v) == %#v but expected %#v", td.path, res, td.isOnlyFilename)
		}
	}
}
