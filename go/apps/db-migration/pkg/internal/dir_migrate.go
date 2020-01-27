package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	nurl "net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/golang-migrate/migrate/v4/source"
)

// Regex matches the following pattern:
//  123_name.up.ext
//  123_name.down.ext
var Regex = regexp.MustCompile(`^([0-9]+)_(.*)`)

// Parse returns Migration for matching Regex pattern.
func Parse(raw string) (*uint, *string, error) {
	m := Regex.FindStringSubmatch(raw)
	if len(m) == 3 {
		versionUint64, err := strconv.ParseUint(m[1], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		version := uint(versionUint64)
		return &version, &m[2], nil
	}
	return nil, nil, source.ErrParse
}

func init() {
	source.Register("dir", &File{})
}

type File struct {
	url        string
	path       string
	migrations *source.Migrations
}

func (f *File) Open(url string) (source.Driver, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}

	// concat host and path to restore full path
	// host might be `.`
	p := u.Host + u.Path

	if len(p) == 0 {
		// default to current directory if no path
		var wd string
		wd, err = os.Getwd()
		if err != nil {
			return nil, err
		}
		p = wd

	} else if p[0:1] == "." || p[0:1] != "/" {
		// make path absolute if relative
		var abs string
		abs, err = filepath.Abs(p)
		if err != nil {
			return nil, err
		}
		p = abs
	}

	// scan directory
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	nf := &File{
		url:        url,
		path:       p,
		migrations: source.NewMigrations(),
	}

	for _, fi := range files {
		if fi.IsDir() {
			v, id, err := Parse(fi.Name())
			if err != nil {
				continue // ignore files that we can't parse
			}

			//Up file exists
			up := &source.Migration{
				Version:    *v,
				Identifier: *id,
				Direction:  source.Up,
				Raw:        filepath.Join(fi.Name(), "up.sql"),
			}

			down := &source.Migration{
				Version:    *v,
				Identifier: *id,
				Direction:  source.Down,
				Raw:        filepath.Join(fi.Name(), "down.sql"),
			}

			err = nf.appendMigration(up)
			if err != nil {
				return nil, err
			}
			err = nf.appendMigration(down)
			if err != nil {
				return nil, err
			}

		}
	}
	return nf, nil
}

func (f *File) appendMigration(m *source.Migration) error {
	if stat, err := os.Stat(path.Join(f.path, m.Raw)); err == nil && !stat.IsDir() {
		if !f.migrations.Append(m) {
			return fmt.Errorf("unable to parse file %v", m)
		}
	}
	return nil
}

func (f *File) Close() error {
	// nothing do to here
	return nil
}

func (f *File) First() (uint, error) {
	v, ok := f.migrations.First()
	if !ok {
		return 0, &os.PathError{Op: "first", Path: f.path, Err: os.ErrNotExist}
	}
	return v, nil

}

func (f *File) Prev(version uint) (uint, error) {
	v, ok := f.migrations.Prev(version)

	if !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("prev for version %v", version), Path: f.path, Err: os.ErrNotExist}
	}

	return v, nil
}

func (f *File) Next(version uint) (uint, error) {
	v, ok := f.migrations.Next(version)
	if !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("next for version %v", version), Path: f.path, Err: os.ErrNotExist}
	}
	return v, nil
}

func (f *File) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := f.migrations.Up(version); ok {
		r, err := os.Open(path.Join(f.path, m.Raw))
		if err != nil {
			return nil, "", err
		}
		return r, m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: f.path, Err: os.ErrNotExist}
}

func (f *File) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := f.migrations.Down(version); ok {
		r, err := os.Open(path.Join(f.path, m.Raw))
		if err != nil {
			return nil, "", err
		}
		return r, m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: f.path, Err: os.ErrNotExist}
}
