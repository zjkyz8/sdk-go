package http

import (
	"io"
	"os"
	"path/filepath"
)

type FileItem struct {
	Src  io.Reader
	Name string
}

func (fi FileItem) Close() error {
	c, b := fi.Src.(io.Closer)
	if b {
		return c.Close()
	}
	return nil
}

func (fi FileItem) GetName() string {
	if fi.Name == "" && fi.Src != nil {
		file, b := fi.Src.(*os.File)
		if b {
			return filepath.Base(file.Name())
		}
	}
	return fi.Name
}
