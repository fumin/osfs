// Package osfs exposes the operating system's file system as an io/fs.ReadDirFS and io/fs.StatFS.
package osfs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// DirFS is a operating's file system.
type DirFS struct {
	Dir string
}

// Open opens a file.
func (fsys DirFS) Open(name string) (fs.File, error) {
	fpath := fsys.path(name)
	f, err := os.Open(fpath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s", fpath))
	}
	return f, nil
}

// ReadDir reads a directory..
func (fsys DirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	fpath := fsys.path(name)
	info, err := os.ReadDir(fpath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s", fpath))
	}
	return info, nil
}

// Stat returns the information of a file.
func (fsys DirFS) Stat(name string) (fs.FileInfo, error) {
	fpath := fsys.path(name)
	info, err := os.Stat(fpath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s", fpath))
	}
	return info, nil
}

func (fsys DirFS) path(name string) string {
	return filepath.Join(fsys.Dir, filepath.FromSlash(name))
}
