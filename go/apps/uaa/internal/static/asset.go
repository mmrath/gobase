package static

import (
	"io/ioutil"
	"path"
)

func AssetFunc(prefix string) func(name string) ([]byte, error) {
	fs := HTTP
	return func(name string) ([]byte, error) {
		f, err := fs.Open(path.Join(prefix, name))
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
	}
}

func AssetNamesFunc(prefix string) (func() []string, error) {
	files, err := WalkDirs(prefix, true)
	if err != nil {
		return nil, err
	}
	return func() []string {
		return files
	}, nil
}
