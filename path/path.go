package path

import "github.com/rantool-team/go-utils/file"

type Path struct {
	Path string
}

func GenPath(path string) Path {
	p := Path{
		Path: path,
	}
	return p
}

func (p Path) IsFile() bool {
	return file.IsFile(p.Path)
}

func (p Path) IsFolder() bool {
	return !p.IsFile()
}
