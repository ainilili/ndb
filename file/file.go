package file

import "os"

type File struct {
	path string
	file *os.File
}

func New(path string, flag int) (*File, error) {
	f, err := os.OpenFile(path, flag, os.FileMode(0744))
	if err != nil {
		return nil, err
	}
	return &File{
		path: path,
		file: f,
	}, nil
}

func (f *File) Write(b []byte) (int, error) {
	return f.file.Write(b)
}

func (f *File) WriteAt(b []byte, off int64) (int, error) {
	return f.file.WriteAt(b, off)
}

func (f *File) Read(b []byte) (int, error) {
	return f.file.Read(b)
}

func (f *File) ReadAt(b []byte, off int64) (int, error) {
	return f.file.ReadAt(b, off)
}
