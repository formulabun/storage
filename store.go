package storage

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

const rootName = "kartfilecache"

var root fs.FS

var dirmode = fs.ModeDir | 0755

func rootPath() string {
	cache, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	return path.Join(cache, rootName)
}

func init() {
	err := os.Mkdir(rootPath(), dirmode)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		panic(err)
	}
	root = os.DirFS(rootPath())
}

func validate(item Storable) error {
	k := item.ToKey()
	if k.Scheme != "" {
		return errors.New("Scheme must be empty")
	}
	if k.Opaque != "" {
		return errors.New("Opaque must be empty")
	}
	if k.User != nil {
		return errors.New("User must be empty")
	}
	if k.Host != "" {
		return errors.New("Host must be empty")
	}
	if k.RawQuery != "" {
		return errors.New("Query must be empty")
	}
	if k.Fragment != "" {
		return errors.New("Fragment must be empty")
	}
	if !fs.ValidPath(path.Dir(k.Path[1:])) {
		return errors.New("Invalid path")
	}
	return nil
}

func Store(item Storable, data io.Reader) error {
	err := validate(item)
	key := item.ToKey()
	if err != nil {
		return fmt.Errorf("Invalid storage url: %s", err)
	}

	p := path.Join(rootPath(), key.Path)
	err = os.MkdirAll(path.Dir(p), dirmode)
	if err != nil {
		return err
	}
	outFile, err := os.Create(p)
	if err != nil {
		return err
	}

	_, err = io.Copy(outFile, data)

	return err
}

func Get(item Storable) (io.ReadCloser, error) {
	err := validate(item)
	if err != nil {
		return nil, fmt.Errorf("Invalid storage url: %s", err)
	}
	key := item.ToKey()

	return root.Open(key.Path[1:])
}

func Has(item Storable) bool {
	if err := validate(item); err != nil {
		return false
	}

	p := path.Join(rootPath(), item.ToKey().Path)
	_, err := os.Stat(p)
	if err != nil {
		return false
	}
	return true
}
