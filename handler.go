package storage

import (
	"log"
	"net/http"
	"os"
	"path"
)

type StoreHandler struct {
}

// Serve a file from the store. The URL path of the request must be equal to a URL produced by a storable
// URLs with only a filename are served with the a file of that filename
func (h StoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := handlePath(r.URL.Path)
	if p == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path.Join(rootPath(), p))
}

func handlePath(url string) string {
	if !urlIsOnlyFilename(url) {
		return url
	}

	// replace /filename with /filename/checksum/filename
	pathToDir := path.Join(rootPath(), url)
	entries, err := os.ReadDir(pathToDir)

	if err != nil || len(entries) == 0 {
		log.Println(err)
		return ""
	}

	pathToFile := path.Join(url, entries[0].Name(), url)
	return pathToFile
}

func urlIsOnlyFilename(urlPath string) bool {
	d, p := path.Split(path.Clean(urlPath))
	return (d == "/") != (p == "")
}
