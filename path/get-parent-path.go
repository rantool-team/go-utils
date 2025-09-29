package path

import "path/filepath"

func (p Path) GetParentPath() Path {
	parentPath := filepath.Dir(p.Path)
	newPath := Path{
		Path: parentPath,
	}

	return newPath
}
