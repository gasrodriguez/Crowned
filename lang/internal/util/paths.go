package util

import (
	"io/fs"
	"os"
	"path/filepath"
)

type paths struct {
	filters    []string
	extensions *Set
	files      *Set
}

func (o *paths) match(path string) bool {
	for i := range o.filters {
		match, err := filepath.Match(o.filters[i], path)
		CheckError(err)
		if match {
			return true
		}
	}
	return false
}

func (o *paths) walkDirFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() {
		if o.match(path) {
			return filepath.SkipDir
		} else {
			return nil
		}
	}
	if o.extensions.Contains(filepath.Ext(path)) {
		o.files.Add(path)
	}
	return nil
}

func FileSet(root string, extensions *Set, filters []string) *Set {
	p := paths{}
	for i := range filters {
		p.filters = append(p.filters, filepath.FromSlash(filters[i]))
	}
	p.extensions = extensions
	p.files = NewSet()

	err := os.Chdir(root)
	CheckError(err)
	err = filepath.WalkDir(".", p.walkDirFunc)
	CheckError(err)
	return p.files
}

func DirSet(files *Set) *Set {
	dirs := NewSet()
	for _, path := range files.Keys() {
		dirs.Add(filepath.Dir(path))
	}
	return dirs
}
